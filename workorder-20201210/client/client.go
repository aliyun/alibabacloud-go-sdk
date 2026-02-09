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
	client.EndpointRule = dara.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": dara.String("workorder.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1": dara.String("workorder.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("workorder"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - GetTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketTemplateResponse
func (client *Client) GetTicketTemplateWithOptions(request *GetTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTicketTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicketTemplate"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTicketTemplateRequest
//
// @return GetTicketTemplateResponse
func (client *Client) GetTicketTemplate(request *GetTicketTemplateRequest) (_result *GetTicketTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTicketTemplateResponse{}
	_body, _err := client.GetTicketTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 是否首次访问BBS工单
//
// @param request - IsFirstBbsTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsFirstBbsTicketResponse
func (client *Client) IsFirstBbsTicketWithOptions(request *IsFirstBbsTicketRequest, runtime *dara.RuntimeOptions) (_result *IsFirstBbsTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsFirstBbsTicket"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsFirstBbsTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 是否首次访问BBS工单
//
// @param request - IsFirstBbsTicketRequest
//
// @return IsFirstBbsTicketResponse
func (client *Client) IsFirstBbsTicket(request *IsFirstBbsTicketRequest) (_result *IsFirstBbsTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IsFirstBbsTicketResponse{}
	_body, _err := client.IsFirstBbsTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SuggestCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuggestCategoryResponse
func (client *Client) SuggestCategoryWithOptions(request *SuggestCategoryRequest, runtime *dara.RuntimeOptions) (_result *SuggestCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cna) {
		query["Cna"] = request.Cna
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DistributionChannel) {
		query["DistributionChannel"] = request.DistributionChannel
	}

	if !dara.IsNil(request.Referrer) {
		query["Referrer"] = request.Referrer
	}

	if !dara.IsNil(request.SubDistributionChannel) {
		query["SubDistributionChannel"] = request.SubDistributionChannel
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.UseHotSuggestCatchAll) {
		query["UseHotSuggestCatchAll"] = request.UseHotSuggestCatchAll
	}

	if !dara.IsNil(request.XGatewayExtendInfo) {
		query["XGatewayExtendInfo"] = request.XGatewayExtendInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuggestCategory"),
		Version:     dara.String("2020-12-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuggestCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuggestCategoryRequest
//
// @return SuggestCategoryResponse
func (client *Client) SuggestCategory(request *SuggestCategoryRequest) (_result *SuggestCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuggestCategoryResponse{}
	_body, _err := client.SuggestCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
