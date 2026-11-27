// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Modifies the resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Flink AI service.
//
// Description:
//
// Disables the Flink AI service.
//
// > **AI/Agent invocation guidelines**: If you call this OpenAPI operation in automated scenarios such as AI assistants or Code Agents, install and use the [official skill provided by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the operation directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI/Agent scenarios.
//
// @param request - CloseFlinkAiServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseFlinkAiServiceResponse
func (client *Client) CloseFlinkAiServiceWithContext(ctx context.Context, request *CloseFlinkAiServiceRequest, runtime *dara.RuntimeOptions) (_result *CloseFlinkAiServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseFlinkAiService"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseFlinkAiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts a subscription instance to a hybrid billing instance.
//
// Description:
//
// *Before using this operation, make sure you fully understand the billing method and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/hybrid-pricing) of hybrid billing for Fully Managed Flink.**
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
// Converts a pay-as-you-go workspace to a subscription workspace.
//
// Description:
//
// *Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/switch-from-pay-as-you-go-to-subscription) of fully managed Flink before you call this operation.**
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
// Converts a subscription instance to pay-as-you-go.
//
// Description:
//
// Before using this operation, make sure that you fully understand the [billing methods and pricing](https://www.alibabacloud.com/help/en/flink/product-overview/switch-from-subscription-to-pay-as-you-go) of Alibaba Cloud Realtime Compute for Apache Flink.
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
// Creates a subscription or pay-as-you-go fully managed Flink workspace.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/billing-overview) of fully managed Flink.**
//
// > **AI/Agent invocation guide**: If you call this OpenAPI operation in automated scenarios such as AI assistants or Code Agents, we recommend that you install and use the [official Skill encapsulated by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the operation directly, append the AlibabaCloud-Agent-Skills/custom identifier to the User-Agent in the request header. This helps the platform continuously improve support for AI/Agent scenarios.
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
	if !dara.IsNil(tmpReq.DefaultHaNamespaceResourceSpec) {
		request.DefaultHaNamespaceResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultHaNamespaceResourceSpec, dara.String("DefaultHaNamespaceResourceSpec"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DefaultNamespaceResourceSpec) {
		request.DefaultNamespaceResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultNamespaceResourceSpec, dara.String("DefaultNamespaceResourceSpec"), dara.String("json"))
	}

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

	if !dara.IsNil(request.DefaultHaNamespaceResourceSpecShrink) {
		body["DefaultHaNamespaceResourceSpec"] = request.DefaultHaNamespaceResourceSpecShrink
	}

	if !dara.IsNil(request.DefaultNamespaceResourceSpecShrink) {
		body["DefaultNamespaceResourceSpec"] = request.DefaultNamespaceResourceSpecShrink
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
// Creates a workspace in a fully managed Flink instance.
//
// Description:
//
// Creates a workspace in a fully managed Flink instance.
//
// > **AI/Agent invocation guide**: If you call this OpenAPI operation in automated scenarios such as AI assistants or Code Agents, install and use the [official skill provided by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the operation directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI/Agent scenarios.
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
// Releases a pay-as-you-go Realtime Compute for Apache Flink workspace.
//
// Description:
//
// *Before using this operation, make sure that you fully understand the billing method and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/refund-policy) of Realtime Compute for Apache Flink.**
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
// Deletes a project workspace.
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
// Queries the details of one or more fully managed Flink workspaces.
//
// Description:
//
// > You cannot release a workspace in the following scenarios: - In pay-as-you-go mode, the workspace is in the Creating, Releasing, or Creation Timed Out state. - If the billing method is subscription, unsubscribe from the resources on the Unsubscription page in the User Center.
//
// > **AI/Agent invocation guide**: If you call this API operation in automated scenarios such as AI assistants or Code Agents, we recommend that you install and use the [officially encapsulated skill](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API operation directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent header of the request. This helps the platform continuously improve support for AI/Agent scenarios.
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
// Queries the namespace information of a specific instance.
//
// Description:
//
// You can call DescribeNamespaces to query the namespace information of a specific instance.
//
// > **AI/Agent invocation guide**: If you call this OpenAPI in automated scenarios such as AI assistants or Code Agents, install and use the [official platform-encapsulated skill](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header to help the platform continuously improve support for AI and agent scenarios.
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
// Queries the zones available for purchase.
//
// Description:
//
// Queries the zones available for purchase.
//
// > **AI/Agent invocation guidelines**: If you call this OpenAPI in automated scenarios such as AI assistants or Code Agents, we recommend that you install and use the [official skill provided by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI and agent scenarios.
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
// Retrieves the status of the Flink AI service.
//
// Description:
//
// Retrieves the status of the Flink AI service.
//
// > **AI/Agent invocation guide**: If you call this OpenAPI in automated scenarios such as AI assistants or Code Agents, install and use the [official platform-encapsulated skill](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI and Agent scenarios.
//
// @param request - GetFlinkAiServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlinkAiServiceResponse
func (client *Client) GetFlinkAiServiceWithContext(ctx context.Context, request *GetFlinkAiServiceRequest, runtime *dara.RuntimeOptions) (_result *GetFlinkAiServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlinkAiService"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlinkAiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of the free quota for Flink AI services.
//
// Description:
//
// Queries the usage of the free quota for Flink AI services.
//
// > **AI/Agent invocation guide**: If you call this OpenAPI in automated scenarios such as AI assistants or Code Agents, install and use the [official platform-encapsulated skill](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI and Agent scenarios.
//
// @param request - GetFlinkAiServiceFreeQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlinkAiServiceFreeQuotaResponse
func (client *Client) GetFlinkAiServiceFreeQuotaWithContext(ctx context.Context, request *GetFlinkAiServiceFreeQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetFlinkAiServiceFreeQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlinkAiServiceFreeQuota"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlinkAiServiceFreeQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource tags. You can query keys by tag values, query values by keys, or retrieve all tags used in a fully managed Flink workspace.
//
// Description:
//
// You can call ListTagResources to query resource tags. You can query keys by tag values, query values by keys, or retrieve all tags used in a fully managed Flink workspace.
//
// > **AI/Agent invocation guide**: If you call this API in automated scenarios such as AI assistants or Code Agents, install and use the [official skill provided by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI and agent scenarios.
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
// Modifies the shutdown protection for Flink AI services.
//
// Description:
//
// Modifies the shutdown protection for Flink AI services.
//
// > **AI/Agent invocation guide**: If you call this API in automated scenarios such as AI assistants or Code Agents, install and use the [official platform-encapsulated skill](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI/Agent scenarios.
//
// @param request - ModifyAiServiceProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAiServiceProtectionResponse
func (client *Client) ModifyAiServiceProtectionWithContext(ctx context.Context, request *ModifyAiServiceProtectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAiServiceProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtection) {
		body["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAiServiceProtection"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAiServiceProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maximum pay-as-you-go resource limit for a hybrid billing instance.
//
// Description:
//
// Make sure that you are familiar with the [billing methods](https://www.alibabacloud.com/help/en/flink/user-guide/reconfigure-resources#task-2507532) and pricing of Realtime Compute for Apache Flink before you call this operation.
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

// Summary:
//
// Scales the resources of a workspace.
//
// Description:
//
// Make sure that you fully understand the [billing methods](https://www.alibabacloud.com/help/en/flink/user-guide/reconfigure-resources#task-2507532) and pricing of Realtime Compute for Apache Flink before you call this operation.
//
// @param tmpReq - ModifyInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpecWithContext(ctx context.Context, tmpReq *ModifyInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceSpecShrinkRequest{}
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
		Action:      dara.String("ModifyInstanceSpec"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSpecResponse{}
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
// Modifies the vSwitches available to a Flink workspace.
//
// Description:
//
// Before using this operation, make sure that you fully understand the restrictions on [modifying vSwitches](https://www.alibabacloud.com/help/en/flink/user-guide/modify-a-vswitch) in Realtime Compute for Apache Flink.
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
// Modifies the resource configurations of a project namespace in a pay-as-you-go, subscription, or hybrid billing workspace.
//
// Description:
//
// >When calling this operation, note the following: - The order status must be normal, that is, OrderType=NORMAL. - When decreasing the quota, the specified resource specification quantity cannot be less than the quantity already in use. For details about the project resource specifications before and after the change, call [DescribeNamespaces](https://help.aliyun.com/document_detail/323441.html).
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
// Modifies the compute resource specifications of a subscription fully managed Flink instance. If the adjusted specifications are smaller than the current specifications, a scale-down operation is performed. If the adjusted specifications are larger than the current specifications, a scale-up operation is performed.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/subscription) of fully managed Flink.**
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
// Modifies the specifications of a namespace in a subscription instance.
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
// Activates the Flink AI service.
//
// Description:
//
// Activates the Flink AI service.
//
// > **AI/Agent invocation guide**: If you call this OpenAPI in automated scenarios such as AI assistants or Code Agents, install and use the [official skill encapsulated by the platform](https://skills.aliyun.com/skills/alibabacloud-flink-instance-manage). If you integrate the OpenAPI client and call the API directly, include the AlibabaCloud-Agent-Skills/custom identifier in the User-Agent request header. This helps the platform continuously improve support for AI/Agent scenarios.
//
// @param request - OpenFlinkAiServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenFlinkAiServiceResponse
func (client *Client) OpenFlinkAiServiceWithContext(ctx context.Context, request *OpenFlinkAiServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenFlinkAiServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenFlinkAiService"),
		Version:     dara.String("2021-10-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenFlinkAiServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for changing the billing method of an instance from pay-as-you-go to subscription.
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
// Queries the estimated refund amount after your Flink instance is switched from subscription to pay-as-you-go. After the switch, your usage is metered hourly and billed accordingly.
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
// Retrieves the pricing for purchasing a workspace for the current account.
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
// Queries the price for upgrading or downgrading an instance.
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
// Queries the price for renewing a subscription workspace.
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
// Renews a subscription fully managed Flink workspace.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/help/en/flink/product-overview/renewal-policy) of fully managed Flink.**
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
// Call this API to add tags to resources.
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
// Deletes resource tags.
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
