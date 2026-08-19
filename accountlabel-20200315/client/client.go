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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("accountlabel"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomerLabelResponse
func (client *Client) AddCustomerLabelWithOptions(request *AddCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *AddCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endtime) {
		query["Endtime"] = request.Endtime
	}

	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddCustomerLabelRequest
//
// @return AddCustomerLabelResponse
func (client *Client) AddCustomerLabel(request *AddCustomerLabelRequest) (_result *AddCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCustomerLabelResponse{}
	_body, _err := client.AddCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 传入多个labelSeries查询标签
//
// @param tmpReq - BatchFetchAccountLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchFetchAccountLabelResponse
func (client *Client) BatchFetchAccountLabelWithOptions(tmpReq *BatchFetchAccountLabelRequest, runtime *dara.RuntimeOptions) (_result *BatchFetchAccountLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchFetchAccountLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LabelSeriesList) {
		request.LabelSeriesListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSeriesList, dara.String("LabelSeriesList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Instant) {
		query["Instant"] = request.Instant
	}

	if !dara.IsNil(request.LabelSeriesListShrink) {
		query["LabelSeriesList"] = request.LabelSeriesListShrink
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchFetchAccountLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchFetchAccountLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 传入多个labelSeries查询标签
//
// @param request - BatchFetchAccountLabelRequest
//
// @return BatchFetchAccountLabelResponse
func (client *Client) BatchFetchAccountLabel(request *BatchFetchAccountLabelRequest) (_result *BatchFetchAccountLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchFetchAccountLabelResponse{}
	_body, _err := client.BatchFetchAccountLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomerLabelResponse
func (client *Client) DeleteCustomerLabelWithOptions(request *DeleteCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomerLabelRequest
//
// @return DeleteCustomerLabelResponse
func (client *Client) DeleteCustomerLabel(request *DeleteCustomerLabelRequest) (_result *DeleteCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomerLabelResponse{}
	_body, _err := client.DeleteCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerLabelResponse
func (client *Client) QueryCustomerLabelWithOptions(request *QueryCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Instant) {
		query["Instant"] = request.Instant
	}

	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerLabel"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomerLabelRequest
//
// @return QueryCustomerLabelResponse
func (client *Client) QueryCustomerLabel(request *QueryCustomerLabelRequest) (_result *QueryCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCustomerLabelResponse{}
	_body, _err := client.QueryCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCustomerLabelByConfigGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerLabelByConfigGroupResponse
func (client *Client) QueryCustomerLabelByConfigGroupWithOptions(request *QueryCustomerLabelByConfigGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerLabelByConfigGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerLabelByConfigGroup"),
		Version:     dara.String("2020-03-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerLabelByConfigGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomerLabelByConfigGroupRequest
//
// @return QueryCustomerLabelByConfigGroupResponse
func (client *Client) QueryCustomerLabelByConfigGroup(request *QueryCustomerLabelByConfigGroupRequest) (_result *QueryCustomerLabelByConfigGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCustomerLabelByConfigGroupResponse{}
	_body, _err := client.QueryCustomerLabelByConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
