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
	client.Endpoint, _err = client.GetEndpoint(dara.String("yunjian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CreateDemandPlanRequest
//
// @param headers - CreateDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDemandPlanResponse
func (client *Client) CreateDemandPlanWithOptions(request *CreateDemandPlanRequest, headers *CreateDemandPlanHeaders, runtime *dara.RuntimeOptions) (_result *CreateDemandPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.DemandType) {
		body["demandType"] = request.DemandType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.TargetCid) {
		body["targetCid"] = request.TargetCid
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDemandPlan"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/saveUrgentDemandPlanItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDemandPlanRequest
//
// @return CreateDemandPlanResponse
func (client *Client) CreateDemandPlan(request *CreateDemandPlanRequest) (_result *CreateDemandPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDemandPlanHeaders{}
	_result = &CreateDemandPlanResponse{}
	_body, _err := client.CreateDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建plan2.0版本
//
// @param request - CreateDemandPlanV2Request
//
// @param headers - CreateDemandPlanV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDemandPlanV2Response
func (client *Client) CreateDemandPlanV2WithOptions(request *CreateDemandPlanV2Request, headers *CreateDemandPlanV2Headers, runtime *dara.RuntimeOptions) (_result *CreateDemandPlanV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.DemandType) {
		body["demandType"] = request.DemandType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	if !dara.IsNil(request.TargetCid) {
		body["targetCid"] = request.TargetCid
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDemandPlanV2"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/saveDemandPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDemandPlanV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建plan2.0版本
//
// @param request - CreateDemandPlanV2Request
//
// @return CreateDemandPlanV2Response
func (client *Client) CreateDemandPlanV2(request *CreateDemandPlanV2Request) (_result *CreateDemandPlanV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDemandPlanV2Headers{}
	_result = &CreateDemandPlanV2Response{}
	_body, _err := client.CreateDemandPlanV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单ite 删除
//
// @param request - DeleteUrgentDemandItemRequest
//
// @param headers - DeleteUrgentDemandItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUrgentDemandItemResponse
func (client *Client) DeleteUrgentDemandItemWithOptions(request *DeleteUrgentDemandItemRequest, headers *DeleteUrgentDemandItemHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUrgentDemandItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Modifier) {
		query["modifier"] = request.Modifier
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUrgentDemandItem"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/deleteUrgentDemandItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUrgentDemandItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单ite 删除
//
// @param request - DeleteUrgentDemandItemRequest
//
// @return DeleteUrgentDemandItemResponse
func (client *Client) DeleteUrgentDemandItem(request *DeleteUrgentDemandItemRequest) (_result *DeleteUrgentDemandItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteUrgentDemandItemHeaders{}
	_result = &DeleteUrgentDemandItemResponse{}
	_body, _err := client.DeleteUrgentDemandItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单plan删除
//
// @param request - DeleteUrgentDemandPlanRequest
//
// @param headers - DeleteUrgentDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUrgentDemandPlanResponse
func (client *Client) DeleteUrgentDemandPlanWithOptions(request *DeleteUrgentDemandPlanRequest, headers *DeleteUrgentDemandPlanHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUrgentDemandPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Modifier) {
		query["modifier"] = request.Modifier
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUrgentDemandPlan"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/deleteUrgentDemandPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUrgentDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单plan删除
//
// @param request - DeleteUrgentDemandPlanRequest
//
// @return DeleteUrgentDemandPlanResponse
func (client *Client) DeleteUrgentDemandPlan(request *DeleteUrgentDemandPlanRequest) (_result *DeleteUrgentDemandPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteUrgentDemandPlanHeaders{}
	_result = &DeleteUrgentDemandPlanResponse{}
	_body, _err := client.DeleteUrgentDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 交付信息同步
//
// @param request - DeliveryItemDetailSynRequest
//
// @param headers - DeliveryItemDetailSynHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeliveryItemDetailSynResponse
func (client *Client) DeliveryItemDetailSynWithOptions(request *DeliveryItemDetailSynRequest, headers *DeliveryItemDetailSynHeaders, runtime *dara.RuntimeOptions) (_result *DeliveryItemDetailSynResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.DeliveryItemDetailDTOS) {
		body["deliveryItemDetailDTOS"] = request.DeliveryItemDetailDTOS
	}

	if !dara.IsNil(request.DeliveryItemId) {
		body["deliveryItemId"] = request.DeliveryItemId
	}

	if !dara.IsNil(request.DeliveryPlanId) {
		body["deliveryPlanId"] = request.DeliveryPlanId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeliveryItemDetailSyn"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/supply/deliveryItemDataSync"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeliveryItemDetailSynResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 交付信息同步
//
// @param request - DeliveryItemDetailSynRequest
//
// @return DeliveryItemDetailSynResponse
func (client *Client) DeliveryItemDetailSyn(request *DeliveryItemDetailSynRequest) (_result *DeliveryItemDetailSynResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeliveryItemDetailSynHeaders{}
	_result = &DeliveryItemDetailSynResponse{}
	_body, _err := client.DeliveryItemDetailSynWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询报备单中报备项列表
//
// @param request - GetUrgentDemandItemListRequest
//
// @param headers - GetUrgentDemandItemListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandItemListResponse
func (client *Client) GetUrgentDemandItemListWithOptions(request *GetUrgentDemandItemListRequest, headers *GetUrgentDemandItemListHeaders, runtime *dara.RuntimeOptions) (_result *GetUrgentDemandItemListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["commodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CommodityTypeCode) {
		body["commodityTypeCode"] = request.CommodityTypeCode
	}

	if !dara.IsNil(request.Current) {
		body["current"] = request.Current
	}

	if !dara.IsNil(request.PlanId) {
		body["planId"] = request.PlanId
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.Size) {
		body["size"] = request.Size
	}

	if !dara.IsNil(request.Zone) {
		body["zone"] = request.Zone
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUrgentDemandItemList"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/getUrgentDemandItemList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUrgentDemandItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报备单中报备项列表
//
// @param request - GetUrgentDemandItemListRequest
//
// @return GetUrgentDemandItemListResponse
func (client *Client) GetUrgentDemandItemList(request *GetUrgentDemandItemListRequest) (_result *GetUrgentDemandItemListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUrgentDemandItemListHeaders{}
	_result = &GetUrgentDemandItemListResponse{}
	_body, _err := client.GetUrgentDemandItemListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// getUrgentDemandPlanDetail
//
// @param request - GetUrgentDemandPlanDetailRequest
//
// @param headers - GetUrgentDemandPlanDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandPlanDetailResponse
func (client *Client) GetUrgentDemandPlanDetailWithOptions(request *GetUrgentDemandPlanDetailRequest, headers *GetUrgentDemandPlanDetailHeaders, runtime *dara.RuntimeOptions) (_result *GetUrgentDemandPlanDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		body["planId"] = request.PlanId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUrgentDemandPlanDetail"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/getUrgentDemandPlanDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUrgentDemandPlanDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// getUrgentDemandPlanDetail
//
// @param request - GetUrgentDemandPlanDetailRequest
//
// @return GetUrgentDemandPlanDetailResponse
func (client *Client) GetUrgentDemandPlanDetail(request *GetUrgentDemandPlanDetailRequest) (_result *GetUrgentDemandPlanDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUrgentDemandPlanDetailHeaders{}
	_result = &GetUrgentDemandPlanDetailResponse{}
	_body, _err := client.GetUrgentDemandPlanDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询报备单列表
//
// @param request - GetUrgentDemandPlanListRequest
//
// @param headers - GetUrgentDemandPlanListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandPlanListResponse
func (client *Client) GetUrgentDemandPlanListWithOptions(request *GetUrgentDemandPlanListRequest, headers *GetUrgentDemandPlanListHeaders, runtime *dara.RuntimeOptions) (_result *GetUrgentDemandPlanListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["current"] = request.Current
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PlanType) {
		body["planType"] = request.PlanType
	}

	if !dara.IsNil(request.Size) {
		body["size"] = request.Size
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUrgentDemandPlanList"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/getUrgentDemandPlanList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUrgentDemandPlanListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报备单列表
//
// @param request - GetUrgentDemandPlanListRequest
//
// @return GetUrgentDemandPlanListResponse
func (client *Client) GetUrgentDemandPlanList(request *GetUrgentDemandPlanListRequest) (_result *GetUrgentDemandPlanListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUrgentDemandPlanListHeaders{}
	_result = &GetUrgentDemandPlanListResponse{}
	_body, _err := client.GetUrgentDemandPlanListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 报备列表查询
//
// @param request - PageDemandPlanWithDemandInfoRequest
//
// @param headers - PageDemandPlanWithDemandInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageDemandPlanWithDemandInfoResponse
func (client *Client) PageDemandPlanWithDemandInfoWithOptions(request *PageDemandPlanWithDemandInfoRequest, headers *PageDemandPlanWithDemandInfoHeaders, runtime *dara.RuntimeOptions) (_result *PageDemandPlanWithDemandInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalStatus) {
		body["approvalStatus"] = request.ApprovalStatus
	}

	if !dara.IsNil(request.CreateTimeEnd) {
		body["createTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		body["createTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.CreatorStaffId) {
		body["creatorStaffId"] = request.CreatorStaffId
	}

	if !dara.IsNil(request.DemandDeliveryStatus) {
		body["demandDeliveryStatus"] = request.DemandDeliveryStatus
	}

	if !dara.IsNil(request.DemandId) {
		body["demandId"] = request.DemandId
	}

	if !dara.IsNil(request.DemandPlanId) {
		body["demandPlanId"] = request.DemandPlanId
	}

	if !dara.IsNil(request.DemandPlanStatus) {
		body["demandPlanStatus"] = request.DemandPlanStatus
	}

	if !dara.IsNil(request.Operator) {
		body["operator"] = request.Operator
	}

	if !dara.IsNil(request.PageNum) {
		body["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoCode) {
		body["roCode"] = request.RoCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PageDemandPlanWithDemandInfo"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/getDemandPlanList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PageDemandPlanWithDemandInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 报备列表查询
//
// @param request - PageDemandPlanWithDemandInfoRequest
//
// @return PageDemandPlanWithDemandInfoResponse
func (client *Client) PageDemandPlanWithDemandInfo(request *PageDemandPlanWithDemandInfoRequest) (_result *PageDemandPlanWithDemandInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PageDemandPlanWithDemandInfoHeaders{}
	_result = &PageDemandPlanWithDemandInfoResponse{}
	_body, _err := client.PageDemandPlanWithDemandInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ecs资源方案
//
// @param request - PushResourcePlanRequest
//
// @param headers - PushResourcePlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourcePlanResponse
func (client *Client) PushResourcePlanWithOptions(request *PushResourcePlanRequest, headers *PushResourcePlanHeaders, runtime *dara.RuntimeOptions) (_result *PushResourcePlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BufferCnt) {
		body["bufferCnt"] = request.BufferCnt
	}

	if !dara.IsNil(request.DemandCount) {
		body["demandCount"] = request.DemandCount
	}

	if !dara.IsNil(request.DemandId) {
		body["demandId"] = request.DemandId
	}

	if !dara.IsNil(request.MethodList) {
		body["methodList"] = request.MethodList
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.RequireCnt) {
		body["requireCnt"] = request.RequireCnt
	}

	if !dara.IsNil(request.SubDemandId) {
		body["subDemandId"] = request.SubDemandId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushResourcePlan"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/supply/resourcePlan/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResourcePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ecs资源方案
//
// @param request - PushResourcePlanRequest
//
// @return PushResourcePlanResponse
func (client *Client) PushResourcePlan(request *PushResourcePlanRequest) (_result *PushResourcePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushResourcePlanHeaders{}
	_result = &PushResourcePlanResponse{}
	_body, _err := client.PushResourcePlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询accountId下所有存在交付状态（包括部分）的报备数据, 以及开通数据情况
//
// @param request - QueryDeliveredSupplyItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeliveredSupplyItemsResponse
func (client *Client) QueryDeliveredSupplyItemsWithOptions(request *QueryDeliveredSupplyItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryDeliveredSupplyItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.CommodityTypeCode) {
		query["commodityTypeCode"] = request.CommodityTypeCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeliveredSupplyItems"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/supply/queryDeliveredSupplyItems"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &QueryDeliveredSupplyItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询accountId下所有存在交付状态（包括部分）的报备数据, 以及开通数据情况
//
// @param request - QueryDeliveredSupplyItemsRequest
//
// @return QueryDeliveredSupplyItemsResponse
func (client *Client) QueryDeliveredSupplyItems(request *QueryDeliveredSupplyItemsRequest) (_result *QueryDeliveredSupplyItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDeliveredSupplyItemsResponse{}
	_body, _err := client.QueryDeliveredSupplyItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账单预算数据
//
// @param request - QueryPeriodBudgetBillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPeriodBudgetBillResponse
func (client *Client) QueryPeriodBudgetBillWithOptions(request *QueryPeriodBudgetBillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryPeriodBudgetBillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectIds) {
		query["objectIds"] = request.ObjectIds
	}

	if !dara.IsNil(request.ObjectType) {
		query["objectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Period) {
		query["period"] = request.Period
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPeriodBudgetBill"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/annual/budget/queryPeriodBudgetBill"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPeriodBudgetBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账单预算数据
//
// @param request - QueryPeriodBudgetBillRequest
//
// @return QueryPeriodBudgetBillResponse
func (client *Client) QueryPeriodBudgetBill(request *QueryPeriodBudgetBillRequest) (_result *QueryPeriodBudgetBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPeriodBudgetBillResponse{}
	_body, _err := client.QueryPeriodBudgetBillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单item新增
//
// @param request - SaveUrgentDemandItemRequest
//
// @param headers - SaveUrgentDemandItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveUrgentDemandItemResponse
func (client *Client) SaveUrgentDemandItemWithOptions(request *SaveUrgentDemandItemRequest, headers *SaveUrgentDemandItemHeaders, runtime *dara.RuntimeOptions) (_result *SaveUrgentDemandItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.Creator) {
		body["creator"] = request.Creator
	}

	if !dara.IsNil(request.EffectTime) {
		body["effectTime"] = request.EffectTime
	}

	if !dara.IsNil(request.Modifier) {
		body["modifier"] = request.Modifier
	}

	if !dara.IsNil(request.NetworkType) {
		body["networkType"] = request.NetworkType
	}

	if !dara.IsNil(request.PayDuration) {
		body["payDuration"] = request.PayDuration
	}

	if !dara.IsNil(request.PayDurationUnit) {
		body["payDurationUnit"] = request.PayDurationUnit
	}

	if !dara.IsNil(request.PayType) {
		body["payType"] = request.PayType
	}

	if !dara.IsNil(request.PlanId) {
		body["planId"] = request.PlanId
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.UrgentDemandEbsRequest) {
		body["urgentDemandEbsRequest"] = request.UrgentDemandEbsRequest
	}

	if !dara.IsNil(request.UrgentDemandEcsRequest) {
		body["urgentDemandEcsRequest"] = request.UrgentDemandEcsRequest
	}

	if !dara.IsNil(request.Zone) {
		body["zone"] = request.Zone
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveUrgentDemandItem"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/saveUrgentDemandItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveUrgentDemandItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单item新增
//
// @param request - SaveUrgentDemandItemRequest
//
// @return SaveUrgentDemandItemResponse
func (client *Client) SaveUrgentDemandItem(request *SaveUrgentDemandItemRequest) (_result *SaveUrgentDemandItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SaveUrgentDemandItemHeaders{}
	_result = &SaveUrgentDemandItemResponse{}
	_body, _err := client.SaveUrgentDemandItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// submitUrgentDemandPlan
//
// @param request - SubmitUrgentDemandPlanRequest
//
// @param headers - SubmitUrgentDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitUrgentDemandPlanResponse
func (client *Client) SubmitUrgentDemandPlanWithOptions(request *SubmitUrgentDemandPlanRequest, headers *SubmitUrgentDemandPlanHeaders, runtime *dara.RuntimeOptions) (_result *SubmitUrgentDemandPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PlanId) {
		query["planId"] = request.PlanId
	}

	if !dara.IsNil(request.UserId) {
		query["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.YunUserId) {
		realHeaders["Yun-User-Id"] = dara.String(dara.ToString(dara.StringValue(headers.YunUserId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitUrgentDemandPlan"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/demand/urgent/submitUrgentDemandPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitUrgentDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// submitUrgentDemandPlan
//
// @param request - SubmitUrgentDemandPlanRequest
//
// @return SubmitUrgentDemandPlanResponse
func (client *Client) SubmitUrgentDemandPlan(request *SubmitUrgentDemandPlanRequest) (_result *SubmitUrgentDemandPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SubmitUrgentDemandPlanHeaders{}
	_result = &SubmitUrgentDemandPlanResponse{}
	_body, _err := client.SubmitUrgentDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云产品交付决策反馈回调
//
// @param request - AcceptFulfillmentDecisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptFulfillmentDecisionResponse
func (client *Client) AcceptFulfillmentDecisionWithOptions(request *AcceptFulfillmentDecisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AcceptFulfillmentDecisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DecisionConclusion) {
		body["DecisionConclusion"] = request.DecisionConclusion
	}

	if !dara.IsNil(request.DecisionType) {
		body["DecisionType"] = request.DecisionType
	}

	if !dara.IsNil(request.OrderId) {
		body["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("acceptFulfillmentDecision"),
		Version:     dara.String("2021-12-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/acceptFulfillmentDecision"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptFulfillmentDecisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云产品交付决策反馈回调
//
// @param request - AcceptFulfillmentDecisionRequest
//
// @return AcceptFulfillmentDecisionResponse
func (client *Client) AcceptFulfillmentDecision(request *AcceptFulfillmentDecisionRequest) (_result *AcceptFulfillmentDecisionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AcceptFulfillmentDecisionResponse{}
	_body, _err := client.AcceptFulfillmentDecisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
