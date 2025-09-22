// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AddSumRecordFlowPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSumRecordFlowPopResponse
func (client *Client) AddSumRecordFlowPopWithContext(ctx context.Context, request *AddSumRecordFlowPopRequest, runtime *dara.RuntimeOptions) (_result *AddSumRecordFlowPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindExhibitorRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindExhibitorRfidPopResponse
func (client *Client) BindExhibitorRfidPopWithContext(ctx context.Context, request *BindExhibitorRfidPopRequest, runtime *dara.RuntimeOptions) (_result *BindExhibitorRfidPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindGuestRfidPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindGuestRfidPopResponse
func (client *Client) BindGuestRfidPopWithContext(ctx context.Context, request *BindGuestRfidPopRequest, runtime *dara.RuntimeOptions) (_result *BindGuestRfidPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckNFCBindPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckNFCBindPopResponse
func (client *Client) CheckNFCBindPopWithContext(ctx context.Context, request *CheckNFCBindPopRequest, runtime *dara.RuntimeOptions) (_result *CheckNFCBindPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestCredentialsRecordResponse
func (client *Client) FindGuestCredentialsRecordWithContext(ctx context.Context, request *FindGuestCredentialsRecordRequest, runtime *dara.RuntimeOptions) (_result *FindGuestCredentialsRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindGuestTicketRecordResponse
func (client *Client) FindGuestTicketRecordWithContext(ctx context.Context, request *FindGuestTicketRecordRequest, runtime *dara.RuntimeOptions) (_result *FindGuestTicketRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAllActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllActivityInfoResponse
func (client *Client) QueryAllActivityInfoWithContext(ctx context.Context, request *QueryAllActivityInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAllActivityInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOrderSessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderSessionListPopResponse
func (client *Client) QueryOrderSessionListPopWithContext(ctx context.Context, request *QueryOrderSessionListPopRequest, runtime *dara.RuntimeOptions) (_result *QueryOrderSessionListPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionByActivityIdPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionByActivityIdPopResponse
func (client *Client) QuerySessionByActivityIdPopWithContext(ctx context.Context, request *QuerySessionByActivityIdPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionByActivityIdPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySessionListPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionListPopResponse
func (client *Client) QuerySessionListPopWithContext(ctx context.Context, request *QuerySessionListPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySessionListPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySignInRecordPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySignInRecordPopResponse
func (client *Client) QuerySignInRecordPopWithContext(ctx context.Context, request *QuerySignInRecordPopRequest, runtime *dara.RuntimeOptions) (_result *QuerySignInRecordPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySingleActivityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleActivityInfoResponse
func (client *Client) QuerySingleActivityInfoWithContext(ctx context.Context, request *QuerySingleActivityInfoRequest, runtime *dara.RuntimeOptions) (_result *QuerySingleActivityInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncSignInInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSignInInfoResponse
func (client *Client) SyncSignInInfoWithContext(ctx context.Context, request *SyncSignInInfoRequest, runtime *dara.RuntimeOptions) (_result *SyncSignInInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TicketOrCredentialsSignInPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketOrCredentialsSignInPopResponse
func (client *Client) TicketOrCredentialsSignInPopWithContext(ctx context.Context, request *TicketOrCredentialsSignInPopRequest, runtime *dara.RuntimeOptions) (_result *TicketOrCredentialsSignInPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCredentialsStatusPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialsStatusPopResponse
func (client *Client) UpdateCredentialsStatusPopWithContext(ctx context.Context, request *UpdateCredentialsStatusPopRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialsStatusPopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTicketRecordByticketCodePopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketRecordByticketCodePopResponse
func (client *Client) UpdateTicketRecordByticketCodePopWithContext(ctx context.Context, request *UpdateTicketRecordByticketCodePopRequest, runtime *dara.RuntimeOptions) (_result *UpdateTicketRecordByticketCodePopResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
