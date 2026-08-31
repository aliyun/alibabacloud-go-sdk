// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a customer note.
//
// @param request - CustomerNoteCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerNoteCreateResponse
func (client *Client) CustomerNoteCreateWithContext(ctx context.Context, request *CustomerNoteCreateRequest, runtime *dara.RuntimeOptions) (_result *CustomerNoteCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactInformation) {
		body["ContactInformation"] = request.ContactInformation
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.CustomerName) {
		body["CustomerName"] = request.CustomerName
	}

	if !dara.IsNil(request.CustomerUid) {
		body["CustomerUid"] = request.CustomerUid
	}

	if !dara.IsNil(request.NoteContent) {
		body["NoteContent"] = request.NoteContent
	}

	if !dara.IsNil(request.TouchDate) {
		body["TouchDate"] = request.TouchDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerNoteCreate"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerNoteCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a customer note.
//
// @param request - CustomerNoteEditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerNoteEditResponse
func (client *Client) CustomerNoteEditWithContext(ctx context.Context, request *CustomerNoteEditRequest, runtime *dara.RuntimeOptions) (_result *CustomerNoteEditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContactInformation) {
		body["ContactInformation"] = request.ContactInformation
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.NoteContent) {
		body["NoteContent"] = request.NoteContent
	}

	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	if !dara.IsNil(request.TouchDate) {
		body["TouchDate"] = request.TouchDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerNoteEdit"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerNoteEditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a paged query list of customer notes with paging support.
//
// @param request - CustomerNoteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerNoteListResponse
func (client *Client) CustomerNoteListWithContext(ctx context.Context, request *CustomerNoteListRequest, runtime *dara.RuntimeOptions) (_result *CustomerNoteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomerUid) {
		body["CustomerUid"] = request.CustomerUid
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerNoteList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerNoteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a customer note.
//
// @param request - CustomerNoteListDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerNoteListDetailResponse
func (client *Client) CustomerNoteListDetailWithContext(ctx context.Context, request *CustomerNoteListDetailRequest, runtime *dara.RuntimeOptions) (_result *CustomerNoteListDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NoteId) {
		body["NoteId"] = request.NoteId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerNoteListDetail"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerNoteListDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries exported bill files.
//
// @param request - GetBillDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBillDetailFileListResponse
func (client *Client) GetBillDetailFileListWithContext(ctx context.Context, request *GetBillDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetBillDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBillDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBillDetailFileListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the commission details of a partner.
//
// @param request - GetCommissionDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommissionDetailFileListResponse
func (client *Client) GetCommissionDetailFileListWithContext(ctx context.Context, request *GetCommissionDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetCommissionDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommissionDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommissionDetailFileListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries partner customer acquisition orders.
//
// @param tmpReq - GetCustomerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrderListResponse
func (client *Client) GetCustomerOrderListWithContext(ctx context.Context, tmpReq *GetCustomerOrderListRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerOrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCustomerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderTypeList) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, dara.String("OrderTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerAccount) {
		query["CustomerAccount"] = request.CustomerAccount
	}

	if !dara.IsNil(request.CustomerUid) {
		query["CustomerUid"] = request.CustomerUid
	}

	if !dara.IsNil(request.OrderCreateAfter) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !dara.IsNil(request.OrderCreateBefore) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OrderPayAfter) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !dara.IsNil(request.OrderPayBefore) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.OrderTypeListShrink) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayAmountAfter) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !dara.IsNil(request.PayAmountBefore) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RamAccountForCustomerManager) {
		query["RamAccountForCustomerManager"] = request.RamAccountForCustomerManager
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerOrderList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerOrderListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the commission details of an international partner.
//
// Description:
//
// Ensure that the current caller identity is a T1 distribution partner.
//
// <notice>Available only for international site.</notice>.
//
// @param request - GetIntlCommissionDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntlCommissionDetailFileListResponse
func (client *Client) GetIntlCommissionDetailFileListWithContext(ctx context.Context, request *GetIntlCommissionDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetIntlCommissionDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntlCommissionDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntlCommissionDetailFileListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the partner renewal rate.
//
// @param request - GetRenewalRateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenewalRateListResponse
func (client *Client) GetRenewalRateListWithContext(ctx context.Context, request *GetRenewalRateListRequest, runtime *dara.RuntimeOptions) (_result *GetRenewalRateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FiscalYearAndQuarter) {
		query["FiscalYearAndQuarter"] = request.FiscalYearAndQuarter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRenewalRateList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRenewalRateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of secondary distributors.
//
// @param request - GetSubPartnerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerListResponse
func (client *Client) GetSubPartnerListWithContext(ctx context.Context, request *GetSubPartnerListRequest, runtime *dara.RuntimeOptions) (_result *GetSubPartnerListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SubPartnerCompanyName) {
		query["SubPartnerCompanyName"] = request.SubPartnerCompanyName
	}

	if !dara.IsNil(request.SubPartnerPid) {
		query["SubPartnerPid"] = request.SubPartnerPid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubPartnerList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubPartnerListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries channel expansion orders.
//
// @param tmpReq - GetSubPartnerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerOrderListResponse
func (client *Client) GetSubPartnerOrderListWithContext(ctx context.Context, tmpReq *GetSubPartnerOrderListRequest, runtime *dara.RuntimeOptions) (_result *GetSubPartnerOrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSubPartnerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderTypeList) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, dara.String("OrderTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderCreateAfter) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !dara.IsNil(request.OrderCreateBefore) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OrderPayAfter) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !dara.IsNil(request.OrderPayBefore) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.OrderTypeListShrink) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayAmountAfter) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !dara.IsNil(request.PayAmountBefore) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubPartnerName) {
		query["SubPartnerName"] = request.SubPartnerName
	}

	if !dara.IsNil(request.SubPartnerUid) {
		query["SubPartnerUid"] = request.SubPartnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubPartnerOrderList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubPartnerOrderListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
