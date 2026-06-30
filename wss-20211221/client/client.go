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
		"cn-shanghai":    dara.String("wss.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": dara.String("wss.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("wss"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Use this API to order, renew, and modify specific products, such as monthly resource plans for Elastic Desktop Service (EDS) Enterprise Edition.
//
// Description:
//
// <props="china">
//
// Before calling this API, make sure you understand how Wuying Workspace is billed and its [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws).
//
// <props="intl">
//
// Before calling this API, make sure you understand how Wuying Workspace is billed and its [pricing](https://www.alibabacloud.com/zh/product/cloud-desktop?#J_8623712560).
//
// If you do not specify automatic payment, this API does not handle the payment. You must use the returned order ID to construct a payment URL. The order becomes active and the resource is provisioned only after the payment is complete.
//
// @param tmpReq - CreateMultiOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiOrderResponse
func (client *Client) CreateMultiOrderWithOptions(tmpReq *CreateMultiOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMultiOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelCookie) {
		query["ChannelCookie"] = request.ChannelCookie
	}

	if !dara.IsNil(request.OrderItems) {
		query["OrderItems"] = request.OrderItems
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PropertiesShrink) {
		query["Properties"] = request.PropertiesShrink
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultiOrder"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultiOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use this API to order, renew, and modify specific products, such as monthly resource plans for Elastic Desktop Service (EDS) Enterprise Edition.
//
// Description:
//
// <props="china">
//
// Before calling this API, make sure you understand how Wuying Workspace is billed and its [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws).
//
// <props="intl">
//
// Before calling this API, make sure you understand how Wuying Workspace is billed and its [pricing](https://www.alibabacloud.com/zh/product/cloud-desktop?#J_8623712560).
//
// If you do not specify automatic payment, this API does not handle the payment. You must use the returned order ID to construct a payment URL. The order becomes active and the resource is provisioned only after the payment is complete.
//
// @param request - CreateMultiOrderRequest
//
// @return CreateMultiOrderResponse
func (client *Client) CreateMultiOrder(request *CreateMultiOrderRequest) (_result *CreateMultiOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultiOrderResponse{}
	_body, _err := client.CreateMultiOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of agents and their usage information.
//
// Description:
//
// ## Usage notes
//
// - **Pagination**: This operation supports pagination by using the `NextToken` and `MaxResults` parameters. For the first request, set `NextToken` to an empty string.
//
// - **Filtering**: Use the `AgentType` and `AgentIds` parameters to filter the results.
//
// - **Status filtering**: Use the `Status` parameter to filter agents by status. Valid values are 0 (deleted) and 1 (active).
//
// - **Sorting**: By default, the results are sorted by `id` in ascending order.
//
// - **Additional parameter for anonymous edition**: The `FillInstance` parameter automatically populates the ID of the JVS_COPILOT agent that is associated with the current user.
//
// @param request - DescribeCreditPackageAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCreditPackageAgentsResponse
func (client *Client) DescribeCreditPackageAgentsWithOptions(request *DescribeCreditPackageAgentsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreditPackageAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentIds) {
		query["AgentIds"] = request.AgentIds
	}

	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCreditPackageAgents"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCreditPackageAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of agents and their usage information.
//
// Description:
//
// ## Usage notes
//
// - **Pagination**: This operation supports pagination by using the `NextToken` and `MaxResults` parameters. For the first request, set `NextToken` to an empty string.
//
// - **Filtering**: Use the `AgentType` and `AgentIds` parameters to filter the results.
//
// - **Status filtering**: Use the `Status` parameter to filter agents by status. Valid values are 0 (deleted) and 1 (active).
//
// - **Sorting**: By default, the results are sorted by `id` in ascending order.
//
// - **Additional parameter for anonymous edition**: The `FillInstance` parameter automatically populates the ID of the JVS_COPILOT agent that is associated with the current user.
//
// @param request - DescribeCreditPackageAgentsRequest
//
// @return DescribeCreditPackageAgentsResponse
func (client *Client) DescribeCreditPackageAgents(request *DescribeCreditPackageAgentsRequest) (_result *DescribeCreditPackageAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCreditPackageAgentsResponse{}
	_body, _err := client.DescribeCreditPackageAgentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries credit usage by a specified dimension such as user, credit package, or agent.
//
// Description:
//
// ## Operation description
//
// This API operation queries credit usage details based on the dimension specified by `UsageType` (User, CreditPackage, or Agent). The response includes the total, remaining, and used credits of the current credit package, hourly consumption samples, alert thresholds, period quotas, and other information.
//
// - **User**: User dimension. Returns the aggregated usage and remaining credits across all active credit packages for the current user.
//
// - **CreditPackage**: Credit package dimension. Returns the total, remaining, and consumption samples for a specified credit package instance.
//
// - **Agent**: Agent dimension. Returns the cumulative usage, current period usage, quota, alert, and other information for a specified agent.
//
// **Notes**:
//
// - The `InstanceIds` parameter can be omitted when `UsageType=User`. Set this parameter to the credit package instance ID when `UsageType=CreditPackage`, or to the AgentId when `UsageType=Agent`.
//
// - Anonymous requests support the `FillInstance` parameter. If `InstanceIds` is not explicitly specified and `FillInstance=true`, the server automatically populates the bound `JVS_COPILOT` AgentId based on the current logon `wyId`.
//
// - Time window constants: The `dayUsedCredit` statistics window is `now - ONE_DAY_MILLIS`, and the `weekUsedCredit` statistics window is `now - ONE_WEEK_MILLIS`.
//
// - The consumption samples in `currentCreditConsumeList` are aggregated by hour and may have an asynchronous synchronization delay of up to 5 minutes.
//
// @param request - DescribeCreditUsageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCreditUsageInfoResponse
func (client *Client) DescribeCreditUsageInfoWithOptions(request *DescribeCreditUsageInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreditUsageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.UsageType) {
		query["UsageType"] = request.UsageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCreditUsageInfo"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCreditUsageInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries credit usage by a specified dimension such as user, credit package, or agent.
//
// Description:
//
// ## Operation description
//
// This API operation queries credit usage details based on the dimension specified by `UsageType` (User, CreditPackage, or Agent). The response includes the total, remaining, and used credits of the current credit package, hourly consumption samples, alert thresholds, period quotas, and other information.
//
// - **User**: User dimension. Returns the aggregated usage and remaining credits across all active credit packages for the current user.
//
// - **CreditPackage**: Credit package dimension. Returns the total, remaining, and consumption samples for a specified credit package instance.
//
// - **Agent**: Agent dimension. Returns the cumulative usage, current period usage, quota, alert, and other information for a specified agent.
//
// **Notes**:
//
// - The `InstanceIds` parameter can be omitted when `UsageType=User`. Set this parameter to the credit package instance ID when `UsageType=CreditPackage`, or to the AgentId when `UsageType=Agent`.
//
// - Anonymous requests support the `FillInstance` parameter. If `InstanceIds` is not explicitly specified and `FillInstance=true`, the server automatically populates the bound `JVS_COPILOT` AgentId based on the current logon `wyId`.
//
// - Time window constants: The `dayUsedCredit` statistics window is `now - ONE_DAY_MILLIS`, and the `weekUsedCredit` statistics window is `now - ONE_WEEK_MILLIS`.
//
// - The consumption samples in `currentCreditConsumeList` are aggregated by hour and may have an asynchronous synchronization delay of up to 5 minutes.
//
// @param request - DescribeCreditUsageInfoRequest
//
// @return DescribeCreditUsageInfoResponse
func (client *Client) DescribeCreditUsageInfo(request *DescribeCreditUsageInfoRequest) (_result *DescribeCreditUsageInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCreditUsageInfoResponse{}
	_body, _err := client.DescribeCreditUsageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves resource deduction and usage statistics based on specified criteria.
//
// Description:
//
// ## Request
//
// - This API supports GET and POST methods.
//
// - The `periods` parameter is a JSON array of `PeriodParam` objects, each containing the `periodUnit` and `baseTime` fields.
//
// - The `resourceTypes` parameter is a JSON array of resource type strings.
//
// - The `startTime` and `endTime` parameters are timestamps that define the query\\"s time range.
//
// - The `nextToken`, `maxResults`, `pageNo`, and `pageSize` parameters control pagination and the number of results to return.
//
// - The API throws a `SalesClientException` if the `resourceTypes` parameter contains an invalid value or if the `periods` parameter fails JSON parsing.
//
// @param request - DescribeDeductionStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeductionStatisticResponse
func (client *Client) DescribeDeductionStatisticWithOptions(request *DescribeDeductionStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeductionStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Periods) {
		query["Periods"] = request.Periods
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeductionStatistic"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeductionStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves resource deduction and usage statistics based on specified criteria.
//
// Description:
//
// ## Request
//
// - This API supports GET and POST methods.
//
// - The `periods` parameter is a JSON array of `PeriodParam` objects, each containing the `periodUnit` and `baseTime` fields.
//
// - The `resourceTypes` parameter is a JSON array of resource type strings.
//
// - The `startTime` and `endTime` parameters are timestamps that define the query\\"s time range.
//
// - The `nextToken`, `maxResults`, `pageNo`, and `pageSize` parameters control pagination and the number of results to return.
//
// - The API throws a `SalesClientException` if the `resourceTypes` parameter contains an invalid value or if the `periods` parameter fails JSON parsing.
//
// @param request - DescribeDeductionStatisticRequest
//
// @return DescribeDeductionStatisticResponse
func (client *Client) DescribeDeductionStatistic(request *DescribeDeductionStatisticRequest) (_result *DescribeDeductionStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeductionStatisticResponse{}
	_body, _err := client.DescribeDeductionStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about delivery addresses.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddressWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDeliveryAddressResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeliveryAddress"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeliveryAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about delivery addresses.
//
// @return DescribeDeliveryAddressResponse
func (client *Client) DescribeDeliveryAddress() (_result *DescribeDeliveryAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeliveryAddressResponse{}
	_body, _err := client.DescribeDeliveryAddressWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries prices for Elastic Desktop Service products, covering order types such as purchase, renewal, configuration change, and cancellation.
//
// Description:
//
// <props="china">
//
// Before using this interface, ensure you understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) for Wuying Workspace.
//
// @param request - DescribeMultiPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiPriceResponse
func (client *Client) DescribeMultiPriceWithOptions(request *DescribeMultiPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultiPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderItems) {
		query["OrderItems"] = request.OrderItems
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PackageCode) {
		query["PackageCode"] = request.PackageCode
	}

	if !dara.IsNil(request.ResellerOwnerUid) {
		query["ResellerOwnerUid"] = request.ResellerOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMultiPrice"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMultiPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries prices for Elastic Desktop Service products, covering order types such as purchase, renewal, configuration change, and cancellation.
//
// Description:
//
// <props="china">
//
// Before using this interface, ensure you understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) for Wuying Workspace.
//
// @param request - DescribeMultiPriceRequest
//
// @return DescribeMultiPriceResponse
func (client *Client) DescribeMultiPrice(request *DescribeMultiPriceRequest) (_result *DescribeMultiPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMultiPriceResponse{}
	_body, _err := client.DescribeMultiPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query deduction details for time-based packages.
//
// @param request - DescribePackageDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductionsWithOptions(request *DescribePackageDeductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribePackageDeductionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PackageIds) {
		query["PackageIds"] = request.PackageIds
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePackageDeductions"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackageDeductionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query deduction details for time-based packages.
//
// @param request - DescribePackageDeductionsRequest
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductions(request *DescribePackageDeductionsRequest) (_result *DescribePackageDeductionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackageDeductionsResponse{}
	_body, _err := client.DescribePackageDeductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an instance.
//
// @param request - ModifyInstancePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstancePropertiesWithOptions(request *ModifyInstancePropertiesRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstancePropertiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceProperties"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstancePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an instance.
//
// @param request - ModifyInstancePropertiesRequest
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstanceProperties(request *ModifyInstancePropertiesRequest) (_result *ModifyInstancePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstancePropertiesResponse{}
	_body, _err := client.ModifyInstancePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the credit quota for specified Agents.
//
// Description:
//
// ## Description
//
// This operation sets the credit quota for one or more Agents of a specific type.
//
// ### Usage notes
//
// - The `AgentType` parameter specifies the type of Agent to which the quota applies, such as `JVSClaw` or `OpenClaw`.
//
// - The `AgentIds` parameter is an array of up to 100 Agent IDs.
//
// - The `CreditQuota` parameter specifies the credit quota for each Agent.
//
// ### Examples
//
// @param request - SetAgentCreditQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAgentCreditQuotaResponse
func (client *Client) SetAgentCreditQuotaWithOptions(request *SetAgentCreditQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetAgentCreditQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentIds) {
		query["AgentIds"] = request.AgentIds
	}

	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CreditQuota) {
		query["CreditQuota"] = request.CreditQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAgentCreditQuota"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAgentCreditQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the credit quota for specified Agents.
//
// Description:
//
// ## Description
//
// This operation sets the credit quota for one or more Agents of a specific type.
//
// ### Usage notes
//
// - The `AgentType` parameter specifies the type of Agent to which the quota applies, such as `JVSClaw` or `OpenClaw`.
//
// - The `AgentIds` parameter is an array of up to 100 Agent IDs.
//
// - The `CreditQuota` parameter specifies the credit quota for each Agent.
//
// ### Examples
//
// @param request - SetAgentCreditQuotaRequest
//
// @return SetAgentCreditQuotaResponse
func (client *Client) SetAgentCreditQuota(request *SetAgentCreditQuotaRequest) (_result *SetAgentCreditQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAgentCreditQuotaResponse{}
	_body, _err := client.SetAgentCreditQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
