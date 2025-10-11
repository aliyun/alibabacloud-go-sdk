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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("resourcecenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilterWithOptions(request *AssociateDefaultFilterRequest, runtime *dara.RuntimeOptions) (_result *AssociateDefaultFilterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilter(request *AssociateDefaultFilterRequest) (_result *AssociateDefaultFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.AssociateDefaultFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDeliveryChannelWithOptions(request *CreateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDeliveryChannelResponse
func (client *Client) CreateDeliveryChannel(request *CreateDeliveryChannelRequest) (_result *CreateDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeliveryChannelResponse{}
	_body, _err := client.CreateDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateFilterWithOptions(request *CreateFilterRequest, runtime *dara.RuntimeOptions) (_result *CreateFilterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateFilterResponse
func (client *Client) CreateFilter(request *CreateFilterRequest) (_result *CreateFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFilterResponse{}
	_body, _err := client.CreateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateMultiAccountDeliveryChannelWithOptions(request *CreateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateMultiAccountDeliveryChannelResponse
func (client *Client) CreateMultiAccountDeliveryChannel(request *CreateMultiAccountDeliveryChannelRequest) (_result *CreateMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.CreateMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSavedQueryWithOptions(request *CreateSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *CreateSavedQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQuery(request *CreateSavedQueryRequest) (_result *CreateSavedQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CreateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDeliveryChannelWithOptions(request *DeleteDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDeliveryChannelResponse
func (client *Client) DeleteDeliveryChannel(request *DeleteDeliveryChannelRequest) (_result *DeleteDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDeliveryChannelResponse{}
	_body, _err := client.DeleteDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFilterWithOptions(request *DeleteFilterRequest, runtime *dara.RuntimeOptions) (_result *DeleteFilterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFilterResponse
func (client *Client) DeleteFilter(request *DeleteFilterRequest) (_result *DeleteFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFilterResponse{}
	_body, _err := client.DeleteFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMultiAccountDeliveryChannelWithOptions(request *DeleteMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultiAccountDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteMultiAccountDeliveryChannelResponse
func (client *Client) DeleteMultiAccountDeliveryChannel(request *DeleteMultiAccountDeliveryChannelRequest) (_result *DeleteMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMultiAccountDeliveryChannelResponse{}
	_body, _err := client.DeleteMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSavedQueryWithOptions(request *DeleteSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSavedQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQuery(request *DeleteSavedQueryRequest) (_result *DeleteSavedQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.DeleteSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - DisableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenterWithOptions(runtime *dara.RuntimeOptions) (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableMultiAccountResourceCenter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenter() (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.DisableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @param request - DisableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenterWithOptions(runtime *dara.RuntimeOptions) (_result *DisableResourceCenterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableResourceCenter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenter() (_result *DisableResourceCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.DisableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @param request - DisassociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilterWithOptions(runtime *dara.RuntimeOptions) (_result *DisassociateDefaultFilterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateDefaultFilter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilter() (_result *DisassociateDefaultFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.DisassociateDefaultFilterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @param request - EnableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenterWithOptions(runtime *dara.RuntimeOptions) (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMultiAccountResourceCenter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenter() (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.EnableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @param request - EnableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenterWithOptions(runtime *dara.RuntimeOptions) (_result *EnableResourceCenterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableResourceCenter"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenter() (_result *EnableResourceCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.EnableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteMultiAccountSQLQueryWithOptions(request *ExecuteMultiAccountSQLQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQuery(request *ExecuteMultiAccountSQLQueryRequest) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.ExecuteMultiAccountSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteSQLQueryWithOptions(request *ExecuteSQLQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteSQLQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQuery(request *ExecuteSQLQueryRequest) (_result *ExecuteSQLQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.ExecuteSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeliveryChannelWithOptions(request *GetDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeliveryChannelResponse
func (client *Client) GetDeliveryChannel(request *GetDeliveryChannelRequest) (_result *GetDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeliveryChannelResponse{}
	_body, _err := client.GetDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeliveryChannelStatisticsWithOptions(request *GetDeliveryChannelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeliveryChannelStatisticsResponse
func (client *Client) GetDeliveryChannelStatistics(request *GetDeliveryChannelStatisticsRequest) (_result *GetDeliveryChannelStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeliveryChannelStatisticsResponse{}
	_body, _err := client.GetDeliveryChannelStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetExampleQueryWithOptions(request *GetExampleQueryRequest, runtime *dara.RuntimeOptions) (_result *GetExampleQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetExampleQueryResponse
func (client *Client) GetExampleQuery(request *GetExampleQueryRequest) (_result *GetExampleQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.GetExampleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiAccountDeliveryChannelWithOptions(request *GetMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMultiAccountDeliveryChannelResponse
func (client *Client) GetMultiAccountDeliveryChannel(request *GetMultiAccountDeliveryChannelRequest) (_result *GetMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiAccountDeliveryChannelResponse{}
	_body, _err := client.GetMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiAccountDeliveryChannelStatisticsWithOptions(request *GetMultiAccountDeliveryChannelStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountDeliveryChannelStatisticsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMultiAccountDeliveryChannelStatisticsResponse
func (client *Client) GetMultiAccountDeliveryChannelStatistics(request *GetMultiAccountDeliveryChannelStatisticsRequest) (_result *GetMultiAccountDeliveryChannelStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiAccountDeliveryChannelStatisticsResponse{}
	_body, _err := client.GetMultiAccountDeliveryChannelStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - GetMultiAccountResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiAccountResourceCenterServiceStatus"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatus() (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.GetMultiAccountResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiAccountResourceConfigurationWithOptions(request *GetMultiAccountResourceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfiguration(request *GetMultiAccountResourceConfigurationRequest) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.GetMultiAccountResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiAccountResourceCountsWithOptions(request *GetMultiAccountResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetMultiAccountResourceCountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMultiAccountResourceCountsResponse
func (client *Client) GetMultiAccountResourceCounts(request *GetMultiAccountResourceCountsRequest) (_result *GetMultiAccountResourceCountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiAccountResourceCountsResponse{}
	_body, _err := client.GetMultiAccountResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @param request - GetResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetResourceCenterServiceStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceCenterServiceStatus"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatus() (_result *GetResourceCenterServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.GetResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceConfigurationWithOptions(request *GetResourceConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfiguration(request *GetResourceConfigurationRequest) (_result *GetResourceConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.GetResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceCountsWithOptions(request *GetResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceCountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetResourceCountsResponse
func (client *Client) GetResourceCounts(request *GetResourceCountsRequest) (_result *GetResourceCountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.GetResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSavedQueryWithOptions(request *GetSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *GetSavedQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSavedQueryResponse
func (client *Client) GetSavedQuery(request *GetSavedQueryRequest) (_result *GetSavedQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.GetSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDeliveryChannelsWithOptions(request *ListDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryChannelsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDeliveryChannelsResponse
func (client *Client) ListDeliveryChannels(request *ListDeliveryChannelsRequest) (_result *ListDeliveryChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeliveryChannelsResponse{}
	_body, _err := client.ListDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListExampleQueriesWithOptions(request *ListExampleQueriesRequest, runtime *dara.RuntimeOptions) (_result *ListExampleQueriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueries(request *ListExampleQueriesRequest) (_result *ListExampleQueriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.ListExampleQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @param request - ListFiltersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFiltersResponse
func (client *Client) ListFiltersWithOptions(runtime *dara.RuntimeOptions) (_result *ListFiltersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListFilters"),
		Version:     dara.String("2022-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @return ListFiltersResponse
func (client *Client) ListFilters() (_result *ListFiltersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFiltersResponse{}
	_body, _err := client.ListFiltersWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiAccountDeliveryChannelsWithOptions(request *ListMultiAccountDeliveryChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMultiAccountDeliveryChannelsResponse
func (client *Client) ListMultiAccountDeliveryChannels(request *ListMultiAccountDeliveryChannelsRequest) (_result *ListMultiAccountDeliveryChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiAccountDeliveryChannelsResponse{}
	_body, _err := client.ListMultiAccountDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiAccountResourceGroupsWithOptions(request *ListMultiAccountResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroups(request *ListMultiAccountResourceGroupsRequest) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.ListMultiAccountResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiAccountResourceRelationshipsWithOptions(request *ListMultiAccountResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMultiAccountResourceRelationshipsResponse
func (client *Client) ListMultiAccountResourceRelationships(request *ListMultiAccountResourceRelationshipsRequest) (_result *ListMultiAccountResourceRelationshipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiAccountResourceRelationshipsResponse{}
	_body, _err := client.ListMultiAccountResourceRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiAccountTagKeysWithOptions(request *ListMultiAccountTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeys(request *ListMultiAccountTagKeysRequest) (_result *ListMultiAccountTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.ListMultiAccountTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiAccountTagValuesWithOptions(request *ListMultiAccountTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValues(request *ListMultiAccountTagValuesRequest) (_result *ListMultiAccountTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.ListMultiAccountTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceRelationshipsWithOptions(request *ListResourceRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceRelationshipsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResourceRelationshipsResponse
func (client *Client) ListResourceRelationships(request *ListResourceRelationshipsRequest) (_result *ListResourceRelationshipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.ListResourceRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSavedQueriesWithOptions(request *ListSavedQueriesRequest, runtime *dara.RuntimeOptions) (_result *ListSavedQueriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueries(request *ListSavedQueriesRequest) (_result *ListSavedQueriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.ListSavedQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchMultiAccountResourcesWithOptions(request *SearchMultiAccountResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResources(request *SearchMultiAccountResourcesRequest) (_result *SearchMultiAccountResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.SearchMultiAccountResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchResourcesWithOptions(request *SearchResourcesRequest, runtime *dara.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchResourcesResponse
func (client *Client) SearchResources(request *SearchResourcesRequest) (_result *SearchResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchResourcesResponse{}
	_body, _err := client.SearchResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateDeliveryChannelWithOptions(request *UpdateDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateDeliveryChannelResponse
func (client *Client) UpdateDeliveryChannel(request *UpdateDeliveryChannelRequest) (_result *UpdateDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDeliveryChannelResponse{}
	_body, _err := client.UpdateDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFilterWithOptions(request *UpdateFilterRequest, runtime *dara.RuntimeOptions) (_result *UpdateFilterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFilterResponse
func (client *Client) UpdateFilter(request *UpdateFilterRequest) (_result *UpdateFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFilterResponse{}
	_body, _err := client.UpdateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMultiAccountDeliveryChannelWithOptions(request *UpdateMultiAccountDeliveryChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateMultiAccountDeliveryChannelResponse
func (client *Client) UpdateMultiAccountDeliveryChannel(request *UpdateMultiAccountDeliveryChannelRequest) (_result *UpdateMultiAccountDeliveryChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMultiAccountDeliveryChannelResponse{}
	_body, _err := client.UpdateMultiAccountDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateSavedQueryWithOptions(request *UpdateSavedQueryRequest, runtime *dara.RuntimeOptions) (_result *UpdateSavedQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQuery(request *UpdateSavedQueryRequest) (_result *UpdateSavedQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.UpdateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
