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
		"cn-qingdao":                  dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-beijing":                  dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("bastionhost.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                dara.String("bastionhost.cn-huhehaote.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai":                 dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-heyuan":                   dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-guangzhou":                dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-chengdu":                  dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hongkong":                 dara.String("bastionhost.cn-hongkong.aliyuncs.com"),
		"ap-northeast-1":              dara.String("bastionhost.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1":              dara.String("bastionhost.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              dara.String("bastionhost.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              dara.String("bastionhost.ap-southeast-5.aliyuncs.com"),
		"us-east-1":                   dara.String("bastionhost.us-east-1.aliyuncs.com"),
		"us-west-1":                   dara.String("bastionhost.us-west-1.aliyuncs.com"),
		"eu-west-1":                   dara.String("bastionhost.eu-west-1.aliyuncs.com"),
		"eu-central-1":                dara.String("bastionhost.eu-central-1.aliyuncs.com"),
		"me-east-1":                   dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("yundun-bastionhost.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("yundun-bastionhost.aliyuncs.com"),
		"ap-south-1":                  dara.String("yundun-bastionhost.aliyuncs.com"),
		"ap-southeast-2":              dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-edge-1":                   dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-fujian":                   dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-wuhan":                    dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-yushanfang":               dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("yundun-bastionhost.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("yundun-bastionhost.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("yundun-bastionhost.aliyuncs.com"),
		"cn-zhengzhou-jva":            dara.String("yundun-bastionhost.aliyuncs.com"),
		"ap-northeast-2":              dara.String("bastionhost.ap-northeast-2.aliyuncs.com"),
		"ap-southeast-6":              dara.String("bastionhost.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-7":              dara.String("bastionhost.ap-southeast-7.aliyuncs.com"),
		"na-south-1":                  dara.String("bastionhost.na-south-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("yundun-bastionhost"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// If an O&M engineer runs a command that is configured in a command approval control policy, the administrator receives an approval request for the command in the Bastionhost console. The command can be run only after the administrator approves the request. If the administrator rejects the request, the command cannot be run.
//
// Description:
//
// Approves a command execution request from an O&M engineer.
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AcceptApproveCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptApproveCommandResponse
func (client *Client) AcceptApproveCommandWithOptions(request *AcceptApproveCommandRequest, runtime *dara.RuntimeOptions) (_result *AcceptApproveCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandId) {
		query["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptApproveCommand"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If an O&M engineer runs a command that is configured in a command approval control policy, the administrator receives an approval request for the command in the Bastionhost console. The command can be run only after the administrator approves the request. If the administrator rejects the request, the command cannot be run.
//
// Description:
//
// Approves a command execution request from an O&M engineer.
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AcceptApproveCommandRequest
//
// @return AcceptApproveCommandResponse
func (client *Client) AcceptApproveCommand(request *AcceptApproveCommandRequest) (_result *AcceptApproveCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.AcceptApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If an administrator enables O&M approval in a control policy, O&M engineers must submit an O&M request and obtain administrator approval before they can log on to assets.
//
// Description:
//
// Approves an O&M request submitted by an O&M engineer.
//
// The China site Chinese version of this China Chinese China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China.
//
// @param request - AcceptOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptOperationTicketResponse
func (client *Client) AcceptOperationTicketWithOptions(request *AcceptOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *AcceptOperationTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EffectCount) {
		query["EffectCount"] = request.EffectCount
	}

	if !dara.IsNil(request.EffectEndTime) {
		query["EffectEndTime"] = request.EffectEndTime
	}

	if !dara.IsNil(request.EffectStartTime) {
		query["EffectStartTime"] = request.EffectStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationTicketId) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If an administrator enables O&M approval in a control policy, O&M engineers must submit an O&M request and obtain administrator approval before they can log on to assets.
//
// Description:
//
// Approves an O&M request submitted by an O&M engineer.
//
// The China site Chinese version of this China Chinese China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China China.
//
// @param request - AcceptOperationTicketRequest
//
// @return AcceptOperationTicketResponse
func (client *Client) AcceptOperationTicket(request *AcceptOperationTicketRequest) (_result *AcceptOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.AcceptOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple database instances to a specified asset group.
//
// @param request - AddDatabasesToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatabasesToGroupResponse
func (client *Client) AddDatabasesToGroupWithOptions(request *AddDatabasesToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDatabasesToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDatabasesToGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDatabasesToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple database instances to a specified asset group.
//
// @param request - AddDatabasesToGroupRequest
//
// @return AddDatabasesToGroupResponse
func (client *Client) AddDatabasesToGroup(request *AddDatabasesToGroupRequest) (_result *AddDatabasesToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDatabasesToGroupResponse{}
	_body, _err := client.AddDatabasesToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple hosts to a specified asset group in a batch.
//
// Description:
//
// Adds multiple hosts to a specified host group in a batch. By adding multiple hosts to a host group, you can centrally manage these hosts and grant permissions in a batch.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AddHostsToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHostsToGroupResponse
func (client *Client) AddHostsToGroupWithOptions(request *AddHostsToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddHostsToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddHostsToGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple hosts to a specified asset group in a batch.
//
// Description:
//
// Adds multiple hosts to a specified host group in a batch. By adding multiple hosts to a host group, you can centrally manage these hosts and grant permissions in a batch.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AddHostsToGroupRequest
//
// @return AddHostsToGroupResponse
func (client *Client) AddHostsToGroup(request *AddHostsToGroupRequest) (_result *AddHostsToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.AddHostsToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds member accounts in a resource directory.
//
// @param request - AddInstanceRdMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceRdMemberResponse
func (client *Client) AddInstanceRdMemberWithOptions(request *AddInstanceRdMemberRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceRdMemberResponse, _err error) {
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

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInstanceRdMember"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInstanceRdMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds member accounts in a resource directory.
//
// @param request - AddInstanceRdMemberRequest
//
// @return AddInstanceRdMemberResponse
func (client *Client) AddInstanceRdMember(request *AddInstanceRdMemberRequest) (_result *AddInstanceRdMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddInstanceRdMemberResponse{}
	_body, _err := client.AddInstanceRdMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds users to a user group in a batch.
//
// Description:
//
// ### Operation description
//
// This operation adds users to a user group in a batch. After you create a user group by calling the [CreateUserGroup](https://help.aliyun.com/document_detail/204596.html) operation, you can call this operation to add multiple users to the user group for batch authorization and management.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AddUsersToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithOptions(request *AddUsersToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUsersToGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to a user group in a batch.
//
// Description:
//
// ### Operation description
//
// This operation adds users to a user group in a batch. After you create a user group by calling the [CreateUserGroup](https://help.aliyun.com/document_detail/204596.html) operation, you can call this operation to add multiple users to the user group for batch authorization and management.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AddUsersToGroupRequest
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroup(request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Approves an O&M task execution request.
//
// @param request - AllowOperationTaskApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllowOperationTaskApprovalResponse
func (client *Client) AllowOperationTaskApprovalWithOptions(request *AllowOperationTaskApprovalRequest, runtime *dara.RuntimeOptions) (_result *AllowOperationTaskApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveId) {
		query["ApproveId"] = request.ApproveId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllowOperationTaskApproval"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllowOperationTaskApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Approves an O&M task execution request.
//
// @param request - AllowOperationTaskApprovalRequest
//
// @return AllowOperationTaskApprovalResponse
func (client *Client) AllowOperationTaskApproval(request *AllowOperationTaskApprovalRequest) (_result *AllowOperationTaskApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllowOperationTaskApprovalResponse{}
	_body, _err := client.AllowOperationTaskApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDatabaseAccountsToUserResponse
func (client *Client) AttachDatabaseAccountsToUserWithOptions(request *AttachDatabaseAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachDatabaseAccountsToUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDatabaseAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDatabaseAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserRequest
//
// @return AttachDatabaseAccountsToUserResponse
func (client *Client) AttachDatabaseAccountsToUser(request *AttachDatabaseAccountsToUserRequest) (_result *AttachDatabaseAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDatabaseAccountsToUserResponse{}
	_body, _err := client.AttachDatabaseAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a user group permissions on databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDatabaseAccountsToUserGroupResponse
func (client *Client) AttachDatabaseAccountsToUserGroupWithOptions(request *AttachDatabaseAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachDatabaseAccountsToUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDatabaseAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDatabaseAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a user group permissions on databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserGroupRequest
//
// @return AttachDatabaseAccountsToUserGroupResponse
func (client *Client) AttachDatabaseAccountsToUserGroup(request *AttachDatabaseAccountsToUserGroupRequest) (_result *AttachDatabaseAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDatabaseAccountsToUserGroupResponse{}
	_body, _err := client.AttachDatabaseAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates host accounts with a host shared key.
//
// @param request - AttachHostAccountsToHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToHostShareKeyResponse
func (client *Client) AttachHostAccountsToHostShareKeyWithOptions(request *AttachHostAccountsToHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountIds) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates host accounts with a host shared key.
//
// @param request - AttachHostAccountsToHostShareKeyRequest
//
// @return AttachHostAccountsToHostShareKeyResponse
func (client *Client) AttachHostAccountsToHostShareKey(request *AttachHostAccountsToHostShareKeyRequest) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.AttachHostAccountsToHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes assets for a user. After the authorization, the user can use Bastionhost to perform O&M operations on the authorized assets.
//
// @param request - AttachHostAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToUserResponse
func (client *Client) AttachHostAccountsToUserWithOptions(request *AttachHostAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes assets for a user. After the authorization, the user can use Bastionhost to perform O&M operations on the authorized assets.
//
// @param request - AttachHostAccountsToUserRequest
//
// @return AttachHostAccountsToUserResponse
func (client *Client) AttachHostAccountsToUser(request *AttachHostAccountsToUserRequest) (_result *AttachHostAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.AttachHostAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a user group permissions on hosts and host accounts.
//
// Description:
//
// After you grant a user group permissions on hosts and host accounts, all users in the user group can access the authorized hosts.
//
// @param request - AttachHostAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToUserGroupResponse
func (client *Client) AttachHostAccountsToUserGroupWithOptions(request *AttachHostAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a user group permissions on hosts and host accounts.
//
// Description:
//
// After you grant a user group permissions on hosts and host accounts, all users in the user group can access the authorized hosts.
//
// @param request - AttachHostAccountsToUserGroupRequest
//
// @return AttachHostAccountsToUserGroupResponse
func (client *Client) AttachHostAccountsToUserGroup(request *AttachHostAccountsToUserGroupRequest) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on host groups and host accounts to a user.
//
// @param request - AttachHostGroupAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostGroupAccountsToUserResponse
func (client *Client) AttachHostGroupAccountsToUserWithOptions(request *AttachHostGroupAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostGroupAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on host groups and host accounts to a user.
//
// @param request - AttachHostGroupAccountsToUserRequest
//
// @return AttachHostGroupAccountsToUserResponse
func (client *Client) AttachHostGroupAccountsToUser(request *AttachHostGroupAccountsToUserRequest) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on host groups and host accounts to a user group.
//
// @param request - AttachHostGroupAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostGroupAccountsToUserGroupResponse
func (client *Client) AttachHostGroupAccountsToUserGroupWithOptions(request *AttachHostGroupAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostGroupAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on host groups and host accounts to a user group.
//
// @param request - AttachHostGroupAccountsToUserGroupRequest
//
// @return AttachHostGroupAccountsToUserGroupResponse
func (client *Client) AttachHostGroupAccountsToUserGroup(request *AttachHostGroupAccountsToUserGroupRequest) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates members with a project.
//
// @param request - AttachMembersToProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachMembersToProjectResponse
func (client *Client) AttachMembersToProjectWithOptions(request *AttachMembersToProjectRequest, runtime *dara.RuntimeOptions) (_result *AttachMembersToProjectResponse, _err error) {
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

	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachMembersToProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachMembersToProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates members with a project.
//
// @param request - AttachMembersToProjectRequest
//
// @return AttachMembersToProjectResponse
func (client *Client) AttachMembersToProject(request *AttachMembersToProjectRequest) (_result *AttachMembersToProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachMembersToProjectResponse{}
	_body, _err := client.AttachMembersToProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates project authorization administrators.
//
// @param request - AttachProjectManagersToProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachProjectManagersToProjectAuthorizationResponse
func (client *Client) AttachProjectManagersToProjectAuthorizationWithOptions(request *AttachProjectManagersToProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *AttachProjectManagersToProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Principals) {
		query["Principals"] = request.Principals
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachProjectManagersToProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachProjectManagersToProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates project authorization administrators.
//
// @param request - AttachProjectManagersToProjectAuthorizationRequest
//
// @return AttachProjectManagersToProjectAuthorizationResponse
func (client *Client) AttachProjectManagersToProjectAuthorization(request *AttachProjectManagersToProjectAuthorizationRequest) (_result *AttachProjectManagersToProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachProjectManagersToProjectAuthorizationResponse{}
	_body, _err := client.AttachProjectManagersToProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the projects that can be managed in a project authorization. You can call this operation only when the ScopeType parameter returned by the GetProjectAuthorization operation for the corresponding project authorization is 2.
//
// @param request - AttachProjectsToProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachProjectsToProjectAuthorizationResponse
func (client *Client) AttachProjectsToProjectAuthorizationWithOptions(request *AttachProjectsToProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *AttachProjectsToProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachProjectsToProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachProjectsToProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the projects that can be managed in a project authorization. You can call this operation only when the ScopeType parameter returned by the GetProjectAuthorization operation for the corresponding project authorization is 2.
//
// @param request - AttachProjectsToProjectAuthorizationRequest
//
// @return AttachProjectsToProjectAuthorizationResponse
func (client *Client) AttachProjectsToProjectAuthorization(request *AttachProjectsToProjectAuthorizationRequest) (_result *AttachProjectsToProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachProjectsToProjectAuthorizationResponse{}
	_body, _err := client.AttachProjectsToProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates projects with the scope of an audit policy. You can invoke this operation only when the ScopeType parameter returned by the GetReviewRule operation for the corresponding audit policy is 2.
//
// @param request - AttachProjectsToReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachProjectsToReviewRuleResponse
func (client *Client) AttachProjectsToReviewRuleWithOptions(request *AttachProjectsToReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *AttachProjectsToReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachProjectsToReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachProjectsToReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates projects with the scope of an audit policy. You can invoke this operation only when the ScopeType parameter returned by the GetReviewRule operation for the corresponding audit policy is 2.
//
// @param request - AttachProjectsToReviewRuleRequest
//
// @return AttachProjectsToReviewRuleResponse
func (client *Client) AttachProjectsToReviewRule(request *AttachProjectsToReviewRuleRequest) (_result *AttachProjectsToReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachProjectsToReviewRuleResponse{}
	_body, _err := client.AttachProjectsToReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches auditors to an audit policy.
//
// @param request - AttachReviewersToReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachReviewersToReviewRuleResponse
func (client *Client) AttachReviewersToReviewRuleWithOptions(request *AttachReviewersToReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *AttachReviewersToReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.Principals) {
		query["Principals"] = request.Principals
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachReviewersToReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachReviewersToReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches auditors to an audit policy.
//
// @param request - AttachReviewersToReviewRuleRequest
//
// @return AttachReviewersToReviewRuleResponse
func (client *Client) AttachReviewersToReviewRule(request *AttachReviewersToReviewRuleRequest) (_result *AttachReviewersToReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachReviewersToReviewRuleResponse{}
	_body, _err := client.AttachReviewersToReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the security groups that control inbound and outbound network traffic for a Bastionhost instance.
//
// @param request - ConfigInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceSecurityGroupsResponse
func (client *Client) ConfigInstanceSecurityGroupsWithOptions(request *ConfigInstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedSecurityGroups) {
		query["AuthorizedSecurityGroups"] = request.AuthorizedSecurityGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceSecurityGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the security groups that control inbound and outbound network traffic for a Bastionhost instance.
//
// @param request - ConfigInstanceSecurityGroupsRequest
//
// @return ConfigInstanceSecurityGroupsResponse
func (client *Client) ConfigInstanceSecurityGroups(request *ConfigInstanceSecurityGroupsRequest) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.ConfigInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds specified public IP addresses to the whitelist of a Bastionhost instance after public network access is enabled to allow public network traffic to access the instance.
//
// Description:
//
// Configures the public IP address whitelist for a Bastionhost instance. By default, all public IP addresses are allowed to access the Bastionhost instance. To restrict public network access, call this operation to add specified public IP addresses to the whitelist and restrict public network traffic to the Bastionhost instance.
//
// The queries per second (QPS) limit for a single user is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - ConfigInstanceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceWhiteListResponse
func (client *Client) ConfigInstanceWhiteListWithOptions(request *ConfigInstanceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceWhiteListResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WhiteList) {
		query["WhiteList"] = request.WhiteList
	}

	if !dara.IsNil(request.WhiteListPolicies) {
		query["WhiteListPolicies"] = request.WhiteListPolicies
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceWhiteList"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds specified public IP addresses to the whitelist of a Bastionhost instance after public network access is enabled to allow public network traffic to access the instance.
//
// Description:
//
// Configures the public IP address whitelist for a Bastionhost instance. By default, all public IP addresses are allowed to access the Bastionhost instance. To restrict public network access, call this operation to add specified public IP addresses to the whitelist and restrict public network traffic to the Bastionhost instance.
//
// The queries per second (QPS) limit for a single user is 30 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - ConfigInstanceWhiteListRequest
//
// @return ConfigInstanceWhiteListResponse
func (client *Client) ConfigInstanceWhiteList(request *ConfigInstanceWhiteListRequest) (_result *ConfigInstanceWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.ConfigInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports database assets into a bastion host. Supported database types include MySQL, SQL Server, PostgreSQL, and Oracle for ApsaraDB RDS instances, MySQL, PostgreSQL, and PolarDB for PostgreSQL (Compatible with Oracle) for PolarDB clusters, and MySQL, SQL Server, PostgreSQL, and Oracle for self-managed databases.
//
// @param request - CreateDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabasePort) {
		query["DatabasePort"] = request.DatabasePort
	}

	if !dara.IsNil(request.DatabasePrivateAddress) {
		query["DatabasePrivateAddress"] = request.DatabasePrivateAddress
	}

	if !dara.IsNil(request.DatabasePublicAddress) {
		query["DatabasePublicAddress"] = request.DatabasePublicAddress
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceMemberId) {
		query["InstanceMemberId"] = request.InstanceMemberId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.PolarDBEndpointType) {
		query["PolarDBEndpointType"] = request.PolarDBEndpointType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceRegionId) {
		query["SourceInstanceRegionId"] = request.SourceInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports database assets into a bastion host. Supported database types include MySQL, SQL Server, PostgreSQL, and Oracle for ApsaraDB RDS instances, MySQL, PostgreSQL, and PolarDB for PostgreSQL (Compatible with Oracle) for PolarDB clusters, and MySQL, SQL Server, PostgreSQL, and Oracle for self-managed databases.
//
// @param request - CreateDatabaseRequest
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After a database is created, you can create a database account for it. After the account is created, O&M engineers can use the account to log on to and manage the database.
//
// @param request - CreateDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseAccountResponse
func (client *Client) CreateDatabaseAccountWithOptions(request *CreateDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginAttribute) {
		query["LoginAttribute"] = request.LoginAttribute
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After a database is created, you can create a database account for it. After the account is created, O&M engineers can use the account to log on to and manage the database.
//
// @param request - CreateDatabaseAccountRequest
//
// @return CreateDatabaseAccountResponse
func (client *Client) CreateDatabaseAccount(request *CreateDatabaseAccountRequest) (_result *CreateDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseAccountResponse{}
	_body, _err := client.CreateDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration backup export task. Only one configuration backup export task can run at a time for a Bastionhost instance.
//
// @param request - CreateExportConfigJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportConfigJobResponse
func (client *Client) CreateExportConfigJobWithOptions(request *CreateExportConfigJobRequest, runtime *dara.RuntimeOptions) (_result *CreateExportConfigJobResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExportConfigJob"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExportConfigJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration backup export task. Only one configuration backup export task can run at a time for a Bastionhost instance.
//
// @param request - CreateExportConfigJobRequest
//
// @return CreateExportConfigJobResponse
func (client *Client) CreateExportConfigJob(request *CreateExportConfigJobRequest) (_result *CreateExportConfigJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExportConfigJobResponse{}
	_body, _err := client.CreateExportConfigJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a host in a Bastionhost instance. Bastionhost supports O&M for hosts from different sources, including Alibaba Cloud ECS instances, on-premises IDC servers, and servers on other clouds. Before you perform O&M on a host through Bastionhost, you must first import the host into Bastionhost. You can call this operation to create a host that you want to manage in Bastionhost.
//
// @param request - CreateHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostResponse
func (client *Client) CreateHostWithOptions(request *CreateHostRequest, runtime *dara.RuntimeOptions) (_result *CreateHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HostPrivateAddress) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !dara.IsNil(request.HostPublicAddress) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceMemberId) {
		query["InstanceMemberId"] = request.InstanceMemberId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a host in a Bastionhost instance. Bastionhost supports O&M for hosts from different sources, including Alibaba Cloud ECS instances, on-premises IDC servers, and servers on other clouds. Before you perform O&M on a host through Bastionhost, you must first import the host into Bastionhost. You can call this operation to create a host that you want to manage in Bastionhost.
//
// @param request - CreateHostRequest
//
// @return CreateHostResponse
func (client *Client) CreateHost(request *CreateHostRequest) (_result *CreateHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostResponse{}
	_body, _err := client.CreateHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After you add a host to Bastionhost, you can create a host account for the host to manage the existing account of the host in Bastionhost. After a host account is created, O&M engineers can use the account to log on to the host through Bastionhost for O&M operations.
//
// @param request - CreateHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostAccountResponse
func (client *Client) CreateHostAccountWithOptions(request *CreateHostAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateHostAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RotationMode) {
		query["RotationMode"] = request.RotationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you add a host to Bastionhost, you can create a host account for the host to manage the existing account of the host in Bastionhost. After a host account is created, O&M engineers can use the account to log on to the host through Bastionhost for O&M operations.
//
// @param request - CreateHostAccountRequest
//
// @return CreateHostAccountResponse
func (client *Client) CreateHostAccount(request *CreateHostAccountRequest) (_result *CreateHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CreateHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can create different asset groups based on your business requirements, add assets of the same type to an asset group, and manage assets by category and perform batch operations.
//
// @param request - CreateHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostGroupResponse
func (client *Client) CreateHostGroupWithOptions(request *CreateHostGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create different asset groups based on your business requirements, add assets of the same type to an asset group, and manage assets by category and perform batch operations.
//
// @param request - CreateHostGroupRequest
//
// @return CreateHostGroupResponse
func (client *Client) CreateHostGroup(request *CreateHostGroupRequest) (_result *CreateHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CreateHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The shared key feature of Bastionhost allows you to store private keys used for host logon in Bastionhost. After a private key is stored, it can be shared across multiple host accounts to improve host account management efficiency.
//
// @param request - CreateHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostShareKeyResponse
func (client *Client) CreateHostShareKeyWithOptions(request *CreateHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyName) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The shared key feature of Bastionhost allows you to store private keys used for host logon in Bastionhost. After a private key is stored, it can be shared across multiple host accounts to improve host account management efficiency.
//
// @param request - CreateHostShareKeyRequest
//
// @return CreateHostShareKeyResponse
func (client *Client) CreateHostShareKey(request *CreateHostShareKeyRequest) (_result *CreateHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CreateHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// To centrally manage assets that are distributed across different network environments or that cannot communicate with the virtual private cloud (VPC) where Bastionhost is deployed, use the network domain feature of Bastionhost. You can configure a proxy server for these assets, create a network domain in Bastionhost, add the proxy server to the network domain, and then add the assets to the network domain. This way, you can manage the assets through Bastionhost.
//
// @param request - CreateNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkDomainResponse
func (client *Client) CreateNetworkDomainWithOptions(request *CreateNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Proxies) {
		query["Proxies"] = request.Proxies
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To centrally manage assets that are distributed across different network environments or that cannot communicate with the virtual private cloud (VPC) where Bastionhost is deployed, use the network domain feature of Bastionhost. You can configure a proxy server for these assets, create a network domain in Bastionhost, add the proxy server to the network domain, and then add the assets to the network domain. This way, you can manage the assets through Bastionhost.
//
// @param request - CreateNetworkDomainRequest
//
// @return CreateNetworkDomainResponse
func (client *Client) CreateNetworkDomain(request *CreateNetworkDomainRequest) (_result *CreateNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkDomainResponse{}
	_body, _err := client.CreateNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an O&M request. When an administrator enables O&M approval in a control policy, O&M engineers must create an O&M request and obtain administrator approval before performing O&M operations.
//
// @param request - CreateOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOperationTicketResponse
func (client *Client) CreateOperationTicketWithOptions(request *CreateOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateOperationTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveComment) {
		query["ApproveComment"] = request.ApproveComment
	}

	if !dara.IsNil(request.AssetAccountName) {
		query["AssetAccountName"] = request.AssetAccountName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.EffectEndTime) {
		query["EffectEndTime"] = request.EffectEndTime
	}

	if !dara.IsNil(request.EffectStartTime) {
		query["EffectStartTime"] = request.EffectStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsOneTimeEffect) {
		query["IsOneTimeEffect"] = request.IsOneTimeEffect
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an O&M request. When an administrator enables O&M approval in a control policy, O&M engineers must create an O&M request and obtain administrator approval before performing O&M operations.
//
// @param request - CreateOperationTicketRequest
//
// @return CreateOperationTicketResponse
func (client *Client) CreateOperationTicket(request *CreateOperationTicketRequest) (_result *CreateOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOperationTicketResponse{}
	_body, _err := client.CreateOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manages O&M behaviors by configuring command control, command approval, protocol control, and access control policies to effectively prevent users from executing high-risk commands or performing misoperations, ensuring O&M security.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages O&M behaviors by configuring command control, command approval, protocol control, and access control policies to effectively prevent users from executing high-risk commands or performing misoperations, ensuring O&M security.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// @param request - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project authorization.
//
// @param request - CreateProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectAuthorizationResponse
func (client *Client) CreateProjectAuthorizationWithOptions(request *CreateProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectAuthorizationResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project authorization.
//
// @param request - CreateProjectAuthorizationRequest
//
// @return CreateProjectAuthorizationResponse
func (client *Client) CreateProjectAuthorization(request *CreateProjectAuthorizationRequest) (_result *CreateProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectAuthorizationResponse{}
	_body, _err := client.CreateProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project notification recipient.
//
// @param request - CreateProjectNotifyReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectNotifyReceiverResponse
func (client *Client) CreateProjectNotifyReceiverWithOptions(request *CreateProjectNotifyReceiverRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectNotifyReceiverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectNotifyReceiver"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectNotifyReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project notification recipient.
//
// @param request - CreateProjectNotifyReceiverRequest
//
// @return CreateProjectNotifyReceiverResponse
func (client *Client) CreateProjectNotifyReceiver(request *CreateProjectNotifyReceiverRequest) (_result *CreateProjectNotifyReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectNotifyReceiverResponse{}
	_body, _err := client.CreateProjectNotifyReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an audit policy.
//
// @param request - CreateReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReviewRuleResponse
func (client *Client) CreateReviewRuleWithOptions(request *CreateReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an audit policy.
//
// @param request - CreateReviewRuleRequest
//
// @return CreateReviewRuleResponse
func (client *Client) CreateReviewRule(request *CreateReviewRuleRequest) (_result *CreateReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReviewRuleResponse{}
	_body, _err := client.CreateReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an authorization rule to grant multiple users batch access to assets and set the validity period for their access. The authorization rule feature improves the efficiency of managing users and assets, and controls the time period during which users can access assets.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an authorization rule to grant multiple users batch access to assets and set the validity period for their access. The authorization rule feature improves the efficiency of managing users and assets, and controls the time period during which users can access assets.
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a bastion host.
//
// Description:
//
// Adds local users, Resource Access Management (RAM) users, Active Directory (AD)-authenticated users, or Lightweight Directory Access Protocol (LDAP)-authenticated users to a bastion host. After adding a user, O\\&M engineers can log on to the bastion host with that account to perform O\\&M operations on authorized hosts.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LanguageStatus) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.MobileCountryCode) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !dara.IsNil(request.NeedResetPassword) {
		query["NeedResetPassword"] = request.NeedResetPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceUserId) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !dara.IsNil(request.TwoFactorStatus) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a bastion host.
//
// Description:
//
// Adds local users, Resource Access Management (RAM) users, Active Directory (AD)-authenticated users, or Lightweight Directory Access Protocol (LDAP)-authenticated users to a bastion host. After adding a user, O\\&M engineers can log on to the bastion host with that account to perform O\\&M operations on authorized hosts.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user group for a bastion host.
//
// Description:
//
// Creates a user group for a bastion host. Bastion host administrators can use this operation to create a user group and then call the [AddUsersToGroup](https://help.aliyun.com/document_detail/204600.html) operation to add users to the user group in batches. Adding users to the same user group enables unified authorization and management of bastion host users.
//
// The China site Chinese edition single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group for a bastion host.
//
// Description:
//
// Creates a user group for a bastion host. Bastion host administrators can use this operation to create a user group and then call the [AddUsersToGroup](https://help.aliyun.com/document_detail/204600.html) operation to add users to the user group in batches. Adding users to the same user group enables unified authorization and management of bastion host users.
//
// The China site Chinese edition single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public key for a bastion host user and hosts the public key in the bastion host. This way, O\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\&M engineers can use the private key that corresponds to the public key to log on to the bastion host from an O\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\&M client.
//
// Description:
//
// You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
//
// @param request - CreateUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPublicKeyResponse
func (client *Client) CreateUserPublicKeyWithOptions(request *CreateUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	if !dara.IsNil(request.PublicKeyName) {
		query["PublicKeyName"] = request.PublicKeyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public key for a bastion host user and hosts the public key in the bastion host. This way, O\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\&M engineers can use the private key that corresponds to the public key to log on to the bastion host from an O\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\&M client.
//
// Description:
//
// You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
//
// @param request - CreateUserPublicKeyRequest
//
// @return CreateUserPublicKeyResponse
func (client *Client) CreateUserPublicKey(request *CreateUserPublicKeyRequest) (_result *CreateUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CreateUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single database instance.
//
// @param request - DeleteDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single database instance.
//
// @param request - DeleteDatabaseRequest
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single database account.
//
// @param request - DeleteDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseAccountResponse
func (client *Client) DeleteDatabaseAccountWithOptions(request *DeleteDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single database account.
//
// @param request - DeleteDatabaseAccountRequest
//
// @return DeleteDatabaseAccountResponse
func (client *Client) DeleteDatabaseAccount(request *DeleteDatabaseAccountRequest) (_result *DeleteDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseAccountResponse{}
	_body, _err := client.DeleteDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single host.
//
// @param request - DeleteHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostResponse
func (client *Client) DeleteHostWithOptions(request *DeleteHostRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single host.
//
// @param request - DeleteHostRequest
//
// @return DeleteHostResponse
func (client *Client) DeleteHost(request *DeleteHostRequest) (_result *DeleteHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostResponse{}
	_body, _err := client.DeleteHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single host account.
//
// Description:
//
// ### Operation description
//
// This operation deletes a single host account. If a host account is no longer in use, call this operation to delete the host account that is configured in Bastionhost.
//
// > After you delete a host account, you must manually enter the host account and password when you log on to the host through Bastionhost.
//
// ### Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Use this operation appropriately.
//
// @param request - DeleteHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostAccountResponse
func (client *Client) DeleteHostAccountWithOptions(request *DeleteHostAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single host account.
//
// Description:
//
// ### Operation description
//
// This operation deletes a single host account. If a host account is no longer in use, call this operation to delete the host account that is configured in Bastionhost.
//
// > After you delete a host account, you must manually enter the host account and password when you log on to the host through Bastionhost.
//
// ### Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Use this operation appropriately.
//
// @param request - DeleteHostAccountRequest
//
// @return DeleteHostAccountResponse
func (client *Client) DeleteHostAccount(request *DeleteHostAccountRequest) (_result *DeleteHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.DeleteHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single asset group.
//
// Description:
//
// This operation deletes a single host group. If all hosts in a host group on a bastion host no longer need to be managed through the bastion host, you can call this operation to delete the host group.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostGroupResponse
func (client *Client) DeleteHostGroupWithOptions(request *DeleteHostGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single asset group.
//
// Description:
//
// This operation deletes a single host group. If all hosts in a host group on a bastion host no longer need to be managed through the bastion host, you can call this operation to delete the host group.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteHostGroupRequest
//
// @return DeleteHostGroupResponse
func (client *Client) DeleteHostGroup(request *DeleteHostGroupRequest) (_result *DeleteHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.DeleteHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a host shared key.
//
// @param request - DeleteHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostShareKeyResponse
func (client *Client) DeleteHostShareKeyWithOptions(request *DeleteHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a host shared key.
//
// @param request - DeleteHostShareKeyRequest
//
// @return DeleteHostShareKeyResponse
func (client *Client) DeleteHostShareKey(request *DeleteHostShareKeyRequest) (_result *DeleteHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.DeleteHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single network domain.
//
// @param request - DeleteNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkDomainResponse
func (client *Client) DeleteNetworkDomainWithOptions(request *DeleteNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkDomainResponse, _err error) {
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

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single network domain.
//
// @param request - DeleteNetworkDomainRequest
//
// @return DeleteNetworkDomainResponse
func (client *Client) DeleteNetworkDomain(request *DeleteNetworkDomainRequest) (_result *DeleteNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkDomainResponse{}
	_body, _err := client.DeleteNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single control policy.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single control policy.
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a project authorization.
//
// @param request - DeleteProjectAuthorizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectAuthorizationsResponse
func (client *Client) DeleteProjectAuthorizationsWithOptions(request *DeleteProjectAuthorizationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectAuthorizationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationIds) {
		query["AuthorizationIds"] = request.AuthorizationIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectAuthorizations"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a project authorization.
//
// @param request - DeleteProjectAuthorizationsRequest
//
// @return DeleteProjectAuthorizationsResponse
func (client *Client) DeleteProjectAuthorizations(request *DeleteProjectAuthorizationsRequest) (_result *DeleteProjectAuthorizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectAuthorizationsResponse{}
	_body, _err := client.DeleteProjectAuthorizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes notification recipients from a project.
//
// @param request - DeleteProjectNotifyReceiversRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectNotifyReceiversResponse
func (client *Client) DeleteProjectNotifyReceiversWithOptions(request *DeleteProjectNotifyReceiversRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectNotifyReceiversResponse, _err error) {
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

	if !dara.IsNil(request.ReceiverIds) {
		query["ReceiverIds"] = request.ReceiverIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectNotifyReceivers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectNotifyReceiversResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes notification recipients from a project.
//
// @param request - DeleteProjectNotifyReceiversRequest
//
// @return DeleteProjectNotifyReceiversResponse
func (client *Client) DeleteProjectNotifyReceivers(request *DeleteProjectNotifyReceiversRequest) (_result *DeleteProjectNotifyReceiversResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectNotifyReceiversResponse{}
	_body, _err := client.DeleteProjectNotifyReceiversWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes projects in batches.
//
// @param request - DeleteProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectsResponse
func (client *Client) DeleteProjectsWithOptions(request *DeleteProjectsRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectsResponse, _err error) {
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

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjects"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes projects in batches.
//
// @param request - DeleteProjectsRequest
//
// @return DeleteProjectsResponse
func (client *Client) DeleteProjects(request *DeleteProjectsRequest) (_result *DeleteProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectsResponse{}
	_body, _err := client.DeleteProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an audit policy.
//
// @param request - DeleteReviewRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReviewRulesResponse
func (client *Client) DeleteReviewRulesWithOptions(request *DeleteReviewRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteReviewRulesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleIds) {
		query["ReviewRuleIds"] = request.ReviewRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteReviewRules"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReviewRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an audit policy.
//
// @param request - DeleteReviewRulesRequest
//
// @return DeleteReviewRulesResponse
func (client *Client) DeleteReviewRules(request *DeleteReviewRulesRequest) (_result *DeleteReviewRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReviewRulesResponse{}
	_body, _err := client.DeleteReviewRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single authorization rule.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single authorization rule.
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a bastion host user.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bastion host user.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a single user group from a bastion host.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single user group from a bastion host.
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a public key from the specified user.
//
// Description:
//
// You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
//
// @param request - DeleteUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPublicKeyResponse
func (client *Client) DeleteUserPublicKeyWithOptions(request *DeleteUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPublicKeyResponse, _err error) {
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

	if !dara.IsNil(request.PublicKeyId) {
		query["PublicKeyId"] = request.PublicKeyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a public key from the specified user.
//
// Description:
//
// You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
//
// @param request - DeleteUserPublicKeyRequest
//
// @return DeleteUserPublicKeyResponse
func (client *Client) DeleteUserPublicKey(request *DeleteUserPublicKeyRequest) (_result *DeleteUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.DeleteUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the attributes of an instance, such as the instance ID and description.
//
// Description:
//
// You can directly call this operation to purchase an HSM instance. For prepaid products, the payment is deducted synchronously.
//
// @param request - DescribeInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAttribute"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the attributes of an instance, such as the instance ID and description.
//
// Description:
//
// You can directly call this operation to purchase an HSM instance. For prepaid products, the payment is deducted synchronously.
//
// @param request - DescribeInstanceAttributeRequest
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of instances.
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2019-12-09"),
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
// Queries the list of instances.
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
// Queries the Alibaba Cloud regions that support Bastionhost instances.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-12-09"),
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
// Queries the Alibaba Cloud regions that support Bastionhost instances.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user.
//
// @param request - DetachDatabaseAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDatabaseAccountsFromUserResponse
func (client *Client) DetachDatabaseAccountsFromUserWithOptions(request *DetachDatabaseAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachDatabaseAccountsFromUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDatabaseAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDatabaseAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user.
//
// @param request - DetachDatabaseAccountsFromUserRequest
//
// @return DetachDatabaseAccountsFromUserResponse
func (client *Client) DetachDatabaseAccountsFromUser(request *DetachDatabaseAccountsFromUserRequest) (_result *DetachDatabaseAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDatabaseAccountsFromUserResponse{}
	_body, _err := client.DetachDatabaseAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes databases and database accounts that are authorized to a user group.
//
// @param request - DetachDatabaseAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDatabaseAccountsFromUserGroupResponse
func (client *Client) DetachDatabaseAccountsFromUserGroupWithOptions(request *DetachDatabaseAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachDatabaseAccountsFromUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDatabaseAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDatabaseAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes databases and database accounts that are authorized to a user group.
//
// @param request - DetachDatabaseAccountsFromUserGroupRequest
//
// @return DetachDatabaseAccountsFromUserGroupResponse
func (client *Client) DetachDatabaseAccountsFromUserGroup(request *DetachDatabaseAccountsFromUserGroupRequest) (_result *DetachDatabaseAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDatabaseAccountsFromUserGroupResponse{}
	_body, _err := client.DetachDatabaseAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the association between host accounts and a host shared key.
//
// @param request - DetachHostAccountsFromHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromHostShareKeyResponse
func (client *Client) DetachHostAccountsFromHostShareKeyWithOptions(request *DetachHostAccountsFromHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountIds) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the association between host accounts and a host shared key.
//
// @param request - DetachHostAccountsFromHostShareKeyRequest
//
// @return DetachHostAccountsFromHostShareKeyResponse
func (client *Client) DetachHostAccountsFromHostShareKey(request *DetachHostAccountsFromHostShareKeyRequest) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.DetachHostAccountsFromHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on hosts and host accounts from a user.
//
// @param request - DetachHostAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromUserResponse
func (client *Client) DetachHostAccountsFromUserWithOptions(request *DetachHostAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on hosts and host accounts from a user.
//
// @param request - DetachHostAccountsFromUserRequest
//
// @return DetachHostAccountsFromUserResponse
func (client *Client) DetachHostAccountsFromUser(request *DetachHostAccountsFromUserRequest) (_result *DetachHostAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.DetachHostAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes authorized hosts and host accounts from a user group.
//
// @param request - DetachHostAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromUserGroupResponse
func (client *Client) DetachHostAccountsFromUserGroupWithOptions(request *DetachHostAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes authorized hosts and host accounts from a user group.
//
// @param request - DetachHostAccountsFromUserGroupRequest
//
// @return DetachHostAccountsFromUserGroupResponse
func (client *Client) DetachHostAccountsFromUserGroup(request *DetachHostAccountsFromUserGroupRequest) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes host groups and host accounts from the list of host groups and host accounts that a user is authorized to manage.
//
// @param request - DetachHostGroupAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostGroupAccountsFromUserResponse
func (client *Client) DetachHostGroupAccountsFromUserWithOptions(request *DetachHostGroupAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostGroupAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes host groups and host accounts from the list of host groups and host accounts that a user is authorized to manage.
//
// @param request - DetachHostGroupAccountsFromUserRequest
//
// @return DetachHostGroupAccountsFromUserResponse
func (client *Client) DetachHostGroupAccountsFromUser(request *DetachHostGroupAccountsFromUserRequest) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on one or more host groups and host accounts from a user group.
//
// @param request - DetachHostGroupAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostGroupAccountsFromUserGroupResponse
func (client *Client) DetachHostGroupAccountsFromUserGroupWithOptions(request *DetachHostGroupAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostGroupAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on one or more host groups and host accounts from a user group.
//
// @param request - DetachHostGroupAccountsFromUserGroupRequest
//
// @return DetachHostGroupAccountsFromUserGroupResponse
func (client *Client) DetachHostGroupAccountsFromUserGroup(request *DetachHostGroupAccountsFromUserGroupRequest) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes members from a project.
//
// @param request - DetachMembersFromProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachMembersFromProjectResponse
func (client *Client) DetachMembersFromProjectWithOptions(request *DetachMembersFromProjectRequest, runtime *dara.RuntimeOptions) (_result *DetachMembersFromProjectResponse, _err error) {
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

	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachMembersFromProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachMembersFromProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes members from a project.
//
// @param request - DetachMembersFromProjectRequest
//
// @return DetachMembersFromProjectResponse
func (client *Client) DetachMembersFromProject(request *DetachMembersFromProjectRequest) (_result *DetachMembersFromProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachMembersFromProjectResponse{}
	_body, _err := client.DetachMembersFromProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a project administrator from project authorization management.
//
// @param request - DetachProjectManagersFromProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachProjectManagersFromProjectAuthorizationResponse
func (client *Client) DetachProjectManagersFromProjectAuthorizationWithOptions(request *DetachProjectManagersFromProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *DetachProjectManagersFromProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PrincipalIds) {
		query["PrincipalIds"] = request.PrincipalIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachProjectManagersFromProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachProjectManagersFromProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a project administrator from project authorization management.
//
// @param request - DetachProjectManagersFromProjectAuthorizationRequest
//
// @return DetachProjectManagersFromProjectAuthorizationResponse
func (client *Client) DetachProjectManagersFromProjectAuthorization(request *DetachProjectManagersFromProjectAuthorizationRequest) (_result *DetachProjectManagersFromProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachProjectManagersFromProjectAuthorizationResponse{}
	_body, _err := client.DetachProjectManagersFromProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes projects from the scope of a project authorization. This operation can be called only when the ScopeType parameter returned by the GetProjectAuthorization operation is 2.
//
// @param request - DetachProjectsFromProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachProjectsFromProjectAuthorizationResponse
func (client *Client) DetachProjectsFromProjectAuthorizationWithOptions(request *DetachProjectsFromProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *DetachProjectsFromProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachProjectsFromProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachProjectsFromProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes projects from the scope of a project authorization. This operation can be called only when the ScopeType parameter returned by the GetProjectAuthorization operation is 2.
//
// @param request - DetachProjectsFromProjectAuthorizationRequest
//
// @return DetachProjectsFromProjectAuthorizationResponse
func (client *Client) DetachProjectsFromProjectAuthorization(request *DetachProjectsFromProjectAuthorizationRequest) (_result *DetachProjectsFromProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachProjectsFromProjectAuthorizationResponse{}
	_body, _err := client.DetachProjectsFromProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes projects from the scope of an audit policy. This operation can be called only when the ScopeType parameter in the corresponding audit policy returned by the GetReviewRule operation has a value of 2.
//
// @param request - DetachProjectsFromReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachProjectsFromReviewRuleResponse
func (client *Client) DetachProjectsFromReviewRuleWithOptions(request *DetachProjectsFromReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *DetachProjectsFromReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachProjectsFromReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachProjectsFromReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes projects from the scope of an audit policy. This operation can be called only when the ScopeType parameter in the corresponding audit policy returned by the GetReviewRule operation has a value of 2.
//
// @param request - DetachProjectsFromReviewRuleRequest
//
// @return DetachProjectsFromReviewRuleResponse
func (client *Client) DetachProjectsFromReviewRule(request *DetachProjectsFromReviewRuleRequest) (_result *DetachProjectsFromReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachProjectsFromReviewRuleResponse{}
	_body, _err := client.DetachProjectsFromReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds an auditor from an audit rule.
//
// @param request - DetachReviewersFromReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachReviewersFromReviewRuleResponse
func (client *Client) DetachReviewersFromReviewRuleWithOptions(request *DetachReviewersFromReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *DetachReviewersFromReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.PrincipalIds) {
		query["PrincipalIds"] = request.PrincipalIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachReviewersFromReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachReviewersFromReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an auditor from an audit rule.
//
// @param request - DetachReviewersFromReviewRuleRequest
//
// @return DetachReviewersFromReviewRuleResponse
func (client *Client) DetachReviewersFromReviewRule(request *DetachReviewersFromReviewRuleRequest) (_result *DetachReviewersFromReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachReviewersFromReviewRuleResponse{}
	_body, _err := client.DetachReviewersFromReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables public network access for a Bastionhost instance.
//
// @param request - DisableInstancePublicAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstancePublicAccessResponse
func (client *Client) DisableInstancePublicAccessWithOptions(request *DisableInstancePublicAccessRequest, runtime *dara.RuntimeOptions) (_result *DisableInstancePublicAccessResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInstancePublicAccess"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables public network access for a Bastionhost instance.
//
// @param request - DisableInstancePublicAccessRequest
//
// @return DisableInstancePublicAccessResponse
func (client *Client) DisableInstancePublicAccess(request *DisableInstancePublicAccessRequest) (_result *DisableInstancePublicAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DisableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an authorization rule.
//
// @param request - DisableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRuleResponse
func (client *Client) DisableRuleWithOptions(request *DisableRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an authorization rule.
//
// @param request - DisableRuleRequest
//
// @return DisableRuleResponse
func (client *Client) DisableRule(request *DisableRuleRequest) (_result *DisableRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableRuleResponse{}
	_body, _err := client.DisableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables public network access for a specified Bastionhost instance.
//
// @param request - EnableInstancePublicAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstancePublicAccessResponse
func (client *Client) EnableInstancePublicAccessWithOptions(request *EnableInstancePublicAccessRequest, runtime *dara.RuntimeOptions) (_result *EnableInstancePublicAccessResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInstancePublicAccess"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables public network access for a specified Bastionhost instance.
//
// @param request - EnableInstancePublicAccessRequest
//
// @return EnableInstancePublicAccessResponse
func (client *Client) EnableInstancePublicAccess(request *EnableInstancePublicAccessRequest) (_result *EnableInstancePublicAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.EnableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an authorization rule.
//
// @param request - EnableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRuleResponse
func (client *Client) EnableRuleWithOptions(request *EnableRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an authorization rule.
//
// @param request - EnableRuleRequest
//
// @return EnableRuleResponse
func (client *Client) EnableRule(request *EnableRuleRequest) (_result *EnableRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableRuleResponse{}
	_body, _err := client.EnableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a host or database O&M token to perform O&M operations on assets.
//
// @param request - GenerateAssetOperationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAssetOperationTokenResponse
func (client *Client) GenerateAssetOperationTokenWithOptions(request *GenerateAssetOperationTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateAssetOperationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetAccountId) {
		query["AssetAccountId"] = request.AssetAccountId
	}

	if !dara.IsNil(request.AssetAccountName) {
		query["AssetAccountName"] = request.AssetAccountName
	}

	if !dara.IsNil(request.AssetAccountPassword) {
		query["AssetAccountPassword"] = request.AssetAccountPassword
	}

	if !dara.IsNil(request.AssetAccountProtocolName) {
		query["AssetAccountProtocolName"] = request.AssetAccountProtocolName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.AssetType) {
		query["AssetType"] = request.AssetType
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginAttribute) {
		query["LoginAttribute"] = request.LoginAttribute
	}

	if !dara.IsNil(request.OperationMode) {
		query["OperationMode"] = request.OperationMode
	}

	if !dara.IsNil(request.OperationNote) {
		query["OperationNote"] = request.OperationNote
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SsoClient) {
		query["SsoClient"] = request.SsoClient
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAssetOperationToken"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAssetOperationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a host or database O&M token to perform O&M operations on assets.
//
// @param request - GenerateAssetOperationTokenRequest
//
// @return GenerateAssetOperationTokenResponse
func (client *Client) GenerateAssetOperationToken(request *GenerateAssetOperationTokenRequest) (_result *GenerateAssetOperationTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAssetOperationTokenResponse{}
	_body, _err := client.GenerateAssetOperationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified O&M task.
//
// @param request - GetAutoOpsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoOpsTaskResponse
func (client *Client) GetAutoOpsTaskWithOptions(request *GetAutoOpsTaskRequest, runtime *dara.RuntimeOptions) (_result *GetAutoOpsTaskResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoOpsTask"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoOpsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified O&M task.
//
// @param request - GetAutoOpsTaskRequest
//
// @return GetAutoOpsTaskResponse
func (client *Client) GetAutoOpsTask(request *GetAutoOpsTaskRequest) (_result *GetAutoOpsTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoOpsTaskResponse{}
	_body, _err := client.GetAutoOpsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified database.
//
// @param request - GetDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified database.
//
// @param request - GetDatabaseRequest
//
// @return GetDatabaseResponse
func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified database account.
//
// @param request - GetDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseAccountResponse
func (client *Client) GetDatabaseAccountWithOptions(request *GetDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified database account.
//
// @param request - GetDatabaseAccountRequest
//
// @return GetDatabaseAccountResponse
func (client *Client) GetDatabaseAccount(request *GetDatabaseAccountRequest) (_result *GetDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseAccountResponse{}
	_body, _err := client.GetDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a configuration backup export task. This operation is supported only for Bastionhost V3.2.44 and later.
//
// @param request - GetExportConfigJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExportConfigJobResponse
func (client *Client) GetExportConfigJobWithOptions(request *GetExportConfigJobRequest, runtime *dara.RuntimeOptions) (_result *GetExportConfigJobResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExportConfigJob"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExportConfigJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a configuration backup export task. This operation is supported only for Bastionhost V3.2.44 and later.
//
// @param request - GetExportConfigJobRequest
//
// @return GetExportConfigJobResponse
func (client *Client) GetExportConfigJob(request *GetExportConfigJobRequest) (_result *GetExportConfigJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExportConfigJobResponse{}
	_body, _err := client.GetExportConfigJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Includes information such as the hostname, source, host address, and protocol port.
//
// @param request - GetHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostResponse
func (client *Client) GetHostWithOptions(request *GetHostRequest, runtime *dara.RuntimeOptions) (_result *GetHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Includes information such as the hostname, source, host address, and protocol port.
//
// @param request - GetHostRequest
//
// @return GetHostResponse
func (client *Client) GetHost(request *GetHostRequest) (_result *GetHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostResponse{}
	_body, _err := client.GetHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified host account.
//
// @param request - GetHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostAccountResponse
func (client *Client) GetHostAccountWithOptions(request *GetHostAccountRequest, runtime *dara.RuntimeOptions) (_result *GetHostAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified host account.
//
// @param request - GetHostAccountRequest
//
// @return GetHostAccountResponse
func (client *Client) GetHostAccount(request *GetHostAccountRequest) (_result *GetHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostAccountResponse{}
	_body, _err := client.GetHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified asset group.
//
// @param request - GetHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostGroupResponse
func (client *Client) GetHostGroupWithOptions(request *GetHostGroupRequest, runtime *dara.RuntimeOptions) (_result *GetHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified asset group.
//
// @param request - GetHostGroupRequest
//
// @return GetHostGroupResponse
func (client *Client) GetHostGroup(request *GetHostGroupRequest) (_result *GetHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostGroupResponse{}
	_body, _err := client.GetHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a host shared key.
//
// @param request - GetHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostShareKeyResponse
func (client *Client) GetHostShareKeyWithOptions(request *GetHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *GetHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a host shared key.
//
// @param request - GetHostShareKeyRequest
//
// @return GetHostShareKeyResponse
func (client *Client) GetHostShareKey(request *GetHostShareKeyRequest) (_result *GetHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.GetHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Active Directory (AD) authentication configurations of a Bastionhost instance.
//
// Description:
//
// ### Usage notes
//
// This operation queries the Active Directory (AD) authentication configurations of a Bastionhost instance. You can use the AD authentication feature of Bastionhost to synchronize users from an AD server to Bastionhost. After the synchronization is complete, users can perform O&M operations on servers by using Bastionhost. For enterprise-level managed operations, you can also refer to Alibaba Cloud Managed Services.
//
// ### QPS limits
//
// The queries per second (QPS) limit on this operation is 10 per user. If the number of API invocations exceeds the limit, throttling is triggered, which may affect your business. Invoke this operation at a reasonable frequency.
//
// @param request - GetInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceADAuthServerResponse
func (client *Client) GetInstanceADAuthServerWithOptions(request *GetInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceADAuthServerResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserSourceId) {
		query["UserSourceId"] = request.UserSourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Active Directory (AD) authentication configurations of a Bastionhost instance.
//
// Description:
//
// ### Usage notes
//
// This operation queries the Active Directory (AD) authentication configurations of a Bastionhost instance. You can use the AD authentication feature of Bastionhost to synchronize users from an AD server to Bastionhost. After the synchronization is complete, users can perform O&M operations on servers by using Bastionhost. For enterprise-level managed operations, you can also refer to Alibaba Cloud Managed Services.
//
// ### QPS limits
//
// The queries per second (QPS) limit on this operation is 10 per user. If the number of API invocations exceeds the limit, throttling is triggered, which may affect your business. Invoke this operation at a reasonable frequency.
//
// @param request - GetInstanceADAuthServerRequest
//
// @return GetInstanceADAuthServerResponse
func (client *Client) GetInstanceADAuthServer(request *GetInstanceADAuthServerRequest) (_result *GetInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.GetInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of Lightweight Directory Access Protocol (LDAP) authentication on a bastion host.
//
// @param request - GetInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLDAPAuthServerResponse
func (client *Client) GetInstanceLDAPAuthServerWithOptions(request *GetInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of Lightweight Directory Access Protocol (LDAP) authentication on a bastion host.
//
// @param request - GetInstanceLDAPAuthServerRequest
//
// @return GetInstanceLDAPAuthServerResponse
func (client *Client) GetInstanceLDAPAuthServer(request *GetInstanceLDAPAuthServerRequest) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.GetInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage usage of a Bastionhost instance.
//
// @param request - GetInstanceStoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceStoreInfoResponse
func (client *Client) GetInstanceStoreInfoWithOptions(request *GetInstanceStoreInfoRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceStoreInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceStoreInfo"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceStoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage usage of a Bastionhost instance.
//
// @param request - GetInstanceStoreInfoRequest
//
// @return GetInstanceStoreInfoResponse
func (client *Client) GetInstanceStoreInfo(request *GetInstanceStoreInfoRequest) (_result *GetInstanceStoreInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceStoreInfoResponse{}
	_body, _err := client.GetInstanceStoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of two-factor authentication on a bastion host.
//
// Description:
//
// Queries the two-factor authentication settings of a bastion host. When two-factor authentication is enabled, Bastionhost sends a verification code to a local user who logs on to the bastion host. The local user must enter a valid username, password, and verification code to log on, which reduces the security risks of account information leaks.
//
// ### Limit
//
// This operation can be called up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered and your business may be affected. Take note of this limit when you call this operation.
//
// @param request - GetInstanceTwoFactorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceTwoFactorResponse
func (client *Client) GetInstanceTwoFactorWithOptions(request *GetInstanceTwoFactorRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceTwoFactorResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceTwoFactor"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of two-factor authentication on a bastion host.
//
// Description:
//
// Queries the two-factor authentication settings of a bastion host. When two-factor authentication is enabled, Bastionhost sends a verification code to a local user who logs on to the bastion host. The local user must enter a valid username, password, and verification code to log on, which reduces the security risks of account information leaks.
//
// ### Limit
//
// This operation can be called up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered and your business may be affected. Take note of this limit when you call this operation.
//
// @param request - GetInstanceTwoFactorRequest
//
// @return GetInstanceTwoFactorResponse
func (client *Client) GetInstanceTwoFactor(request *GetInstanceTwoFactorRequest) (_result *GetInstanceTwoFactorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.GetInstanceTwoFactorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified network domain.
//
// @param request - GetNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkDomainResponse
func (client *Client) GetNetworkDomainWithOptions(request *GetNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckProxyState) {
		query["CheckProxyState"] = request.CheckProxyState
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified network domain.
//
// @param request - GetNetworkDomainRequest
//
// @return GetNetworkDomainResponse
func (client *Client) GetNetworkDomain(request *GetNetworkDomainRequest) (_result *GetNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkDomainResponse{}
	_body, _err := client.GetNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified control policy.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified control policy.
//
// @param request - GetPolicyRequest
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the asset scope of a specified control policy.
//
// @param request - GetPolicyAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyAssetScopeResponse
func (client *Client) GetPolicyAssetScopeWithOptions(request *GetPolicyAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyAssetScopeResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyAssetScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyAssetScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the asset scope of a specified control policy.
//
// @param request - GetPolicyAssetScopeRequest
//
// @return GetPolicyAssetScopeResponse
func (client *Client) GetPolicyAssetScope(request *GetPolicyAssetScopeRequest) (_result *GetPolicyAssetScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyAssetScopeResponse{}
	_body, _err := client.GetPolicyAssetScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the effective user scope of a specified control policy.
//
// @param request - GetPolicyUserScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyUserScopeResponse
func (client *Client) GetPolicyUserScopeWithOptions(request *GetPolicyUserScopeRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyUserScopeResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyUserScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyUserScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the effective user scope of a specified control policy.
//
// @param request - GetPolicyUserScopeRequest
//
// @return GetPolicyUserScopeResponse
func (client *Client) GetPolicyUserScope(request *GetPolicyUserScopeRequest) (_result *GetPolicyUserScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyUserScopeResponse{}
	_body, _err := client.GetPolicyUserScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified project.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified project.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a project authorization.
//
// @param request - GetProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectAuthorizationResponse
func (client *Client) GetProjectAuthorizationWithOptions(request *GetProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *GetProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a project authorization.
//
// @param request - GetProjectAuthorizationRequest
//
// @return GetProjectAuthorizationResponse
func (client *Client) GetProjectAuthorization(request *GetProjectAuthorizationRequest) (_result *GetProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectAuthorizationResponse{}
	_body, _err := client.GetProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an audit policy.
//
// @param request - GetReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReviewRuleResponse
func (client *Client) GetReviewRuleWithOptions(request *GetReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *GetReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an audit policy.
//
// @param request - GetReviewRuleRequest
//
// @return GetReviewRuleResponse
func (client *Client) GetReviewRule(request *GetReviewRuleRequest) (_result *GetReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetReviewRuleResponse{}
	_body, _err := client.GetReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified authorization rule.
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified authorization rule.
//
// @param request - GetRuleRequest
//
// @return GetRuleResponse
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified Bastionhost user.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified Bastionhost user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified Bastionhost user group.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified Bastionhost user group.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports KMS secrets for a specified host.
//
// @param tmpReq - ImportKMSSecretsForHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKMSSecretsForHostResponse
func (client *Client) ImportKMSSecretsForHostWithOptions(tmpReq *ImportKMSSecretsForHostRequest, runtime *dara.RuntimeOptions) (_result *ImportKMSSecretsForHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportKMSSecretsForHostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Secrets) {
		request.SecretsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Secrets, dara.String("Secrets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecretsShrink) {
		query["Secrets"] = request.SecretsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportKMSSecretsForHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportKMSSecretsForHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports KMS secrets for a specified host.
//
// @param request - ImportKMSSecretsForHostRequest
//
// @return ImportKMSSecretsForHostResponse
func (client *Client) ImportKMSSecretsForHost(request *ImportKMSSecretsForHostRequest) (_result *ImportKMSSecretsForHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportKMSSecretsForHostResponse{}
	_body, _err := client.ImportKMSSecretsForHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of commands that require approval.
//
// Description:
//
// This API operation allows administrators to retrieve the list of commands that require approval.
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
//
// @param request - ListApproveCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApproveCommandsResponse
func (client *Client) ListApproveCommandsWithOptions(request *ListApproveCommandsRequest, runtime *dara.RuntimeOptions) (_result *ListApproveCommandsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApproveCommands"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of commands that require approval.
//
// Description:
//
// This API operation allows administrators to retrieve the list of commands that require approval.
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
//
// @param request - ListApproveCommandsRequest
//
// @return ListApproveCommandsResponse
func (client *Client) ListApproveCommands(request *ListApproveCommandsRequest) (_result *ListApproveCommandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.ListApproveCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of database accounts.
//
// @param request - ListDatabaseAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsResponse
func (client *Client) ListDatabaseAccountsWithOptions(request *ListDatabaseAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of database accounts.
//
// @param request - ListDatabaseAccountsRequest
//
// @return ListDatabaseAccountsResponse
func (client *Client) ListDatabaseAccounts(request *ListDatabaseAccountsRequest) (_result *ListDatabaseAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsResponse{}
	_body, _err := client.ListDatabaseAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsForUserResponse
func (client *Client) ListDatabaseAccountsForUserWithOptions(request *ListDatabaseAccountsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccountsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserRequest
//
// @return ListDatabaseAccountsForUserResponse
func (client *Client) ListDatabaseAccountsForUser(request *ListDatabaseAccountsForUserRequest) (_result *ListDatabaseAccountsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsForUserResponse{}
	_body, _err := client.ListDatabaseAccountsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized and unauthorized database accounts for a user group.
//
// @param request - ListDatabaseAccountsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsForUserGroupResponse
func (client *Client) ListDatabaseAccountsForUserGroupWithOptions(request *ListDatabaseAccountsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccountsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized and unauthorized database accounts for a user group.
//
// @param request - ListDatabaseAccountsForUserGroupRequest
//
// @return ListDatabaseAccountsForUserGroupResponse
func (client *Client) ListDatabaseAccountsForUserGroup(request *ListDatabaseAccountsForUserGroupRequest) (_result *ListDatabaseAccountsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsForUserGroupResponse{}
	_body, _err := client.ListDatabaseAccountsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases of a specified Bastionhost instance.
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases of a specified Bastionhost instance.
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of databases authorized for a user.
//
// @param request - ListDatabasesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesForUserResponse
func (client *Client) ListDatabasesForUserWithOptions(request *ListDatabasesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabasesForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of databases authorized for a user.
//
// @param request - ListDatabasesForUserRequest
//
// @return ListDatabasesForUserResponse
func (client *Client) ListDatabasesForUser(request *ListDatabasesForUserRequest) (_result *ListDatabasesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesForUserResponse{}
	_body, _err := client.ListDatabasesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases authorized for a user group.
//
// @param request - ListDatabasesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesForUserGroupResponse
func (client *Client) ListDatabasesForUserGroupWithOptions(request *ListDatabasesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabasesForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases authorized for a user group.
//
// @param request - ListDatabasesForUserGroupRequest
//
// @return ListDatabasesForUserGroupResponse
func (client *Client) ListDatabasesForUserGroup(request *ListDatabasesForUserGroupRequest) (_result *ListDatabasesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesForUserGroupResponse{}
	_body, _err := client.ListDatabasesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of host accounts.
//
// @param request - ListHostAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsResponse
func (client *Client) ListHostAccountsWithOptions(request *ListHostAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of host accounts.
//
// @param request - ListHostAccountsRequest
//
// @return ListHostAccountsResponse
func (client *Client) ListHostAccounts(request *ListHostAccountsRequest) (_result *ListHostAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.ListHostAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the host accounts associated with a host shared key.
//
// @param request - ListHostAccountsForHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForHostShareKeyResponse
func (client *Client) ListHostAccountsForHostShareKeyWithOptions(request *ListHostAccountsForHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the host accounts associated with a host shared key.
//
// @param request - ListHostAccountsForHostShareKeyRequest
//
// @return ListHostAccountsForHostShareKeyResponse
func (client *Client) ListHostAccountsForHostShareKey(request *ListHostAccountsForHostShareKeyRequest) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.ListHostAccountsForHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized and unauthorized host accounts for a specified user on a specified host.
//
// @param request - ListHostAccountsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForUserResponse
func (client *Client) ListHostAccountsForUserWithOptions(request *ListHostAccountsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized and unauthorized host accounts for a specified user on a specified host.
//
// @param request - ListHostAccountsForUserRequest
//
// @return ListHostAccountsForUserResponse
func (client *Client) ListHostAccountsForUser(request *ListHostAccountsForUserRequest) (_result *ListHostAccountsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.ListHostAccountsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorized and unauthorized host accounts of a specified user group on a specified host.
//
// @param request - ListHostAccountsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForUserGroupResponse
func (client *Client) ListHostAccountsForUserGroupWithOptions(request *ListHostAccountsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorized and unauthorized host accounts of a specified user group on a specified host.
//
// @param request - ListHostAccountsForUserGroupRequest
//
// @return ListHostAccountsForUserGroupResponse
func (client *Client) ListHostAccountsForUserGroup(request *ListHostAccountsForUserGroupRequest) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.ListHostAccountsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of host accounts authorized for a specified user in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupAccountNamesForUserResponse
func (client *Client) ListHostGroupAccountNamesForUserWithOptions(request *ListHostGroupAccountNamesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupAccountNamesForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of host accounts authorized for a specified user in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserRequest
//
// @return ListHostGroupAccountNamesForUserResponse
func (client *Client) ListHostGroupAccountNamesForUser(request *ListHostGroupAccountNamesForUserRequest) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of host accounts authorized for a specified user group in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupAccountNamesForUserGroupResponse
func (client *Client) ListHostGroupAccountNamesForUserGroupWithOptions(request *ListHostGroupAccountNamesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupAccountNamesForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of host accounts authorized for a specified user group in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserGroupRequest
//
// @return ListHostGroupAccountNamesForUserGroupResponse
func (client *Client) ListHostGroupAccountNamesForUserGroup(request *ListHostGroupAccountNamesForUserGroupRequest) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the asset groups under a specified bastion host.
//
// @param request - ListHostGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsResponse
func (client *Client) ListHostGroupsWithOptions(request *ListHostGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the asset groups under a specified bastion host.
//
// @param request - ListHostGroupsRequest
//
// @return ListHostGroupsResponse
func (client *Client) ListHostGroups(request *ListHostGroupsRequest) (_result *ListHostGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.ListHostGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorized or unauthorized host groups for a specified Bastionhost user.
//
// @param request - ListHostGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsForUserResponse
func (client *Client) ListHostGroupsForUserWithOptions(request *ListHostGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorized or unauthorized host groups for a specified Bastionhost user.
//
// @param request - ListHostGroupsForUserRequest
//
// @return ListHostGroupsForUserResponse
func (client *Client) ListHostGroupsForUser(request *ListHostGroupsForUserRequest) (_result *ListHostGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.ListHostGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorized or unauthorized host groups for a specified Bastionhost user group.
//
// @param request - ListHostGroupsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsForUserGroupResponse
func (client *Client) ListHostGroupsForUserGroupWithOptions(request *ListHostGroupsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorized or unauthorized host groups for a specified Bastionhost user group.
//
// @param request - ListHostGroupsForUserGroupRequest
//
// @return ListHostGroupsForUserGroupResponse
func (client *Client) ListHostGroupsForUserGroup(request *ListHostGroupsForUserGroupRequest) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.ListHostGroupsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of host shared keys.
//
// @param request - ListHostShareKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostShareKeysResponse
func (client *Client) ListHostShareKeysWithOptions(request *ListHostShareKeysRequest, runtime *dara.RuntimeOptions) (_result *ListHostShareKeysResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostShareKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of host shared keys.
//
// @param request - ListHostShareKeysRequest
//
// @return ListHostShareKeysResponse
func (client *Client) ListHostShareKeys(request *ListHostShareKeysRequest) (_result *ListHostShareKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.ListHostShareKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hosts in a specified Bastionhost instance.
//
// @param request - ListHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsResponse
func (client *Client) ListHostsWithOptions(request *ListHostsRequest, runtime *dara.RuntimeOptions) (_result *ListHostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHosts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hosts in a specified Bastionhost instance.
//
// @param request - ListHostsRequest
//
// @return ListHostsResponse
func (client *Client) ListHosts(request *ListHostsRequest) (_result *ListHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsResponse{}
	_body, _err := client.ListHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized or unauthorized hosts for a specified Bastionhost user.
//
// @param request - ListHostsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsForUserResponse
func (client *Client) ListHostsForUserWithOptions(request *ListHostsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized or unauthorized hosts for a specified Bastionhost user.
//
// @param request - ListHostsForUserRequest
//
// @return ListHostsForUserResponse
func (client *Client) ListHostsForUser(request *ListHostsForUserRequest) (_result *ListHostsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.ListHostsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized or unauthorized hosts for a specified Bastionhost user group.
//
// @param request - ListHostsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsForUserGroupResponse
func (client *Client) ListHostsForUserGroupWithOptions(request *ListHostsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostsForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized or unauthorized hosts for a specified Bastionhost user group.
//
// @param request - ListHostsForUserGroupRequest
//
// @return ListHostsForUserGroupResponse
func (client *Client) ListHostsForUserGroup(request *ListHostsForUserGroupRequest) (_result *ListHostsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.ListHostsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all importable KMS credentials for a specified host.
//
// @param request - ListImportableKMSSecretsForHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImportableKMSSecretsForHostResponse
func (client *Client) ListImportableKMSSecretsForHostWithOptions(request *ListImportableKMSSecretsForHostRequest, runtime *dara.RuntimeOptions) (_result *ListImportableKMSSecretsForHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImportableKMSSecretsForHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImportableKMSSecretsForHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all importable KMS credentials for a specified host.
//
// @param request - ListImportableKMSSecretsForHostRequest
//
// @return ListImportableKMSSecretsForHostResponse
func (client *Client) ListImportableKMSSecretsForHost(request *ListImportableKMSSecretsForHostRequest) (_result *ListImportableKMSSecretsForHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImportableKMSSecretsForHostResponse{}
	_body, _err := client.ListImportableKMSSecretsForHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of member accounts in a resource directory.
//
// @param request - ListInstanceRdMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRdMembersResponse
func (client *Client) ListInstanceRdMembersWithOptions(request *ListInstanceRdMembersRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceRdMembersResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceRdMembers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceRdMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of member accounts in a resource directory.
//
// @param request - ListInstanceRdMembersRequest
//
// @return ListInstanceRdMembersResponse
func (client *Client) ListInstanceRdMembers(request *ListInstanceRdMembersRequest) (_result *ListInstanceRdMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceRdMembersResponse{}
	_body, _err := client.ListInstanceRdMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of member accounts that are not assigned to other projects.
//
// @param request - ListMembersNotForProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersNotForProjectResponse
func (client *Client) ListMembersNotForProjectWithOptions(request *ListMembersNotForProjectRequest, runtime *dara.RuntimeOptions) (_result *ListMembersNotForProjectResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMembersNotForProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMembersNotForProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of member accounts that are not assigned to other projects.
//
// @param request - ListMembersNotForProjectRequest
//
// @return ListMembersNotForProjectResponse
func (client *Client) ListMembersNotForProject(request *ListMembersNotForProjectRequest) (_result *ListMembersNotForProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMembersNotForProjectResponse{}
	_body, _err := client.ListMembersNotForProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network domains of a specified Bastionhost instance.
//
// @param request - ListNetworkDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkDomainsResponse
func (client *Client) ListNetworkDomainsWithOptions(request *ListNetworkDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkDomainsResponse, _err error) {
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

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkDomains"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network domains of a specified Bastionhost instance.
//
// @param request - ListNetworkDomainsRequest
//
// @return ListNetworkDomainsResponse
func (client *Client) ListNetworkDomains(request *ListNetworkDomainsRequest) (_result *ListNetworkDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkDomainsResponse{}
	_body, _err := client.ListNetworkDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of authorized database accounts.
//
// @param request - ListOperationDatabaseAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationDatabaseAccountsResponse
func (client *Client) ListOperationDatabaseAccountsWithOptions(request *ListOperationDatabaseAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationDatabaseAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationProjectId) {
		query["OperationProjectId"] = request.OperationProjectId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationDatabaseAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationDatabaseAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of authorized database accounts.
//
// @param request - ListOperationDatabaseAccountsRequest
//
// @return ListOperationDatabaseAccountsResponse
func (client *Client) ListOperationDatabaseAccounts(request *ListOperationDatabaseAccountsRequest) (_result *ListOperationDatabaseAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationDatabaseAccountsResponse{}
	_body, _err := client.ListOperationDatabaseAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of databases authorized for the current Resource Access Management (RAM) user.
//
// @param request - ListOperationDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationDatabasesResponse
func (client *Client) ListOperationDatabasesWithOptions(request *ListOperationDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListOperationDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationDatabases"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of databases authorized for the current Resource Access Management (RAM) user.
//
// @param request - ListOperationDatabasesRequest
//
// @return ListOperationDatabasesResponse
func (client *Client) ListOperationDatabases(request *ListOperationDatabasesRequest) (_result *ListOperationDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationDatabasesResponse{}
	_body, _err := client.ListOperationDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the host accounts for which the current Resource Access Management (RAM) user has authorization.
//
// @param request - ListOperationHostAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationHostAccountsResponse
func (client *Client) ListOperationHostAccountsWithOptions(request *ListOperationHostAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationHostAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationHostAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationHostAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the host accounts for which the current Resource Access Management (RAM) user has authorization.
//
// @param request - ListOperationHostAccountsRequest
//
// @return ListOperationHostAccountsResponse
func (client *Client) ListOperationHostAccounts(request *ListOperationHostAccountsRequest) (_result *ListOperationHostAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationHostAccountsResponse{}
	_body, _err := client.ListOperationHostAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of hosts with authorization for the current Resource Access Management (RAM) user.
//
// @param request - ListOperationHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationHostsResponse
func (client *Client) ListOperationHostsWithOptions(request *ListOperationHostsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationHostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.OperationProjectId) {
		query["OperationProjectId"] = request.OperationProjectId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationHosts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of hosts with authorization for the current Resource Access Management (RAM) user.
//
// @param request - ListOperationHostsRequest
//
// @return ListOperationHostsResponse
func (client *Client) ListOperationHosts(request *ListOperationHostsRequest) (_result *ListOperationHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationHostsResponse{}
	_body, _err := client.ListOperationHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of O&M tickets pending approval.
//
// Description:
//
// This API operation allows administrators to retrieve a list of O&M tickets pending approval.
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param request - ListOperationTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationTicketsResponse
func (client *Client) ListOperationTicketsWithOptions(request *ListOperationTicketsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationTicketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetAddress) {
		query["AssetAddress"] = request.AssetAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationTickets"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of O&M tickets pending approval.
//
// Description:
//
// This API operation allows administrators to retrieve a list of O&M tickets pending approval.
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param request - ListOperationTicketsRequest
//
// @return ListOperationTicketsResponse
func (client *Client) ListOperationTickets(request *ListOperationTicketsRequest) (_result *ListOperationTicketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.ListOperationTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the control policies.
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the control policies.
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of project authorizations.
//
// @param request - ListProjectAuthorizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectAuthorizationsResponse
func (client *Client) ListProjectAuthorizationsWithOptions(request *ListProjectAuthorizationsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectAuthorizationsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectAuthorizations"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of project authorizations.
//
// @param request - ListProjectAuthorizationsRequest
//
// @return ListProjectAuthorizationsResponse
func (client *Client) ListProjectAuthorizations(request *ListProjectAuthorizationsRequest) (_result *ListProjectAuthorizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectAuthorizationsResponse{}
	_body, _err := client.ListProjectAuthorizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of notification recipients for a project.
//
// @param request - ListProjectNotifyReceiversRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectNotifyReceiversResponse
func (client *Client) ListProjectNotifyReceiversWithOptions(request *ListProjectNotifyReceiversRequest, runtime *dara.RuntimeOptions) (_result *ListProjectNotifyReceiversResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectNotifyReceivers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectNotifyReceiversResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of notification recipients for a project.
//
// @param request - ListProjectNotifyReceiversRequest
//
// @return ListProjectNotifyReceiversResponse
func (client *Client) ListProjectNotifyReceivers(request *ListProjectNotifyReceiversRequest) (_result *ListProjectNotifyReceiversResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectNotifyReceiversResponse{}
	_body, _err := client.ListProjectNotifyReceiversWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of projects.
//
// @param request - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of projects.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the projects that the current user can manage.
//
// @param request - ListProjectsForProjectAuthorizationPrincipalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsForProjectAuthorizationPrincipalResponse
func (client *Client) ListProjectsForProjectAuthorizationPrincipalWithOptions(request *ListProjectsForProjectAuthorizationPrincipalRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsForProjectAuthorizationPrincipalResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectsForProjectAuthorizationPrincipal"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsForProjectAuthorizationPrincipalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the projects that the current user can manage.
//
// @param request - ListProjectsForProjectAuthorizationPrincipalRequest
//
// @return ListProjectsForProjectAuthorizationPrincipalResponse
func (client *Client) ListProjectsForProjectAuthorizationPrincipal(request *ListProjectsForProjectAuthorizationPrincipalRequest) (_result *ListProjectsForProjectAuthorizationPrincipalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsForProjectAuthorizationPrincipalResponse{}
	_body, _err := client.ListProjectsForProjectAuthorizationPrincipalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of projects that are not associated with a project authorization.
//
// @param request - ListProjectsNotForProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsNotForProjectAuthorizationResponse
func (client *Client) ListProjectsNotForProjectAuthorizationWithOptions(request *ListProjectsNotForProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsNotForProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectsNotForProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsNotForProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of projects that are not associated with a project authorization.
//
// @param request - ListProjectsNotForProjectAuthorizationRequest
//
// @return ListProjectsNotForProjectAuthorizationResponse
func (client *Client) ListProjectsNotForProjectAuthorization(request *ListProjectsNotForProjectAuthorizationRequest) (_result *ListProjectsNotForProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsNotForProjectAuthorizationResponse{}
	_body, _err := client.ListProjectsNotForProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of projects that are not associated with an audit policy.
//
// @param request - ListProjectsNotForReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsNotForReviewRuleResponse
func (client *Client) ListProjectsNotForReviewRuleWithOptions(request *ListProjectsNotForReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsNotForReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectsNotForReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsNotForReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of projects that are not associated with an audit policy.
//
// @param request - ListProjectsNotForReviewRuleRequest
//
// @return ListProjectsNotForReviewRuleResponse
func (client *Client) ListProjectsNotForReviewRule(request *ListProjectsNotForReviewRuleRequest) (_result *ListProjectsNotForReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsNotForReviewRuleResponse{}
	_body, _err := client.ListProjectsNotForReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the audit policy list.
//
// @param request - ListReviewRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReviewRulesResponse
func (client *Client) ListReviewRulesWithOptions(request *ListReviewRulesRequest, runtime *dara.RuntimeOptions) (_result *ListReviewRulesResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReviewRules"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReviewRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit policy list.
//
// @param request - ListReviewRulesRequest
//
// @return ListReviewRulesResponse
func (client *Client) ListReviewRules(request *ListReviewRulesRequest) (_result *ListReviewRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReviewRulesResponse{}
	_body, _err := client.ListReviewRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorization rules of a specified Bastionhost instance.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleState) {
		query["RuleState"] = request.RuleState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization rules of a specified Bastionhost instance.
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags bound to a resource.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags bound to a resource.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are bound to one or more Bastionhost instances.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are bound to one or more Bastionhost instances.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of pending automated O&M approvals.
//
// @param request - ListTodoOpsTaskApprovalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTodoOpsTaskApprovalsResponse
func (client *Client) ListTodoOpsTaskApprovalsWithOptions(request *ListTodoOpsTaskApprovalsRequest, runtime *dara.RuntimeOptions) (_result *ListTodoOpsTaskApprovalsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScheduleType) {
		query["ScheduleType"] = request.ScheduleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTodoOpsTaskApprovals"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTodoOpsTaskApprovalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of pending automated O&M approvals.
//
// @param request - ListTodoOpsTaskApprovalsRequest
//
// @return ListTodoOpsTaskApprovalsResponse
func (client *Client) ListTodoOpsTaskApprovals(request *ListTodoOpsTaskApprovalsRequest) (_result *ListTodoOpsTaskApprovalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTodoOpsTaskApprovalsResponse{}
	_body, _err := client.ListTodoOpsTaskApprovalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the user groups of a specified Bastionhost instance.
//
// @param request - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user groups of a specified Bastionhost instance.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of user groups that a user has joined.
//
// @param request - ListUserGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForUserResponse
func (client *Client) ListUserGroupsForUserWithOptions(request *ListUserGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsForUserResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PagingOff) {
		query["PagingOff"] = request.PagingOff
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of user groups that a user has joined.
//
// @param request - ListUserGroupsForUserRequest
//
// @return ListUserGroupsForUserResponse
func (client *Client) ListUserGroupsForUser(request *ListUserGroupsForUserRequest) (_result *ListUserGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsForUserResponse{}
	_body, _err := client.ListUserGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the SSH public keys that are configured for a Bastionhost user. Users can use these public keys for key-based authentication when logging on to hosts through Bastionhost.
//
// @param request - ListUserPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPublicKeysResponse
func (client *Client) ListUserPublicKeysWithOptions(request *ListUserPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListUserPublicKeysResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPublicKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the SSH public keys that are configured for a Bastionhost user. Users can use these public keys for key-based authentication when logging on to hosts through Bastionhost.
//
// @param request - ListUserPublicKeysRequest
//
// @return ListUserPublicKeysResponse
func (client *Client) ListUserPublicKeys(request *ListUserPublicKeysRequest) (_result *ListUserPublicKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.ListUserPublicKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of users for a specified bastion host.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceUserId) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserState) {
		query["UserState"] = request.UserState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of users for a specified bastion host.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Locks one or more users of a bastion host.
//
// Description:
//
// # Description
//
// Locks one or more bastion host users to prevent them from performing O\\&M operations. If a user does not need to use a bastion host within a specific period of time, lock the user account. Locked users cannot log on to or perform O\\&M operations on the hosts on which they have permissions. To unlock users later, use the [UnlockUsers](https://help.aliyun.com/document_detail/204590.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockUsersResponse
func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *dara.RuntimeOptions) (_result *LockUsersResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockUsers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks one or more users of a bastion host.
//
// Description:
//
// # Description
//
// Locks one or more bastion host users to prevent them from performing O\\&M operations. If a user does not need to use a bastion host within a specific period of time, lock the user account. Locked users cannot log on to or perform O\\&M operations on the hosts on which they have permissions. To unlock users later, use the [UnlockUsers](https://help.aliyun.com/document_detail/204590.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockUsersRequest
//
// @return LockUsersResponse
func (client *Client) LockUsers(request *LockUsersRequest) (_result *LockUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockUsersResponse{}
	_body, _err := client.LockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a database.
//
// @param request - ModifyDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabaseWithOptions(request *ModifyDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabasePort) {
		query["DatabasePort"] = request.DatabasePort
	}

	if !dara.IsNil(request.DatabasePrivateAddress) {
		query["DatabasePrivateAddress"] = request.DatabasePrivateAddress
	}

	if !dara.IsNil(request.DatabasePublicAddress) {
		query["DatabasePublicAddress"] = request.DatabasePublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a database.
//
// @param request - ModifyDatabaseRequest
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabase(request *ModifyDatabaseRequest) (_result *ModifyDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.ModifyDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a database account.
//
// @param request - ModifyDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseAccountResponse
func (client *Client) ModifyDatabaseAccountWithOptions(request *ModifyDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a database account.
//
// @param request - ModifyDatabaseAccountRequest
//
// @return ModifyDatabaseAccountResponse
func (client *Client) ModifyDatabaseAccount(request *ModifyDatabaseAccountRequest) (_result *ModifyDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabaseAccountResponse{}
	_body, _err := client.ModifyDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves multiple database instances to a specified project.
//
// @param request - ModifyDatabasesProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabasesProjectResponse
func (client *Client) ModifyDatabasesProjectWithOptions(request *ModifyDatabasesProjectRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabasesProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetProjectId) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabasesProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabasesProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves multiple database instances to a specified project.
//
// @param request - ModifyDatabasesProjectRequest
//
// @return ModifyDatabasesProjectResponse
func (client *Client) ModifyDatabasesProject(request *ModifyDatabasesProjectRequest) (_result *ModifyDatabasesProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabasesProjectResponse{}
	_body, _err := client.ModifyDatabasesProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the address, name, operating system type, and description of a host.
//
// Description:
//
// This operation updates the basic information of on-premises hosts, ECS hosts, and hosts in dedicated ApsaraDB RDS clusters.
//
// > If you update the basic information of an ECS host or a host in a dedicated ApsaraDB RDS cluster, the changes may be overwritten. This is because Bastionhost periodically synchronizes the basic information of ECS hosts and hosts in dedicated ApsaraDB RDS clusters.
//
// @param request - ModifyHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostResponse
func (client *Client) ModifyHostWithOptions(request *ModifyHostRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HostPrivateAddress) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !dara.IsNil(request.HostPublicAddress) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PrefKex) {
		query["PrefKex"] = request.PrefKex
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the address, name, operating system type, and description of a host.
//
// Description:
//
// This operation updates the basic information of on-premises hosts, ECS hosts, and hosts in dedicated ApsaraDB RDS clusters.
//
// > If you update the basic information of an ECS host or a host in a dedicated ApsaraDB RDS cluster, the changes may be overwritten. This is because Bastionhost periodically synchronizes the basic information of ECS hosts and hosts in dedicated ApsaraDB RDS clusters.
//
// @param request - ModifyHostRequest
//
// @return ModifyHostResponse
func (client *Client) ModifyHost(request *ModifyHostRequest) (_result *ModifyHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostResponse{}
	_body, _err := client.ModifyHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies host account information, including the name, password, and private key of a host account.
//
// @param request - ModifyHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostAccountResponse
func (client *Client) ModifyHostAccountWithOptions(request *ModifyHostAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RotationMode) {
		query["RotationMode"] = request.RotationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies host account information, including the name, password, and private key of a host account.
//
// @param request - ModifyHostAccountRequest
//
// @return ModifyHostAccountResponse
func (client *Client) ModifyHostAccount(request *ModifyHostAccountRequest) (_result *ModifyHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.ModifyHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name or description of an asset group.
//
// @param request - ModifyHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostGroupResponse
func (client *Client) ModifyHostGroupWithOptions(request *ModifyHostGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name or description of an asset group.
//
// @param request - ModifyHostGroupRequest
//
// @return ModifyHostGroupResponse
func (client *Client) ModifyHostGroup(request *ModifyHostGroupRequest) (_result *ModifyHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.ModifyHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a host shared key.
//
// @param request - ModifyHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostShareKeyResponse
func (client *Client) ModifyHostShareKeyWithOptions(request *ModifyHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostShareKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.HostShareKeyName) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a host shared key.
//
// @param request - ModifyHostShareKeyRequest
//
// @return ModifyHostShareKeyResponse
func (client *Client) ModifyHostShareKey(request *ModifyHostShareKeyRequest) (_result *ModifyHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.ModifyHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the endpoint type used for host O&M. Bastionhost supports O&M over the Internet and private networks.
//
// @param request - ModifyHostsActiveAddressTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostsActiveAddressTypeResponse
func (client *Client) ModifyHostsActiveAddressTypeWithOptions(request *ModifyHostsActiveAddressTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostsActiveAddressType"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the endpoint type used for host O&M. Bastionhost supports O&M over the Internet and private networks.
//
// @param request - ModifyHostsActiveAddressTypeRequest
//
// @return ModifyHostsActiveAddressTypeResponse
func (client *Client) ModifyHostsActiveAddressType(request *ModifyHostsActiveAddressTypeRequest) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.ModifyHostsActiveAddressTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the ports used by Bastionhost for the RDP and SSH protocols on hosts. By default, Bastionhost uses port 3389 for RDP and port 22 for SSH. If you have configured custom ports on your hosts, you can call this operation to update the port settings.
//
// Description:
//
// ### Usage notes
//
// This operation modifies the ports of O&M protocols for multiple hosts. If you consider the standard ports of O&M protocols insecure and vulnerable to attacks, such as port 22 for SSH, you can call this operation to specify custom ports.
//
// > Ports 0 to 1024 are reserved for Bastionhost. Do not specify a reserved port as a custom port.
//
// ### QPS limits
//
// You can call this operation up to 10 times per second. Requests that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyHostsPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostsPortResponse
func (client *Client) ModifyHostsPortWithOptions(request *ModifyHostsPortRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostsPortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostsPort"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the ports used by Bastionhost for the RDP and SSH protocols on hosts. By default, Bastionhost uses port 3389 for RDP and port 22 for SSH. If you have configured custom ports on your hosts, you can call this operation to update the port settings.
//
// Description:
//
// ### Usage notes
//
// This operation modifies the ports of O&M protocols for multiple hosts. If you consider the standard ports of O&M protocols insecure and vulnerable to attacks, such as port 22 for SSH, you can call this operation to specify custom ports.
//
// > Ports 0 to 1024 are reserved for Bastionhost. Do not specify a reserved port as a custom port.
//
// ### QPS limits
//
// You can call this operation up to 10 times per second. Requests that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyHostsPortRequest
//
// @return ModifyHostsPortResponse
func (client *Client) ModifyHostsPort(request *ModifyHostsPortRequest) (_result *ModifyHostsPortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.ModifyHostsPortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves hosts to a different project in a batch.
//
// @param request - ModifyHostsProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostsProjectResponse
func (client *Client) ModifyHostsProjectWithOptions(request *ModifyHostsProjectRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostsProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetProjectId) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostsProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves hosts to a different project in a batch.
//
// @param request - ModifyHostsProjectRequest
//
// @return ModifyHostsProjectResponse
func (client *Client) ModifyHostsProject(request *ModifyHostsProjectRequest) (_result *ModifyHostsProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostsProjectResponse{}
	_body, _err := client.ModifyHostsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the Active Directory (AD) authentication server for a Bastionhost instance.
//
// @param request - ModifyInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceADAuthServerResponse
func (client *Client) ModifyInstanceADAuthServerWithOptions(request *ModifyInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EmailMapping) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.IsSyncEmailAttr) {
		query["IsSyncEmailAttr"] = request.IsSyncEmailAttr
	}

	if !dara.IsNil(request.IsSyncMobileAttr) {
		query["IsSyncMobileAttr"] = request.IsSyncMobileAttr
	}

	if !dara.IsNil(request.IsSyncNameAttr) {
		query["IsSyncNameAttr"] = request.IsSyncNameAttr
	}

	if !dara.IsNil(request.IsSyncOuAsUserGroup) {
		query["IsSyncOuAsUserGroup"] = request.IsSyncOuAsUserGroup
	}

	if !dara.IsNil(request.MobileMapping) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !dara.IsNil(request.NameMapping) {
		query["NameMapping"] = request.NameMapping
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.ServerName) {
		query["ServerName"] = request.ServerName
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	if !dara.IsNil(request.SyncInterval) {
		query["SyncInterval"] = request.SyncInterval
	}

	if !dara.IsNil(request.UserSourceId) {
		query["UserSourceId"] = request.UserSourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the Active Directory (AD) authentication server for a Bastionhost instance.
//
// @param request - ModifyInstanceADAuthServerRequest
//
// @return ModifyInstanceADAuthServerResponse
func (client *Client) ModifyInstanceADAuthServer(request *ModifyInstanceADAuthServerRequest) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.ModifyInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a specified bastion host instance.
//
// Description:
//
// You can call this operation to purchase an Encryption Service instance. This is a prepaid product, and the operation deducts the fee synchronously.
//
// @param request - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a specified bastion host instance.
//
// Description:
//
// You can call this operation to purchase an Encryption Service instance. This is a prepaid product, and the operation deducts the fee synchronously.
//
// @param request - ModifyInstanceAttributeRequest
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the LDAP authentication server for a Bastionhost instance.
//
// @param request - ModifyInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceLDAPAuthServerResponse
func (client *Client) ModifyInstanceLDAPAuthServerWithOptions(request *ModifyInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.EmailMapping) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.LoginNameMapping) {
		query["LoginNameMapping"] = request.LoginNameMapping
	}

	if !dara.IsNil(request.MobileMapping) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !dara.IsNil(request.NameMapping) {
		query["NameMapping"] = request.NameMapping
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the LDAP authentication server for a Bastionhost instance.
//
// @param request - ModifyInstanceLDAPAuthServerRequest
//
// @return ModifyInstanceLDAPAuthServerResponse
func (client *Client) ModifyInstanceLDAPAuthServer(request *ModifyInstanceLDAPAuthServerRequest) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.ModifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the two-factor authentication settings of a bastion host.
//
// @param request - ModifyInstanceTwoFactorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceTwoFactorResponse
func (client *Client) ModifyInstanceTwoFactorWithOptions(request *ModifyInstanceTwoFactorRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableTwoFactor) {
		query["EnableTwoFactor"] = request.EnableTwoFactor
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SkipTwoFactorTime) {
		query["SkipTwoFactorTime"] = request.SkipTwoFactorTime
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceTwoFactor"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the two-factor authentication settings of a bastion host.
//
// @param request - ModifyInstanceTwoFactorRequest
//
// @return ModifyInstanceTwoFactorResponse
func (client *Client) ModifyInstanceTwoFactor(request *ModifyInstanceTwoFactorRequest) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.ModifyInstanceTwoFactorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a network domain.
//
// @param request - ModifyNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkDomainResponse
func (client *Client) ModifyNetworkDomainWithOptions(request *ModifyNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Proxies) {
		query["Proxies"] = request.Proxies
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a network domain.
//
// @param request - ModifyNetworkDomainRequest
//
// @return ModifyNetworkDomainResponse
func (client *Client) ModifyNetworkDomain(request *ModifyNetworkDomainRequest) (_result *ModifyNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNetworkDomainResponse{}
	_body, _err := client.ModifyNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a control policy.
//
// @param request - ModifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicyWithOptions(request *ModifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicy"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of a control policy.
//
// @param request - ModifyPolicyRequest
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicy(request *ModifyPolicyRequest) (_result *ModifyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.ModifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a project.
//
// @param request - ModifyProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProjectResponse
func (client *Client) ModifyProjectWithOptions(request *ModifyProjectRequest, runtime *dara.RuntimeOptions) (_result *ModifyProjectResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProject"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a project.
//
// @param request - ModifyProjectRequest
//
// @return ModifyProjectResponse
func (client *Client) ModifyProject(request *ModifyProjectRequest) (_result *ModifyProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyProjectResponse{}
	_body, _err := client.ModifyProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies project authorization.
//
// @param request - ModifyProjectAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProjectAuthorizationResponse
func (client *Client) ModifyProjectAuthorizationWithOptions(request *ModifyProjectAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ModifyProjectAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProjectAuthorization"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProjectAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies project authorization.
//
// @param request - ModifyProjectAuthorizationRequest
//
// @return ModifyProjectAuthorizationResponse
func (client *Client) ModifyProjectAuthorization(request *ModifyProjectAuthorizationRequest) (_result *ModifyProjectAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyProjectAuthorizationResponse{}
	_body, _err := client.ModifyProjectAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the notification recipient of a project.
//
// @param request - ModifyProjectNotifyReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProjectNotifyReceiverResponse
func (client *Client) ModifyProjectNotifyReceiverWithOptions(request *ModifyProjectNotifyReceiverRequest, runtime *dara.RuntimeOptions) (_result *ModifyProjectNotifyReceiverResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProjectNotifyReceiver"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProjectNotifyReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the notification recipient of a project.
//
// @param request - ModifyProjectNotifyReceiverRequest
//
// @return ModifyProjectNotifyReceiverResponse
func (client *Client) ModifyProjectNotifyReceiver(request *ModifyProjectNotifyReceiverRequest) (_result *ModifyProjectNotifyReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyProjectNotifyReceiverResponse{}
	_body, _err := client.ModifyProjectNotifyReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an audit policy.
//
// @param request - ModifyReviewRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReviewRuleResponse
func (client *Client) ModifyReviewRuleWithOptions(request *ModifyReviewRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyReviewRuleResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReviewRuleId) {
		query["ReviewRuleId"] = request.ReviewRuleId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyReviewRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyReviewRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an audit policy.
//
// @param request - ModifyReviewRuleRequest
//
// @return ModifyReviewRuleResponse
func (client *Client) ModifyReviewRule(request *ModifyReviewRuleRequest) (_result *ModifyReviewRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyReviewRuleResponse{}
	_body, _err := client.ModifyReviewRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information of an authorization rule.
//
// @param request - ModifyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleResponse
func (client *Client) ModifyRuleWithOptions(request *ModifyRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an authorization rule.
//
// @param request - ModifyRuleRequest
//
// @return ModifyRuleResponse
func (client *Client) ModifyRule(request *ModifyRuleRequest) (_result *ModifyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRuleResponse{}
	_body, _err := client.ModifyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a user of a bastion host.
//
// @param request - ModifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserResponse
func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LanguageStatus) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.MobileCountryCode) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !dara.IsNil(request.NeedResetPassword) {
		query["NeedResetPassword"] = request.NeedResetPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !dara.IsNil(request.TwoFactorStatus) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a user of a bastion host.
//
// @param request - ModifyUserRequest
//
// @return ModifyUserResponse
func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a user group.
//
// @param request - ModifyUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserGroupResponse
func (client *Client) ModifyUserGroupWithOptions(request *ModifyUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a user group.
//
// @param request - ModifyUserGroupRequest
//
// @return ModifyUserGroupResponse
func (client *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (_result *ModifyUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.ModifyUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the public key of the user.
//
// @param request - ModifyUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserPublicKeyResponse
func (client *Client) ModifyUserPublicKeyWithOptions(request *ModifyUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	if !dara.IsNil(request.PublicKeyId) {
		query["PublicKeyId"] = request.PublicKeyId
	}

	if !dara.IsNil(request.PublicKeyName) {
		query["PublicKeyName"] = request.PublicKeyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the public key of the user.
//
// @param request - ModifyUserPublicKeyRequest
//
// @return ModifyUserPublicKeyResponse
func (client *Client) ModifyUserPublicKey(request *ModifyUserPublicKeyRequest) (_result *ModifyUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserPublicKeyResponse{}
	_body, _err := client.ModifyUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves multiple databases to a specified network domain.
//
// @param request - MoveDatabasesToNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDatabasesToNetworkDomainResponse
func (client *Client) MoveDatabasesToNetworkDomainWithOptions(request *MoveDatabasesToNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *MoveDatabasesToNetworkDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveDatabasesToNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveDatabasesToNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves multiple databases to a specified network domain.
//
// @param request - MoveDatabasesToNetworkDomainRequest
//
// @return MoveDatabasesToNetworkDomainResponse
func (client *Client) MoveDatabasesToNetworkDomain(request *MoveDatabasesToNetworkDomainRequest) (_result *MoveDatabasesToNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveDatabasesToNetworkDomainResponse{}
	_body, _err := client.MoveDatabasesToNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves multiple hosts to a specified network domain.
//
// @param request - MoveHostsToNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveHostsToNetworkDomainResponse
func (client *Client) MoveHostsToNetworkDomainWithOptions(request *MoveHostsToNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *MoveHostsToNetworkDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveHostsToNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveHostsToNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves multiple hosts to a specified network domain.
//
// @param request - MoveHostsToNetworkDomainRequest
//
// @return MoveHostsToNetworkDomainResponse
func (client *Client) MoveHostsToNetworkDomain(request *MoveHostsToNetworkDomainRequest) (_result *MoveHostsToNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveHostsToNetworkDomainResponse{}
	_body, _err := client.MoveHostsToNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a Bastionhost instance to a specified resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a Bastionhost instance to a specified resource group.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects a command execution request. If an O&M engineer executes a command configured in a command approval policy, the administrator receives an approval request in the Bastionhost console. The command can be executed only after the administrator approves it. If the request is rejected, the command cannot be executed.
//
// Description:
//
// This API operation allows administrators to reject command execution requests from O&M engineers.
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency.
//
// @param request - RejectApproveCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectApproveCommandResponse
func (client *Client) RejectApproveCommandWithOptions(request *RejectApproveCommandRequest, runtime *dara.RuntimeOptions) (_result *RejectApproveCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandId) {
		query["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectApproveCommand"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects a command execution request. If an O&M engineer executes a command configured in a command approval policy, the administrator receives an approval request in the Bastionhost console. The command can be executed only after the administrator approves it. If the request is rejected, the command cannot be executed.
//
// Description:
//
// This API operation allows administrators to reject command execution requests from O&M engineers.
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency.
//
// @param request - RejectApproveCommandRequest
//
// @return RejectApproveCommandResponse
func (client *Client) RejectApproveCommand(request *RejectApproveCommandRequest) (_result *RejectApproveCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.RejectApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects an O&M task execution request.
//
// @param request - RejectOperationTaskApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectOperationTaskApprovalResponse
func (client *Client) RejectOperationTaskApprovalWithOptions(request *RejectOperationTaskApprovalRequest, runtime *dara.RuntimeOptions) (_result *RejectOperationTaskApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveId) {
		query["ApproveId"] = request.ApproveId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectOperationTaskApproval"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectOperationTaskApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects an O&M task execution request.
//
// @param request - RejectOperationTaskApprovalRequest
//
// @return RejectOperationTaskApprovalResponse
func (client *Client) RejectOperationTaskApproval(request *RejectOperationTaskApprovalRequest) (_result *RejectOperationTaskApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectOperationTaskApprovalResponse{}
	_body, _err := client.RejectOperationTaskApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects an O&M application. If O&M approval is enabled in the control policy configured by the administrator, O&M engineers must apply for O&M and obtain administrator approval before logging on to assets.
//
// Description:
//
// Rejects an O&M application submitted by an O&M engineer.
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RejectOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectOperationTicketResponse
func (client *Client) RejectOperationTicketWithOptions(request *RejectOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *RejectOperationTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationTicketId) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects an O&M application. If O&M approval is enabled in the control policy configured by the administrator, O&M engineers must apply for O&M and obtain administrator approval before logging on to assets.
//
// Description:
//
// Rejects an O&M application submitted by an O&M engineer.
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RejectOperationTicketRequest
//
// @return RejectOperationTicketResponse
func (client *Client) RejectOperationTicket(request *RejectOperationTicketRequest) (_result *RejectOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.RejectOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple databases from an asset group at a time.
//
// @param request - RemoveDatabasesFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDatabasesFromGroupResponse
func (client *Client) RemoveDatabasesFromGroupWithOptions(request *RemoveDatabasesFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveDatabasesFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDatabasesFromGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDatabasesFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple databases from an asset group at a time.
//
// @param request - RemoveDatabasesFromGroupRequest
//
// @return RemoveDatabasesFromGroupResponse
func (client *Client) RemoveDatabasesFromGroup(request *RemoveDatabasesFromGroupRequest) (_result *RemoveDatabasesFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDatabasesFromGroupResponse{}
	_body, _err := client.RemoveDatabasesFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple hosts from a specified host group.
//
// Description:
//
// This operation removes multiple hosts from a host group. If some hosts no longer require unified management through a host group, you can use this operation to remove them.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - RemoveHostsFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveHostsFromGroupResponse
func (client *Client) RemoveHostsFromGroupWithOptions(request *RemoveHostsFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveHostsFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveHostsFromGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple hosts from a specified host group.
//
// Description:
//
// This operation removes multiple hosts from a host group. If some hosts no longer require unified management through a host group, you can use this operation to remove them.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - RemoveHostsFromGroupRequest
//
// @return RemoveHostsFromGroupResponse
func (client *Client) RemoveHostsFromGroup(request *RemoveHostsFromGroupRequest) (_result *RemoveHostsFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.RemoveHostsFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Resource Directory (RD) member account from a Bastionhost instance. You must remove all assets of the member account from Bastionhost before you call this operation.
//
// Description:
//
// Before removing a member account, you must purge the ECS, RDS, and PolarDB instance assets that belong to the member account from Bastionhost.
//
// @param request - RemoveInstanceRdMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstanceRdMemberResponse
func (client *Client) RemoveInstanceRdMemberWithOptions(request *RemoveInstanceRdMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstanceRdMemberResponse, _err error) {
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

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstanceRdMember"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstanceRdMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a Resource Directory (RD) member account from a Bastionhost instance. You must remove all assets of the member account from Bastionhost before you call this operation.
//
// Description:
//
// Before removing a member account, you must purge the ECS, RDS, and PolarDB instance assets that belong to the member account from Bastionhost.
//
// @param request - RemoveInstanceRdMemberRequest
//
// @return RemoveInstanceRdMemberResponse
func (client *Client) RemoveInstanceRdMember(request *RemoveInstanceRdMemberRequest) (_result *RemoveInstanceRdMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveInstanceRdMemberResponse{}
	_body, _err := client.RemoveInstanceRdMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple users from a user group.
//
// Description:
//
// ### Usage notes
//
// This operation removes multiple users from a user group. You can call this operation to remove users from a user group in batches when users resign or transfer to other departments and their permissions need to be revoked, or when users switch to another user group.
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
//
// @param request - RemoveUsersFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithOptions(request *RemoveUsersFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsersFromGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple users from a user group.
//
// Description:
//
// ### Usage notes
//
// This operation removes multiple users from a user group. You can call this operation to remove users from a user group in batches when users resign or transfer to other departments and their permissions need to be revoked, or when users switch to another user group.
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
//
// @param request - RemoveUsersFromGroupRequest
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroup(request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews an O&M token. Each call extends the validity of the O&M token by 1 hour.
//
// @param request - RenewAssetOperationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAssetOperationTokenResponse
func (client *Client) RenewAssetOperationTokenWithOptions(request *RenewAssetOperationTokenRequest, runtime *dara.RuntimeOptions) (_result *RenewAssetOperationTokenResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.TokenId) {
		query["TokenId"] = request.TokenId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAssetOperationToken"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAssetOperationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews an O&M token. Each call extends the validity of the O&M token by 1 hour.
//
// @param request - RenewAssetOperationTokenRequest
//
// @return RenewAssetOperationTokenResponse
func (client *Client) RenewAssetOperationToken(request *RenewAssetOperationTokenRequest) (_result *RenewAssetOperationTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAssetOperationTokenResponse{}
	_body, _err := client.RenewAssetOperationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purges the logon credentials (password or SSH private key) of a specified host account.
//
// @param request - ResetHostAccountCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetHostAccountCredentialResponse
func (client *Client) ResetHostAccountCredentialWithOptions(request *ResetHostAccountCredentialRequest, runtime *dara.RuntimeOptions) (_result *ResetHostAccountCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialType) {
		query["CredentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetHostAccountCredential"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purges the logon credentials (password or SSH private key) of a specified host account.
//
// @param request - ResetHostAccountCredentialRequest
//
// @return ResetHostAccountCredentialResponse
func (client *Client) ResetHostAccountCredential(request *ResetHostAccountCredentialRequest) (_result *ResetHostAccountCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.ResetHostAccountCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the logon time limit settings for a specified control policy.
//
// @param tmpReq - SetPolicyAccessTimeRangeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyAccessTimeRangeConfigResponse
func (client *Client) SetPolicyAccessTimeRangeConfigWithOptions(tmpReq *SetPolicyAccessTimeRangeConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyAccessTimeRangeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetPolicyAccessTimeRangeConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessTimeRangeConfig) {
		request.AccessTimeRangeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessTimeRangeConfig, dara.String("AccessTimeRangeConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTimeRangeConfigShrink) {
		query["AccessTimeRangeConfig"] = request.AccessTimeRangeConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyAccessTimeRangeConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyAccessTimeRangeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the logon time limit settings for a specified control policy.
//
// @param request - SetPolicyAccessTimeRangeConfigRequest
//
// @return SetPolicyAccessTimeRangeConfigResponse
func (client *Client) SetPolicyAccessTimeRangeConfig(request *SetPolicyAccessTimeRangeConfigRequest) (_result *SetPolicyAccessTimeRangeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyAccessTimeRangeConfigResponse{}
	_body, _err := client.SetPolicyAccessTimeRangeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the O&M approval configuration for a specified control policy.
//
// @param tmpReq - SetPolicyApprovalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyApprovalConfigResponse
func (client *Client) SetPolicyApprovalConfigWithOptions(tmpReq *SetPolicyApprovalConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyApprovalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetPolicyApprovalConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovalConfig) {
		request.ApprovalConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovalConfig, dara.String("ApprovalConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalConfigShrink) {
		query["ApprovalConfig"] = request.ApprovalConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyApprovalConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyApprovalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the O&M approval configuration for a specified control policy.
//
// @param request - SetPolicyApprovalConfigRequest
//
// @return SetPolicyApprovalConfigResponse
func (client *Client) SetPolicyApprovalConfig(request *SetPolicyApprovalConfigRequest) (_result *SetPolicyApprovalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyApprovalConfigResponse{}
	_body, _err := client.SetPolicyApprovalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the effective asset scope for a specified control policy.
//
// @param request - SetPolicyAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyAssetScopeResponse
func (client *Client) SetPolicyAssetScopeWithOptions(request *SetPolicyAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyAssetScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyAssetScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyAssetScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the effective asset scope for a specified control policy.
//
// @param request - SetPolicyAssetScopeRequest
//
// @return SetPolicyAssetScopeResponse
func (client *Client) SetPolicyAssetScope(request *SetPolicyAssetScopeRequest) (_result *SetPolicyAssetScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyAssetScopeResponse{}
	_body, _err := client.SetPolicyAssetScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the commands that are allowed or denied, and the commands that require approval for the users and hosts to which a specified policy applies.
//
// @param tmpReq - SetPolicyCommandConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyCommandConfigResponse
func (client *Client) SetPolicyCommandConfigWithOptions(tmpReq *SetPolicyCommandConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyCommandConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetPolicyCommandConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CommandConfig) {
		request.CommandConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommandConfig, dara.String("CommandConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandConfigShrink) {
		query["CommandConfig"] = request.CommandConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyCommandConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyCommandConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the commands that are allowed or denied, and the commands that require approval for the users and hosts to which a specified policy applies.
//
// @param request - SetPolicyCommandConfigRequest
//
// @return SetPolicyCommandConfigResponse
func (client *Client) SetPolicyCommandConfig(request *SetPolicyCommandConfigRequest) (_result *SetPolicyCommandConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyCommandConfigResponse{}
	_body, _err := client.SetPolicyCommandConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures whether source IP addresses can access hosts to which the current policy applies.
//
// @param tmpReq - SetPolicyIPAclConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyIPAclConfigResponse
func (client *Client) SetPolicyIPAclConfigWithOptions(tmpReq *SetPolicyIPAclConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyIPAclConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetPolicyIPAclConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IPAclConfig) {
		request.IPAclConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IPAclConfig, dara.String("IPAclConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IPAclConfigShrink) {
		query["IPAclConfig"] = request.IPAclConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyIPAclConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyIPAclConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures whether source IP addresses can access hosts to which the current policy applies.
//
// @param request - SetPolicyIPAclConfigRequest
//
// @return SetPolicyIPAclConfigResponse
func (client *Client) SetPolicyIPAclConfig(request *SetPolicyIPAclConfigRequest) (_result *SetPolicyIPAclConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyIPAclConfigResponse{}
	_body, _err := client.SetPolicyIPAclConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the RDP, SSH, and SFTP options of a control policy.
//
// @param tmpReq - SetPolicyProtocolConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyProtocolConfigResponse
func (client *Client) SetPolicyProtocolConfigWithOptions(tmpReq *SetPolicyProtocolConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyProtocolConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetPolicyProtocolConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProtocolConfig) {
		request.ProtocolConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProtocolConfig, dara.String("ProtocolConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProtocolConfigShrink) {
		query["ProtocolConfig"] = request.ProtocolConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyProtocolConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyProtocolConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the RDP, SSH, and SFTP options of a control policy.
//
// @param request - SetPolicyProtocolConfigRequest
//
// @return SetPolicyProtocolConfigResponse
func (client *Client) SetPolicyProtocolConfig(request *SetPolicyProtocolConfigRequest) (_result *SetPolicyProtocolConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyProtocolConfigResponse{}
	_body, _err := client.SetPolicyProtocolConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the user scope for a specified control policy.
//
// @param request - SetPolicyUserScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyUserScopeResponse
func (client *Client) SetPolicyUserScopeWithOptions(request *SetPolicyUserScopeRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyUserScopeResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyUserScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyUserScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the user scope for a specified control policy.
//
// @param request - SetPolicyUserScopeRequest
//
// @return SetPolicyUserScopeResponse
func (client *Client) SetPolicyUserScope(request *SetPolicyUserScopeRequest) (_result *SetPolicyUserScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyUserScopeResponse{}
	_body, _err := client.SetPolicyUserScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a specified Bastionhost instance.
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientSecurityGroupIds) {
		query["ClientSecurityGroupIds"] = request.ClientSecurityGroupIds
	}

	if !dara.IsNil(request.EnablePortalPrivateAccess) {
		query["EnablePortalPrivateAccess"] = request.EnablePortalPrivateAccess
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !dara.IsNil(request.SlaveVswitchId) {
		query["SlaveVswitchId"] = request.SlaveVswitchId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a specified Bastionhost instance.
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and binds tags for specified Bastionhost instances.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and binds tags for specified Bastionhost instances.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks one or more users of a bastion host.
//
// Description:
//
// Unlocks previously locked bastion host users, allowing them to resume O\\&M operations. Use this operation after locking users with the [LockUsers](https://help.aliyun.com/document_detail/204591.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *dara.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockUsers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks one or more users of a bastion host.
//
// Description:
//
// Unlocks previously locked bastion host users, allowing them to resume O\\&M operations. Use this operation after locking users with the [LockUsers](https://help.aliyun.com/document_detail/204591.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockUsersRequest
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsers(request *UnlockUsersRequest) (_result *UnlockUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockUsersResponse{}
	_body, _err := client.UnlockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds and deletes tags from a specified Bastionhost instance in batches.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds and deletes tags from a specified Bastionhost instance in batches.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the connectivity to the Active Directory (AD) authentication server configured for a Bastionhost instance. Use this operation to validate your AD settings before you save them.
//
// @param request - VerifyInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyInstanceADAuthServerResponse
func (client *Client) VerifyInstanceADAuthServerWithOptions(request *VerifyInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the connectivity to the Active Directory (AD) authentication server configured for a Bastionhost instance. Use this operation to validate your AD settings before you save them.
//
// @param request - VerifyInstanceADAuthServerRequest
//
// @return VerifyInstanceADAuthServerResponse
func (client *Client) VerifyInstanceADAuthServer(request *VerifyInstanceADAuthServerRequest) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.VerifyInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the connectivity to the Lightweight Directory Access Protocol (LDAP) authentication server configured for a Bastionhost instance. Use this operation to validate your LDAP settings before you save them.
//
// @param request - VerifyInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyInstanceLDAPAuthServerResponse
func (client *Client) VerifyInstanceLDAPAuthServerWithOptions(request *VerifyInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the connectivity to the Lightweight Directory Access Protocol (LDAP) authentication server configured for a Bastionhost instance. Use this operation to validate your LDAP settings before you save them.
//
// @param request - VerifyInstanceLDAPAuthServerRequest
//
// @return VerifyInstanceLDAPAuthServerResponse
func (client *Client) VerifyInstanceLDAPAuthServer(request *VerifyInstanceLDAPAuthServerRequest) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.VerifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
