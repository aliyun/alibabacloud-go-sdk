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
	client.Endpoint, _err = client.GetEndpoint(dara.String("marketing_event"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddSumRecordFlowPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSumRecordFlowPopResponse
func (client *Client) AddSumRecordFlowPopWithOptions(request *AddSumRecordFlowPopRequest, runtime *dara.RuntimeOptions) (_result *AddSumRecordFlowPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ConferenceName) {
		query["ConferenceName"] = request.ConferenceName
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.EntryName) {
		query["EntryName"] = request.EntryName
	}

	if !dara.IsNil(request.Idcard) {
		query["Idcard"] = request.Idcard
	}

	if !dara.IsNil(request.SignTime) {
		query["SignTime"] = request.SignTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSumRecordFlowPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSumRecordFlowPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddSumRecordFlowPopRequest
//
// @return AddSumRecordFlowPopResponse
func (client *Client) AddSumRecordFlowPop(request *AddSumRecordFlowPopRequest) (_result *AddSumRecordFlowPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSumRecordFlowPopResponse{}
	_body, _err := client.AddSumRecordFlowPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindExhibitorRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindExhibitorRfidPopResponse
func (client *Client) BindExhibitorRfidPopWithOptions(request *BindExhibitorRfidPopRequest, runtime *dara.RuntimeOptions) (_result *BindExhibitorRfidPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.GmtCreate) {
		query["GmtCreate"] = request.GmtCreate
	}

	if !dara.IsNil(request.GmtModified) {
		query["GmtModified"] = request.GmtModified
	}

	if !dara.IsNil(request.GuestTicketRecordId) {
		query["GuestTicketRecordId"] = request.GuestTicketRecordId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Rfid) {
		query["Rfid"] = request.Rfid
	}

	if !dara.IsNil(request.TicketCode) {
		query["TicketCode"] = request.TicketCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindExhibitorRfidPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindExhibitorRfidPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindExhibitorRfidPopRequest
//
// @return BindExhibitorRfidPopResponse
func (client *Client) BindExhibitorRfidPop(request *BindExhibitorRfidPopRequest) (_result *BindExhibitorRfidPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindExhibitorRfidPopResponse{}
	_body, _err := client.BindExhibitorRfidPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindGuestRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindGuestRfidPopResponse
func (client *Client) BindGuestRfidPopWithOptions(request *BindGuestRfidPopRequest, runtime *dara.RuntimeOptions) (_result *BindGuestRfidPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.GmtCreate) {
		query["GmtCreate"] = request.GmtCreate
	}

	if !dara.IsNil(request.GmtModified) {
		query["GmtModified"] = request.GmtModified
	}

	if !dara.IsNil(request.GuestTicketRecordId) {
		query["GuestTicketRecordId"] = request.GuestTicketRecordId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Rfid) {
		query["Rfid"] = request.Rfid
	}

	if !dara.IsNil(request.TicketCode) {
		query["TicketCode"] = request.TicketCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindGuestRfidPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindGuestRfidPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindGuestRfidPopRequest
//
// @return BindGuestRfidPopResponse
func (client *Client) BindGuestRfidPop(request *BindGuestRfidPopRequest) (_result *BindGuestRfidPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindGuestRfidPopResponse{}
	_body, _err := client.BindGuestRfidPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CheckNFCBindPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckNFCBindPopResponse
func (client *Client) CheckNFCBindPopWithOptions(request *CheckNFCBindPopRequest, runtime *dara.RuntimeOptions) (_result *CheckNFCBindPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.NfcId) {
		query["NfcId"] = request.NfcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckNFCBindPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckNFCBindPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckNFCBindPopRequest
//
// @return CheckNFCBindPopResponse
func (client *Client) CheckNFCBindPop(request *CheckNFCBindPopRequest) (_result *CheckNFCBindPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckNFCBindPopResponse{}
	_body, _err := client.CheckNFCBindPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拉取领证人员记录
//
// @param request - FindGuestCredentialsRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestCredentialsRecordResponse
func (client *Client) FindGuestCredentialsRecordWithOptions(request *FindGuestCredentialsRecordRequest, runtime *dara.RuntimeOptions) (_result *FindGuestCredentialsRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.DateTimeString) {
		query["DateTimeString"] = request.DateTimeString
	}

	if !dara.IsNil(request.EndDateTime) {
		query["EndDateTime"] = request.EndDateTime
	}

	if !dara.IsNil(request.StartDateTime) {
		query["StartDateTime"] = request.StartDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindGuestCredentialsRecord"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindGuestCredentialsRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取领证人员记录
//
// @param request - FindGuestCredentialsRecordRequest
//
// @return FindGuestCredentialsRecordResponse
func (client *Client) FindGuestCredentialsRecord(request *FindGuestCredentialsRecordRequest) (_result *FindGuestCredentialsRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindGuestCredentialsRecordResponse{}
	_body, _err := client.FindGuestCredentialsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云栖大会获取取票人信息list接口
//
// @param request - FindGuestTicketRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestTicketRecordResponse
func (client *Client) FindGuestTicketRecordWithOptions(request *FindGuestTicketRecordRequest, runtime *dara.RuntimeOptions) (_result *FindGuestTicketRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.DateTimeString) {
		query["DateTimeString"] = request.DateTimeString
	}

	if !dara.IsNil(request.EndDateTime) {
		query["EndDateTime"] = request.EndDateTime
	}

	if !dara.IsNil(request.StartDateTime) {
		query["StartDateTime"] = request.StartDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindGuestTicketRecord"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindGuestTicketRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云栖大会获取取票人信息list接口
//
// @param request - FindGuestTicketRecordRequest
//
// @return FindGuestTicketRecordResponse
func (client *Client) FindGuestTicketRecord(request *FindGuestTicketRecordRequest) (_result *FindGuestTicketRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindGuestTicketRecordResponse{}
	_body, _err := client.FindGuestTicketRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAllActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllActivityInfoResponse
func (client *Client) QueryAllActivityInfoWithOptions(request *QueryAllActivityInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAllActivityInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAllActivityInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAllActivityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAllActivityInfoRequest
//
// @return QueryAllActivityInfoResponse
func (client *Client) QueryAllActivityInfo(request *QueryAllActivityInfoRequest) (_result *QueryAllActivityInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAllActivityInfoResponse{}
	_body, _err := client.QueryAllActivityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryOrderSessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderSessionListPopResponse
func (client *Client) QueryOrderSessionListPopWithOptions(request *QueryOrderSessionListPopRequest, runtime *dara.RuntimeOptions) (_result *QueryOrderSessionListPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.NfcId) {
		query["NfcId"] = request.NfcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOrderSessionListPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrderSessionListPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOrderSessionListPopRequest
//
// @return QueryOrderSessionListPopResponse
func (client *Client) QueryOrderSessionListPop(request *QueryOrderSessionListPopRequest) (_result *QueryOrderSessionListPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryOrderSessionListPopResponse{}
	_body, _err := client.QueryOrderSessionListPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询新加坡千问大会票证信息
//
// @param request - QueryQwenConferenceSgTicketPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQwenConferenceSgTicketPopResponse
func (client *Client) QueryQwenConferenceSgTicketPopWithOptions(request *QueryQwenConferenceSgTicketPopRequest, runtime *dara.RuntimeOptions) (_result *QueryQwenConferenceSgTicketPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TicketToken) {
		query["TicketToken"] = request.TicketToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryQwenConferenceSgTicketPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryQwenConferenceSgTicketPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询新加坡千问大会票证信息
//
// @param request - QueryQwenConferenceSgTicketPopRequest
//
// @return QueryQwenConferenceSgTicketPopResponse
func (client *Client) QueryQwenConferenceSgTicketPop(request *QueryQwenConferenceSgTicketPopRequest) (_result *QueryQwenConferenceSgTicketPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryQwenConferenceSgTicketPopResponse{}
	_body, _err := client.QueryQwenConferenceSgTicketPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 千问大会搜索票据信息
//
// @param request - QueryQwenConferenceSgTicketSearchPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQwenConferenceSgTicketSearchPopResponse
func (client *Client) QueryQwenConferenceSgTicketSearchPopWithOptions(request *QueryQwenConferenceSgTicketSearchPopRequest, runtime *dara.RuntimeOptions) (_result *QueryQwenConferenceSgTicketSearchPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryQwenConferenceSgTicketSearchPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryQwenConferenceSgTicketSearchPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 千问大会搜索票据信息
//
// @param request - QueryQwenConferenceSgTicketSearchPopRequest
//
// @return QueryQwenConferenceSgTicketSearchPopResponse
func (client *Client) QueryQwenConferenceSgTicketSearchPop(request *QueryQwenConferenceSgTicketSearchPopRequest) (_result *QueryQwenConferenceSgTicketSearchPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryQwenConferenceSgTicketSearchPopResponse{}
	_body, _err := client.QueryQwenConferenceSgTicketSearchPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySessionByActivityIdPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionByActivityIdPopResponse
func (client *Client) QuerySessionByActivityIdPopWithOptions(request *QuerySessionByActivityIdPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionByActivityIdPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionByActivityIdPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionByActivityIdPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionByActivityIdPopRequest
//
// @return QuerySessionByActivityIdPopResponse
func (client *Client) QuerySessionByActivityIdPop(request *QuerySessionByActivityIdPopRequest) (_result *QuerySessionByActivityIdPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySessionByActivityIdPopResponse{}
	_body, _err := client.QuerySessionByActivityIdPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionListPopResponse
func (client *Client) QuerySessionListPopWithOptions(request *QuerySessionListPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionListPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.NfcId) {
		query["NfcId"] = request.NfcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionListPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionListPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionListPopRequest
//
// @return QuerySessionListPopResponse
func (client *Client) QuerySessionListPop(request *QuerySessionListPopRequest) (_result *QuerySessionListPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySessionListPopResponse{}
	_body, _err := client.QuerySessionListPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySignInRecordPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySignInRecordPopResponse
func (client *Client) QuerySignInRecordPopWithOptions(request *QuerySignInRecordPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySignInRecordPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySignInRecordPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySignInRecordPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySignInRecordPopRequest
//
// @return QuerySignInRecordPopResponse
func (client *Client) QuerySignInRecordPop(request *QuerySignInRecordPopRequest) (_result *QuerySignInRecordPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySignInRecordPopResponse{}
	_body, _err := client.QuerySignInRecordPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySingleActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleActivityInfoResponse
func (client *Client) QuerySingleActivityInfoWithOptions(request *QuerySingleActivityInfoRequest, runtime *dara.RuntimeOptions) (_result *QuerySingleActivityInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySingleActivityInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySingleActivityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySingleActivityInfoRequest
//
// @return QuerySingleActivityInfoResponse
func (client *Client) QuerySingleActivityInfo(request *QuerySingleActivityInfoRequest) (_result *QuerySingleActivityInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySingleActivityInfoResponse{}
	_body, _err := client.QuerySingleActivityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SyncSignInInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSignInInfoResponse
func (client *Client) SyncSignInInfoWithOptions(request *SyncSignInInfoRequest, runtime *dara.RuntimeOptions) (_result *SyncSignInInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncSignInInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncSignInInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncSignInInfoRequest
//
// @return SyncSignInInfoResponse
func (client *Client) SyncSignInInfo(request *SyncSignInInfoRequest) (_result *SyncSignInInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncSignInInfoResponse{}
	_body, _err := client.SyncSignInInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TicketOrCredentialsSignInPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketOrCredentialsSignInPopResponse
func (client *Client) TicketOrCredentialsSignInPopWithOptions(request *TicketOrCredentialsSignInPopRequest, runtime *dara.RuntimeOptions) (_result *TicketOrCredentialsSignInPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ConferenceName) {
		query["ConferenceName"] = request.ConferenceName
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.EntryName) {
		query["EntryName"] = request.EntryName
	}

	if !dara.IsNil(request.Idcard) {
		query["Idcard"] = request.Idcard
	}

	if !dara.IsNil(request.SignTime) {
		query["SignTime"] = request.SignTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TicketOrCredentialsSignInPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TicketOrCredentialsSignInPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TicketOrCredentialsSignInPopRequest
//
// @return TicketOrCredentialsSignInPopResponse
func (client *Client) TicketOrCredentialsSignInPop(request *TicketOrCredentialsSignInPopRequest) (_result *TicketOrCredentialsSignInPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TicketOrCredentialsSignInPopResponse{}
	_body, _err := client.TicketOrCredentialsSignInPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCredentialsStatusPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialsStatusPopResponse
func (client *Client) UpdateCredentialsStatusPopWithOptions(request *UpdateCredentialsStatusPopRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialsStatusPopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ProxyRecipientName) {
		query["ProxyRecipientName"] = request.ProxyRecipientName
	}

	if !dara.IsNil(request.ProxyRecipientPhoneNumber) {
		query["ProxyRecipientPhoneNumber"] = request.ProxyRecipientPhoneNumber
	}

	if !dara.IsNil(request.ReceiptLocation) {
		query["ReceiptLocation"] = request.ReceiptLocation
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredentialsStatusPop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialsStatusPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCredentialsStatusPopRequest
//
// @return UpdateCredentialsStatusPopResponse
func (client *Client) UpdateCredentialsStatusPop(request *UpdateCredentialsStatusPopRequest) (_result *UpdateCredentialsStatusPopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCredentialsStatusPopResponse{}
	_body, _err := client.UpdateCredentialsStatusPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTicketRecordByticketCodePopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketRecordByticketCodePopResponse
func (client *Client) UpdateTicketRecordByticketCodePopWithOptions(request *UpdateTicketRecordByticketCodePopRequest, runtime *dara.RuntimeOptions) (_result *UpdateTicketRecordByticketCodePopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgendaId) {
		query["AgendaId"] = request.AgendaId
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.Event) {
		query["Event"] = request.Event
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTicketRecordByticketCodePop"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTicketRecordByticketCodePopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTicketRecordByticketCodePopRequest
//
// @return UpdateTicketRecordByticketCodePopResponse
func (client *Client) UpdateTicketRecordByticketCodePop(request *UpdateTicketRecordByticketCodePopRequest) (_result *UpdateTicketRecordByticketCodePopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTicketRecordByticketCodePopResponse{}
	_body, _err := client.UpdateTicketRecordByticketCodePopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
