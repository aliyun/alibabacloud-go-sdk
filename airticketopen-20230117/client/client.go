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
// Queries the account fund flow list.
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
// Queries the account fund flow list.
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
// Recommends ancillary products.
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
// Recommends ancillary products.
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
// Creates a booking order.
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
// Creates a booking order.
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
// Cancels an unpaid order.
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
// Cancels an unpaid order.
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
// Submits a change application.
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
// Submits a change application.
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
// Cancel the change order.
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
// Cancel the change order.
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
// Confirms a flight change order.
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
// Confirms a flight change order.
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
// Retrieves the details of a flight change order.
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
// Retrieves the details of a flight change order.
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
// Queries a paging list of change order summaries by buyer account.
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
// Queries a paging list of change order summaries by buyer account.
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
// Queries the list of change orders by the original order number.
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
// Queries the list of change orders by the original order number.
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
// Collects lowest-price flight information.
//
// Description:
//
// Collects lowest-price flight information.
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
// Collects lowest-price flight information.
//
// Description:
//
// Collects lowest-price flight information.
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
// Find richer quote information for the itinerary, including free baggage allowance, refund and change rules, and baggage through-check rules.
//
// Description:
//
// In the input parameters, choose either solution_id or journey_param_list. solution_id must be obtained from the Search API.
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
// Find richer quote information for the itinerary, including free baggage allowance, refund and change rules, and baggage through-check rules.
//
// Description:
//
// In the input parameters, choose either solution_id or journey_param_list. solution_id must be obtained from the Search API.
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
// Uploads a file as an attachment image. The file size is limited to 300 KB or less.
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
// Uploads a file as an attachment image. The file size is limited to 300 KB or less.
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
// Queries flight change information by order number.
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
// Queries flight change information by order number.
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
// Obtains a token for API calls. The token is valid for 2 hours.
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
// Obtains a token for API calls. The token is valid for 2 hours.
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
// Queries luggage through-check information for an itinerary.
//
// Description:
//
// Queries luggage through-check information for an itinerary. Provide itinerary information as input, and the API returns whether luggage through-check is supported for the itinerary. Luggage through-check applies to transfer and stopover scenarios.
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
// Queries luggage through-check information for an itinerary.
//
// Description:
//
// Queries luggage through-check information for an itinerary. Provide itinerary information as input, and the API returns whether luggage through-check is supported for the itinerary. Luggage through-check applies to transfer and stopover scenarios.
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
// Queries order details.
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
// Queries order details.
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
// Queries the order list.
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
// Queries the order list.
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
// Verifies seat availability and pricing. If the price has changed, the developer can proceed with Book at the updated price. If the price has not changed, the order is placed at the original price.
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
// Verifies seat availability and pricing. If the price has changed, the developer can proceed with Book at the updated price. If the price has not changed, the order is placed at the original price.
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
// Submits a refund application for an air ticket.
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
// Submits a refund application for an air ticket.
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
// Retrieves the details of a refund order.
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
// Retrieves the details of a refund order.
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
// Queries the details of refund orders.
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
// Queries the details of refund orders.
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
// Searches for flight quotes and returns the lowest price across multiple flights. Note that the response of this operation does not include refund and change rules, free baggage allowance, or baggage through-check rules.
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
// Searches for flight quotes and returns the lowest price across multiple flights. Note that the response of this operation does not include refund and change rules, free baggage allowance, or baggage through-check rules.
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
// Search and quote prices, currently providing the lowest price across multiple flights. Note that this API response includes refund/change rules, free baggage allowance, and baggage through-check rules.
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
// Search and quote prices, currently providing the lowest price across multiple flights. Note that this API response includes refund/change rules, free baggage allowance, and baggage through-check rules.
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
// Pays for and issues a ticket.
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
// Pays for and issues a ticket.
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
// Performs a pre-ticketing check. This operation is optional.
//
// Description:
//
// Performs a pre-ticketing check. This operation is optional.
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
// Performs a pre-ticketing check. This operation is optional.
//
// Description:
//
// Performs a pre-ticketing check. This operation is optional.
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
// Queries transit visa requirements for a flight itinerary. You provide flight information, and the API returns whether a transit visa is required for the itinerary. Only transfer or stopover segments are valid input parameters (transfers or stopovers passing through a third country). The supported passenger type defaults to Chinese mainland travelers.
//
// Description:
//
// Queries transit visa requirements for a flight itinerary. You provide flight information, and the API returns whether a transit visa is required for the itinerary. Only transfer or stopover segments are valid input parameters (transfers or stopovers passing through a third country). The supported passenger type defaults to Chinese mainland travelers.
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
// Queries transit visa requirements for a flight itinerary. You provide flight information, and the API returns whether a transit visa is required for the itinerary. Only transfer or stopover segments are valid input parameters (transfers or stopovers passing through a third country). The supported passenger type defaults to Chinese mainland travelers.
//
// Description:
//
// Queries transit visa requirements for a flight itinerary. You provide flight information, and the API returns whether a transit visa is required for the itinerary. Only transfer or stopover segments are valid input parameters (transfers or stopovers passing through a third country). The supported passenger type defaults to Chinese mainland travelers.
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

// Summary:
//
// 申请退款
//
// @param request - ApplyRefundRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyRefundResponse
func (client *Client) ApplyRefundWithOptions(request *ApplyRefundRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyRefundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.RefundReason) {
		body["RefundReason"] = request.RefundReason
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("applyRefund"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/applyRefund"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请退款
//
// @param request - ApplyRefundRequest
//
// @return ApplyRefundResponse
func (client *Client) ApplyRefund(request *ApplyRefundRequest) (_result *ApplyRefundResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyRefundResponse{}
	_body, _err := client.ApplyRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询酒店详情
//
// @param tmpReq - BatchGetHotelDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetHotelDetailResponse
func (client *Client) BatchGetHotelDetailWithOptions(tmpReq *BatchGetHotelDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetHotelDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetHotelDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("batchGetHotelDetail"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/batchGetHotelDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetHotelDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询酒店详情
//
// @param request - BatchGetHotelDetailRequest
//
// @return BatchGetHotelDetailResponse
func (client *Client) BatchGetHotelDetail(request *BatchGetHotelDetailRequest) (_result *BatchGetHotelDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetHotelDetailResponse{}
	_body, _err := client.BatchGetHotelDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消或退款
//
// @param request - CancelOrRefundRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrRefundResponse
func (client *Client) CancelOrRefundWithOptions(request *CancelOrRefundRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelOrRefundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("cancelOrRefund"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cancelOrRefund"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOrRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消或退款
//
// @param request - CancelOrRefundRequest
//
// @return CancelOrRefundResponse
func (client *Client) CancelOrRefund(request *CancelOrRefundRequest) (_result *CancelOrRefundResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelOrRefundResponse{}
	_body, _err := client.CancelOrRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消订单
//
// @param request - CancelOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderResponse
func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("cancelOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cancelOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消订单
//
// @param request - CancelOrderRequest
//
// @return CancelOrderResponse
func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创单并支付
//
// @param tmpReq - CreateAndPayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndPayResponse
func (client *Client) CreateAndPayWithOptions(tmpReq *CreateAndPayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAndPayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAndPayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("Contact"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Guests) {
		request.GuestsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Guests, dara.String("Guests"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ContactShrink) {
		body["Contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.GuestsShrink) {
		body["Guests"] = request.GuestsShrink
	}

	if !dara.IsNil(request.ItemOfferId) {
		body["ItemOfferId"] = request.ItemOfferId
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("createAndPay"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/createAndPay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndPayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创单并支付
//
// @param request - CreateAndPayRequest
//
// @return CreateAndPayResponse
func (client *Client) CreateAndPay(request *CreateAndPayRequest) (_result *CreateAndPayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAndPayResponse{}
	_body, _err := client.CreateAndPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param tmpReq - CreateOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(tmpReq *CreateOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("Contact"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Guests) {
		request.GuestsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Guests, dara.String("Guests"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ContactShrink) {
		body["Contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.GuestsShrink) {
		body["Guests"] = request.GuestsShrink
	}

	if !dara.IsNil(request.ItemOfferId) {
		body["ItemOfferId"] = request.ItemOfferId
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("createOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/createOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建订单
//
// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a refund.
//
// @param request - GlobalHotelApplyRefundRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelApplyRefundResponse
func (client *Client) GlobalHotelApplyRefundWithOptions(request *GlobalHotelApplyRefundRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelApplyRefundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.RefundReason) {
		body["RefundReason"] = request.RefundReason
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelApplyRefund"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelApplyRefund"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelApplyRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a refund.
//
// @param request - GlobalHotelApplyRefundRequest
//
// @return GlobalHotelApplyRefundResponse
func (client *Client) GlobalHotelApplyRefund(request *GlobalHotelApplyRefundRequest) (_result *GlobalHotelApplyRefundResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelApplyRefundResponse{}
	_body, _err := client.GlobalHotelApplyRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries hotel details in batches.
//
// @param tmpReq - GlobalHotelBatchGetHotelDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelBatchGetHotelDetailResponse
func (client *Client) GlobalHotelBatchGetHotelDetailWithOptions(tmpReq *GlobalHotelBatchGetHotelDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelBatchGetHotelDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelBatchGetHotelDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelBatchGetHotelDetail"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelBatchGetHotelDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelBatchGetHotelDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries hotel details in batches.
//
// @param request - GlobalHotelBatchGetHotelDetailRequest
//
// @return GlobalHotelBatchGetHotelDetailResponse
func (client *Client) GlobalHotelBatchGetHotelDetail(request *GlobalHotelBatchGetHotelDetailRequest) (_result *GlobalHotelBatchGetHotelDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelBatchGetHotelDetailResponse{}
	_body, _err := client.GlobalHotelBatchGetHotelDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels or refunds an order.
//
// @param request - GlobalHotelCancelOrRefundRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelCancelOrRefundResponse
func (client *Client) GlobalHotelCancelOrRefundWithOptions(request *GlobalHotelCancelOrRefundRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelCancelOrRefundResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelCancelOrRefund"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelCancelOrRefund"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelCancelOrRefundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels or refunds an order.
//
// @param request - GlobalHotelCancelOrRefundRequest
//
// @return GlobalHotelCancelOrRefundResponse
func (client *Client) GlobalHotelCancelOrRefund(request *GlobalHotelCancelOrRefundRequest) (_result *GlobalHotelCancelOrRefundResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelCancelOrRefundResponse{}
	_body, _err := client.GlobalHotelCancelOrRefundWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an order.
//
// @param request - GlobalHotelCancelOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelCancelOrderResponse
func (client *Client) GlobalHotelCancelOrderWithOptions(request *GlobalHotelCancelOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelCancelOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelCancelOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelCancelOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelCancelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an order.
//
// @param request - GlobalHotelCancelOrderRequest
//
// @return GlobalHotelCancelOrderResponse
func (client *Client) GlobalHotelCancelOrder(request *GlobalHotelCancelOrderRequest) (_result *GlobalHotelCancelOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelCancelOrderResponse{}
	_body, _err := client.GlobalHotelCancelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an order and processes the payment.
//
// @param tmpReq - GlobalHotelCreateAndPayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelCreateAndPayResponse
func (client *Client) GlobalHotelCreateAndPayWithOptions(tmpReq *GlobalHotelCreateAndPayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelCreateAndPayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelCreateAndPayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("Contact"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Guests) {
		request.GuestsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Guests, dara.String("Guests"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ContactShrink) {
		body["Contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.GuestsShrink) {
		body["Guests"] = request.GuestsShrink
	}

	if !dara.IsNil(request.ItemOfferId) {
		body["ItemOfferId"] = request.ItemOfferId
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelCreateAndPay"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelCreateAndPay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelCreateAndPayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an order and processes the payment.
//
// @param request - GlobalHotelCreateAndPayRequest
//
// @return GlobalHotelCreateAndPayResponse
func (client *Client) GlobalHotelCreateAndPay(request *GlobalHotelCreateAndPayRequest) (_result *GlobalHotelCreateAndPayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelCreateAndPayResponse{}
	_body, _err := client.GlobalHotelCreateAndPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an order.
//
// @param tmpReq - GlobalHotelCreateOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelCreateOrderResponse
func (client *Client) GlobalHotelCreateOrderWithOptions(tmpReq *GlobalHotelCreateOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelCreateOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelCreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Contact) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, dara.String("Contact"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Guests) {
		request.GuestsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Guests, dara.String("Guests"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ContactShrink) {
		body["Contact"] = request.ContactShrink
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.GuestsShrink) {
		body["Guests"] = request.GuestsShrink
	}

	if !dara.IsNil(request.ItemOfferId) {
		body["ItemOfferId"] = request.ItemOfferId
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelCreateOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelCreateOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelCreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an order.
//
// @param request - GlobalHotelCreateOrderRequest
//
// @return GlobalHotelCreateOrderResponse
func (client *Client) GlobalHotelCreateOrder(request *GlobalHotelCreateOrderRequest) (_result *GlobalHotelCreateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelCreateOrderResponse{}
	_body, _err := client.GlobalHotelCreateOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Processes a distribution payment.
//
// @param request - GlobalHotelPayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelPayResponse
func (client *Client) GlobalHotelPayWithOptions(request *GlobalHotelPayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelPayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelPay"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelPay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelPayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Processes a distribution payment.
//
// @param request - GlobalHotelPayRequest
//
// @return GlobalHotelPayResponse
func (client *Client) GlobalHotelPay(request *GlobalHotelPayRequest) (_result *GlobalHotelPayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelPayResponse{}
	_body, _err := client.GlobalHotelPayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the availability of hotel rate plans.
//
// @param tmpReq - GlobalHotelQueryAvailabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelQueryAvailabilityResponse
func (client *Client) GlobalHotelQueryAvailabilityWithOptions(tmpReq *GlobalHotelQueryAvailabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelQueryAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelQueryAvailabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.AdultCount) {
		body["AdultCount"] = request.AdultCount
	}

	if !dara.IsNil(request.CheckInDate) {
		body["CheckInDate"] = request.CheckInDate
	}

	if !dara.IsNil(request.CheckOutDate) {
		body["CheckOutDate"] = request.CheckOutDate
	}

	if !dara.IsNil(request.ChildCount) {
		body["ChildCount"] = request.ChildCount
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelQueryAvailability"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelQueryAvailability"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelQueryAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the availability of hotel rate plans.
//
// @param request - GlobalHotelQueryAvailabilityRequest
//
// @return GlobalHotelQueryAvailabilityResponse
func (client *Client) GlobalHotelQueryAvailability(request *GlobalHotelQueryAvailabilityRequest) (_result *GlobalHotelQueryAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelQueryAvailabilityResponse{}
	_body, _err := client.GlobalHotelQueryAvailabilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries calendar-based rate availability for hotels in batch.
//
// @param tmpReq - GlobalHotelQueryCalendarAvailabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelQueryCalendarAvailabilityResponse
func (client *Client) GlobalHotelQueryCalendarAvailabilityWithOptions(tmpReq *GlobalHotelQueryCalendarAvailabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelQueryCalendarAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelQueryCalendarAvailabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.AdultCount) {
		body["AdultCount"] = request.AdultCount
	}

	if !dara.IsNil(request.CheckInDateEnd) {
		body["CheckInDateEnd"] = request.CheckInDateEnd
	}

	if !dara.IsNil(request.CheckInDateStart) {
		body["CheckInDateStart"] = request.CheckInDateStart
	}

	if !dara.IsNil(request.ChildCount) {
		body["ChildCount"] = request.ChildCount
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelQueryCalendarAvailability"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelQueryCalendarAvailability"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelQueryCalendarAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries calendar-based rate availability for hotels in batch.
//
// @param request - GlobalHotelQueryCalendarAvailabilityRequest
//
// @return GlobalHotelQueryCalendarAvailabilityResponse
func (client *Client) GlobalHotelQueryCalendarAvailability(request *GlobalHotelQueryCalendarAvailabilityRequest) (_result *GlobalHotelQueryCalendarAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelQueryCalendarAvailabilityResponse{}
	_body, _err := client.GlobalHotelQueryCalendarAvailabilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an order.
//
// @param request - GlobalHotelQueryOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelQueryOrderResponse
func (client *Client) GlobalHotelQueryOrderWithOptions(request *GlobalHotelQueryOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelQueryOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelQueryOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelQueryOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelQueryOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an order.
//
// @param request - GlobalHotelQueryOrderRequest
//
// @return GlobalHotelQueryOrderResponse
func (client *Client) GlobalHotelQueryOrder(request *GlobalHotelQueryOrderRequest) (_result *GlobalHotelQueryOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelQueryOrderResponse{}
	_body, _err := client.GlobalHotelQueryOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries city administrative divisions (in Chinese and English) by paging.
//
// @param request - GlobalHotelSearchCityPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelSearchCityPageResponse
func (client *Client) GlobalHotelSearchCityPageWithOptions(request *GlobalHotelSearchCityPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelSearchCityPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Count) {
		body["Count"] = request.Count
	}

	if !dara.IsNil(request.CountryCode) {
		body["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelSearchCityPage"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelSearchCityPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelSearchCityPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries city administrative divisions (in Chinese and English) by paging.
//
// @param request - GlobalHotelSearchCityPageRequest
//
// @return GlobalHotelSearchCityPageResponse
func (client *Client) GlobalHotelSearchCityPage(request *GlobalHotelSearchCityPageRequest) (_result *GlobalHotelSearchCityPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelSearchCityPageResponse{}
	_body, _err := client.GlobalHotelSearchCityPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of the hotel list by city with paging support.
//
// @param request - GlobalHotelSearchHotelListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelSearchHotelListResponse
func (client *Client) GlobalHotelSearchHotelListWithOptions(request *GlobalHotelSearchHotelListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelSearchHotelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.CityCode) {
		body["CityCode"] = request.CityCode
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelSearchHotelList"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelSearchHotelList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelSearchHotelListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of the hotel list by city with paging support.
//
// @param request - GlobalHotelSearchHotelListRequest
//
// @return GlobalHotelSearchHotelListResponse
func (client *Client) GlobalHotelSearchHotelList(request *GlobalHotelSearchHotelListRequest) (_result *GlobalHotelSearchHotelListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelSearchHotelListResponse{}
	_body, _err := client.GlobalHotelSearchHotelListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates the price of a hotel offer.
//
// @param tmpReq - GlobalHotelValidatePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalHotelValidatePriceResponse
func (client *Client) GlobalHotelValidatePriceWithOptions(tmpReq *GlobalHotelValidatePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GlobalHotelValidatePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GlobalHotelValidatePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Adults) {
		body["Adults"] = request.Adults
	}

	if !dara.IsNil(request.Children) {
		body["Children"] = request.Children
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.ItemOfferKey) {
		body["ItemOfferKey"] = request.ItemOfferKey
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("globalHotelValidatePrice"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotelValidatePrice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GlobalHotelValidatePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the price of a hotel offer.
//
// @param request - GlobalHotelValidatePriceRequest
//
// @return GlobalHotelValidatePriceResponse
func (client *Client) GlobalHotelValidatePrice(request *GlobalHotelValidatePriceRequest) (_result *GlobalHotelValidatePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalHotelValidatePriceResponse{}
	_body, _err := client.GlobalHotelValidatePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分销支付
//
// @param request - PayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PayResponse
func (client *Client) PayWithOptions(request *PayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("pay"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pay"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分销支付
//
// @param request - PayRequest
//
// @return PayResponse
func (client *Client) Pay(request *PayRequest) (_result *PayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PayResponse{}
	_body, _err := client.PayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询酒店报价可用性
//
// @param tmpReq - QueryAvailabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailabilityResponse
func (client *Client) QueryAvailabilityWithOptions(tmpReq *QueryAvailabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryAvailabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.AdultCount) {
		body["AdultCount"] = request.AdultCount
	}

	if !dara.IsNil(request.CheckInDate) {
		body["CheckInDate"] = request.CheckInDate
	}

	if !dara.IsNil(request.CheckOutDate) {
		body["CheckOutDate"] = request.CheckOutDate
	}

	if !dara.IsNil(request.ChildCount) {
		body["ChildCount"] = request.ChildCount
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("queryAvailability"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queryAvailability"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询酒店报价可用性
//
// @param request - QueryAvailabilityRequest
//
// @return QueryAvailabilityResponse
func (client *Client) QueryAvailability(request *QueryAvailabilityRequest) (_result *QueryAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAvailabilityResponse{}
	_body, _err := client.QueryAvailabilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量日历报价查询
//
// @param tmpReq - QueryCalendarAvailabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCalendarAvailabilityResponse
func (client *Client) QueryCalendarAvailabilityWithOptions(tmpReq *QueryCalendarAvailabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryCalendarAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryCalendarAvailabilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StandardHotelIds) {
		request.StandardHotelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardHotelIds, dara.String("StandardHotelIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.AdultCount) {
		body["AdultCount"] = request.AdultCount
	}

	if !dara.IsNil(request.CheckInDateEnd) {
		body["CheckInDateEnd"] = request.CheckInDateEnd
	}

	if !dara.IsNil(request.CheckInDateStart) {
		body["CheckInDateStart"] = request.CheckInDateStart
	}

	if !dara.IsNil(request.ChildCount) {
		body["ChildCount"] = request.ChildCount
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.StandardHotelIdsShrink) {
		body["StandardHotelIds"] = request.StandardHotelIdsShrink
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("queryCalendarAvailability"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queryCalendarAvailability"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCalendarAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量日历报价查询
//
// @param request - QueryCalendarAvailabilityRequest
//
// @return QueryCalendarAvailabilityResponse
func (client *Client) QueryCalendarAvailability(request *QueryCalendarAvailabilityRequest) (_result *QueryCalendarAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCalendarAvailabilityResponse{}
	_body, _err := client.QueryCalendarAvailabilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - QueryOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderResponse
func (client *Client) QueryOrderWithOptions(request *QueryOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.ExternalOrderNo) {
		body["ExternalOrderNo"] = request.ExternalOrderNo
	}

	if !dara.IsNil(request.OrderNo) {
		body["OrderNo"] = request.OrderNo
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("queryOrder"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/queryOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订单
//
// @param request - QueryOrderRequest
//
// @return QueryOrderResponse
func (client *Client) QueryOrder(request *QueryOrderRequest) (_result *QueryOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryOrderResponse{}
	_body, _err := client.QueryOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询城市行政区划（中英文）
//
// @param request - SearchCityPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCityPageResponse
func (client *Client) SearchCityPageWithOptions(request *SearchCityPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchCityPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Count) {
		body["Count"] = request.Count
	}

	if !dara.IsNil(request.CountryCode) {
		body["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("searchCityPage"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/searchCityPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCityPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询城市行政区划（中英文）
//
// @param request - SearchCityPageRequest
//
// @return SearchCityPageResponse
func (client *Client) SearchCityPage(request *SearchCityPageRequest) (_result *SearchCityPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchCityPageResponse{}
	_body, _err := client.SearchCityPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按城市分页查询酒店列表
//
// @param request - SearchHotelListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchHotelListResponse
func (client *Client) SearchHotelListWithOptions(request *SearchHotelListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchHotelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.CityCode) {
		body["CityCode"] = request.CityCode
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("searchHotelList"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/globalHotel/searchHotelList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchHotelListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按城市分页查询酒店列表
//
// @param request - SearchHotelListRequest
//
// @return SearchHotelListResponse
func (client *Client) SearchHotelList(request *SearchHotelListRequest) (_result *SearchHotelListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchHotelListResponse{}
	_body, _err := client.SearchHotelListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验价
//
// @param tmpReq - ValidatePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidatePriceResponse
func (client *Client) ValidatePriceWithOptions(tmpReq *ValidatePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidatePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ValidatePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChildrenAges) {
		request.ChildrenAgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChildrenAges, dara.String("ChildrenAges"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		body["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.Adults) {
		body["Adults"] = request.Adults
	}

	if !dara.IsNil(request.Children) {
		body["Children"] = request.Children
	}

	if !dara.IsNil(request.ChildrenAgesShrink) {
		body["ChildrenAges"] = request.ChildrenAgesShrink
	}

	if !dara.IsNil(request.ItemOfferKey) {
		body["ItemOfferKey"] = request.ItemOfferKey
	}

	if !dara.IsNil(request.RoomCount) {
		body["RoomCount"] = request.RoomCount
	}

	if !dara.IsNil(request.TracerId) {
		body["TracerId"] = request.TracerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("validatePrice"),
		Version:     dara.String("2023-01-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/validatePrice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidatePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验价
//
// @param request - ValidatePriceRequest
//
// @return ValidatePriceResponse
func (client *Client) ValidatePrice(request *ValidatePriceRequest) (_result *ValidatePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidatePriceResponse{}
	_body, _err := client.ValidatePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
