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

	if !dara.IsNil(request.DeviceCount) {
		query["DeviceCount"] = request.DeviceCount
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
// 删除指标平台
//
// @param request - DeleteAgentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentPlatformResponse
func (client *Client) DeleteAgentPlatformWithOptions(request *DeleteAgentPlatformRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentPlatformResponse, _err error) {
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
		Action:      dara.String("DeleteAgentPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指标平台
//
// @param request - DeleteAgentPlatformRequest
//
// @return DeleteAgentPlatformResponse
func (client *Client) DeleteAgentPlatform(request *DeleteAgentPlatformRequest) (_result *DeleteAgentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAgentPlatformResponse{}
	_body, _err := client.DeleteAgentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除具身智能平台
//
// @param request - DeleteEmbodiedAIPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEmbodiedAIPlatformResponse
func (client *Client) DeleteEmbodiedAIPlatformWithOptions(request *DeleteEmbodiedAIPlatformRequest, runtime *dara.RuntimeOptions) (_result *DeleteEmbodiedAIPlatformResponse, _err error) {
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
		Action:      dara.String("DeleteEmbodiedAIPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEmbodiedAIPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除具身智能平台
//
// @param request - DeleteEmbodiedAIPlatformRequest
//
// @return DeleteEmbodiedAIPlatformResponse
func (client *Client) DeleteEmbodiedAIPlatform(request *DeleteEmbodiedAIPlatformRequest) (_result *DeleteEmbodiedAIPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEmbodiedAIPlatformResponse{}
	_body, _err := client.DeleteEmbodiedAIPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对ADB-MySQL提供产品RAG检索和实例分析、运维诊断
//
// @param request - DescribeChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChatMessageResponse
func (client *Client) DescribeChatMessageWithSSE(request *DescribeChatMessageRequest, runtime *dara.RuntimeOptions, _yield chan *DescribeChatMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.describeChatMessageWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// 对ADB-MySQL提供产品RAG检索和实例分析、运维诊断
//
// @param request - DescribeChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChatMessageResponse
func (client *Client) DescribeChatMessageWithOptions(request *DescribeChatMessageRequest, runtime *dara.RuntimeOptions) (_result *DescribeChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChatMessage"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对ADB-MySQL提供产品RAG检索和实例分析、运维诊断
//
// @param request - DescribeChatMessageRequest
//
// @return DescribeChatMessageResponse
func (client *Client) DescribeChatMessage(request *DescribeChatMessageRequest) (_result *DescribeChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChatMessageResponse{}
	_body, _err := client.DescribeChatMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询具身智能平台设备资源分配方案
//
// @param request - DescribeEapDeviceResourceAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEapDeviceResourceAllocationResponse
func (client *Client) DescribeEapDeviceResourceAllocationWithOptions(request *DescribeEapDeviceResourceAllocationRequest, runtime *dara.RuntimeOptions) (_result *DescribeEapDeviceResourceAllocationResponse, _err error) {
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

	if !dara.IsNil(request.DeviceCount) {
		query["DeviceCount"] = request.DeviceCount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEapDeviceResourceAllocation"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEapDeviceResourceAllocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询具身智能平台设备资源分配方案
//
// @param request - DescribeEapDeviceResourceAllocationRequest
//
// @return DescribeEapDeviceResourceAllocationResponse
func (client *Client) DescribeEapDeviceResourceAllocation(request *DescribeEapDeviceResourceAllocationRequest) (_result *DescribeEapDeviceResourceAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEapDeviceResourceAllocationResponse{}
	_body, _err := client.DescribeEapDeviceResourceAllocationWithOptions(request, runtime)
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
// 解锁具身智能平台
//
// @param request - LockEmbodiedAIPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockEmbodiedAIPlatformResponse
func (client *Client) LockEmbodiedAIPlatformWithOptions(request *LockEmbodiedAIPlatformRequest, runtime *dara.RuntimeOptions) (_result *LockEmbodiedAIPlatformResponse, _err error) {
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
		Action:      dara.String("LockEmbodiedAIPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockEmbodiedAIPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解锁具身智能平台
//
// @param request - LockEmbodiedAIPlatformRequest
//
// @return LockEmbodiedAIPlatformResponse
func (client *Client) LockEmbodiedAIPlatform(request *LockEmbodiedAIPlatformRequest) (_result *LockEmbodiedAIPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockEmbodiedAIPlatformResponse{}
	_body, _err := client.LockEmbodiedAIPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改变配指标平台
//
// @param tmpReq - ModifyAgentPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAgentPlatformResponse
func (client *Client) ModifyAgentPlatformWithOptions(tmpReq *ModifyAgentPlatformRequest, runtime *dara.RuntimeOptions) (_result *ModifyAgentPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAgentPlatformShrinkRequest{}
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
		Action:      dara.String("ModifyAgentPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAgentPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改变配指标平台
//
// @param request - ModifyAgentPlatformRequest
//
// @return ModifyAgentPlatformResponse
func (client *Client) ModifyAgentPlatform(request *ModifyAgentPlatformRequest) (_result *ModifyAgentPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAgentPlatformResponse{}
	_body, _err := client.ModifyAgentPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变配具身智能平台
//
// @param tmpReq - ModifyEmbodiedAIPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEmbodiedAIPlatformResponse
func (client *Client) ModifyEmbodiedAIPlatformWithOptions(tmpReq *ModifyEmbodiedAIPlatformRequest, runtime *dara.RuntimeOptions) (_result *ModifyEmbodiedAIPlatformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyEmbodiedAIPlatformShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RayConfig) {
		request.RayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RayConfig, dara.String("RayConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DeviceCount) {
		query["DeviceCount"] = request.DeviceCount
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
		Action:      dara.String("ModifyEmbodiedAIPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEmbodiedAIPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变配具身智能平台
//
// @param request - ModifyEmbodiedAIPlatformRequest
//
// @return ModifyEmbodiedAIPlatformResponse
func (client *Client) ModifyEmbodiedAIPlatform(request *ModifyEmbodiedAIPlatformRequest) (_result *ModifyEmbodiedAIPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEmbodiedAIPlatformResponse{}
	_body, _err := client.ModifyEmbodiedAIPlatformWithOptions(request, runtime)
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

// Summary:
//
// 解锁具身智能平台
//
// @param request - UnlockEmbodiedAIPlatformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockEmbodiedAIPlatformResponse
func (client *Client) UnlockEmbodiedAIPlatformWithOptions(request *UnlockEmbodiedAIPlatformRequest, runtime *dara.RuntimeOptions) (_result *UnlockEmbodiedAIPlatformResponse, _err error) {
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
		Action:      dara.String("UnlockEmbodiedAIPlatform"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockEmbodiedAIPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解锁具身智能平台
//
// @param request - UnlockEmbodiedAIPlatformRequest
//
// @return UnlockEmbodiedAIPlatformResponse
func (client *Client) UnlockEmbodiedAIPlatform(request *UnlockEmbodiedAIPlatformRequest) (_result *UnlockEmbodiedAIPlatformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockEmbodiedAIPlatformResponse{}
	_body, _err := client.UnlockEmbodiedAIPlatformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) describeChatMessageWithSSE_opYieldFunc(_yield chan *DescribeChatMessageResponse, _yieldErr chan error, request *DescribeChatMessageRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChatMessage"),
		Version:     dara.String("2025-08-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
