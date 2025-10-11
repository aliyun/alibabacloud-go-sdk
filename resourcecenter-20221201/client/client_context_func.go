// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilterWithContext(ctx context.Context, request *AssociateDefaultFilterRequest, runtime *dara.RuntimeOptions) (_result *AssociateDefaultFilterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterName) {
		query["FilterName"] = request.FilterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateDefaultFilter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the ResourceSnapshotDelivery.CustomExpression parameter empty.
//
//   - Custom delivery: Set the ResourceSnapshotDelivery.CustomExpression parameter to an appropriate value.
//
// @param request - CreateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryChannelResponse
func (client *Client) CreateDeliveryChannelWithContext(ctx context.Context, request *CreateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelDescription) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !dara.IsNil(request.DeliveryChannelFilter) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.ResourceChangeDelivery) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !dara.IsNil(request.ResourceSnapshotDelivery) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a filter.
//
// @param request - CreateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilterResponse
func (client *Client) CreateFilterWithContext(ctx context.Context, request *CreateFilterRequest, runtime *dara.RuntimeOptions) (_result *CreateFilterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterConfiguration) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !dara.IsNil(request.FilterName) {
		query["FilterName"] = request.FilterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFilter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a multi-account delivery channel.
//
// Description:
//
// In Resource Center, you can create multi-account delivery channels by using the management account of your resource directory or the delegated administrator account of Resource Center to deliver resource configuration change events and scheduled resource snapshots within the members in your resource directory to Object Storage Service (OSS) or Simple Log Service. Then, other Alibaba Cloud services consume standardized resource information from OSS or Simple Log Service.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - CreateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiAccountDeliveryChannelResponse
func (client *Client) CreateMultiAccountDeliveryChannelWithContext(ctx context.Context, request *CreateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelDescription) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !dara.IsNil(request.DeliveryChannelFilter) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.ResourceChangeDelivery) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !dara.IsNil(request.ResourceSnapshotDelivery) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultiAccountDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom query template.
//
// @param request - CreateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQueryWithContext(ctx context.Context, request *CreateSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *CreateSavedQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSavedQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a single-account delivery channel.
//
// @param request - DeleteDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryChannelResponse
func (client *Client) DeleteDeliveryChannelWithContext(ctx context.Context, request *DeleteDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a filter.
//
// @param request - DeleteFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilterResponse
func (client *Client) DeleteFilterWithContext(ctx context.Context, request *DeleteFilterRequest, runtime *dara.RuntimeOptions) (_result *DeleteFilterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterName) {
		query["FilterName"] = request.FilterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFilter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-account delivery channel.
//
// @param request - DeleteMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiAccountDeliveryChannelResponse
func (client *Client) DeleteMultiAccountDeliveryChannelWithContext(ctx context.Context, request *DeleteMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultiAccountDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultiAccountDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom query template.
//
// @param request - DeleteSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQueryWithContext(ctx context.Context, request *DeleteSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSavedQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSavedQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query resources across accounts.
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQueryWithContext(ctx context.Context, request *ExecuteMultiAccountSQLQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteMultiAccountSQLQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query the resources that can be accessed within the current account.
//
// @param request - ExecuteSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQueryWithContext(ctx context.Context, request *ExecuteSQLQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSQLQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSQLQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a single-account delivery channel.
//
// @param request - GetDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelResponse
func (client *Client) GetDeliveryChannelWithContext(ctx context.Context, request *GetDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on a single-account delivery channel.
//
// @param request - GetDeliveryChannelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelStatisticsResponse
func (client *Client) GetDeliveryChannelStatisticsWithContext(ctx context.Context, request *GetDeliveryChannelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeliveryChannelStatistics"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeliveryChannelStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a sample query template.
//
// @param request - GetExampleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExampleQueryResponse
func (client *Client) GetExampleQueryWithContext(ctx context.Context, request *GetExampleQueryRequest, runtime *dara.RuntimeOptions) (_result *GetExampleQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExampleQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountDeliveryChannelResponse
func (client *Client) GetMultiAccountDeliveryChannelWithContext(ctx context.Context, request *GetMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiAccountDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on a multi-account delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountDeliveryChannelStatisticsResponse
func (client *Client) GetMultiAccountDeliveryChannelStatisticsWithContext(ctx context.Context, request *GetMultiAccountDeliveryChannelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiAccountDeliveryChannelStatistics"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiAccountDeliveryChannelStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the management account or a member of a resource directory.
//
// @param request - GetMultiAccountResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfigurationWithContext(ctx context.Context, request *GetMultiAccountResourceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiAccountResourceConfiguration"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取多账号资源数量
//
// @param request - GetMultiAccountResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCountsResponse
func (client *Client) GetMultiAccountResourceCountsWithContext(ctx context.Context, request *GetMultiAccountResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceCountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.GroupByKey) {
		query["GroupByKey"] = request.GroupByKey
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiAccountResourceCounts"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiAccountResourceCountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the current account.
//
// @param request - GetResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfigurationWithContext(ctx context.Context, request *GetResourceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceConfiguration"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of resources on which the current account has access permissions.
//
// @param request - GetResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCountsWithContext(ctx context.Context, request *GetResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceCountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.GroupByKey) {
		query["GroupByKey"] = request.GroupByKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceCounts"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom query template.
//
// @param request - GetSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavedQueryResponse
func (client *Client) GetSavedQueryWithContext(ctx context.Context, request *GetSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *GetSavedQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavedQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of single-account delivery channels.
//
// @param request - ListDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryChannelsResponse
func (client *Client) ListDeliveryChannelsWithContext(ctx context.Context, request *ListDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryChannelsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeliveryChannels"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeliveryChannelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all sample query templates.
//
// @param request - ListExampleQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueriesWithContext(ctx context.Context, request *ListExampleQueriesRequest, runtime *dara.RuntimeOptions) (_result *ListExampleQueriesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExampleQueries"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of multi-account delivery channels.
//
// @param request - ListMultiAccountDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountDeliveryChannelsResponse
func (client *Client) ListMultiAccountDeliveryChannelsWithContext(ctx context.Context, request *ListMultiAccountDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiAccountDeliveryChannels"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiAccountDeliveryChannelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource groups within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroupsWithContext(ctx context.Context, request *ListMultiAccountResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiAccountResourceGroups"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationships between resources within the management account or members of your resource directory.
//
// Description:
//
//	  Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// @param request - ListMultiAccountResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceRelationshipsResponse
func (client *Client) ListMultiAccountResourceRelationshipsWithContext(ctx context.Context, request *ListMultiAccountResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RelatedResourceFilter) {
		query["RelatedResourceFilter"] = request.RelatedResourceFilter
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiAccountResourceRelationships"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiAccountResourceRelationshipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the management account or a member of your resource directory.
//
// @param request - ListMultiAccountTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeysWithContext(ctx context.Context, request *ListMultiAccountTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiAccountTagKeys"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValuesWithContext(ctx context.Context, request *ListMultiAccountTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagValue) {
		query["TagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiAccountTagValues"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resource relationships on which the current account has access permissions.
//
// Description:
//
//	  You can call this operation to query only the resource relationships on which the current account has access permissions.
//
//		- By default, this operation returns up to 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the query scope. For information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only entries that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Entries that meet any value of the filter condition are returned.
//
// @param request - ListResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRelationshipsResponse
func (client *Client) ListResourceRelationshipsWithContext(ctx context.Context, request *ListResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceRelationshipsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RelatedResourceFilter) {
		query["RelatedResourceFilter"] = request.RelatedResourceFilter
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
		Action:      dara.String("ListResourceRelationships"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the metadata of resource types
//
// @param request - ListResourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithContext(ctx context.Context, request *ListResourceTypesRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypes"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom query templates.
//
// @param request - ListSavedQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueriesWithContext(ctx context.Context, request *ListSavedQueriesRequest, runtime *dara.RuntimeOptions) (_result *ListSavedQueriesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSavedQueries"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the current account.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the current account.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagValue) {
		query["TagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for resources within the management account or members of a resource directory.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.alibabacloud.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchMultiAccountResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResourcesWithContext(ctx context.Context, request *SearchMultiAccountResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SortCriterion) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMultiAccountResources"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for resources that you can access within the current account.
//
// Description:
//
//	  You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
//		- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
//		- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
//		- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
//		- You can visit [Sample Code Center](https://api.aliyun.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResourcesResponse
func (client *Client) SearchResourcesWithContext(ctx context.Context, request *SearchResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SortCriterion) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchResources"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a single-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeliveryChannelResponse
func (client *Client) UpdateDeliveryChannelWithContext(ctx context.Context, request *UpdateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelDescription) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !dara.IsNil(request.DeliveryChannelFilter) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.ResourceChangeDelivery) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !dara.IsNil(request.ResourceSnapshotDelivery) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a filter.
//
// @param request - UpdateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFilterResponse
func (client *Client) UpdateFilterWithContext(ctx context.Context, request *UpdateFilterRequest, runtime *dara.RuntimeOptions) (_result *UpdateFilterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterConfiguration) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !dara.IsNil(request.FilterName) {
		query["FilterName"] = request.FilterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFilter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a multi-account delivery channel.
//
// Description:
//
// Resource delivery supports the delivery of resource configuration change events and scheduled resource snapshots.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
//   - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
//   - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiAccountDeliveryChannelResponse
func (client *Client) UpdateMultiAccountDeliveryChannelWithContext(ctx context.Context, request *UpdateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryChannelDescription) {
		query["DeliveryChannelDescription"] = request.DeliveryChannelDescription
	}

	if !dara.IsNil(request.DeliveryChannelFilter) {
		query["DeliveryChannelFilter"] = request.DeliveryChannelFilter
	}

	if !dara.IsNil(request.DeliveryChannelId) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !dara.IsNil(request.DeliveryChannelName) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !dara.IsNil(request.ResourceChangeDelivery) {
		query["ResourceChangeDelivery"] = request.ResourceChangeDelivery
	}

	if !dara.IsNil(request.ResourceSnapshotDelivery) {
		query["ResourceSnapshotDelivery"] = request.ResourceSnapshotDelivery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultiAccountDeliveryChannel"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom query template.
//
// @param request - UpdateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQueryWithContext(ctx context.Context, request *UpdateSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *UpdateSavedQueryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSavedQuery"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
