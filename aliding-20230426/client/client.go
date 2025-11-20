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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aliding"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加日程参与者
//
// @param tmpReq - AddAttendeeRequest
//
// @param tmpHeader - AddAttendeeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAttendeeResponse
func (client *Client) AddAttendeeWithOptions(tmpReq *AddAttendeeRequest, tmpHeader *AddAttendeeHeaders, runtime *dara.RuntimeOptions) (_result *AddAttendeeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加日程参与者
//
// @param request - AddAttendeeRequest
//
// @return AddAttendeeResponse
func (client *Client) AddAttendee(request *AddAttendeeRequest) (_result *AddAttendeeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddAttendeeHeaders{}
	_result = &AddAttendeeResponse{}
	_body, _err := client.AddAttendeeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddDriveSpaceWithOptions(tmpReq *AddDriveSpaceRequest, tmpHeader *AddDriveSpaceHeaders, runtime *dara.RuntimeOptions) (_result *AddDriveSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddDriveSpaceRequest
//
// @return AddDriveSpaceResponse
func (client *Client) AddDriveSpace(request *AddDriveSpaceRequest) (_result *AddDriveSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddDriveSpaceHeaders{}
	_result = &AddDriveSpaceResponse{}
	_body, _err := client.AddDriveSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddFolderWithOptions(tmpReq *AddFolderRequest, tmpHeader *AddFolderHeaders, runtime *dara.RuntimeOptions) (_result *AddFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddFolderRequest
//
// @return AddFolderResponse
func (client *Client) AddFolder(request *AddFolderRequest) (_result *AddFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddFolderHeaders{}
	_result = &AddFolderResponse{}
	_body, _err := client.AddFolderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddMeetingRoomsWithOptions(tmpReq *AddMeetingRoomsRequest, tmpHeader *AddMeetingRoomsHeaders, runtime *dara.RuntimeOptions) (_result *AddMeetingRoomsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddMeetingRoomsRequest
//
// @return AddMeetingRoomsResponse
func (client *Client) AddMeetingRooms(request *AddMeetingRoomsRequest) (_result *AddMeetingRoomsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddMeetingRoomsHeaders{}
	_result = &AddMeetingRoomsResponse{}
	_body, _err := client.AddMeetingRoomsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddMultiDimTableWithOptions(tmpReq *AddMultiDimTableRequest, tmpHeader *AddMultiDimTableHeaders, runtime *dara.RuntimeOptions) (_result *AddMultiDimTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddMultiDimTableRequest
//
// @return AddMultiDimTableResponse
func (client *Client) AddMultiDimTable(request *AddMultiDimTableRequest) (_result *AddMultiDimTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddMultiDimTableHeaders{}
	_result = &AddMultiDimTableResponse{}
	_body, _err := client.AddMultiDimTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddPermissionWithOptions(tmpReq *AddPermissionRequest, tmpHeader *AddPermissionHeaders, runtime *dara.RuntimeOptions) (_result *AddPermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddPermissionRequest
//
// @return AddPermissionResponse
func (client *Client) AddPermission(request *AddPermissionRequest) (_result *AddPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddPermissionHeaders{}
	_result = &AddPermissionResponse{}
	_body, _err := client.AddPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddRecordPermissionWithOptions(tmpReq *AddRecordPermissionRequest, tmpHeader *AddRecordPermissionHeaders, runtime *dara.RuntimeOptions) (_result *AddRecordPermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddRecordPermissionRequest
//
// @return AddRecordPermissionResponse
func (client *Client) AddRecordPermission(request *AddRecordPermissionRequest) (_result *AddRecordPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddRecordPermissionHeaders{}
	_result = &AddRecordPermissionResponse{}
	_body, _err := client.AddRecordPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddScenegroupMemberWithOptions(request *AddScenegroupMemberRequest, tmpHeader *AddScenegroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *AddScenegroupMemberResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddScenegroupMemberResponse
func (client *Client) AddScenegroupMember(request *AddScenegroupMemberRequest) (_result *AddScenegroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddScenegroupMemberHeaders{}
	_result = &AddScenegroupMemberResponse{}
	_body, _err := client.AddScenegroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddTicketMemoWithOptions(tmpReq *AddTicketMemoRequest, tmpHeader *AddTicketMemoHeaders, runtime *dara.RuntimeOptions) (_result *AddTicketMemoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddTicketMemoRequest
//
// @return AddTicketMemoResponse
func (client *Client) AddTicketMemo(request *AddTicketMemoRequest) (_result *AddTicketMemoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddTicketMemoHeaders{}
	_result = &AddTicketMemoResponse{}
	_body, _err := client.AddTicketMemoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddWorkspaceWithOptions(tmpReq *AddWorkspaceRequest, tmpHeader *AddWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddWorkspaceRequest
//
// @return AddWorkspaceResponse
func (client *Client) AddWorkspace(request *AddWorkspaceRequest) (_result *AddWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddWorkspaceHeaders{}
	_result = &AddWorkspaceResponse{}
	_body, _err := client.AddWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddWorkspaceDocMembersWithOptions(tmpReq *AddWorkspaceDocMembersRequest, tmpHeader *AddWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceDocMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddWorkspaceDocMembersRequest
//
// @return AddWorkspaceDocMembersResponse
func (client *Client) AddWorkspaceDocMembers(request *AddWorkspaceDocMembersRequest) (_result *AddWorkspaceDocMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddWorkspaceDocMembersHeaders{}
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.AddWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddWorkspaceMembersWithOptions(tmpReq *AddWorkspaceMembersRequest, tmpHeader *AddWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *AddWorkspaceMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddWorkspaceMembersRequest
//
// @return AddWorkspaceMembersResponse
func (client *Client) AddWorkspaceMembers(request *AddWorkspaceMembersRequest) (_result *AddWorkspaceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddWorkspaceMembersHeaders{}
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.AddWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssignTicketWithOptions(tmpReq *AssignTicketRequest, tmpHeader *AssignTicketHeaders, runtime *dara.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AssignTicketRequest
//
// @return AssignTicketResponse
func (client *Client) AssignTicket(request *AssignTicketRequest) (_result *AssignTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AssignTicketHeaders{}
	_result = &AssignTicketResponse{}
	_body, _err := client.AssignTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AuthorizeSkillWithOptions(tmpReq *AuthorizeSkillRequest, tmpHeader *AuthorizeSkillHeaders, runtime *dara.RuntimeOptions) (_result *AuthorizeSkillResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AuthorizeSkillRequest
//
// @return AuthorizeSkillResponse
func (client *Client) AuthorizeSkill(request *AuthorizeSkillRequest) (_result *AuthorizeSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuthorizeSkillHeaders{}
	_result = &AuthorizeSkillResponse{}
	_body, _err := client.AuthorizeSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchGetFormDataByIdListWithOptions(tmpReq *BatchGetFormDataByIdListRequest, tmpHeader *BatchGetFormDataByIdListHeaders, runtime *dara.RuntimeOptions) (_result *BatchGetFormDataByIdListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchGetFormDataByIdListRequest
//
// @return BatchGetFormDataByIdListResponse
func (client *Client) BatchGetFormDataByIdList(request *BatchGetFormDataByIdListRequest) (_result *BatchGetFormDataByIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchGetFormDataByIdListHeaders{}
	_result = &BatchGetFormDataByIdListResponse{}
	_body, _err := client.BatchGetFormDataByIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchRemovalByFormInstanceIdListWithOptions(tmpReq *BatchRemovalByFormInstanceIdListRequest, tmpHeader *BatchRemovalByFormInstanceIdListHeaders, runtime *dara.RuntimeOptions) (_result *BatchRemovalByFormInstanceIdListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchRemovalByFormInstanceIdListRequest
//
// @return BatchRemovalByFormInstanceIdListResponse
func (client *Client) BatchRemovalByFormInstanceIdList(request *BatchRemovalByFormInstanceIdListRequest) (_result *BatchRemovalByFormInstanceIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchRemovalByFormInstanceIdListHeaders{}
	_result = &BatchRemovalByFormInstanceIdListResponse{}
	_body, _err := client.BatchRemovalByFormInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchSaveFormDataWithOptions(tmpReq *BatchSaveFormDataRequest, tmpHeader *BatchSaveFormDataHeaders, runtime *dara.RuntimeOptions) (_result *BatchSaveFormDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchSaveFormDataRequest
//
// @return BatchSaveFormDataResponse
func (client *Client) BatchSaveFormData(request *BatchSaveFormDataRequest) (_result *BatchSaveFormDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchSaveFormDataHeaders{}
	_result = &BatchSaveFormDataResponse{}
	_body, _err := client.BatchSaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchUpdateFormDataByInstanceIdWithOptions(tmpReq *BatchUpdateFormDataByInstanceIdRequest, tmpHeader *BatchUpdateFormDataByInstanceIdHeaders, runtime *dara.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchUpdateFormDataByInstanceIdRequest
//
// @return BatchUpdateFormDataByInstanceIdResponse
func (client *Client) BatchUpdateFormDataByInstanceId(request *BatchUpdateFormDataByInstanceIdRequest) (_result *BatchUpdateFormDataByInstanceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchUpdateFormDataByInstanceIdHeaders{}
	_result = &BatchUpdateFormDataByInstanceIdResponse{}
	_body, _err := client.BatchUpdateFormDataByInstanceIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchUpdateFormDataByInstanceMapWithOptions(tmpReq *BatchUpdateFormDataByInstanceMapRequest, tmpHeader *BatchUpdateFormDataByInstanceMapHeaders, runtime *dara.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceMapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchUpdateFormDataByInstanceMapRequest
//
// @return BatchUpdateFormDataByInstanceMapResponse
func (client *Client) BatchUpdateFormDataByInstanceMap(request *BatchUpdateFormDataByInstanceMapRequest) (_result *BatchUpdateFormDataByInstanceMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &BatchUpdateFormDataByInstanceMapHeaders{}
	_result = &BatchUpdateFormDataByInstanceMapResponse{}
	_body, _err := client.BatchUpdateFormDataByInstanceMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelScheduleConferenceWithOptions(tmpReq *CancelScheduleConferenceRequest, tmpHeader *CancelScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CancelScheduleConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CancelScheduleConferenceRequest
//
// @return CancelScheduleConferenceResponse
func (client *Client) CancelScheduleConference(request *CancelScheduleConferenceRequest) (_result *CancelScheduleConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CancelScheduleConferenceHeaders{}
	_result = &CancelScheduleConferenceResponse{}
	_body, _err := client.CancelScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChangeDingTalkIdWithOptions(tmpReq *ChangeDingTalkIdRequest, tmpHeader *ChangeDingTalkIdHeaders, runtime *dara.RuntimeOptions) (_result *ChangeDingTalkIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ChangeDingTalkIdRequest
//
// @return ChangeDingTalkIdResponse
func (client *Client) ChangeDingTalkId(request *ChangeDingTalkIdRequest) (_result *ChangeDingTalkIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ChangeDingTalkIdHeaders{}
	_result = &ChangeDingTalkIdResponse{}
	_body, _err := client.ChangeDingTalkIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckAlibabaStaffWithOptions(tmpReq *CheckAlibabaStaffRequest, tmpHeader *CheckAlibabaStaffHeaders, runtime *dara.RuntimeOptions) (_result *CheckAlibabaStaffResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CheckAlibabaStaffRequest
//
// @return CheckAlibabaStaffResponse
func (client *Client) CheckAlibabaStaff(request *CheckAlibabaStaffRequest) (_result *CheckAlibabaStaffResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckAlibabaStaffHeaders{}
	_result = &CheckAlibabaStaffResponse{}
	_body, _err := client.CheckAlibabaStaffWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckUserIsGroupMemberWithOptions(request *CheckUserIsGroupMemberRequest, tmpHeader *CheckUserIsGroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *CheckUserIsGroupMemberResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckUserIsGroupMemberResponse
func (client *Client) CheckUserIsGroupMember(request *CheckUserIsGroupMemberRequest) (_result *CheckUserIsGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckUserIsGroupMemberHeaders{}
	_result = &CheckUserIsGroupMemberResponse{}
	_body, _err := client.CheckUserIsGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ClearWithOptions(tmpReq *ClearRequest, tmpHeader *ClearHeaders, runtime *dara.RuntimeOptions) (_result *ClearResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ClearRequest
//
// @return ClearResponse
func (client *Client) Clear(request *ClearRequest) (_result *ClearResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ClearHeaders{}
	_result = &ClearResponse{}
	_body, _err := client.ClearWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ClearDataWithOptions(tmpReq *ClearDataRequest, tmpHeader *ClearDataHeaders, runtime *dara.RuntimeOptions) (_result *ClearDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ClearDataRequest
//
// @return ClearDataResponse
func (client *Client) ClearData(request *ClearDataRequest) (_result *ClearDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ClearDataHeaders{}
	_result = &ClearDataResponse{}
	_body, _err := client.ClearDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CloseVideoConferenceWithOptions(tmpReq *CloseVideoConferenceRequest, tmpHeader *CloseVideoConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CloseVideoConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CloseVideoConferenceRequest
//
// @return CloseVideoConferenceResponse
func (client *Client) CloseVideoConference(request *CloseVideoConferenceRequest) (_result *CloseVideoConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CloseVideoConferenceHeaders{}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.CloseVideoConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CommentListReportWithOptions(tmpReq *CommentListReportRequest, tmpHeader *CommentListReportHeaders, runtime *dara.RuntimeOptions) (_result *CommentListReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CommentListReportRequest
//
// @return CommentListReportResponse
func (client *Client) CommentListReport(request *CommentListReportRequest) (_result *CommentListReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CommentListReportHeaders{}
	_result = &CommentListReportResponse{}
	_body, _err := client.CommentListReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CommitFileWithOptions(tmpReq *CommitFileRequest, tmpHeader *CommitFileHeaders, runtime *dara.RuntimeOptions) (_result *CommitFileResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CommitFileRequest
//
// @return CommitFileResponse
func (client *Client) CommitFile(request *CommitFileRequest) (_result *CommitFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CommitFileHeaders{}
	_result = &CommitFileResponse{}
	_body, _err := client.CommitFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CopyDentryWithOptions(tmpReq *CopyDentryRequest, tmpHeader *CopyDentryHeaders, runtime *dara.RuntimeOptions) (_result *CopyDentryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CopyDentryRequest
//
// @return CopyDentryResponse
func (client *Client) CopyDentry(request *CopyDentryRequest) (_result *CopyDentryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CopyDentryHeaders{}
	_result = &CopyDentryResponse{}
	_body, _err := client.CopyDentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - CreateAlidingAssistantRequest
//
// @param tmpHeader - CreateAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlidingAssistantResponse
func (client *Client) CreateAlidingAssistantWithOptions(tmpReq *CreateAlidingAssistantRequest, tmpHeader *CreateAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *CreateAlidingAssistantResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAlidingAssistantRequest
//
// @return CreateAlidingAssistantResponse
func (client *Client) CreateAlidingAssistant(request *CreateAlidingAssistantRequest) (_result *CreateAlidingAssistantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateAlidingAssistantHeaders{}
	_result = &CreateAlidingAssistantResponse{}
	_body, _err := client.CreateAlidingAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDeliveryPlanWithOptions(tmpReq *CreateDeliveryPlanRequest, tmpHeader *CreateDeliveryPlanHeaders, runtime *dara.RuntimeOptions) (_result *CreateDeliveryPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateDeliveryPlanRequest
//
// @return CreateDeliveryPlanResponse
func (client *Client) CreateDeliveryPlan(request *CreateDeliveryPlanRequest) (_result *CreateDeliveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDeliveryPlanHeaders{}
	_result = &CreateDeliveryPlanResponse{}
	_body, _err := client.CreateDeliveryPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDingtalkPersonalTodoTaskWithOptions(tmpReq *CreateDingtalkPersonalTodoTaskRequest, tmpHeader *CreateDingtalkPersonalTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateDingtalkPersonalTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateDingtalkPersonalTodoTaskRequest
//
// @return CreateDingtalkPersonalTodoTaskResponse
func (client *Client) CreateDingtalkPersonalTodoTask(request *CreateDingtalkPersonalTodoTaskRequest) (_result *CreateDingtalkPersonalTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateDingtalkPersonalTodoTaskHeaders{}
	_result = &CreateDingtalkPersonalTodoTaskResponse{}
	_body, _err := client.CreateDingtalkPersonalTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateEventWithOptions(tmpReq *CreateEventRequest, tmpHeader *CreateEventHeaders, runtime *dara.RuntimeOptions) (_result *CreateEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateEventRequest
//
// @return CreateEventResponse
func (client *Client) CreateEvent(request *CreateEventRequest) (_result *CreateEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateEventHeaders{}
	_result = &CreateEventResponse{}
	_body, _err := client.CreateEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLiveWithOptions(tmpReq *CreateLiveRequest, tmpHeader *CreateLiveHeaders, runtime *dara.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateLiveRequest
//
// @return CreateLiveResponse
func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateLiveHeaders{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateMeetingRoomWithOptions(tmpReq *CreateMeetingRoomRequest, tmpHeader *CreateMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *CreateMeetingRoomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateMeetingRoomRequest
//
// @return CreateMeetingRoomResponse
func (client *Client) CreateMeetingRoom(request *CreateMeetingRoomRequest) (_result *CreateMeetingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateMeetingRoomHeaders{}
	_result = &CreateMeetingRoomResponse{}
	_body, _err := client.CreateMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateMeetingRoomGroupWithOptions(tmpReq *CreateMeetingRoomGroupRequest, tmpHeader *CreateMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *CreateMeetingRoomGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateMeetingRoomGroupRequest
//
// @return CreateMeetingRoomGroupResponse
func (client *Client) CreateMeetingRoomGroup(request *CreateMeetingRoomGroupRequest) (_result *CreateMeetingRoomGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateMeetingRoomGroupHeaders{}
	_result = &CreateMeetingRoomGroupResponse{}
	_body, _err := client.CreateMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateMessageWithOptions(request *CreateMessageRequest, headers *CreateMessageHeaders, runtime *dara.RuntimeOptions) (_result *CreateMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateMessageResponse
func (client *Client) CreateMessage(request *CreateMessageRequest) (_result *CreateMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateMessageHeaders{}
	_result = &CreateMessageResponse{}
	_body, _err := client.CreateMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateMultiDimTableFieldWithOptions(tmpReq *CreateMultiDimTableFieldRequest, tmpHeader *CreateMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *CreateMultiDimTableFieldResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateMultiDimTableFieldRequest
//
// @return CreateMultiDimTableFieldResponse
func (client *Client) CreateMultiDimTableField(request *CreateMultiDimTableFieldRequest) (_result *CreateMultiDimTableFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateMultiDimTableFieldHeaders{}
	_result = &CreateMultiDimTableFieldResponse{}
	_body, _err := client.CreateMultiDimTableFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateOrUpdateFormDataWithOptions(request *CreateOrUpdateFormDataRequest, tmpHeader *CreateOrUpdateFormDataHeaders, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateFormDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormData(request *CreateOrUpdateFormDataRequest) (_result *CreateOrUpdateFormDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateOrUpdateFormDataHeaders{}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.CreateOrUpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateOrgHonorTemplateWithOptions(tmpReq *CreateOrgHonorTemplateRequest, tmpHeader *CreateOrgHonorTemplateHeaders, runtime *dara.RuntimeOptions) (_result *CreateOrgHonorTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateOrgHonorTemplateRequest
//
// @return CreateOrgHonorTemplateResponse
func (client *Client) CreateOrgHonorTemplate(request *CreateOrgHonorTemplateRequest) (_result *CreateOrgHonorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateOrgHonorTemplateHeaders{}
	_result = &CreateOrgHonorTemplateResponse{}
	_body, _err := client.CreateOrgHonorTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePersonalTodoTaskWithOptions(tmpReq *CreatePersonalTodoTaskRequest, tmpHeader *CreatePersonalTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreatePersonalTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreatePersonalTodoTaskRequest
//
// @return CreatePersonalTodoTaskResponse
func (client *Client) CreatePersonalTodoTask(request *CreatePersonalTodoTaskRequest) (_result *CreatePersonalTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreatePersonalTodoTaskHeaders{}
	_result = &CreatePersonalTodoTaskResponse{}
	_body, _err := client.CreatePersonalTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateReportWithOptions(tmpReq *CreateReportRequest, tmpHeader *CreateReportHeaders, runtime *dara.RuntimeOptions) (_result *CreateReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateReportRequest
//
// @return CreateReportResponse
func (client *Client) CreateReport(request *CreateReportRequest) (_result *CreateReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateReportHeaders{}
	_result = &CreateReportResponse{}
	_body, _err := client.CreateReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRunWithSSE(request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions, _yield chan *CreateRunResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createRunWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) CreateRunWithOptions(request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions) (_result *CreateRunResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRunResponse
func (client *Client) CreateRun(request *CreateRunRequest) (_result *CreateRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateRunHeaders{}
	_result = &CreateRunResponse{}
	_body, _err := client.CreateRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateScenegroupWithOptions(request *CreateScenegroupRequest, tmpHeader *CreateScenegroupHeaders, runtime *dara.RuntimeOptions) (_result *CreateScenegroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateScenegroupResponse
func (client *Client) CreateScenegroup(request *CreateScenegroupRequest) (_result *CreateScenegroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateScenegroupHeaders{}
	_result = &CreateScenegroupResponse{}
	_body, _err := client.CreateScenegroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateScheduleConferenceWithOptions(tmpReq *CreateScheduleConferenceRequest, tmpHeader *CreateScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CreateScheduleConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateScheduleConferenceRequest
//
// @return CreateScheduleConferenceResponse
func (client *Client) CreateScheduleConference(request *CreateScheduleConferenceRequest) (_result *CreateScheduleConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateScheduleConferenceHeaders{}
	_result = &CreateScheduleConferenceResponse{}
	_body, _err := client.CreateScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSearchDomeWithOptions(tmpReq *CreateSearchDomeRequest, tmpHeader *CreateSearchDomeHeaders, runtime *dara.RuntimeOptions) (_result *CreateSearchDomeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateSearchDomeRequest
//
// @return CreateSearchDomeResponse
func (client *Client) CreateSearchDome(request *CreateSearchDomeRequest) (_result *CreateSearchDomeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSearchDomeHeaders{}
	_result = &CreateSearchDomeResponse{}
	_body, _err := client.CreateSearchDomeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSearchKeywordWithOptions(tmpReq *CreateSearchKeywordRequest, tmpHeader *CreateSearchKeywordHeaders, runtime *dara.RuntimeOptions) (_result *CreateSearchKeywordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateSearchKeywordRequest
//
// @return CreateSearchKeywordResponse
func (client *Client) CreateSearchKeyword(request *CreateSearchKeywordRequest) (_result *CreateSearchKeywordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSearchKeywordHeaders{}
	_result = &CreateSearchKeywordResponse{}
	_body, _err := client.CreateSearchKeywordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSheetWithOptions(tmpReq *CreateSheetRequest, tmpHeader *CreateSheetHeaders, runtime *dara.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateSheetRequest
//
// @return CreateSheetResponse
func (client *Client) CreateSheet(request *CreateSheetRequest) (_result *CreateSheetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSheetHeaders{}
	_result = &CreateSheetResponse{}
	_body, _err := client.CreateSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSubscribedCalendarWithOptions(tmpReq *CreateSubscribedCalendarRequest, tmpHeader *CreateSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *CreateSubscribedCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateSubscribedCalendarRequest
//
// @return CreateSubscribedCalendarResponse
func (client *Client) CreateSubscribedCalendar(request *CreateSubscribedCalendarRequest) (_result *CreateSubscribedCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSubscribedCalendarHeaders{}
	_result = &CreateSubscribedCalendarResponse{}
	_body, _err := client.CreateSubscribedCalendarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateThreadWithOptions(request *CreateThreadRequest, headers *CreateThreadHeaders, runtime *dara.RuntimeOptions) (_result *CreateThreadResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateThreadResponse
func (client *Client) CreateThread(request *CreateThreadRequest) (_result *CreateThreadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateThreadHeaders{}
	_result = &CreateThreadResponse{}
	_body, _err := client.CreateThreadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTicketWithOptions(tmpReq *CreateTicketRequest, tmpHeader *CreateTicketHeaders, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateTicketHeaders{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTodoTaskWithOptions(tmpReq *CreateTodoTaskRequest, tmpHeader *CreateTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateTodoTaskRequest
//
// @return CreateTodoTaskResponse
func (client *Client) CreateTodoTask(request *CreateTodoTaskRequest) (_result *CreateTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateTodoTaskHeaders{}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CreateTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateVideoConferenceWithOptions(tmpReq *CreateVideoConferenceRequest, tmpHeader *CreateVideoConferenceHeaders, runtime *dara.RuntimeOptions) (_result *CreateVideoConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateVideoConferenceRequest
//
// @return CreateVideoConferenceResponse
func (client *Client) CreateVideoConference(request *CreateVideoConferenceRequest) (_result *CreateVideoConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateVideoConferenceHeaders{}
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.CreateVideoConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateWorkspaceWithOptions(tmpReq *CreateWorkspaceRequest, tmpHeader *CreateWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateWorkspaceHeaders{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateWorkspaceDocWithOptions(tmpReq *CreateWorkspaceDocRequest, tmpHeader *CreateWorkspaceDocHeaders, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceDocResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateWorkspaceDocRequest
//
// @return CreateWorkspaceDocResponse
func (client *Client) CreateWorkspaceDoc(request *CreateWorkspaceDocRequest) (_result *CreateWorkspaceDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateWorkspaceDocHeaders{}
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.CreateWorkspaceDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - DeleteAlidingAssistantRequest
//
// @param tmpHeader - DeleteAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlidingAssistantResponse
func (client *Client) DeleteAlidingAssistantWithOptions(tmpReq *DeleteAlidingAssistantRequest, tmpHeader *DeleteAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *DeleteAlidingAssistantResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAlidingAssistantRequest
//
// @return DeleteAlidingAssistantResponse
func (client *Client) DeleteAlidingAssistant(request *DeleteAlidingAssistantRequest) (_result *DeleteAlidingAssistantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteAlidingAssistantHeaders{}
	_result = &DeleteAlidingAssistantResponse{}
	_body, _err := client.DeleteAlidingAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteColumnsWithOptions(tmpReq *DeleteColumnsRequest, tmpHeader *DeleteColumnsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteColumnsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteColumnsRequest
//
// @return DeleteColumnsResponse
func (client *Client) DeleteColumns(request *DeleteColumnsRequest) (_result *DeleteColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteColumnsHeaders{}
	_result = &DeleteColumnsResponse{}
	_body, _err := client.DeleteColumnsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDentryWithOptions(tmpReq *DeleteDentryRequest, tmpHeader *DeleteDentryHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDentryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteDentryRequest
//
// @return DeleteDentryResponse
func (client *Client) DeleteDentry(request *DeleteDentryRequest) (_result *DeleteDentryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteDentryHeaders{}
	_result = &DeleteDentryResponse{}
	_body, _err := client.DeleteDentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDriveSpaceWithOptions(tmpReq *DeleteDriveSpaceRequest, tmpHeader *DeleteDriveSpaceHeaders, runtime *dara.RuntimeOptions) (_result *DeleteDriveSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteDriveSpaceRequest
//
// @return DeleteDriveSpaceResponse
func (client *Client) DeleteDriveSpace(request *DeleteDriveSpaceRequest) (_result *DeleteDriveSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteDriveSpaceHeaders{}
	_result = &DeleteDriveSpaceResponse{}
	_body, _err := client.DeleteDriveSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteEventWithOptions(request *DeleteEventRequest, tmpHeader *DeleteEventHeaders, runtime *dara.RuntimeOptions) (_result *DeleteEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteEventResponse
func (client *Client) DeleteEvent(request *DeleteEventRequest) (_result *DeleteEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteEventHeaders{}
	_result = &DeleteEventResponse{}
	_body, _err := client.DeleteEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFormDataWithOptions(request *DeleteFormDataRequest, tmpHeader *DeleteFormDataHeaders, runtime *dara.RuntimeOptions) (_result *DeleteFormDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFormDataResponse
func (client *Client) DeleteFormData(request *DeleteFormDataRequest) (_result *DeleteFormDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteFormDataHeaders{}
	_result = &DeleteFormDataResponse{}
	_body, _err := client.DeleteFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param tmpHeader - DeleteInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, tmpHeader *DeleteInstanceHeaders, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteInstanceHeaders{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteLiveWithOptions(tmpReq *DeleteLiveRequest, tmpHeader *DeleteLiveHeaders, runtime *dara.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteLiveRequest
//
// @return DeleteLiveResponse
func (client *Client) DeleteLive(request *DeleteLiveRequest) (_result *DeleteLiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteLiveHeaders{}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DeleteLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMeetingRoomWithOptions(tmpReq *DeleteMeetingRoomRequest, tmpHeader *DeleteMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMeetingRoomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteMeetingRoomRequest
//
// @return DeleteMeetingRoomResponse
func (client *Client) DeleteMeetingRoom(request *DeleteMeetingRoomRequest) (_result *DeleteMeetingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMeetingRoomHeaders{}
	_result = &DeleteMeetingRoomResponse{}
	_body, _err := client.DeleteMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMeetingRoomGroupWithOptions(tmpReq *DeleteMeetingRoomGroupRequest, tmpHeader *DeleteMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMeetingRoomGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteMeetingRoomGroupRequest
//
// @return DeleteMeetingRoomGroupResponse
func (client *Client) DeleteMeetingRoomGroup(request *DeleteMeetingRoomGroupRequest) (_result *DeleteMeetingRoomGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMeetingRoomGroupHeaders{}
	_result = &DeleteMeetingRoomGroupResponse{}
	_body, _err := client.DeleteMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMultiDimTableFieldWithOptions(tmpReq *DeleteMultiDimTableFieldRequest, tmpHeader *DeleteMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMultiDimTableFieldResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteMultiDimTableFieldRequest
//
// @return DeleteMultiDimTableFieldResponse
func (client *Client) DeleteMultiDimTableField(request *DeleteMultiDimTableFieldRequest) (_result *DeleteMultiDimTableFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMultiDimTableFieldHeaders{}
	_result = &DeleteMultiDimTableFieldResponse{}
	_body, _err := client.DeleteMultiDimTableFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMultiDimTableRecordsWithOptions(tmpReq *DeleteMultiDimTableRecordsRequest, tmpHeader *DeleteMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteMultiDimTableRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteMultiDimTableRecordsRequest
//
// @return DeleteMultiDimTableRecordsResponse
func (client *Client) DeleteMultiDimTableRecords(request *DeleteMultiDimTableRecordsRequest) (_result *DeleteMultiDimTableRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteMultiDimTableRecordsHeaders{}
	_result = &DeleteMultiDimTableRecordsResponse{}
	_body, _err := client.DeleteMultiDimTableRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePermissionWithOptions(tmpReq *DeletePermissionRequest, tmpHeader *DeletePermissionHeaders, runtime *dara.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeletePermissionRequest
//
// @return DeletePermissionResponse
func (client *Client) DeletePermission(request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRowsWithOptions(tmpReq *DeleteRowsRequest, tmpHeader *DeleteRowsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteRowsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteRowsRequest
//
// @return DeleteRowsResponse
func (client *Client) DeleteRows(request *DeleteRowsRequest) (_result *DeleteRowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteRowsHeaders{}
	_result = &DeleteRowsResponse{}
	_body, _err := client.DeleteRowsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteScenegroupMemberWithOptions(request *DeleteScenegroupMemberRequest, tmpHeader *DeleteScenegroupMemberHeaders, runtime *dara.RuntimeOptions) (_result *DeleteScenegroupMemberResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteScenegroupMemberResponse
func (client *Client) DeleteScenegroupMember(request *DeleteScenegroupMemberRequest) (_result *DeleteScenegroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteScenegroupMemberHeaders{}
	_result = &DeleteScenegroupMemberResponse{}
	_body, _err := client.DeleteScenegroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSheetWithOptions(tmpReq *DeleteSheetRequest, tmpHeader *DeleteSheetHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSheetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteSheetRequest
//
// @return DeleteSheetResponse
func (client *Client) DeleteSheet(request *DeleteSheetRequest) (_result *DeleteSheetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSheetHeaders{}
	_result = &DeleteSheetResponse{}
	_body, _err := client.DeleteSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteSubscribedCalendarWithOptions(request *DeleteSubscribedCalendarRequest, tmpHeader *DeleteSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSubscribedCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteSubscribedCalendarResponse
func (client *Client) DeleteSubscribedCalendar(request *DeleteSubscribedCalendarRequest) (_result *DeleteSubscribedCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSubscribedCalendarHeaders{}
	_result = &DeleteSubscribedCalendarResponse{}
	_body, _err := client.DeleteSubscribedCalendarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTodoTaskWithOptions(tmpReq *DeleteTodoTaskRequest, tmpHeader *DeleteTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteTodoTaskRequest
//
// @return DeleteTodoTaskResponse
func (client *Client) DeleteTodoTask(request *DeleteTodoTaskRequest) (_result *DeleteTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteTodoTaskHeaders{}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DeleteTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteWorkspaceDocMembersWithOptions(tmpReq *DeleteWorkspaceDocMembersRequest, tmpHeader *DeleteWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteWorkspaceDocMembersRequest
//
// @return DeleteWorkspaceDocMembersResponse
func (client *Client) DeleteWorkspaceDocMembers(request *DeleteWorkspaceDocMembersRequest) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteWorkspaceDocMembersHeaders{}
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.DeleteWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteWorkspaceMembersWithOptions(tmpReq *DeleteWorkspaceMembersRequest, tmpHeader *DeleteWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteWorkspaceMembersRequest
//
// @return DeleteWorkspaceMembersResponse
func (client *Client) DeleteWorkspaceMembers(request *DeleteWorkspaceMembersRequest) (_result *DeleteWorkspaceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteWorkspaceMembersHeaders{}
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.DeleteWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocBlocksDeleteWithOptions(tmpReq *DocBlocksDeleteRequest, tmpHeader *DocBlocksDeleteHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksDeleteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DocBlocksDeleteRequest
//
// @return DocBlocksDeleteResponse
func (client *Client) DocBlocksDelete(request *DocBlocksDeleteRequest) (_result *DocBlocksDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DocBlocksDeleteHeaders{}
	_result = &DocBlocksDeleteResponse{}
	_body, _err := client.DocBlocksDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocBlocksModifyWithOptions(tmpReq *DocBlocksModifyRequest, tmpHeader *DocBlocksModifyHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksModifyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DocBlocksModifyRequest
//
// @return DocBlocksModifyResponse
func (client *Client) DocBlocksModify(request *DocBlocksModifyRequest) (_result *DocBlocksModifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DocBlocksModifyHeaders{}
	_result = &DocBlocksModifyResponse{}
	_body, _err := client.DocBlocksModifyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocBlocksQueryWithOptions(tmpReq *DocBlocksQueryRequest, tmpHeader *DocBlocksQueryHeaders, runtime *dara.RuntimeOptions) (_result *DocBlocksQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DocBlocksQueryRequest
//
// @return DocBlocksQueryResponse
func (client *Client) DocBlocksQuery(request *DocBlocksQueryRequest) (_result *DocBlocksQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DocBlocksQueryHeaders{}
	_result = &DocBlocksQueryResponse{}
	_body, _err := client.DocBlocksQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocUpdateContentWithOptions(tmpReq *DocUpdateContentRequest, tmpHeader *DocUpdateContentHeaders, runtime *dara.RuntimeOptions) (_result *DocUpdateContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DocUpdateContentRequest
//
// @return DocUpdateContentResponse
func (client *Client) DocUpdateContent(request *DocUpdateContentRequest) (_result *DocUpdateContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DocUpdateContentHeaders{}
	_result = &DocUpdateContentResponse{}
	_body, _err := client.DocUpdateContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteBatchTaskWithOptions(request *ExecuteBatchTaskRequest, tmpHeader *ExecuteBatchTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteBatchTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteBatchTaskResponse
func (client *Client) ExecuteBatchTask(request *ExecuteBatchTaskRequest) (_result *ExecuteBatchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExecuteBatchTaskHeaders{}
	_result = &ExecuteBatchTaskResponse{}
	_body, _err := client.ExecuteBatchTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecutePlatformTaskWithOptions(request *ExecutePlatformTaskRequest, tmpHeader *ExecutePlatformTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecutePlatformTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecutePlatformTaskResponse
func (client *Client) ExecutePlatformTask(request *ExecutePlatformTaskRequest) (_result *ExecutePlatformTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExecutePlatformTaskHeaders{}
	_result = &ExecutePlatformTaskResponse{}
	_body, _err := client.ExecutePlatformTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteTaskWithOptions(request *ExecuteTaskRequest, tmpHeader *ExecuteTaskHeaders, runtime *dara.RuntimeOptions) (_result *ExecuteTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteTaskResponse
func (client *Client) ExecuteTask(request *ExecuteTaskRequest) (_result *ExecuteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExecuteTaskHeaders{}
	_result = &ExecuteTaskResponse{}
	_body, _err := client.ExecuteTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExpandGroupCapacityWithOptions(tmpReq *ExpandGroupCapacityRequest, tmpHeader *ExpandGroupCapacityHeaders, runtime *dara.RuntimeOptions) (_result *ExpandGroupCapacityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ExpandGroupCapacityRequest
//
// @return ExpandGroupCapacityResponse
func (client *Client) ExpandGroupCapacity(request *ExpandGroupCapacityRequest) (_result *ExpandGroupCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ExpandGroupCapacityHeaders{}
	_result = &ExpandGroupCapacityResponse{}
	_body, _err := client.ExpandGroupCapacityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FinishTicketWithOptions(tmpReq *FinishTicketRequest, tmpHeader *FinishTicketHeaders, runtime *dara.RuntimeOptions) (_result *FinishTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FinishTicketRequest
//
// @return FinishTicketResponse
func (client *Client) FinishTicket(request *FinishTicketRequest) (_result *FinishTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FinishTicketHeaders{}
	_result = &FinishTicketResponse{}
	_body, _err := client.FinishTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetActivityListWithOptions(request *GetActivityListRequest, tmpHeader *GetActivityListHeaders, runtime *dara.RuntimeOptions) (_result *GetActivityListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetActivityListResponse
func (client *Client) GetActivityList(request *GetActivityListRequest) (_result *GetActivityListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetActivityListHeaders{}
	_result = &GetActivityListResponse{}
	_body, _err := client.GetActivityListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAlidingAssistantInfoWithOptions(tmpReq *GetAlidingAssistantInfoRequest, tmpHeader *GetAlidingAssistantInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetAlidingAssistantInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetAlidingAssistantInfoRequest
//
// @return GetAlidingAssistantInfoResponse
func (client *Client) GetAlidingAssistantInfo(request *GetAlidingAssistantInfoRequest) (_result *GetAlidingAssistantInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAlidingAssistantInfoHeaders{}
	_result = &GetAlidingAssistantInfoResponse{}
	_body, _err := client.GetAlidingAssistantInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAllSheetsWithOptions(tmpReq *GetAllSheetsRequest, tmpHeader *GetAllSheetsHeaders, runtime *dara.RuntimeOptions) (_result *GetAllSheetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetAllSheetsRequest
//
// @return GetAllSheetsResponse
func (client *Client) GetAllSheets(request *GetAllSheetsRequest) (_result *GetAllSheetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAllSheetsHeaders{}
	_result = &GetAllSheetsResponse{}
	_body, _err := client.GetAllSheetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAssistantCapabilityWithOptions(request *GetAssistantCapabilityRequest, headers *GetAssistantCapabilityHeaders, runtime *dara.RuntimeOptions) (_result *GetAssistantCapabilityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAssistantCapabilityResponse
func (client *Client) GetAssistantCapability(request *GetAssistantCapabilityRequest) (_result *GetAssistantCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAssistantCapabilityHeaders{}
	_result = &GetAssistantCapabilityResponse{}
	_body, _err := client.GetAssistantCapabilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - GetCardTemplateRequest
//
// @param tmpHeader - GetCardTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardTemplateResponse
func (client *Client) GetCardTemplateWithOptions(tmpReq *GetCardTemplateRequest, tmpHeader *GetCardTemplateHeaders, runtime *dara.RuntimeOptions) (_result *GetCardTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCardTemplateRequest
//
// @return GetCardTemplateResponse
func (client *Client) GetCardTemplate(request *GetCardTemplateRequest) (_result *GetCardTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCardTemplateHeaders{}
	_result = &GetCardTemplateResponse{}
	_body, _err := client.GetCardTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetConversaionSpaceWithOptions(tmpReq *GetConversaionSpaceRequest, tmpHeader *GetConversaionSpaceHeaders, runtime *dara.RuntimeOptions) (_result *GetConversaionSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetConversaionSpaceRequest
//
// @return GetConversaionSpaceResponse
func (client *Client) GetConversaionSpace(request *GetConversaionSpaceRequest) (_result *GetConversaionSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetConversaionSpaceHeaders{}
	_result = &GetConversaionSpaceResponse{}
	_body, _err := client.GetConversaionSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCorpAccomplishmentTasksWithOptions(request *GetCorpAccomplishmentTasksRequest, tmpHeader *GetCorpAccomplishmentTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetCorpAccomplishmentTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCorpAccomplishmentTasksResponse
func (client *Client) GetCorpAccomplishmentTasks(request *GetCorpAccomplishmentTasksRequest) (_result *GetCorpAccomplishmentTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCorpAccomplishmentTasksHeaders{}
	_result = &GetCorpAccomplishmentTasksResponse{}
	_body, _err := client.GetCorpAccomplishmentTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCorpTasksWithOptions(request *GetCorpTasksRequest, tmpHeader *GetCorpTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetCorpTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCorpTasksResponse
func (client *Client) GetCorpTasks(request *GetCorpTasksRequest) (_result *GetCorpTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCorpTasksHeaders{}
	_result = &GetCorpTasksResponse{}
	_body, _err := client.GetCorpTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - GetDeptNoRequest
//
// @param tmpHeader - GetDeptNoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeptNoResponse
func (client *Client) GetDeptNoWithOptions(tmpReq *GetDeptNoRequest, tmpHeader *GetDeptNoHeaders, runtime *dara.RuntimeOptions) (_result *GetDeptNoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDeptNoRequest
//
// @return GetDeptNoResponse
func (client *Client) GetDeptNo(request *GetDeptNoRequest) (_result *GetDeptNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeptNoHeaders{}
	_result = &GetDeptNoResponse{}
	_body, _err := client.GetDeptNoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocContentWithOptions(tmpReq *GetDocContentRequest, tmpHeader *GetDocContentHeaders, runtime *dara.RuntimeOptions) (_result *GetDocContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetDocContentRequest
//
// @return GetDocContentResponse
func (client *Client) GetDocContent(request *GetDocContentRequest) (_result *GetDocContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDocContentHeaders{}
	_result = &GetDocContentResponse{}
	_body, _err := client.GetDocContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocContentTakIdWithOptions(tmpReq *GetDocContentTakIdRequest, tmpHeader *GetDocContentTakIdHeaders, runtime *dara.RuntimeOptions) (_result *GetDocContentTakIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetDocContentTakIdRequest
//
// @return GetDocContentTakIdResponse
func (client *Client) GetDocContentTakId(request *GetDocContentTakIdRequest) (_result *GetDocContentTakIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDocContentTakIdHeaders{}
	_result = &GetDocContentTakIdResponse{}
	_body, _err := client.GetDocContentTakIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEventWithOptions(request *GetEventRequest, tmpHeader *GetEventHeaders, runtime *dara.RuntimeOptions) (_result *GetEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEventResponse
func (client *Client) GetEvent(request *GetEventRequest) (_result *GetEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetEventHeaders{}
	_result = &GetEventResponse{}
	_body, _err := client.GetEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFieldDefByUuidWithOptions(request *GetFieldDefByUuidRequest, tmpHeader *GetFieldDefByUuidHeaders, runtime *dara.RuntimeOptions) (_result *GetFieldDefByUuidResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFieldDefByUuidResponse
func (client *Client) GetFieldDefByUuid(request *GetFieldDefByUuidRequest) (_result *GetFieldDefByUuidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFieldDefByUuidHeaders{}
	_result = &GetFieldDefByUuidResponse{}
	_body, _err := client.GetFieldDefByUuidWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFileDownloadInfoWithOptions(tmpReq *GetFileDownloadInfoRequest, tmpHeader *GetFileDownloadInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetFileDownloadInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetFileDownloadInfoRequest
//
// @return GetFileDownloadInfoResponse
func (client *Client) GetFileDownloadInfo(request *GetFileDownloadInfoRequest) (_result *GetFileDownloadInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFileDownloadInfoHeaders{}
	_result = &GetFileDownloadInfoResponse{}
	_body, _err := client.GetFileDownloadInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFileUploadInfoWithOptions(tmpReq *GetFileUploadInfoRequest, tmpHeader *GetFileUploadInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetFileUploadInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetFileUploadInfoRequest
//
// @return GetFileUploadInfoResponse
func (client *Client) GetFileUploadInfo(request *GetFileUploadInfoRequest) (_result *GetFileUploadInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFileUploadInfoHeaders{}
	_result = &GetFileUploadInfoResponse{}
	_body, _err := client.GetFileUploadInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFormComponentDefinitionListWithOptions(request *GetFormComponentDefinitionListRequest, tmpHeader *GetFormComponentDefinitionListHeaders, runtime *dara.RuntimeOptions) (_result *GetFormComponentDefinitionListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFormComponentDefinitionListResponse
func (client *Client) GetFormComponentDefinitionList(request *GetFormComponentDefinitionListRequest) (_result *GetFormComponentDefinitionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFormComponentDefinitionListHeaders{}
	_result = &GetFormComponentDefinitionListResponse{}
	_body, _err := client.GetFormComponentDefinitionListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFormDataByIDWithOptions(request *GetFormDataByIDRequest, tmpHeader *GetFormDataByIDHeaders, runtime *dara.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFormDataByIDResponse
func (client *Client) GetFormDataByID(request *GetFormDataByIDRequest) (_result *GetFormDataByIDResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFormDataByIDHeaders{}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.GetFormDataByIDWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFormListInAppWithOptions(request *GetFormListInAppRequest, tmpHeader *GetFormListInAppHeaders, runtime *dara.RuntimeOptions) (_result *GetFormListInAppResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFormListInAppResponse
func (client *Client) GetFormListInApp(request *GetFormListInAppRequest) (_result *GetFormListInAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetFormListInAppHeaders{}
	_result = &GetFormListInAppResponse{}
	_body, _err := client.GetFormListInAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGroupLiveListWithOptions(tmpReq *GetGroupLiveListRequest, tmpHeader *GetGroupLiveListHeaders, runtime *dara.RuntimeOptions) (_result *GetGroupLiveListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetGroupLiveListRequest
//
// @return GetGroupLiveListResponse
func (client *Client) GetGroupLiveList(request *GetGroupLiveListRequest) (_result *GetGroupLiveListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetGroupLiveListHeaders{}
	_result = &GetGroupLiveListResponse{}
	_body, _err := client.GetGroupLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInnerGroupMembersWithOptions(request *GetInnerGroupMembersRequest, tmpHeader *GetInnerGroupMembersHeaders, runtime *dara.RuntimeOptions) (_result *GetInnerGroupMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInnerGroupMembersResponse
func (client *Client) GetInnerGroupMembers(request *GetInnerGroupMembersRequest) (_result *GetInnerGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetInnerGroupMembersHeaders{}
	_result = &GetInnerGroupMembersResponse{}
	_body, _err := client.GetInnerGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceByIdWithOptions(request *GetInstanceByIdRequest, tmpHeader *GetInstanceByIdHeaders, runtime *dara.RuntimeOptions) (_result *GetInstanceByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceById(request *GetInstanceByIdRequest) (_result *GetInstanceByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetInstanceByIdHeaders{}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.GetInstanceByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstanceIdListWithOptions(request *GetInstanceIdListRequest, tmpHeader *GetInstanceIdListHeaders, runtime *dara.RuntimeOptions) (_result *GetInstanceIdListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdList(request *GetInstanceIdListRequest) (_result *GetInstanceIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetInstanceIdListHeaders{}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.GetInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstancesWithOptions(request *GetInstancesRequest, tmpHeader *GetInstancesHeaders, runtime *dara.RuntimeOptions) (_result *GetInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstancesResponse
func (client *Client) GetInstances(request *GetInstancesRequest) (_result *GetInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetInstancesHeaders{}
	_result = &GetInstancesResponse{}
	_body, _err := client.GetInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInstancesByIdListWithOptions(request *GetInstancesByIdListRequest, tmpHeader *GetInstancesByIdListHeaders, runtime *dara.RuntimeOptions) (_result *GetInstancesByIdListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInstancesByIdListResponse
func (client *Client) GetInstancesByIdList(request *GetInstancesByIdListRequest) (_result *GetInstancesByIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetInstancesByIdListHeaders{}
	_result = &GetInstancesByIdListResponse{}
	_body, _err := client.GetInstancesByIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetLiveReplayUrlWithOptions(tmpReq *GetLiveReplayUrlRequest, tmpHeader *GetLiveReplayUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetLiveReplayUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetLiveReplayUrlRequest
//
// @return GetLiveReplayUrlResponse
func (client *Client) GetLiveReplayUrl(request *GetLiveReplayUrlRequest) (_result *GetLiveReplayUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetLiveReplayUrlHeaders{}
	_result = &GetLiveReplayUrlResponse{}
	_body, _err := client.GetLiveReplayUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMeCorpSubmissionWithOptions(request *GetMeCorpSubmissionRequest, tmpHeader *GetMeCorpSubmissionHeaders, runtime *dara.RuntimeOptions) (_result *GetMeCorpSubmissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMeCorpSubmissionResponse
func (client *Client) GetMeCorpSubmission(request *GetMeCorpSubmissionRequest) (_result *GetMeCorpSubmissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMeCorpSubmissionHeaders{}
	_result = &GetMeCorpSubmissionResponse{}
	_body, _err := client.GetMeCorpSubmissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMeetingRoomsScheduleWithOptions(tmpReq *GetMeetingRoomsScheduleRequest, tmpHeader *GetMeetingRoomsScheduleHeaders, runtime *dara.RuntimeOptions) (_result *GetMeetingRoomsScheduleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMeetingRoomsScheduleRequest
//
// @return GetMeetingRoomsScheduleResponse
func (client *Client) GetMeetingRoomsSchedule(request *GetMeetingRoomsScheduleRequest) (_result *GetMeetingRoomsScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMeetingRoomsScheduleHeaders{}
	_result = &GetMeetingRoomsScheduleResponse{}
	_body, _err := client.GetMeetingRoomsScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMineWorkspaceWithOptions(tmpReq *GetMineWorkspaceRequest, tmpHeader *GetMineWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *GetMineWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMineWorkspaceRequest
//
// @return GetMineWorkspaceResponse
func (client *Client) GetMineWorkspace(request *GetMineWorkspaceRequest) (_result *GetMineWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMineWorkspaceHeaders{}
	_result = &GetMineWorkspaceResponse{}
	_body, _err := client.GetMineWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiDimTableAllFieldsWithOptions(tmpReq *GetMultiDimTableAllFieldsRequest, tmpHeader *GetMultiDimTableAllFieldsHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableAllFieldsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMultiDimTableAllFieldsRequest
//
// @return GetMultiDimTableAllFieldsResponse
func (client *Client) GetMultiDimTableAllFields(request *GetMultiDimTableAllFieldsRequest) (_result *GetMultiDimTableAllFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMultiDimTableAllFieldsHeaders{}
	_result = &GetMultiDimTableAllFieldsResponse{}
	_body, _err := client.GetMultiDimTableAllFieldsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiDimTableAllSheetsWithOptions(tmpReq *GetMultiDimTableAllSheetsRequest, tmpHeader *GetMultiDimTableAllSheetsHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableAllSheetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMultiDimTableAllSheetsRequest
//
// @return GetMultiDimTableAllSheetsResponse
func (client *Client) GetMultiDimTableAllSheets(request *GetMultiDimTableAllSheetsRequest) (_result *GetMultiDimTableAllSheetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMultiDimTableAllSheetsHeaders{}
	_result = &GetMultiDimTableAllSheetsResponse{}
	_body, _err := client.GetMultiDimTableAllSheetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiDimTableRecordWithOptions(tmpReq *GetMultiDimTableRecordRequest, tmpHeader *GetMultiDimTableRecordHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMultiDimTableRecordRequest
//
// @return GetMultiDimTableRecordResponse
func (client *Client) GetMultiDimTableRecord(request *GetMultiDimTableRecordRequest) (_result *GetMultiDimTableRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMultiDimTableRecordHeaders{}
	_result = &GetMultiDimTableRecordResponse{}
	_body, _err := client.GetMultiDimTableRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultiDimTableSheetWithOptions(tmpReq *GetMultiDimTableSheetRequest, tmpHeader *GetMultiDimTableSheetHeaders, runtime *dara.RuntimeOptions) (_result *GetMultiDimTableSheetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMultiDimTableSheetRequest
//
// @return GetMultiDimTableSheetResponse
func (client *Client) GetMultiDimTableSheet(request *GetMultiDimTableSheetRequest) (_result *GetMultiDimTableSheetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMultiDimTableSheetHeaders{}
	_result = &GetMultiDimTableSheetResponse{}
	_body, _err := client.GetMultiDimTableSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMultipartFileUploadInfosWithOptions(tmpReq *GetMultipartFileUploadInfosRequest, tmpHeader *GetMultipartFileUploadInfosHeaders, runtime *dara.RuntimeOptions) (_result *GetMultipartFileUploadInfosResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetMultipartFileUploadInfosRequest
//
// @return GetMultipartFileUploadInfosResponse
func (client *Client) GetMultipartFileUploadInfos(request *GetMultipartFileUploadInfosRequest) (_result *GetMultipartFileUploadInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetMultipartFileUploadInfosHeaders{}
	_result = &GetMultipartFileUploadInfosResponse{}
	_body, _err := client.GetMultipartFileUploadInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNewestInnerGroupsWithOptions(tmpReq *GetNewestInnerGroupsRequest, tmpHeader *GetNewestInnerGroupsHeaders, runtime *dara.RuntimeOptions) (_result *GetNewestInnerGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetNewestInnerGroupsRequest
//
// @return GetNewestInnerGroupsResponse
func (client *Client) GetNewestInnerGroups(request *GetNewestInnerGroupsRequest) (_result *GetNewestInnerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetNewestInnerGroupsHeaders{}
	_result = &GetNewestInnerGroupsResponse{}
	_body, _err := client.GetNewestInnerGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNodeWithOptions(tmpReq *GetNodeRequest, tmpHeader *GetNodeHeaders, runtime *dara.RuntimeOptions) (_result *GetNodeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetNodeHeaders{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNodeByUrlWithOptions(tmpReq *GetNodeByUrlRequest, tmpHeader *GetNodeByUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetNodeByUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetNodeByUrlRequest
//
// @return GetNodeByUrlResponse
func (client *Client) GetNodeByUrl(request *GetNodeByUrlRequest) (_result *GetNodeByUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetNodeByUrlHeaders{}
	_result = &GetNodeByUrlResponse{}
	_body, _err := client.GetNodeByUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNodesWithOptions(tmpReq *GetNodesRequest, tmpHeader *GetNodesHeaders, runtime *dara.RuntimeOptions) (_result *GetNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetNodesRequest
//
// @return GetNodesResponse
func (client *Client) GetNodes(request *GetNodesRequest) (_result *GetNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetNodesHeaders{}
	_result = &GetNodesResponse{}
	_body, _err := client.GetNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNotifyMeWithOptions(request *GetNotifyMeRequest, tmpHeader *GetNotifyMeHeaders, runtime *dara.RuntimeOptions) (_result *GetNotifyMeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetNotifyMeResponse
func (client *Client) GetNotifyMe(request *GetNotifyMeRequest) (_result *GetNotifyMeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetNotifyMeHeaders{}
	_result = &GetNotifyMeResponse{}
	_body, _err := client.GetNotifyMeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOpenUrlWithOptions(request *GetOpenUrlRequest, tmpHeader *GetOpenUrlHeaders, runtime *dara.RuntimeOptions) (_result *GetOpenUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOpenUrlResponse
func (client *Client) GetOpenUrl(request *GetOpenUrlRequest) (_result *GetOpenUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOpenUrlHeaders{}
	_result = &GetOpenUrlResponse{}
	_body, _err := client.GetOpenUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOperationRecordsWithOptions(request *GetOperationRecordsRequest, tmpHeader *GetOperationRecordsHeaders, runtime *dara.RuntimeOptions) (_result *GetOperationRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOperationRecordsResponse
func (client *Client) GetOperationRecords(request *GetOperationRecordsRequest) (_result *GetOperationRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOperationRecordsHeaders{}
	_result = &GetOperationRecordsResponse{}
	_body, _err := client.GetOperationRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - GetOrgLiveListRequest
//
// @param tmpHeader - GetOrgLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgLiveListResponse
func (client *Client) GetOrgLiveListWithOptions(tmpReq *GetOrgLiveListRequest, tmpHeader *GetOrgLiveListHeaders, runtime *dara.RuntimeOptions) (_result *GetOrgLiveListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOrgLiveListRequest
//
// @return GetOrgLiveListResponse
func (client *Client) GetOrgLiveList(request *GetOrgLiveListRequest) (_result *GetOrgLiveListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOrgLiveListHeaders{}
	_result = &GetOrgLiveListResponse{}
	_body, _err := client.GetOrgLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOrgOrWebOpenDocContentTaskIdWithOptions(tmpReq *GetOrgOrWebOpenDocContentTaskIdRequest, tmpHeader *GetOrgOrWebOpenDocContentTaskIdHeaders, runtime *dara.RuntimeOptions) (_result *GetOrgOrWebOpenDocContentTaskIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetOrgOrWebOpenDocContentTaskIdRequest
//
// @return GetOrgOrWebOpenDocContentTaskIdResponse
func (client *Client) GetOrgOrWebOpenDocContentTaskId(request *GetOrgOrWebOpenDocContentTaskIdRequest) (_result *GetOrgOrWebOpenDocContentTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOrgOrWebOpenDocContentTaskIdHeaders{}
	_result = &GetOrgOrWebOpenDocContentTaskIdResponse{}
	_body, _err := client.GetOrgOrWebOpenDocContentTaskIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProcessDefinitionWithOptions(request *GetProcessDefinitionRequest, tmpHeader *GetProcessDefinitionHeaders, runtime *dara.RuntimeOptions) (_result *GetProcessDefinitionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProcessDefinitionResponse
func (client *Client) GetProcessDefinition(request *GetProcessDefinitionRequest) (_result *GetProcessDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetProcessDefinitionHeaders{}
	_result = &GetProcessDefinitionResponse{}
	_body, _err := client.GetProcessDefinitionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRangeWithOptions(tmpReq *GetRangeRequest, tmpHeader *GetRangeHeaders, runtime *dara.RuntimeOptions) (_result *GetRangeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetRangeRequest
//
// @return GetRangeResponse
func (client *Client) GetRange(request *GetRangeRequest) (_result *GetRangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetRangeHeaders{}
	_result = &GetRangeResponse{}
	_body, _err := client.GetRangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRelatedWorkspacesWithOptions(tmpReq *GetRelatedWorkspacesRequest, tmpHeader *GetRelatedWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *GetRelatedWorkspacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetRelatedWorkspacesRequest
//
// @return GetRelatedWorkspacesResponse
func (client *Client) GetRelatedWorkspaces(request *GetRelatedWorkspacesRequest) (_result *GetRelatedWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetRelatedWorkspacesHeaders{}
	_result = &GetRelatedWorkspacesResponse{}
	_body, _err := client.GetRelatedWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetReportTemplateByNameWithOptions(tmpReq *GetReportTemplateByNameRequest, tmpHeader *GetReportTemplateByNameHeaders, runtime *dara.RuntimeOptions) (_result *GetReportTemplateByNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetReportTemplateByNameRequest
//
// @return GetReportTemplateByNameResponse
func (client *Client) GetReportTemplateByName(request *GetReportTemplateByNameRequest) (_result *GetReportTemplateByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetReportTemplateByNameHeaders{}
	_result = &GetReportTemplateByNameResponse{}
	_body, _err := client.GetReportTemplateByNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetReportUnReadCountWithOptions(tmpReq *GetReportUnReadCountRequest, tmpHeader *GetReportUnReadCountHeaders, runtime *dara.RuntimeOptions) (_result *GetReportUnReadCountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetReportUnReadCountRequest
//
// @return GetReportUnReadCountResponse
func (client *Client) GetReportUnReadCount(request *GetReportUnReadCountRequest) (_result *GetReportUnReadCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetReportUnReadCountHeaders{}
	_result = &GetReportUnReadCountResponse{}
	_body, _err := client.GetReportUnReadCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRunningTasksWithOptions(request *GetRunningTasksRequest, tmpHeader *GetRunningTasksHeaders, runtime *dara.RuntimeOptions) (_result *GetRunningTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRunningTasksResponse
func (client *Client) GetRunningTasks(request *GetRunningTasksRequest) (_result *GetRunningTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetRunningTasksHeaders{}
	_result = &GetRunningTasksResponse{}
	_body, _err := client.GetRunningTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetScheduleWithOptions(tmpReq *GetScheduleRequest, tmpHeader *GetScheduleHeaders, runtime *dara.RuntimeOptions) (_result *GetScheduleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetScheduleRequest
//
// @return GetScheduleResponse
func (client *Client) GetSchedule(request *GetScheduleRequest) (_result *GetScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetScheduleHeaders{}
	_result = &GetScheduleResponse{}
	_body, _err := client.GetScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSheetWithOptions(tmpReq *GetSheetRequest, tmpHeader *GetSheetHeaders, runtime *dara.RuntimeOptions) (_result *GetSheetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetSheetRequest
//
// @return GetSheetResponse
func (client *Client) GetSheet(request *GetSheetRequest) (_result *GetSheetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSheetHeaders{}
	_result = &GetSheetResponse{}
	_body, _err := client.GetSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSheetContentJobIdWithOptions(tmpReq *GetSheetContentJobIdRequest, tmpHeader *GetSheetContentJobIdHeaders, runtime *dara.RuntimeOptions) (_result *GetSheetContentJobIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetSheetContentJobIdRequest
//
// @return GetSheetContentJobIdResponse
func (client *Client) GetSheetContentJobId(request *GetSheetContentJobIdRequest) (_result *GetSheetContentJobIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSheetContentJobIdHeaders{}
	_result = &GetSheetContentJobIdResponse{}
	_body, _err := client.GetSheetContentJobIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSpaceDirectoriesWithOptions(tmpReq *GetSpaceDirectoriesRequest, tmpHeader *GetSpaceDirectoriesHeaders, runtime *dara.RuntimeOptions) (_result *GetSpaceDirectoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetSpaceDirectoriesRequest
//
// @return GetSpaceDirectoriesResponse
func (client *Client) GetSpaceDirectories(request *GetSpaceDirectoriesRequest) (_result *GetSpaceDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSpaceDirectoriesHeaders{}
	_result = &GetSpaceDirectoriesResponse{}
	_body, _err := client.GetSpaceDirectoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSubscribedCalendarWithOptions(request *GetSubscribedCalendarRequest, tmpHeader *GetSubscribedCalendarHeaders, runtime *dara.RuntimeOptions) (_result *GetSubscribedCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSubscribedCalendarResponse
func (client *Client) GetSubscribedCalendar(request *GetSubscribedCalendarRequest) (_result *GetSubscribedCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetSubscribedCalendarHeaders{}
	_result = &GetSubscribedCalendarResponse{}
	_body, _err := client.GetSubscribedCalendarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTaskCopiesWithOptions(request *GetTaskCopiesRequest, tmpHeader *GetTaskCopiesHeaders, runtime *dara.RuntimeOptions) (_result *GetTaskCopiesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTaskCopiesResponse
func (client *Client) GetTaskCopies(request *GetTaskCopiesRequest) (_result *GetTaskCopiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetTaskCopiesHeaders{}
	_result = &GetTaskCopiesResponse{}
	_body, _err := client.GetTaskCopiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTemplateListByUserIdWithOptions(tmpReq *GetTemplateListByUserIdRequest, tmpHeader *GetTemplateListByUserIdHeaders, runtime *dara.RuntimeOptions) (_result *GetTemplateListByUserIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetTemplateListByUserIdRequest
//
// @return GetTemplateListByUserIdResponse
func (client *Client) GetTemplateListByUserId(request *GetTemplateListByUserIdRequest) (_result *GetTemplateListByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetTemplateListByUserIdHeaders{}
	_result = &GetTemplateListByUserIdResponse{}
	_body, _err := client.GetTemplateListByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTicketWithOptions(tmpReq *GetTicketRequest, tmpHeader *GetTicketHeaders, runtime *dara.RuntimeOptions) (_result *GetTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetTicketRequest
//
// @return GetTicketResponse
func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetTicketHeaders{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTodoTaskWithOptions(tmpReq *GetTodoTaskRequest, tmpHeader *GetTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *GetTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetTodoTaskRequest
//
// @return GetTodoTaskResponse
func (client *Client) GetTodoTask(request *GetTodoTaskRequest) (_result *GetTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetTodoTaskHeaders{}
	_result = &GetTodoTaskResponse{}
	_body, _err := client.GetTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserWithOptions(tmpReq *GetUserRequest, tmpHeader *GetUserHeaders, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserIdWithOptions(tmpReq *GetUserIdRequest, tmpHeader *GetUserIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetUserIdRequest
//
// @return GetUserIdResponse
func (client *Client) GetUserId(request *GetUserIdRequest) (_result *GetUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdHeaders{}
	_result = &GetUserIdResponse{}
	_body, _err := client.GetUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserIdByOrgIdAndStaffIdWithOptions(tmpReq *GetUserIdByOrgIdAndStaffIdRequest, tmpHeader *GetUserIdByOrgIdAndStaffIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByOrgIdAndStaffIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetUserIdByOrgIdAndStaffIdRequest
//
// @return GetUserIdByOrgIdAndStaffIdResponse
func (client *Client) GetUserIdByOrgIdAndStaffId(request *GetUserIdByOrgIdAndStaffIdRequest) (_result *GetUserIdByOrgIdAndStaffIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdByOrgIdAndStaffIdHeaders{}
	_result = &GetUserIdByOrgIdAndStaffIdResponse{}
	_body, _err := client.GetUserIdByOrgIdAndStaffIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserLatestPlanWithOptions(tmpReq *GetUserLatestPlanRequest, tmpHeader *GetUserLatestPlanHeaders, runtime *dara.RuntimeOptions) (_result *GetUserLatestPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetUserLatestPlanRequest
//
// @return GetUserLatestPlanResponse
func (client *Client) GetUserLatestPlan(request *GetUserLatestPlanRequest) (_result *GetUserLatestPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserLatestPlanHeaders{}
	_result = &GetUserLatestPlanResponse{}
	_body, _err := client.GetUserLatestPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkspaceWithOptions(tmpReq *GetWorkspaceRequest, tmpHeader *GetWorkspaceHeaders, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetWorkspaceHeaders{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkspacesWithOptions(tmpReq *GetWorkspacesRequest, tmpHeader *GetWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *GetWorkspacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetWorkspacesRequest
//
// @return GetWorkspacesResponse
func (client *Client) GetWorkspaces(request *GetWorkspacesRequest) (_result *GetWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetWorkspacesHeaders{}
	_result = &GetWorkspacesResponse{}
	_body, _err := client.GetWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GrantHonorWithOptions(tmpReq *GrantHonorRequest, tmpHeader *GrantHonorHeaders, runtime *dara.RuntimeOptions) (_result *GrantHonorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GrantHonorRequest
//
// @return GrantHonorResponse
func (client *Client) GrantHonor(request *GrantHonorRequest) (_result *GrantHonorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GrantHonorHeaders{}
	_result = &GrantHonorResponse{}
	_body, _err := client.GrantHonorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InitMultipartFileUploadWithOptions(tmpReq *InitMultipartFileUploadRequest, tmpHeader *InitMultipartFileUploadHeaders, runtime *dara.RuntimeOptions) (_result *InitMultipartFileUploadResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InitMultipartFileUploadRequest
//
// @return InitMultipartFileUploadResponse
func (client *Client) InitMultipartFileUpload(request *InitMultipartFileUploadRequest) (_result *InitMultipartFileUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InitMultipartFileUploadHeaders{}
	_result = &InitMultipartFileUploadResponse{}
	_body, _err := client.InitMultipartFileUploadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertColumnsBeforeWithOptions(tmpReq *InsertColumnsBeforeRequest, tmpHeader *InsertColumnsBeforeHeaders, runtime *dara.RuntimeOptions) (_result *InsertColumnsBeforeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertColumnsBeforeRequest
//
// @return InsertColumnsBeforeResponse
func (client *Client) InsertColumnsBefore(request *InsertColumnsBeforeRequest) (_result *InsertColumnsBeforeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertColumnsBeforeHeaders{}
	_result = &InsertColumnsBeforeResponse{}
	_body, _err := client.InsertColumnsBeforeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertContentWithOptionsWithOptions(tmpReq *InsertContentWithOptionsRequest, tmpHeader *InsertContentWithOptionsHeaders, runtime *dara.RuntimeOptions) (_result *InsertContentWithOptionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertContentWithOptionsRequest
//
// @return InsertContentWithOptionsResponse
func (client *Client) InsertContentWithOptions(request *InsertContentWithOptionsRequest) (_result *InsertContentWithOptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertContentWithOptionsHeaders{}
	_result = &InsertContentWithOptionsResponse{}
	_body, _err := client.InsertContentWithOptionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertDropDownListWithOptions(tmpReq *InsertDropDownListRequest, tmpHeader *InsertDropDownListHeaders, runtime *dara.RuntimeOptions) (_result *InsertDropDownListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertDropDownListRequest
//
// @return InsertDropDownListResponse
func (client *Client) InsertDropDownList(request *InsertDropDownListRequest) (_result *InsertDropDownListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertDropDownListHeaders{}
	_result = &InsertDropDownListResponse{}
	_body, _err := client.InsertDropDownListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertMultiDimTableRecordWithOptions(tmpReq *InsertMultiDimTableRecordRequest, tmpHeader *InsertMultiDimTableRecordHeaders, runtime *dara.RuntimeOptions) (_result *InsertMultiDimTableRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertMultiDimTableRecordRequest
//
// @return InsertMultiDimTableRecordResponse
func (client *Client) InsertMultiDimTableRecord(request *InsertMultiDimTableRecordRequest) (_result *InsertMultiDimTableRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertMultiDimTableRecordHeaders{}
	_result = &InsertMultiDimTableRecordResponse{}
	_body, _err := client.InsertMultiDimTableRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InsertRowsBeforeWithOptions(tmpReq *InsertRowsBeforeRequest, tmpHeader *InsertRowsBeforeHeaders, runtime *dara.RuntimeOptions) (_result *InsertRowsBeforeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InsertRowsBeforeRequest
//
// @return InsertRowsBeforeResponse
func (client *Client) InsertRowsBefore(request *InsertRowsBeforeRequest) (_result *InsertRowsBeforeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InsertRowsBeforeHeaders{}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.InsertRowsBeforeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InviteUsersWithOptions(tmpReq *InviteUsersRequest, tmpHeader *InviteUsersHeaders, runtime *dara.RuntimeOptions) (_result *InviteUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InviteUsersRequest
//
// @return InviteUsersResponse
func (client *Client) InviteUsers(request *InviteUsersRequest) (_result *InviteUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InviteUsersHeaders{}
	_result = &InviteUsersResponse{}
	_body, _err := client.InviteUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InvokeAssistantWithSSE(request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions, _yield chan *InvokeAssistantResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeAssistantWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) InvokeAssistantWithOptions(request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions) (_result *InvokeAssistantResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return InvokeAssistantResponse
func (client *Client) InvokeAssistant(request *InvokeAssistantRequest) (_result *InvokeAssistantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InvokeAssistantHeaders{}
	_result = &InvokeAssistantResponse{}
	_body, _err := client.InvokeAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InvokeSkillWithSSE(tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions, _yield chan *InvokeSkillResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeSkillWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, tmpHeader, runtime)
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
func (client *Client) InvokeSkillWithOptions(tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions) (_result *InvokeSkillResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - InvokeSkillRequest
//
// @return InvokeSkillResponse
func (client *Client) InvokeSkill(request *InvokeSkillRequest) (_result *InvokeSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InvokeSkillHeaders{}
	_result = &InvokeSkillResponse{}
	_body, _err := client.InvokeSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, tmpHeader *ListApplicationHeaders, runtime *dara.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListApplicationResponse
func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListApplicationHeaders{}
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCalendarsWithOptions(tmpReq *ListCalendarsRequest, tmpHeader *ListCalendarsHeaders, runtime *dara.RuntimeOptions) (_result *ListCalendarsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListCalendarsRequest
//
// @return ListCalendarsResponse
func (client *Client) ListCalendars(request *ListCalendarsRequest) (_result *ListCalendarsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCalendarsHeaders{}
	_result = &ListCalendarsResponse{}
	_body, _err := client.ListCalendarsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDentriesWithOptions(tmpReq *ListDentriesRequest, tmpHeader *ListDentriesHeaders, runtime *dara.RuntimeOptions) (_result *ListDentriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListDentriesRequest
//
// @return ListDentriesResponse
func (client *Client) ListDentries(request *ListDentriesRequest) (_result *ListDentriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDentriesHeaders{}
	_result = &ListDentriesResponse{}
	_body, _err := client.ListDentriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDriveSpacesWithOptions(tmpReq *ListDriveSpacesRequest, tmpHeader *ListDriveSpacesHeaders, runtime *dara.RuntimeOptions) (_result *ListDriveSpacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListDriveSpacesRequest
//
// @return ListDriveSpacesResponse
func (client *Client) ListDriveSpaces(request *ListDriveSpacesRequest) (_result *ListDriveSpacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDriveSpacesHeaders{}
	_result = &ListDriveSpacesResponse{}
	_body, _err := client.ListDriveSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEventsWithOptions(request *ListEventsRequest, tmpHeader *ListEventsHeaders, runtime *dara.RuntimeOptions) (_result *ListEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEventsResponse
func (client *Client) ListEvents(request *ListEventsRequest) (_result *ListEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListEventsHeaders{}
	_result = &ListEventsResponse{}
	_body, _err := client.ListEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListEventsViewWithOptions(request *ListEventsViewRequest, tmpHeader *ListEventsViewHeaders, runtime *dara.RuntimeOptions) (_result *ListEventsViewResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListEventsViewResponse
func (client *Client) ListEventsView(request *ListEventsViewRequest) (_result *ListEventsViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListEventsViewHeaders{}
	_result = &ListEventsViewResponse{}
	_body, _err := client.ListEventsViewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFormRemarksWithOptions(tmpReq *ListFormRemarksRequest, tmpHeader *ListFormRemarksHeaders, runtime *dara.RuntimeOptions) (_result *ListFormRemarksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFormRemarksRequest
//
// @return ListFormRemarksResponse
func (client *Client) ListFormRemarks(request *ListFormRemarksRequest) (_result *ListFormRemarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListFormRemarksHeaders{}
	_result = &ListFormRemarksResponse{}
	_body, _err := client.ListFormRemarksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMessageWithOptions(request *ListMessageRequest, headers *ListMessageHeaders, runtime *dara.RuntimeOptions) (_result *ListMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListMessageResponse
func (client *Client) ListMessage(request *ListMessageRequest) (_result *ListMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListMessageHeaders{}
	_result = &ListMessageResponse{}
	_body, _err := client.ListMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListMultiDimTableRecordsWithOptions(tmpReq *ListMultiDimTableRecordsRequest, tmpHeader *ListMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *ListMultiDimTableRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListMultiDimTableRecordsRequest
//
// @return ListMultiDimTableRecordsResponse
func (client *Client) ListMultiDimTableRecords(request *ListMultiDimTableRecordsRequest) (_result *ListMultiDimTableRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListMultiDimTableRecordsHeaders{}
	_result = &ListMultiDimTableRecordsResponse{}
	_body, _err := client.ListMultiDimTableRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListNavigationByFormTypeWithOptions(request *ListNavigationByFormTypeRequest, tmpHeader *ListNavigationByFormTypeHeaders, runtime *dara.RuntimeOptions) (_result *ListNavigationByFormTypeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListNavigationByFormTypeResponse
func (client *Client) ListNavigationByFormType(request *ListNavigationByFormTypeRequest) (_result *ListNavigationByFormTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListNavigationByFormTypeHeaders{}
	_result = &ListNavigationByFormTypeResponse{}
	_body, _err := client.ListNavigationByFormTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, tmpHeader *ListNodesHeaders, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListNodesHeaders{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPermissionsWithOptions(tmpReq *ListPermissionsRequest, tmpHeader *ListPermissionsHeaders, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListPermissionsRequest
//
// @return ListPermissionsResponse
func (client *Client) ListPermissions(request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListPermissionsHeaders{}
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListReportWithOptions(tmpReq *ListReportRequest, tmpHeader *ListReportHeaders, runtime *dara.RuntimeOptions) (_result *ListReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListReportRequest
//
// @return ListReportResponse
func (client *Client) ListReport(request *ListReportRequest) (_result *ListReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListReportHeaders{}
	_result = &ListReportResponse{}
	_body, _err := client.ListReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSkillWithOptions(request *ListSkillRequest, tmpHeader *ListSkillHeaders, runtime *dara.RuntimeOptions) (_result *ListSkillResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSkillResponse
func (client *Client) ListSkill(request *ListSkillRequest) (_result *ListSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSkillHeaders{}
	_result = &ListSkillResponse{}
	_body, _err := client.ListSkillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTableDataByFormInstanceIdTableIdWithOptions(request *ListTableDataByFormInstanceIdTableIdRequest, tmpHeader *ListTableDataByFormInstanceIdTableIdHeaders, runtime *dara.RuntimeOptions) (_result *ListTableDataByFormInstanceIdTableIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTableDataByFormInstanceIdTableIdResponse
func (client *Client) ListTableDataByFormInstanceIdTableId(request *ListTableDataByFormInstanceIdTableIdRequest) (_result *ListTableDataByFormInstanceIdTableIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListTableDataByFormInstanceIdTableIdHeaders{}
	_result = &ListTableDataByFormInstanceIdTableIdResponse{}
	_body, _err := client.ListTableDataByFormInstanceIdTableIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTeamsWithOptions(tmpReq *ListTeamsRequest, tmpHeader *ListTeamsHeaders, runtime *dara.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListTeamsRequest
//
// @return ListTeamsResponse
func (client *Client) ListTeams(request *ListTeamsRequest) (_result *ListTeamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListTeamsHeaders{}
	_result = &ListTeamsResponse{}
	_body, _err := client.ListTeamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTemplateWithOptions(tmpReq *ListTemplateRequest, tmpHeader *ListTemplateHeaders, runtime *dara.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListTemplateRequest
//
// @return ListTemplateResponse
func (client *Client) ListTemplate(request *ListTemplateRequest) (_result *ListTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListTemplateHeaders{}
	_result = &ListTemplateResponse{}
	_body, _err := client.ListTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTicketOperateRecordWithOptions(tmpReq *ListTicketOperateRecordRequest, tmpHeader *ListTicketOperateRecordHeaders, runtime *dara.RuntimeOptions) (_result *ListTicketOperateRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListTicketOperateRecordRequest
//
// @return ListTicketOperateRecordResponse
func (client *Client) ListTicketOperateRecord(request *ListTicketOperateRecordRequest) (_result *ListTicketOperateRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListTicketOperateRecordHeaders{}
	_result = &ListTicketOperateRecordResponse{}
	_body, _err := client.ListTicketOperateRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListWorkspacesWithOptions(tmpReq *ListWorkspacesRequest, tmpHeader *ListWorkspacesHeaders, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListWorkspacesHeaders{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PatchEventWithOptions(tmpReq *PatchEventRequest, tmpHeader *PatchEventHeaders, runtime *dara.RuntimeOptions) (_result *PatchEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - PatchEventRequest
//
// @return PatchEventResponse
func (client *Client) PatchEvent(request *PatchEventRequest) (_result *PatchEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PatchEventHeaders{}
	_result = &PatchEventResponse{}
	_body, _err := client.PatchEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryCloudRecordTextWithOptions(tmpReq *QueryCloudRecordTextRequest, tmpHeader *QueryCloudRecordTextHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryCloudRecordTextRequest
//
// @return QueryCloudRecordTextResponse
func (client *Client) QueryCloudRecordText(request *QueryCloudRecordTextRequest) (_result *QueryCloudRecordTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryCloudRecordTextHeaders{}
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.QueryCloudRecordTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryCloudRecordVideoWithOptions(tmpReq *QueryCloudRecordVideoRequest, tmpHeader *QueryCloudRecordVideoHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordVideoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryCloudRecordVideoRequest
//
// @return QueryCloudRecordVideoResponse
func (client *Client) QueryCloudRecordVideo(request *QueryCloudRecordVideoRequest) (_result *QueryCloudRecordVideoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryCloudRecordVideoHeaders{}
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.QueryCloudRecordVideoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryCloudRecordVideoPlayInfoWithOptions(tmpReq *QueryCloudRecordVideoPlayInfoRequest, tmpHeader *QueryCloudRecordVideoPlayInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryCloudRecordVideoPlayInfoRequest
//
// @return QueryCloudRecordVideoPlayInfoResponse
func (client *Client) QueryCloudRecordVideoPlayInfo(request *QueryCloudRecordVideoPlayInfoRequest) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryCloudRecordVideoPlayInfoHeaders{}
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.QueryCloudRecordVideoPlayInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryConferenceInfoWithOptions(request *QueryConferenceInfoRequest, tmpHeader *QueryConferenceInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryConferenceInfoResponse
func (client *Client) QueryConferenceInfo(request *QueryConferenceInfoRequest) (_result *QueryConferenceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryConferenceInfoHeaders{}
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.QueryConferenceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryConferenceInfoByRoomCodeWithOptions(tmpReq *QueryConferenceInfoByRoomCodeRequest, tmpHeader *QueryConferenceInfoByRoomCodeHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceInfoByRoomCodeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryConferenceInfoByRoomCodeRequest
//
// @return QueryConferenceInfoByRoomCodeResponse
func (client *Client) QueryConferenceInfoByRoomCode(request *QueryConferenceInfoByRoomCodeRequest) (_result *QueryConferenceInfoByRoomCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryConferenceInfoByRoomCodeHeaders{}
	_result = &QueryConferenceInfoByRoomCodeResponse{}
	_body, _err := client.QueryConferenceInfoByRoomCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryConferenceMembersWithOptions(tmpReq *QueryConferenceMembersRequest, tmpHeader *QueryConferenceMembersHeaders, runtime *dara.RuntimeOptions) (_result *QueryConferenceMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryConferenceMembersRequest
//
// @return QueryConferenceMembersResponse
func (client *Client) QueryConferenceMembers(request *QueryConferenceMembersRequest) (_result *QueryConferenceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryConferenceMembersHeaders{}
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.QueryConferenceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDentriesInfoWithOptions(tmpReq *QueryDentriesInfoRequest, tmpHeader *QueryDentriesInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryDentriesInfoResponse, _err error) {
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

	if !dara.IsNil(request.UnionId) {
		body["UnionId"] = request.UnionId
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryDentriesInfoRequest
//
// @return QueryDentriesInfoResponse
func (client *Client) QueryDentriesInfo(request *QueryDentriesInfoRequest) (_result *QueryDentriesInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryDentriesInfoHeaders{}
	_result = &QueryDentriesInfoResponse{}
	_body, _err := client.QueryDentriesInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryDentryWithOptions(tmpReq *QueryDentryRequest, tmpHeader *QueryDentryHeaders, runtime *dara.RuntimeOptions) (_result *QueryDentryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryDentryRequest
//
// @return QueryDentryResponse
func (client *Client) QueryDentry(request *QueryDentryRequest) (_result *QueryDentryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryDentryHeaders{}
	_result = &QueryDentryResponse{}
	_body, _err := client.QueryDentryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryGroupLiveInfoWithOptions(tmpReq *QueryGroupLiveInfoRequest, tmpHeader *QueryGroupLiveInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryGroupLiveInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryGroupLiveInfoRequest
//
// @return QueryGroupLiveInfoResponse
func (client *Client) QueryGroupLiveInfo(request *QueryGroupLiveInfoRequest) (_result *QueryGroupLiveInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryGroupLiveInfoHeaders{}
	_result = &QueryGroupLiveInfoResponse{}
	_body, _err := client.QueryGroupLiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryLiveInfoWithOptions(tmpReq *QueryLiveInfoRequest, tmpHeader *QueryLiveInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryLiveInfoRequest
//
// @return QueryLiveInfoResponse
func (client *Client) QueryLiveInfo(request *QueryLiveInfoRequest) (_result *QueryLiveInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryLiveInfoHeaders{}
	_result = &QueryLiveInfoResponse{}
	_body, _err := client.QueryLiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryLiveWatchDetailWithOptions(tmpReq *QueryLiveWatchDetailRequest, tmpHeader *QueryLiveWatchDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveWatchDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryLiveWatchDetailRequest
//
// @return QueryLiveWatchDetailResponse
func (client *Client) QueryLiveWatchDetail(request *QueryLiveWatchDetailRequest) (_result *QueryLiveWatchDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryLiveWatchDetailHeaders{}
	_result = &QueryLiveWatchDetailResponse{}
	_body, _err := client.QueryLiveWatchDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryLiveWatchUserListWithOptions(tmpReq *QueryLiveWatchUserListRequest, tmpHeader *QueryLiveWatchUserListHeaders, runtime *dara.RuntimeOptions) (_result *QueryLiveWatchUserListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryLiveWatchUserListRequest
//
// @return QueryLiveWatchUserListResponse
func (client *Client) QueryLiveWatchUserList(request *QueryLiveWatchUserListRequest) (_result *QueryLiveWatchUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryLiveWatchUserListHeaders{}
	_result = &QueryLiveWatchUserListResponse{}
	_body, _err := client.QueryLiveWatchUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMeetingRoomWithOptions(tmpReq *QueryMeetingRoomRequest, tmpHeader *QueryMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMeetingRoomRequest
//
// @return QueryMeetingRoomResponse
func (client *Client) QueryMeetingRoom(request *QueryMeetingRoomRequest) (_result *QueryMeetingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMeetingRoomHeaders{}
	_result = &QueryMeetingRoomResponse{}
	_body, _err := client.QueryMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMeetingRoomGroupWithOptions(tmpReq *QueryMeetingRoomGroupRequest, tmpHeader *QueryMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMeetingRoomGroupRequest
//
// @return QueryMeetingRoomGroupResponse
func (client *Client) QueryMeetingRoomGroup(request *QueryMeetingRoomGroupRequest) (_result *QueryMeetingRoomGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMeetingRoomGroupHeaders{}
	_result = &QueryMeetingRoomGroupResponse{}
	_body, _err := client.QueryMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMeetingRoomGroupListWithOptions(tmpReq *QueryMeetingRoomGroupListRequest, tmpHeader *QueryMeetingRoomGroupListHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomGroupListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMeetingRoomGroupListRequest
//
// @return QueryMeetingRoomGroupListResponse
func (client *Client) QueryMeetingRoomGroupList(request *QueryMeetingRoomGroupListRequest) (_result *QueryMeetingRoomGroupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMeetingRoomGroupListHeaders{}
	_result = &QueryMeetingRoomGroupListResponse{}
	_body, _err := client.QueryMeetingRoomGroupListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMeetingRoomListWithOptions(tmpReq *QueryMeetingRoomListRequest, tmpHeader *QueryMeetingRoomListHeaders, runtime *dara.RuntimeOptions) (_result *QueryMeetingRoomListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMeetingRoomListRequest
//
// @return QueryMeetingRoomListResponse
func (client *Client) QueryMeetingRoomList(request *QueryMeetingRoomListRequest) (_result *QueryMeetingRoomListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMeetingRoomListHeaders{}
	_result = &QueryMeetingRoomListResponse{}
	_body, _err := client.QueryMeetingRoomListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMinutesWithOptions(tmpReq *QueryMinutesRequest, tmpHeader *QueryMinutesHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMinutesRequest
//
// @return QueryMinutesResponse
func (client *Client) QueryMinutes(request *QueryMinutesRequest) (_result *QueryMinutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMinutesHeaders{}
	_result = &QueryMinutesResponse{}
	_body, _err := client.QueryMinutesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMinutesSummaryWithOptions(tmpReq *QueryMinutesSummaryRequest, tmpHeader *QueryMinutesSummaryHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMinutesSummaryRequest
//
// @return QueryMinutesSummaryResponse
func (client *Client) QueryMinutesSummary(request *QueryMinutesSummaryRequest) (_result *QueryMinutesSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMinutesSummaryHeaders{}
	_result = &QueryMinutesSummaryResponse{}
	_body, _err := client.QueryMinutesSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryMinutesTextWithOptions(tmpReq *QueryMinutesTextRequest, tmpHeader *QueryMinutesTextHeaders, runtime *dara.RuntimeOptions) (_result *QueryMinutesTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryMinutesTextRequest
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesText(request *QueryMinutesTextRequest) (_result *QueryMinutesTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMinutesTextHeaders{}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.QueryMinutesTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryOrgHonorsWithOptions(tmpReq *QueryOrgHonorsRequest, tmpHeader *QueryOrgHonorsHeaders, runtime *dara.RuntimeOptions) (_result *QueryOrgHonorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryOrgHonorsRequest
//
// @return QueryOrgHonorsResponse
func (client *Client) QueryOrgHonors(request *QueryOrgHonorsRequest) (_result *QueryOrgHonorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryOrgHonorsHeaders{}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.QueryOrgHonorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryOrgTodoTasksWithOptions(tmpReq *QueryOrgTodoTasksRequest, tmpHeader *QueryOrgTodoTasksHeaders, runtime *dara.RuntimeOptions) (_result *QueryOrgTodoTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryOrgTodoTasksRequest
//
// @return QueryOrgTodoTasksResponse
func (client *Client) QueryOrgTodoTasks(request *QueryOrgTodoTasksRequest) (_result *QueryOrgTodoTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryOrgTodoTasksHeaders{}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.QueryOrgTodoTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryRecordMinutesUrlWithOptions(tmpReq *QueryRecordMinutesUrlRequest, tmpHeader *QueryRecordMinutesUrlHeaders, runtime *dara.RuntimeOptions) (_result *QueryRecordMinutesUrlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryRecordMinutesUrlRequest
//
// @return QueryRecordMinutesUrlResponse
func (client *Client) QueryRecordMinutesUrl(request *QueryRecordMinutesUrlRequest) (_result *QueryRecordMinutesUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryRecordMinutesUrlHeaders{}
	_result = &QueryRecordMinutesUrlResponse{}
	_body, _err := client.QueryRecordMinutesUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryReportDetailWithOptions(tmpReq *QueryReportDetailRequest, tmpHeader *QueryReportDetailHeaders, runtime *dara.RuntimeOptions) (_result *QueryReportDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryReportDetailRequest
//
// @return QueryReportDetailResponse
func (client *Client) QueryReportDetail(request *QueryReportDetailRequest) (_result *QueryReportDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryReportDetailHeaders{}
	_result = &QueryReportDetailResponse{}
	_body, _err := client.QueryReportDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryScheduleConferenceWithOptions(tmpReq *QueryScheduleConferenceRequest, tmpHeader *QueryScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *QueryScheduleConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryScheduleConferenceRequest
//
// @return QueryScheduleConferenceResponse
func (client *Client) QueryScheduleConference(request *QueryScheduleConferenceRequest) (_result *QueryScheduleConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryScheduleConferenceHeaders{}
	_result = &QueryScheduleConferenceResponse{}
	_body, _err := client.QueryScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryScheduleConferenceInfoWithOptions(tmpReq *QueryScheduleConferenceInfoRequest, tmpHeader *QueryScheduleConferenceInfoHeaders, runtime *dara.RuntimeOptions) (_result *QueryScheduleConferenceInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryScheduleConferenceInfoRequest
//
// @return QueryScheduleConferenceInfoResponse
func (client *Client) QueryScheduleConferenceInfo(request *QueryScheduleConferenceInfoRequest) (_result *QueryScheduleConferenceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryScheduleConferenceInfoHeaders{}
	_result = &QueryScheduleConferenceInfoResponse{}
	_body, _err := client.QueryScheduleConferenceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryUserHonorsWithOptions(tmpReq *QueryUserHonorsRequest, tmpHeader *QueryUserHonorsHeaders, runtime *dara.RuntimeOptions) (_result *QueryUserHonorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryUserHonorsRequest
//
// @return QueryUserHonorsResponse
func (client *Client) QueryUserHonors(request *QueryUserHonorsRequest) (_result *QueryUserHonorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryUserHonorsHeaders{}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.QueryUserHonorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RecallHonorWithOptions(tmpReq *RecallHonorRequest, tmpHeader *RecallHonorHeaders, runtime *dara.RuntimeOptions) (_result *RecallHonorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RecallHonorRequest
//
// @return RecallHonorResponse
func (client *Client) RecallHonor(request *RecallHonorRequest) (_result *RecallHonorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RecallHonorHeaders{}
	_result = &RecallHonorResponse{}
	_body, _err := client.RecallHonorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReceiverListReportWithOptions(tmpReq *ReceiverListReportRequest, tmpHeader *ReceiverListReportHeaders, runtime *dara.RuntimeOptions) (_result *ReceiverListReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ReceiverListReportRequest
//
// @return ReceiverListReportResponse
func (client *Client) ReceiverListReport(request *ReceiverListReportRequest) (_result *ReceiverListReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ReceiverListReportHeaders{}
	_result = &ReceiverListReportResponse{}
	_body, _err := client.ReceiverListReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RedirectTaskWithOptions(request *RedirectTaskRequest, tmpHeader *RedirectTaskHeaders, runtime *dara.RuntimeOptions) (_result *RedirectTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RedirectTaskResponse
func (client *Client) RedirectTask(request *RedirectTaskRequest) (_result *RedirectTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RedirectTaskHeaders{}
	_result = &RedirectTaskResponse{}
	_body, _err := client.RedirectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveAttendeeWithOptions(tmpReq *RemoveAttendeeRequest, tmpHeader *RemoveAttendeeHeaders, runtime *dara.RuntimeOptions) (_result *RemoveAttendeeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RemoveAttendeeRequest
//
// @return RemoveAttendeeResponse
func (client *Client) RemoveAttendee(request *RemoveAttendeeRequest) (_result *RemoveAttendeeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveAttendeeHeaders{}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.RemoveAttendeeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveMeetingRoomsWithOptions(tmpReq *RemoveMeetingRoomsRequest, tmpHeader *RemoveMeetingRoomsHeaders, runtime *dara.RuntimeOptions) (_result *RemoveMeetingRoomsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RemoveMeetingRoomsRequest
//
// @return RemoveMeetingRoomsResponse
func (client *Client) RemoveMeetingRooms(request *RemoveMeetingRoomsRequest) (_result *RemoveMeetingRoomsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveMeetingRoomsHeaders{}
	_result = &RemoveMeetingRoomsResponse{}
	_body, _err := client.RemoveMeetingRoomsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RespondEventWithOptions(tmpReq *RespondEventRequest, tmpHeader *RespondEventHeaders, runtime *dara.RuntimeOptions) (_result *RespondEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - RespondEventRequest
//
// @return RespondEventResponse
func (client *Client) RespondEvent(request *RespondEventRequest) (_result *RespondEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RespondEventHeaders{}
	_result = &RespondEventResponse{}
	_body, _err := client.RespondEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RetrieveRunWithOptions(request *RetrieveRunRequest, headers *RetrieveRunHeaders, runtime *dara.RuntimeOptions) (_result *RetrieveRunResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RetrieveRunResponse
func (client *Client) RetrieveRun(request *RetrieveRunRequest) (_result *RetrieveRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RetrieveRunHeaders{}
	_result = &RetrieveRunResponse{}
	_body, _err := client.RetrieveRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveContentWithOptions(tmpReq *SaveContentRequest, tmpHeader *SaveContentHeaders, runtime *dara.RuntimeOptions) (_result *SaveContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SaveContentRequest
//
// @return SaveContentResponse
func (client *Client) SaveContent(request *SaveContentRequest) (_result *SaveContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SaveContentHeaders{}
	_result = &SaveContentResponse{}
	_body, _err := client.SaveContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, tmpHeader *SaveFormDataHeaders, runtime *dara.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SaveFormDataResponse
func (client *Client) SaveFormData(request *SaveFormDataRequest) (_result *SaveFormDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SaveFormDataHeaders{}
	_result = &SaveFormDataResponse{}
	_body, _err := client.SaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveFormRemarkWithOptions(request *SaveFormRemarkRequest, tmpHeader *SaveFormRemarkHeaders, runtime *dara.RuntimeOptions) (_result *SaveFormRemarkResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SaveFormRemarkResponse
func (client *Client) SaveFormRemark(request *SaveFormRemarkRequest) (_result *SaveFormRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SaveFormRemarkHeaders{}
	_result = &SaveFormRemarkResponse{}
	_body, _err := client.SaveFormRemarkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchEmployeeFieldValuesWithOptions(request *SearchEmployeeFieldValuesRequest, tmpHeader *SearchEmployeeFieldValuesHeaders, runtime *dara.RuntimeOptions) (_result *SearchEmployeeFieldValuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchEmployeeFieldValuesResponse
func (client *Client) SearchEmployeeFieldValues(request *SearchEmployeeFieldValuesRequest) (_result *SearchEmployeeFieldValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchEmployeeFieldValuesHeaders{}
	_result = &SearchEmployeeFieldValuesResponse{}
	_body, _err := client.SearchEmployeeFieldValuesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchFormDataIdListWithOptions(request *SearchFormDataIdListRequest, tmpHeader *SearchFormDataIdListHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataIdListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdList(request *SearchFormDataIdListRequest) (_result *SearchFormDataIdListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchFormDataIdListHeaders{}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.SearchFormDataIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchFormDataSecondGenerationWithOptions(request *SearchFormDataSecondGenerationRequest, tmpHeader *SearchFormDataSecondGenerationHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataSecondGenerationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGeneration(request *SearchFormDataSecondGenerationRequest) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationHeaders{}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.SearchFormDataSecondGenerationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchFormDataSecondGenerationNoTableFieldWithOptions(request *SearchFormDataSecondGenerationNoTableFieldRequest, tmpHeader *SearchFormDataSecondGenerationNoTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDataSecondGenerationNoTableFieldResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchFormDataSecondGenerationNoTableFieldResponse
func (client *Client) SearchFormDataSecondGenerationNoTableField(request *SearchFormDataSecondGenerationNoTableFieldRequest) (_result *SearchFormDataSecondGenerationNoTableFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationNoTableFieldHeaders{}
	_result = &SearchFormDataSecondGenerationNoTableFieldResponse{}
	_body, _err := client.SearchFormDataSecondGenerationNoTableFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchFormDatasWithOptions(request *SearchFormDatasRequest, tmpHeader *SearchFormDatasHeaders, runtime *dara.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchFormDatasResponse
func (client *Client) SearchFormDatas(request *SearchFormDatasRequest) (_result *SearchFormDatasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchFormDatasHeaders{}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.SearchFormDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchInnerGroupsWithOptions(request *SearchInnerGroupsRequest, tmpHeader *SearchInnerGroupsHeaders, runtime *dara.RuntimeOptions) (_result *SearchInnerGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchInnerGroupsResponse
func (client *Client) SearchInnerGroups(request *SearchInnerGroupsRequest) (_result *SearchInnerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchInnerGroupsHeaders{}
	_result = &SearchInnerGroupsResponse{}
	_body, _err := client.SearchInnerGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendBannerWithOptions(tmpReq *SendBannerRequest, tmpHeader *SendBannerHeaders, runtime *dara.RuntimeOptions) (_result *SendBannerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SendBannerRequest
//
// @return SendBannerResponse
func (client *Client) SendBanner(request *SendBannerRequest) (_result *SendBannerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendBannerHeaders{}
	_result = &SendBannerResponse{}
	_body, _err := client.SendBannerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendPopupWithOptions(tmpReq *SendPopupRequest, tmpHeader *SendPopupHeaders, runtime *dara.RuntimeOptions) (_result *SendPopupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SendPopupRequest
//
// @return SendPopupResponse
func (client *Client) SendPopup(request *SendPopupRequest) (_result *SendPopupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendPopupHeaders{}
	_result = &SendPopupResponse{}
	_body, _err := client.SendPopupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendSearchShadeWithOptions(tmpReq *SendSearchShadeRequest, tmpHeader *SendSearchShadeHeaders, runtime *dara.RuntimeOptions) (_result *SendSearchShadeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SendSearchShadeRequest
//
// @return SendSearchShadeResponse
func (client *Client) SendSearchShade(request *SendSearchShadeRequest) (_result *SendSearchShadeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendSearchShadeHeaders{}
	_result = &SendSearchShadeResponse{}
	_body, _err := client.SendSearchShadeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetColumnsVisibilityWithOptions(tmpReq *SetColumnsVisibilityRequest, tmpHeader *SetColumnsVisibilityHeaders, runtime *dara.RuntimeOptions) (_result *SetColumnsVisibilityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SetColumnsVisibilityRequest
//
// @return SetColumnsVisibilityResponse
func (client *Client) SetColumnsVisibility(request *SetColumnsVisibilityRequest) (_result *SetColumnsVisibilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SetColumnsVisibilityHeaders{}
	_result = &SetColumnsVisibilityResponse{}
	_body, _err := client.SetColumnsVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetConferenceHostsWithOptions(tmpReq *SetConferenceHostsRequest, tmpHeader *SetConferenceHostsHeaders, runtime *dara.RuntimeOptions) (_result *SetConferenceHostsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SetConferenceHostsRequest
//
// @return SetConferenceHostsResponse
func (client *Client) SetConferenceHosts(request *SetConferenceHostsRequest) (_result *SetConferenceHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SetConferenceHostsHeaders{}
	_result = &SetConferenceHostsResponse{}
	_body, _err := client.SetConferenceHostsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetRowsVisibilityWithOptions(tmpReq *SetRowsVisibilityRequest, tmpHeader *SetRowsVisibilityHeaders, runtime *dara.RuntimeOptions) (_result *SetRowsVisibilityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SetRowsVisibilityRequest
//
// @return SetRowsVisibilityResponse
func (client *Client) SetRowsVisibility(request *SetRowsVisibilityRequest) (_result *SetRowsVisibilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SetRowsVisibilityHeaders{}
	_result = &SetRowsVisibilityResponse{}
	_body, _err := client.SetRowsVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SimpleListReportWithOptions(tmpReq *SimpleListReportRequest, tmpHeader *SimpleListReportHeaders, runtime *dara.RuntimeOptions) (_result *SimpleListReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SimpleListReportRequest
//
// @return SimpleListReportResponse
func (client *Client) SimpleListReport(request *SimpleListReportRequest) (_result *SimpleListReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SimpleListReportHeaders{}
	_result = &SimpleListReportResponse{}
	_body, _err := client.SimpleListReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartCloudRecordWithOptions(tmpReq *StartCloudRecordRequest, tmpHeader *StartCloudRecordHeaders, runtime *dara.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StartCloudRecordRequest
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecord(request *StartCloudRecordRequest) (_result *StartCloudRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartCloudRecordHeaders{}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.StartCloudRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, tmpHeader *StartInstanceHeaders, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartInstanceHeaders{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartMinutesWithOptions(tmpReq *StartMinutesRequest, tmpHeader *StartMinutesHeaders, runtime *dara.RuntimeOptions) (_result *StartMinutesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StartMinutesRequest
//
// @return StartMinutesResponse
func (client *Client) StartMinutes(request *StartMinutesRequest) (_result *StartMinutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StartMinutesHeaders{}
	_result = &StartMinutesResponse{}
	_body, _err := client.StartMinutesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StatisticsListByTypeReportWithOptions(tmpReq *StatisticsListByTypeReportRequest, tmpHeader *StatisticsListByTypeReportHeaders, runtime *dara.RuntimeOptions) (_result *StatisticsListByTypeReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StatisticsListByTypeReportRequest
//
// @return StatisticsListByTypeReportResponse
func (client *Client) StatisticsListByTypeReport(request *StatisticsListByTypeReportRequest) (_result *StatisticsListByTypeReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StatisticsListByTypeReportHeaders{}
	_result = &StatisticsListByTypeReportResponse{}
	_body, _err := client.StatisticsListByTypeReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StatisticsReportWithOptions(tmpReq *StatisticsReportRequest, tmpHeader *StatisticsReportHeaders, runtime *dara.RuntimeOptions) (_result *StatisticsReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StatisticsReportRequest
//
// @return StatisticsReportResponse
func (client *Client) StatisticsReport(request *StatisticsReportRequest) (_result *StatisticsReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StatisticsReportHeaders{}
	_result = &StatisticsReportResponse{}
	_body, _err := client.StatisticsReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopCloudRecordWithOptions(tmpReq *StopCloudRecordRequest, tmpHeader *StopCloudRecordHeaders, runtime *dara.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StopCloudRecordRequest
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecord(request *StopCloudRecordRequest) (_result *StopCloudRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StopCloudRecordHeaders{}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.StopCloudRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopMinutesWithOptions(tmpReq *StopMinutesRequest, tmpHeader *StopMinutesHeaders, runtime *dara.RuntimeOptions) (_result *StopMinutesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - StopMinutesRequest
//
// @return StopMinutesResponse
func (client *Client) StopMinutes(request *StopMinutesRequest) (_result *StopMinutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &StopMinutesHeaders{}
	_result = &StopMinutesResponse{}
	_body, _err := client.StopMinutesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubscribeCalendarWithOptions(request *SubscribeCalendarRequest, tmpHeader *SubscribeCalendarHeaders, runtime *dara.RuntimeOptions) (_result *SubscribeCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubscribeCalendarResponse
func (client *Client) SubscribeCalendar(request *SubscribeCalendarRequest) (_result *SubscribeCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SubscribeCalendarHeaders{}
	_result = &SubscribeCalendarResponse{}
	_body, _err := client.SubscribeCalendarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubscribeEventWithOptions(tmpReq *SubscribeEventRequest, tmpHeader *SubscribeEventHeaders, runtime *dara.RuntimeOptions) (_result *SubscribeEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SubscribeEventRequest
//
// @return SubscribeEventResponse
func (client *Client) SubscribeEvent(request *SubscribeEventRequest) (_result *SubscribeEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SubscribeEventHeaders{}
	_result = &SubscribeEventResponse{}
	_body, _err := client.SubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - SwitchMainOrgRequest
//
// @param tmpHeader - SwitchMainOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchMainOrgResponse
func (client *Client) SwitchMainOrgWithOptions(tmpReq *SwitchMainOrgRequest, tmpHeader *SwitchMainOrgHeaders, runtime *dara.RuntimeOptions) (_result *SwitchMainOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchMainOrgRequest
//
// @return SwitchMainOrgResponse
func (client *Client) SwitchMainOrg(request *SwitchMainOrgRequest) (_result *SwitchMainOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SwitchMainOrgHeaders{}
	_result = &SwitchMainOrgResponse{}
	_body, _err := client.SwitchMainOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SyncDingTypeWithOptions(tmpReq *SyncDingTypeRequest, tmpHeader *SyncDingTypeHeaders, runtime *dara.RuntimeOptions) (_result *SyncDingTypeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SyncDingTypeRequest
//
// @return SyncDingTypeResponse
func (client *Client) SyncDingType(request *SyncDingTypeRequest) (_result *SyncDingTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SyncDingTypeHeaders{}
	_result = &SyncDingTypeResponse{}
	_body, _err := client.SyncDingTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TerminateInstanceWithOptions(request *TerminateInstanceRequest, tmpHeader *TerminateInstanceHeaders, runtime *dara.RuntimeOptions) (_result *TerminateInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TerminateInstanceResponse
func (client *Client) TerminateInstance(request *TerminateInstanceRequest) (_result *TerminateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &TerminateInstanceHeaders{}
	_result = &TerminateInstanceResponse{}
	_body, _err := client.TerminateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TransferTicketWithOptions(tmpReq *TransferTicketRequest, tmpHeader *TransferTicketHeaders, runtime *dara.RuntimeOptions) (_result *TransferTicketResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - TransferTicketRequest
//
// @return TransferTicketResponse
func (client *Client) TransferTicket(request *TransferTicketRequest) (_result *TransferTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &TransferTicketHeaders{}
	_result = &TransferTicketResponse{}
	_body, _err := client.TransferTicketWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnsubscribeCalendarWithOptions(request *UnsubscribeCalendarRequest, tmpHeader *UnsubscribeCalendarHeaders, runtime *dara.RuntimeOptions) (_result *UnsubscribeCalendarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnsubscribeCalendarResponse
func (client *Client) UnsubscribeCalendar(request *UnsubscribeCalendarRequest) (_result *UnsubscribeCalendarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UnsubscribeCalendarHeaders{}
	_result = &UnsubscribeCalendarResponse{}
	_body, _err := client.UnsubscribeCalendarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnsubscribeEventWithOptions(tmpReq *UnsubscribeEventRequest, tmpHeader *UnsubscribeEventHeaders, runtime *dara.RuntimeOptions) (_result *UnsubscribeEventResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UnsubscribeEventRequest
//
// @return UnsubscribeEventResponse
func (client *Client) UnsubscribeEvent(request *UnsubscribeEventRequest) (_result *UnsubscribeEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UnsubscribeEventHeaders{}
	_result = &UnsubscribeEventResponse{}
	_body, _err := client.UnsubscribeEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateAlidingAssistantRequest
//
// @param tmpHeader - UpdateAlidingAssistantHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlidingAssistantResponse
func (client *Client) UpdateAlidingAssistantWithOptions(tmpReq *UpdateAlidingAssistantRequest, tmpHeader *UpdateAlidingAssistantHeaders, runtime *dara.RuntimeOptions) (_result *UpdateAlidingAssistantResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAlidingAssistantRequest
//
// @return UpdateAlidingAssistantResponse
func (client *Client) UpdateAlidingAssistant(request *UpdateAlidingAssistantRequest) (_result *UpdateAlidingAssistantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateAlidingAssistantHeaders{}
	_result = &UpdateAlidingAssistantResponse{}
	_body, _err := client.UpdateAlidingAssistantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateConvExtensionRequest
//
// @param tmpHeader - UpdateConvExtensionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConvExtensionResponse
func (client *Client) UpdateConvExtensionWithOptions(tmpReq *UpdateConvExtensionRequest, tmpHeader *UpdateConvExtensionHeaders, runtime *dara.RuntimeOptions) (_result *UpdateConvExtensionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateConvExtensionRequest
//
// @return UpdateConvExtensionResponse
func (client *Client) UpdateConvExtension(request *UpdateConvExtensionRequest) (_result *UpdateConvExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateConvExtensionHeaders{}
	_result = &UpdateConvExtensionResponse{}
	_body, _err := client.UpdateConvExtensionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFormDataWithOptions(request *UpdateFormDataRequest, tmpHeader *UpdateFormDataHeaders, runtime *dara.RuntimeOptions) (_result *UpdateFormDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFormDataResponse
func (client *Client) UpdateFormData(request *UpdateFormDataRequest) (_result *UpdateFormDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateFormDataHeaders{}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.UpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, tmpHeader *UpdateInstanceHeaders, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateInstanceHeaders{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateLiveWithOptions(tmpReq *UpdateLiveRequest, tmpHeader *UpdateLiveHeaders, runtime *dara.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateLiveRequest
//
// @return UpdateLiveResponse
func (client *Client) UpdateLive(request *UpdateLiveRequest) (_result *UpdateLiveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateLiveHeaders{}
	_result = &UpdateLiveResponse{}
	_body, _err := client.UpdateLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMeetingRoomWithOptions(tmpReq *UpdateMeetingRoomRequest, tmpHeader *UpdateMeetingRoomHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMeetingRoomResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMeetingRoomRequest
//
// @return UpdateMeetingRoomResponse
func (client *Client) UpdateMeetingRoom(request *UpdateMeetingRoomRequest) (_result *UpdateMeetingRoomResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMeetingRoomHeaders{}
	_result = &UpdateMeetingRoomResponse{}
	_body, _err := client.UpdateMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMeetingRoomGroupWithOptions(tmpReq *UpdateMeetingRoomGroupRequest, tmpHeader *UpdateMeetingRoomGroupHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMeetingRoomGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMeetingRoomGroupRequest
//
// @return UpdateMeetingRoomGroupResponse
func (client *Client) UpdateMeetingRoomGroup(request *UpdateMeetingRoomGroupRequest) (_result *UpdateMeetingRoomGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMeetingRoomGroupHeaders{}
	_result = &UpdateMeetingRoomGroupResponse{}
	_body, _err := client.UpdateMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMultiDimTableWithOptions(tmpReq *UpdateMultiDimTableRequest, tmpHeader *UpdateMultiDimTableHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMultiDimTableRequest
//
// @return UpdateMultiDimTableResponse
func (client *Client) UpdateMultiDimTable(request *UpdateMultiDimTableRequest) (_result *UpdateMultiDimTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMultiDimTableHeaders{}
	_result = &UpdateMultiDimTableResponse{}
	_body, _err := client.UpdateMultiDimTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMultiDimTableFieldWithOptions(tmpReq *UpdateMultiDimTableFieldRequest, tmpHeader *UpdateMultiDimTableFieldHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableFieldResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMultiDimTableFieldRequest
//
// @return UpdateMultiDimTableFieldResponse
func (client *Client) UpdateMultiDimTableField(request *UpdateMultiDimTableFieldRequest) (_result *UpdateMultiDimTableFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMultiDimTableFieldHeaders{}
	_result = &UpdateMultiDimTableFieldResponse{}
	_body, _err := client.UpdateMultiDimTableFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMultiDimTableRecordsWithOptions(tmpReq *UpdateMultiDimTableRecordsRequest, tmpHeader *UpdateMultiDimTableRecordsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateMultiDimTableRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateMultiDimTableRecordsRequest
//
// @return UpdateMultiDimTableRecordsResponse
func (client *Client) UpdateMultiDimTableRecords(request *UpdateMultiDimTableRecordsRequest) (_result *UpdateMultiDimTableRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateMultiDimTableRecordsHeaders{}
	_result = &UpdateMultiDimTableRecordsResponse{}
	_body, _err := client.UpdateMultiDimTableRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePermissionWithOptions(tmpReq *UpdatePermissionRequest, tmpHeader *UpdatePermissionHeaders, runtime *dara.RuntimeOptions) (_result *UpdatePermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdatePermissionRequest
//
// @return UpdatePermissionResponse
func (client *Client) UpdatePermission(request *UpdatePermissionRequest) (_result *UpdatePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdatePermissionHeaders{}
	_result = &UpdatePermissionResponse{}
	_body, _err := client.UpdatePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRangeWithOptions(tmpReq *UpdateRangeRequest, tmpHeader *UpdateRangeHeaders, runtime *dara.RuntimeOptions) (_result *UpdateRangeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateRangeRequest
//
// @return UpdateRangeResponse
func (client *Client) UpdateRange(request *UpdateRangeRequest) (_result *UpdateRangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateRangeHeaders{}
	_result = &UpdateRangeResponse{}
	_body, _err := client.UpdateRangeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateScheduleConfSettingsWithOptions(tmpReq *UpdateScheduleConfSettingsRequest, tmpHeader *UpdateScheduleConfSettingsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduleConfSettingsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateScheduleConfSettingsRequest
//
// @return UpdateScheduleConfSettingsResponse
func (client *Client) UpdateScheduleConfSettings(request *UpdateScheduleConfSettingsRequest) (_result *UpdateScheduleConfSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateScheduleConfSettingsHeaders{}
	_result = &UpdateScheduleConfSettingsResponse{}
	_body, _err := client.UpdateScheduleConfSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateScheduleConferenceWithOptions(tmpReq *UpdateScheduleConferenceRequest, tmpHeader *UpdateScheduleConferenceHeaders, runtime *dara.RuntimeOptions) (_result *UpdateScheduleConferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateScheduleConferenceRequest
//
// @return UpdateScheduleConferenceResponse
func (client *Client) UpdateScheduleConference(request *UpdateScheduleConferenceRequest) (_result *UpdateScheduleConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateScheduleConferenceHeaders{}
	_result = &UpdateScheduleConferenceResponse{}
	_body, _err := client.UpdateScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateStatusWithOptions(tmpReq *UpdateStatusRequest, tmpHeader *UpdateStatusHeaders, runtime *dara.RuntimeOptions) (_result *UpdateStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateStatusRequest
//
// @return UpdateStatusResponse
func (client *Client) UpdateStatus(request *UpdateStatusRequest) (_result *UpdateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateStatusHeaders{}
	_result = &UpdateStatusResponse{}
	_body, _err := client.UpdateStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateSubscribedCalendarsWithOptions(tmpReq *UpdateSubscribedCalendarsRequest, tmpHeader *UpdateSubscribedCalendarsHeaders, runtime *dara.RuntimeOptions) (_result *UpdateSubscribedCalendarsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateSubscribedCalendarsRequest
//
// @return UpdateSubscribedCalendarsResponse
func (client *Client) UpdateSubscribedCalendars(request *UpdateSubscribedCalendarsRequest) (_result *UpdateSubscribedCalendarsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateSubscribedCalendarsHeaders{}
	_result = &UpdateSubscribedCalendarsResponse{}
	_body, _err := client.UpdateSubscribedCalendarsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTodoTaskWithOptions(tmpReq *UpdateTodoTaskRequest, tmpHeader *UpdateTodoTaskHeaders, runtime *dara.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateTodoTaskRequest
//
// @return UpdateTodoTaskResponse
func (client *Client) UpdateTodoTask(request *UpdateTodoTaskRequest) (_result *UpdateTodoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateTodoTaskHeaders{}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.UpdateTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTodoTaskExecutorStatusWithOptions(tmpReq *UpdateTodoTaskExecutorStatusRequest, tmpHeader *UpdateTodoTaskExecutorStatusHeaders, runtime *dara.RuntimeOptions) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateTodoTaskExecutorStatusRequest
//
// @return UpdateTodoTaskExecutorStatusResponse
func (client *Client) UpdateTodoTaskExecutorStatus(request *UpdateTodoTaskExecutorStatusRequest) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateTodoTaskExecutorStatusHeaders{}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.UpdateTodoTaskExecutorStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateUserAvatarWithOptions(request *UpdateUserAvatarRequest, tmpHeader *UpdateUserAvatarHeaders, runtime *dara.RuntimeOptions) (_result *UpdateUserAvatarResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateUserAvatarResponse
func (client *Client) UpdateUserAvatar(request *UpdateUserAvatarRequest) (_result *UpdateUserAvatarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateUserAvatarHeaders{}
	_result = &UpdateUserAvatarResponse{}
	_body, _err := client.UpdateUserAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateVideoConferenceSettingWithOptions(tmpReq *UpdateVideoConferenceSettingRequest, tmpHeader *UpdateVideoConferenceSettingHeaders, runtime *dara.RuntimeOptions) (_result *UpdateVideoConferenceSettingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateVideoConferenceSettingRequest
//
// @return UpdateVideoConferenceSettingResponse
func (client *Client) UpdateVideoConferenceSetting(request *UpdateVideoConferenceSettingRequest) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateVideoConferenceSettingHeaders{}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.UpdateVideoConferenceSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateWorkspaceDocMembersWithOptions(tmpReq *UpdateWorkspaceDocMembersRequest, tmpHeader *UpdateWorkspaceDocMembersHeaders, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateWorkspaceDocMembersRequest
//
// @return UpdateWorkspaceDocMembersResponse
func (client *Client) UpdateWorkspaceDocMembers(request *UpdateWorkspaceDocMembersRequest) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateWorkspaceDocMembersHeaders{}
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.UpdateWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateWorkspaceMembersWithOptions(tmpReq *UpdateWorkspaceMembersRequest, tmpHeader *UpdateWorkspaceMembersHeaders, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateWorkspaceMembersRequest
//
// @return UpdateWorkspaceMembersResponse
func (client *Client) UpdateWorkspaceMembers(request *UpdateWorkspaceMembersRequest) (_result *UpdateWorkspaceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateWorkspaceMembersHeaders{}
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.UpdateWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UploadMediaWithOptions(tmpReq *UploadMediaRequest, tmpHeader *UploadMediaHeaders, runtime *dara.RuntimeOptions) (_result *UploadMediaResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UploadMediaRequest
//
// @return UploadMediaResponse
func (client *Client) UploadMedia(request *UploadMediaRequest) (_result *UploadMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UploadMediaHeaders{}
	_result = &UploadMediaResponse{}
	_body, _err := client.UploadMediaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) WearOrgHonorWithOptions(tmpReq *WearOrgHonorRequest, tmpHeader *WearOrgHonorHeaders, runtime *dara.RuntimeOptions) (_result *WearOrgHonorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - WearOrgHonorRequest
//
// @return WearOrgHonorResponse
func (client *Client) WearOrgHonor(request *WearOrgHonorRequest) (_result *WearOrgHonorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &WearOrgHonorHeaders{}
	_result = &WearOrgHonorResponse{}
	_body, _err := client.WearOrgHonorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) createRunWithSSE_opYieldFunc(_yield chan *CreateRunResponse, _yieldErr chan error, request *CreateRunRequest, headers *CreateRunHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) invokeAssistantWithSSE_opYieldFunc(_yield chan *InvokeAssistantResponse, _yieldErr chan error, request *InvokeAssistantRequest, headers *InvokeAssistantHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) invokeSkillWithSSE_opYieldFunc(_yield chan *InvokeSkillResponse, _yieldErr chan error, tmpReq *InvokeSkillRequest, tmpHeader *InvokeSkillHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
