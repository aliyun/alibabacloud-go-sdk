// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("wss")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
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
// Creates an order, renews, or modifies the configuration of specific products such as the monthly duration package for Wuying Workspace Enterprise Edition.
//
// Description:
//
// <props="china">Make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Wuying Workspace before you call this operation.
//
// <props="intl">Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/cloud-desktop?#J_8623712560) of Wuying Workspace before you call this operation.
//
// If automatic payment is not specified, this operation does not process the payment. You must use the order ID returned by this operation to construct a payment redirect URL and complete the payment. The order takes effect and resources are provisioned only after the payment is completed.
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
// Creates an order, renews, or modifies the configuration of specific products such as the monthly duration package for Wuying Workspace Enterprise Edition.
//
// Description:
//
// <props="china">Make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Wuying Workspace before you call this operation.
//
// <props="intl">Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/cloud-desktop?#J_8623712560) of Wuying Workspace before you call this operation.
//
// If automatic payment is not specified, this operation does not process the payment. You must use the order ID returned by this operation to construct a payment redirect URL and complete the payment. The order takes effect and resources are provisioned only after the payment is completed.
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
// 查询 AI 调用明细
//
// Description:
//
// ## 请求说明
//
// - 该接口用于分页查询当前阿里云账号下指定业务版本、地域和时间范围内的AI模型调用明细。
//
// - 支持按模型调用请求ID、请求模型和用户ID进行筛选。
//
// - 查询结果按调用时间倒序返回。
//
// - 默认查询北京时间当日零点至当前时间的数据，查询区间不超过30天且不得早于当前时间向前推算3个自然月。
//
// - 调用方仅能查询当前调用身份所属主账号的数据，不能指定其他账号。
//
// - `PageNumber`与`PageSize`的乘积不得超过1000000。
//
// - `BizType`不传时默认为企业版。
//
// - `GatewayRegion`不传时国内站默认上海，国际站默认新加坡。
//
// @param request - DescribeAICallDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAICallDetailsResponse
func (client *Client) DescribeAICallDetailsWithOptions(request *DescribeAICallDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAICallDetailsResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GatewayRegion) {
		query["GatewayRegion"] = request.GatewayRegion
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchRequestId) {
		query["SearchRequestId"] = request.SearchRequestId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WyIds) {
		query["WyIds"] = request.WyIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAICallDetails"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAICallDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 AI 调用明细
//
// Description:
//
// ## 请求说明
//
// - 该接口用于分页查询当前阿里云账号下指定业务版本、地域和时间范围内的AI模型调用明细。
//
// - 支持按模型调用请求ID、请求模型和用户ID进行筛选。
//
// - 查询结果按调用时间倒序返回。
//
// - 默认查询北京时间当日零点至当前时间的数据，查询区间不超过30天且不得早于当前时间向前推算3个自然月。
//
// - 调用方仅能查询当前调用身份所属主账号的数据，不能指定其他账号。
//
// - `PageNumber`与`PageSize`的乘积不得超过1000000。
//
// - `BizType`不传时默认为企业版。
//
// - `GatewayRegion`不传时国内站默认上海，国际站默认新加坡。
//
// @param request - DescribeAICallDetailsRequest
//
// @return DescribeAICallDetailsResponse
func (client *Client) DescribeAICallDetails(request *DescribeAICallDetailsRequest) (_result *DescribeAICallDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAICallDetailsResponse{}
	_body, _err := client.DescribeAICallDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of agents and usage summary information under the current username.
//
// Description:
//
// ## Operation description
//
// - **Paging**: Use the `NextToken` and `MaxResults` parameters for paging. Set `NextToken` to an empty character string for the first request.
//
// - **Filtering**: Use the `AgentType` and `AgentIds` parameters to filter the returned agent list.
//
// - **Status filtering**: Use the `Status` parameter to filter agents by status. Valid values: 0 (deleted) and 1 (active).
//
// - **Sorting**: Results are sorted by `id` in ascending order by default.
//
// - **Additional parameter for anonymous edition**: The `FillInstance` parameter automatically populates the bound JVS_COPILOT AgentId of the currently logged-on user.
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
// Queries the list of agents and usage summary information under the current username.
//
// Description:
//
// ## Operation description
//
// - **Paging**: Use the `NextToken` and `MaxResults` parameters for paging. Set `NextToken` to an empty character string for the first request.
//
// - **Filtering**: Use the `AgentType` and `AgentIds` parameters to filter the returned agent list.
//
// - **Status filtering**: Use the `Status` parameter to filter agents by status. Valid values: 0 (deleted) and 1 (active).
//
// - **Sorting**: Results are sorted by `id` in ascending order by default.
//
// - **Additional parameter for anonymous edition**: The `FillInstance` parameter automatically populates the bound JVS_COPILOT AgentId of the currently logged-on user.
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
// Queries the credit usage of users, credit packages, or agents based on a specified dimension.
//
// Description:
//
// ## Request description
//
// This API queries credit usage details based on the dimension specified by `UsageType` (User, CreditPackage, Agent, or EndUser). The response includes the total amount, remaining amount, and used amount of the current credit package, hourly consumption samples, alert thresholds, and periodic quotas.
//
// - **User**: The user dimension. Returns the total used amount and remaining amount of all active credit packages for the current user.
//
// - **CreditPackage**: The credit package dimension. Returns the total amount, remaining amount, and consumption samples of a specified credit package instance.
//
// - **Agent**: The agent dimension. Returns the cumulative used amount, current period used amount, quota, and alerts for a specified agent.
//
// - **EndUser**: The Wuying user dimension. Returns the shared credit usage, pay-as-you-go credit usage, and total credit usage for the current period of a Wuying user.
//
// **Note**:
//
// - The `InstanceIds` parameter can be omitted when `UsageType` is set to `User`. Set it to the credit package instance ID when `UsageType` is `CreditPackage`, the agent ID when `UsageType` is `Agent`, and the Wuying ID when `UsageType` is `EndUser`.
//
// - Anonymous requests support the `FillInstance` parameter. If `InstanceIds` is not explicitly provided and `FillInstance` is set to `true`, the server automatically populates the bound `JVS_COPILOT` agent ID based on the currently logged-in `wyId`.
//
// - Time window constants: The statistical window for `dayUsedCredit` is `now - ONE_DAY_MILLIS`, and the statistical window for `weekUsedCredit` is `now - ONE_WEEK_MILLIS`.
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

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
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
// Queries the credit usage of users, credit packages, or agents based on a specified dimension.
//
// Description:
//
// ## Request description
//
// This API queries credit usage details based on the dimension specified by `UsageType` (User, CreditPackage, Agent, or EndUser). The response includes the total amount, remaining amount, and used amount of the current credit package, hourly consumption samples, alert thresholds, and periodic quotas.
//
// - **User**: The user dimension. Returns the total used amount and remaining amount of all active credit packages for the current user.
//
// - **CreditPackage**: The credit package dimension. Returns the total amount, remaining amount, and consumption samples of a specified credit package instance.
//
// - **Agent**: The agent dimension. Returns the cumulative used amount, current period used amount, quota, and alerts for a specified agent.
//
// - **EndUser**: The Wuying user dimension. Returns the shared credit usage, pay-as-you-go credit usage, and total credit usage for the current period of a Wuying user.
//
// **Note**:
//
// - The `InstanceIds` parameter can be omitted when `UsageType` is set to `User`. Set it to the credit package instance ID when `UsageType` is `CreditPackage`, the agent ID when `UsageType` is `Agent`, and the Wuying ID when `UsageType` is `EndUser`.
//
// - Anonymous requests support the `FillInstance` parameter. If `InstanceIds` is not explicitly provided and `FillInstance` is set to `true`, the server automatically populates the bound `JVS_COPILOT` agent ID based on the currently logged-in `wyId`.
//
// - Time window constants: The statistical window for `dayUsedCredit` is `now - ONE_DAY_MILLIS`, and the statistical window for `weekUsedCredit` is `now - ONE_WEEK_MILLIS`.
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
// Retrieves resource deduction and usage statistics based on specified conditions.
//
// Description:
//
// ## Request description
//
// - This operation supports GET and POST methods.
//
// - The `periods` parameter is in JSON array format. Each element is a `PeriodParam` object that contains the `periodUnit` and `baseTime` fields.
//
// - The `resourceTypes` parameter is in JSON array format and contains multiple resource type strings.
//
// - The `startTime` and `endTime` parameters are in timestamp format and specify the time range for the query.
//
// - Pagination parameters include `nextToken`, `maxResults`, `pageNo`, and `pageSize`, which control the number and pagination of returned results.
//
// - If the `resourceTypes` parameter contains invalid values or the `periods` parameter fails JSON parsing, a `SalesClientException` exception is thrown.
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
// Retrieves resource deduction and usage statistics based on specified conditions.
//
// Description:
//
// ## Request description
//
// - This operation supports GET and POST methods.
//
// - The `periods` parameter is in JSON array format. Each element is a `PeriodParam` object that contains the `periodUnit` and `baseTime` fields.
//
// - The `resourceTypes` parameter is in JSON array format and contains multiple resource type strings.
//
// - The `startTime` and `endTime` parameters are in timestamp format and specify the time range for the query.
//
// - Pagination parameters include `nextToken`, `maxResults`, `pageNo`, and `pageSize`, which control the number and pagination of returned results.
//
// - If the `resourceTypes` parameter contains invalid values or the `periods` parameter fails JSON parsing, a `SalesClientException` exception is thrown.
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
// Queries logistics address information.
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
// Queries logistics address information.
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
// Queries deduction group information.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Elastic Desktop Service.
//
// @param request - DescribeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
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

	if !dara.IsNil(request.GroupResourceTypes) {
		query["GroupResourceTypes"] = request.GroupResourceTypes
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.States) {
		query["States"] = request.States
	}

	if !dara.IsNil(request.TimeUsingPeriod) {
		query["TimeUsingPeriod"] = request.TimeUsingPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroups"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries deduction group information.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Elastic Desktop Service.
//
// @param request - DescribeGroupsRequest
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroups(request *DescribeGroupsRequest) (_result *DescribeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.DescribeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the renewal methods.
//
// Description:
//
// Filters instances by resource type, instance, and region, and uses paging to query the auto-renewal status.
//
// @param request - DescribeInstanceAutoRenewalAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAutoRenewalAttributeResponse
func (client *Client) DescribeInstanceAutoRenewalAttributeWithOptions(request *DescribeInstanceAutoRenewalAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewal) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceTypeList) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !dara.IsNil(request.SearchRegionId) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAutoRenewalAttribute"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the renewal methods.
//
// Description:
//
// Filters instances by resource type, instance, and region, and uses paging to query the auto-renewal status.
//
// @param request - DescribeInstanceAutoRenewalAttributeRequest
//
// @return DescribeInstanceAutoRenewalAttributeResponse
func (client *Client) DescribeInstanceAutoRenewalAttribute(request *DescribeInstanceAutoRenewalAttributeRequest) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.DescribeInstanceAutoRenewalAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the prices of Elastic Desktop Service products for order types such as new purchases, renewals, configuration changes, and unsubscriptions.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Elastic Desktop Service.
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
// Queries the prices of Elastic Desktop Service products for order types such as new purchases, renewals, configuration changes, and unsubscriptions.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of Elastic Desktop Service.
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
// Queries the deduction details of core-hour packages.
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
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.AgentTypes) {
		query["AgentTypes"] = request.AgentTypes
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ConsumeWyId) {
		query["ConsumeWyId"] = request.ConsumeWyId
	}

	if !dara.IsNil(request.ConsumeWyIds) {
		query["ConsumeWyIds"] = request.ConsumeWyIds
	}

	if !dara.IsNil(request.DeductionRequestId) {
		query["DeductionRequestId"] = request.DeductionRequestId
	}

	if !dara.IsNil(request.DeductionTypes) {
		query["DeductionTypes"] = request.DeductionTypes
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupResourceTypes) {
		query["GroupResourceTypes"] = request.GroupResourceTypes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdType) {
		query["InstanceIdType"] = request.InstanceIdType
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
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

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.RunIdList) {
		query["RunIdList"] = request.RunIdList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WyId) {
		query["WyId"] = request.WyId
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
// Queries the deduction details of core-hour packages.
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
// Queries AI credit deductions.
//
// Description:
//
// ## Request description
//
// - This operation supports the GET and POST methods.
//
// - The periods parameter is in JSON array format. Each element is a PeriodParam object that contains the periodUnit and baseTime fields.
//
// - The resourceTypes parameter is in JSON array format and contains multiple resource type strings.
//
// - The startTime and endTime parameters are in timestamp format and specify the time range for the query.
//
// - The pagination parameters include nextToken, maxResults, pageNo, and pageSize, which are used to control the number of returned results and pagination.
//
// - If the resourceTypes parameter contains invalid values or the JSON parsing of the periods parameter fails, a SalesClientException is thrown.
//
// @param request - DescribeRunIdDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRunIdDeductionsResponse
func (client *Client) DescribeRunIdDeductionsWithOptions(request *DescribeRunIdDeductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRunIdDeductionsResponse, _err error) {
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

	if !dara.IsNil(request.AgentTypes) {
		query["AgentTypes"] = request.AgentTypes
	}

	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ConsumeWyId) {
		query["ConsumeWyId"] = request.ConsumeWyId
	}

	if !dara.IsNil(request.ConsumeWyIds) {
		query["ConsumeWyIds"] = request.ConsumeWyIds
	}

	if !dara.IsNil(request.DeductionRequestId) {
		query["DeductionRequestId"] = request.DeductionRequestId
	}

	if !dara.IsNil(request.DeductionTypes) {
		query["DeductionTypes"] = request.DeductionTypes
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupByFields) {
		query["GroupByFields"] = request.GroupByFields
	}

	if !dara.IsNil(request.GroupResourceTypes) {
		query["GroupResourceTypes"] = request.GroupResourceTypes
	}

	if !dara.IsNil(request.GroupSeparator) {
		query["GroupSeparator"] = request.GroupSeparator
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdType) {
		query["InstanceIdType"] = request.InstanceIdType
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.RunId) {
		query["RunId"] = request.RunId
	}

	if !dara.IsNil(request.RunIdList) {
		query["RunIdList"] = request.RunIdList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WyId) {
		query["WyId"] = request.WyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRunIdDeductions"),
		Version:     dara.String("2021-12-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRunIdDeductionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries AI credit deductions.
//
// Description:
//
// ## Request description
//
// - This operation supports the GET and POST methods.
//
// - The periods parameter is in JSON array format. Each element is a PeriodParam object that contains the periodUnit and baseTime fields.
//
// - The resourceTypes parameter is in JSON array format and contains multiple resource type strings.
//
// - The startTime and endTime parameters are in timestamp format and specify the time range for the query.
//
// - The pagination parameters include nextToken, maxResults, pageNo, and pageSize, which are used to control the number of returned results and pagination.
//
// - If the resourceTypes parameter contains invalid values or the JSON parsing of the periods parameter fails, a SalesClientException is thrown.
//
// @param request - DescribeRunIdDeductionsRequest
//
// @return DescribeRunIdDeductionsResponse
func (client *Client) DescribeRunIdDeductions(request *DescribeRunIdDeductionsRequest) (_result *DescribeRunIdDeductionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRunIdDeductionsResponse{}
	_body, _err := client.DescribeRunIdDeductionsWithOptions(request, runtime)
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

	if !dara.IsNil(request.Properties) {
		query["Properties"] = request.Properties
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
// Sets the credit usage quota for a specified Agent.
//
// Description:
//
// ## Operation description
//
// You can use this API operation to set the credit usage quota for a specific type of Agent. Batch setting of quotas for multiple Agents is supported.
//
// ### Before you begin
//
// - The `AgentType` parameter specifies the type of Agent for which you want to set the quota, such as `JVSClaw` or `OpenClaw`.
//
// - The `AgentIds` parameter is an array that can contain up to 100 Agent IDs.
//
// - The `CreditQuota` parameter specifies the credit usage quota for each Agent.
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

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
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
// Sets the credit usage quota for a specified Agent.
//
// Description:
//
// ## Operation description
//
// You can use this API operation to set the credit usage quota for a specific type of Agent. Batch setting of quotas for multiple Agents is supported.
//
// ### Before you begin
//
// - The `AgentType` parameter specifies the type of Agent for which you want to set the quota, such as `JVSClaw` or `OpenClaw`.
//
// - The `AgentIds` parameter is an array that can contain up to 100 Agent IDs.
//
// - The `CreditQuota` parameter specifies the credit usage quota for each Agent.
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
