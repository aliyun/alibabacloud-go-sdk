// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加日程参与者
//
// @param tmpReq - AddAttendeeRequest
//
// @param tmpHeader - AddAttendeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAttendeeResponse
func (client *Client) AddAttendeeWithContext(ctx context.Context, tmpReq *AddAttendeeRequest, tmpHeader *AddAttendeeHeaders, runtime *dara.RuntimeOptions) (_result *AddAttendeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAttendeeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddAttendeeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AttendeesToAdd) {
		request.AttendeesToAddShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttendeesToAdd, dara.String("AttendeesToAdd"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttendeesToAddShrink) {
		body["AttendeesToAdd"] = request.AttendeesToAddShrink
	}

	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.ChatNotification) {
		body["chatNotification"] = request.ChatNotification
	}

	if !dara.IsNil(request.PushNotification) {
		body["pushNotification"] = request.PushNotification
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAttendee"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/addAttendee"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAttendeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建钉盘空间
//
// @param tmpReq - AddDriveSpaceRequest
//
// @param tmpHeader - AddDriveSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDriveSpaceResponse
func (client *Client) AddDriveSpaceWithContext(ctx context.Context, tmpReq *AddDriveSpaceRequest, tmpHeader *AddDriveSpaceHeaders, runtime *dara.RuntimeOptions) (_result *AddDriveSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddDriveSpaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddDriveSpaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDriveSpace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/addDriveSpace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDriveSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文件夹
//
// @param tmpReq - AddFolderRequest
//
// @param tmpHeader - AddFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFolderResponse
func (client *Client) AddFolderWithContext(ctx context.Context, tmpReq *AddFolderRequest, tmpHeader *AddFolderHeaders, runtime *dara.RuntimeOptions) (_result *AddFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddFolderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddFolderShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.ParentId) {
		body["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFolder"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/addFolder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预定会议室
//
// @param tmpReq - AddMeetingRoomsRequest
//
// @param tmpHeader - AddMeetingRoomsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMeetingRoomsResponse
func (client *Client) AddMeetingRoomsWithContext(ctx context.Context, tmpReq *AddMeetingRoomsRequest, tmpHeader *AddMeetingRoomsHeaders, runtime *dara.RuntimeOptions) (_result *AddMeetingRoomsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddMeetingRoomsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddMeetingRoomsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MeetingRoomsToAdd) {
		request.MeetingRoomsToAddShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MeetingRoomsToAdd, dara.String("MeetingRoomsToAdd"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.MeetingRoomsToAddShrink) {
		body["MeetingRoomsToAdd"] = request.MeetingRoomsToAddShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMeetingRooms"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/addMeetingRooms"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMeetingRoomsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增数据表
//
// @param tmpReq - AddMultiDimTableRequest
//
// @param tmpHeader - AddMultiDimTableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMultiDimTableResponse
func (client *Client) AddMultiDimTableWithContext(ctx context.Context, tmpReq *AddMultiDimTableRequest, tmpHeader *AddMultiDimTableHeaders, runtime *dara.RuntimeOptions) (_result *AddMultiDimTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddMultiDimTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddMultiDimTableShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Fields) {
		request.FieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Fields, dara.String("Fields"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.FieldsShrink) {
		body["Fields"] = request.FieldsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMultiDimTable"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/addMultiDimTable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMultiDimTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文件权限
//
// @param tmpReq - AddPermissionRequest
//
// @param tmpHeader - AddPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPermissionResponse
func (client *Client) AddPermissionWithContext(ctx context.Context, tmpReq *AddPermissionRequest, tmpHeader *AddPermissionHeaders, runtime *dara.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddPermissionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.RoleId) {
		body["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPermission"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/addPermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加闪记权限
//
// @param tmpReq - AddRecordPermissionRequest
//
// @param tmpHeader - AddRecordPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordPermissionResponse
func (client *Client) AddRecordPermissionWithContext(ctx context.Context, tmpReq *AddRecordPermissionRequest, tmpHeader *AddRecordPermissionHeaders, runtime *dara.RuntimeOptions) (_result *AddRecordPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddRecordPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddRecordPermissionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ConferenceId) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecordPermission"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/addRecordPermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecordPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增群成员
//
// @param request - AddScenegroupMemberRequest
//
// @param tmpHeader - AddScenegroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddScenegroupMemberResponse
func (client *Client) AddScenegroupMemberWithContext(ctx context.Context, request *AddScenegroupMemberRequest, tmpHeader *AddScenegroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *AddScenegroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &AddScenegroupMemberShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddScenegroupMember"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/addScenegroupMember"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddScenegroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工单添加备注
//
// @param tmpReq - AddTicketMemoRequest
//
// @param tmpHeader - AddTicketMemoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTicketMemoResponse
func (client *Client) AddTicketMemoWithContext(ctx context.Context, tmpReq *AddTicketMemoRequest, tmpHeader *AddTicketMemoHeaders, runtime *dara.RuntimeOptions) (_result *AddTicketMemoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTicketMemoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddTicketMemoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketMemo) {
		request.TicketMemoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketMemo, dara.String("TicketMemo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.TicketMemoShrink) {
		body["TicketMemo"] = request.TicketMemoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTicketMemo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/addTicketMemo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTicketMemoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建知识库
//
// @param tmpReq - AddWorkspaceRequest
//
// @param tmpHeader - AddWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceResponse
func (client *Client) AddWorkspaceWithContext(ctx context.Context, tmpReq *AddWorkspaceRequest, tmpHeader *AddWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddWorkspaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWorkspace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/addWorkspace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加知识库文档成员
//
// @param tmpReq - AddWorkspaceDocMembersRequest
//
// @param tmpHeader - AddWorkspaceDocMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceDocMembersResponse
func (client *Client) AddWorkspaceDocMembersWithContext(ctx context.Context, tmpReq *AddWorkspaceDocMembersRequest, tmpHeader *AddWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceDocMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWorkspaceDocMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/addWorkspaceDocMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加知识库成员
//
// @param tmpReq - AddWorkspaceMembersRequest
//
// @param tmpHeader - AddWorkspaceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceMembersResponse
func (client *Client) AddWorkspaceMembersWithContext(ctx context.Context, tmpReq *AddWorkspaceMembersRequest, tmpHeader *AddWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWorkspaceMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/addWorkspaceMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指派工单
//
// @param tmpReq - AssignTicketRequest
//
// @param tmpHeader - AssignTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignTicketResponse
func (client *Client) AssignTicketWithContext(ctx context.Context, tmpReq *AssignTicketRequest, tmpHeader *AssignTicketHeaders, runtime *dara.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssignTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AssignTicketShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notify) {
		request.NotifyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notify, dara.String("Notify"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProcessorUserIds) {
		request.ProcessorUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProcessorUserIds, dara.String("ProcessorUserIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketMemo) {
		request.TicketMemoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketMemo, dara.String("TicketMemo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NotifyShrink) {
		body["Notify"] = request.NotifyShrink
	}

	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.ProcessorUserIdsShrink) {
		body["ProcessorUserIds"] = request.ProcessorUserIdsShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.TicketMemoShrink) {
		body["TicketMemo"] = request.TicketMemoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignTicket"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/assignTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验AI技能调用权限
//
// @param tmpReq - AuthorizeSkillRequest
//
// @param tmpHeader - AuthorizeSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeSkillResponse
func (client *Client) AuthorizeSkillWithContext(ctx context.Context, tmpReq *AuthorizeSkillRequest, tmpHeader *AuthorizeSkillHeaders, runtime *dara.RuntimeOptions) (_result *AuthorizeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthorizeSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AuthorizeSkillShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PermissionCodes) {
		request.PermissionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PermissionCodes, dara.String("PermissionCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PermissionCodesShrink) {
		body["PermissionCodes"] = request.PermissionCodesShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeSkill"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/skill/authorizeSkill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取表单实例数据
//
// @param tmpReq - BatchGetFormDataByIdListRequest
//
// @param tmpHeader - BatchGetFormDataByIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetFormDataByIdListResponse
func (client *Client) BatchGetFormDataByIdListWithContext(ctx context.Context, tmpReq *BatchGetFormDataByIdListRequest, tmpHeader *BatchGetFormDataByIdListHeaders, runtime *dara.RuntimeOptions) (_result *BatchGetFormDataByIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetFormDataByIdListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &BatchGetFormDataByIdListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FormInstanceIdList) {
		request.FormInstanceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FormInstanceIdList, dara.String("FormInstanceIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormInstanceIdListShrink) {
		body["FormInstanceIdList"] = request.FormInstanceIdListShrink
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.NeedFormInstanceValue) {
		body["NeedFormInstanceValue"] = request.NeedFormInstanceValue
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetFormDataByIdList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/batchGetFormDataByIdList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetFormDataByIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除表单实例
//
// @param tmpReq - BatchRemovalByFormInstanceIdListRequest
//
// @param tmpHeader - BatchRemovalByFormInstanceIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRemovalByFormInstanceIdListResponse
func (client *Client) BatchRemovalByFormInstanceIdListWithContext(ctx context.Context, tmpReq *BatchRemovalByFormInstanceIdListRequest, tmpHeader *BatchRemovalByFormInstanceIdListHeaders, runtime *dara.RuntimeOptions) (_result *BatchRemovalByFormInstanceIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchRemovalByFormInstanceIdListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &BatchRemovalByFormInstanceIdListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FormInstanceIdList) {
		request.FormInstanceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FormInstanceIdList, dara.String("FormInstanceIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AsynchronousExecution) {
		body["AsynchronousExecution"] = request.AsynchronousExecution
	}

	if !dara.IsNil(request.ExecuteExpression) {
		body["ExecuteExpression"] = request.ExecuteExpression
	}

	if !dara.IsNil(request.FormInstanceIdListShrink) {
		body["FormInstanceIdList"] = request.FormInstanceIdListShrink
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRemovalByFormInstanceIdList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/batchRemovalByFormInstanceIdList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRemovalByFormInstanceIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建表单实例
//
// @param tmpReq - BatchSaveFormDataRequest
//
// @param tmpHeader - BatchSaveFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSaveFormDataResponse
func (client *Client) BatchSaveFormDataWithContext(ctx context.Context, tmpReq *BatchSaveFormDataRequest, tmpHeader *BatchSaveFormDataHeaders, runtime *dara.RuntimeOptions) (_result *BatchSaveFormDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchSaveFormDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &BatchSaveFormDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FormDataJsonList) {
		request.FormDataJsonListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FormDataJsonList, dara.String("FormDataJsonList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AsynchronousExecution) {
		body["AsynchronousExecution"] = request.AsynchronousExecution
	}

	if !dara.IsNil(request.FormDataJsonListShrink) {
		body["FormDataJsonList"] = request.FormDataJsonListShrink
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.KeepRunningAfterException) {
		body["KeepRunningAfterException"] = request.KeepRunningAfterException
	}

	if !dara.IsNil(request.NoExecuteExpression) {
		body["NoExecuteExpression"] = request.NoExecuteExpression
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSaveFormData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/batchSaveFormData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSaveFormDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新表单实例内的组件值
//
// @param tmpReq - BatchUpdateFormDataByInstanceIdRequest
//
// @param tmpHeader - BatchUpdateFormDataByInstanceIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFormDataByInstanceIdResponse
func (client *Client) BatchUpdateFormDataByInstanceIdWithContext(ctx context.Context, tmpReq *BatchUpdateFormDataByInstanceIdRequest, tmpHeader *BatchUpdateFormDataByInstanceIdHeaders, runtime *dara.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateFormDataByInstanceIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &BatchUpdateFormDataByInstanceIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FormInstanceIdList) {
		request.FormInstanceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FormInstanceIdList, dara.String("FormInstanceIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AsynchronousExecution) {
		body["AsynchronousExecution"] = request.AsynchronousExecution
	}

	if !dara.IsNil(request.FormInstanceIdListShrink) {
		body["FormInstanceIdList"] = request.FormInstanceIdListShrink
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.IgnoreEmpty) {
		body["IgnoreEmpty"] = request.IgnoreEmpty
	}

	if !dara.IsNil(request.NoExecuteExpression) {
		body["NoExecuteExpression"] = request.NoExecuteExpression
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.UpdateFormDataJson) {
		body["UpdateFormDataJson"] = request.UpdateFormDataJson
	}

	if !dara.IsNil(request.UseLatestFormSchemaVersion) {
		body["UseLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateFormDataByInstanceId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/batchUpdateFormDataByInstanceId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateFormDataByInstanceIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过表单实例数据批量更新表单实例
//
// @param tmpReq - BatchUpdateFormDataByInstanceMapRequest
//
// @param tmpHeader - BatchUpdateFormDataByInstanceMapHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateFormDataByInstanceMapResponse
func (client *Client) BatchUpdateFormDataByInstanceMapWithContext(ctx context.Context, tmpReq *BatchUpdateFormDataByInstanceMapRequest, tmpHeader *BatchUpdateFormDataByInstanceMapHeaders, runtime *dara.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateFormDataByInstanceMapShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &BatchUpdateFormDataByInstanceMapShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateFormDataJsonMap) {
		request.UpdateFormDataJsonMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateFormDataJsonMap, dara.String("UpdateFormDataJsonMap"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AsynchronousExecution) {
		body["AsynchronousExecution"] = request.AsynchronousExecution
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.IgnoreEmpty) {
		body["IgnoreEmpty"] = request.IgnoreEmpty
	}

	if !dara.IsNil(request.NoExecuteExpression) {
		body["NoExecuteExpression"] = request.NoExecuteExpression
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.UpdateFormDataJsonMapShrink) {
		body["UpdateFormDataJsonMap"] = request.UpdateFormDataJsonMapShrink
	}

	if !dara.IsNil(request.UseLatestFormSchemaVersion) {
		body["UseLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateFormDataByInstanceMap"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/batchUpdateFormDataByInstanceMap"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateFormDataByInstanceMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消预约会议
//
// @param tmpReq - CancelScheduleConferenceRequest
//
// @param tmpHeader - CancelScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelScheduleConferenceResponse
func (client *Client) CancelScheduleConferenceWithContext(ctx context.Context, tmpReq *CancelScheduleConferenceRequest, tmpHeader *CancelScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CancelScheduleConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CancelScheduleConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CancelScheduleConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ScheduleConferenceId) {
		body["ScheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelScheduleConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/cancelScheduleConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelScheduleConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改阿里钉号
//
// @param tmpReq - ChangeDingTalkIdRequest
//
// @param tmpHeader - ChangeDingTalkIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDingTalkIdResponse
func (client *Client) ChangeDingTalkIdWithContext(ctx context.Context, tmpReq *ChangeDingTalkIdRequest, tmpHeader *ChangeDingTalkIdHeaders, runtime *dara.RuntimeOptions) (_result *ChangeDingTalkIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChangeDingTalkIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ChangeDingTalkIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DingTalkId) {
		body["DingTalkId"] = request.DingTalkId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDingTalkId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/user/changeDingTalkId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDingTalkIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验阿里员工
//
// @param tmpReq - CheckAlibabaStaffRequest
//
// @param tmpHeader - CheckAlibabaStaffHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAlibabaStaffResponse
func (client *Client) CheckAlibabaStaffWithContext(ctx context.Context, tmpReq *CheckAlibabaStaffRequest, tmpHeader *CheckAlibabaStaffHeaders, runtime *dara.RuntimeOptions) (_result *CheckAlibabaStaffResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckAlibabaStaffShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CheckAlibabaStaffShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAlibabaStaff"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/checkAlibabaStaff"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAlibabaStaffResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户是否为企业内部群成员
//
// @param request - CheckUserIsGroupMemberRequest
//
// @param tmpHeader - CheckUserIsGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUserIsGroupMemberResponse
func (client *Client) CheckUserIsGroupMemberWithContext(ctx context.Context, request *CheckUserIsGroupMemberRequest, tmpHeader *CheckUserIsGroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *CheckUserIsGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &CheckUserIsGroupMemberShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUserIsGroupMember"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/checkUserIsGroupMember"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUserIsGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除单元格所有内容
//
// @param tmpReq - ClearRequest
//
// @param tmpHeader - ClearHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearResponse
func (client *Client) ClearWithContext(ctx context.Context, tmpReq *ClearRequest, tmpHeader *ClearHeaders, runtime *dara.RuntimeOptions) (_result *ClearResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ClearShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ClearShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RangeAddress) {
		body["RangeAddress"] = request.RangeAddress
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Clear"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/clear"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除单元格数据
//
// @param tmpReq - ClearDataRequest
//
// @param tmpHeader - ClearDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearDataResponse
func (client *Client) ClearDataWithContext(ctx context.Context, tmpReq *ClearDataRequest, tmpHeader *ClearDataHeaders, runtime *dara.RuntimeOptions) (_result *ClearDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ClearDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ClearDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RangeAddress) {
		body["RangeAddress"] = request.RangeAddress
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/clearData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭视频会议
//
// @param tmpReq - CloseVideoConferenceRequest
//
// @param tmpHeader - CloseVideoConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseVideoConferenceResponse
func (client *Client) CloseVideoConferenceWithContext(ctx context.Context, tmpReq *CloseVideoConferenceRequest, tmpHeader *CloseVideoConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CloseVideoConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloseVideoConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CloseVideoConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseVideoConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/closeVideoConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志评论列表
//
// @param tmpReq - CommentListReportRequest
//
// @param tmpHeader - CommentListReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommentListReportResponse
func (client *Client) CommentListReportWithContext(ctx context.Context, tmpReq *CommentListReportRequest, tmpHeader *CommentListReportHeaders, runtime *dara.RuntimeOptions) (_result *CommentListReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CommentListReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CommentListReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommentListReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/commentListReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommentListReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交文件
//
// @param tmpReq - CommitFileRequest
//
// @param tmpHeader - CommitFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitFileResponse
func (client *Client) CommitFileWithContext(ctx context.Context, tmpReq *CommitFileRequest, tmpHeader *CommitFileHeaders, runtime *dara.RuntimeOptions) (_result *CommitFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CommitFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CommitFileShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.ParentDentryUuid) {
		body["ParentDentryUuid"] = request.ParentDentryUuid
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UploadKey) {
		body["UploadKey"] = request.UploadKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommitFile"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/commitFile"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建知识库节点副本
//
// @param tmpReq - CopyDentryRequest
//
// @param tmpHeader - CopyDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDentryResponse
func (client *Client) CopyDentryWithContext(ctx context.Context, tmpReq *CopyDentryRequest, tmpHeader *CopyDentryHeaders, runtime *dara.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CopyDentryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CopyDentryShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TargetSpaceId) {
		body["TargetSpaceId"] = request.TargetSpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ToNextDentryId) {
		body["ToNextDentryId"] = request.ToNextDentryId
	}

	if !dara.IsNil(request.ToParentDentryId) {
		body["ToParentDentryId"] = request.ToParentDentryId
	}

	if !dara.IsNil(request.ToPrevDentryId) {
		body["ToPrevDentryId"] = request.ToPrevDentryId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyDentry"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/copyDentry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyDentryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过NodeId创建知识库节点副本
//
// @param tmpReq - CopyDentryByNodeIdRequest
//
// @param tmpHeader - CopyDentryByNodeIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDentryByNodeIdResponse
func (client *Client) CopyDentryByNodeIdWithContext(ctx context.Context, tmpReq *CopyDentryByNodeIdRequest, tmpHeader *CopyDentryByNodeIdHeaders, runtime *dara.RuntimeOptions) (_result *CopyDentryByNodeIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CopyDentryByNodeIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CopyDentryByNodeIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ToNextNodeId) {
		body["ToNextNodeId"] = request.ToNextNodeId
	}

	if !dara.IsNil(request.ToParentNodeId) {
		body["ToParentNodeId"] = request.ToParentNodeId
	}

	if !dara.IsNil(request.ToPrevNodeId) {
		body["ToPrevNodeId"] = request.ToPrevNodeId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyDentryByNodeId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/copyDentryByNodeId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyDentryByNodeIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - CreateAlidingAssistantRequest
//
// @param tmpHeader - CreateAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlidingAssistantResponse
func (client *Client) CreateAlidingAssistantWithContext(ctx context.Context, tmpReq *CreateAlidingAssistantRequest, tmpHeader *CreateAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *CreateAlidingAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAlidingAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateAlidingAssistantShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecommendPrompts) {
		request.RecommendPromptsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecommendPrompts, dara.String("RecommendPrompts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		body["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExtShrink) {
		body["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Icon) {
		body["Icon"] = request.Icon
	}

	if !dara.IsNil(request.Instructions) {
		body["Instructions"] = request.Instructions
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecommendPromptsShrink) {
		body["RecommendPrompts"] = request.RecommendPromptsShrink
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIdentityId) {
		body["SourceIdentityId"] = request.SourceIdentityId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WelcomeContent) {
		body["WelcomeContent"] = request.WelcomeContent
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlidingAssistant"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/aiagent/createAlidingAssistant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlidingAssistantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布钉钉投放活动
//
// @param tmpReq - CreateDeliveryPlanRequest
//
// @param tmpHeader - CreateDeliveryPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryPlanResponse
func (client *Client) CreateDeliveryPlanWithContext(ctx context.Context, tmpReq *CreateDeliveryPlanRequest, tmpHeader *CreateDeliveryPlanHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeliveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDeliveryPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateDeliveryPlanShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResId) {
		body["ResId"] = request.ResId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeliveryPlan"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/createDeliveryPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeliveryPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限创建钉钉个人待办
//
// @param tmpReq - CreateDingtalkPersonalTodoTaskRequest
//
// @param tmpHeader - CreateDingtalkPersonalTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDingtalkPersonalTodoTaskResponse
func (client *Client) CreateDingtalkPersonalTodoTaskWithContext(ctx context.Context, tmpReq *CreateDingtalkPersonalTodoTaskRequest, tmpHeader *CreateDingtalkPersonalTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateDingtalkPersonalTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDingtalkPersonalTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateDingtalkPersonalTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutorIds) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, dara.String("ExecutorIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotifyConfigs) {
		request.NotifyConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyConfigs, dara.String("NotifyConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ParticipantIds) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, dara.String("ParticipantIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DueTime) {
		body["DueTime"] = request.DueTime
	}

	if !dara.IsNil(request.ExecutorIdsShrink) {
		body["ExecutorIds"] = request.ExecutorIdsShrink
	}

	if !dara.IsNil(request.NotifyConfigsShrink) {
		body["NotifyConfigs"] = request.NotifyConfigsShrink
	}

	if !dara.IsNil(request.ParticipantIdsShrink) {
		body["ParticipantIds"] = request.ParticipantIdsShrink
	}

	if !dara.IsNil(request.Subject) {
		body["Subject"] = request.Subject
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserToken) {
		body["UserToken"] = request.UserToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDingtalkPersonalTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/createDingtalkPersonalTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDingtalkPersonalTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建日程
//
// @param tmpReq - CreateEventRequest
//
// @param tmpHeader - CreateEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventResponse
func (client *Client) CreateEventWithContext(ctx context.Context, tmpReq *CreateEventRequest, tmpHeader *CreateEventHeaders, runtime *dara.RuntimeOptions) (_result *CreateEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Attendees) {
		request.AttendeesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attendees, dara.String("Attendees"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CardInstances) {
		request.CardInstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CardInstances, dara.String("CardInstances"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.End) {
		request.EndShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.End, dara.String("End"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Location) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, dara.String("Location"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OnlineMeetingInfo) {
		request.OnlineMeetingInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OnlineMeetingInfo, dara.String("OnlineMeetingInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Recurrence) {
		request.RecurrenceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Recurrence, dara.String("Recurrence"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Reminders) {
		request.RemindersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Reminders, dara.String("Reminders"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RichTextDescription) {
		request.RichTextDescriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RichTextDescription, dara.String("RichTextDescription"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UiConfigs) {
		request.UiConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UiConfigs, dara.String("UiConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Start) {
		request.StartShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Start, dara.String("start"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttendeesShrink) {
		body["Attendees"] = request.AttendeesShrink
	}

	if !dara.IsNil(request.CardInstancesShrink) {
		body["CardInstances"] = request.CardInstancesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndShrink) {
		body["End"] = request.EndShrink
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.IsAllDay) {
		body["IsAllDay"] = request.IsAllDay
	}

	if !dara.IsNil(request.LocationShrink) {
		body["Location"] = request.LocationShrink
	}

	if !dara.IsNil(request.OnlineMeetingInfoShrink) {
		body["OnlineMeetingInfo"] = request.OnlineMeetingInfoShrink
	}

	if !dara.IsNil(request.RecurrenceShrink) {
		body["Recurrence"] = request.RecurrenceShrink
	}

	if !dara.IsNil(request.RemindersShrink) {
		body["Reminders"] = request.RemindersShrink
	}

	if !dara.IsNil(request.RichTextDescriptionShrink) {
		body["RichTextDescription"] = request.RichTextDescriptionShrink
	}

	if !dara.IsNil(request.Summary) {
		body["Summary"] = request.Summary
	}

	if !dara.IsNil(request.UiConfigsShrink) {
		body["UiConfigs"] = request.UiConfigsShrink
	}

	if !dara.IsNil(request.CalendarId) {
		body["calendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.StartShrink) {
		body["start"] = request.StartShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/createEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建直播
//
// @param tmpReq - CreateLiveRequest
//
// @param tmpHeader - CreateLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveResponse
func (client *Client) CreateLiveWithContext(ctx context.Context, tmpReq *CreateLiveRequest, tmpHeader *CreateLiveHeaders, runtime *dara.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLiveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateLiveShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverUrl) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !dara.IsNil(request.Introduction) {
		body["Introduction"] = request.Introduction
	}

	if !dara.IsNil(request.PreEndTime) {
		body["PreEndTime"] = request.PreEndTime
	}

	if !dara.IsNil(request.PreStartTime) {
		body["PreStartTime"] = request.PreStartTime
	}

	if !dara.IsNil(request.PublicType) {
		body["PublicType"] = request.PublicType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLive"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/createLive"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会议室
//
// @param tmpReq - CreateMeetingRoomRequest
//
// @param tmpHeader - CreateMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMeetingRoomResponse
func (client *Client) CreateMeetingRoomWithContext(ctx context.Context, tmpReq *CreateMeetingRoomRequest, tmpHeader *CreateMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *CreateMeetingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMeetingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateMeetingRoomShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReservationAuthority) {
		request.ReservationAuthorityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReservationAuthority, dara.String("ReservationAuthority"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoomLabelIds) {
		request.RoomLabelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomLabelIds, dara.String("RoomLabelIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoomLocation) {
		request.RoomLocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomLocation, dara.String("RoomLocation"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableCycleReservation) {
		body["EnableCycleReservation"] = request.EnableCycleReservation
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsvRoomId) {
		body["IsvRoomId"] = request.IsvRoomId
	}

	if !dara.IsNil(request.ReservationAuthorityShrink) {
		body["ReservationAuthority"] = request.ReservationAuthorityShrink
	}

	if !dara.IsNil(request.RoomCapacity) {
		body["RoomCapacity"] = request.RoomCapacity
	}

	if !dara.IsNil(request.RoomLabelIdsShrink) {
		body["RoomLabelIds"] = request.RoomLabelIdsShrink
	}

	if !dara.IsNil(request.RoomLocationShrink) {
		body["RoomLocation"] = request.RoomLocationShrink
	}

	if !dara.IsNil(request.RoomName) {
		body["RoomName"] = request.RoomName
	}

	if !dara.IsNil(request.RoomPicture) {
		body["RoomPicture"] = request.RoomPicture
	}

	if !dara.IsNil(request.RoomStatus) {
		body["RoomStatus"] = request.RoomStatus
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMeetingRoom"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/createMeetingRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMeetingRoomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会议室分组
//
// @param tmpReq - CreateMeetingRoomGroupRequest
//
// @param tmpHeader - CreateMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMeetingRoomGroupResponse
func (client *Client) CreateMeetingRoomGroupWithContext(ctx context.Context, tmpReq *CreateMeetingRoomGroupRequest, tmpHeader *CreateMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *CreateMeetingRoomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMeetingRoomGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateMeetingRoomGroupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ParentGroupId) {
		body["ParentGroupId"] = request.ParentGroupId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMeetingRoomGroup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/createMeetingRoomGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMeetingRoomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消息
//
// @param request - CreateMessageRequest
//
// @param headers - CreateMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageResponse
func (client *Client) CreateMessageWithContext(ctx context.Context, request *CreateMessageRequest, headers *CreateMessageHeaders, runtime *dara.RuntimeOptions) (_result *CreateMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMessage"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/createMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建字段
//
// @param tmpReq - CreateMultiDimTableFieldRequest
//
// @param tmpHeader - CreateMultiDimTableFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiDimTableFieldResponse
func (client *Client) CreateMultiDimTableFieldWithContext(ctx context.Context, tmpReq *CreateMultiDimTableFieldRequest, tmpHeader *CreateMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *CreateMultiDimTableFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMultiDimTableFieldShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateMultiDimTableFieldShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Property) {
		request.PropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Property, dara.String("Property"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PropertyShrink) {
		body["Property"] = request.PropertyShrink
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultiDimTableField"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/createMultiDimTableField"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultiDimTableFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或更新表单实例
//
// @param request - CreateOrUpdateFormDataRequest
//
// @param tmpHeader - CreateOrUpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormDataWithContext(ctx context.Context, request *CreateOrUpdateFormDataRequest, tmpHeader *CreateOrUpdateFormDataHeaders, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateFormDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &CreateOrUpdateFormDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormDataJson) {
		body["FormDataJson"] = request.FormDataJson
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.NoExecuteExpression) {
		body["NoExecuteExpression"] = request.NoExecuteExpression
	}

	if !dara.IsNil(request.SearchCondition) {
		body["SearchCondition"] = request.SearchCondition
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateFormData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/createOrUpdateFormData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建荣誉勋章模板
//
// @param tmpReq - CreateOrgHonorTemplateRequest
//
// @param tmpHeader - CreateOrgHonorTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrgHonorTemplateResponse
func (client *Client) CreateOrgHonorTemplateWithContext(ctx context.Context, tmpReq *CreateOrgHonorTemplateRequest, tmpHeader *CreateOrgHonorTemplateHeaders, runtime *dara.RuntimeOptions) (_result *CreateOrgHonorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrgHonorTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateOrgHonorTemplateShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.AvatarFrameMediaId) {
		body["avatarFrameMediaId"] = request.AvatarFrameMediaId
	}

	if !dara.IsNil(request.DefaultBgColor) {
		body["defaultBgColor"] = request.DefaultBgColor
	}

	if !dara.IsNil(request.MedalDesc) {
		body["medalDesc"] = request.MedalDesc
	}

	if !dara.IsNil(request.MedalMediaId) {
		body["medalMediaId"] = request.MedalMediaId
	}

	if !dara.IsNil(request.MedalName) {
		body["medalName"] = request.MedalName
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrgHonorTemplate"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/createOrgHonorTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrgHonorTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建钉钉个人待办任务
//
// @param tmpReq - CreatePersonalTodoTaskRequest
//
// @param tmpHeader - CreatePersonalTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePersonalTodoTaskResponse
func (client *Client) CreatePersonalTodoTaskWithContext(ctx context.Context, tmpReq *CreatePersonalTodoTaskRequest, tmpHeader *CreatePersonalTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreatePersonalTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePersonalTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreatePersonalTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutorIds) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, dara.String("ExecutorIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotifyConfigs) {
		request.NotifyConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyConfigs, dara.String("NotifyConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ParticipantIds) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, dara.String("ParticipantIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DueTime) {
		body["DueTime"] = request.DueTime
	}

	if !dara.IsNil(request.ExecutorIdsShrink) {
		body["ExecutorIds"] = request.ExecutorIdsShrink
	}

	if !dara.IsNil(request.NotifyConfigsShrink) {
		body["NotifyConfigs"] = request.NotifyConfigsShrink
	}

	if !dara.IsNil(request.ParticipantIdsShrink) {
		body["ParticipantIds"] = request.ParticipantIdsShrink
	}

	if !dara.IsNil(request.ReminderTimeStamp) {
		body["ReminderTimeStamp"] = request.ReminderTimeStamp
	}

	if !dara.IsNil(request.Subject) {
		body["Subject"] = request.Subject
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePersonalTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/createPersonalTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePersonalTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建日志
//
// @param tmpReq - CreateReportRequest
//
// @param tmpHeader - CreateReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportResponse
func (client *Client) CreateReportWithContext(ctx context.Context, tmpReq *CreateReportRequest, tmpHeader *CreateReportHeaders, runtime *dara.RuntimeOptions) (_result *CreateReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("Contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ToCids) {
		request.ToCidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToCids, dara.String("ToCids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ToUserids) {
		request.ToUseridsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToUserids, dara.String("ToUserids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentsShrink) {
		body["Contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.DdFrom) {
		body["DdFrom"] = request.DdFrom
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ToChat) {
		body["ToChat"] = request.ToChat
	}

	if !dara.IsNil(request.ToCidsShrink) {
		body["ToCids"] = request.ToCidsShrink
	}

	if !dara.IsNil(request.ToUseridsShrink) {
		body["ToUserids"] = request.ToUseridsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/createReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建运行
//
// @param request - CreateRunRequest
//
// @param headers - CreateRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRunResponse
func (client *Client) CreateRunWithSSECtx(ctx context.Context, request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions, _yield chan *CreateRunResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createRunWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 创建运行
//
// @param request - CreateRunRequest
//
// @param headers - CreateRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRunResponse
func (client *Client) CreateRunWithContext(ctx context.Context, request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions) (_result *CreateRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowStructViewContent) {
		body["allowStructViewContent"] = request.AllowStructViewContent
	}

	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRun"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/createRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景群
//
// @param request - CreateScenegroupRequest
//
// @param tmpHeader - CreateScenegroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScenegroupResponse
func (client *Client) CreateScenegroupWithContext(ctx context.Context, request *CreateScenegroupRequest, tmpHeader *CreateScenegroupHeaders, runtime *dara.RuntimeOptions) (_result *CreateScenegroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &CreateScenegroupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddFriendForbidden) {
		body["AddFriendForbidden"] = request.AddFriendForbidden
	}

	if !dara.IsNil(request.AllMembersCanCreateCalendar) {
		body["AllMembersCanCreateCalendar"] = request.AllMembersCanCreateCalendar
	}

	if !dara.IsNil(request.AllMembersCanCreateMcsConf) {
		body["AllMembersCanCreateMcsConf"] = request.AllMembersCanCreateMcsConf
	}

	if !dara.IsNil(request.ChatBannedType) {
		body["ChatBannedType"] = request.ChatBannedType
	}

	if !dara.IsNil(request.GroupEmailDisabled) {
		body["GroupEmailDisabled"] = request.GroupEmailDisabled
	}

	if !dara.IsNil(request.GroupLiveSwitch) {
		body["GroupLiveSwitch"] = request.GroupLiveSwitch
	}

	if !dara.IsNil(request.Icon) {
		body["Icon"] = request.Icon
	}

	if !dara.IsNil(request.ManagementType) {
		body["ManagementType"] = request.ManagementType
	}

	if !dara.IsNil(request.MembersToAdminChat) {
		body["MembersToAdminChat"] = request.MembersToAdminChat
	}

	if !dara.IsNil(request.MentionAllAuthority) {
		body["MentionAllAuthority"] = request.MentionAllAuthority
	}

	if !dara.IsNil(request.OnlyAdminCanDing) {
		body["OnlyAdminCanDing"] = request.OnlyAdminCanDing
	}

	if !dara.IsNil(request.OnlyAdminCanSetMsgTop) {
		body["OnlyAdminCanSetMsgTop"] = request.OnlyAdminCanSetMsgTop
	}

	if !dara.IsNil(request.Searchable) {
		body["Searchable"] = request.Searchable
	}

	if !dara.IsNil(request.ShowHistoryType) {
		body["ShowHistoryType"] = request.ShowHistoryType
	}

	if !dara.IsNil(request.SubadminIds) {
		body["SubadminIds"] = request.SubadminIds
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.ValidationType) {
		body["ValidationType"] = request.ValidationType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScenegroup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/createScenegroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScenegroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预约会议
//
// @param tmpReq - CreateScheduleConferenceRequest
//
// @param tmpHeader - CreateScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduleConferenceResponse
func (client *Client) CreateScheduleConferenceWithContext(ctx context.Context, tmpReq *CreateScheduleConferenceRequest, tmpHeader *CreateScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CreateScheduleConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScheduleConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateScheduleConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfSettingModel) {
		request.ScheduleConfSettingModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfSettingModel, dara.String("ScheduleConfSettingModel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ScheduleConfSettingModelShrink) {
		body["ScheduleConfSettingModel"] = request.ScheduleConfSettingModelShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduleConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/createScheduleConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduleConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布钉钉搜索穹顶
//
// @param tmpReq - CreateSearchDomeRequest
//
// @param tmpHeader - CreateSearchDomeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchDomeResponse
func (client *Client) CreateSearchDomeWithContext(ctx context.Context, tmpReq *CreateSearchDomeRequest, tmpHeader *CreateSearchDomeHeaders, runtime *dara.RuntimeOptions) (_result *CreateSearchDomeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSearchDomeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSearchDomeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResId) {
		body["ResId"] = request.ResId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSearchDome"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/createSearchDome"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSearchDomeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布钉钉搜索关键词
//
// @param tmpReq - CreateSearchKeywordRequest
//
// @param tmpHeader - CreateSearchKeywordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchKeywordResponse
func (client *Client) CreateSearchKeywordWithContext(ctx context.Context, tmpReq *CreateSearchKeywordRequest, tmpHeader *CreateSearchKeywordHeaders, runtime *dara.RuntimeOptions) (_result *CreateSearchKeywordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSearchKeywordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSearchKeywordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResId) {
		body["ResId"] = request.ResId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSearchKeyword"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/createSearchKeyword"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSearchKeywordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作表
//
// @param tmpReq - CreateSheetRequest
//
// @param tmpHeader - CreateSheetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSheetResponse
func (client *Client) CreateSheetWithContext(ctx context.Context, tmpReq *CreateSheetRequest, tmpHeader *CreateSheetHeaders, runtime *dara.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSheet"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/createSheet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSheetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建订阅日历
//
// @param tmpReq - CreateSubscribedCalendarRequest
//
// @param tmpHeader - CreateSubscribedCalendarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubscribedCalendarResponse
func (client *Client) CreateSubscribedCalendarWithContext(ctx context.Context, tmpReq *CreateSubscribedCalendarRequest, tmpHeader *CreateSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *CreateSubscribedCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSubscribedCalendarShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSubscribedCalendarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubscribeScope) {
		request.SubscribeScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubscribeScope, dara.String("SubscribeScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ManagersShrink) {
		body["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SubscribeScopeShrink) {
		body["SubscribeScope"] = request.SubscribeScopeShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubscribedCalendar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/createSubscribedCalendar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubscribedCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建线程
//
// @param request - CreateThreadRequest
//
// @param headers - CreateThreadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateThreadResponse
func (client *Client) CreateThreadWithContext(ctx context.Context, request *CreateThreadRequest, headers *CreateThreadHeaders, runtime *dara.RuntimeOptions) (_result *CreateThreadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ClientEnum) {
		body["clientEnum"] = request.ClientEnum
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateThread"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/createThread"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateThreadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工单
//
// @param tmpReq - CreateTicketRequest
//
// @param tmpHeader - CreateTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, tmpReq *CreateTicketRequest, tmpHeader *CreateTicketHeaders, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateTicketShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notify) {
		request.NotifyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notify, dara.String("Notify"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProcessorUserIds) {
		request.ProcessorUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProcessorUserIds, dara.String("ProcessorUserIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SceneContext) {
		request.SceneContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneContext, dara.String("SceneContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomFields) {
		body["CustomFields"] = request.CustomFields
	}

	if !dara.IsNil(request.NotifyShrink) {
		body["Notify"] = request.NotifyShrink
	}

	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTemplateBizId) {
		body["OpenTemplateBizId"] = request.OpenTemplateBizId
	}

	if !dara.IsNil(request.ProcessorUserIdsShrink) {
		body["ProcessorUserIds"] = request.ProcessorUserIdsShrink
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SceneContextShrink) {
		body["SceneContext"] = request.SceneContextShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/createTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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
// 创建代办
//
// @param tmpReq - CreateTodoTaskRequest
//
// @param tmpHeader - CreateTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTodoTaskResponse
func (client *Client) CreateTodoTaskWithContext(ctx context.Context, tmpReq *CreateTodoTaskRequest, tmpHeader *CreateTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ActionList) {
		request.ActionListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionList, dara.String("actionList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ContentFieldList) {
		request.ContentFieldListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContentFieldList, dara.String("contentFieldList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DetailUrl) {
		request.DetailUrlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetailUrl, dara.String("detailUrl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutorIds) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, dara.String("executorIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotifyConfigs) {
		request.NotifyConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyConfigs, dara.String("notifyConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ParticipantIds) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, dara.String("participantIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RemindNotifyConfigs) {
		request.RemindNotifyConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemindNotifyConfigs, dara.String("remindNotifyConfigs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperatorId) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ActionListShrink) {
		body["actionList"] = request.ActionListShrink
	}

	if !dara.IsNil(request.ContentFieldListShrink) {
		body["contentFieldList"] = request.ContentFieldListShrink
	}

	if !dara.IsNil(request.CreatorId) {
		body["creatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DetailUrlShrink) {
		body["detailUrl"] = request.DetailUrlShrink
	}

	if !dara.IsNil(request.DueTime) {
		body["dueTime"] = request.DueTime
	}

	if !dara.IsNil(request.ExecutorIdsShrink) {
		body["executorIds"] = request.ExecutorIdsShrink
	}

	if !dara.IsNil(request.IsOnlyShowExecutor) {
		body["isOnlyShowExecutor"] = request.IsOnlyShowExecutor
	}

	if !dara.IsNil(request.NotifyConfigsShrink) {
		body["notifyConfigs"] = request.NotifyConfigsShrink
	}

	if !dara.IsNil(request.ParticipantIdsShrink) {
		body["participantIds"] = request.ParticipantIdsShrink
	}

	if !dara.IsNil(request.Priority) {
		body["priority"] = request.Priority
	}

	if !dara.IsNil(request.RemindNotifyConfigsShrink) {
		body["remindNotifyConfigs"] = request.RemindNotifyConfigsShrink
	}

	if !dara.IsNil(request.ReminderTimeStamp) {
		body["reminderTimeStamp"] = request.ReminderTimeStamp
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/createTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建视频会议
//
// @param tmpReq - CreateVideoConferenceRequest
//
// @param tmpHeader - CreateVideoConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoConferenceResponse
func (client *Client) CreateVideoConferenceWithContext(ctx context.Context, tmpReq *CreateVideoConferenceRequest, tmpHeader *CreateVideoConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CreateVideoConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVideoConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateVideoConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InviteUserIds) {
		request.InviteUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InviteUserIds, dara.String("InviteUserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfTitle) {
		body["ConfTitle"] = request.ConfTitle
	}

	if !dara.IsNil(request.InviteCaller) {
		body["InviteCaller"] = request.InviteCaller
	}

	if !dara.IsNil(request.InviteUserIdsShrink) {
		body["InviteUserIds"] = request.InviteUserIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/createVideoConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建知识库
//
// @param tmpReq - CreateWorkspaceRequest
//
// @param tmpHeader - CreateWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithContext(ctx context.Context, tmpReq *CreateWorkspaceRequest, tmpHeader *CreateWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateWorkspaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/createWorkspace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建知识库文档
//
// @param tmpReq - CreateWorkspaceDocRequest
//
// @param tmpHeader - CreateWorkspaceDocHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceDocResponse
func (client *Client) CreateWorkspaceDocWithContext(ctx context.Context, tmpReq *CreateWorkspaceDocRequest, tmpHeader *CreateWorkspaceDocHeaders, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkspaceDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateWorkspaceDocShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentNodeId) {
		body["ParentNodeId"] = request.ParentNodeId
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspaceDoc"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/createWorkspaceDoc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - DeleteAlidingAssistantRequest
//
// @param tmpHeader - DeleteAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlidingAssistantResponse
func (client *Client) DeleteAlidingAssistantWithContext(ctx context.Context, tmpReq *DeleteAlidingAssistantRequest, tmpHeader *DeleteAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *DeleteAlidingAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAlidingAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteAlidingAssistantShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["AssistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlidingAssistant"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/aiagent/deleteAlidingAssistant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlidingAssistantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定列
//
// @param tmpReq - DeleteColumnsRequest
//
// @param tmpHeader - DeleteColumnsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteColumnsResponse
func (client *Client) DeleteColumnsWithContext(ctx context.Context, tmpReq *DeleteColumnsRequest, tmpHeader *DeleteColumnsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteColumnsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteColumnsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Column) {
		body["Column"] = request.Column
	}

	if !dara.IsNil(request.ColumnCount) {
		body["ColumnCount"] = request.ColumnCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteColumns"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteColumns"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteColumnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件或文件夹
//
// @param tmpReq - DeleteDentryRequest
//
// @param tmpHeader - DeleteDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDentryResponse
func (client *Client) DeleteDentryWithContext(ctx context.Context, tmpReq *DeleteDentryRequest, tmpHeader *DeleteDentryHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDentryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDentryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteDentryShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ToRecycleBin) {
		body["ToRecycleBin"] = request.ToRecycleBin
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDentry"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteDentry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDentryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除钉盘空间
//
// @param tmpReq - DeleteDriveSpaceRequest
//
// @param tmpHeader - DeleteDriveSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDriveSpaceResponse
func (client *Client) DeleteDriveSpaceWithContext(ctx context.Context, tmpReq *DeleteDriveSpaceRequest, tmpHeader *DeleteDriveSpaceHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDriveSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDriveSpaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteDriveSpaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDriveSpace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteDriveSpace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDriveSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除日程
//
// @param request - DeleteEventRequest
//
// @param tmpHeader - DeleteEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventResponse
func (client *Client) DeleteEventWithContext(ctx context.Context, request *DeleteEventRequest, tmpHeader *DeleteEventHeaders, runtime *dara.RuntimeOptions) (_result *DeleteEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &DeleteEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.PushNotification) {
		body["pushNotification"] = request.PushNotification
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/deleteEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除表单数据
//
// @param request - DeleteFormDataRequest
//
// @param tmpHeader - DeleteFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFormDataResponse
func (client *Client) DeleteFormDataWithContext(ctx context.Context, request *DeleteFormDataRequest, tmpHeader *DeleteFormDataHeaders, runtime *dara.RuntimeOptions) (_result *DeleteFormDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &DeleteFormDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormInstanceId) {
		body["FormInstanceId"] = request.FormInstanceId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFormData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/deleteFormData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFormDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param tmpHeader - DeleteInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, tmpHeader *DeleteInstanceHeaders, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &DeleteInstanceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/deleteInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除直播
//
// @param tmpReq - DeleteLiveRequest
//
// @param tmpHeader - DeleteLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveResponse
func (client *Client) DeleteLiveWithContext(ctx context.Context, tmpReq *DeleteLiveRequest, tmpHeader *DeleteLiveHeaders, runtime *dara.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteLiveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteLiveShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLive"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/deleteLive"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会议室
//
// @param tmpReq - DeleteMeetingRoomRequest
//
// @param tmpHeader - DeleteMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMeetingRoomResponse
func (client *Client) DeleteMeetingRoomWithContext(ctx context.Context, tmpReq *DeleteMeetingRoomRequest, tmpHeader *DeleteMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMeetingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMeetingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteMeetingRoomShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoomId) {
		body["RoomId"] = request.RoomId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMeetingRoom"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/deleteMeetingRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMeetingRoomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会议室分组
//
// @param tmpReq - DeleteMeetingRoomGroupRequest
//
// @param tmpHeader - DeleteMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMeetingRoomGroupResponse
func (client *Client) DeleteMeetingRoomGroupWithContext(ctx context.Context, tmpReq *DeleteMeetingRoomGroupRequest, tmpHeader *DeleteMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMeetingRoomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMeetingRoomGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteMeetingRoomGroupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMeetingRoomGroup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/deleteMeetingRoomGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMeetingRoomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除字段
//
// @param tmpReq - DeleteMultiDimTableFieldRequest
//
// @param tmpHeader - DeleteMultiDimTableFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiDimTableFieldResponse
func (client *Client) DeleteMultiDimTableFieldWithContext(ctx context.Context, tmpReq *DeleteMultiDimTableFieldRequest, tmpHeader *DeleteMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMultiDimTableFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMultiDimTableFieldShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteMultiDimTableFieldShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.FieldIdOrName) {
		body["FieldIdOrName"] = request.FieldIdOrName
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultiDimTableField"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/deleteMultiDimTableField"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultiDimTableFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除多行记录
//
// @param tmpReq - DeleteMultiDimTableRecordsRequest
//
// @param tmpHeader - DeleteMultiDimTableRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiDimTableRecordsResponse
func (client *Client) DeleteMultiDimTableRecordsWithContext(ctx context.Context, tmpReq *DeleteMultiDimTableRecordsRequest, tmpHeader *DeleteMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMultiDimTableRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMultiDimTableRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteMultiDimTableRecordsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecordIds) {
		request.RecordIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordIds, dara.String("RecordIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.RecordIdsShrink) {
		body["RecordIds"] = request.RecordIdsShrink
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultiDimTableRecords"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/deleteMultiDimTableRecords"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultiDimTableRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件权限
//
// @param tmpReq - DeletePermissionRequest
//
// @param tmpHeader - DeletePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePermissionResponse
func (client *Client) DeletePermissionWithContext(ctx context.Context, tmpReq *DeletePermissionRequest, tmpHeader *DeletePermissionHeaders, runtime *dara.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeletePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeletePermissionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.RoleId) {
		body["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePermission"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deletePermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定行
//
// @param tmpReq - DeleteRowsRequest
//
// @param tmpHeader - DeleteRowsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRowsResponse
func (client *Client) DeleteRowsWithContext(ctx context.Context, tmpReq *DeleteRowsRequest, tmpHeader *DeleteRowsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteRowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteRowsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Row) {
		body["Row"] = request.Row
	}

	if !dara.IsNil(request.RowCount) {
		body["RowCount"] = request.RowCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRows"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteRows"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除群成员
//
// @param request - DeleteScenegroupMemberRequest
//
// @param tmpHeader - DeleteScenegroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScenegroupMemberResponse
func (client *Client) DeleteScenegroupMemberWithContext(ctx context.Context, request *DeleteScenegroupMemberRequest, tmpHeader *DeleteScenegroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *DeleteScenegroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &DeleteScenegroupMemberShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	if !dara.IsNil(request.UserIds) {
		body["UserIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScenegroupMember"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/deleteScenegroupMember"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScenegroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个工作表
//
// @param tmpReq - DeleteSheetRequest
//
// @param tmpHeader - DeleteSheetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSheetResponse
func (client *Client) DeleteSheetWithContext(ctx context.Context, tmpReq *DeleteSheetRequest, tmpHeader *DeleteSheetHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSheetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSheet"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteSheet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSheetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除订阅日历
//
// @param request - DeleteSubscribedCalendarRequest
//
// @param tmpHeader - DeleteSubscribedCalendarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubscribedCalendarResponse
func (client *Client) DeleteSubscribedCalendarWithContext(ctx context.Context, request *DeleteSubscribedCalendarRequest, tmpHeader *DeleteSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSubscribedCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &DeleteSubscribedCalendarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubscribedCalendar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/deleteSubscribedCalendar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubscribedCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代办
//
// @param tmpReq - DeleteTodoTaskRequest
//
// @param tmpHeader - DeleteTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTodoTaskResponse
func (client *Client) DeleteTodoTaskWithContext(ctx context.Context, tmpReq *DeleteTodoTaskRequest, tmpHeader *DeleteTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.OperatorId) {
		body["operatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/deleteTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除知识库文档成员
//
// @param tmpReq - DeleteWorkspaceDocMembersRequest
//
// @param tmpHeader - DeleteWorkspaceDocMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceDocMembersResponse
func (client *Client) DeleteWorkspaceDocMembersWithContext(ctx context.Context, tmpReq *DeleteWorkspaceDocMembersRequest, tmpHeader *DeleteWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspaceDocMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteWorkspaceDocMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除知识库成员
//
// @param tmpReq - DeleteWorkspaceMembersRequest
//
// @param tmpHeader - DeleteWorkspaceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceMembersResponse
func (client *Client) DeleteWorkspaceMembersWithContext(ctx context.Context, tmpReq *DeleteWorkspaceMembersRequest, tmpHeader *DeleteWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspaceMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/deleteWorkspaceMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除块元素
//
// @param tmpReq - DocBlocksDeleteRequest
//
// @param tmpHeader - DocBlocksDeleteHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocBlocksDeleteResponse
func (client *Client) DocBlocksDeleteWithContext(ctx context.Context, tmpReq *DocBlocksDeleteRequest, tmpHeader *DocBlocksDeleteHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DocBlocksDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DocBlocksDeleteShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockId) {
		body["BlockId"] = request.BlockId
	}

	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocBlocksDelete"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/docBlocksDelete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocBlocksDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新块元素
//
// @param tmpReq - DocBlocksModifyRequest
//
// @param tmpHeader - DocBlocksModifyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocBlocksModifyResponse
func (client *Client) DocBlocksModifyWithContext(ctx context.Context, tmpReq *DocBlocksModifyRequest, tmpHeader *DocBlocksModifyHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksModifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DocBlocksModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DocBlocksModifyShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Element) {
		request.ElementShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Element, dara.String("Element"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockId) {
		body["BlockId"] = request.BlockId
	}

	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.ElementShrink) {
		body["Element"] = request.ElementShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocBlocksModify"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/docBlocksModify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocBlocksModifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询块元素
//
// @param tmpReq - DocBlocksQueryRequest
//
// @param tmpHeader - DocBlocksQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocBlocksQueryResponse
func (client *Client) DocBlocksQueryWithContext(ctx context.Context, tmpReq *DocBlocksQueryRequest, tmpHeader *DocBlocksQueryHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DocBlocksQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DocBlocksQueryShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockType) {
		body["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.DocKey) {
		body["DocKey"] = request.DocKey
	}

	if !dara.IsNil(request.EndIndex) {
		body["EndIndex"] = request.EndIndex
	}

	if !dara.IsNil(request.StartIndex) {
		body["StartIndex"] = request.StartIndex
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocBlocksQuery"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/docBlocksQuery"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocBlocksQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 覆写文档
//
// @param tmpReq - DocUpdateContentRequest
//
// @param tmpHeader - DocUpdateContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocUpdateContentResponse
func (client *Client) DocUpdateContentWithContext(ctx context.Context, tmpReq *DocUpdateContentRequest, tmpHeader *DocUpdateContentHeaders, runtime *dara.RuntimeOptions) (_result *DocUpdateContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DocUpdateContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DocUpdateContentShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.DocKey) {
		body["DocKey"] = request.DocKey
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocUpdateContent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/docUpdateContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocUpdateContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量执行宜搭审批任务
//
// @param request - ExecuteBatchTaskRequest
//
// @param tmpHeader - ExecuteBatchTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteBatchTaskResponse
func (client *Client) ExecuteBatchTaskWithContext(ctx context.Context, request *ExecuteBatchTaskRequest, tmpHeader *ExecuteBatchTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ExecuteBatchTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.OutResult) {
		body["OutResult"] = request.OutResult
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TaskInformationList) {
		body["TaskInformationList"] = request.TaskInformationList
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteBatchTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/executeBatchTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行宜搭的审批任务
//
// @param request - ExecutePlatformTaskRequest
//
// @param tmpHeader - ExecutePlatformTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecutePlatformTaskResponse
func (client *Client) ExecutePlatformTaskWithContext(ctx context.Context, request *ExecutePlatformTaskRequest, tmpHeader *ExecutePlatformTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecutePlatformTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ExecutePlatformTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormDataJson) {
		body["FormDataJson"] = request.FormDataJson
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.NoExecuteExpressions) {
		body["NoExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !dara.IsNil(request.OutResult) {
		body["OutResult"] = request.OutResult
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecutePlatformTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/executePlatformTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecutePlatformTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同意或拒绝宜搭审批任务(执行审批任务)
//
// @param request - ExecuteTaskRequest
//
// @param tmpHeader - ExecuteTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTaskResponse
func (client *Client) ExecuteTaskWithContext(ctx context.Context, request *ExecuteTaskRequest, tmpHeader *ExecuteTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ExecuteTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DigitalSignUrl) {
		body["DigitalSignUrl"] = request.DigitalSignUrl
	}

	if !dara.IsNil(request.FormDataJson) {
		body["FormDataJson"] = request.FormDataJson
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.NoExecuteExpressions) {
		body["NoExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !dara.IsNil(request.OutResult) {
		body["OutResult"] = request.OutResult
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/executeTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群扩容
//
// @param tmpReq - ExpandGroupCapacityRequest
//
// @param tmpHeader - ExpandGroupCapacityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandGroupCapacityResponse
func (client *Client) ExpandGroupCapacityWithContext(ctx context.Context, tmpReq *ExpandGroupCapacityRequest, tmpHeader *ExpandGroupCapacityHeaders, runtime *dara.RuntimeOptions) (_result *ExpandGroupCapacityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExpandGroupCapacityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ExpandGroupCapacityShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpandGroupCapacity"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/im/expandGroupCapacity"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpandGroupCapacityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完结工单
//
// @param tmpReq - FinishTicketRequest
//
// @param tmpHeader - FinishTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishTicketResponse
func (client *Client) FinishTicketWithContext(ctx context.Context, tmpReq *FinishTicketRequest, tmpHeader *FinishTicketHeaders, runtime *dara.RuntimeOptions) (_result *FinishTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FinishTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &FinishTicketShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notify) {
		request.NotifyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notify, dara.String("Notify"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketMemo) {
		request.TicketMemoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketMemo, dara.String("TicketMemo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NotifyShrink) {
		body["Notify"] = request.NotifyShrink
	}

	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.TicketMemoShrink) {
		body["TicketMemo"] = request.TicketMemoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishTicket"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/finishTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程设计的节点信息
//
// @param request - GetActivityListRequest
//
// @param tmpHeader - GetActivityListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActivityListResponse
func (client *Client) GetActivityListWithContext(ctx context.Context, request *GetActivityListRequest, tmpHeader *GetActivityListHeaders, runtime *dara.RuntimeOptions) (_result *GetActivityListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetActivityListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessCode) {
		body["ProcessCode"] = request.ProcessCode
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetActivityList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getActivityList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetActivityListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取阿里钉ai助理信息
//
// @param tmpReq - GetAlidingAssistantInfoRequest
//
// @param tmpHeader - GetAlidingAssistantInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlidingAssistantInfoResponse
func (client *Client) GetAlidingAssistantInfoWithContext(ctx context.Context, tmpReq *GetAlidingAssistantInfoRequest, tmpHeader *GetAlidingAssistantInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetAlidingAssistantInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAlidingAssistantInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetAlidingAssistantInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["AssistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlidingAssistantInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/aiagent/getAlidingAssistantInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlidingAssistantInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有工作表
//
// @param tmpReq - GetAllSheetsRequest
//
// @param tmpHeader - GetAllSheetsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllSheetsResponse
func (client *Client) GetAllSheetsWithContext(ctx context.Context, tmpReq *GetAllSheetsRequest, tmpHeader *GetAllSheetsHeaders, runtime *dara.RuntimeOptions) (_result *GetAllSheetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAllSheetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetAllSheetsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAllSheets"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getAllSheets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAllSheetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取助理能力
//
// @param request - GetAssistantCapabilityRequest
//
// @param headers - GetAssistantCapabilityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssistantCapabilityResponse
func (client *Client) GetAssistantCapabilityWithContext(ctx context.Context, request *GetAssistantCapabilityRequest, headers *GetAssistantCapabilityHeaders, runtime *dara.RuntimeOptions) (_result *GetAssistantCapabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.Timeout) {
		body["timeout"] = request.Timeout
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssistantCapability"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/getAssistantCapability"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssistantCapabilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - GetCardTemplateRequest
//
// @param tmpHeader - GetCardTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardTemplateResponse
func (client *Client) GetCardTemplateWithContext(ctx context.Context, tmpReq *GetCardTemplateRequest, tmpHeader *GetCardTemplateHeaders, runtime *dara.RuntimeOptions) (_result *GetCardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCardTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetCardTemplateShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCardTemplate"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/card/getCardTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCardTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取群存储空间信息
//
// @param tmpReq - GetConversaionSpaceRequest
//
// @param tmpHeader - GetConversaionSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversaionSpaceResponse
func (client *Client) GetConversaionSpaceWithContext(ctx context.Context, tmpReq *GetConversaionSpaceRequest, tmpHeader *GetConversaionSpaceHeaders, runtime *dara.RuntimeOptions) (_result *GetConversaionSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetConversaionSpaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetConversaionSpaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConversaionSpace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getConversaionSpace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConversaionSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织内已完成的审批任务
//
// @param request - GetCorpAccomplishmentTasksRequest
//
// @param tmpHeader - GetCorpAccomplishmentTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCorpAccomplishmentTasksResponse
func (client *Client) GetCorpAccomplishmentTasksWithContext(ctx context.Context, request *GetCorpAccomplishmentTasksRequest, tmpHeader *GetCorpAccomplishmentTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetCorpAccomplishmentTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetCorpAccomplishmentTasksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppTypes) {
		body["AppTypes"] = request.AppTypes
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCorpAccomplishmentTasks"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getCorpAccomplishmentTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCorpAccomplishmentTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务列表（组织维度）
//
// @param request - GetCorpTasksRequest
//
// @param tmpHeader - GetCorpTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCorpTasksResponse
func (client *Client) GetCorpTasksWithContext(ctx context.Context, request *GetCorpTasksRequest, tmpHeader *GetCorpTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetCorpTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetCorpTasksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppTypes) {
		body["AppTypes"] = request.AppTypes
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCorpTasks"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getCorpTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCorpTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - GetDeptNoRequest
//
// @param tmpHeader - GetDeptNoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeptNoResponse
func (client *Client) GetDeptNoWithContext(ctx context.Context, tmpReq *GetDeptNoRequest, tmpHeader *GetDeptNoHeaders, runtime *dara.RuntimeOptions) (_result *GetDeptNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeptNoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDeptNoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.DeptId) {
		body["deptId"] = request.DeptId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeptNo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/dept/getDeptNo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeptNoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉会议信息
//
// @param tmpReq - GetDingtalkMeetingInfoRequest
//
// @param tmpHeader - GetDingtalkMeetingInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkMeetingInfoResponse
func (client *Client) GetDingtalkMeetingInfoWithContext(ctx context.Context, tmpReq *GetDingtalkMeetingInfoRequest, tmpHeader *GetDingtalkMeetingInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkMeetingInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkMeetingInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkMeetingInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkMeetingInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkMeetingInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkMeetingInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉会议列表
//
// @param tmpReq - GetDingtalkMeetingListRequest
//
// @param tmpHeader - GetDingtalkMeetingListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkMeetingListResponse
func (client *Client) GetDingtalkMeetingListWithContext(ctx context.Context, tmpReq *GetDingtalkMeetingListRequest, tmpHeader *GetDingtalkMeetingListHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkMeetingListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkMeetingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkMeetingListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoomCode) {
		body["roomCode"] = request.RoomCode
	}

	if !dara.IsNil(request.RoomName) {
		body["roomName"] = request.RoomName
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.WorkNo) {
		body["workNo"] = request.WorkNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkMeetingList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkMeetingList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkMeetingListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉会议成员事件
//
// @param tmpReq - GetDingtalkMeetingMemberEventRequest
//
// @param tmpHeader - GetDingtalkMeetingMemberEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkMeetingMemberEventResponse
func (client *Client) GetDingtalkMeetingMemberEventWithContext(ctx context.Context, tmpReq *GetDingtalkMeetingMemberEventRequest, tmpHeader *GetDingtalkMeetingMemberEventHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkMeetingMemberEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkMeetingMemberEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkMeetingMemberEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.BeginTime) {
		body["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.WorkNo) {
		body["workNo"] = request.WorkNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkMeetingMemberEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkMeetingMemberEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkMeetingMemberEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉会议成员列表
//
// @param tmpReq - GetDingtalkMeetingMemberListRequest
//
// @param tmpHeader - GetDingtalkMeetingMemberListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkMeetingMemberListResponse
func (client *Client) GetDingtalkMeetingMemberListWithContext(ctx context.Context, tmpReq *GetDingtalkMeetingMemberListRequest, tmpHeader *GetDingtalkMeetingMemberListHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkMeetingMemberListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkMeetingMemberListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkMeetingMemberListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ClusterName) {
		body["clusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkMeetingMemberList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkMeetingMemberList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkMeetingMemberListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会议指标数据
//
// @param tmpReq - GetDingtalkMeetingMetricDataRequest
//
// @param tmpHeader - GetDingtalkMeetingMetricDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkMeetingMetricDataResponse
func (client *Client) GetDingtalkMeetingMetricDataWithContext(ctx context.Context, tmpReq *GetDingtalkMeetingMetricDataRequest, tmpHeader *GetDingtalkMeetingMetricDataHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkMeetingMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkMeetingMetricDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkMeetingMetricDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.BeginTime) {
		body["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.TypeName) {
		body["typeName"] = request.TypeName
	}

	if !dara.IsNil(request.WorkNo) {
		body["workNo"] = request.WorkNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkMeetingMetricData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkMeetingMetricData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkMeetingMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉投屏信息
//
// @param tmpReq - GetDingtalkProjectionInfoRequest
//
// @param tmpHeader - GetDingtalkProjectionInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkProjectionInfoResponse
func (client *Client) GetDingtalkProjectionInfoWithContext(ctx context.Context, tmpReq *GetDingtalkProjectionInfoRequest, tmpHeader *GetDingtalkProjectionInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkProjectionInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkProjectionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkProjectionInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Client) {
		body["client"] = request.Client
	}

	if !dara.IsNil(request.EndTs) {
		body["endTs"] = request.EndTs
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.PubWorkNo) {
		body["pubWorkNo"] = request.PubWorkNo
	}

	if !dara.IsNil(request.RoomId) {
		body["roomId"] = request.RoomId
	}

	if !dara.IsNil(request.StartTs) {
		body["startTs"] = request.StartTs
	}

	if !dara.IsNil(request.SubUid) {
		body["subUid"] = request.SubUid
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkProjectionInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkProjectionInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkProjectionInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉钉投屏列表
//
// @param tmpReq - GetDingtalkProjectionListRequest
//
// @param tmpHeader - GetDingtalkProjectionListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingtalkProjectionListResponse
func (client *Client) GetDingtalkProjectionListWithContext(ctx context.Context, tmpReq *GetDingtalkProjectionListRequest, tmpHeader *GetDingtalkProjectionListHeaders, runtime *dara.RuntimeOptions) (_result *GetDingtalkProjectionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDingtalkProjectionListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDingtalkProjectionListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectorWorkNo) {
		body["projectorWorkNo"] = request.ProjectorWorkNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingtalkProjectionList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getDingtalkProjectionList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingtalkProjectionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容
//
// @param tmpReq - GetDocContentRequest
//
// @param tmpHeader - GetDocContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocContentResponse
func (client *Client) GetDocContentWithContext(ctx context.Context, tmpReq *GetDocContentRequest, tmpHeader *GetDocContentHeaders, runtime *dara.RuntimeOptions) (_result *GetDocContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDocContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDocContentShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.TargetFormat) {
		body["TargetFormat"] = request.TargetFormat
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserToken) {
		body["userToken"] = request.UserToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocContent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getDocContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取文档内容taskId
//
// @param tmpReq - GetDocContentTakIdRequest
//
// @param tmpHeader - GetDocContentTakIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocContentTakIdResponse
func (client *Client) GetDocContentTakIdWithContext(ctx context.Context, tmpReq *GetDocContentTakIdRequest, tmpHeader *GetDocContentTakIdHeaders, runtime *dara.RuntimeOptions) (_result *GetDocContentTakIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDocContentTakIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetDocContentTakIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.GenerateCp) {
		body["GenerateCp"] = request.GenerateCp
	}

	if !dara.IsNil(request.TargetFormat) {
		body["TargetFormat"] = request.TargetFormat
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocContentTakId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getDocContentTakId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocContentTakIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个日程详情
//
// @param request - GetEventRequest
//
// @param tmpHeader - GetEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventResponse
func (client *Client) GetEventWithContext(ctx context.Context, request *GetEventRequest, tmpHeader *GetEventHeaders, runtime *dara.RuntimeOptions) (_result *GetEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxAttendees) {
		query["MaxAttendees"] = request.MaxAttendees
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/getEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单内的组件信息
//
// @param request - GetFieldDefByUuidRequest
//
// @param tmpHeader - GetFieldDefByUuidHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFieldDefByUuidResponse
func (client *Client) GetFieldDefByUuidWithContext(ctx context.Context, request *GetFieldDefByUuidRequest, tmpHeader *GetFieldDefByUuidHeaders, runtime *dara.RuntimeOptions) (_result *GetFieldDefByUuidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetFieldDefByUuidShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFieldDefByUuid"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getFieldDefByUuid"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFieldDefByUuidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件下载信息
//
// @param tmpReq - GetFileDownloadInfoRequest
//
// @param tmpHeader - GetFileDownloadInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileDownloadInfoResponse
func (client *Client) GetFileDownloadInfoWithContext(ctx context.Context, tmpReq *GetFileDownloadInfoRequest, tmpHeader *GetFileDownloadInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetFileDownloadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetFileDownloadInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetFileDownloadInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileDownloadInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getFileDownloadInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件上传信息
//
// @param tmpReq - GetFileUploadInfoRequest
//
// @param tmpHeader - GetFileUploadInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileUploadInfoResponse
func (client *Client) GetFileUploadInfoWithContext(ctx context.Context, tmpReq *GetFileUploadInfoRequest, tmpHeader *GetFileUploadInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetFileUploadInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetFileUploadInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.ParentDentryUuid) {
		body["ParentDentryUuid"] = request.ParentDentryUuid
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileUploadInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getFileUploadInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单组件定义列表
//
// @param request - GetFormComponentDefinitionListRequest
//
// @param tmpHeader - GetFormComponentDefinitionListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormComponentDefinitionListResponse
func (client *Client) GetFormComponentDefinitionListWithContext(ctx context.Context, request *GetFormComponentDefinitionListRequest, tmpHeader *GetFormComponentDefinitionListHeaders, runtime *dara.RuntimeOptions) (_result *GetFormComponentDefinitionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetFormComponentDefinitionListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFormComponentDefinitionList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getFormComponentDefinitionList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFormComponentDefinitionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表单数据
//
// @param request - GetFormDataByIDRequest
//
// @param tmpHeader - GetFormDataByIDHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormDataByIDResponse
func (client *Client) GetFormDataByIDWithContext(ctx context.Context, request *GetFormDataByIDRequest, tmpHeader *GetFormDataByIDHeaders, runtime *dara.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetFormDataByIDShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFormDataByID"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getFormDataByID"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定应用下的表单列表
//
// @param request - GetFormListInAppRequest
//
// @param tmpHeader - GetFormListInAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormListInAppResponse
func (client *Client) GetFormListInAppWithContext(ctx context.Context, request *GetFormListInAppRequest, tmpHeader *GetFormListInAppHeaders, runtime *dara.RuntimeOptions) (_result *GetFormListInAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetFormListInAppShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormTypes) {
		body["FormTypes"] = request.FormTypes
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFormListInApp"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getFormListInApp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFormListInAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群内直播信息(最早支持2024年01月数据)
//
// @param tmpReq - GetGroupLiveListRequest
//
// @param tmpHeader - GetGroupLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupLiveListResponse
func (client *Client) GetGroupLiveListWithContext(ctx context.Context, tmpReq *GetGroupLiveListRequest, tmpHeader *GetGroupLiveListHeaders, runtime *dara.RuntimeOptions) (_result *GetGroupLiveListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetGroupLiveListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetGroupLiveListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroupLiveList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getGroupLiveList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupLiveListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业内部群成员
//
// @param request - GetInnerGroupMembersRequest
//
// @param tmpHeader - GetInnerGroupMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInnerGroupMembersResponse
func (client *Client) GetInnerGroupMembersWithContext(ctx context.Context, request *GetInnerGroupMembersRequest, tmpHeader *GetInnerGroupMembersHeaders, runtime *dara.RuntimeOptions) (_result *GetInnerGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetInnerGroupMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OpenConversationId) {
		body["OpenConversationId"] = request.OpenConversationId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInnerGroupMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/getInnerGroupMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInnerGroupMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据流程实例ID获取流程实例
//
// @param request - GetInstanceByIdRequest
//
// @param tmpHeader - GetInstanceByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceByIdWithContext(ctx context.Context, request *GetInstanceByIdRequest, tmpHeader *GetInstanceByIdHeaders, runtime *dara.RuntimeOptions) (_result *GetInstanceByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetInstanceByIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceById"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getInstanceById"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例ID列表
//
// @param request - GetInstanceIdListRequest
//
// @param tmpHeader - GetInstanceIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdListWithContext(ctx context.Context, request *GetInstanceIdListRequest, tmpHeader *GetInstanceIdListHeaders, runtime *dara.RuntimeOptions) (_result *GetInstanceIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetInstanceIdListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ApprovedResult) {
		body["ApprovedResult"] = request.ApprovedResult
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.InstanceStatus) {
		body["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchFieldJson) {
		body["SearchFieldJson"] = request.SearchFieldJson
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceIdList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getInstanceIdList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程实例
//
// @param request - GetInstancesRequest
//
// @param tmpHeader - GetInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancesResponse
func (client *Client) GetInstancesWithContext(ctx context.Context, request *GetInstancesRequest, tmpHeader *GetInstancesHeaders, runtime *dara.RuntimeOptions) (_result *GetInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetInstancesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ApprovedResult) {
		body["ApprovedResult"] = request.ApprovedResult
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.InstanceStatus) {
		body["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OrderConfigJson) {
		body["OrderConfigJson"] = request.OrderConfigJson
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchFieldJson) {
		body["SearchFieldJson"] = request.SearchFieldJson
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstances"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getInstances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据实例 ID 列表批量获取流程实例详情(批量获取流程实例列表)
//
// @param request - GetInstancesByIdListRequest
//
// @param tmpHeader - GetInstancesByIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancesByIdListResponse
func (client *Client) GetInstancesByIdListWithContext(ctx context.Context, request *GetInstancesByIdListRequest, tmpHeader *GetInstancesByIdListHeaders, runtime *dara.RuntimeOptions) (_result *GetInstancesByIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetInstancesByIdListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessInstanceIds) {
		body["ProcessInstanceIds"] = request.ProcessInstanceIds
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstancesByIdList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getInstancesByIdList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstancesByIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播的可下载回放地址
//
// @param tmpReq - GetLiveReplayUrlRequest
//
// @param tmpHeader - GetLiveReplayUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveReplayUrlResponse
func (client *Client) GetLiveReplayUrlWithContext(ctx context.Context, tmpReq *GetLiveReplayUrlRequest, tmpHeader *GetLiveReplayUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetLiveReplayUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetLiveReplayUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetLiveReplayUrlShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveReplayUrl"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getLiveReplayUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveReplayUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织内某人提交的任务
//
// @param request - GetMeCorpSubmissionRequest
//
// @param tmpHeader - GetMeCorpSubmissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMeCorpSubmissionResponse
func (client *Client) GetMeCorpSubmissionWithContext(ctx context.Context, request *GetMeCorpSubmissionRequest, tmpHeader *GetMeCorpSubmissionHeaders, runtime *dara.RuntimeOptions) (_result *GetMeCorpSubmissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetMeCorpSubmissionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppTypes) {
		body["AppTypes"] = request.AppTypes
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMeCorpSubmission"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getMeCorpSubmission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMeCorpSubmissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会议室忙闲信息
//
// @param tmpReq - GetMeetingRoomsScheduleRequest
//
// @param tmpHeader - GetMeetingRoomsScheduleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMeetingRoomsScheduleResponse
func (client *Client) GetMeetingRoomsScheduleWithContext(ctx context.Context, tmpReq *GetMeetingRoomsScheduleRequest, tmpHeader *GetMeetingRoomsScheduleHeaders, runtime *dara.RuntimeOptions) (_result *GetMeetingRoomsScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMeetingRoomsScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMeetingRoomsScheduleShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoomIds) {
		request.RoomIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomIds, dara.String("RoomIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RoomIdsShrink) {
		body["RoomIds"] = request.RoomIdsShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMeetingRoomsSchedule"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/getMeetingRoomsSchedule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMeetingRoomsScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取我的文档知识库信息
//
// @param tmpReq - GetMineWorkspaceRequest
//
// @param tmpHeader - GetMineWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMineWorkspaceResponse
func (client *Client) GetMineWorkspaceWithContext(ctx context.Context, tmpReq *GetMineWorkspaceRequest, tmpHeader *GetMineWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *GetMineWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMineWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMineWorkspaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMineWorkspace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getMineWorkspace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMineWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有字段
//
// @param tmpReq - GetMultiDimTableAllFieldsRequest
//
// @param tmpHeader - GetMultiDimTableAllFieldsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiDimTableAllFieldsResponse
func (client *Client) GetMultiDimTableAllFieldsWithContext(ctx context.Context, tmpReq *GetMultiDimTableAllFieldsRequest, tmpHeader *GetMultiDimTableAllFieldsHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableAllFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMultiDimTableAllFieldsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMultiDimTableAllFieldsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiDimTableAllFields"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/getMultiDimTableAllFields"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiDimTableAllFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有数据表
//
// @param tmpReq - GetMultiDimTableAllSheetsRequest
//
// @param tmpHeader - GetMultiDimTableAllSheetsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiDimTableAllSheetsResponse
func (client *Client) GetMultiDimTableAllSheetsWithContext(ctx context.Context, tmpReq *GetMultiDimTableAllSheetsRequest, tmpHeader *GetMultiDimTableAllSheetsHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableAllSheetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMultiDimTableAllSheetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMultiDimTableAllSheetsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiDimTableAllSheets"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/getMultiDimTableAllSheets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiDimTableAllSheetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取记录
//
// @param tmpReq - GetMultiDimTableRecordRequest
//
// @param tmpHeader - GetMultiDimTableRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiDimTableRecordResponse
func (client *Client) GetMultiDimTableRecordWithContext(ctx context.Context, tmpReq *GetMultiDimTableRecordRequest, tmpHeader *GetMultiDimTableRecordHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMultiDimTableRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMultiDimTableRecordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.RecordId) {
		body["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiDimTableRecord"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/getMultiDimTableRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiDimTableRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据表
//
// @param tmpReq - GetMultiDimTableSheetRequest
//
// @param tmpHeader - GetMultiDimTableSheetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiDimTableSheetResponse
func (client *Client) GetMultiDimTableSheetWithContext(ctx context.Context, tmpReq *GetMultiDimTableSheetRequest, tmpHeader *GetMultiDimTableSheetHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableSheetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMultiDimTableSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMultiDimTableSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiDimTableSheet"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/getMultiDimTableSheet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiDimTableSheetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件分片上传信息
//
// @param tmpReq - GetMultipartFileUploadInfosRequest
//
// @param tmpHeader - GetMultipartFileUploadInfosHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultipartFileUploadInfosResponse
func (client *Client) GetMultipartFileUploadInfosWithContext(ctx context.Context, tmpReq *GetMultipartFileUploadInfosRequest, tmpHeader *GetMultipartFileUploadInfosHeaders, runtime *dara.RuntimeOptions) (_result *GetMultipartFileUploadInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetMultipartFileUploadInfosShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetMultipartFileUploadInfosShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PartNumbers) {
		request.PartNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartNumbers, dara.String("PartNumbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.PartNumbersShrink) {
		body["PartNumbers"] = request.PartNumbersShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UploadKey) {
		body["UploadKey"] = request.UploadKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultipartFileUploadInfos"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/GetMultipartFileUploadInfos"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultipartFileUploadInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最近活跃的企业内部群列表
//
// @param tmpReq - GetNewestInnerGroupsRequest
//
// @param tmpHeader - GetNewestInnerGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNewestInnerGroupsResponse
func (client *Client) GetNewestInnerGroupsWithContext(ctx context.Context, tmpReq *GetNewestInnerGroupsRequest, tmpHeader *GetNewestInnerGroupsHeaders, runtime *dara.RuntimeOptions) (_result *GetNewestInnerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNewestInnerGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetNewestInnerGroupsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNewestInnerGroups"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/getNewestInnerGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNewestInnerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点
//
// @param tmpReq - GetNodeRequest
//
// @param tmpHeader - GetNodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithContext(ctx context.Context, tmpReq *GetNodeRequest, tmpHeader *GetNodeHeaders, runtime *dara.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetNodeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithPermissionRole) {
		body["WithPermissionRole"] = request.WithPermissionRole
	}

	if !dara.IsNil(request.WithStatisticalInfo) {
		body["WithStatisticalInfo"] = request.WithStatisticalInfo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNode"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getNode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过链接获取节点
//
// @param tmpReq - GetNodeByUrlRequest
//
// @param tmpHeader - GetNodeByUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeByUrlResponse
func (client *Client) GetNodeByUrlWithContext(ctx context.Context, tmpReq *GetNodeByUrlRequest, tmpHeader *GetNodeByUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetNodeByUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNodeByUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetNodeByUrlShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeByUrl"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getNodeByUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeByUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取节点
//
// @param tmpReq - GetNodesRequest
//
// @param tmpHeader - GetNodesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodesResponse
func (client *Client) GetNodesWithContext(ctx context.Context, tmpReq *GetNodesRequest, tmpHeader *GetNodesHeaders, runtime *dara.RuntimeOptions) (_result *GetNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetNodesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeIds) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, dara.String("NodeIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeIdsShrink) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodes"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getNodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发送给用户的通知
//
// @param request - GetNotifyMeRequest
//
// @param tmpHeader - GetNotifyMeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotifyMeResponse
func (client *Client) GetNotifyMeWithContext(ctx context.Context, request *GetNotifyMeRequest, tmpHeader *GetNotifyMeHeaders, runtime *dara.RuntimeOptions) (_result *GetNotifyMeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetNotifyMeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppTypes) {
		body["AppTypes"] = request.AppTypes
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.InstanceCreateFromTimeGMT) {
		body["InstanceCreateFromTimeGMT"] = request.InstanceCreateFromTimeGMT
	}

	if !dara.IsNil(request.InstanceCreateToTimeGMT) {
		body["InstanceCreateToTimeGMT"] = request.InstanceCreateToTimeGMT
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotifyMe"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getNotifyMe"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotifyMeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取宜搭附件临时免登地址
//
// @param request - GetOpenUrlRequest
//
// @param tmpHeader - GetOpenUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenUrlResponse
func (client *Client) GetOpenUrlWithContext(ctx context.Context, request *GetOpenUrlRequest, tmpHeader *GetOpenUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetOpenUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetOpenUrlShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpenUrl"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getOpenUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpenUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审批记录
//
// @param request - GetOperationRecordsRequest
//
// @param tmpHeader - GetOperationRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationRecordsResponse
func (client *Client) GetOperationRecordsWithContext(ctx context.Context, request *GetOperationRecordsRequest, tmpHeader *GetOperationRecordsHeaders, runtime *dara.RuntimeOptions) (_result *GetOperationRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetOperationRecordsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationRecords"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getOperationRecords"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - GetOrgLiveListRequest
//
// @param tmpHeader - GetOrgLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgLiveListResponse
func (client *Client) GetOrgLiveListWithContext(ctx context.Context, tmpReq *GetOrgLiveListRequest, tmpHeader *GetOrgLiveListHeaders, runtime *dara.RuntimeOptions) (_result *GetOrgLiveListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOrgLiveListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetOrgLiveListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrgLiveList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/getOrgLiveList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrgLiveListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限获取组织或者互联网公开文档内容taskId
//
// @param tmpReq - GetOrgOrWebOpenDocContentTaskIdRequest
//
// @param tmpHeader - GetOrgOrWebOpenDocContentTaskIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgOrWebOpenDocContentTaskIdResponse
func (client *Client) GetOrgOrWebOpenDocContentTaskIdWithContext(ctx context.Context, tmpReq *GetOrgOrWebOpenDocContentTaskIdRequest, tmpHeader *GetOrgOrWebOpenDocContentTaskIdHeaders, runtime *dara.RuntimeOptions) (_result *GetOrgOrWebOpenDocContentTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOrgOrWebOpenDocContentTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetOrgOrWebOpenDocContentTaskIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.GenerateCp) {
		body["GenerateCp"] = request.GenerateCp
	}

	if !dara.IsNil(request.ScopeType) {
		body["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.TargetFormat) {
		body["TargetFormat"] = request.TargetFormat
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrgOrWebOpenDocContentTaskId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getOrgOrWebOpenDocContentTaskId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrgOrWebOpenDocContentTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程定义
//
// @param request - GetProcessDefinitionRequest
//
// @param tmpHeader - GetProcessDefinitionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessDefinitionResponse
func (client *Client) GetProcessDefinitionWithContext(ctx context.Context, request *GetProcessDefinitionRequest, tmpHeader *GetProcessDefinitionHeaders, runtime *dara.RuntimeOptions) (_result *GetProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetProcessDefinitionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.NameSpace) {
		body["NameSpace"] = request.NameSpace
	}

	if !dara.IsNil(request.OrderNumber) {
		body["OrderNumber"] = request.OrderNumber
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProcessDefinition"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getProcessDefinition"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单元格区域
//
// @param tmpReq - GetRangeRequest
//
// @param tmpHeader - GetRangeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRangeResponse
func (client *Client) GetRangeWithContext(ctx context.Context, tmpReq *GetRangeRequest, tmpHeader *GetRangeHeaders, runtime *dara.RuntimeOptions) (_result *GetRangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetRangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetRangeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RangeAddress) {
		body["RangeAddress"] = request.RangeAddress
	}

	if !dara.IsNil(request.Select) {
		body["Select"] = request.Select
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRange"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getRange"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRangeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户有权限的知识库列表(旧)
//
// @param tmpReq - GetRelatedWorkspacesRequest
//
// @param tmpHeader - GetRelatedWorkspacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRelatedWorkspacesResponse
func (client *Client) GetRelatedWorkspacesWithContext(ctx context.Context, tmpReq *GetRelatedWorkspacesRequest, tmpHeader *GetRelatedWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *GetRelatedWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetRelatedWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetRelatedWorkspacesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IncludeRecent) {
		body["IncludeRecent"] = request.IncludeRecent
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRelatedWorkspaces"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getRelatedWorkspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRelatedWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模板详情
//
// @param tmpReq - GetReportTemplateByNameRequest
//
// @param tmpHeader - GetReportTemplateByNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportTemplateByNameResponse
func (client *Client) GetReportTemplateByNameWithContext(ctx context.Context, tmpReq *GetReportTemplateByNameRequest, tmpHeader *GetReportTemplateByNameHeaders, runtime *dara.RuntimeOptions) (_result *GetReportTemplateByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetReportTemplateByNameShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetReportTemplateByNameShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReportTemplateByName"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/getReportTemplateByName"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportTemplateByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取员工有多少数量的日志（一个月内）是未读状态
//
// @param tmpReq - GetReportUnReadCountRequest
//
// @param tmpHeader - GetReportUnReadCountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportUnReadCountResponse
func (client *Client) GetReportUnReadCountWithContext(ctx context.Context, tmpReq *GetReportUnReadCountRequest, tmpHeader *GetReportUnReadCountHeaders, runtime *dara.RuntimeOptions) (_result *GetReportUnReadCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetReportUnReadCountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetReportUnReadCountShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReportUnReadCount"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/getReportUnReadCount"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportUnReadCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流程运行任务（VPC）
//
// @param request - GetRunningTasksRequest
//
// @param tmpHeader - GetRunningTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunningTasksResponse
func (client *Client) GetRunningTasksWithContext(ctx context.Context, request *GetRunningTasksRequest, tmpHeader *GetRunningTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetRunningTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetRunningTasksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRunningTasks"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getRunningTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRunningTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户忙闲信息
//
// @param tmpReq - GetScheduleRequest
//
// @param tmpHeader - GetScheduleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduleResponse
func (client *Client) GetScheduleWithContext(ctx context.Context, tmpReq *GetScheduleRequest, tmpHeader *GetScheduleHeaders, runtime *dara.RuntimeOptions) (_result *GetScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetScheduleShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["UserIds"] = request.UserIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSchedule"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/getSchedule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作表
//
// @param tmpReq - GetSheetRequest
//
// @param tmpHeader - GetSheetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSheetResponse
func (client *Client) GetSheetWithContext(ctx context.Context, tmpReq *GetSheetRequest, tmpHeader *GetSheetHeaders, runtime *dara.RuntimeOptions) (_result *GetSheetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSheet"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/getSheet"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSheetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表格文档内容jobId
//
// @param tmpReq - GetSheetContentJobIdRequest
//
// @param tmpHeader - GetSheetContentJobIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSheetContentJobIdResponse
func (client *Client) GetSheetContentJobIdWithContext(ctx context.Context, tmpReq *GetSheetContentJobIdRequest, tmpHeader *GetSheetContentJobIdHeaders, runtime *dara.RuntimeOptions) (_result *GetSheetContentJobIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSheetContentJobIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetSheetContentJobIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSheetContentJobId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getSheetContentJobId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSheetContentJobIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询知识库下的目录结构
//
// @param tmpReq - GetSpaceDirectoriesRequest
//
// @param tmpHeader - GetSpaceDirectoriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceDirectoriesResponse
func (client *Client) GetSpaceDirectoriesWithContext(ctx context.Context, tmpReq *GetSpaceDirectoriesRequest, tmpHeader *GetSpaceDirectoriesHeaders, runtime *dara.RuntimeOptions) (_result *GetSpaceDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSpaceDirectoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetSpaceDirectoriesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpaceDirectories"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getSpaceDirectories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpaceDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个订阅日历详情
//
// @param request - GetSubscribedCalendarRequest
//
// @param tmpHeader - GetSubscribedCalendarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscribedCalendarResponse
func (client *Client) GetSubscribedCalendarWithContext(ctx context.Context, request *GetSubscribedCalendarRequest, tmpHeader *GetSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *GetSubscribedCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetSubscribedCalendarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscribedCalendar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/getSubscribedCalendar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscribedCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询抄送我的任务列表（应用维度）
//
// @param request - GetTaskCopiesRequest
//
// @param tmpHeader - GetTaskCopiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskCopiesResponse
func (client *Client) GetTaskCopiesWithContext(ctx context.Context, request *GetTaskCopiesRequest, tmpHeader *GetTaskCopiesHeaders, runtime *dara.RuntimeOptions) (_result *GetTaskCopiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &GetTaskCopiesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessCodes) {
		body["ProcessCodes"] = request.ProcessCodes
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskCopies"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/getTaskCopies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskCopiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户可见的日志模板
//
// @param tmpReq - GetTemplateListByUserIdRequest
//
// @param tmpHeader - GetTemplateListByUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateListByUserIdResponse
func (client *Client) GetTemplateListByUserIdWithContext(ctx context.Context, tmpReq *GetTemplateListByUserIdRequest, tmpHeader *GetTemplateListByUserIdHeaders, runtime *dara.RuntimeOptions) (_result *GetTemplateListByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTemplateListByUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetTemplateListByUserIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateListByUserId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/getTemplateListByUserId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateListByUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定工单详情
//
// @param tmpReq - GetTicketRequest
//
// @param tmpHeader - GetTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketResponse
func (client *Client) GetTicketWithContext(ctx context.Context, tmpReq *GetTicketRequest, tmpHeader *GetTicketHeaders, runtime *dara.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetTicketShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicket"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/getTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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
// 获取钉钉待办任务详情
//
// @param tmpReq - GetTodoTaskRequest
//
// @param tmpHeader - GetTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodoTaskResponse
func (client *Client) GetTodoTaskWithContext(ctx context.Context, tmpReq *GetTodoTaskRequest, tmpHeader *GetTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *GetTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/getTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专属账号信息
//
// @param tmpReq - GetUserRequest
//
// @param tmpHeader - GetUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, tmpReq *GetUserRequest, tmpHeader *GetUserHeaders, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetUserShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/getUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据unionId获取用户userId
//
// @param tmpReq - GetUserIdRequest
//
// @param tmpHeader - GetUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdResponse
func (client *Client) GetUserIdWithContext(ctx context.Context, tmpReq *GetUserIdRequest, tmpHeader *GetUserIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetUserIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UnionId) {
		body["UnionId"] = request.UnionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/getUserId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据orgId和staffId获取用户userId
//
// @param tmpReq - GetUserIdByOrgIdAndStaffIdRequest
//
// @param tmpHeader - GetUserIdByOrgIdAndStaffIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByOrgIdAndStaffIdResponse
func (client *Client) GetUserIdByOrgIdAndStaffIdWithContext(ctx context.Context, tmpReq *GetUserIdByOrgIdAndStaffIdRequest, tmpHeader *GetUserIdByOrgIdAndStaffIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByOrgIdAndStaffIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserIdByOrgIdAndStaffIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetUserIdByOrgIdAndStaffIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrgId) {
		body["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserIdByOrgIdAndStaffId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/getUserIdByOrgIdAndStaffId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdByOrgIdAndStaffIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户最新的有效的专属账号迁移方案
//
// @param tmpReq - GetUserLatestPlanRequest
//
// @param tmpHeader - GetUserLatestPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserLatestPlanResponse
func (client *Client) GetUserLatestPlanWithContext(ctx context.Context, tmpReq *GetUserLatestPlanRequest, tmpHeader *GetUserLatestPlanHeaders, runtime *dara.RuntimeOptions) (_result *GetUserLatestPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserLatestPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetUserLatestPlanShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserLatestPlan"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/indepding/getUserLatestPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserLatestPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库
//
// @param tmpReq - GetWorkspaceRequest
//
// @param tmpHeader - GetWorkspaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithContext(ctx context.Context, tmpReq *GetWorkspaceRequest, tmpHeader *GetWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetWorkspaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithPermissionRole) {
		body["WithPermissionRole"] = request.WithPermissionRole
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getWorkspace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取知识库
//
// @param tmpReq - GetWorkspacesRequest
//
// @param tmpHeader - GetWorkspacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspacesResponse
func (client *Client) GetWorkspacesWithContext(ctx context.Context, tmpReq *GetWorkspacesRequest, tmpHeader *GetWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *GetWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetWorkspacesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WorkspaceIds) {
		request.WorkspaceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WorkspaceIds, dara.String("WorkspaceIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceIdsShrink) {
		body["WorkspaceIds"] = request.WorkspaceIdsShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaces"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/getWorkspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授予勋章
//
// @param tmpReq - GrantHonorRequest
//
// @param tmpHeader - GrantHonorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantHonorResponse
func (client *Client) GrantHonorWithContext(ctx context.Context, tmpReq *GrantHonorRequest, tmpHeader *GrantHonorHeaders, runtime *dara.RuntimeOptions) (_result *GrantHonorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantHonorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GrantHonorShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenConversationIds) {
		request.OpenConversationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenConversationIds, dara.String("openConversationIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReceiverUserIds) {
		request.ReceiverUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReceiverUserIds, dara.String("receiverUserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ExpirationTime) {
		body["expirationTime"] = request.ExpirationTime
	}

	if !dara.IsNil(request.GrantReason) {
		body["grantReason"] = request.GrantReason
	}

	if !dara.IsNil(request.GranterName) {
		body["granterName"] = request.GranterName
	}

	if !dara.IsNil(request.HonorId) {
		body["honorId"] = request.HonorId
	}

	if !dara.IsNil(request.NoticeAnnouncer) {
		body["noticeAnnouncer"] = request.NoticeAnnouncer
	}

	if !dara.IsNil(request.NoticeSingle) {
		body["noticeSingle"] = request.NoticeSingle
	}

	if !dara.IsNil(request.OpenConversationIdsShrink) {
		body["openConversationIds"] = request.OpenConversationIdsShrink
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.ReceiverUserIdsShrink) {
		body["receiverUserIds"] = request.ReceiverUserIdsShrink
	}

	if !dara.IsNil(request.SenderUserId) {
		body["senderUserId"] = request.SenderUserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantHonor"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/grantHonor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantHonorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化文件分片上传
//
// @param tmpReq - InitMultipartFileUploadRequest
//
// @param tmpHeader - InitMultipartFileUploadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitMultipartFileUploadResponse
func (client *Client) InitMultipartFileUploadWithContext(ctx context.Context, tmpReq *InitMultipartFileUploadRequest, tmpHeader *InitMultipartFileUploadHeaders, runtime *dara.RuntimeOptions) (_result *InitMultipartFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InitMultipartFileUploadShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InitMultipartFileUploadShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.ParentDentryUuid) {
		body["ParentDentryUuid"] = request.ParentDentryUuid
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitMultipartFileUpload"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/initMultipartFileUpload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitMultipartFileUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在指定列左侧插入若干列
//
// @param tmpReq - InsertColumnsBeforeRequest
//
// @param tmpHeader - InsertColumnsBeforeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertColumnsBeforeResponse
func (client *Client) InsertColumnsBeforeWithContext(ctx context.Context, tmpReq *InsertColumnsBeforeRequest, tmpHeader *InsertColumnsBeforeHeaders, runtime *dara.RuntimeOptions) (_result *InsertColumnsBeforeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertColumnsBeforeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertColumnsBeforeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Column) {
		body["Column"] = request.Column
	}

	if !dara.IsNil(request.ColumnCount) {
		body["ColumnCount"] = request.ColumnCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertColumnsBefore"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/insertColumnsBefore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertColumnsBeforeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档中插入内容
//
// @param tmpReq - InsertContentWithOptionsRequest
//
// @param tmpHeader - InsertContentWithOptionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertContentWithOptionsResponse
func (client *Client) InsertContentWithOptionsWithContext(ctx context.Context, tmpReq *InsertContentWithOptionsRequest, tmpHeader *InsertContentWithOptionsHeaders, runtime *dara.RuntimeOptions) (_result *InsertContentWithOptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertContentWithOptionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertContentWithOptionsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.DocumentId) {
		body["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.Index) {
		body["Index"] = request.Index
	}

	if !dara.IsNil(request.PathShrink) {
		body["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertContentWithOptions"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/insertContentWithOptions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertContentWithOptionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插入下拉列表
//
// @param tmpReq - InsertDropDownListRequest
//
// @param tmpHeader - InsertDropDownListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertDropDownListResponse
func (client *Client) InsertDropDownListWithContext(ctx context.Context, tmpReq *InsertDropDownListRequest, tmpHeader *InsertDropDownListHeaders, runtime *dara.RuntimeOptions) (_result *InsertDropDownListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertDropDownListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertDropDownListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OptionsShrink) {
		body["Options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.RangeAddress) {
		body["RangeAddress"] = request.RangeAddress
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertDropDownList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/insertDropDownList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertDropDownListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增记录
//
// @param tmpReq - InsertMultiDimTableRecordRequest
//
// @param tmpHeader - InsertMultiDimTableRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertMultiDimTableRecordResponse
func (client *Client) InsertMultiDimTableRecordWithContext(ctx context.Context, tmpReq *InsertMultiDimTableRecordRequest, tmpHeader *InsertMultiDimTableRecordHeaders, runtime *dara.RuntimeOptions) (_result *InsertMultiDimTableRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertMultiDimTableRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertMultiDimTableRecordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Records) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, dara.String("Records"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.RecordsShrink) {
		body["Records"] = request.RecordsShrink
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertMultiDimTableRecord"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/insertMultiDimTableRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertMultiDimTableRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在指定行上方插入若干行
//
// @param tmpReq - InsertRowsBeforeRequest
//
// @param tmpHeader - InsertRowsBeforeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertRowsBeforeResponse
func (client *Client) InsertRowsBeforeWithContext(ctx context.Context, tmpReq *InsertRowsBeforeRequest, tmpHeader *InsertRowsBeforeHeaders, runtime *dara.RuntimeOptions) (_result *InsertRowsBeforeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InsertRowsBeforeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertRowsBeforeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Row) {
		body["Row"] = request.Row
	}

	if !dara.IsNil(request.RowCount) {
		body["RowCount"] = request.RowCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertRowsBefore"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/insertRowsBefore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 邀请用户入会
//
// @param tmpReq - InviteUsersRequest
//
// @param tmpHeader - InviteUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteUsersResponse
func (client *Client) InviteUsersWithContext(ctx context.Context, tmpReq *InviteUsersRequest, tmpHeader *InviteUsersHeaders, runtime *dara.RuntimeOptions) (_result *InviteUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InviteUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InviteUsersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InviteeList) {
		request.InviteeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InviteeList, dara.String("InviteeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PhoneInviteeList) {
		request.PhoneInviteeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneInviteeList, dara.String("phoneInviteeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InviteeListShrink) {
		body["InviteeList"] = request.InviteeListShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.PhoneInviteeListShrink) {
		body["phoneInviteeList"] = request.PhoneInviteeListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InviteUsers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/inviteUsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InviteUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用助理
//
// @param request - InvokeAssistantRequest
//
// @param headers - InvokeAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeAssistantResponse
func (client *Client) InvokeAssistantWithSSECtx(ctx context.Context, request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions, _yield chan *InvokeAssistantResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeAssistantWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 调用助理
//
// @param request - InvokeAssistantRequest
//
// @param headers - InvokeAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeAssistantResponse
func (client *Client) InvokeAssistantWithContext(ctx context.Context, request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions) (_result *InvokeAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ClientEnum) {
		body["clientEnum"] = request.ClientEnum
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAssistant"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/invokeAssistant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeAssistantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用AI技能
//
// @param tmpReq - InvokeSkillRequest
//
// @param tmpHeader - InvokeSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeSkillResponse
func (client *Client) InvokeSkillWithSSECtx(ctx context.Context, tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions, _yield chan *InvokeSkillResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeSkillWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, tmpHeader, runtime)
	return
}

// Summary:
//
// 调用AI技能
//
// @param tmpReq - InvokeSkillRequest
//
// @param tmpHeader - InvokeSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeSkillResponse
func (client *Client) InvokeSkillWithContext(ctx context.Context, tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions) (_result *InvokeSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InvokeSkillShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.SkillId) {
		body["SkillId"] = request.SkillId
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.SourceIdOfAssistantId) {
		body["sourceIdOfAssistantId"] = request.SourceIdOfAssistantId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeSkill"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/skill/invoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询宜搭应用列表
//
// @param request - ListApplicationRequest
//
// @param tmpHeader - ListApplicationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationResponse
func (client *Client) ListApplicationWithContext(ctx context.Context, request *ListApplicationRequest, tmpHeader *ListApplicationHeaders, runtime *dara.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListApplicationShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppFilter) {
		body["AppFilter"] = request.AppFilter
	}

	if !dara.IsNil(request.AppNameSearchKeyword) {
		body["AppNameSearchKeyword"] = request.AppNameSearchKeyword
	}

	if !dara.IsNil(request.CorpId) {
		body["CorpId"] = request.CorpId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplication"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/listApplication"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日历
//
// @param tmpReq - ListCalendarsRequest
//
// @param tmpHeader - ListCalendarsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalendarsResponse
func (client *Client) ListCalendarsWithContext(ctx context.Context, tmpReq *ListCalendarsRequest, tmpHeader *ListCalendarsHeaders, runtime *dara.RuntimeOptions) (_result *ListCalendarsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCalendarsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListCalendarsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalendars"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/listCalendars"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalendarsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件或文件夹列表
//
// @param tmpReq - ListDentriesRequest
//
// @param tmpHeader - ListDentriesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDentriesResponse
func (client *Client) ListDentriesWithContext(ctx context.Context, tmpReq *ListDentriesRequest, tmpHeader *ListDentriesHeaders, runtime *dara.RuntimeOptions) (_result *ListDentriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDentriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListDentriesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.ParentId) {
		body["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithThumbnail) {
		body["WithThumbnail"] = request.WithThumbnail
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDentries"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/listDentries"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDentriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取钉盘空间列表
//
// @param tmpReq - ListDriveSpacesRequest
//
// @param tmpHeader - ListDriveSpacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDriveSpacesResponse
func (client *Client) ListDriveSpacesWithContext(ctx context.Context, tmpReq *ListDriveSpacesRequest, tmpHeader *ListDriveSpacesHeaders, runtime *dara.RuntimeOptions) (_result *ListDriveSpacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDriveSpacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListDriveSpacesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SpaceType) {
		body["SpaceType"] = request.SpaceType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDriveSpaces"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/listDriveSpaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDriveSpacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日程列表
//
// @param request - ListEventsRequest
//
// @param tmpHeader - ListEventsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventsResponse
func (client *Client) ListEventsWithContext(ctx context.Context, request *ListEventsRequest, tmpHeader *ListEventsHeaders, runtime *dara.RuntimeOptions) (_result *ListEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListEventsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.MaxAttendees) {
		body["MaxAttendees"] = request.MaxAttendees
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SeriesMasterId) {
		body["SeriesMasterId"] = request.SeriesMasterId
	}

	if !dara.IsNil(request.ShowDeleted) {
		body["ShowDeleted"] = request.ShowDeleted
	}

	if !dara.IsNil(request.SyncToken) {
		body["SyncToken"] = request.SyncToken
	}

	if !dara.IsNil(request.TimeMax) {
		body["TimeMax"] = request.TimeMax
	}

	if !dara.IsNil(request.TimeMin) {
		body["TimeMin"] = request.TimeMin
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvents"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/listEvents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日程视图
//
// @param request - ListEventsViewRequest
//
// @param tmpHeader - ListEventsViewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventsViewResponse
func (client *Client) ListEventsViewWithContext(ctx context.Context, request *ListEventsViewRequest, tmpHeader *ListEventsViewHeaders, runtime *dara.RuntimeOptions) (_result *ListEventsViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListEventsViewShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.MaxAttendees) {
		body["MaxAttendees"] = request.MaxAttendees
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TimeMax) {
		body["TimeMax"] = request.TimeMax
	}

	if !dara.IsNil(request.TimeMin) {
		body["TimeMin"] = request.TimeMin
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEventsView"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/listEventsView"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventsViewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询宜搭表单实例的评论
//
// @param tmpReq - ListFormRemarksRequest
//
// @param tmpHeader - ListFormRemarksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFormRemarksResponse
func (client *Client) ListFormRemarksWithContext(ctx context.Context, tmpReq *ListFormRemarksRequest, tmpHeader *ListFormRemarksHeaders, runtime *dara.RuntimeOptions) (_result *ListFormRemarksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFormRemarksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListFormRemarksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FormInstanceIdList) {
		request.FormInstanceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FormInstanceIdList, dara.String("FormInstanceIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormInstanceIdListShrink) {
		body["FormInstanceIdList"] = request.FormInstanceIdListShrink
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFormRemarks"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/listFormRemarks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFormRemarksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消息
//
// @param request - ListMessageRequest
//
// @param headers - ListMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageResponse
func (client *Client) ListMessageWithContext(ctx context.Context, request *ListMessageRequest, headers *ListMessageHeaders, runtime *dara.RuntimeOptions) (_result *ListMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessage"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/listMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出多行记录
//
// @param tmpReq - ListMultiDimTableRecordsRequest
//
// @param tmpHeader - ListMultiDimTableRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiDimTableRecordsResponse
func (client *Client) ListMultiDimTableRecordsWithContext(ctx context.Context, tmpReq *ListMultiDimTableRecordsRequest, tmpHeader *ListMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *ListMultiDimTableRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMultiDimTableRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListMultiDimTableRecordsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiDimTableRecords"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/listMultiDimTableRecords"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiDimTableRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用下的页面列表
//
// @param request - ListNavigationByFormTypeRequest
//
// @param tmpHeader - ListNavigationByFormTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNavigationByFormTypeResponse
func (client *Client) ListNavigationByFormTypeWithContext(ctx context.Context, request *ListNavigationByFormTypeRequest, tmpHeader *ListNavigationByFormTypeHeaders, runtime *dara.RuntimeOptions) (_result *ListNavigationByFormTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListNavigationByFormTypeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormType) {
		body["FormType"] = request.FormType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNavigationByFormType"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/listNavigationByFormType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNavigationByFormTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点列表
//
// @param tmpReq - ListNodesRequest
//
// @param tmpHeader - ListNodesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, tmpReq *ListNodesRequest, tmpHeader *ListNodesHeaders, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListNodesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentNodeId) {
		body["ParentNodeId"] = request.ParentNodeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithPermissionRole) {
		body["WithPermissionRole"] = request.WithPermissionRole
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/listNodes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件权限列表
//
// @param tmpReq - ListPermissionsRequest
//
// @param tmpHeader - ListPermissionsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithContext(ctx context.Context, tmpReq *ListPermissionsRequest, tmpHeader *ListPermissionsHeaders, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListPermissionsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissions"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/listPermissions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户发出的日志列表
//
// @param tmpReq - ListReportRequest
//
// @param tmpHeader - ListReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportResponse
func (client *Client) ListReportWithContext(ctx context.Context, tmpReq *ListReportRequest, tmpHeader *ListReportHeaders, runtime *dara.RuntimeOptions) (_result *ListReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ModifiedEndTime) {
		body["ModifiedEndTime"] = request.ModifiedEndTime
	}

	if !dara.IsNil(request.ModifiedStartTime) {
		body["ModifiedStartTime"] = request.ModifiedStartTime
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/listReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出AI技能
//
// @param request - ListSkillRequest
//
// @param tmpHeader - ListSkillHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillResponse
func (client *Client) ListSkillWithContext(ctx context.Context, request *ListSkillRequest, tmpHeader *ListSkillHeaders, runtime *dara.RuntimeOptions) (_result *ListSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListSkillShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["groupId"] = request.GroupId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkill"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/skill/listSkill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取子表组件数据
//
// @param request - ListTableDataByFormInstanceIdTableIdRequest
//
// @param tmpHeader - ListTableDataByFormInstanceIdTableIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableDataByFormInstanceIdTableIdResponse
func (client *Client) ListTableDataByFormInstanceIdTableIdWithContext(ctx context.Context, request *ListTableDataByFormInstanceIdTableIdRequest, tmpHeader *ListTableDataByFormInstanceIdTableIdHeaders, runtime *dara.RuntimeOptions) (_result *ListTableDataByFormInstanceIdTableIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &ListTableDataByFormInstanceIdTableIdShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormInstanceId) {
		body["FormInstanceId"] = request.FormInstanceId
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TableFieldId) {
		body["TableFieldId"] = request.TableFieldId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableDataByFormInstanceIdTableId"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/listTableDataByFormInstanceIdTableId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableDataByFormInstanceIdTableIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识小组列表
//
// @param tmpReq - ListTeamsRequest
//
// @param tmpHeader - ListTeamsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamsResponse
func (client *Client) ListTeamsWithContext(ctx context.Context, tmpReq *ListTeamsRequest, tmpHeader *ListTeamsHeaders, runtime *dara.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTeamsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListTeamsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeams"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/listTeams"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档模板
//
// @param tmpReq - ListTemplateRequest
//
// @param tmpHeader - ListTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateResponse
func (client *Client) ListTemplateWithContext(ctx context.Context, tmpReq *ListTemplateRequest, tmpHeader *ListTemplateHeaders, runtime *dara.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListTemplateShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplate"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/listTemplate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工单操作记录
//
// @param tmpReq - ListTicketOperateRecordRequest
//
// @param tmpHeader - ListTicketOperateRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketOperateRecordResponse
func (client *Client) ListTicketOperateRecordWithContext(ctx context.Context, tmpReq *ListTicketOperateRecordRequest, tmpHeader *ListTicketOperateRecordHeaders, runtime *dara.RuntimeOptions) (_result *ListTicketOperateRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTicketOperateRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListTicketOperateRecordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTicketOperateRecord"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/listTicketOperateRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库列表
//
// @param tmpReq - ListWorkspacesRequest
//
// @param tmpHeader - ListWorkspacesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithContext(ctx context.Context, tmpReq *ListWorkspacesRequest, tmpHeader *ListWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkspacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ListWorkspacesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.TeamId) {
		body["TeamId"] = request.TeamId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithPermissionRole) {
		body["WithPermissionRole"] = request.WithPermissionRole
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/listWorkspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改日程
//
// @param tmpReq - PatchEventRequest
//
// @param tmpHeader - PatchEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchEventResponse
func (client *Client) PatchEventWithContext(ctx context.Context, tmpReq *PatchEventRequest, tmpHeader *PatchEventHeaders, runtime *dara.RuntimeOptions) (_result *PatchEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PatchEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &PatchEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Attendees) {
		request.AttendeesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attendees, dara.String("Attendees"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CardInstances) {
		request.CardInstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CardInstances, dara.String("CardInstances"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.End) {
		request.EndShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.End, dara.String("End"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Location) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, dara.String("Location"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Recurrence) {
		request.RecurrenceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Recurrence, dara.String("Recurrence"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Reminders) {
		request.RemindersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Reminders, dara.String("Reminders"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Start) {
		request.StartShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Start, dara.String("Start"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("categories"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OnlineMeetingInfo) {
		request.OnlineMeetingInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OnlineMeetingInfo, dara.String("onlineMeetingInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RichTextDescription) {
		request.RichTextDescriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RichTextDescription, dara.String("richTextDescription"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UiConfigs) {
		request.UiConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UiConfigs, dara.String("uiConfigs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		query["categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.FreeBusyStatus) {
		query["freeBusyStatus"] = request.FreeBusyStatus
	}

	if !dara.IsNil(request.OnlineMeetingInfoShrink) {
		query["onlineMeetingInfo"] = request.OnlineMeetingInfoShrink
	}

	if !dara.IsNil(request.RichTextDescriptionShrink) {
		query["richTextDescription"] = request.RichTextDescriptionShrink
	}

	if !dara.IsNil(request.UiConfigsShrink) {
		query["uiConfigs"] = request.UiConfigsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttendeesShrink) {
		body["Attendees"] = request.AttendeesShrink
	}

	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.CardInstancesShrink) {
		body["CardInstances"] = request.CardInstancesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndShrink) {
		body["End"] = request.EndShrink
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.IsAllDay) {
		body["IsAllDay"] = request.IsAllDay
	}

	if !dara.IsNil(request.LocationShrink) {
		body["Location"] = request.LocationShrink
	}

	if !dara.IsNil(request.RecurrenceShrink) {
		body["Recurrence"] = request.RecurrenceShrink
	}

	if !dara.IsNil(request.RemindersShrink) {
		body["Reminders"] = request.RemindersShrink
	}

	if !dara.IsNil(request.StartShrink) {
		body["Start"] = request.StartShrink
	}

	if !dara.IsNil(request.Summary) {
		body["Summary"] = request.Summary
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/patchEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PatchEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议录制中的文本信息
//
// @param tmpReq - QueryCloudRecordTextRequest
//
// @param tmpHeader - QueryCloudRecordTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordTextResponse
func (client *Client) QueryCloudRecordTextWithContext(ctx context.Context, tmpReq *QueryCloudRecordTextRequest, tmpHeader *QueryCloudRecordTextHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryCloudRecordTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryCloudRecordTextShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		body["Direction"] = request.Direction
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCloudRecordText"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryCloudRecordText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议录制的详情信息
//
// @param tmpReq - QueryCloudRecordVideoRequest
//
// @param tmpHeader - QueryCloudRecordVideoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordVideoResponse
func (client *Client) QueryCloudRecordVideoWithContext(ctx context.Context, tmpReq *QueryCloudRecordVideoRequest, tmpHeader *QueryCloudRecordVideoHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryCloudRecordVideoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryCloudRecordVideoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCloudRecordVideo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryCloudRecordVideo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议录制中的视频信息
//
// @param tmpReq - QueryCloudRecordVideoPlayInfoRequest
//
// @param tmpHeader - QueryCloudRecordVideoPlayInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordVideoPlayInfoResponse
func (client *Client) QueryCloudRecordVideoPlayInfoWithContext(ctx context.Context, tmpReq *QueryCloudRecordVideoPlayInfoRequest, tmpHeader *QueryCloudRecordVideoPlayInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryCloudRecordVideoPlayInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryCloudRecordVideoPlayInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConferenceId) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.MediaId) {
		body["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCloudRecordVideoPlayInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryCloudRecordVideoPlayInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频会议信息
//
// @param request - QueryConferenceInfoRequest
//
// @param tmpHeader - QueryConferenceInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceInfoResponse
func (client *Client) QueryConferenceInfoWithContext(ctx context.Context, request *QueryConferenceInfoRequest, tmpHeader *QueryConferenceInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &QueryConferenceInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConferenceInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryConferenceInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据会议码查询视频会议信息
//
// @param tmpReq - QueryConferenceInfoByRoomCodeRequest
//
// @param tmpHeader - QueryConferenceInfoByRoomCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceInfoByRoomCodeResponse
func (client *Client) QueryConferenceInfoByRoomCodeWithContext(ctx context.Context, tmpReq *QueryConferenceInfoByRoomCodeRequest, tmpHeader *QueryConferenceInfoByRoomCodeHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceInfoByRoomCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryConferenceInfoByRoomCodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryConferenceInfoByRoomCodeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RoomCode) {
		body["roomCode"] = request.RoomCode
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConferenceInfoByRoomCode"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryConferenceInfoByRoomCode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConferenceInfoByRoomCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频会议成员
//
// @param tmpReq - QueryConferenceMembersRequest
//
// @param tmpHeader - QueryConferenceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceMembersResponse
func (client *Client) QueryConferenceMembersWithContext(ctx context.Context, tmpReq *QueryConferenceMembersRequest, tmpHeader *QueryConferenceMembersHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryConferenceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryConferenceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConferenceMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryConferenceMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件或文件夹信息
//
// @param tmpReq - QueryDentriesInfoRequest
//
// @param tmpHeader - QueryDentriesInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDentriesInfoResponse
func (client *Client) QueryDentriesInfoWithContext(ctx context.Context, tmpReq *QueryDentriesInfoRequest, tmpHeader *QueryDentriesInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryDentriesInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryDentriesInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryDentriesInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AppIdsForAppProperties) {
		request.AppIdsForAppPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppIdsForAppProperties, dara.String("AppIdsForAppProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppIdsForAppPropertiesShrink) {
		body["AppIdsForAppProperties"] = request.AppIdsForAppPropertiesShrink
	}

	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WithThumbnail) {
		body["WithThumbnail"] = request.WithThumbnail
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDentriesInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/queryDentriesInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDentriesInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询知识库节点信息
//
// @param tmpReq - QueryDentryRequest
//
// @param tmpHeader - QueryDentryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDentryResponse
func (client *Client) QueryDentryWithContext(ctx context.Context, tmpReq *QueryDentryRequest, tmpHeader *QueryDentryHeaders, runtime *dara.RuntimeOptions) (_result *QueryDentryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryDentryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryDentryShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryId) {
		body["DentryId"] = request.DentryId
	}

	if !dara.IsNil(request.IncludeSpace) {
		body["IncludeSpace"] = request.IncludeSpace
	}

	if !dara.IsNil(request.SpaceId) {
		body["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDentry"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v2/documents/queryDentry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDentryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群直播详情
//
// @param tmpReq - QueryGroupLiveInfoRequest
//
// @param tmpHeader - QueryGroupLiveInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupLiveInfoResponse
func (client *Client) QueryGroupLiveInfoWithContext(ctx context.Context, tmpReq *QueryGroupLiveInfoRequest, tmpHeader *QueryGroupLiveInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryGroupLiveInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryGroupLiveInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryGroupLiveInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AnchorUnionId) {
		body["AnchorUnionId"] = request.AnchorUnionId
	}

	if !dara.IsNil(request.LiveUuid) {
		body["LiveUuid"] = request.LiveUuid
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGroupLiveInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryGroupLiveInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGroupLiveInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播信息
//
// @param tmpReq - QueryLiveInfoRequest
//
// @param tmpHeader - QueryLiveInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveInfoResponse
func (client *Client) QueryLiveInfoWithContext(ctx context.Context, tmpReq *QueryLiveInfoRequest, tmpHeader *QueryLiveInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryLiveInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryLiveInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLiveInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryLiveInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLiveInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播的观看数据
//
// @param tmpReq - QueryLiveWatchDetailRequest
//
// @param tmpHeader - QueryLiveWatchDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveWatchDetailResponse
func (client *Client) QueryLiveWatchDetailWithContext(ctx context.Context, tmpReq *QueryLiveWatchDetailRequest, tmpHeader *QueryLiveWatchDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveWatchDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryLiveWatchDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryLiveWatchDetailShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLiveWatchDetail"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryLiveWatchDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLiveWatchDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播观看人员信息
//
// @param tmpReq - QueryLiveWatchUserListRequest
//
// @param tmpHeader - QueryLiveWatchUserListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveWatchUserListResponse
func (client *Client) QueryLiveWatchUserListWithContext(ctx context.Context, tmpReq *QueryLiveWatchUserListRequest, tmpHeader *QueryLiveWatchUserListHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveWatchUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryLiveWatchUserListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryLiveWatchUserListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLiveWatchUserList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryLiveWatchUserList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLiveWatchUserListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议室详情
//
// @param tmpReq - QueryMeetingRoomRequest
//
// @param tmpHeader - QueryMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomResponse
func (client *Client) QueryMeetingRoomWithContext(ctx context.Context, tmpReq *QueryMeetingRoomRequest, tmpHeader *QueryMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMeetingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMeetingRoomShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoomId) {
		body["RoomId"] = request.RoomId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMeetingRoom"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMeetingRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMeetingRoomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议室分组信息
//
// @param tmpReq - QueryMeetingRoomGroupRequest
//
// @param tmpHeader - QueryMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomGroupResponse
func (client *Client) QueryMeetingRoomGroupWithContext(ctx context.Context, tmpReq *QueryMeetingRoomGroupRequest, tmpHeader *QueryMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMeetingRoomGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMeetingRoomGroupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMeetingRoomGroup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMeetingRoomGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMeetingRoomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议室分组列表
//
// @param tmpReq - QueryMeetingRoomGroupListRequest
//
// @param tmpHeader - QueryMeetingRoomGroupListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomGroupListResponse
func (client *Client) QueryMeetingRoomGroupListWithContext(ctx context.Context, tmpReq *QueryMeetingRoomGroupListRequest, tmpHeader *QueryMeetingRoomGroupListHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMeetingRoomGroupListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMeetingRoomGroupListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMeetingRoomGroupList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMeetingRoomGroupList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMeetingRoomGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议室列表
//
// @param tmpReq - QueryMeetingRoomListRequest
//
// @param tmpHeader - QueryMeetingRoomListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomListResponse
func (client *Client) QueryMeetingRoomListWithContext(ctx context.Context, tmpReq *QueryMeetingRoomListRequest, tmpHeader *QueryMeetingRoomListHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMeetingRoomListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMeetingRoomListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMeetingRoomList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMeetingRoomList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMeetingRoomListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记录音
//
// @param tmpReq - QueryMinutesRequest
//
// @param tmpHeader - QueryMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesResponse
func (client *Client) QueryMinutesWithContext(ctx context.Context, tmpReq *QueryMinutesRequest, tmpHeader *QueryMinutesHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMinutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMinutesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMinutes"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMinutes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMinutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议闪记智能纪要
//
// @param tmpReq - QueryMinutesSummaryRequest
//
// @param tmpHeader - QueryMinutesSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesSummaryResponse
func (client *Client) QueryMinutesSummaryWithContext(ctx context.Context, tmpReq *QueryMinutesSummaryRequest, tmpHeader *QueryMinutesSummaryHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMinutesSummaryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMinutesSummaryShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SummaryTypeList) {
		request.SummaryTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SummaryTypeList, dara.String("summaryTypeList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.SummaryTypeListShrink) {
		body["summaryTypeList"] = request.SummaryTypeListShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMinutesSummary"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMinutesSummary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMinutesSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议闪记的文本信息
//
// @param tmpReq - QueryMinutesTextRequest
//
// @param tmpHeader - QueryMinutesTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesTextWithContext(ctx context.Context, tmpReq *QueryMinutesTextRequest, tmpHeader *QueryMinutesTextHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMinutesTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryMinutesTextShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.Direction) {
		body["direction"] = request.Direction
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMinutesText"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryMinutesText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业荣誉
//
// @param tmpReq - QueryOrgHonorsRequest
//
// @param tmpHeader - QueryOrgHonorsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgHonorsResponse
func (client *Client) QueryOrgHonorsWithContext(ctx context.Context, tmpReq *QueryOrgHonorsRequest, tmpHeader *QueryOrgHonorsHeaders, runtime *dara.RuntimeOptions) (_result *QueryOrgHonorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryOrgHonorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryOrgHonorsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOrgHonors"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/queryOrgHonors"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业代办
//
// @param tmpReq - QueryOrgTodoTasksRequest
//
// @param tmpHeader - QueryOrgTodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgTodoTasksResponse
func (client *Client) QueryOrgTodoTasksWithContext(ctx context.Context, tmpReq *QueryOrgTodoTasksRequest, tmpHeader *QueryOrgTodoTasksHeaders, runtime *dara.RuntimeOptions) (_result *QueryOrgTodoTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryOrgTodoTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryOrgTodoTasksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.IsDone) {
		body["isDone"] = request.IsDone
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOrgTodoTasks"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/queryOrgTodoTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记的链接，可通过链接查看闪记内容
//
// @param tmpReq - QueryRecordMinutesUrlRequest
//
// @param tmpHeader - QueryRecordMinutesUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordMinutesUrlResponse
func (client *Client) QueryRecordMinutesUrlWithContext(ctx context.Context, tmpReq *QueryRecordMinutesUrlRequest, tmpHeader *QueryRecordMinutesUrlHeaders, runtime *dara.RuntimeOptions) (_result *QueryRecordMinutesUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryRecordMinutesUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryRecordMinutesUrlShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ConferenceId) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecordMinutesUrl"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryRecordMinutesUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecordMinutesUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志详情
//
// @param tmpReq - QueryReportDetailRequest
//
// @param tmpHeader - QueryReportDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReportDetailResponse
func (client *Client) QueryReportDetailWithContext(ctx context.Context, tmpReq *QueryReportDetailRequest, tmpHeader *QueryReportDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryReportDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryReportDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryReportDetailShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReportDetail"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/queryReportDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReportDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询钉钉机器人退订工号
//
// @param request - QueryRobotUnsubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotUnsubscriptionResponse
func (client *Client) QueryRobotUnsubscriptionWithContext(ctx context.Context, request *QueryRobotUnsubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryRobotUnsubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RobotCode) {
		body["RobotCode"] = request.RobotCode
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotUnsubscription"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/robot/queryRobotUnsubscription"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotUnsubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预约会议
//
// @param tmpReq - QueryScheduleConferenceRequest
//
// @param tmpHeader - QueryScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConferenceResponse
func (client *Client) QueryScheduleConferenceWithContext(ctx context.Context, tmpReq *QueryScheduleConferenceRequest, tmpHeader *QueryScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *QueryScheduleConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryScheduleConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryScheduleConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ScheduleConferenceId) {
		body["scheduleConferenceId"] = request.ScheduleConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryScheduleConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryScheduleConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryScheduleConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预约会议历史会议信息
//
// @param tmpReq - QueryScheduleConferenceInfoRequest
//
// @param tmpHeader - QueryScheduleConferenceInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConferenceInfoResponse
func (client *Client) QueryScheduleConferenceInfoWithContext(ctx context.Context, tmpReq *QueryScheduleConferenceInfoRequest, tmpHeader *QueryScheduleConferenceInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryScheduleConferenceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryScheduleConferenceInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryScheduleConferenceInfoShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ScheduleConferenceId) {
		body["ScheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryScheduleConferenceInfo"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/queryScheduleConferenceInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryScheduleConferenceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询员工勋章列表
//
// @param tmpReq - QueryUserHonorsRequest
//
// @param tmpHeader - QueryUserHonorsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserHonorsResponse
func (client *Client) QueryUserHonorsWithContext(ctx context.Context, tmpReq *QueryUserHonorsRequest, tmpHeader *QueryUserHonorsHeaders, runtime *dara.RuntimeOptions) (_result *QueryUserHonorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryUserHonorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryUserHonorsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserHonors"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/queryUserHonors"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 收回勋章
//
// @param tmpReq - RecallHonorRequest
//
// @param tmpHeader - RecallHonorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallHonorResponse
func (client *Client) RecallHonorWithContext(ctx context.Context, tmpReq *RecallHonorRequest, tmpHeader *RecallHonorHeaders, runtime *dara.RuntimeOptions) (_result *RecallHonorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RecallHonorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &RecallHonorShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.HonorId) {
		body["honorId"] = request.HonorId
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecallHonor"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/recallHonor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecallHonorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志接收人员列表
//
// @param tmpReq - ReceiverListReportRequest
//
// @param tmpHeader - ReceiverListReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReceiverListReportResponse
func (client *Client) ReceiverListReportWithContext(ctx context.Context, tmpReq *ReceiverListReportRequest, tmpHeader *ReceiverListReportHeaders, runtime *dara.RuntimeOptions) (_result *ReceiverListReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReceiverListReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &ReceiverListReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReceiverListReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/receiverListReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReceiverListReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转交任务
//
// @param request - RedirectTaskRequest
//
// @param tmpHeader - RedirectTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedirectTaskResponse
func (client *Client) RedirectTaskWithContext(ctx context.Context, request *RedirectTaskRequest, tmpHeader *RedirectTaskHeaders, runtime *dara.RuntimeOptions) (_result *RedirectTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &RedirectTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ByManager) {
		body["ByManager"] = request.ByManager
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.NowActionExecutorId) {
		body["NowActionExecutorId"] = request.NowActionExecutorId
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedirectTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/redirectTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedirectTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除日程参与者
//
// @param tmpReq - RemoveAttendeeRequest
//
// @param tmpHeader - RemoveAttendeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAttendeeResponse
func (client *Client) RemoveAttendeeWithContext(ctx context.Context, tmpReq *RemoveAttendeeRequest, tmpHeader *RemoveAttendeeHeaders, runtime *dara.RuntimeOptions) (_result *RemoveAttendeeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveAttendeeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &RemoveAttendeeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AttendeesToRemove) {
		request.AttendeesToRemoveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttendeesToRemove, dara.String("AttendeesToRemove"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttendeesToRemoveShrink) {
		body["AttendeesToRemove"] = request.AttendeesToRemoveShrink
	}

	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAttendee"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/removeAttendee"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消预定会议室
//
// @param tmpReq - RemoveMeetingRoomsRequest
//
// @param tmpHeader - RemoveMeetingRoomsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMeetingRoomsResponse
func (client *Client) RemoveMeetingRoomsWithContext(ctx context.Context, tmpReq *RemoveMeetingRoomsRequest, tmpHeader *RemoveMeetingRoomsHeaders, runtime *dara.RuntimeOptions) (_result *RemoveMeetingRoomsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveMeetingRoomsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &RemoveMeetingRoomsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MeetingRoomsToRemove) {
		request.MeetingRoomsToRemoveShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MeetingRoomsToRemove, dara.String("MeetingRoomsToRemove"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.MeetingRoomsToRemoveShrink) {
		body["MeetingRoomsToRemove"] = request.MeetingRoomsToRemoveShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveMeetingRooms"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/removeMeetingRooms"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveMeetingRoomsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置日程响应邀请状态
//
// @param tmpReq - RespondEventRequest
//
// @param tmpHeader - RespondEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RespondEventResponse
func (client *Client) RespondEventWithContext(ctx context.Context, tmpReq *RespondEventRequest, tmpHeader *RespondEventHeaders, runtime *dara.RuntimeOptions) (_result *RespondEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RespondEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &RespondEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.EventId) {
		body["EventId"] = request.EventId
	}

	if !dara.IsNil(request.ResponseStatus) {
		body["ResponseStatus"] = request.ResponseStatus
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RespondEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/respondEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RespondEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询运行
//
// @param request - RetrieveRunRequest
//
// @param headers - RetrieveRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveRunResponse
func (client *Client) RetrieveRunWithContext(ctx context.Context, request *RetrieveRunRequest, headers *RetrieveRunHeaders, runtime *dara.RuntimeOptions) (_result *RetrieveRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveRun"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/retrieveRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存日志内容
//
// @param tmpReq - SaveContentRequest
//
// @param tmpHeader - SaveContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContentResponse
func (client *Client) SaveContentWithContext(ctx context.Context, tmpReq *SaveContentRequest, tmpHeader *SaveContentHeaders, runtime *dara.RuntimeOptions) (_result *SaveContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SaveContentShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Contents) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, dara.String("Contents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentsShrink) {
		body["Contents"] = request.ContentsShrink
	}

	if !dara.IsNil(request.DdFrom) {
		body["DdFrom"] = request.DdFrom
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveContent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/saveContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存表单数据
//
// @param request - SaveFormDataRequest
//
// @param tmpHeader - SaveFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFormDataResponse
func (client *Client) SaveFormDataWithContext(ctx context.Context, request *SaveFormDataRequest, tmpHeader *SaveFormDataHeaders, runtime *dara.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SaveFormDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormDataJson) {
		body["FormDataJson"] = request.FormDataJson
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveFormData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/saveFormData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveFormDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交表单或流程实例下的评论
//
// @param request - SaveFormRemarkRequest
//
// @param tmpHeader - SaveFormRemarkHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFormRemarkResponse
func (client *Client) SaveFormRemarkWithContext(ctx context.Context, request *SaveFormRemarkRequest, tmpHeader *SaveFormRemarkHeaders, runtime *dara.RuntimeOptions) (_result *SaveFormRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SaveFormRemarkShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.AtUserId) {
		body["AtUserId"] = request.AtUserId
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.FormInstanceId) {
		body["FormInstanceId"] = request.FormInstanceId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ReplyId) {
		body["ReplyId"] = request.ReplyId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveFormRemark"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/saveFormRemark"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveFormRemarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取员工组件的值
//
// @param request - SearchEmployeeFieldValuesRequest
//
// @param tmpHeader - SearchEmployeeFieldValuesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchEmployeeFieldValuesResponse
func (client *Client) SearchEmployeeFieldValuesWithContext(ctx context.Context, request *SearchEmployeeFieldValuesRequest, tmpHeader *SearchEmployeeFieldValuesHeaders, runtime *dara.RuntimeOptions) (_result *SearchEmployeeFieldValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchEmployeeFieldValuesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.SearchFieldJson) {
		body["SearchFieldJson"] = request.SearchFieldJson
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.TargetFieldJson) {
		body["TargetFieldJson"] = request.TargetFieldJson
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchEmployeeFieldValues"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/searchEmployeeFieldValues"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchEmployeeFieldValuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取多个表单实例ID
//
// @param request - SearchFormDataIdListRequest
//
// @param tmpHeader - SearchFormDataIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdListWithContext(ctx context.Context, request *SearchFormDataIdListRequest, tmpHeader *SearchFormDataIdListHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataIdListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchFormDataIdListShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchFieldJson) {
		body["SearchFieldJson"] = request.SearchFieldJson
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFormDataIdList"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/searchFormDataIdList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过高级查询条件获取表单实例数据（包括子表单组件数据）
//
// @param request - SearchFormDataSecondGenerationRequest
//
// @param tmpHeader - SearchFormDataSecondGenerationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGenerationWithContext(ctx context.Context, request *SearchFormDataSecondGenerationRequest, tmpHeader *SearchFormDataSecondGenerationHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchFormDataSecondGenerationShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OrderConfigJson) {
		body["OrderConfigJson"] = request.OrderConfigJson
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchCondition) {
		body["SearchCondition"] = request.SearchCondition
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFormDataSecondGeneration"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/searchFormDataSecondGeneration"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过高级查询条件获取表单实例数据（不包括子表单组件数据）
//
// @param request - SearchFormDataSecondGenerationNoTableFieldRequest
//
// @param tmpHeader - SearchFormDataSecondGenerationNoTableFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataSecondGenerationNoTableFieldResponse
func (client *Client) SearchFormDataSecondGenerationNoTableFieldWithContext(ctx context.Context, request *SearchFormDataSecondGenerationNoTableFieldRequest, tmpHeader *SearchFormDataSecondGenerationNoTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataSecondGenerationNoTableFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchFormDataSecondGenerationNoTableFieldShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OrderConfigJson) {
		body["OrderConfigJson"] = request.OrderConfigJson
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchCondition) {
		body["SearchCondition"] = request.SearchCondition
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFormDataSecondGenerationNoTableField"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/searchFormDataSecondGenerationNoTableField"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFormDataSecondGenerationNoTableFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表单实例数据
//
// @param request - SearchFormDatasRequest
//
// @param tmpHeader - SearchFormDatasHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDatasResponse
func (client *Client) SearchFormDatasWithContext(ctx context.Context, request *SearchFormDatasRequest, tmpHeader *SearchFormDatasHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchFormDatasShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.CreateFromTimeGMT) {
		body["CreateFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !dara.IsNil(request.CreateToTimeGMT) {
		body["CreateToTimeGMT"] = request.CreateToTimeGMT
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DynamicOrder) {
		body["DynamicOrder"] = request.DynamicOrder
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ModifiedFromTimeGMT) {
		body["ModifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !dara.IsNil(request.ModifiedToTimeGMT) {
		body["ModifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !dara.IsNil(request.OriginatorId) {
		body["OriginatorId"] = request.OriginatorId
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchFieldJson) {
		body["SearchFieldJson"] = request.SearchFieldJson
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFormDatas"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/searchFormDatas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据关键词搜索企业内部群
//
// @param request - SearchInnerGroupsRequest
//
// @param tmpHeader - SearchInnerGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInnerGroupsResponse
func (client *Client) SearchInnerGroupsWithContext(ctx context.Context, request *SearchInnerGroupsRequest, tmpHeader *SearchInnerGroupsHeaders, runtime *dara.RuntimeOptions) (_result *SearchInnerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SearchInnerGroupsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.SearchKey) {
		body["SearchKey"] = request.SearchKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchInnerGroups"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/im/searchInnerGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchInnerGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送钉钉Banner通知
//
// @param tmpReq - SendBannerRequest
//
// @param tmpHeader - SendBannerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBannerResponse
func (client *Client) SendBannerWithContext(ctx context.Context, tmpReq *SendBannerRequest, tmpHeader *SendBannerHeaders, runtime *dara.RuntimeOptions) (_result *SendBannerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendBannerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SendBannerShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendBanner"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/sendBanner"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendBannerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送钉钉封屏弹窗
//
// @param tmpReq - SendPopupRequest
//
// @param tmpHeader - SendPopupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendPopupResponse
func (client *Client) SendPopupWithContext(ctx context.Context, tmpReq *SendPopupRequest, tmpHeader *SendPopupHeaders, runtime *dara.RuntimeOptions) (_result *SendPopupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendPopupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SendPopupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendPopup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/sendPopup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendPopupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送钉钉搜索底纹
//
// @param tmpReq - SendSearchShadeRequest
//
// @param tmpHeader - SendSearchShadeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSearchShadeResponse
func (client *Client) SendSearchShadeWithContext(ctx context.Context, tmpReq *SendSearchShadeRequest, tmpHeader *SendSearchShadeHeaders, runtime *dara.RuntimeOptions) (_result *SendSearchShadeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendSearchShadeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SendSearchShadeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSearchShade"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/watt/sendSearchShade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSearchShadeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定列隐藏
//
// @param tmpReq - SetColumnsVisibilityRequest
//
// @param tmpHeader - SetColumnsVisibilityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetColumnsVisibilityResponse
func (client *Client) SetColumnsVisibilityWithContext(ctx context.Context, tmpReq *SetColumnsVisibilityRequest, tmpHeader *SetColumnsVisibilityHeaders, runtime *dara.RuntimeOptions) (_result *SetColumnsVisibilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetColumnsVisibilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SetColumnsVisibilityShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Column) {
		body["Column"] = request.Column
	}

	if !dara.IsNil(request.ColumnCount) {
		body["ColumnCount"] = request.ColumnCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetColumnsVisibility"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/setColumnsVisibility"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetColumnsVisibilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置联席主持人
//
// @param tmpReq - SetConferenceHostsRequest
//
// @param tmpHeader - SetConferenceHostsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetConferenceHostsResponse
func (client *Client) SetConferenceHostsWithContext(ctx context.Context, tmpReq *SetConferenceHostsRequest, tmpHeader *SetConferenceHostsHeaders, runtime *dara.RuntimeOptions) (_result *SetConferenceHostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetConferenceHostsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SetConferenceHostsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["UserIds"] = request.UserIdsShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetConferenceHosts"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/setConferenceHosts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetConferenceHostsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定行隐藏
//
// @param tmpReq - SetRowsVisibilityRequest
//
// @param tmpHeader - SetRowsVisibilityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRowsVisibilityResponse
func (client *Client) SetRowsVisibilityWithContext(ctx context.Context, tmpReq *SetRowsVisibilityRequest, tmpHeader *SetRowsVisibilityHeaders, runtime *dara.RuntimeOptions) (_result *SetRowsVisibilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetRowsVisibilityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SetRowsVisibilityShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Row) {
		body["Row"] = request.Row
	}

	if !dara.IsNil(request.RowCount) {
		body["RowCount"] = request.RowCount
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRowsVisibility"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/setRowsVisibility"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRowsVisibilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户发送日志的概要信息
//
// @param tmpReq - SimpleListReportRequest
//
// @param tmpHeader - SimpleListReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SimpleListReportResponse
func (client *Client) SimpleListReportWithContext(ctx context.Context, tmpReq *SimpleListReportRequest, tmpHeader *SimpleListReportHeaders, runtime *dara.RuntimeOptions) (_result *SimpleListReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SimpleListReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SimpleListReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SimpleListReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/simpleListReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SimpleListReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启视频会议云录制
//
// @param tmpReq - StartCloudRecordRequest
//
// @param tmpHeader - StartCloudRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecordWithContext(ctx context.Context, tmpReq *StartCloudRecordRequest, tmpHeader *StartCloudRecordHeaders, runtime *dara.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StartCloudRecordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.SmallWindowPosition) {
		body["SmallWindowPosition"] = request.SmallWindowPosition
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCloudRecord"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/startCloudRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起宜搭审批流程
//
// @param request - StartInstanceRequest
//
// @param tmpHeader - StartInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithContext(ctx context.Context, request *StartInstanceRequest, tmpHeader *StartInstanceHeaders, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &StartInstanceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DepartmentId) {
		body["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.FormDataJson) {
		body["FormDataJson"] = request.FormDataJson
	}

	if !dara.IsNil(request.FormUuid) {
		body["FormUuid"] = request.FormUuid
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessCode) {
		body["ProcessCode"] = request.ProcessCode
	}

	if !dara.IsNil(request.ProcessData) {
		body["ProcessData"] = request.ProcessData
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/startInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启闪记
//
// @param tmpReq - StartMinutesRequest
//
// @param tmpHeader - StartMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMinutesResponse
func (client *Client) StartMinutesWithContext(ctx context.Context, tmpReq *StartMinutesRequest, tmpHeader *StartMinutesHeaders, runtime *dara.RuntimeOptions) (_result *StartMinutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartMinutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StartMinutesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	if !dara.IsNil(request.OwnerUserId) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.RecordAudio) {
		body["recordAudio"] = request.RecordAudio
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMinutes"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/startMinutes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMinutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志相关人员列表
//
// @param tmpReq - StatisticsListByTypeReportRequest
//
// @param tmpHeader - StatisticsListByTypeReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StatisticsListByTypeReportResponse
func (client *Client) StatisticsListByTypeReportWithContext(ctx context.Context, tmpReq *StatisticsListByTypeReportRequest, tmpHeader *StatisticsListByTypeReportHeaders, runtime *dara.RuntimeOptions) (_result *StatisticsListByTypeReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StatisticsListByTypeReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StatisticsListByTypeReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StatisticsListByTypeReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/statisticsListByTypeReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StatisticsListByTypeReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志统计数据
//
// @param tmpReq - StatisticsReportRequest
//
// @param tmpHeader - StatisticsReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StatisticsReportResponse
func (client *Client) StatisticsReportWithContext(ctx context.Context, tmpReq *StatisticsReportRequest, tmpHeader *StatisticsReportHeaders, runtime *dara.RuntimeOptions) (_result *StatisticsReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StatisticsReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StatisticsReportShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		body["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StatisticsReport"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/log/statisticsReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StatisticsReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止视频会议云录制
//
// @param tmpReq - StopCloudRecordRequest
//
// @param tmpHeader - StopCloudRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecordWithContext(ctx context.Context, tmpReq *StopCloudRecordRequest, tmpHeader *StopCloudRecordHeaders, runtime *dara.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StopCloudRecordShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCloudRecord"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/stopCloudRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停闪记
//
// @param tmpReq - StopMinutesRequest
//
// @param tmpHeader - StopMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMinutesResponse
func (client *Client) StopMinutesWithContext(ctx context.Context, tmpReq *StopMinutesRequest, tmpHeader *StopMinutesHeaders, runtime *dara.RuntimeOptions) (_result *StopMinutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopMinutesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &StopMinutesShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMinutes"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/stopMinutes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMinutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅公共日历
//
// @param request - SubscribeCalendarRequest
//
// @param tmpHeader - SubscribeCalendarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeCalendarResponse
func (client *Client) SubscribeCalendarWithContext(ctx context.Context, request *SubscribeCalendarRequest, tmpHeader *SubscribeCalendarHeaders, runtime *dara.RuntimeOptions) (_result *SubscribeCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &SubscribeCalendarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubscribeCalendar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/subscribeCalendar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscribeCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅文件变更事件
//
// @param tmpReq - SubscribeEventRequest
//
// @param tmpHeader - SubscribeEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeEventResponse
func (client *Client) SubscribeEventWithContext(ctx context.Context, tmpReq *SubscribeEventRequest, tmpHeader *SubscribeEventHeaders, runtime *dara.RuntimeOptions) (_result *SubscribeEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubscribeEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SubscribeEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeId) {
		body["ScopeId"] = request.ScopeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubscribeEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/subscribeEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscribeEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - SwitchMainOrgRequest
//
// @param tmpHeader - SwitchMainOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchMainOrgResponse
func (client *Client) SwitchMainOrgWithContext(ctx context.Context, tmpReq *SwitchMainOrgRequest, tmpHeader *SwitchMainOrgHeaders, runtime *dara.RuntimeOptions) (_result *SwitchMainOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SwitchMainOrgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SwitchMainOrgShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TargetOrgId) {
		body["TargetOrgId"] = request.TargetOrgId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchMainOrg"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/user/switchMainOrg"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchMainOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步钉钉账号类型
//
// @param tmpReq - SyncDingTypeRequest
//
// @param tmpHeader - SyncDingTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDingTypeResponse
func (client *Client) SyncDingTypeWithContext(ctx context.Context, tmpReq *SyncDingTypeRequest, tmpHeader *SyncDingTypeHeaders, runtime *dara.RuntimeOptions) (_result *SyncDingTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncDingTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &SyncDingTypeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DingType) {
		body["DingType"] = request.DingType
	}

	if !dara.IsNil(request.IsDimission) {
		body["IsDimission"] = request.IsDimission
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkNo) {
		body["WorkNo"] = request.WorkNo
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDingType"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/indepding/syncDingType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDingTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止流程实例
//
// @param request - TerminateInstanceRequest
//
// @param tmpHeader - TerminateInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateInstanceResponse
func (client *Client) TerminateInstanceWithContext(ctx context.Context, request *TerminateInstanceRequest, tmpHeader *TerminateInstanceHeaders, runtime *dara.RuntimeOptions) (_result *TerminateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &TerminateInstanceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateInstance"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/terminateInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转交工单
//
// @param tmpReq - TransferTicketRequest
//
// @param tmpHeader - TransferTicketHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferTicketResponse
func (client *Client) TransferTicketWithContext(ctx context.Context, tmpReq *TransferTicketRequest, tmpHeader *TransferTicketHeaders, runtime *dara.RuntimeOptions) (_result *TransferTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TransferTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &TransferTicketShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notify) {
		request.NotifyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notify, dara.String("Notify"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProcessorUserIds) {
		request.ProcessorUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProcessorUserIds, dara.String("ProcessorUserIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TicketMemo) {
		request.TicketMemoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketMemo, dara.String("TicketMemo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NotifyShrink) {
		body["Notify"] = request.NotifyShrink
	}

	if !dara.IsNil(request.OpenTeamId) {
		body["OpenTeamId"] = request.OpenTeamId
	}

	if !dara.IsNil(request.OpenTicketId) {
		body["OpenTicketId"] = request.OpenTicketId
	}

	if !dara.IsNil(request.ProcessorUserIdsShrink) {
		body["ProcessorUserIds"] = request.ProcessorUserIdsShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.TicketMemoShrink) {
		body["TicketMemo"] = request.TicketMemoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferTicket"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ticket/transferTicket"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferTicketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消订阅公共日历
//
// @param request - UnsubscribeCalendarRequest
//
// @param tmpHeader - UnsubscribeCalendarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeCalendarResponse
func (client *Client) UnsubscribeCalendarWithContext(ctx context.Context, request *UnsubscribeCalendarRequest, tmpHeader *UnsubscribeCalendarHeaders, runtime *dara.RuntimeOptions) (_result *UnsubscribeCalendarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &UnsubscribeCalendarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnsubscribeCalendar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/unsubscribeCalendar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnsubscribeCalendarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消订阅文件变更事件
//
// @param tmpReq - UnsubscribeEventRequest
//
// @param tmpHeader - UnsubscribeEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeEventResponse
func (client *Client) UnsubscribeEventWithContext(ctx context.Context, tmpReq *UnsubscribeEventRequest, tmpHeader *UnsubscribeEventHeaders, runtime *dara.RuntimeOptions) (_result *UnsubscribeEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnsubscribeEventShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UnsubscribeEventShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.ScopeId) {
		body["ScopeId"] = request.ScopeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnsubscribeEvent"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/unsubscribeEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - UpdateAlidingAssistantRequest
//
// @param tmpHeader - UpdateAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlidingAssistantResponse
func (client *Client) UpdateAlidingAssistantWithContext(ctx context.Context, tmpReq *UpdateAlidingAssistantRequest, tmpHeader *UpdateAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *UpdateAlidingAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAlidingAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateAlidingAssistantShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Feature) {
		request.FeatureShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Feature, dara.String("Feature"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecommendPrompts) {
		request.RecommendPromptsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecommendPrompts, dara.String("RecommendPrompts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["AssistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExtShrink) {
		body["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.FallbackContent) {
		body["FallbackContent"] = request.FallbackContent
	}

	if !dara.IsNil(request.FeatureShrink) {
		body["Feature"] = request.FeatureShrink
	}

	if !dara.IsNil(request.Icon) {
		body["Icon"] = request.Icon
	}

	if !dara.IsNil(request.Instructions) {
		body["Instructions"] = request.Instructions
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecommendPromptsShrink) {
		body["RecommendPrompts"] = request.RecommendPromptsShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WelcomeContent) {
		body["WelcomeContent"] = request.WelcomeContent
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlidingAssistant"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/aiagent/updateAlidingAssistant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlidingAssistantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - UpdateConvExtensionRequest
//
// @param tmpHeader - UpdateConvExtensionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConvExtensionResponse
func (client *Client) UpdateConvExtensionWithContext(ctx context.Context, tmpReq *UpdateConvExtensionRequest, tmpHeader *UpdateConvExtensionHeaders, runtime *dara.RuntimeOptions) (_result *UpdateConvExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateConvExtensionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateConvExtensionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StaffIdList) {
		request.StaffIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StaffIdList, dara.String("StaffIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MobileUrl) {
		body["MobileUrl"] = request.MobileUrl
	}

	if !dara.IsNil(request.PcUrl) {
		body["PcUrl"] = request.PcUrl
	}

	if !dara.IsNil(request.StaffIdListShrink) {
		body["StaffIdList"] = request.StaffIdListShrink
	}

	if !dara.IsNil(request.SystemUid) {
		body["SystemUid"] = request.SystemUid
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConvExtension"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/conversation/updateConvExtension"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConvExtensionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新表单数据
//
// @param request - UpdateFormDataRequest
//
// @param tmpHeader - UpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFormDataResponse
func (client *Client) UpdateFormDataWithContext(ctx context.Context, request *UpdateFormDataRequest, tmpHeader *UpdateFormDataHeaders, runtime *dara.RuntimeOptions) (_result *UpdateFormDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &UpdateFormDataShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.FormInstanceId) {
		body["FormInstanceId"] = request.FormInstanceId
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.UpdateFormDataJson) {
		body["UpdateFormDataJson"] = request.UpdateFormDataJson
	}

	if !dara.IsNil(request.UseLatestVersion) {
		body["UseLatestVersion"] = request.UseLatestVersion
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFormData"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/updateFormData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流程实例
//
// @param request - UpdateInstanceRequest
//
// @param tmpHeader - UpdateInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest, tmpHeader *UpdateInstanceHeaders, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &UpdateInstanceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	if !dara.IsNil(request.UpdateFormDataJson) {
		body["UpdateFormDataJson"] = request.UpdateFormDataJson
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/updateInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改直播属性信息
//
// @param tmpReq - UpdateLiveRequest
//
// @param tmpHeader - UpdateLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveResponse
func (client *Client) UpdateLiveWithContext(ctx context.Context, tmpReq *UpdateLiveRequest, tmpHeader *UpdateLiveHeaders, runtime *dara.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLiveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateLiveShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverUrl) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !dara.IsNil(request.Introduction) {
		body["Introduction"] = request.Introduction
	}

	if !dara.IsNil(request.LiveId) {
		body["LiveId"] = request.LiveId
	}

	if !dara.IsNil(request.PreEndTime) {
		body["PreEndTime"] = request.PreEndTime
	}

	if !dara.IsNil(request.PreStartTime) {
		body["PreStartTime"] = request.PreStartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLive"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateLive"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会议室信息
//
// @param tmpReq - UpdateMeetingRoomRequest
//
// @param tmpHeader - UpdateMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeetingRoomResponse
func (client *Client) UpdateMeetingRoomWithContext(ctx context.Context, tmpReq *UpdateMeetingRoomRequest, tmpHeader *UpdateMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMeetingRoomResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMeetingRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateMeetingRoomShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReservationAuthority) {
		request.ReservationAuthorityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReservationAuthority, dara.String("ReservationAuthority"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoomLabelIds) {
		request.RoomLabelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomLabelIds, dara.String("RoomLabelIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RoomLocation) {
		request.RoomLocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomLocation, dara.String("RoomLocation"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableCycleReservation) {
		body["EnableCycleReservation"] = request.EnableCycleReservation
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsvRoomId) {
		body["IsvRoomId"] = request.IsvRoomId
	}

	if !dara.IsNil(request.ReservationAuthorityShrink) {
		body["ReservationAuthority"] = request.ReservationAuthorityShrink
	}

	if !dara.IsNil(request.RoomCapacity) {
		body["RoomCapacity"] = request.RoomCapacity
	}

	if !dara.IsNil(request.RoomId) {
		body["RoomId"] = request.RoomId
	}

	if !dara.IsNil(request.RoomLabelIdsShrink) {
		body["RoomLabelIds"] = request.RoomLabelIdsShrink
	}

	if !dara.IsNil(request.RoomLocationShrink) {
		body["RoomLocation"] = request.RoomLocationShrink
	}

	if !dara.IsNil(request.RoomName) {
		body["RoomName"] = request.RoomName
	}

	if !dara.IsNil(request.RoomPicture) {
		body["RoomPicture"] = request.RoomPicture
	}

	if !dara.IsNil(request.RoomStatus) {
		body["RoomStatus"] = request.RoomStatus
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMeetingRoom"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateMeetingRoom"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMeetingRoomResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会议室分组信息
//
// @param tmpReq - UpdateMeetingRoomGroupRequest
//
// @param tmpHeader - UpdateMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeetingRoomGroupResponse
func (client *Client) UpdateMeetingRoomGroupWithContext(ctx context.Context, tmpReq *UpdateMeetingRoomGroupRequest, tmpHeader *UpdateMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMeetingRoomGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMeetingRoomGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateMeetingRoomGroupShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMeetingRoomGroup"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateMeetingRoomGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMeetingRoomGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据表
//
// @param tmpReq - UpdateMultiDimTableRequest
//
// @param tmpHeader - UpdateMultiDimTableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiDimTableResponse
func (client *Client) UpdateMultiDimTableWithContext(ctx context.Context, tmpReq *UpdateMultiDimTableRequest, tmpHeader *UpdateMultiDimTableHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMultiDimTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateMultiDimTableShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultiDimTable"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/updateMultiDimTable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultiDimTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新字段
//
// @param tmpReq - UpdateMultiDimTableFieldRequest
//
// @param tmpHeader - UpdateMultiDimTableFieldHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiDimTableFieldResponse
func (client *Client) UpdateMultiDimTableFieldWithContext(ctx context.Context, tmpReq *UpdateMultiDimTableFieldRequest, tmpHeader *UpdateMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMultiDimTableFieldShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateMultiDimTableFieldShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Property) {
		request.PropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Property, dara.String("Property"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.FieldIdOrName) {
		body["FieldIdOrName"] = request.FieldIdOrName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PropertyShrink) {
		body["Property"] = request.PropertyShrink
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultiDimTableField"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/updateMultiDimTableField"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultiDimTableFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新多行记录
//
// @param tmpReq - UpdateMultiDimTableRecordsRequest
//
// @param tmpHeader - UpdateMultiDimTableRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultiDimTableRecordsResponse
func (client *Client) UpdateMultiDimTableRecordsWithContext(ctx context.Context, tmpReq *UpdateMultiDimTableRecordsRequest, tmpHeader *UpdateMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMultiDimTableRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateMultiDimTableRecordsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecordIds) {
		request.RecordIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordIds, dara.String("RecordIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseId) {
		body["BaseId"] = request.BaseId
	}

	if !dara.IsNil(request.RecordIdsShrink) {
		body["RecordIds"] = request.RecordIdsShrink
	}

	if !dara.IsNil(request.SheetIdOrName) {
		body["SheetIdOrName"] = request.SheetIdOrName
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultiDimTableRecords"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/table/updateMultiDimTableRecords"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultiDimTableRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文件权限
//
// @param tmpReq - UpdatePermissionRequest
//
// @param tmpHeader - UpdatePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermissionWithContext(ctx context.Context, tmpReq *UpdatePermissionRequest, tmpHeader *UpdatePermissionHeaders, runtime *dara.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdatePermissionShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Option) {
		request.OptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Option, dara.String("Option"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DentryUuid) {
		body["DentryUuid"] = request.DentryUuid
	}

	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.OptionShrink) {
		body["Option"] = request.OptionShrink
	}

	if !dara.IsNil(request.RoleId) {
		body["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePermission"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/updatePermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新单元格区域
//
// @param tmpReq - UpdateRangeRequest
//
// @param tmpHeader - UpdateRangeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRangeResponse
func (client *Client) UpdateRangeWithContext(ctx context.Context, tmpReq *UpdateRangeRequest, tmpHeader *UpdateRangeHeaders, runtime *dara.RuntimeOptions) (_result *UpdateRangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateRangeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BackgroundColors) {
		request.BackgroundColorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackgroundColors, dara.String("BackgroundColors"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hyperlinks) {
		request.HyperlinksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hyperlinks, dara.String("Hyperlinks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundColorsShrink) {
		body["BackgroundColors"] = request.BackgroundColorsShrink
	}

	if !dara.IsNil(request.HyperlinksShrink) {
		body["Hyperlinks"] = request.HyperlinksShrink
	}

	if !dara.IsNil(request.NumberFormat) {
		body["NumberFormat"] = request.NumberFormat
	}

	if !dara.IsNil(request.RangeAddress) {
		body["RangeAddress"] = request.RangeAddress
	}

	if !dara.IsNil(request.SheetId) {
		body["SheetId"] = request.SheetId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ValuesShrink) {
		body["Values"] = request.ValuesShrink
	}

	if !dara.IsNil(request.WorkbookId) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRange"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/updateRange"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRangeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新预约会议设置
//
// @param tmpReq - UpdateScheduleConfSettingsRequest
//
// @param tmpHeader - UpdateScheduleConfSettingsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleConfSettingsResponse
func (client *Client) UpdateScheduleConfSettingsWithContext(ctx context.Context, tmpReq *UpdateScheduleConfSettingsRequest, tmpHeader *UpdateScheduleConfSettingsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduleConfSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateScheduleConfSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateScheduleConfSettingsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfSettingModel) {
		request.ScheduleConfSettingModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfSettingModel, dara.String("ScheduleConfSettingModel"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ScheduleConfSettingModelShrink) {
		body["ScheduleConfSettingModel"] = request.ScheduleConfSettingModelShrink
	}

	if !dara.IsNil(request.ScheduleConferenceId) {
		body["ScheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduleConfSettings"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateScheduleConfSettings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduleConfSettingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新预约会议
//
// @param tmpReq - UpdateScheduleConferenceRequest
//
// @param tmpHeader - UpdateScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleConferenceResponse
func (client *Client) UpdateScheduleConferenceWithContext(ctx context.Context, tmpReq *UpdateScheduleConferenceRequest, tmpHeader *UpdateScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduleConferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateScheduleConferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateScheduleConferenceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ScheduleConferenceId) {
		body["ScheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduleConference"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateScheduleConference"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduleConferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新状态
//
// @param tmpReq - UpdateStatusRequest
//
// @param tmpHeader - UpdateStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStatusResponse
func (client *Client) UpdateStatusWithContext(ctx context.Context, tmpReq *UpdateStatusRequest, tmpHeader *UpdateStatusHeaders, runtime *dara.RuntimeOptions) (_result *UpdateStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateStatusShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ErrorLines) {
		request.ErrorLinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ErrorLines, dara.String("ErrorLines"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		body["AppType"] = request.AppType
	}

	if !dara.IsNil(request.ErrorLinesShrink) {
		body["ErrorLines"] = request.ErrorLinesShrink
	}

	if !dara.IsNil(request.ImportSequence) {
		body["ImportSequence"] = request.ImportSequence
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.SystemToken) {
		body["SystemToken"] = request.SystemToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStatus"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/yida/updateStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新订阅日历
//
// @param tmpReq - UpdateSubscribedCalendarsRequest
//
// @param tmpHeader - UpdateSubscribedCalendarsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscribedCalendarsResponse
func (client *Client) UpdateSubscribedCalendarsWithContext(ctx context.Context, tmpReq *UpdateSubscribedCalendarsRequest, tmpHeader *UpdateSubscribedCalendarsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateSubscribedCalendarsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSubscribedCalendarsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateSubscribedCalendarsShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubscribeScope) {
		request.SubscribeScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubscribeScope, dara.String("SubscribeScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CalendarId) {
		body["CalendarId"] = request.CalendarId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ManagersShrink) {
		body["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SubscribeScopeShrink) {
		body["SubscribeScope"] = request.SubscribeScopeShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubscribedCalendars"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/calendar/updateSubscribedCalendars"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubscribedCalendarsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代办
//
// @param tmpReq - UpdateTodoTaskRequest
//
// @param tmpHeader - UpdateTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTodoTaskResponse
func (client *Client) UpdateTodoTaskWithContext(ctx context.Context, tmpReq *UpdateTodoTaskRequest, tmpHeader *UpdateTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutorIds) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, dara.String("executorIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ParticipantIds) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, dara.String("participantIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Done) {
		body["done"] = request.Done
	}

	if !dara.IsNil(request.DueTime) {
		body["dueTime"] = request.DueTime
	}

	if !dara.IsNil(request.ExecutorIdsShrink) {
		body["executorIds"] = request.ExecutorIdsShrink
	}

	if !dara.IsNil(request.ParticipantIdsShrink) {
		body["participantIds"] = request.ParticipantIdsShrink
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTodoTask"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/updateTodoTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代办执行者状态
//
// @param tmpReq - UpdateTodoTaskExecutorStatusRequest
//
// @param tmpHeader - UpdateTodoTaskExecutorStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTodoTaskExecutorStatusResponse
func (client *Client) UpdateTodoTaskExecutorStatusWithContext(ctx context.Context, tmpReq *UpdateTodoTaskExecutorStatusRequest, tmpHeader *UpdateTodoTaskExecutorStatusHeaders, runtime *dara.RuntimeOptions) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTodoTaskExecutorStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateTodoTaskExecutorStatusShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutorStatusList) {
		request.ExecutorStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorStatusList, dara.String("executorStatusList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ExecutorStatusListShrink) {
		body["executorStatusList"] = request.ExecutorStatusListShrink
	}

	if !dara.IsNil(request.OperatorId) {
		body["operatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTodoTaskExecutorStatus"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/task/updateTodoTaskExecutorStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新企业账号用户头像
//
// @param request - UpdateUserAvatarRequest
//
// @param tmpHeader - UpdateUserAvatarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserAvatarResponse
func (client *Client) UpdateUserAvatarWithContext(ctx context.Context, request *UpdateUserAvatarRequest, tmpHeader *UpdateUserAvatarHeaders, runtime *dara.RuntimeOptions) (_result *UpdateUserAvatarResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	headers := &UpdateUserAvatarShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AvatarMediaId) {
		body["AvatarMediaId"] = request.AvatarMediaId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserAvatar"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/contact/updateUserAvatar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserAvatarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置正在进行中的视频会议属性
//
// @param tmpReq - UpdateVideoConferenceSettingRequest
//
// @param tmpHeader - UpdateVideoConferenceSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoConferenceSettingResponse
func (client *Client) UpdateVideoConferenceSettingWithContext(ctx context.Context, tmpReq *UpdateVideoConferenceSettingRequest, tmpHeader *UpdateVideoConferenceSettingHeaders, runtime *dara.RuntimeOptions) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVideoConferenceSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateVideoConferenceSettingShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowUnmuteSelf) {
		body["AllowUnmuteSelf"] = request.AllowUnmuteSelf
	}

	if !dara.IsNil(request.AutoTransferHost) {
		body["AutoTransferHost"] = request.AutoTransferHost
	}

	if !dara.IsNil(request.ForbiddenShareScreen) {
		body["ForbiddenShareScreen"] = request.ForbiddenShareScreen
	}

	if !dara.IsNil(request.LockConference) {
		body["LockConference"] = request.LockConference
	}

	if !dara.IsNil(request.MuteAll) {
		body["MuteAll"] = request.MuteAll
	}

	if !dara.IsNil(request.OnlyInternalEmployeesJoin) {
		body["OnlyInternalEmployeesJoin"] = request.OnlyInternalEmployeesJoin
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.ConferenceId) {
		body["conferenceId"] = request.ConferenceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVideoConferenceSetting"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/ysp/updateVideoConferenceSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改知识库文档成员权限
//
// @param tmpReq - UpdateWorkspaceDocMembersRequest
//
// @param tmpHeader - UpdateWorkspaceDocMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceDocMembersResponse
func (client *Client) UpdateWorkspaceDocMembersWithContext(ctx context.Context, tmpReq *UpdateWorkspaceDocMembersRequest, tmpHeader *UpdateWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspaceDocMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/updateWorkspaceDocMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新知识库成员权限
//
// @param tmpReq - UpdateWorkspaceMembersRequest
//
// @param tmpHeader - UpdateWorkspaceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceMembersResponse
func (client *Client) UpdateWorkspaceMembersWithContext(ctx context.Context, tmpReq *UpdateWorkspaceMembersRequest, tmpHeader *UpdateWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		body["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspaceMembers"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dingtalk/v1/documents/updateWorkspaceMembers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传媒体<br/>
//
// @param tmpReq - UploadMediaRequest
//
// @param tmpHeader - UploadMediaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMediaResponse
func (client *Client) UploadMediaWithContext(ctx context.Context, tmpReq *UploadMediaRequest, tmpHeader *UploadMediaHeaders, runtime *dara.RuntimeOptions) (_result *UploadMediaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadMediaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UploadMediaShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.MediaName) {
		body["mediaName"] = request.MediaName
	}

	if !dara.IsNil(request.MediaType) {
		body["mediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMedia"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/documents/uploadMedia"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMediaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 穿戴勋章
//
// @param tmpReq - WearOrgHonorRequest
//
// @param tmpHeader - WearOrgHonorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WearOrgHonorResponse
func (client *Client) WearOrgHonorWithContext(ctx context.Context, tmpReq *WearOrgHonorRequest, tmpHeader *WearOrgHonorHeaders, runtime *dara.RuntimeOptions) (_result *WearOrgHonorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &WearOrgHonorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &WearOrgHonorShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantContext) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, dara.String("TenantContext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantContextShrink) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !dara.IsNil(request.HonorId) {
		body["honorId"] = request.HonorId
	}

	if !dara.IsNil(request.OrgId) {
		body["orgId"] = request.OrgId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	if !dara.IsNil(request.Wear) {
		body["wear"] = request.Wear
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WearOrgHonor"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aliding/v1/honor/wearOrgHonor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WearOrgHonorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) createRunWithSSECtx_opYieldFunc(_yield chan *CreateRunResponse, _yieldErr chan error, ctx context.Context, request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowStructViewContent) {
		body["allowStructViewContent"] = request.AllowStructViewContent
	}

	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRun"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/createRun"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) invokeAssistantWithSSECtx_opYieldFunc(_yield chan *InvokeAssistantResponse, _yieldErr chan error, ctx context.Context, request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssistantId) {
		body["assistantId"] = request.AssistantId
	}

	if !dara.IsNil(request.ClientEnum) {
		body["clientEnum"] = request.ClientEnum
	}

	if !dara.IsNil(request.ExtLoginUser) {
		body["extLoginUser"] = request.ExtLoginUser
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.OriginalAssistantId) {
		body["originalAssistantId"] = request.OriginalAssistantId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SourceIdOfOriginalAssistantId) {
		body["sourceIdOfOriginalAssistantId"] = request.SourceIdOfOriginalAssistantId
	}

	if !dara.IsNil(request.SourceTypeOfOriginalAssistantId) {
		body["sourceTypeOfOriginalAssistantId"] = request.SourceTypeOfOriginalAssistantId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountId) {
		realHeaders["accountId"] = dara.String(dara.ToString(dara.StringValue(headers.AccountId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAssistant"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/assistant/invokeAssistant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) invokeSkillWithSSECtx_opYieldFunc(_yield chan *InvokeSkillResponse, _yieldErr chan error, ctx context.Context, tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &InvokeSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InvokeSkillShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !dara.IsNil(tmpHeader.AccountContext) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, dara.String("AccountContext"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.SkillId) {
		body["SkillId"] = request.SkillId
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.SourceIdOfAssistantId) {
		body["sourceIdOfAssistantId"] = request.SourceIdOfAssistantId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.AccountContextShrink) {
		realHeaders["AccountContext"] = dara.String(dara.Stringify(dara.StringValue(headers.AccountContextShrink)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeSkill"),
		Version:     dara.String("2023-04-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ai/v1/skill/invoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
