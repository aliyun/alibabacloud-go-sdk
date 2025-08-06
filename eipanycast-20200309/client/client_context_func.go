// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an Anycast elastic IP address (Anycast EIP).
//
// @param request - AllocateAnycastEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateAnycastEipAddressResponse
func (client *Client) AllocateAnycastEipAddressWithContext(ctx context.Context, request *AllocateAnycastEipAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocateAnycastEipAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceLocation) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateAnycastEipAddress"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateAnycastEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an Anycast elastic IP address (Anycast EIP) with an endpoint.
//
// @param request - AssociateAnycastEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateAnycastEipAddressResponse
func (client *Client) AssociateAnycastEipAddressWithContext(ctx context.Context, request *AssociateAnycastEipAddressRequest, runtime *dara.RuntimeOptions) (_result *AssociateAnycastEipAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.AssociationMode) {
		query["AssociationMode"] = request.AssociationMode
	}

	if !dara.IsNil(request.BindInstanceId) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !dara.IsNil(request.BindInstanceRegionId) {
		query["BindInstanceRegionId"] = request.BindInstanceRegionId
	}

	if !dara.IsNil(request.BindInstanceType) {
		query["BindInstanceType"] = request.BindInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.PopLocations) {
		query["PopLocations"] = request.PopLocations
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateAnycastEipAddress"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateAnycastEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改AnycastEIp实例资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2020-03-09"),
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
// Queries Anycast elastic IP addresses (Anycast EIPs) associated with specified instance IP addresses or instance IDs, including instance status, maximum bandwidth, and associated resources.
//
// @param request - DescribeAnycastEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnycastEipAddressResponse
func (client *Client) DescribeAnycastEipAddressWithContext(ctx context.Context, request *DescribeAnycastEipAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnycastEipAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.BindInstanceId) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnycastEipAddress"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnycastEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the access points in an area.
//
// @param request - DescribeAnycastPopLocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnycastPopLocationsResponse
func (client *Client) DescribeAnycastPopLocationsWithContext(ctx context.Context, request *DescribeAnycastPopLocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnycastPopLocationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceLocation) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnycastPopLocations"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnycastPopLocationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where you can associate Anycast elastic IP addresses (Anycast EIPs) with endpoints.
//
// @param request - DescribeAnycastServerRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnycastServerRegionsResponse
func (client *Client) DescribeAnycastServerRegionsWithContext(ctx context.Context, request *DescribeAnycastServerRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnycastServerRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceLocation) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnycastServerRegions"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnycastServerRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Anycast elastic IP addresses (Anycast EIPs) in an access area, including instance status, maximum bandwidth, and associated resources.
//
// @param request - ListAnycastEipAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnycastEipAddressesResponse
func (client *Client) ListAnycastEipAddressesWithContext(ctx context.Context, request *ListAnycastEipAddressesRequest, runtime *dara.RuntimeOptions) (_result *ListAnycastEipAddressesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastEipAddress) {
		query["AnycastEipAddress"] = request.AnycastEipAddress
	}

	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.AnycastIds) {
		query["AnycastIds"] = request.AnycastIds
	}

	if !dara.IsNil(request.BindInstanceIds) {
		query["BindInstanceIds"] = request.BindInstanceIds
	}

	if !dara.IsNil(request.BusinessStatus) {
		query["BusinessStatus"] = request.BusinessStatus
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceLocation) {
		query["ServiceLocation"] = request.ServiceLocation
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnycastEipAddresses"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnycastEipAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Version:     dara.String("2020-03-09"),
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
// Modifies the name and description of an Anycast elastic IP address (Anycast EIP).
//
// @param request - ModifyAnycastEipAddressAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAnycastEipAddressAttributeResponse
func (client *Client) ModifyAnycastEipAddressAttributeWithContext(ctx context.Context, request *ModifyAnycastEipAddressAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAnycastEipAddressAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAnycastEipAddressAttribute"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAnycastEipAddressAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maximum bandwidth of an Anycast elastic IP address (Anycast EIP).
//
// @param request - ModifyAnycastEipAddressSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAnycastEipAddressSpecResponse
func (client *Client) ModifyAnycastEipAddressSpecWithContext(ctx context.Context, request *ModifyAnycastEipAddressSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyAnycastEipAddressSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAnycastEipAddressSpec"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAnycastEipAddressSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an Anycast elastic IP address (Anycast EIP).
//
// @param request - ReleaseAnycastEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseAnycastEipAddressResponse
func (client *Client) ReleaseAnycastEipAddressWithContext(ctx context.Context, request *ReleaseAnycastEipAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleaseAnycastEipAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseAnycastEipAddress"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseAnycastEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Version:     dara.String("2020-03-09"),
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
// Disassociates an Anycast elastic IP address (Anycast EIP) from an endpoint.
//
// @param request - UnassociateAnycastEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateAnycastEipAddressResponse
func (client *Client) UnassociateAnycastEipAddressWithContext(ctx context.Context, request *UnassociateAnycastEipAddressRequest, runtime *dara.RuntimeOptions) (_result *UnassociateAnycastEipAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.BindInstanceId) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !dara.IsNil(request.BindInstanceRegionId) {
		query["BindInstanceRegionId"] = request.BindInstanceRegionId
	}

	if !dara.IsNil(request.BindInstanceType) {
		query["BindInstanceType"] = request.BindInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassociateAnycastEipAddress"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassociateAnycastEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from Anycast EIPs.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Version:     dara.String("2020-03-09"),
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

// Summary:
//
// If an Anycast EIP is associated with endpoints in different regions, you can change the access points that are mapped to an endpoint. You can call UpdateAnycastEipAddressAssociations to add or delete endpoints to modify the mappings between endpoints and access points.
//
// @param request - UpdateAnycastEipAddressAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAnycastEipAddressAssociationsResponse
func (client *Client) UpdateAnycastEipAddressAssociationsWithContext(ctx context.Context, request *UpdateAnycastEipAddressAssociationsRequest, runtime *dara.RuntimeOptions) (_result *UpdateAnycastEipAddressAssociationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnycastId) {
		query["AnycastId"] = request.AnycastId
	}

	if !dara.IsNil(request.AssociationMode) {
		query["AssociationMode"] = request.AssociationMode
	}

	if !dara.IsNil(request.BindInstanceId) {
		query["BindInstanceId"] = request.BindInstanceId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.PopLocationAddList) {
		query["PopLocationAddList"] = request.PopLocationAddList
	}

	if !dara.IsNil(request.PopLocationDeleteList) {
		query["PopLocationDeleteList"] = request.PopLocationDeleteList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAnycastEipAddressAssociations"),
		Version:     dara.String("2020-03-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAnycastEipAddressAssociationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
