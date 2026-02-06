// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 申请：变更目的地/恢复/暂停
//
// @param tmpReq - AddAddressRecoverSuspendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAddressRecoverSuspendResponse
func (client *Client) AddAddressRecoverSuspendWithContext(ctx context.Context, tmpReq *AddAddressRecoverSuspendRequest, runtime *dara.RuntimeOptions) (_result *AddAddressRecoverSuspendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAddressRecoverSuspendShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuditRecord) {
		request.AuditRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuditRecord, dara.String("AuditRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordShrink) {
		query["AuditRecord"] = request.AuditRecordShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RequestType) {
		query["RequestType"] = request.RequestType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAddressRecoverSuspend"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAddressRecoverSuspendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册viber账号，开户
//
// @param tmpReq - AddAuditViberOpenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuditViberOpenResponse
func (client *Client) AddAuditViberOpenWithContext(ctx context.Context, tmpReq *AddAuditViberOpenRequest, runtime *dara.RuntimeOptions) (_result *AddAuditViberOpenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAuditViberOpenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuditRecord) {
		request.AuditRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuditRecord, dara.String("AuditRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordShrink) {
		query["AuditRecord"] = request.AuditRecordShrink
	}

	if !dara.IsNil(request.AuditResult) {
		query["AuditResult"] = request.AuditResult
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAuditViberOpen"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAuditViberOpenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AddChatGroup
//
// @param request - AddChatGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChatGroupResponse
func (client *Client) AddChatGroupWithContext(ctx context.Context, request *AddChatGroupRequest, runtime *dara.RuntimeOptions) (_result *AddChatGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddChatGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddChatGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AddChatGroupInviteLink
//
// @param request - AddChatGroupInviteLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChatGroupInviteLinkResponse
func (client *Client) AddChatGroupInviteLinkWithContext(ctx context.Context, request *AddChatGroupInviteLinkRequest, runtime *dara.RuntimeOptions) (_result *AddChatGroupInviteLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddChatGroupInviteLink"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddChatGroupInviteLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a phone number for a WhatsApp Business account (WABA).
//
// @param request - AddChatappPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChatappPhoneNumberResponse
func (client *Client) AddChatappPhoneNumberWithContext(ctx context.Context, request *AddChatappPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *AddChatappPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cc) {
		query["Cc"] = request.Cc
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PreValidateId) {
		query["PreValidateId"] = request.PreValidateId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VerifiedName) {
		query["VerifiedName"] = request.VerifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddChatappPhoneNumber"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddChatappPhoneNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑联系人-新增联系人
//
// @param tmpReq - AddContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddContactsResponse
func (client *Client) AddContactsWithContext(ctx context.Context, tmpReq *AddContactsRequest, runtime *dara.RuntimeOptions) (_result *AddContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.ContactDetails) {
		query["ContactDetails"] = request.ContactDetails
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Groups) {
		query["Groups"] = request.Groups
	}

	if !dara.IsNil(request.NeedUpdate) {
		query["NeedUpdate"] = request.NeedUpdate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddContacts"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加自定义受众(这个接口需要从镇元直接配置）
//
// @param tmpReq - AddCustomAudienceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomAudienceUserResponse
func (client *Client) AddCustomAudienceUserWithContext(ctx context.Context, tmpReq *AddCustomAudienceUserRequest, runtime *dara.RuntimeOptions) (_result *AddCustomAudienceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddCustomAudienceUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Users) {
		request.UsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Users, dara.String("Users"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.BatchLastFlag) {
		query["BatchLastFlag"] = request.BatchLastFlag
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustomAudienceId) {
		query["CustomAudienceId"] = request.CustomAudienceId
	}

	if !dara.IsNil(request.EstimatedNumTotal) {
		query["EstimatedNumTotal"] = request.EstimatedNumTotal
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UsersShrink) {
		query["Users"] = request.UsersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomAudienceUser"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomAudienceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加群组
//
// @param tmpReq - AddGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupResponse
func (client *Client) AddGroupWithContext(ctx context.Context, tmpReq *AddGroupRequest, runtime *dara.RuntimeOptions) (_result *AddGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.ContactDetails) {
		query["ContactDetails"] = request.ContactDetails
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增营销活动
//
// @param tmpReq - AddMarketingFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMarketingFlowResponse
func (client *Client) AddMarketingFlowWithContext(ctx context.Context, tmpReq *AddMarketingFlowRequest, runtime *dara.RuntimeOptions) (_result *AddMarketingFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddMarketingFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityDesc) {
		query["ActivityDesc"] = request.ActivityDesc
	}

	if !dara.IsNil(request.ActivityName) {
		query["ActivityName"] = request.ActivityName
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExecutionType) {
		query["ExecutionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamFlag) {
		query["ParamFlag"] = request.ParamFlag
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["Params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.RelatedFlowCode) {
		query["RelatedFlowCode"] = request.RelatedFlowCode
	}

	if !dara.IsNil(request.RelatedFlowName) {
		query["RelatedFlowName"] = request.RelatedFlowName
	}

	if !dara.IsNil(request.RelatedGroupId) {
		query["RelatedGroupId"] = request.RelatedGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMarketingFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMarketingFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定DM账号
//
// @param tmpReq - BindDmAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindDmAccountResponse
func (client *Client) BindDmAccountWithContext(ctx context.Context, tmpReq *BindDmAccountRequest, runtime *dara.RuntimeOptions) (_result *BindDmAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BindDmAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtendAttr) {
		request.ExtendAttrShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendAttr, dara.String("ExtendAttr"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountCode) {
		query["AccountCode"] = request.AccountCode
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.ExtendAttrShrink) {
		query["ExtendAttr"] = request.ExtendAttrShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindDmAccount"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindDmAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定ins的page
//
// @param request - BindInstagramPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindInstagramPageResponse
func (client *Client) BindInstagramPageWithContext(ctx context.Context, request *BindInstagramPageRequest, runtime *dara.RuntimeOptions) (_result *BindInstagramPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindInstagramPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindInstagramPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定选择的pageId
//
// @param request - BindMessengerPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindMessengerPageResponse
func (client *Client) BindMessengerPageWithContext(ctx context.Context, request *BindMessengerPageRequest, runtime *dara.RuntimeOptions) (_result *BindMessengerPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindMessengerPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindMessengerPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds the WhatsApp Business account with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappBindWabaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappBindWabaResponse
func (client *Client) ChatappBindWabaWithContext(ctx context.Context, request *ChatappBindWabaRequest, runtime *dara.RuntimeOptions) (_result *ChatappBindWabaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappBindWaba"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappBindWabaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a phone number for migration.
//
// Description:
//
// The space ID of the RAM user within the independent software vendor (ISV) account.
//
// @param request - ChatappMigrationRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappMigrationRegisterResponse
func (client *Client) ChatappMigrationRegisterWithContext(ctx context.Context, request *ChatappMigrationRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappMigrationRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappMigrationRegister"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappMigrationRegisterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies a specified phone number for migration.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappMigrationVerifiedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappMigrationVerifiedResponse
func (client *Client) ChatappMigrationVerifiedWithContext(ctx context.Context, request *ChatappMigrationVerifiedRequest, runtime *dara.RuntimeOptions) (_result *ChatappMigrationVerifiedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappMigrationVerified"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappMigrationVerifiedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deregisters a phone number from a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberDeregisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappPhoneNumberDeregisterResponse
func (client *Client) ChatappPhoneNumberDeregisterWithContext(ctx context.Context, request *ChatappPhoneNumberDeregisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappPhoneNumberDeregisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappPhoneNumberDeregister"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappPhoneNumberDeregisterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappPhoneNumberRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappPhoneNumberRegisterResponse
func (client *Client) ChatappPhoneNumberRegisterWithContext(ctx context.Context, request *ChatappPhoneNumberRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappPhoneNumberRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappPhoneNumberRegister"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappPhoneNumberRegisterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes phone numbers.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappSyncPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappSyncPhoneNumberResponse
func (client *Client) ChatappSyncPhoneNumberWithContext(ctx context.Context, request *ChatappSyncPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ChatappSyncPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappSyncPhoneNumber"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappSyncPhoneNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a phone number with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappVerifyAndRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappVerifyAndRegisterResponse
func (client *Client) ChatappVerifyAndRegisterWithContext(ctx context.Context, request *ChatappVerifyAndRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappVerifyAndRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatappVerifyAndRegister"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappVerifyAndRegisterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制模板
//
// @param request - CopyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyTemplateResponse
func (client *Client) CopyTemplateWithContext(ctx context.Context, request *CopyTemplateRequest, runtime *dara.RuntimeOptions) (_result *CopyTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneTemplateCode) {
		query["SceneTemplateCode"] = request.SceneTemplateCode
	}

	if !dara.IsNil(request.SceneTemplateName) {
		query["SceneTemplateName"] = request.SceneTemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Chatflow
//
// @param tmpReq - CreateChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatFlowResponse
func (client *Client) CreateChatFlowWithContext(ctx context.Context, tmpReq *CreateChatFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowTriggerType) {
		query["FlowTriggerType"] = request.FlowTriggerType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Import and create flow
//
// @param tmpReq - CreateChatFlowByImportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatFlowByImportResponse
func (client *Client) CreateChatFlowByImportWithContext(ctx context.Context, tmpReq *CreateChatFlowByImportRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowByImportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateChatFlowByImportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowViewModel) {
		query["FlowViewModel"] = request.FlowViewModel
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatFlowByImport"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatFlowByImportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create chatFlow log setting
//
// @param request - CreateChatFlowLogSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatFlowLogSettingResponse
func (client *Client) CreateChatFlowLogSettingWithContext(ctx context.Context, request *CreateChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowLogSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatFlowLogSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatFlowLogSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the number.
//
// Description:
//
// The status of the phone number.
//
// @param request - CreateChatappMigrationInitiateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatappMigrationInitiateResponse
func (client *Client) CreateChatappMigrationInitiateWithContext(ctx context.Context, request *CreateChatappMigrationInitiateRequest, runtime *dara.RuntimeOptions) (_result *CreateChatappMigrationInitiateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.MobileNumber) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatappMigrationInitiate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatappMigrationInitiateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The HTTP status code.
//
// \\\\\\\\	- Example: OK. This parameter indicates that the request is successful.
//
// \\\\\\\\	- Other values indicate that the request fails. For more information, see \\\\\\[Error codes]\\\\\\(https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
//
// Description:
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreateChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatappTemplateResponse
func (client *Client) CreateChatappTemplateWithContext(ctx context.Context, tmpReq *CreateChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateChatappTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Example) {
		request.ExampleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Example, dara.String("Example"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowCategoryChange) {
		body["AllowCategoryChange"] = request.AllowCategoryChange
	}

	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.CategoryChangePaused) {
		body["CategoryChangePaused"] = request.CategoryChangePaused
	}

	if !dara.IsNil(request.ComponentsShrink) {
		body["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.ExampleShrink) {
		body["Example"] = request.ExampleShrink
	}

	if !dara.IsNil(request.IsvCode) {
		body["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.MessageSendTtlSeconds) {
		body["MessageSendTtlSeconds"] = request.MessageSendTtlSeconds
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatappTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义受众
//
// @param request - CreateCustomAudienceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomAudienceResponse
func (client *Client) CreateCustomAudienceWithContext(ctx context.Context, request *CreateCustomAudienceRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomAudienceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UploadType) {
		query["UploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomAudience"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomAudienceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithContext(ctx context.Context, tmpReq *CreateFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		query["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.EndpointUri) {
		query["EndpointUri"] = request.EndpointUri
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateFlowVersion
//
// @param tmpReq - CreateFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowVersionResponse
func (client *Client) CreateFlowVersionWithContext(ctx context.Context, tmpReq *CreateFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersionCopyFrom) {
		query["FlowVersionCopyFrom"] = request.FlowVersionCopyFrom
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建实例
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ContactMail) {
		query["ContactMail"] = request.ContactMail
	}

	if !dara.IsNil(request.CountryId) {
		query["CountryId"] = request.CountryId
	}

	if !dara.IsNil(request.FacebookBmId) {
		query["FacebookBmId"] = request.FacebookBmId
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsConfirmAudit) {
		query["IsConfirmAudit"] = request.IsConfirmAudit
	}

	if !dara.IsNil(request.IsvTerms) {
		query["IsvTerms"] = request.IsvTerms
	}

	if !dara.IsNil(request.OfficeAddress) {
		query["OfficeAddress"] = request.OfficeAddress
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预算
//
// @param request - CreateMessageCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageCampaignResponse
func (client *Client) CreateMessageCampaignWithContext(ctx context.Context, request *CreateMessageCampaignRequest, runtime *dara.RuntimeOptions) (_result *CreateMessageCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.Budget) {
		query["Budget"] = request.Budget
	}

	if !dara.IsNil(request.BudgetType) {
		query["BudgetType"] = request.BudgetType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMessageCampaign"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMessageCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 嵌入式授权messenger
//
// @param tmpReq - CreateMessengerPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessengerPageResponse
func (client *Client) CreateMessengerPageWithContext(ctx context.Context, tmpReq *CreateMessengerPageRequest, runtime *dara.RuntimeOptions) (_result *CreateMessengerPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMessengerPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdAccountIds) {
		request.AdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdAccountIds, dara.String("AdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountIdsShrink) {
		query["AdAccountIds"] = request.AdAccountIdsShrink
	}

	if !dara.IsNil(request.AuthenticationCode) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !dara.IsNil(request.BusinessId) {
		query["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMessengerPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMessengerPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quick-response (QR) code that contains a message.
//
// @param request - CreatePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePhoneMessageQrdlResponse
func (client *Client) CreatePhoneMessageQrdlWithContext(ctx context.Context, request *CreatePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *CreatePhoneMessageQrdlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GenerateQrImage) {
		query["GenerateQrImage"] = request.GenerateQrImage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PrefilledMessage) {
		query["PrefilledMessage"] = request.PrefilledMessage
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePhoneMessageQrdl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePhoneMessageQrdlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据嵌入式code获取pageId入库
//
// @param tmpReq - CreateWhatsappConversionApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWhatsappConversionApiResponse
func (client *Client) CreateWhatsappConversionApiWithContext(ctx context.Context, tmpReq *CreateWhatsappConversionApiRequest, runtime *dara.RuntimeOptions) (_result *CreateWhatsappConversionApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWhatsappConversionApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("Permissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PermissionsShrink) {
		query["Permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWhatsappConversionApi"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWhatsappConversionApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Process
//
// @param tmpReq - DeleteChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatFlowResponse
func (client *Client) DeleteChatFlowWithContext(ctx context.Context, tmpReq *DeleteChatFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteChatGroup
//
// @param request - DeleteChatGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatGroupResponse
func (client *Client) DeleteChatGroupWithContext(ctx context.Context, request *DeleteChatGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteChatGroupInviteLink
//
// @param request - DeleteChatGroupInviteLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatGroupInviteLinkResponse
func (client *Client) DeleteChatGroupInviteLinkWithContext(ctx context.Context, request *DeleteChatGroupInviteLinkRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupInviteLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatGroupInviteLink"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatGroupInviteLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteChatGroupParticipants
//
// @param tmpReq - DeleteChatGroupParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatGroupParticipantsResponse
func (client *Client) DeleteChatGroupParticipantsWithContext(ctx context.Context, tmpReq *DeleteChatGroupParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupParticipantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteChatGroupParticipantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.List) {
		request.ListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.List, dara.String("List"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ListShrink) {
		query["List"] = request.ListShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatGroupParticipants"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatGroupParticipantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChatappTemplateResponse
func (client *Client) DeleteChatappTemplateWithContext(ctx context.Context, request *DeleteChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatappTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChatappTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑联系人-删除联系人
//
// @param tmpReq - DeleteContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactsResponse
func (client *Client) DeleteContactsWithContext(ctx context.Context, tmpReq *DeleteContactsRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteContactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.ContactDetails) {
		query["ContactDetails"] = request.ContactDetails
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContacts"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除联系人(选择后删除)
//
// @param request - DeleteContactsByIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactsByIdsResponse
func (client *Client) DeleteContactsByIdsWithContext(ctx context.Context, request *DeleteContactsByIdsRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactsByIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Contacts) {
		query["Contacts"] = request.Contacts
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactsByIds"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactsByIdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Flow. Only Flows in the DRAFT state can be deleted.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithContext(ctx context.Context, request *DeleteFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Flow Version
//
// @param tmpReq - DeleteFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersionWithContext(ctx context.Context, tmpReq *DeleteFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除群组
//
// @param request - DeleteGroupByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupByIdResponse
func (client *Client) DeleteGroupByIdWithContext(ctx context.Context, request *DeleteGroupByIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroupById"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ins的page
//
// @param request - DeleteInstagramPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstagramPageResponse
func (client *Client) DeleteInstagramPageWithContext(ctx context.Context, request *DeleteInstagramPageRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstagramPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstagramPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstagramPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实例
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// 删除营销活动
//
// @param request - DeleteMarketingFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMarketingFlowResponse
func (client *Client) DeleteMarketingFlowWithContext(ctx context.Context, request *DeleteMarketingFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteMarketingFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityCode) {
		query["ActivityCode"] = request.ActivityCode
	}

	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMarketingFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMarketingFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除预算
//
// @param request - DeleteMessageCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageCampaignResponse
func (client *Client) DeleteMessageCampaignWithContext(ctx context.Context, request *DeleteMessageCampaignRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessageCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessageCampaign"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除messenger的page
//
// @param request - DeleteMessengerPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessengerPageResponse
func (client *Client) DeleteMessengerPageWithContext(ctx context.Context, request *DeleteMessengerPageRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessengerPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessengerPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessengerPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quick-response (QR) code that contains a message.
//
// @param request - DeletePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePhoneMessageQrdlResponse
func (client *Client) DeletePhoneMessageQrdlWithContext(ctx context.Context, request *DeletePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *DeletePhoneMessageQrdlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.QrdlCode) {
		query["QrdlCode"] = request.QrdlCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePhoneMessageQrdl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePhoneMessageQrdlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deprecates a Flow.
//
// @param request - DeprecateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprecateFlowResponse
func (client *Client) DeprecateFlowWithContext(ctx context.Context, request *DeprecateFlowRequest, runtime *dara.RuntimeOptions) (_result *DeprecateFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeprecateFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeprecateFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the statistics on the metrics that are related to WhatsApp.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - EnableWhatsappROIMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWhatsappROIMetricResponse
func (client *Client) EnableWhatsappROIMetricWithContext(ctx context.Context, request *EnableWhatsappROIMetricRequest, runtime *dara.RuntimeOptions) (_result *EnableWhatsappROIMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableWhatsappROIMetric"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWhatsappROIMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Bind phone numbers to flow
//
// @param tmpReq - FlowBindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlowBindPhoneResponse
func (client *Client) FlowBindPhoneWithContext(ctx context.Context, tmpReq *FlowBindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowBindPhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FlowBindPhoneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumbers) {
		request.PhoneNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumbers, dara.String("PhoneNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelCode) {
		query["ChannelCode"] = request.ChannelCode
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbersShrink) {
		query["PhoneNumbers"] = request.PhoneNumbersShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlowBindPhone"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlowBindPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Rebind phone number for flow
//
// @param tmpReq - FlowRebindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlowRebindPhoneResponse
func (client *Client) FlowRebindPhoneWithContext(ctx context.Context, tmpReq *FlowRebindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowRebindPhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FlowRebindPhoneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumbers) {
		request.PhoneNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumbers, dara.String("PhoneNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelCode) {
		query["ChannelCode"] = request.ChannelCode
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbersShrink) {
		query["PhoneNumbers"] = request.PhoneNumbersShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlowRebindPhone"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlowRebindPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unbind phone number from flow
//
// @param tmpReq - FlowUnbindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlowUnbindPhoneResponse
func (client *Client) FlowUnbindPhoneWithContext(ctx context.Context, tmpReq *FlowUnbindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowUnbindPhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FlowUnbindPhoneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumbers) {
		request.PhoneNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumbers, dara.String("PhoneNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbersShrink) {
		query["PhoneNumbers"] = request.PhoneNumbersShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlowUnbindPhone"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlowUnbindPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取临时的URL
//
// @param request - GeneratePresignedUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GeneratePresignedUrlResponse
func (client *Client) GeneratePresignedUrlWithContext(ctx context.Context, request *GeneratePresignedUrlRequest, runtime *dara.RuntimeOptions) (_result *GeneratePresignedUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GeneratePresignedUrl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GeneratePresignedUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过类型查询出个人待审核的单子
//
// @param request - GetAuditRequestByTypeUnAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditRequestByTypeUnAuditResponse
func (client *Client) GetAuditRequestByTypeUnAuditWithContext(ctx context.Context, request *GetAuditRequestByTypeUnAuditRequest, runtime *dara.RuntimeOptions) (_result *GetAuditRequestByTypeUnAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RequestType) {
		query["RequestType"] = request.RequestType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuditRequestByTypeUnAudit"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuditRequestByTypeUnAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get ChatFlow Runtime Data
//
// @param tmpReq - GetChatFlowMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatFlowMetricResponse
func (client *Client) GetChatFlowMetricWithContext(ctx context.Context, tmpReq *GetChatFlowMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatFlowMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetChatFlowMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MetricParam) {
		request.MetricParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricParam, dara.String("MetricParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.MetricParamShrink) {
		query["MetricParam"] = request.MetricParamShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatFlowMetric"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatFlowMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query chatFlow template
//
// @param request - GetChatFlowTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatFlowTemplateResponse
func (client *Client) GetChatFlowTemplateWithContext(ctx context.Context, request *GetChatFlowTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetChatFlowTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatFlowTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatFlowTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询ChatApp开通状态
//
// @param request - GetChatappOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappOpenStatusResponse
func (client *Client) GetChatappOpenStatusWithContext(ctx context.Context, request *GetChatappOpenStatusRequest, runtime *dara.RuntimeOptions) (_result *GetChatappOpenStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappOpenStatus"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappOpenStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of messages that are sent by using a phone number by a specific metric.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappPhoneNumberMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappPhoneNumberMetricResponse
func (client *Client) GetChatappPhoneNumberMetricWithContext(ctx context.Context, request *GetChatappPhoneNumberMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatappPhoneNumberMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappPhoneNumberMetric"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappPhoneNumberMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Chatapp号码其它控制
//
// @param request - GetChatappPhoneNumberSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappPhoneNumberSettingResponse
func (client *Client) GetChatappPhoneNumberSettingWithContext(ctx context.Context, request *GetChatappPhoneNumberSettingRequest, runtime *dara.RuntimeOptions) (_result *GetChatappPhoneNumberSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappPhoneNumberSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappPhoneNumberSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this API operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappTemplateDetailResponse
func (client *Client) GetChatappTemplateDetailWithContext(ctx context.Context, request *GetChatappTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetChatappTemplateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappTemplateDetail"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics about a marketing template.
//
// Description:
//
// You can call this operation up to 50 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappTemplateMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappTemplateMetricResponse
func (client *Client) GetChatappTemplateMetricWithContext(ctx context.Context, request *GetChatappTemplateMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatappTemplateMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappTemplateMetric"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappTemplateMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authentication information that is used to upload a file.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappUploadAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappUploadAuthorizationResponse
func (client *Client) GetChatappUploadAuthorizationWithContext(ctx context.Context, request *GetChatappUploadAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *GetChatappUploadAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappUploadAuthorization"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappUploadAuthorizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetChatappVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatappVerifyCodeResponse
func (client *Client) GetChatappVerifyCodeWithContext(ctx context.Context, request *GetChatappVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *GetChatappVerifyCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatappVerifyCode"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatappVerifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCommerceSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommerceSettingResponse
func (client *Client) GetCommerceSettingWithContext(ctx context.Context, request *GetCommerceSettingRequest, runtime *dara.RuntimeOptions) (_result *GetCommerceSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommerceSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommerceSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures welcoming messages, opening remarks, and commands.
//
// Description:
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// This operation will directly obtain data from Facebook, which sets an upper limit on the total number of calls for operations. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetConversationalAutomationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationalAutomationResponse
func (client *Client) GetConversationalAutomationWithContext(ctx context.Context, request *GetConversationalAutomationRequest, runtime *dara.RuntimeOptions) (_result *GetConversationalAutomationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConversationalAutomation"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConversationalAutomationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户来源站点
//
// @param request - GetCustomerSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerSiteResponse
func (client *Client) GetCustomerSiteWithContext(ctx context.Context, request *GetCustomerSiteRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerSite"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerSiteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载excel数据
//
// @param tmpReq - GetDownloadExcelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDownloadExcelListResponse
func (client *Client) GetDownloadExcelListWithContext(ctx context.Context, tmpReq *GetDownloadExcelListRequest, runtime *dara.RuntimeOptions) (_result *GetDownloadExcelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDownloadExcelListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CountryNames) {
		request.CountryNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CountryNames, dara.String("CountryNames"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GroupIds) {
		request.GroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIds, dara.String("GroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.CountryNamesShrink) {
		query["CountryNames"] = request.CountryNamesShrink
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupIdsShrink) {
		query["GroupIds"] = request.GroupIdsShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDownloadExcelList"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDownloadExcelListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ins的page列表
//
// @param request - GetFbInstagramPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFbInstagramPagesResponse
func (client *Client) GetFbInstagramPagesWithContext(ctx context.Context, request *GetFbInstagramPagesRequest, runtime *dara.RuntimeOptions) (_result *GetFbInstagramPagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFbInstagramPages"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFbInstagramPagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取facebook的pageId列表
//
// @param request - GetFbMessengerPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFbMessengerPagesResponse
func (client *Client) GetFbMessengerPagesWithContext(ctx context.Context, request *GetFbMessengerPagesRequest, runtime *dara.RuntimeOptions) (_result *GetFbMessengerPagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFbMessengerPages"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFbMessengerPagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowResponse
func (client *Client) GetFlowWithContext(ctx context.Context, request *GetFlowRequest, runtime *dara.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the JSON content of a Flow.
//
// @param request - GetFlowJSONAssestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowJSONAssestResponse
func (client *Client) GetFlowJSONAssestWithContext(ctx context.Context, request *GetFlowJSONAssestRequest, runtime *dara.RuntimeOptions) (_result *GetFlowJSONAssestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlowJSONAssest"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowJSONAssestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the preview URL of a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetFlowPreviewUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowPreviewUrlResponse
func (client *Client) GetFlowPreviewUrlWithContext(ctx context.Context, request *GetFlowPreviewUrlRequest, runtime *dara.RuntimeOptions) (_result *GetFlowPreviewUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlowPreviewUrl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowPreviewUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群组是否重名
//
// @param tmpReq - GetGroupExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupExistResponse
func (client *Client) GetGroupExistWithContext(ctx context.Context, tmpReq *GetGroupExistRequest, runtime *dara.RuntimeOptions) (_result *GetGroupExistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetGroupExistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroupExist"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupExistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预算指标
//
// @param request - GetMessageCampaignInsightsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageCampaignInsightsResponse
func (client *Client) GetMessageCampaignInsightsWithContext(ctx context.Context, request *GetMessageCampaignInsightsRequest, runtime *dara.RuntimeOptions) (_result *GetMessageCampaignInsightsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageCampaignInsights"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageCampaignInsightsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the verification code for the migration number.
//
// Description:
//
// The single user QPS limit for this interface is 10 times per second. Exceeding the limit may result in restricted API calls, which may affect your business. Please make reasonable calls.
//
// @param request - GetMigrationVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMigrationVerifyCodeResponse
func (client *Client) GetMigrationVerifyCodeWithContext(ctx context.Context, request *GetMigrationVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *GetMigrationVerifyCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMigrationVerifyCode"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMigrationVerifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains permissions based on the authorization code obtained from embedded signup.
//
// @param tmpReq - GetPermissionByCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionByCodeResponse
func (client *Client) GetPermissionByCodeWithContext(ctx context.Context, tmpReq *GetPermissionByCodeRequest, runtime *dara.RuntimeOptions) (_result *GetPermissionByCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPermissionByCodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("Permissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PermissionsShrink) {
		query["Permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPermissionByCode"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPermissionByCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the encryption public key of a phone number.
//
// @param request - GetPhoneEncryptionPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneEncryptionPublicKeyResponse
func (client *Client) GetPhoneEncryptionPublicKeyWithContext(ctx context.Context, request *GetPhoneEncryptionPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *GetPhoneEncryptionPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhoneEncryptionPublicKey"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhoneEncryptionPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification status of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneNumberVerificationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneNumberVerificationStatusResponse
func (client *Client) GetPhoneNumberVerificationStatusWithContext(ctx context.Context, request *GetPhoneNumberVerificationStatusRequest, runtime *dara.RuntimeOptions) (_result *GetPhoneNumberVerificationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhoneNumberVerificationStatus"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhoneNumberVerificationStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载保证函模板的地址
//
// @param request - GetPledgeTemplateAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPledgeTemplateAddressResponse
func (client *Client) GetPledgeTemplateAddressWithContext(ctx context.Context, request *GetPledgeTemplateAddressRequest, runtime *dara.RuntimeOptions) (_result *GetPledgeTemplateAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.IndustryType) {
		query["IndustryType"] = request.IndustryType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPledgeTemplateAddress"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPledgeTemplateAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the ID of a pre-registered phone number used for embedded signup without the need to re-obtain a verification code.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPreValidatePhoneIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPreValidatePhoneIdResponse
func (client *Client) GetPreValidatePhoneIdWithContext(ctx context.Context, request *GetPreValidatePhoneIdRequest, runtime *dara.RuntimeOptions) (_result *GetPreValidatePhoneIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.VerifyCode) {
		body["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPreValidatePhoneId"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPreValidatePhoneIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过requestNo查询申请单
//
// @param request - GetViberByRequestNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViberByRequestNoResponse
func (client *Client) GetViberByRequestNoWithContext(ctx context.Context, request *GetViberByRequestNoRequest, runtime *dara.RuntimeOptions) (_result *GetViberByRequestNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RequestNo) {
		query["RequestNo"] = request.RequestNo
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetViberByRequestNo"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetViberByRequestNoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 可以申请暂停的次数
//
// @param request - GetViberPauseTimesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetViberPauseTimesResponse
func (client *Client) GetViberPauseTimesWithContext(ctx context.Context, request *GetViberPauseTimesRequest, runtime *dara.RuntimeOptions) (_result *GetViberPauseTimesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetViberPauseTimes"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetViberPauseTimesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the product catalogs that are associated with a WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetWhatsappConnectionCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWhatsappConnectionCatalogResponse
func (client *Client) GetWhatsappConnectionCatalogWithContext(ctx context.Context, request *GetWhatsappConnectionCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetWhatsappConnectionCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWhatsappConnectionCatalog"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWhatsappConnectionCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取嵌入式授权page
//
// @param request - GetWhatsappConversionApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWhatsappConversionApiResponse
func (client *Client) GetWhatsappConversionApiWithContext(ctx context.Context, request *GetWhatsappConversionApiRequest, runtime *dara.RuntimeOptions) (_result *GetWhatsappConversionApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWhatsappConversionApi"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWhatsappConversionApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the messaging health status of different types of nodes.
//
// Description:
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// This operation will directly obtain data from Facebook, which sets an upper limit on the total number of calls for operations. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetWhatsappHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWhatsappHealthStatusResponse
func (client *Client) GetWhatsappHealthStatusWithContext(ctx context.Context, request *GetWhatsappHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *GetWhatsappHealthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWhatsappHealthStatus"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWhatsappHealthStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the application ID under the ISV account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - IsvGetAppIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsvGetAppIdResponse
func (client *Client) IsvGetAppIdWithContext(ctx context.Context, request *IsvGetAppIdRequest, runtime *dara.RuntimeOptions) (_result *IsvGetAppIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IntlVersion) {
		query["IntlVersion"] = request.IntlVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Permissions) {
		query["Permissions"] = request.Permissions
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsvGetAppId"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IsvGetAppIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群组列表
//
// @param tmpReq - ListAllGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllGroupsResponse
func (client *Client) ListAllGroupsWithContext(ctx context.Context, tmpReq *ListAllGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListAllGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAllGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllGroups"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询绑定的dm账号
//
// @param request - ListBindDmAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindDmAccountResponse
func (client *Client) ListBindDmAccountWithContext(ctx context.Context, request *ListBindDmAccountRequest, runtime *dara.RuntimeOptions) (_result *ListBindDmAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindDmAccount"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindDmAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Bound List Based on flowCode
//
// Description:
//
// - You can call this interface to query the list of phone numbers or merchant account IDs bound to a process, or you can view the list in the [**Flow Editor**](https://chatapp.console.aliyun.com/ChatFlowBuilder) > **Settings*	- interface.
//
// - Before calling this interface, make sure that the process you created has already been bound to a phone number or merchant account ID.
//
// - If the process you created is not bound to a phone number or merchant account ID, you can manually bind a phone number or merchant account ID in the [**Flow Editor**](https://chatapp.console.aliyun.com/ChatFlowBuilder) > **Settings*	- interface, or bind it through the [FlowBindPhone](https://help.aliyun.com/document_detail/2937190.html) interface.
//
// @param request - ListBindingRelationsForFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindingRelationsForFlowVersionResponse
func (client *Client) ListBindingRelationsForFlowVersionWithContext(ctx context.Context, request *ListBindingRelationsForFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ListBindingRelationsForFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBindingRelationsForFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBindingRelationsForFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Flows
//
// @param tmpReq - ListChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatFlowResponse
func (client *Client) ListChatFlowWithContext(ctx context.Context, tmpReq *ListChatFlowRequest, runtime *dara.RuntimeOptions) (_result *ListChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowTriggerType) {
		query["FlowTriggerType"] = request.FlowTriggerType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ReturnWithOnlineVersion) {
		query["ReturnWithOnlineVersion"] = request.ReturnWithOnlineVersion
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ChatFlow Template List
//
// @param request - ListChatFlowTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatFlowTemplateResponse
func (client *Client) ListChatFlowTemplateWithContext(ctx context.Context, request *ListChatFlowTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListChatFlowTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatFlowTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatFlowTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListChatGroup
//
// @param tmpReq - ListChatGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatGroupResponse
func (client *Client) ListChatGroupWithContext(ctx context.Context, tmpReq *ListChatGroupRequest, runtime *dara.RuntimeOptions) (_result *ListChatGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChatGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupStatus) {
		query["GroupStatus"] = request.GroupStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListChatGroupParticipants
//
// @param tmpReq - ListChatGroupParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatGroupParticipantsResponse
func (client *Client) ListChatGroupParticipantsWithContext(ctx context.Context, tmpReq *ListChatGroupParticipantsRequest, runtime *dara.RuntimeOptions) (_result *ListChatGroupParticipantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChatGroupParticipantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatGroupParticipants"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatGroupParticipantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消息列表
//
// @param tmpReq - ListChatappMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatappMessageResponse
func (client *Client) ListChatappMessageWithContext(ctx context.Context, tmpReq *ListChatappMessageRequest, runtime *dara.RuntimeOptions) (_result *ListChatappMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChatappMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ClientAcceptStatus) {
		query["ClientAcceptStatus"] = request.ClientAcceptStatus
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeStr) {
		query["EndTimeStr"] = request.EndTimeStr
	}

	if !dara.IsNil(request.EventAction) {
		query["EventAction"] = request.EventAction
	}

	if !dara.IsNil(request.GroupMessageId) {
		query["GroupMessageId"] = request.GroupMessageId
	}

	if !dara.IsNil(request.MessageStatus) {
		query["MessageStatus"] = request.MessageStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeStr) {
		query["StartTimeStr"] = request.StartTimeStr
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.UserNumber) {
		query["UserNumber"] = request.UserNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatappMessage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatappMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries message templates.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ListChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatappTemplateResponse
func (client *Client) ListChatappTemplateWithContext(ctx context.Context, tmpReq *ListChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListChatappTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChatappTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义受众组列表
//
// @param tmpReq - ListCustomAudienceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAudienceResponse
func (client *Client) ListCustomAudienceWithContext(ctx context.Context, tmpReq *ListCustomAudienceRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAudienceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCustomAudienceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustomAudienceId) {
		query["CustomAudienceId"] = request.CustomAudienceId
	}

	if !dara.IsNil(request.CustomAudienceName) {
		query["CustomAudienceName"] = request.CustomAudienceName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAudience"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAudienceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账号下的Dm账号
//
// @param request - ListDmAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDmAccountResponse
func (client *Client) ListDmAccountWithContext(ctx context.Context, request *ListDmAccountRequest, runtime *dara.RuntimeOptions) (_result *ListDmAccountResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SendType) {
		query["SendType"] = request.SendType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDmAccount"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDmAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询DM的tag
//
// @param request - ListDmTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDmTagResponse
func (client *Client) ListDmTagWithContext(ctx context.Context, request *ListDmTagRequest, runtime *dara.RuntimeOptions) (_result *ListDmTagResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDmTag"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDmTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询facebook帖子列表
//
// @param request - ListFacebookPostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFacebookPostsResponse
func (client *Client) ListFacebookPostsWithContext(ctx context.Context, request *ListFacebookPostsRequest, runtime *dara.RuntimeOptions) (_result *ListFacebookPostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFacebookPosts"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFacebookPostsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Flows.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ListFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowResponse
func (client *Client) ListFlowWithContext(ctx context.Context, tmpReq *ListFlowRequest, runtime *dara.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListFlowNodeGroup
//
// @param request - ListFlowNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowNodeGroupResponse
func (client *Client) ListFlowNodeGroupWithContext(ctx context.Context, request *ListFlowNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *ListFlowNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowNodeGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowNodeGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListFlowNodePrototypeV2
//
// @param request - ListFlowNodePrototypeV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowNodePrototypeV2Response
func (client *Client) ListFlowNodePrototypeV2WithContext(ctx context.Context, request *ListFlowNodePrototypeV2Request, runtime *dara.RuntimeOptions) (_result *ListFlowNodePrototypeV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.GroupCode) {
		query["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowNodePrototypeV2"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowNodePrototypeV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Flow Versions
//
// @param tmpReq - ListFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowVersionResponse
func (client *Client) ListFlowVersionWithContext(ctx context.Context, tmpReq *ListFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ListFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ins的page
//
// @param request - ListInstagramPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstagramPageResponse
func (client *Client) ListInstagramPageWithContext(ctx context.Context, request *ListInstagramPageRequest, runtime *dara.RuntimeOptions) (_result *ListInstagramPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstagramPage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstagramPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询instagram帖子列表
//
// @param request - ListInstagramPostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstagramPostsResponse
func (client *Client) ListInstagramPostsWithContext(ctx context.Context, request *ListInstagramPostsRequest, runtime *dara.RuntimeOptions) (_result *ListInstagramPostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstagramPosts"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstagramPostsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithContext(ctx context.Context, request *ListInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.FilterStr) {
		query["FilterStr"] = request.FilterStr
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SubmitTime) {
		query["SubmitTime"] = request.SubmitTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstance"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询活动列表
//
// @param tmpReq - ListMarketingFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMarketingFlowResponse
func (client *Client) ListMarketingFlowWithContext(ctx context.Context, tmpReq *ListMarketingFlowRequest, runtime *dara.RuntimeOptions) (_result *ListMarketingFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMarketingFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityCode) {
		query["ActivityCode"] = request.ActivityCode
	}

	if !dara.IsNil(request.ActivityName) {
		query["ActivityName"] = request.ActivityName
	}

	if !dara.IsNil(request.ActivityStatus) {
		query["ActivityStatus"] = request.ActivityStatus
	}

	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RelatedFlowCode) {
		query["RelatedFlowCode"] = request.RelatedFlowCode
	}

	if !dara.IsNil(request.RelatedGroupId) {
		query["RelatedGroupId"] = request.RelatedGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMarketingFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMarketingFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预算列表
//
// @param tmpReq - ListMessageCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageCampaignResponse
func (client *Client) ListMessageCampaignWithContext(ctx context.Context, tmpReq *ListMessageCampaignRequest, runtime *dara.RuntimeOptions) (_result *ListMessageCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMessageCampaignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.CampaignName) {
		query["CampaignName"] = request.CampaignName
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageCampaign"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订阅token
//
// @param request - ListMessengerSubscriptionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessengerSubscriptionTokenResponse
func (client *Client) ListMessengerSubscriptionTokenWithContext(ctx context.Context, request *ListMessengerSubscriptionTokenRequest, runtime *dara.RuntimeOptions) (_result *ListMessengerSubscriptionTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustomAudienceId) {
		query["CustomAudienceId"] = request.CustomAudienceId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.PageKey) {
		query["PageKey"] = request.PageKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessengerSubscriptionToken"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessengerSubscriptionTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Page绑定的广告账户列表
//
// @param request - ListPageAdAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPageAdAccountResponse
func (client *Client) ListPageAdAccountWithContext(ctx context.Context, request *ListPageAdAccountRequest, runtime *dara.RuntimeOptions) (_result *ListPageAdAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPageAdAccount"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPageAdAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a list of quick-response (QR) codes that contain messages.
//
// @param request - ListPhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPhoneMessageQrdlResponse
func (client *Client) ListPhoneMessageQrdlWithContext(ctx context.Context, request *ListPhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *ListPhoneMessageQrdlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPhoneMessageQrdl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPhoneMessageQrdlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries products in a product catalog.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductResponse
func (client *Client) ListProductWithContext(ctx context.Context, request *ListProductRequest, runtime *dara.RuntimeOptions) (_result *ListProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.After) {
		query["After"] = request.After
	}

	if !dara.IsNil(request.Before) {
		query["Before"] = request.Before
	}

	if !dara.IsNil(request.CatalogId) {
		query["CatalogId"] = request.CatalogId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProduct"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the product catalogs on the Business Manager platform of Meta.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListProductCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductCatalogResponse
func (client *Client) ListProductCatalogWithContext(ctx context.Context, request *ListProductCatalogRequest, runtime *dara.RuntimeOptions) (_result *ListProductCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.After) {
		query["After"] = request.After
	}

	if !dara.IsNil(request.Before) {
		query["Before"] = request.Before
	}

	if !dara.IsNil(request.BusinessId) {
		query["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductCatalog"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示viber申请单服务号卡片
//
// @param request - ListViberServiceMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListViberServiceMessageResponse
func (client *Client) ListViberServiceMessageWithContext(ctx context.Context, request *ListViberServiceMessageRequest, runtime *dara.RuntimeOptions) (_result *ListViberServiceMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListViberServiceMessage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListViberServiceMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The code of the message template.
//
// Description:
//
// The name of the message template.
//
// @param tmpReq - ModifyChatappTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyChatappTemplateResponse
func (client *Client) ModifyChatappTemplateWithContext(ctx context.Context, tmpReq *ModifyChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyChatappTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyChatappTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Example) {
		request.ExampleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Example, dara.String("Example"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["Category"] = request.Category
	}

	if !dara.IsNil(request.CategoryChangePaused) {
		body["CategoryChangePaused"] = request.CategoryChangePaused
	}

	if !dara.IsNil(request.ComponentsShrink) {
		body["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.ExampleShrink) {
		body["Example"] = request.ExampleShrink
	}

	if !dara.IsNil(request.IsvCode) {
		body["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.MessageSendTtlSeconds) {
		body["MessageSendTtlSeconds"] = request.MessageSendTtlSeconds
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyChatappTemplate"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyChatappTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改模板上的一些属性
//
// @param request - ModifyChatappTemplatePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyChatappTemplatePropertiesResponse
func (client *Client) ModifyChatappTemplatePropertiesWithContext(ctx context.Context, request *ModifyChatappTemplatePropertiesRequest, runtime *dara.RuntimeOptions) (_result *ModifyChatappTemplatePropertiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSend) {
		query["AllowSend"] = request.AllowSend
	}

	if !dara.IsNil(request.CategoryChangePaused) {
		query["CategoryChangePaused"] = request.CategoryChangePaused
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyChatappTemplateProperties"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyChatappTemplatePropertiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - ModifyFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowResponse
func (client *Client) ModifyFlowWithContext(ctx context.Context, tmpReq *ModifyFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		query["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.EndpointUri) {
		query["EndpointUri"] = request.EndpointUri
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// # ModifyPhoneBusinessProfile
//
// @param tmpReq - ModifyPhoneBusinessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPhoneBusinessProfileResponse
func (client *Client) ModifyPhoneBusinessProfileWithContext(ctx context.Context, tmpReq *ModifyPhoneBusinessProfileRequest, runtime *dara.RuntimeOptions) (_result *ModifyPhoneBusinessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPhoneBusinessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Websites) {
		request.WebsitesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Websites, dara.String("Websites"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.About) {
		query["About"] = request.About
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ProfilePictureUrl) {
		query["ProfilePictureUrl"] = request.ProfilePictureUrl
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Vertical) {
		query["Vertical"] = request.Vertical
	}

	if !dara.IsNil(request.WebsitesShrink) {
		query["Websites"] = request.WebsitesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPhoneBusinessProfile"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPhoneBusinessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 联系人变更群组
//
// @param tmpReq - MoveContactToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveContactToGroupResponse
func (client *Client) MoveContactToGroupWithContext(ctx context.Context, tmpReq *MoveContactToGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveContactToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MoveContactToGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.Contacts) {
		query["Contacts"] = request.Contacts
	}

	if !dara.IsNil(request.LinkExistGroups) {
		query["LinkExistGroups"] = request.LinkExistGroups
	}

	if !dara.IsNil(request.LinkNewGroups) {
		query["LinkNewGroups"] = request.LinkNewGroups
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveContactToGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveContactToGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Offline Flow Version
//
// @param tmpReq - OfflineFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineFlowVersionResponse
func (client *Client) OfflineFlowVersionWithContext(ctx context.Context, tmpReq *OfflineFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *OfflineFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflineFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Online Flow Version
//
// @param tmpReq - OnlineFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineFlowVersionResponse
func (client *Client) OnlineFlowVersionWithContext(ctx context.Context, tmpReq *OnlineFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *OnlineFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OnlineFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通Chatapp服务
//
// @param request - OpenChatappServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenChatappServiceResponse
func (client *Client) OpenChatappServiceWithContext(ctx context.Context, request *OpenChatappServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenChatappServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenChatappService"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenChatappServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停服务
//
// @param request - PauseMarketingFLowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseMarketingFLowResponse
func (client *Client) PauseMarketingFLowWithContext(ctx context.Context, request *PauseMarketingFLowRequest, runtime *dara.RuntimeOptions) (_result *PauseMarketingFLowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityCode) {
		query["ActivityCode"] = request.ActivityCode
	}

	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseMarketingFLow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseMarketingFLowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a Flow.
//
// Description:
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PublishFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFlowResponse
func (client *Client) PublishFlowWithContext(ctx context.Context, request *PublishFlowRequest, runtime *dara.RuntimeOptions) (_result *PublishFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the WhatsApp Business account you associate with ChatApp.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappBindWabaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatappBindWabaResponse
func (client *Client) QueryChatappBindWabaWithContext(ctx context.Context, request *QueryChatappBindWabaRequest, runtime *dara.RuntimeOptions) (_result *QueryChatappBindWabaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryChatappBindWaba"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChatappBindWabaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries phone numbers that receive messages and statuses of these numbers under a specified user.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryChatappPhoneNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatappPhoneNumbersResponse
func (client *Client) QueryChatappPhoneNumbersWithContext(ctx context.Context, request *QueryChatappPhoneNumbersRequest, runtime *dara.RuntimeOptions) (_result *QueryChatappPhoneNumbersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryChatappPhoneNumbers"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChatappPhoneNumbersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例
//
// @param request - QueryInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceResponse
func (client *Client) QueryInstanceWithContext(ctx context.Context, request *QueryInstanceRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstance"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询营销消息是否生效
//
// @param request - QueryMMLActiveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMMLActiveResponse
func (client *Client) QueryMMLActiveWithContext(ctx context.Context, request *QueryMMLActiveRequest, runtime *dara.RuntimeOptions) (_result *QueryMMLActiveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMMLActive"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMMLActiveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business information of the account to which a specified phone number is bound.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryPhoneBusinessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneBusinessProfileResponse
func (client *Client) QueryPhoneBusinessProfileWithContext(ctx context.Context, request *QueryPhoneBusinessProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneBusinessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPhoneBusinessProfile"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPhoneBusinessProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the business information about the WhatsApp Business account (WABA).
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryWabaBusinessInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWabaBusinessInfoResponse
func (client *Client) QueryWabaBusinessInfoWithContext(ctx context.Context, request *QueryWabaBusinessInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryWabaBusinessInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWabaBusinessInfo"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWabaBusinessInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve Flow
//
// @param tmpReq - ReadChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadChatFlowResponse
func (client *Client) ReadChatFlowWithContext(ctx context.Context, tmpReq *ReadChatFlowRequest, runtime *dara.RuntimeOptions) (_result *ReadChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View chatFlow log settings
//
// @param request - ReadChatFlowLogSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadChatFlowLogSettingResponse
func (client *Client) ReadChatFlowLogSettingWithContext(ctx context.Context, request *ReadChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *ReadChatFlowLogSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadChatFlowLogSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadChatFlowLogSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Flow Version
//
// @param tmpReq - ReadFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadFlowVersionResponse
func (client *Client) ReadFlowVersionWithContext(ctx context.Context, tmpReq *ReadFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ReadFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前群组移除单个联系人
//
// @param request - RemoveContactByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveContactByIdResponse
func (client *Client) RemoveContactByIdWithContext(ctx context.Context, request *RemoveContactByIdRequest, runtime *dara.RuntimeOptions) (_result *RemoveContactByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveContactById"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveContactByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请求Whatsapp Conversion api
//
// @param tmpReq - RequestWhatsappConversionApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequestWhatsappConversionApiResponse
func (client *Client) RequestWhatsappConversionApiWithContext(ctx context.Context, tmpReq *RequestWhatsappConversionApiRequest, runtime *dara.RuntimeOptions) (_result *RequestWhatsappConversionApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RequestWhatsappConversionApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequestData) {
		request.RequestDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestData, dara.String("RequestData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.RequestDataShrink) {
		query["RequestData"] = request.RequestDataShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RequestWhatsappConversionApi"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RequestWhatsappConversionApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to multiple phone numbers by using ChatAPP at a time.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// You can send messages to up to 1,000 phone numbers in a single request.
//
// @param tmpReq - SendChatappMassMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatappMassMessageResponse
func (client *Client) SendChatappMassMessageWithContext(ctx context.Context, tmpReq *SendChatappMassMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatappMassMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendChatappMassMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SenderList) {
		request.SenderListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SenderList, dara.String("SenderList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.FallBackContent) {
		query["FallBackContent"] = request.FallBackContent
	}

	if !dara.IsNil(request.FallBackDuration) {
		query["FallBackDuration"] = request.FallBackDuration
	}

	if !dara.IsNil(request.FallBackId) {
		query["FallBackId"] = request.FallBackId
	}

	if !dara.IsNil(request.FallBackRule) {
		query["FallBackRule"] = request.FallBackRule
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SenderListShrink) {
		query["SenderList"] = request.SenderListShrink
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatappMassMessage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendChatappMassMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends messages by using ChatAPP.
//
// Description:
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - SendChatappMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatappMessageResponse
func (client *Client) SendChatappMessageWithContext(ctx context.Context, tmpReq *SendChatappMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatappMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendChatappMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FlowAction) {
		request.FlowActionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlowAction, dara.String("FlowAction"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProductAction) {
		request.ProductActionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductAction, dara.String("ProductAction"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateParams) {
		request.TemplateParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateParams, dara.String("TemplateParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContextMessageId) {
		query["ContextMessageId"] = request.ContextMessageId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		query["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.FallBackContent) {
		query["FallBackContent"] = request.FallBackContent
	}

	if !dara.IsNil(request.FallBackDuration) {
		query["FallBackDuration"] = request.FallBackDuration
	}

	if !dara.IsNil(request.FallBackId) {
		query["FallBackId"] = request.FallBackId
	}

	if !dara.IsNil(request.FallBackRule) {
		query["FallBackRule"] = request.FallBackRule
	}

	if !dara.IsNil(request.FlowActionShrink) {
		query["FlowAction"] = request.FlowActionShrink
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.IsvCode) {
		query["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MessageCampaignId) {
		query["MessageCampaignId"] = request.MessageCampaignId
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.ProductActionShrink) {
		query["ProductAction"] = request.ProductActionShrink
	}

	if !dara.IsNil(request.RecipientType) {
		query["RecipientType"] = request.RecipientType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateParamsShrink) {
		query["TemplateParams"] = request.TemplateParamsShrink
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	if !dara.IsNil(request.TrackingData) {
		query["TrackingData"] = request.TrackingData
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatappMessage"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步flow
//
// @param request - SyncFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncFlowResponse
func (client *Client) SyncFlowWithContext(ctx context.Context, request *SyncFlowRequest, runtime *dara.RuntimeOptions) (_result *SyncFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步查询预算
//
// @param request - SyncMessageCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncMessageCampaignResponse
func (client *Client) SyncMessageCampaignWithContext(ctx context.Context, request *SyncMessageCampaignRequest, runtime *dara.RuntimeOptions) (_result *SyncMessageCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdAccountId) {
		query["AdAccountId"] = request.AdAccountId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncMessageCampaign"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncMessageCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Messenger订阅token
//
// @param request - SyncMessengerSubscriptionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncMessengerSubscriptionTokenResponse
func (client *Client) SyncMessengerSubscriptionTokenWithContext(ctx context.Context, request *SyncMessengerSubscriptionTokenRequest, runtime *dara.RuntimeOptions) (_result *SyncMessengerSubscriptionTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustomAudienceId) {
		query["CustomAudienceId"] = request.CustomAudienceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageId) {
		query["PageId"] = request.PageId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncMessengerSubscriptionToken"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncMessengerSubscriptionTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Trigger an Online ChatFlow
//
// Description:
//
// After triggering an online flow, if your flow contains components that incur costs for cloud products, such as message sending or function calls, please ensure you fully understand the billing methods and prices of the related products before using this interface.
//
// @param tmpReq - TriggerChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerChatFlowResponse
func (client *Client) TriggerChatFlowWithContext(ctx context.Context, tmpReq *TriggerChatFlowRequest, runtime *dara.RuntimeOptions) (_result *TriggerChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TriggerChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Data) {
		request.DataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Data, dara.String("Data"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClaimTimeMillis) {
		query["ClaimTimeMillis"] = request.ClaimTimeMillis
	}

	if !dara.IsNil(request.DataShrink) {
		query["Data"] = request.DataShrink
	}

	if !dara.IsNil(request.DiscardTimeMillis) {
		query["DiscardTimeMillis"] = request.DiscardTimeMillis
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑邮件账号
//
// @param request - UnbindDmAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDmAccountResponse
func (client *Client) UnbindDmAccountWithContext(ctx context.Context, request *UnbindDmAccountRequest, runtime *dara.RuntimeOptions) (_result *UnbindDmAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDmAccount"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDmAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the callback URL of an account.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateAccountWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountWebhookResponse
func (client *Client) UpdateAccountWebhookWithContext(ctx context.Context, request *UpdateAccountWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.HttpFlag) {
		query["HttpFlag"] = request.HttpFlag
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueueFlag) {
		query["QueueFlag"] = request.QueueFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StatusCallbackUrl) {
		query["StatusCallbackUrl"] = request.StatusCallbackUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccountWebhook"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改viber申请单
//
// @param tmpReq - UpdateAuditRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuditRequestResponse
func (client *Client) UpdateAuditRequestWithContext(ctx context.Context, tmpReq *UpdateAuditRequestRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuditRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAuditRequestShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuditRecord) {
		request.AuditRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuditRecord, dara.String("AuditRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordShrink) {
		query["AuditRecord"] = request.AuditRecordShrink
	}

	if !dara.IsNil(request.AuditResult) {
		query["AuditResult"] = request.AuditResult
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RequestNo) {
		query["RequestNo"] = request.RequestNo
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuditRequest"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuditRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Process
//
// @param tmpReq - UpdateChatFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChatFlowResponse
func (client *Client) UpdateChatFlowWithContext(ctx context.Context, tmpReq *UpdateChatFlowRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateChatFlowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChatFlow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChatFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify chatFlow log settings
//
// @param request - UpdateChatFlowLogSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChatFlowLogSettingResponse
func (client *Client) UpdateChatFlowLogSettingWithContext(ctx context.Context, request *UpdateChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatFlowLogSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChatFlowLogSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChatFlowLogSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateChatGroup
//
// @param request - UpdateChatGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChatGroupResponse
func (client *Client) UpdateChatGroupWithContext(ctx context.Context, request *UpdateChatGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProfilePictureFile) {
		query["ProfilePictureFile"] = request.ProfilePictureFile
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChatGroup"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChatGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the business settings of a phone number.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UpdateCommerceSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommerceSettingResponse
func (client *Client) UpdateCommerceSettingWithContext(ctx context.Context, request *UpdateCommerceSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateCommerceSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CartEnable) {
		query["CartEnable"] = request.CartEnable
	}

	if !dara.IsNil(request.CatalogVisible) {
		query["CatalogVisible"] = request.CatalogVisible
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCommerceSetting"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCommerceSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改联系人
//
// @param tmpReq - UpdateContactByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContactByIdResponse
func (client *Client) UpdateContactByIdWithContext(ctx context.Context, tmpReq *UpdateContactByIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateContactByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateContactByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.ContactDetails) {
		query["ContactDetails"] = request.ContactDetails
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContactById"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContactByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies welcoming messages, opening remarks, and commands for a phone number.
//
// Description:
//
// ### [](#qps-)QPS limit
//
// You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// This operation will directly obtain data from Facebook, which sets an upper limit on the total number of calls for operations. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - UpdateConversationalAutomationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConversationalAutomationResponse
func (client *Client) UpdateConversationalAutomationWithContext(ctx context.Context, tmpReq *UpdateConversationalAutomationRequest, runtime *dara.RuntimeOptions) (_result *UpdateConversationalAutomationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateConversationalAutomationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Commands) {
		request.CommandsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commands, dara.String("Commands"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Prompts) {
		request.PromptsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Prompts, dara.String("Prompts"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandsShrink) {
		query["Commands"] = request.CommandsShrink
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.EnableWelcomeMessage) {
		query["EnableWelcomeMessage"] = request.EnableWelcomeMessage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PromptsShrink) {
		query["Prompts"] = request.PromptsShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConversationalAutomation"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConversationalAutomationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Flow by using JSON content.
//
// @param request - UpdateFlowJSONAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowJSONAssetResponse
func (client *Client) UpdateFlowJSONAssetWithContext(ctx context.Context, request *UpdateFlowJSONAssetRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowJSONAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowJSONAsset"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowJSONAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update flow version, used for updating the flow DSL on the canvas
//
// @param tmpReq - UpdateFlowVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowVersionResponse
func (client *Client) UpdateFlowVersionWithContext(ctx context.Context, tmpReq *UpdateFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFlowVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizExtend) {
		request.BizExtendShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizExtend, dara.String("BizExtend"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCode) {
		query["BizCode"] = request.BizCode
	}

	if !dara.IsNil(request.BizExtendShrink) {
		query["BizExtend"] = request.BizExtendShrink
	}

	if !dara.IsNil(request.FlowCode) {
		query["FlowCode"] = request.FlowCode
	}

	if !dara.IsNil(request.FlowVersion) {
		query["FlowVersion"] = request.FlowVersion
	}

	if !dara.IsNil(request.FlowViewModel) {
		query["FlowViewModel"] = request.FlowViewModel
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowVersion"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群组改名
//
// @param request - UpdateGroupNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupNameResponse
func (client *Client) UpdateGroupNameWithContext(ctx context.Context, request *UpdateGroupNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroupName"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactMail) {
		query["ContactMail"] = request.ContactMail
	}

	if !dara.IsNil(request.CountryId) {
		query["CountryId"] = request.CountryId
	}

	if !dara.IsNil(request.FacebookBmId) {
		query["FacebookBmId"] = request.FacebookBmId
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsConfirmAudit) {
		query["IsConfirmAudit"] = request.IsConfirmAudit
	}

	if !dara.IsNil(request.IsvTerms) {
		query["IsvTerms"] = request.IsvTerms
	}

	if !dara.IsNil(request.OfficeAddress) {
		query["OfficeAddress"] = request.OfficeAddress
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// 修改营销活动
//
// @param tmpReq - UpdateMarketingFLowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMarketingFLowResponse
func (client *Client) UpdateMarketingFLowWithContext(ctx context.Context, tmpReq *UpdateMarketingFLowRequest, runtime *dara.RuntimeOptions) (_result *UpdateMarketingFLowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMarketingFLowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityCode) {
		query["ActivityCode"] = request.ActivityCode
	}

	if !dara.IsNil(request.ActivityDesc) {
		query["ActivityDesc"] = request.ActivityDesc
	}

	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.ActivityName) {
		query["ActivityName"] = request.ActivityName
	}

	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExecutionType) {
		query["ExecutionType"] = request.ExecutionType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamFlag) {
		query["ParamFlag"] = request.ParamFlag
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["Params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.RelatedFlowCode) {
		query["RelatedFlowCode"] = request.RelatedFlowCode
	}

	if !dara.IsNil(request.RelatedFlowName) {
		query["RelatedFlowName"] = request.RelatedFlowName
	}

	if !dara.IsNil(request.RelatedGroupId) {
		query["RelatedGroupId"] = request.RelatedGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMarketingFLow"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMarketingFLowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the encryption public key of a phone number.
//
// @param request - UpdatePhoneEncryptionPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneEncryptionPublicKeyResponse
func (client *Client) UpdatePhoneEncryptionPublicKeyWithContext(ctx context.Context, request *UpdatePhoneEncryptionPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneEncryptionPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.EncryptionPublicKey) {
		query["EncryptionPublicKey"] = request.EncryptionPublicKey
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePhoneEncryptionPublicKey"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePhoneEncryptionPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a quick-response (QR) code that contains a message.
//
// @param request - UpdatePhoneMessageQrdlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneMessageQrdlResponse
func (client *Client) UpdatePhoneMessageQrdlWithContext(ctx context.Context, request *UpdatePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneMessageQrdlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.GenerateQrImage) {
		query["GenerateQrImage"] = request.GenerateQrImage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PrefilledMessage) {
		query["PrefilledMessage"] = request.PrefilledMessage
	}

	if !dara.IsNil(request.QrdlCode) {
		query["QrdlCode"] = request.QrdlCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePhoneMessageQrdl"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePhoneMessageQrdlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The HTTP status code returned.
//
// \\	- A value of OK indicates that the call is successful.
//
// \\	- Other values indicate that the call fails. For more information, see [Error codes]\\(~~196974~~).
//
// Description:
//
// The error message returned.
//
// @param request - UpdatePhoneWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePhoneWebhookResponse
func (client *Client) UpdatePhoneWebhookWithContext(ctx context.Context, request *UpdatePhoneWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.HttpFlag) {
		query["HttpFlag"] = request.HttpFlag
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.QueueFlag) {
		query["QueueFlag"] = request.QueueFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StatusCallbackUrl) {
		query["StatusCallbackUrl"] = request.StatusCallbackUrl
	}

	if !dara.IsNil(request.UpCallbackUrl) {
		query["UpCallbackUrl"] = request.UpCallbackUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePhoneWebhook"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePhoneWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新waba的mml状态
//
// @param request - UpdateWabaMmlStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWabaMmlStatusResponse
func (client *Client) UpdateWabaMmlStatusWithContext(ctx context.Context, request *UpdateWabaMmlStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateWabaMmlStatusResponse, _err error) {
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

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WabaId) {
		query["WabaId"] = request.WabaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWabaMmlStatus"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWabaMmlStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Whatsapp 语音电话
//
// @param tmpReq - WhatsappCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WhatsappCallResponse
func (client *Client) WhatsappCallWithContext(ctx context.Context, tmpReq *WhatsappCallRequest, runtime *dara.RuntimeOptions) (_result *WhatsappCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &WhatsappCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Session) {
		request.SessionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Session, dara.String("Session"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessNumber) {
		query["BusinessNumber"] = request.BusinessNumber
	}

	if !dara.IsNil(request.CallAction) {
		query["CallAction"] = request.CallAction
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SessionShrink) {
		query["Session"] = request.SessionShrink
	}

	if !dara.IsNil(request.UserNumber) {
		query["UserNumber"] = request.UserNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WhatsappCall"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WhatsappCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
