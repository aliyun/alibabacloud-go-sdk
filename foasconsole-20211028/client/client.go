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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("foasconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 开通弹性计算
//
// @param tmpReq - ConvertHybridInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertHybridInstanceResponse
func (client *Client) ConvertHybridInstanceWithOptions(tmpReq *ConvertHybridInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertHybridInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConvertHybridInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		query["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertHybridInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertHybridInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通弹性计算
//
// @param request - ConvertHybridInstanceRequest
//
// @return ConvertHybridInstanceResponse
func (client *Client) ConvertHybridInstance(request *ConvertHybridInstanceRequest) (_result *ConvertHybridInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertHybridInstanceResponse{}
	_body, _err := client.ConvertHybridInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按量付费转包年包月
//
// @param tmpReq - ConvertInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstanceWithOptions(tmpReq *ConvertInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConvertInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NamespaceResourceSpecs) {
		request.NamespaceResourceSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NamespaceResourceSpecs, dara.String("NamespaceResourceSpecs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsAutoRenew) {
		body["IsAutoRenew"] = request.IsAutoRenew
	}

	if !dara.IsNil(request.NamespaceResourceSpecsShrink) {
		body["NamespaceResourceSpecs"] = request.NamespaceResourceSpecsShrink
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按量付费转包年包月
//
// @param request - ConvertInstanceRequest
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstance(request *ConvertInstanceRequest) (_result *ConvertInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.ConvertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 包年包月转按量付费
//
// @param request - ConvertPrepayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPrepayInstanceResponse
func (client *Client) ConvertPrepayInstanceWithOptions(request *ConvertPrepayInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertPrepayInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertPrepayInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertPrepayInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 包年包月转按量付费
//
// @param request - ConvertPrepayInstanceRequest
//
// @return ConvertPrepayInstanceResponse
func (client *Client) ConvertPrepayInstance(request *ConvertPrepayInstanceRequest) (_result *ConvertPrepayInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertPrepayInstanceResponse{}
	_body, _err := client.ConvertPrepayInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param tmpReq - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(tmpReq *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaResourceSpec) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, dara.String("HaResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HaVSwitchIds) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, dara.String("HaVSwitchIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Storage) {
		request.StorageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Storage, dara.String("Storage"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArchitectureType) {
		body["ArchitectureType"] = request.ArchitectureType
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Ha) {
		body["Ha"] = request.Ha
	}

	if !dara.IsNil(request.HaResourceSpecShrink) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !dara.IsNil(request.HaVSwitchIdsShrink) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MonitorType) {
		body["MonitorType"] = request.MonitorType
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	if !dara.IsNil(request.StorageShrink) {
		body["Storage"] = request.StorageShrink
	}

	if !dara.IsNil(request.TagShrink) {
		body["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建命名空间
//
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(tmpReq *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Ha) {
		body["Ha"] = request.Ha
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建命名空间
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 释放按量付费的实例
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 释放按量付费的实例
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除namespace
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除namespace
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// instance列表
//
// @param tmpReq - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(tmpReq *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// instance列表
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
// namespace列表
//
// @param tmpReq - DescribeNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespacesResponse
func (client *Client) DescribeNamespacesWithOptions(tmpReq *DescribeNamespacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeNamespacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNamespaces"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// namespace列表
//
// @param request - DescribeNamespacesRequest
//
// @return DescribeNamespacesResponse
func (client *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (_result *DescribeNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.DescribeNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持的region列表
//
// @param request - DescribeSupportedRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportedRegionsResponse
func (client *Client) DescribeSupportedRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeSupportedRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportedRegions"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportedRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持的region列表
//
// @return DescribeSupportedRegionsResponse
func (client *Client) DescribeSupportedRegions() (_result *DescribeSupportedRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportedRegionsResponse{}
	_body, _err := client.DescribeSupportedRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持的zoneId列表
//
// @param request - DescribeSupportedZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportedZonesResponse
func (client *Client) DescribeSupportedZonesWithOptions(request *DescribeSupportedZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportedZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportedZones"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportedZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持的zoneId列表
//
// @param request - DescribeSupportedZonesRequest
//
// @return DescribeSupportedZonesResponse
func (client *Client) DescribeSupportedZones(request *DescribeSupportedZonesRequest) (_result *DescribeSupportedZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportedZonesResponse{}
	_body, _err := client.DescribeSupportedZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举flinkasi标签
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
		Version:     dara.String("2021-10-28"),
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
// 列举flinkasi标签
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
// 对按量弹性实例修改resource quota
//
// @param tmpReq - ModifyElasticResourceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticResourceSpecResponse
func (client *Client) ModifyElasticResourceSpecWithOptions(tmpReq *ModifyElasticResourceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticResourceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyElasticResourceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticResourceSpec"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticResourceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对按量弹性实例修改resource quota
//
// @param request - ModifyElasticResourceSpecRequest
//
// @return ModifyElasticResourceSpecResponse
func (client *Client) ModifyElasticResourceSpec(request *ModifyElasticResourceSpecRequest) (_result *ModifyElasticResourceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticResourceSpecResponse{}
	_body, _err := client.ModifyElasticResourceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyInstanceVswitch is deprecated
//
// Summary:
//
// 修改集群交换机
//
// @param tmpReq - ModifyInstanceVswitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVswitchResponse
func (client *Client) ModifyInstanceVswitchWithOptions(tmpReq *ModifyInstanceVswitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVswitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceVswitchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaVSwitchIds) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, dara.String("HaVSwitchIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HaVSwitchIdsShrink) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceVswitch"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceVswitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyInstanceVswitch is deprecated
//
// Summary:
//
// 修改集群交换机
//
// @param request - ModifyInstanceVswitchRequest
//
// @return ModifyInstanceVswitchResponse
// Deprecated
func (client *Client) ModifyInstanceVswitch(request *ModifyInstanceVswitchRequest) (_result *ModifyInstanceVswitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceVswitchResponse{}
	_body, _err := client.ModifyInstanceVswitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改namespace资源，包含按量和包年包月、混合计费
//
// @param tmpReq - ModifyNamespaceSpecV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNamespaceSpecV2Response
func (client *Client) ModifyNamespaceSpecV2WithOptions(tmpReq *ModifyNamespaceSpecV2Request, runtime *dara.RuntimeOptions) (_result *ModifyNamespaceSpecV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyNamespaceSpecV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ElasticResourceSpec) {
		request.ElasticResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ElasticResourceSpec, dara.String("ElasticResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GuaranteedResourceSpec) {
		request.GuaranteedResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GuaranteedResourceSpec, dara.String("GuaranteedResourceSpec"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Ha) {
		query["Ha"] = request.Ha
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticResourceSpecShrink) {
		body["ElasticResourceSpec"] = request.ElasticResourceSpecShrink
	}

	if !dara.IsNil(request.GuaranteedResourceSpecShrink) {
		body["GuaranteedResourceSpec"] = request.GuaranteedResourceSpecShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNamespaceSpecV2"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNamespaceSpecV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改namespace资源，包含按量和包年包月、混合计费
//
// @param request - ModifyNamespaceSpecV2Request
//
// @return ModifyNamespaceSpecV2Response
func (client *Client) ModifyNamespaceSpecV2(request *ModifyNamespaceSpecV2Request) (_result *ModifyNamespaceSpecV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNamespaceSpecV2Response{}
	_body, _err := client.ModifyNamespaceSpecV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2021-10-28::ModifyInstanceSpec instead.
//
// Summary:
//
// 扩容/缩容
//
// @param tmpReq - ModifyPrepayInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayInstanceSpecResponse
func (client *Client) ModifyPrepayInstanceSpecWithOptions(tmpReq *ModifyPrepayInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPrepayInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaResourceSpec) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, dara.String("HaResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HaVSwitchIds) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, dara.String("HaVSwitchIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Ha) {
		body["Ha"] = request.Ha
	}

	if !dara.IsNil(request.HaResourceSpecShrink) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !dara.IsNil(request.HaVSwitchIdsShrink) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !dara.IsNil(request.HaZoneId) {
		body["HaZoneId"] = request.HaZoneId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPrepayInstanceSpec"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2021-10-28::ModifyInstanceSpec instead.
//
// Summary:
//
// 扩容/缩容
//
// @param request - ModifyPrepayInstanceSpecRequest
//
// @return ModifyPrepayInstanceSpecResponse
// Deprecated
func (client *Client) ModifyPrepayInstanceSpec(request *ModifyPrepayInstanceSpecRequest) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.ModifyPrepayInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2021-10-28::ModifyNamespaceSpec instead.
//
// Summary:
//
// 修改namespace资源分配
//
// @param tmpReq - ModifyPrepayNamespaceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayNamespaceSpecResponse
func (client *Client) ModifyPrepayNamespaceSpecWithOptions(tmpReq *ModifyPrepayNamespaceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPrepayNamespaceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPrepayNamespaceSpec"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPrepayNamespaceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2021-10-28::ModifyNamespaceSpec instead.
//
// Summary:
//
// 修改namespace资源分配
//
// @param request - ModifyPrepayNamespaceSpecRequest
//
// @return ModifyPrepayNamespaceSpecResponse
// Deprecated
func (client *Client) ModifyPrepayNamespaceSpec(request *ModifyPrepayNamespaceSpecRequest) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPrepayNamespaceSpecResponse{}
	_body, _err := client.ModifyPrepayNamespaceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按量付费转包年包月询价
//
// @param tmpReq - QueryConvertInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConvertInstancePriceResponse
func (client *Client) QueryConvertInstancePriceWithOptions(tmpReq *QueryConvertInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryConvertInstancePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryConvertInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NamespaceResourceSpecs) {
		request.NamespaceResourceSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NamespaceResourceSpecs, dara.String("NamespaceResourceSpecs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsAutoRenew) {
		body["IsAutoRenew"] = request.IsAutoRenew
	}

	if !dara.IsNil(request.NamespaceResourceSpecsShrink) {
		body["NamespaceResourceSpecs"] = request.NamespaceResourceSpecsShrink
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConvertInstancePrice"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConvertInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按量付费转包年包月询价
//
// @param request - QueryConvertInstancePriceRequest
//
// @return QueryConvertInstancePriceResponse
func (client *Client) QueryConvertInstancePrice(request *QueryConvertInstancePriceRequest) (_result *QueryConvertInstancePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryConvertInstancePriceResponse{}
	_body, _err := client.QueryConvertInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 包年包月转按量付费询价
//
// @param request - QueryConvertPrepayInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConvertPrepayInstancePriceResponse
func (client *Client) QueryConvertPrepayInstancePriceWithOptions(request *QueryConvertPrepayInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryConvertPrepayInstancePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConvertPrepayInstancePrice"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConvertPrepayInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 包年包月转按量付费询价
//
// @param request - QueryConvertPrepayInstancePriceRequest
//
// @return QueryConvertPrepayInstancePriceResponse
func (client *Client) QueryConvertPrepayInstancePrice(request *QueryConvertPrepayInstancePriceRequest) (_result *QueryConvertPrepayInstancePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryConvertPrepayInstancePriceResponse{}
	_body, _err := client.QueryConvertPrepayInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取创建实例的价格
//
// @param tmpReq - QueryCreateInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCreateInstancePriceResponse
func (client *Client) QueryCreateInstancePriceWithOptions(tmpReq *QueryCreateInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryCreateInstancePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryCreateInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaResourceSpec) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, dara.String("HaResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Storage) {
		request.StorageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Storage, dara.String("Storage"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArchitectureType) {
		body["ArchitectureType"] = request.ArchitectureType
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		body["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Extra) {
		body["Extra"] = request.Extra
	}

	if !dara.IsNil(request.Ha) {
		body["Ha"] = request.Ha
	}

	if !dara.IsNil(request.HaResourceSpecShrink) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	if !dara.IsNil(request.StorageShrink) {
		body["Storage"] = request.StorageShrink
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCreateInstancePrice"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCreateInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取创建实例的价格
//
// @param request - QueryCreateInstancePriceRequest
//
// @return QueryCreateInstancePriceResponse
func (client *Client) QueryCreateInstancePrice(request *QueryCreateInstancePriceRequest) (_result *QueryCreateInstancePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCreateInstancePriceResponse{}
	_body, _err := client.QueryCreateInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询付费类型为包年包月的实例修改资源规格的价格
//
// @param tmpReq - QueryModifyInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyInstancePriceResponse
func (client *Client) QueryModifyInstancePriceWithOptions(tmpReq *QueryModifyInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryModifyInstancePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryModifyInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaResourceSpec) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, dara.String("HaResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HaVSwitchIds) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, dara.String("HaVSwitchIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSpec) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, dara.String("ResourceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Ha) {
		body["Ha"] = request.Ha
	}

	if !dara.IsNil(request.HaResourceSpecShrink) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !dara.IsNil(request.HaVSwitchIdsShrink) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceSpecShrink) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryModifyInstancePrice"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询付费类型为包年包月的实例修改资源规格的价格
//
// @param request - QueryModifyInstancePriceRequest
//
// @return QueryModifyInstancePriceResponse
func (client *Client) QueryModifyInstancePrice(request *QueryModifyInstancePriceRequest) (_result *QueryModifyInstancePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryModifyInstancePriceResponse{}
	_body, _err := client.QueryModifyInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询付费类型为包年包月的实例续费价格
//
// @param request - QueryRenewInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRenewInstancePriceResponse
func (client *Client) QueryRenewInstancePriceWithOptions(request *QueryRenewInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryRenewInstancePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRenewInstancePrice"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRenewInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询付费类型为包年包月的实例续费价格
//
// @param request - QueryRenewInstancePriceRequest
//
// @return QueryRenewInstancePriceResponse
func (client *Client) QueryRenewInstancePrice(request *QueryRenewInstancePriceRequest) (_result *QueryRenewInstancePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRenewInstancePriceResponse{}
	_body, _err := client.QueryRenewInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 续费
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionCode) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.UsePromotionCode) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 续费
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
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
		Version:     dara.String("2021-10-28"),
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
// 打标签接口
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
// flinkasi去标签
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
		Version:     dara.String("2021-10-28"),
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
// flinkasi去标签
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
