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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ecs-workbench"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 查询Workbench终端配置
//
// Description:
//
// 查询Workbench终端配置
//
// @param request - DescribeTerminalSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTerminalSettingsResponse
func (client *Client) DescribeTerminalSettingsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeTerminalSettingsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTerminalSettings"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTerminalSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Workbench终端配置
//
// Description:
//
// 查询Workbench终端配置
//
// @return DescribeTerminalSettingsResponse
func (client *Client) DescribeTerminalSettings() (_result *DescribeTerminalSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTerminalSettingsResponse{}
	_body, _err := client.DescribeTerminalSettingsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Workbench终端配置
//
// Description:
//
// 修改Workbench终端配置
//
// @param tmpReq - ModifyTerminalSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTerminalSettingsResponse
func (client *Client) ModifyTerminalSettingsWithOptions(tmpReq *ModifyTerminalSettingsRequest, runtime *dara.RuntimeOptions) (_result *ModifyTerminalSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyTerminalSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PasswordlessLoginConfig) {
		request.PasswordlessLoginConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PasswordlessLoginConfig, dara.String("PasswordlessLoginConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PasswordlessLoginConfigShrink) {
		query["PasswordlessLoginConfig"] = request.PasswordlessLoginConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTerminalSettings"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTerminalSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Workbench终端配置
//
// Description:
//
// 修改Workbench终端配置
//
// @param request - ModifyTerminalSettingsRequest
//
// @return ModifyTerminalSettingsResponse
func (client *Client) ModifyTerminalSettings(request *ModifyTerminalSettingsRequest) (_result *ModifyTerminalSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTerminalSettingsResponse{}
	_body, _err := client.ModifyTerminalSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
