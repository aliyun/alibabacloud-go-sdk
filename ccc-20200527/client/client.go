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
		"ap-northeast-1":        dara.String("ccc.aliyuncs.com"),
		"ap-south-1":            dara.String("ccc.aliyuncs.com"),
		"ap-southeast-1":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-2":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-3":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-5":        dara.String("ccc.aliyuncs.com"),
		"cn-beijing":            dara.String("ccc.aliyuncs.com"),
		"cn-chengdu":            dara.String("ccc.aliyuncs.com"),
		"cn-hongkong":           dara.String("ccc.aliyuncs.com"),
		"cn-huhehaote":          dara.String("ccc.aliyuncs.com"),
		"cn-qingdao":            dara.String("ccc.aliyuncs.com"),
		"cn-shenzhen":           dara.String("ccc.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("ccc.aliyuncs.com"),
		"eu-central-1":          dara.String("ccc.aliyuncs.com"),
		"eu-west-1":             dara.String("ccc.aliyuncs.com"),
		"me-east-1":             dara.String("ccc.aliyuncs.com"),
		"us-east-1":             dara.String("ccc.aliyuncs.com"),
		"us-west-1":             dara.String("ccc.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("ccc.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("ccc.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("ccc.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("ccc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ccc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CheckServiceLinkedRoleForDeletingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleForDeletingResponse
func (client *Client) CheckServiceLinkedRoleForDeletingWithOptions(request *CheckServiceLinkedRoleForDeletingRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionTaskId) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SPIRegionId) {
		query["SPIRegionId"] = request.SPIRegionId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRoleForDeleting"),
		Version:     dara.String("2020-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckServiceLinkedRoleForDeletingRequest
//
// @return CheckServiceLinkedRoleForDeletingResponse
func (client *Client) CheckServiceLinkedRoleForDeleting(request *CheckServiceLinkedRoleForDeletingRequest) (_result *CheckServiceLinkedRoleForDeletingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleForDeletingResponse{}
	_body, _err := client.CheckServiceLinkedRoleForDeletingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
