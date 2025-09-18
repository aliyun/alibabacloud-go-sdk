// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 方法描述：删除消息
//
// @param request - DelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelMessageResponse
func (client *Client) DelMessageWithContext(ctx context.Context, request *DelMessageRequest, runtime *dara.RuntimeOptions) (_result *DelMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：站内信全部删除（逻辑删除）
//
// @param request - DeleteAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessageWithContext(ctx context.Context, request *DeleteAllMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：分类全部标记为已读，不填则全部标记
//
// @param request - ReadAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessageWithContext(ctx context.Context, request *ReadAllMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadClassNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadClassNameResponse
func (client *Client) ReadClassNameWithContext(ctx context.Context, request *ReadClassNameRequest, runtime *dara.RuntimeOptions) (_result *ReadClassNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadClassName"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadClassNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：消息标记为已读
//
// @param request - ReadMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithContext(ctx context.Context, request *ReadMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息正文
//
// @param request - ReadMessageContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContentWithContext(ctx context.Context, request *ReadMessageContentRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageContent"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息列表
//
// @param request - ReadMessageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageListWithContext(ctx context.Context, request *ReadMessageListRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.Loc) {
		body["Loc"] = request.Loc
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取未读消息总数
//
// @param request - ReadMessageNewTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotalWithContext(ctx context.Context, request *ReadMessageNewTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageNewTotalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageNewTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadNumGroupByClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClassWithContext(ctx context.Context, request *ReadNumGroupByClassRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupByClassResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupByClass"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取所有分类下的信息
//
// @param request - ReadNumGroupTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotalWithContext(ctx context.Context, request *ReadNumGroupTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupTotalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
