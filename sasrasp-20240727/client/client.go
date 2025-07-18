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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("sasrasp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # DescribeAttackProtectionCount
//
// @param request - DescribeAttackProtectionCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttackProtectionCountResponse
func (client *Client) DescribeAttackProtectionCountWithOptions(request *DescribeAttackProtectionCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttackProtectionCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAttackProtectionCount"),
		Version:     dara.String("2024-07-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAttackProtectionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeAttackProtectionCount
//
// @param request - DescribeAttackProtectionCountRequest
//
// @return DescribeAttackProtectionCountResponse
func (client *Client) DescribeAttackProtectionCount(request *DescribeAttackProtectionCountRequest) (_result *DescribeAttackProtectionCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAttackProtectionCountResponse{}
	_body, _err := client.DescribeAttackProtectionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询攻击项
//
// @param request - DescribeAttacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttacksResponse
func (client *Client) DescribeAttacksWithOptions(request *DescribeAttacksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttacksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AttackHostId) {
		query["AttackHostId"] = request.AttackHostId
	}

	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.AttackUrl) {
		query["AttackUrl"] = request.AttackUrl
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.HandlerType) {
		query["HandlerType"] = request.HandlerType
	}

	if !dara.IsNil(request.Hostname) {
		query["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pid) {
		query["Pid"] = request.Pid
	}

	if !dara.IsNil(request.RaspType) {
		query["RaspType"] = request.RaspType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Remote) {
		query["Remote"] = request.Remote
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.UnionId) {
		query["UnionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAttacks"),
		Version:     dara.String("2024-07-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAttacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询攻击项
//
// @param request - DescribeAttacksRequest
//
// @return DescribeAttacksResponse
func (client *Client) DescribeAttacks(request *DescribeAttacksRequest) (_result *DescribeAttacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAttacksResponse{}
	_body, _err := client.DescribeAttacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
