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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-shanghai-finance-1": dara.String("bdrc.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("bdrc.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1":        dara.String("bdrc.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bdrc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Checks data protection scoring rules. Calling this operation triggers an asynchronous task to check whether your resources meet the data protection scoring requirements.
//
// @param request - CheckRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRulesResponse
func (client *Client) CheckRulesWithOptions(request *CheckRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks data protection scoring rules. Calling this operation triggers an asynchronous task to check whether your resources meet the data protection scoring requirements.
//
// @param request - CheckRulesRequest
//
// @return CheckRulesResponse
func (client *Client) CheckRules(request *CheckRulesRequest) (_result *CheckRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckRulesResponse{}
	_body, _err := client.CheckRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CloseBdrcServiceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseBdrcServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CloseBdrcServiceResponse
func (client *Client) CloseBdrcService() (_result *CloseBdrcServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseBdrcServiceResponse{}
	_body, _err := client.CloseBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateProtectionPolicyWithOptions(tmpReq *CreateProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProtectionPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateProtectionPolicyRequest
//
// @return CreateProtectionPolicyResponse
func (client *Client) CreateProtectionPolicy(request *CreateProtectionPolicyRequest) (_result *CreateProtectionPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProtectionPolicyResponse{}
	_body, _err := client.CreateProtectionPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateResourceCategoryWithOptions(request *CreateResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceCategoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateResourceCategoryResponse
func (client *Client) CreateResourceCategory(request *CreateResourceCategoryRequest) (_result *CreateResourceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceCategoryResponse{}
	_body, _err := client.CreateResourceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteProtectionPolicyWithOptions(ProtectionPolicyId *string, request *DeleteProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProtectionPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteProtectionPolicyResponse
func (client *Client) DeleteProtectionPolicy(ProtectionPolicyId *string, request *DeleteProtectionPolicyRequest) (_result *DeleteProtectionPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProtectionPolicyResponse{}
	_body, _err := client.DeleteProtectionPolicyWithOptions(ProtectionPolicyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteResourceCategoryWithOptions(request *DeleteResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceCategoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteResourceCategoryResponse
func (client *Client) DeleteResourceCategory(request *DeleteResourceCategoryRequest) (_result *DeleteResourceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceCategoryResponse{}
	_body, _err := client.DeleteResourceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCheckDetailsWithOptions(request *DescribeCheckDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeCheckDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCheckDetailsResponse
func (client *Client) DescribeCheckDetails(request *DescribeCheckDetailsRequest) (_result *DescribeCheckDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCheckDetailsResponse{}
	_body, _err := client.DescribeCheckDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeProductDataRedundancyTypeStatWithOptions(tmpReq *DescribeProductDataRedundancyTypeStatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProductDataRedundancyTypeStatResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeProductDataRedundancyTypeStatRequest
//
// @return DescribeProductDataRedundancyTypeStatResponse
func (client *Client) DescribeProductDataRedundancyTypeStat(request *DescribeProductDataRedundancyTypeStatRequest) (_result *DescribeProductDataRedundancyTypeStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProductDataRedundancyTypeStatResponse{}
	_body, _err := client.DescribeProductDataRedundancyTypeStatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeProductsWithOptions(tmpReq *DescribeProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeProductsRequest
//
// @return DescribeProductsResponse
func (client *Client) DescribeProducts(request *DescribeProductsRequest) (_result *DescribeProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProductsResponse{}
	_body, _err := client.DescribeProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeResourcesWithOptions(tmpReq *DescribeResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeResourcesRequest
//
// @return DescribeResourcesResponse
func (client *Client) DescribeResources(request *DescribeResourcesRequest) (_result *DescribeResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourcesResponse{}
	_body, _err := client.DescribeResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRulesWithOptions(tmpReq *DescribeRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeRulesRequest
//
// @return DescribeRulesResponse
func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTaskWithOptions(TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTaskResponse
func (client *Client) DescribeTask(TaskId *string) (_result *DescribeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTopRiskyResourcesWithOptions(tmpReq *DescribeTopRiskyResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTopRiskyResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeTopRiskyResourcesRequest
//
// @return DescribeTopRiskyResourcesResponse
func (client *Client) DescribeTopRiskyResources(request *DescribeTopRiskyResourcesRequest) (_result *DescribeTopRiskyResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTopRiskyResourcesResponse{}
	_body, _err := client.DescribeTopRiskyResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisableCheckProductWithOptions(request *DisableCheckProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableCheckProductResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisableCheckProductResponse
func (client *Client) DisableCheckProduct(request *DisableCheckProductRequest) (_result *DisableCheckProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableCheckProductResponse{}
	_body, _err := client.DisableCheckProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisableCheckResourceWithOptions(request *DisableCheckResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableCheckResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisableCheckResourceResponse
func (client *Client) DisableCheckResource(request *DisableCheckResourceRequest) (_result *DisableCheckResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableCheckResourceResponse{}
	_body, _err := client.DisableCheckResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EnableCheckProductWithOptions(request *EnableCheckProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableCheckProductResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EnableCheckProductResponse
func (client *Client) EnableCheckProduct(request *EnableCheckProductRequest) (_result *EnableCheckProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableCheckProductResponse{}
	_body, _err := client.EnableCheckProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EnableCheckResourceWithOptions(request *EnableCheckResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableCheckResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EnableCheckResourceResponse
func (client *Client) EnableCheckResource(request *EnableCheckResourceRequest) (_result *EnableCheckResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableCheckResourceResponse{}
	_body, _err := client.EnableCheckResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetBdrcServiceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBdrcServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetBdrcServiceResponse
func (client *Client) GetBdrcService() (_result *GetBdrcServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBdrcServiceResponse{}
	_body, _err := client.GetBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMessageWithOptions(MessageId *string, request *GetMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMessageResponse
func (client *Client) GetMessage(MessageId *string, request *GetMessageRequest) (_result *GetMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMessageResponse{}
	_body, _err := client.GetMessageWithOptions(MessageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProtectionPolicyWithOptions(ProtectionPolicyId *string, request *GetProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProtectionPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProtectionPolicyResponse
func (client *Client) GetProtectionPolicy(ProtectionPolicyId *string, request *GetProtectionPolicyRequest) (_result *GetProtectionPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProtectionPolicyResponse{}
	_body, _err := client.GetProtectionPolicyWithOptions(ProtectionPolicyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceCategoryWithOptions(request *GetResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceCategoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetResourceCategoryResponse
func (client *Client) GetResourceCategory(request *GetResourceCategoryRequest) (_result *GetResourceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceCategoryResponse{}
	_body, _err := client.GetResourceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMessagesWithOptions(request *ListMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMessagesResponse
func (client *Client) ListMessages(request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProtectionPoliciesWithOptions(request *ListProtectionPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProtectionPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProtectionPoliciesResponse
func (client *Client) ListProtectionPolicies(request *ListProtectionPoliciesRequest) (_result *ListProtectionPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProtectionPoliciesResponse{}
	_body, _err := client.ListProtectionPoliciesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProtectionPolicyApplicationsWithOptions(ProtectionPolicyId *string, request *ListProtectionPolicyApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProtectionPolicyApplicationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProtectionPolicyApplicationsResponse
func (client *Client) ListProtectionPolicyApplications(ProtectionPolicyId *string, request *ListProtectionPolicyApplicationsRequest) (_result *ListProtectionPolicyApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProtectionPolicyApplicationsResponse{}
	_body, _err := client.ListProtectionPolicyApplicationsWithOptions(ProtectionPolicyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceCategoriesWithOptions(request *ListResourceCategoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceCategoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResourceCategoriesResponse
func (client *Client) ListResourceCategories(request *ListResourceCategoriesRequest) (_result *ListResourceCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceCategoriesResponse{}
	_body, _err := client.ListResourceCategoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OpenBdrcServiceWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenBdrcServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OpenBdrcServiceResponse
func (client *Client) OpenBdrcService() (_result *OpenBdrcServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenBdrcServiceResponse{}
	_body, _err := client.OpenBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateProtectionPolicyWithOptions(ProtectionPolicyId *string, tmpReq *UpdateProtectionPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProtectionPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateProtectionPolicyRequest
//
// @return UpdateProtectionPolicyResponse
func (client *Client) UpdateProtectionPolicy(ProtectionPolicyId *string, request *UpdateProtectionPolicyRequest) (_result *UpdateProtectionPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProtectionPolicyResponse{}
	_body, _err := client.UpdateProtectionPolicyWithOptions(ProtectionPolicyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateResourceCategoryWithOptions(request *UpdateResourceCategoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceCategoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateResourceCategoryResponse
func (client *Client) UpdateResourceCategory(request *UpdateResourceCategoryRequest) (_result *UpdateResourceCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceCategoryResponse{}
	_body, _err := client.UpdateResourceCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateResourcesWithOptions(request *UpdateResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateResourcesResponse
func (client *Client) UpdateResources(request *UpdateResourcesRequest) (_result *UpdateResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourcesResponse{}
	_body, _err := client.UpdateResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
