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
	client.Endpoint, _err = client.GetEndpoint(dara.String("adbai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建指标平台
//
// @param tmpReq - CreateAgentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentPlatformResponse
func (client *Client) CreateAgentPlatformWithOptions(tmpReq *CreateAgentPlatformRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentPlatformShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AiPlatformConfig) {
		request.AiPlatformConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AiPlatformConfig, dara.String("AiPlatformConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AiPlatformConfigShrink) {
		query["AiPlatformConfig"] = request.AiPlatformConfigShrink
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("CreateAgentPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建指标平台
//
// @param request - CreateAgentPlatformRequest
//
// @return CreateAgentPlatformResponse
func (client *Client) CreateAgentPlatform(request *CreateAgentPlatformRequest) (_result *CreateAgentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgentPlatformResponse{}
	_body, _err := client.CreateAgentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建具身智能平台
//
// @param tmpReq - CreateEmbodiedAIPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEmbodiedAIPlatformResponse
func (client *Client) CreateEmbodiedAIPlatformWithOptions(tmpReq *CreateEmbodiedAIPlatformRequest, runtime *dara.RuntimeOptions) (_result *CreateEmbodiedAIPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEmbodiedAIPlatformShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RayConfig) {
		request.RayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RayConfig, dara.String("RayConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PlatformName) {
		query["PlatformName"] = request.PlatformName
	}

	if !dara.IsNil(request.RayConfigShrink) {
		query["RayConfig"] = request.RayConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WebserverSpecName) {
		query["WebserverSpecName"] = request.WebserverSpecName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEmbodiedAIPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEmbodiedAIPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建具身智能平台
//
// @param request - CreateEmbodiedAIPlatformRequest
//
// @return CreateEmbodiedAIPlatformResponse
func (client *Client) CreateEmbodiedAIPlatform(request *CreateEmbodiedAIPlatformRequest) (_result *CreateEmbodiedAIPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEmbodiedAIPlatformResponse{}
	_body, _err := client.CreateEmbodiedAIPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询具身智能平台
//
// @param request - DescribeEmbodiedAIPlatformsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEmbodiedAIPlatformsResponse
func (client *Client) DescribeEmbodiedAIPlatformsWithOptions(request *DescribeEmbodiedAIPlatformsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEmbodiedAIPlatformsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlatformName) {
		query["PlatformName"] = request.PlatformName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEmbodiedAIPlatforms"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEmbodiedAIPlatformsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询具身智能平台
//
// @param request - DescribeEmbodiedAIPlatformsRequest
//
// @return DescribeEmbodiedAIPlatformsResponse
func (client *Client) DescribeEmbodiedAIPlatforms(request *DescribeEmbodiedAIPlatformsRequest) (_result *DescribeEmbodiedAIPlatformsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEmbodiedAIPlatformsResponse{}
	_body, _err := client.DescribeEmbodiedAIPlatformsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询具身智能平台资源用量
//
// @param request - GetEmbodiedAIPlatformResourceUsageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmbodiedAIPlatformResourceUsageInfoResponse
func (client *Client) GetEmbodiedAIPlatformResourceUsageInfoWithOptions(request *GetEmbodiedAIPlatformResourceUsageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetEmbodiedAIPlatformResourceUsageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PlatformName) {
		query["PlatformName"] = request.PlatformName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmbodiedAIPlatformResourceUsageInfo"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmbodiedAIPlatformResourceUsageInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询具身智能平台资源用量
//
// @param request - GetEmbodiedAIPlatformResourceUsageInfoRequest
//
// @return GetEmbodiedAIPlatformResourceUsageInfoResponse
func (client *Client) GetEmbodiedAIPlatformResourceUsageInfo(request *GetEmbodiedAIPlatformResourceUsageInfoRequest) (_result *GetEmbodiedAIPlatformResourceUsageInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEmbodiedAIPlatformResourceUsageInfoResponse{}
	_body, _err := client.GetEmbodiedAIPlatformResourceUsageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置具身智能平台密码
//
// @param request - ResetEmbodiedAIPlatformPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetEmbodiedAIPlatformPasswordResponse
func (client *Client) ResetEmbodiedAIPlatformPasswordWithOptions(request *ResetEmbodiedAIPlatformPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetEmbodiedAIPlatformPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PlatformName) {
		query["PlatformName"] = request.PlatformName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetEmbodiedAIPlatformPassword"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetEmbodiedAIPlatformPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置具身智能平台密码
//
// @param request - ResetEmbodiedAIPlatformPasswordRequest
//
// @return ResetEmbodiedAIPlatformPasswordResponse
func (client *Client) ResetEmbodiedAIPlatformPassword(request *ResetEmbodiedAIPlatformPasswordRequest) (_result *ResetEmbodiedAIPlatformPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetEmbodiedAIPlatformPasswordResponse{}
	_body, _err := client.ResetEmbodiedAIPlatformPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
