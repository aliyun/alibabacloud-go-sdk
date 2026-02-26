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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the configurations of multiple resources in your account.
//
// @param request - BatchGetResourceConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetResourceConfigurationsResponse
func (client *Client) BatchGetResourceConfigurationsWithContext(ctx context.Context, request *BatchGetResourceConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *BatchGetResourceConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetResourceConfigurations"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetResourceConfigurationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource delivery channel for the current account.
//
// Description:
//
// Resource delivery supports the scheduled delivery of resource snapshots and the delivery of resource configuration changes.
//
// Scheduled delivery of resource snapshots supports two scenarios:
//
// - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
// - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - CreateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryChannelResponse
func (client *Client) CreateDeliveryChannelWithContext(ctx context.Context, request *CreateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// In Resource Center, you can create multi-account delivery channels by using the management account of your resource directory or the delegated administrator account of Resource Center to deliver resource configuration change events and scheduled resource snapshots within the members in your resource directory to Object Storage Service (OSS) or Simple Log Service (SLS). Then, other Alibaba Cloud services consume standardized resource information from OSS or Simple Log Service.
//
// Scheduled resource snapshots support the following delivery scenarios:
//
// - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
// - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - CreateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiAccountDeliveryChannelResponse
func (client *Client) CreateMultiAccountDeliveryChannelWithContext(ctx context.Context, request *CreateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 执行多账号查询
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQueryWithContext(ctx context.Context, request *ExecuteMultiAccountSQLQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the information about a delivery channel within the current account.
//
// @param request - GetDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelResponse
func (client *Client) GetDeliveryChannelWithContext(ctx context.Context, request *GetDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the statistics of a resource delivery channel in the current account.
//
// @param request - GetDeliveryChannelStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryChannelStatisticsResponse
func (client *Client) GetDeliveryChannelStatisticsWithContext(ctx context.Context, request *GetDeliveryChannelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the information about a cross-account resource delivery channel.
//
// @param request - GetMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountDeliveryChannelResponse
func (client *Client) GetMultiAccountDeliveryChannelWithContext(ctx context.Context, request *GetMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the number of resources within the management account and multiple members of a resource directory.
//
// Description:
//
// You can query only resources supported by Resource Center. For more information, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// @param request - GetMultiAccountResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCountsResponse
func (client *Client) GetMultiAccountResourceCountsWithContext(ctx context.Context, request *GetMultiAccountResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceCountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 获取资源配置
//
// @param request - GetResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfigurationWithContext(ctx context.Context, request *GetResourceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the number of resources in your account that you have permission to access.
//
// Description:
//
// - You can query the number of resources in your account that you have permission to access.
//
// - You can query only the [Alibaba Cloud services and resource types that are supported by Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// @param request - GetResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCountsWithContext(ctx context.Context, request *GetResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceCountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.GroupByKey) {
		query["GroupByKey"] = request.GroupByKey
	}

	if !dara.IsNil(request.IncludeDeletedResources) {
		query["IncludeDeletedResources"] = request.IncludeDeletedResources
	}

	if !dara.IsNil(request.SearchExpression) {
		query["SearchExpression"] = request.SearchExpression
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries a list of delivery channels within the current account.
//
// @param request - ListDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryChannelsResponse
func (client *Client) ListDeliveryChannelsWithContext(ctx context.Context, request *ListDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryChannelsResponse, _err error) {
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
// Queries a list of delivery channels in resource directory mode.
//
// @param request - ListMultiAccountDeliveryChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountDeliveryChannelsResponse
func (client *Client) ListMultiAccountDeliveryChannelsWithContext(ctx context.Context, request *ListMultiAccountDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
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
// Queries the resource groups within the management account or a member in a resource directory.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroupsWithContext(ctx context.Context, request *ListMultiAccountResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 跨账号列出资源关系
//
// @param request - ListMultiAccountResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceRelationshipsResponse
func (client *Client) ListMultiAccountResourceRelationshipsWithContext(ctx context.Context, request *ListMultiAccountResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
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
// 查询多账号标签键
//
// @param request - ListMultiAccountTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeysWithContext(ctx context.Context, request *ListMultiAccountTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 查询多账号标签值
//
// @param request - ListMultiAccountTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValuesWithContext(ctx context.Context, request *ListMultiAccountTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 列出资源关系
//
// @param request - ListResourceRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceRelationshipsResponse
func (client *Client) ListResourceRelationshipsWithContext(ctx context.Context, request *ListResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceRelationshipsResponse, _err error) {
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 查询标签值
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Searches for resources in the management account and multiple member accounts of a resource directory.
//
// Description:
//
// - You can search only for [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// - To search for resources across accounts as a RAM user or RAM role, you must have been attached the `resourcecenter:SearchMultiAccountResources` policy. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
// - By default, this operation returns a maximum of 20 entries. To change the maximum number of entries that can be returned, specify the `MaxResults` parameter.
//
// - If the response does not include `NextToken`, no more data is available. To query the next page of results, set the `NextToken` parameter to the value that was returned from the previous call. If you do not specify the `NextToken` parameter, the first page of data is returned by default.
//
// - You can set one or more filter conditions to narrow the search scope. For information about supported filter parameters and matching methods, see the information below. Multiple filter conditions are joined by a logical `AND`. Only resources that meet all filter conditions are returned. The values within a single filter condition are joined by a logical `OR`. Resources that meet any value for a filter condition are returned.
//
// - For more query examples, visit <props="china">[示例中心](https://api.aliyun.com/api-tools/demo/ResourceCenter) <props="intl">[OpenAPI Portal](https://api.alibabacloud.com/api-tools/demo/ResourceCenter).
//
// @param request - SearchMultiAccountResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResourcesWithContext(ctx context.Context, request *SearchMultiAccountResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Searches for resources in your current account that you are permitted to access.
//
// Description:
//
// - You can search only for resources in your current account that you are permitted to access.
//
// - You can search only for the [Alibaba Cloud services and resource types that are supported by Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// - By default, the SearchResources operation returns a maximum of 20 entries. You can specify the `MaxResults` parameter to change the maximum number of entries that are returned.
//
// - If the response does not include a `NextToken` value, no more results are available. To retrieve the next page of results, include the `NextToken` value from the previous response in your next request. If you do not specify the `NextToken` parameter, the first page of results is returned.
//
// - You can set one or more filter conditions to narrow the search scope. For information about the supported filter parameters and matching methods, see the following sections. Multiple filter conditions are combined by a logical `AND`. Only resources that meet all filter conditions are returned. The values within a filter condition are combined by a logical `OR`. Resources that meet any value of the filter condition are returned.
//
// - For more query examples, see [API Explorer](https://api.aliyun.com/api-tools/demo/ResourceCenter).
//
// @param request - SearchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResourcesResponse
func (client *Client) SearchResourcesWithContext(ctx context.Context, request *SearchResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IncludeDeletedResources) {
		query["IncludeDeletedResources"] = request.IncludeDeletedResources
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

	if !dara.IsNil(request.SearchExpression) {
		query["SearchExpression"] = request.SearchExpression
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
// 更新投递渠道
//
// @param request - UpdateDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeliveryChannelResponse
func (client *Client) UpdateDeliveryChannelWithContext(ctx context.Context, request *UpdateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// - Standard delivery: Leave the `ResourceSnapshotDelivery.CustomExpression` parameter empty.
//
// - Custom delivery: Set the `ResourceSnapshotDelivery.CustomExpression` parameter to an appropriate value.
//
// @param request - UpdateMultiAccountDeliveryChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiAccountDeliveryChannelResponse
func (client *Client) UpdateMultiAccountDeliveryChannelWithContext(ctx context.Context, request *UpdateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
