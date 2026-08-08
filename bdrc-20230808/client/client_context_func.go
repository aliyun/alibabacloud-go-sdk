// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Checks data protection scoring rules. Calling this operation triggers an asynchronous task to check whether your resources meet the data protection scoring requirements.
//
// @param request - CheckRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRulesResponse
func (client *Client) CheckRulesWithContext(ctx context.Context, request *CheckRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceArn) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckRules"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/rules/check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shuts down the Backup and Disaster Recovery Center.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseBdrcServiceResponse
func (client *Client) CloseBdrcServiceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseBdrcServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseBdrcService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/service/close"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseBdrcServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protection policy.
//
// @param tmpReq - CreateProtectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtectionPolicyResponse
func (client *Client) CreateProtectionPolicyWithContext(ctx context.Context, tmpReq *CreateProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProtectionPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProtectionPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BoundResourceCategoryIds) {
		request.BoundResourceCategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BoundResourceCategoryIds, dara.String("BoundResourceCategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubProtectionPolicies) {
		request.SubProtectionPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubProtectionPolicies, dara.String("SubProtectionPolicies"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BoundResourceCategoryIdsShrink) {
		body["BoundResourceCategoryIds"] = request.BoundResourceCategoryIdsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProtectionPolicyName) {
		body["ProtectionPolicyName"] = request.ProtectionPolicyName
	}

	if !dara.IsNil(request.ProtectionPolicyRegionId) {
		body["ProtectionPolicyRegionId"] = request.ProtectionPolicyRegionId
	}

	if !dara.IsNil(request.SubProtectionPoliciesShrink) {
		body["SubProtectionPolicies"] = request.SubProtectionPoliciesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProtectionPolicy"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtectionPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource category.
//
// @param request - CreateResourceCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceCategoryResponse
func (client *Client) CreateResourceCategoryWithContext(ctx context.Context, request *CreateResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCategoryName) {
		body["ResourceCategoryName"] = request.ResourceCategoryName
	}

	if !dara.IsNil(request.ResourceMatcher) {
		body["ResourceMatcher"] = request.ResourceMatcher
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceCategory"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resource-categories/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protection policy.
//
// @param request - DeleteProtectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtectionPolicyResponse
func (client *Client) DeleteProtectionPolicyWithContext(ctx context.Context, ProtectionPolicyId *string, request *DeleteProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProtectionPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtectionPolicy"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies/" + dara.PercentEncode(dara.StringValue(ProtectionPolicyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtectionPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource category.
//
// @param request - DeleteResourceCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceCategoryResponse
func (client *Client) DeleteResourceCategoryWithContext(ctx context.Context, request *DeleteResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCategoryId) {
		body["ResourceCategoryId"] = request.ResourceCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceCategory"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resource-categories/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the check results for data protection rules.
//
// @param request - DescribeCheckDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCheckDetailsResponse
func (client *Client) DescribeCheckDetailsWithContext(ctx context.Context, request *DescribeCheckDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeCheckDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCheckDetails"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/check-details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCheckDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on data redundancy types for a cloud service.
//
// @param tmpReq - DescribeProductDataRedundancyTypeStatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductDataRedundancyTypeStatResponse
func (client *Client) DescribeProductDataRedundancyTypeStatWithContext(ctx context.Context, tmpReq *DescribeProductDataRedundancyTypeStatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProductDataRedundancyTypeStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeProductDataRedundancyTypeStatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceOwnerIds) {
		request.ResourceOwnerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceOwnerIds, dara.String("ResourceOwnerIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceOwnerIdsShrink) {
		query["ResourceOwnerIds"] = request.ResourceOwnerIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductDataRedundancyTypeStat"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/products/data-redundancy-type-stat"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductDataRedundancyTypeStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data protection score status of cloud services.
//
// @param tmpReq - DescribeProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsResponse
func (client *Client) DescribeProductsWithContext(ctx context.Context, tmpReq *DescribeProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeProductsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceOwnerIds) {
		request.ResourceOwnerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceOwnerIds, dara.String("ResourceOwnerIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceOwnerIdsShrink) {
		query["ResourceOwnerIds"] = request.ResourceOwnerIdsShrink
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProducts"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions.
//
// Description:
//
// BDRC本身是中心化的产品，接口用于部分与Region相关的功能使用。
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query resources
//
// @param tmpReq - DescribeResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcesResponse
func (client *Client) DescribeResourcesWithContext(ctx context.Context, tmpReq *DescribeResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceOwnerIds) {
		request.ResourceOwnerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceOwnerIds, dara.String("ResourceOwnerIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataRedundancyType) {
		query["DataRedundancyType"] = request.DataRedundancyType
	}

	if !dara.IsNil(request.FailedRuleTemplate) {
		query["FailedRuleTemplate"] = request.FailedRuleTemplate
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerIdsShrink) {
		query["ResourceOwnerIds"] = request.ResourceOwnerIdsShrink
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StorageClass) {
		query["StorageClass"] = request.StorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResources"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data protection rules.
//
// @param tmpReq - DescribeRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithContext(ctx context.Context, tmpReq *DescribeRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceOwnerIds) {
		request.ResourceOwnerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceOwnerIds, dara.String("ResourceOwnerIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceOwnerIdsShrink) {
		query["ResourceOwnerIds"] = request.ResourceOwnerIdsShrink
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRules"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an asynchronous task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithContext(ctx context.Context, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTask"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves asynchronous tasks.
//
// @param request - DescribeTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the most threatened resources.
//
// @param tmpReq - DescribeTopRiskyResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTopRiskyResourcesResponse
func (client *Client) DescribeTopRiskyResourcesWithContext(ctx context.Context, tmpReq *DescribeTopRiskyResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTopRiskyResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeTopRiskyResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceOwnerIds) {
		request.ResourceOwnerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceOwnerIds, dara.String("ResourceOwnerIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceOwnerIdsShrink) {
		query["ResourceOwnerIds"] = request.ResourceOwnerIdsShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTopRiskyResources"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/top-risky"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTopRiskyResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the data protection score for a cloud service.
//
// @param request - DisableCheckProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCheckProductResponse
func (client *Client) DisableCheckProductWithContext(ctx context.Context, request *DisableCheckProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableCheckProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCheckProduct"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/products/disable-check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCheckProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the data protection score for a resource.
//
// @param request - DisableCheckResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCheckResourceResponse
func (client *Client) DisableCheckResourceWithContext(ctx context.Context, request *DisableCheckResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableCheckResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceArn) {
		body["ResourceArn"] = request.ResourceArn
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCheckResource"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/disable-check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCheckResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables data protection scoring for an Alibaba Cloud service.
//
// @param request - EnableCheckProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCheckProductResponse
func (client *Client) EnableCheckProductWithContext(ctx context.Context, request *EnableCheckProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableCheckProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCheckProduct"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/products/enable-check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCheckProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables data protection scoring for a single resource.
//
// @param request - EnableCheckResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCheckResourceResponse
func (client *Client) EnableCheckResourceWithContext(ctx context.Context, request *EnableCheckResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableCheckResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceArn) {
		body["ResourceArn"] = request.ResourceArn
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCheckResource"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/enable-check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCheckResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启跨账号管理
//
// @param request - EnableCrossAccountManagementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCrossAccountManagementResponse
func (client *Client) EnableCrossAccountManagementWithContext(ctx context.Context, request *EnableCrossAccountManagementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableCrossAccountManagementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCrossAccountManagement"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/cross-accounts/enable-management"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCrossAccountManagementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Backup and Disaster Recovery Center.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBdrcServiceResponse
func (client *Client) GetBdrcServiceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBdrcServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBdrcService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/service"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBdrcServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a single message.
//
// @param request - GetMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageResponse
func (client *Client) GetMessageWithContext(ctx context.Context, MessageId *string, request *GetMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessage"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/messages/" + dara.PercentEncode(dara.StringValue(MessageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a protection policy.
//
// @param request - GetProtectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProtectionPolicyResponse
func (client *Client) GetProtectionPolicyWithContext(ctx context.Context, ProtectionPolicyId *string, request *GetProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProtectionPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProtectionPolicy"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies/" + dara.PercentEncode(dara.StringValue(ProtectionPolicyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProtectionPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a single resource category.
//
// @param request - GetResourceCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCategoryResponse
func (client *Client) GetResourceCategoryWithContext(ctx context.Context, request *GetResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceCategory"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resource-categories/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建跨账号纳管关系
//
// @param request - ListCrossAccountsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrossAccountsResponse
func (client *Client) ListCrossAccountsWithContext(ctx context.Context, request *ListCrossAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCrossAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccountOwnerId) {
		query["CrossAccountOwnerId"] = request.CrossAccountOwnerId
	}

	if !dara.IsNil(request.ManagementMode) {
		query["ManagementMode"] = request.ManagementMode
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrossAccounts"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/cross-accounts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrossAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries messages in batches.
//
// @param request - ListMessagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessagesResponse
func (client *Client) ListMessagesWithContext(ctx context.Context, request *ListMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MessageLevel) {
		query["MessageLevel"] = request.MessageLevel
	}

	if !dara.IsNil(request.MessageTimeEarlierThan) {
		query["MessageTimeEarlierThan"] = request.MessageTimeEarlierThan
	}

	if !dara.IsNil(request.MessageTimeLaterThan) {
		query["MessageTimeLaterThan"] = request.MessageTimeLaterThan
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessages"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of protection policies.
//
// @param request - ListProtectionPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProtectionPoliciesResponse
func (client *Client) ListProtectionPoliciesWithContext(ctx context.Context, request *ListProtectionPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProtectionPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProtectionPolicyId) {
		query["ProtectionPolicyId"] = request.ProtectionPolicyId
	}

	if !dara.IsNil(request.ProtectionPolicyRegionId) {
		query["ProtectionPolicyRegionId"] = request.ProtectionPolicyRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProtectionPolicies"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProtectionPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the application history of the protection policy.
//
// @param request - ListProtectionPolicyApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProtectionPolicyApplicationsResponse
func (client *Client) ListProtectionPolicyApplicationsWithContext(ctx context.Context, ProtectionPolicyId *string, request *ListProtectionPolicyApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProtectionPolicyApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyStatus) {
		query["ApplyStatus"] = request.ApplyStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SubProtectionPolicyType) {
		query["SubProtectionPolicyType"] = request.SubProtectionPolicyType
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProtectionPolicyApplications"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies/" + dara.PercentEncode(dara.StringValue(ProtectionPolicyId)) + "/list-applications"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProtectionPolicyApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of resource categories.
//
// @param request - ListResourceCategoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceCategoriesResponse
func (client *Client) ListResourceCategoriesWithContext(ctx context.Context, request *ListResourceCategoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceCategoryId) {
		query["ResourceCategoryId"] = request.ResourceCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceCategories"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resource-categories/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Backup and Disaster Recovery Center.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenBdrcServiceResponse
func (client *Client) OpenBdrcServiceWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenBdrcServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenBdrcService"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/service/open"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenBdrcServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预检查资源数量
//
// @param tmpReq - PrecheckResourceCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckResourceCountResponse
func (client *Client) PrecheckResourceCountWithContext(ctx context.Context, tmpReq *PrecheckResourceCountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PrecheckResourceCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PrecheckResourceCountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TagResourceMatchers) {
		request.TagResourceMatchersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagResourceMatchers, dara.String("TagResourceMatchers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagResourceMatchersShrink) {
		body["TagResourceMatchers"] = request.TagResourceMatchersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrecheckResourceCount"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/precheck-count"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrecheckResourceCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新跨账号纳管关系
//
// @param tmpReq - UpdateCrossAccountsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrossAccountsResponse
func (client *Client) UpdateCrossAccountsWithContext(ctx context.Context, tmpReq *UpdateCrossAccountsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCrossAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCrossAccountsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateTargets) {
		request.CreateTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTargets, dara.String("CreateTargets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeleteTargets) {
		request.DeleteTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteTargets, dara.String("DeleteTargets"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateTargetsShrink) {
		body["CreateTargets"] = request.CreateTargetsShrink
	}

	if !dara.IsNil(request.DeleteTargetsShrink) {
		body["DeleteTargets"] = request.DeleteTargetsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCrossAccounts"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/cross-accounts"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCrossAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a protection policy.
//
// @param tmpReq - UpdateProtectionPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProtectionPolicyResponse
func (client *Client) UpdateProtectionPolicyWithContext(ctx context.Context, ProtectionPolicyId *string, tmpReq *UpdateProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProtectionPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProtectionPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BoundResourceCategoryIds) {
		request.BoundResourceCategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BoundResourceCategoryIds, dara.String("BoundResourceCategoryIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubProtectionPolicies) {
		request.SubProtectionPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubProtectionPolicies, dara.String("SubProtectionPolicies"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BoundResourceCategoryIdsShrink) {
		body["BoundResourceCategoryIds"] = request.BoundResourceCategoryIdsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProtectionPolicyName) {
		body["ProtectionPolicyName"] = request.ProtectionPolicyName
	}

	if !dara.IsNil(request.SubProtectionPoliciesShrink) {
		body["SubProtectionPolicies"] = request.SubProtectionPoliciesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProtectionPolicy"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/protection-policies/" + dara.PercentEncode(dara.StringValue(ProtectionPolicyId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProtectionPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a resource category.
//
// @param request - UpdateResourceCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceCategoryResponse
func (client *Client) UpdateResourceCategoryWithContext(ctx context.Context, request *UpdateResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCategoryId) {
		body["ResourceCategoryId"] = request.ResourceCategoryId
	}

	if !dara.IsNil(request.ResourceCategoryName) {
		body["ResourceCategoryName"] = request.ResourceCategoryName
	}

	if !dara.IsNil(request.ResourceMatcher) {
		body["ResourceMatcher"] = request.ResourceMatcher
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceCategory"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resource-categories/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource list. When you call this operation, an asynchronous task is triggered to update your resource list and data protection score.
//
// @param request - UpdateResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourcesResponse
func (client *Client) UpdateResourcesWithContext(ctx context.Context, request *UpdateResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResources"),
		Version:     dara.String("2023-08-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
