// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 开通弹性计算
//
// @param tmpReq - ConvertHybridInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertHybridInstanceResponse
func (client *Client) ConvertHybridInstanceWithContext(ctx context.Context, tmpReq *ConvertHybridInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertHybridInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ConvertInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstanceWithContext(ctx context.Context, tmpReq *ConvertInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPrepayInstanceResponse
func (client *Client) ConvertPrepayInstanceWithContext(ctx context.Context, request *ConvertPrepayInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertPrepayInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, tmpReq *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithContext(ctx context.Context, tmpReq *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithContext(ctx context.Context, tmpReq *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DescribeNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespacesResponse
func (client *Client) DescribeNamespacesWithContext(ctx context.Context, tmpReq *DescribeNamespacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportedZonesResponse
func (client *Client) DescribeSupportedZonesWithContext(ctx context.Context, request *DescribeSupportedZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportedZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyElasticResourceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticResourceSpecResponse
func (client *Client) ModifyElasticResourceSpecWithContext(ctx context.Context, tmpReq *ModifyElasticResourceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticResourceSpecResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyInstanceVswitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVswitchResponse
func (client *Client) ModifyInstanceVswitchWithContext(ctx context.Context, tmpReq *ModifyInstanceVswitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVswitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyNamespaceSpecV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNamespaceSpecV2Response
func (client *Client) ModifyNamespaceSpecV2WithContext(ctx context.Context, tmpReq *ModifyNamespaceSpecV2Request, runtime *dara.RuntimeOptions) (_result *ModifyNamespaceSpecV2Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyPrepayInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayInstanceSpecResponse
func (client *Client) ModifyPrepayInstanceSpecWithContext(ctx context.Context, tmpReq *ModifyPrepayInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyPrepayNamespaceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayNamespaceSpecResponse
func (client *Client) ModifyPrepayNamespaceSpecWithContext(ctx context.Context, tmpReq *ModifyPrepayNamespaceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - QueryConvertInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConvertInstancePriceResponse
func (client *Client) QueryConvertInstancePriceWithContext(ctx context.Context, tmpReq *QueryConvertInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryConvertInstancePriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConvertPrepayInstancePriceResponse
func (client *Client) QueryConvertPrepayInstancePriceWithContext(ctx context.Context, request *QueryConvertPrepayInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryConvertPrepayInstancePriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - QueryCreateInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCreateInstancePriceResponse
func (client *Client) QueryCreateInstancePriceWithContext(ctx context.Context, tmpReq *QueryCreateInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryCreateInstancePriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - QueryModifyInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyInstancePriceResponse
func (client *Client) QueryModifyInstancePriceWithContext(ctx context.Context, tmpReq *QueryModifyInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryModifyInstancePriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRenewInstancePriceResponse
func (client *Client) QueryRenewInstancePriceWithContext(ctx context.Context, request *QueryRenewInstancePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryRenewInstancePriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
