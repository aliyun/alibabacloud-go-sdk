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
	client.Endpoint, _err = client.GetEndpoint(dara.String("airticketopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 账号资金流水
//
// @param request - AccountFlowListRequest
//
// @param headers - AccountFlowListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountFlowListResponse
func (client *Client) AccountFlowListWithOptions(request *AccountFlowListRequest, headers *AccountFlowListHeaders, runtime *dara.RuntimeOptions) (_result *AccountFlowListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DayNum) {
		query["day_num"] = request.DayNum
	}

	if !dara.IsNil(request.PageIndex) {
		query["page_index"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.UtcBeginTime) {
		query["utc_begin_time"] = request.UtcBeginTime
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountFlowList"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/account/flow-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountFlowListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号资金流水
//
// @param request - AccountFlowListRequest
//
// @return AccountFlowListResponse
func (client *Client) AccountFlowList(request *AccountFlowListRequest) (_result *AccountFlowListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AccountFlowListHeaders{}
	_result = &AccountFlowListResponse{}
	_body, _err := client.AccountFlowListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Ancillary - Suggestion
//
// Description:
//
// search ancillary for selected solution, you should enter the solution_id returned by enrich.
//
// @param request - AncillarySuggestRequest
//
// @param headers - AncillarySuggestHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AncillarySuggestResponse
func (client *Client) AncillarySuggestWithOptions(request *AncillarySuggestRequest, headers *AncillarySuggestHeaders, runtime *dara.RuntimeOptions) (_result *AncillarySuggestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SolutionId) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AncillarySuggest"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/ancillary/action-suggest"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AncillarySuggestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Ancillary - Suggestion
//
// Description:
//
// search ancillary for selected solution, you should enter the solution_id returned by enrich.
//
// @param request - AncillarySuggestRequest
//
// @return AncillarySuggestResponse
func (client *Client) AncillarySuggest(request *AncillarySuggestRequest) (_result *AncillarySuggestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AncillarySuggestHeaders{}
	_result = &AncillarySuggestResponse{}
	_body, _err := client.AncillarySuggestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Transaction-Reservation
//
// Description:
//
// Enter solution_id returned by enrich, ancillary_id returned by ancillarySuggest(optional), passengers information and contact information, the book interface will create an order wait for pay.
//
// There are two issues should be noticed:
//
// 1. the solution_id must be processed by pricing.
//
// 2. the order created by book interface should be pay within 30 minutes, otherwise the order will be closed.
//
// @param tmpReq - BookRequest
//
// @param headers - BookHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BookResponse
func (client *Client) BookWithOptions(tmpReq *BookRequest, headers *BookHeaders, runtime *dara.RuntimeOptions) (_result *BookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("contact"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerAncillaryPurchaseMapList) {
		request.PassengerAncillaryPurchaseMapListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerAncillaryPurchaseMapList, dara.String("passenger_ancillary_purchase_map_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PassengerList) {
		request.PassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerList, dara.String("passenger_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactShrink) {
		body["contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.OutOrderNum) {
		body["out_order_num"] = request.OutOrderNum
	}

	if !dara.IsNil(request.PassengerAncillaryPurchaseMapListShrink) {
		body["passenger_ancillary_purchase_map_list"] = request.PassengerAncillaryPurchaseMapListShrink
	}

	if !dara.IsNil(request.PassengerListShrink) {
		body["passenger_list"] = request.PassengerListShrink
	}

	if !dara.IsNil(request.SolutionId) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Book"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-book"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Transaction-Reservation
//
// Description:
//
// Enter solution_id returned by enrich, ancillary_id returned by ancillarySuggest(optional), passengers information and contact information, the book interface will create an order wait for pay.
//
// There are two issues should be noticed:
//
// 1. the solution_id must be processed by pricing.
//
// 2. the order created by book interface should be pay within 30 minutes, otherwise the order will be closed.
//
// @param request - BookRequest
//
// @return BookResponse
func (client *Client) Book(request *BookRequest) (_result *BookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BookHeaders{}
	_result = &BookResponse{}
	_body, _err := client.BookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Transaction - Unpaid Cancellation
//
// Description:
//
// close an unpaid order
//
// @param request - CancelRequest
//
// @param headers - CancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelResponse
func (client *Client) CancelWithOptions(request *CancelRequest, headers *CancelHeaders, runtime *dara.RuntimeOptions) (_result *CancelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Cancel"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Transaction - Unpaid Cancellation
//
// Description:
//
// close an unpaid order
//
// @param request - CancelRequest
//
// @return CancelResponse
func (client *Client) Cancel(request *CancelRequest) (_result *CancelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CancelHeaders{}
	_result = &CancelResponse{}
	_body, _err := client.CancelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-Apply
//
// @param tmpReq - ChangeApplyRequest
//
// @param headers - ChangeApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeApplyResponse
func (client *Client) ChangeApplyWithOptions(tmpReq *ChangeApplyRequest, headers *ChangeApplyHeaders, runtime *dara.RuntimeOptions) (_result *ChangeApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChangeApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChangePassengerList) {
		request.ChangePassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangePassengerList, dara.String("change_passenger_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChangedJourneys) {
		request.ChangedJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangedJourneys, dara.String("changed_journeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("contact"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangePassengerListShrink) {
		body["change_passenger_list"] = request.ChangePassengerListShrink
	}

	if !dara.IsNil(request.ChangedJourneysShrink) {
		body["changed_journeys"] = request.ChangedJourneysShrink
	}

	if !dara.IsNil(request.ContactShrink) {
		body["contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeApply"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/action-apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-Apply
//
// @param request - ChangeApplyRequest
//
// @return ChangeApplyResponse
func (client *Client) ChangeApply(request *ChangeApplyRequest) (_result *ChangeApplyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeApplyHeaders{}
	_result = &ChangeApplyResponse{}
	_body, _err := client.ChangeApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-取消
//
// @param request - ChangeCancelRequest
//
// @param headers - ChangeCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancelWithOptions(request *ChangeCancelRequest, headers *ChangeCancelHeaders, runtime *dara.RuntimeOptions) (_result *ChangeCancelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderNum) {
		body["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeCancel"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/action-cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-取消
//
// @param request - ChangeCancelRequest
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancel(request *ChangeCancelRequest) (_result *ChangeCancelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeCancelHeaders{}
	_result = &ChangeCancelResponse{}
	_body, _err := client.ChangeCancelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-确认
//
// @param request - ChangeConfirmRequest
//
// @param headers - ChangeConfirmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeConfirmResponse
func (client *Client) ChangeConfirmWithOptions(request *ChangeConfirmRequest, headers *ChangeConfirmHeaders, runtime *dara.RuntimeOptions) (_result *ChangeConfirmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderNum) {
		body["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeConfirm"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/action-confirm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeConfirmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-确认
//
// @param request - ChangeConfirmRequest
//
// @return ChangeConfirmResponse
func (client *Client) ChangeConfirm(request *ChangeConfirmRequest) (_result *ChangeConfirmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeConfirmHeaders{}
	_result = &ChangeConfirmResponse{}
	_body, _err := client.ChangeConfirmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Change-Detail
//
// @param request - ChangeDetailRequest
//
// @param headers - ChangeDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailResponse
func (client *Client) ChangeDetailWithOptions(request *ChangeDetailRequest, headers *ChangeDetailHeaders, runtime *dara.RuntimeOptions) (_result *ChangeDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderNum) {
		query["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDetail"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Change-Detail
//
// @param request - ChangeDetailRequest
//
// @return ChangeDetailResponse
func (client *Client) ChangeDetail(request *ChangeDetailRequest) (_result *ChangeDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeDetailHeaders{}
	_result = &ChangeDetailResponse{}
	_body, _err := client.ChangeDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签单列表-关于买家账号
//
// @param request - ChangeDetailListOfBuyerRequest
//
// @param headers - ChangeDetailListOfBuyerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailListOfBuyerResponse
func (client *Client) ChangeDetailListOfBuyerWithOptions(request *ChangeDetailListOfBuyerRequest, headers *ChangeDetailListOfBuyerHeaders, runtime *dara.RuntimeOptions) (_result *ChangeDetailListOfBuyerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageIndex) {
		query["page_index"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.UtcCreateBegin) {
		query["utc_create_begin"] = request.UtcCreateBegin
	}

	if !dara.IsNil(request.UtcCreateEnd) {
		query["utc_create_end"] = request.UtcCreateEnd
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDetailListOfBuyer"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/buyer/detail-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDetailListOfBuyerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签单列表-关于买家账号
//
// @param request - ChangeDetailListOfBuyerRequest
//
// @return ChangeDetailListOfBuyerResponse
func (client *Client) ChangeDetailListOfBuyer(request *ChangeDetailListOfBuyerRequest) (_result *ChangeDetailListOfBuyerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeDetailListOfBuyerHeaders{}
	_result = &ChangeDetailListOfBuyerResponse{}
	_body, _err := client.ChangeDetailListOfBuyerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签单列表-关于正向订单
//
// @param request - ChangeDetailListOfOrderNumRequest
//
// @param headers - ChangeDetailListOfOrderNumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailListOfOrderNumResponse
func (client *Client) ChangeDetailListOfOrderNumWithOptions(request *ChangeDetailListOfOrderNumRequest, headers *ChangeDetailListOfOrderNumHeaders, runtime *dara.RuntimeOptions) (_result *ChangeDetailListOfOrderNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		query["order_num"] = request.OrderNum
	}

	if !dara.IsNil(request.PageIndex) {
		query["page_index"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDetailListOfOrderNum"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/change/order-num/detail-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDetailListOfOrderNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签单列表-关于正向订单
//
// @param request - ChangeDetailListOfOrderNumRequest
//
// @return ChangeDetailListOfOrderNumResponse
func (client *Client) ChangeDetailListOfOrderNum(request *ChangeDetailListOfOrderNumRequest) (_result *ChangeDetailListOfOrderNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeDetailListOfOrderNumHeaders{}
	_result = &ChangeDetailListOfOrderNumResponse{}
	_body, _err := client.ChangeDetailListOfOrderNumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据收集-低价航班信息
//
// @param tmpReq - CollectFlightLowestPriceRequest
//
// @param headers - CollectFlightLowestPriceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CollectFlightLowestPriceResponse
func (client *Client) CollectFlightLowestPriceWithOptions(tmpReq *CollectFlightLowestPriceRequest, headers *CollectFlightLowestPriceHeaders, runtime *dara.RuntimeOptions) (_result *CollectFlightLowestPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CollectFlightLowestPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LowestPriceFlightInfoList) {
		request.LowestPriceFlightInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LowestPriceFlightInfoList, dara.String("lowest_price_flight_info_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LowestPriceFlightInfoListShrink) {
		body["lowest_price_flight_info_list"] = request.LowestPriceFlightInfoListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CollectFlightLowestPrice"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/data-collect/flight-lowest-price"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CollectFlightLowestPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据收集-低价航班信息
//
// @param request - CollectFlightLowestPriceRequest
//
// @return CollectFlightLowestPriceResponse
func (client *Client) CollectFlightLowestPrice(request *CollectFlightLowestPriceRequest) (_result *CollectFlightLowestPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CollectFlightLowestPriceHeaders{}
	_result = &CollectFlightLowestPriceResponse{}
	_body, _err := client.CollectFlightLowestPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Search-Enrich
//
// Description:
//
// Choose either `solution_id` or `journey_param_list` in the parameters, and `solution_id` needs to be obtained from the Search interface.
//
// @param tmpReq - EnrichRequest
//
// @param headers - EnrichHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnrichResponse
func (client *Client) EnrichWithOptions(tmpReq *EnrichRequest, headers *EnrichHeaders, runtime *dara.RuntimeOptions) (_result *EnrichResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnrichShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JourneyParamList) {
		request.JourneyParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JourneyParamList, dara.String("journey_param_list"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Adults) {
		body["adults"] = request.Adults
	}

	if !dara.IsNil(request.CabinClass) {
		body["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.Children) {
		body["children"] = request.Children
	}

	if !dara.IsNil(request.Infants) {
		body["infants"] = request.Infants
	}

	if !dara.IsNil(request.JourneyParamListShrink) {
		body["journey_param_list"] = request.JourneyParamListShrink
	}

	if !dara.IsNil(request.SolutionId) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Enrich"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-enrich"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnrichResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Search-Enrich
//
// Description:
//
// Choose either `solution_id` or `journey_param_list` in the parameters, and `solution_id` needs to be obtained from the Search interface.
//
// @param request - EnrichRequest
//
// @return EnrichResponse
func (client *Client) Enrich(request *EnrichRequest) (_result *EnrichResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &EnrichHeaders{}
	_result = &EnrichResponse{}
	_body, _err := client.EnrichWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 附件上传
//
// @param request - FileUploadRequest
//
// @param headers - FileUploadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileUploadResponse
func (client *Client) FileUploadWithOptions(request *FileUploadRequest, headers *FileUploadHeaders, runtime *dara.RuntimeOptions) (_result *FileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileContent) {
		body["file_content"] = request.FileContent
	}

	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileUpload"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/attachment/action-upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 附件上传
//
// @param request - FileUploadRequest
//
// @return FileUploadResponse
func (client *Client) FileUpload(request *FileUploadRequest) (_result *FileUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FileUploadHeaders{}
	_result = &FileUploadResponse{}
	_body, _err := client.FileUploadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航变信息-关于订单
//
// @param request - FlightChangeOfOrderRequest
//
// @param headers - FlightChangeOfOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightChangeOfOrderResponse
func (client *Client) FlightChangeOfOrderWithOptions(request *FlightChangeOfOrderRequest, headers *FlightChangeOfOrderHeaders, runtime *dara.RuntimeOptions) (_result *FlightChangeOfOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		query["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlightChangeOfOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/flightchange/of-order"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FlightChangeOfOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航变信息-关于订单
//
// @param request - FlightChangeOfOrderRequest
//
// @return FlightChangeOfOrderResponse
func (client *Client) FlightChangeOfOrder(request *FlightChangeOfOrderRequest) (_result *FlightChangeOfOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FlightChangeOfOrderHeaders{}
	_result = &FlightChangeOfOrderResponse{}
	_body, _err := client.FlightChangeOfOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["app_key"] = request.AppKey
	}

	if !dara.IsNil(request.AppSecret) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/token"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Token
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航程行李直挂
//
// @param tmpReq - LuggageDirectRequest
//
// @param headers - LuggageDirectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LuggageDirectResponse
func (client *Client) LuggageDirectWithOptions(tmpReq *LuggageDirectRequest, headers *LuggageDirectHeaders, runtime *dara.RuntimeOptions) (_result *LuggageDirectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LuggageDirectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FlightSegmentParamList) {
		request.FlightSegmentParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlightSegmentParamList, dara.String("flight_segment_param_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FlightSegmentParamListShrink) {
		query["flight_segment_param_list"] = request.FlightSegmentParamListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LuggageDirect"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/flight-data/luggage-direct"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LuggageDirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航程行李直挂
//
// @param request - LuggageDirectRequest
//
// @return LuggageDirectResponse
func (client *Client) LuggageDirect(request *LuggageDirectRequest) (_result *LuggageDirectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &LuggageDirectHeaders{}
	_result = &LuggageDirectResponse{}
	_body, _err := client.LuggageDirectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Trade-Order Details
//
// Description:
//
// query order detail
//
// @param request - OrderDetailRequest
//
// @param headers - OrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderDetailResponse
func (client *Client) OrderDetailWithOptions(request *OrderDetailRequest, headers *OrderDetailHeaders, runtime *dara.RuntimeOptions) (_result *OrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		query["order_num"] = request.OrderNum
	}

	if !dara.IsNil(request.OutOrderNum) {
		query["out_order_num"] = request.OutOrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OrderDetail"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/order-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Trade-Order Details
//
// Description:
//
// query order detail
//
// @param request - OrderDetailRequest
//
// @return OrderDetailResponse
func (client *Client) OrderDetail(request *OrderDetailRequest) (_result *OrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &OrderDetailHeaders{}
	_result = &OrderDetailResponse{}
	_body, _err := client.OrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Trade - Order List
//
// Description:
//
// query order list
//
// @param request - OrderListRequest
//
// @param headers - OrderListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderListResponse
func (client *Client) OrderListWithOptions(request *OrderListRequest, headers *OrderListHeaders, runtime *dara.RuntimeOptions) (_result *OrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BookTimeEnd) {
		query["book_time_end"] = request.BookTimeEnd
	}

	if !dara.IsNil(request.BookTimeStart) {
		query["book_time_start"] = request.BookTimeStart
	}

	if !dara.IsNil(request.PageIndex) {
		query["page_index"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OrderList"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/order-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Trade - Order List
//
// Description:
//
// query order list
//
// @param request - OrderListRequest
//
// @return OrderListResponse
func (client *Client) OrderList(request *OrderListRequest) (_result *OrderListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &OrderListHeaders{}
	_result = &OrderListResponse{}
	_body, _err := client.OrderListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Transaction - Seat and Price Verification
//
// Description:
//
// Check is price and remaining seats of solution you selected has changed. You should enter the solution_id returned by enrich.
//
// @param request - PricingRequest
//
// @param headers - PricingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PricingResponse
func (client *Client) PricingWithOptions(request *PricingRequest, headers *PricingHeaders, runtime *dara.RuntimeOptions) (_result *PricingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SolutionId) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Pricing"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-pricing"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PricingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Transaction - Seat and Price Verification
//
// Description:
//
// Check is price and remaining seats of solution you selected has changed. You should enter the solution_id returned by enrich.
//
// @param request - PricingRequest
//
// @return PricingResponse
func (client *Client) Pricing(request *PricingRequest) (_result *PricingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PricingHeaders{}
	_result = &PricingResponse{}
	_body, _err := client.PricingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Ticket Refund - Application
//
// @param tmpReq - RefundApplyRequest
//
// @param headers - RefundApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundApplyResponse
func (client *Client) RefundApplyWithOptions(tmpReq *RefundApplyRequest, headers *RefundApplyHeaders, runtime *dara.RuntimeOptions) (_result *RefundApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefundApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RefundJourneys) {
		request.RefundJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundJourneys, dara.String("refund_journeys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RefundPassengerList) {
		request.RefundPassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundPassengerList, dara.String("refund_passenger_list"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RefundType) {
		request.RefundTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundType, dara.String("refund_type"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	if !dara.IsNil(request.RefundJourneysShrink) {
		body["refund_journeys"] = request.RefundJourneysShrink
	}

	if !dara.IsNil(request.RefundPassengerListShrink) {
		body["refund_passenger_list"] = request.RefundPassengerListShrink
	}

	if !dara.IsNil(request.RefundTypeShrink) {
		body["refund_type"] = request.RefundTypeShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundApply"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/refund/action-apply"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Ticket Refund - Application
//
// @param request - RefundApplyRequest
//
// @return RefundApplyResponse
func (client *Client) RefundApply(request *RefundApplyRequest) (_result *RefundApplyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RefundApplyHeaders{}
	_result = &RefundApplyResponse{}
	_body, _err := client.RefundApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Refund - Detail
//
// @param request - RefundDetailRequest
//
// @param headers - RefundDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundDetailResponse
func (client *Client) RefundDetailWithOptions(request *RefundDetailRequest, headers *RefundDetailHeaders, runtime *dara.RuntimeOptions) (_result *RefundDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RefundOrderNum) {
		query["refund_order_num"] = request.RefundOrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundDetail"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/refund/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Refund - Detail
//
// @param request - RefundDetailRequest
//
// @return RefundDetailResponse
func (client *Client) RefundDetail(request *RefundDetailRequest) (_result *RefundDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RefundDetailHeaders{}
	_result = &RefundDetailResponse{}
	_body, _err := client.RefundDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Refund - Detail List
//
// @param request - RefundDetailListRequest
//
// @param headers - RefundDetailListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundDetailListResponse
func (client *Client) RefundDetailListWithOptions(request *RefundDetailListRequest, headers *RefundDetailListHeaders, runtime *dara.RuntimeOptions) (_result *RefundDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		query["order_num"] = request.OrderNum
	}

	if !dara.IsNil(request.PageIndex) {
		query["page_index"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.RefundCreateBeginTime) {
		query["refund_create_begin_time"] = request.RefundCreateBeginTime
	}

	if !dara.IsNil(request.RefundCreateEndTime) {
		query["refund_create_end_time"] = request.RefundCreateEndTime
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundDetailList"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/refund/detail-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Refund - Detail List
//
// @param request - RefundDetailListRequest
//
// @return RefundDetailListResponse
func (client *Client) RefundDetailList(request *RefundDetailListRequest) (_result *RefundDetailListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RefundDetailListHeaders{}
	_result = &RefundDetailListResponse{}
	_body, _err := client.RefundDetailListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Search
//
// Description:
//
// Enter the information of departure, arrival, departure date, passenger number and cabin, return the lowest price for each flight.
//
// @param tmpReq - SearchRequest
//
// @param headers - SearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResponse
func (client *Client) SearchWithOptions(tmpReq *SearchRequest, headers *SearchHeaders, runtime *dara.RuntimeOptions) (_result *SearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AirLegs) {
		request.AirLegsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AirLegs, dara.String("air_legs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchControlOptions) {
		request.SearchControlOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchControlOptions, dara.String("search_control_options"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Adults) {
		body["adults"] = request.Adults
	}

	if !dara.IsNil(request.AirLegsShrink) {
		body["air_legs"] = request.AirLegsShrink
	}

	if !dara.IsNil(request.CabinClass) {
		body["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.Children) {
		body["children"] = request.Children
	}

	if !dara.IsNil(request.Infants) {
		body["infants"] = request.Infants
	}

	if !dara.IsNil(request.SearchControlOptionsShrink) {
		body["search_control_options"] = request.SearchControlOptionsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Search"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Search
//
// Description:
//
// Enter the information of departure, arrival, departure date, passenger number and cabin, return the lowest price for each flight.
//
// @param request - SearchRequest
//
// @return SearchResponse
func (client *Client) Search(request *SearchRequest) (_result *SearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchHeaders{}
	_result = &SearchResponse{}
	_body, _err := client.SearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 标准搜索
//
// @param tmpReq - StandardSearchRequest
//
// @param headers - StandardSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StandardSearchResponse
func (client *Client) StandardSearchWithOptions(tmpReq *StandardSearchRequest, headers *StandardSearchHeaders, runtime *dara.RuntimeOptions) (_result *StandardSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StandardSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AirLegs) {
		request.AirLegsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AirLegs, dara.String("air_legs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SearchControlOptions) {
		request.SearchControlOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchControlOptions, dara.String("search_control_options"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Adults) {
		body["adults"] = request.Adults
	}

	if !dara.IsNil(request.AirLegsShrink) {
		body["air_legs"] = request.AirLegsShrink
	}

	if !dara.IsNil(request.CabinClass) {
		body["cabin_class"] = request.CabinClass
	}

	if !dara.IsNil(request.Children) {
		body["children"] = request.Children
	}

	if !dara.IsNil(request.Infants) {
		body["infants"] = request.Infants
	}

	if !dara.IsNil(request.SearchControlOptionsShrink) {
		body["search_control_options"] = request.SearchControlOptionsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StandardSearch"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-standardsearch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StandardSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 标准搜索
//
// @param request - StandardSearchRequest
//
// @return StandardSearchResponse
func (client *Client) StandardSearch(request *StandardSearchRequest) (_result *StandardSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StandardSearchHeaders{}
	_result = &StandardSearchResponse{}
	_body, _err := client.StandardSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Transaction - Payment and Ticket Issuance
//
// @param request - TicketingRequest
//
// @param headers - TicketingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketingResponse
func (client *Client) TicketingWithOptions(request *TicketingRequest, headers *TicketingHeaders, runtime *dara.RuntimeOptions) (_result *TicketingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Ticketing"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-ticketing"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Transaction - Payment and Ticket Issuance
//
// @param request - TicketingRequest
//
// @return TicketingResponse
func (client *Client) Ticketing(request *TicketingRequest) (_result *TicketingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &TicketingHeaders{}
	_result = &TicketingResponse{}
	_body, _err := client.TicketingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Transaction - Pre-payment verification
//
// Description:
//
// Pre-check for Ticketing, this interface is optional to use.
//
// @param request - TicketingCheckRequest
//
// @param headers - TicketingCheckHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketingCheckResponse
func (client *Client) TicketingCheckWithOptions(request *TicketingCheckRequest, headers *TicketingCheckHeaders, runtime *dara.RuntimeOptions) (_result *TicketingCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderNum) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketingCheck"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/trade/action-ticketing-check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketingCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Transaction - Pre-payment verification
//
// Description:
//
// Pre-check for Ticketing, this interface is optional to use.
//
// @param request - TicketingCheckRequest
//
// @return TicketingCheckResponse
func (client *Client) TicketingCheck(request *TicketingCheckRequest) (_result *TicketingCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &TicketingCheckHeaders{}
	_result = &TicketingCheckResponse{}
	_body, _err := client.TicketingCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航程过境签
//
// @param tmpReq - TransitVisaRequest
//
// @param headers - TransitVisaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransitVisaResponse
func (client *Client) TransitVisaWithOptions(tmpReq *TransitVisaRequest, headers *TransitVisaHeaders, runtime *dara.RuntimeOptions) (_result *TransitVisaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TransitVisaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FlightSegmentParamList) {
		request.FlightSegmentParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlightSegmentParamList, dara.String("flight_segment_param_list"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FlightSegmentParamListShrink) {
		query["flight_segment_param_list"] = request.FlightSegmentParamListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAirticketAccessToken) {
		realHeaders["x-acs-airticket-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketAccessToken)))
	}

	if !dara.IsNil(headers.XAcsAirticketLanguage) {
		realHeaders["x-acs-airticket-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAirticketLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransitVisa"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/airticket/v1/flight-data/transit-visa"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransitVisaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航程过境签
//
// @param request - TransitVisaRequest
//
// @return TransitVisaResponse
func (client *Client) TransitVisa(request *TransitVisaRequest) (_result *TransitVisaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &TransitVisaHeaders{}
	_result = &TransitVisaResponse{}
	_body, _err := client.TransitVisaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
