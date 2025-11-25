// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # DescribeAttackProtectionCount
//
// @param request - DescribeAttackProtectionCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttackProtectionCountResponse
func (client *Client) DescribeAttackProtectionCountWithContext(ctx context.Context, request *DescribeAttackProtectionCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttackProtectionCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttacksResponse
func (client *Client) DescribeAttacksWithContext(ctx context.Context, request *DescribeAttacksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
