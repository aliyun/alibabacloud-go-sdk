// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Places, renews, or modifies orders for specific products of Elastic Desktop Service (EDS) Enterprise Edition, such as monthly duration packages.
//
// Description:
//
// <props="china">Before you use this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws) of EDS.
//
// <props="intl">Before you use this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/cloud-desktop?#J_8623712560) of EDS.
//
// If automatic payment is not specified, this operation does not process the payment. You must use the order ID returned by this operation to construct a payment redirect URL and complete the payment before the order takes effect and the resources are provisioned.
//
// @param tmpReq - CreateMultiOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiOrderResponse
func (client *Client) CreateMultiOrderWithContext(ctx context.Context, tmpReq *CreateMultiOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Agents and usage summary information under the current username.
//
// Description:
//
// ## Request description
//
// - **Paging support**: Use the `NextToken` and `MaxResults` parameters for paging. Set `NextToken` to an empty character string for the first request.
//
// - **Filtering**: Use the `AgentType` and `AgentIds` parameters to filter the returned Agent list.
//
// - **Status filtering**: Use the `Status` parameter to filter Agents by status (0: deleted, 1: active).
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
func (client *Client) DescribeCreditPackageAgentsWithContext(ctx context.Context, request *DescribeCreditPackageAgentsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreditPackageAgentsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries credit usage details by a specified dimension such as user, credit package, or agent.
//
// Description:
//
// ## Request description
//
// This API queries credit usage details based on the dimension specified by `UsageType` (User / CreditPackage / Agent). The response includes the total credits, remaining credits, used credits, hourly consumption samples, alert thresholds, and period quotas of the current credit package.
//
// - **User**: User dimension. Returns the aggregated usage and remaining credits of all active credit packages for the current user.
//
// - **CreditPackage**: Credit package dimension. Returns the total credits, remaining credits, and consumption samples of a specified credit package instance.
//
// - **Agent**: Agent dimension. Returns the cumulative usage, current period usage, quota, and alert information of a specified agent.
//
// **Notes**:
//
// - The `InstanceIds` parameter can be omitted when `UsageType=User`. Pass a credit package instance ID when `UsageType=CreditPackage`, or pass an AgentId when `UsageType=Agent`.
//
// - Anonymous requests support the `FillInstance` parameter. When `InstanceIds` is not explicitly provided and `FillInstance=true`, the server automatically populates the bound `JVS_COPILOT` AgentId based on the current logon `wyId`.
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
func (client *Client) DescribeCreditUsageInfoWithContext(ctx context.Context, request *DescribeCreditUsageInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreditUsageInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - When the `resourceTypes` parameter contains invalid values or the `periods` parameter fails JSON parsing, a `SalesClientException` exception is thrown.
//
// @param request - DescribeDeductionStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeductionStatisticResponse
func (client *Client) DescribeDeductionStatisticWithContext(ctx context.Context, request *DescribeDeductionStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeductionStatisticResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prices of Elastic Desktop Service products, including prices for new purchases, renewals, specification changes, and unsubscriptions.
//
// Description:
//
// <props="china">Before you call this operation, make sure that you fully understand the billing of Elastic Desktop Service and its [pricing](https://www.aliyun.com/price/product?#/gws/detail/gws).
//
// @param request - DescribeMultiPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiPriceResponse
func (client *Client) DescribeMultiPriceWithContext(ctx context.Context, request *DescribeMultiPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultiPriceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deduction details of a core-hour package.
//
// @param request - DescribePackageDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageDeductionsResponse
func (client *Client) DescribePackageDeductionsWithContext(ctx context.Context, request *DescribePackageDeductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribePackageDeductionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// ## Operation description
//
// - This operation supports GET and POST methods.
//
// - The `periods` parameter is in JSON array format. Each element is a `PeriodParam` object that contains the `periodUnit` and `baseTime` fields.
//
// - The `resourceTypes` parameter is in JSON array format and contains multiple resource type strings.
//
// - The `startTime` and `endTime` parameters are in timestamp format and specify the time range for the query.
//
// - Pagination parameters include `nextToken`, `maxResults`, `pageNo`, and `pageSize`, which control the number of returned results and pagination.
//
// - When the `resourceTypes` parameter contains invalid values or the `periods` parameter fails JSON parsing, a `SalesClientException` exception is thrown.
//
// @param request - DescribeRunIdDeductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRunIdDeductionsResponse
func (client *Client) DescribeRunIdDeductionsWithContext(ctx context.Context, request *DescribeRunIdDeductionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRunIdDeductionsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIdType) {
		query["InstanceIdType"] = request.InstanceIdType
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePropertiesResponse
func (client *Client) ModifyInstancePropertiesWithContext(ctx context.Context, request *ModifyInstancePropertiesRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstancePropertiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAgentCreditQuotaResponse
func (client *Client) SetAgentCreditQuotaWithContext(ctx context.Context, request *SetAgentCreditQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetAgentCreditQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
