// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Closes a ticket.
//
// @param request - CloseTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseTicketResponse
func (client *Client) CloseTicketWithContext(ctx context.Context, request *CloseTicketRequest, runtime *dara.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket.
//
// @param tmpReq - CreateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, tmpReq *CreateTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileNameList) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, dara.String("FileNameList"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.SecretInfo) {
		request.SecretInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecretInfo, dara.String("SecretInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretInfoShrink) {
		query["SecretInfo"] = request.SecretInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		body["Email"] = request.Email
	}

	if !dara.IsNil(request.FileNameListShrink) {
		body["FileNameList"] = request.FileNameListShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evaluates a ticket.
//
// @param request - EvaluateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateTicketResponse
func (client *Client) EvaluateTicketWithContext(ctx context.Context, request *EvaluateTicketRequest, runtime *dara.RuntimeOptions) (_result *EvaluateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Score) {
		body["Score"] = request.Score
	}

	if !dara.IsNil(request.Solved) {
		body["Solved"] = request.Solved
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Object Storage Service (OSS) URL that is used to upload attachments.
//
// @param request - GetAttachmentUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttachmentUploadUrlResponse
func (client *Client) GetAttachmentUploadUrlWithContext(ctx context.Context, request *GetAttachmentUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAttachmentUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttachmentUploadUrl"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttachmentUploadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query tickets.
//
// @param request - GetTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketResponse
func (client *Client) GetTicketWithContext(ctx context.Context, request *GetTicketRequest, runtime *dara.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list data of ticket problem categories.
//
// @param request - ListCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithContext(ctx context.Context, request *ListCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCategories"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the data of the Alibaba Cloud product list.
//
// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithContext(ctx context.Context, request *ListProductsRequest, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the ticket communication records.
//
// @param request - ListTicketNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketNotesResponse
func (client *Client) ListTicketNotesWithContext(ctx context.Context, request *ListTicketNotesRequest, runtime *dara.RuntimeOptions) (_result *ListTicketNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTicketNotes"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain the list of my tickets.
//
// @param tmpReq - ListTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketsResponse
func (client *Client) ListTicketsWithContext(ctx context.Context, tmpReq *ListTicketsRequest, runtime *dara.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TicketIdList) {
		request.TicketIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketIdList, dara.String("TicketIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatusList) {
		body["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.TicketIdListShrink) {
		body["TicketIdList"] = request.TicketIdListShrink
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTickets"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Reopens a ticket
//
// @param request - ReopenTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenTicketResponse
func (client *Client) ReopenTicketWithContext(ctx context.Context, request *ReopenTicketRequest, runtime *dara.RuntimeOptions) (_result *ReopenTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reply to the ticket. You can call the ListTicketNotes operation to obtain the content of the reply.
//
// @param tmpReq - ReplyTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplyTicketResponse
func (client *Client) ReplyTicketWithContext(ctx context.Context, tmpReq *ReplyTicketRequest, runtime *dara.RuntimeOptions) (_result *ReplyTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplyTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileNameList) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, dara.String("FileNameList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FileNameListShrink) {
		query["FileNameList"] = request.FileNameListShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Encrypt) {
		body["Encrypt"] = request.Encrypt
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplyTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplyTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
