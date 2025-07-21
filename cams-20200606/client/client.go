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
	client.Endpoint, _err = client.GetEndpoint(dara.String("cams"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # AddChatGroup
//
// @param request - AddChatGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddChatGroupResponse
func (client *Client) AddChatGroupWithOptions(request *AddChatGroupRequest, runtime *dara.RuntimeOptions) (_result *AddChatGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddChatGroupResponse
func (client *Client) AddChatGroup(request *AddChatGroupRequest) (_result *AddChatGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddChatGroupResponse{}
	_body, _err := client.AddChatGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddChatGroupInviteLinkWithOptions(request *AddChatGroupInviteLinkRequest, runtime *dara.RuntimeOptions) (_result *AddChatGroupInviteLinkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddChatGroupInviteLinkResponse
func (client *Client) AddChatGroupInviteLink(request *AddChatGroupInviteLinkRequest) (_result *AddChatGroupInviteLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddChatGroupInviteLinkResponse{}
	_body, _err := client.AddChatGroupInviteLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddChatappPhoneNumberWithOptions(request *AddChatappPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *AddChatappPhoneNumberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddChatappPhoneNumberResponse
func (client *Client) AddChatappPhoneNumber(request *AddChatappPhoneNumberRequest) (_result *AddChatappPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddChatappPhoneNumberResponse{}
	_body, _err := client.AddChatappPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappBindWabaWithOptions(request *ChatappBindWabaRequest, runtime *dara.RuntimeOptions) (_result *ChatappBindWabaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappBindWabaResponse
func (client *Client) ChatappBindWaba(request *ChatappBindWabaRequest) (_result *ChatappBindWabaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappBindWabaResponse{}
	_body, _err := client.ChatappBindWabaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries WhatsApp Business account (WABA) information after embedded signup. You do not need to call this API operation if you use Version 2 of WhatsApp embedded signup.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappEmbedSignUpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatappEmbedSignUpResponse
func (client *Client) ChatappEmbedSignUpWithOptions(request *ChatappEmbedSignUpRequest, runtime *dara.RuntimeOptions) (_result *ChatappEmbedSignUpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputToken) {
		query["InputToken"] = request.InputToken
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
		Action:      dara.String("ChatappEmbedSignUp"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatappEmbedSignUpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries WhatsApp Business account (WABA) information after embedded signup. You do not need to call this API operation if you use Version 2 of WhatsApp embedded signup.
//
// Description:
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ChatappEmbedSignUpRequest
//
// @return ChatappEmbedSignUpResponse
func (client *Client) ChatappEmbedSignUp(request *ChatappEmbedSignUpRequest) (_result *ChatappEmbedSignUpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappEmbedSignUpResponse{}
	_body, _err := client.ChatappEmbedSignUpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappMigrationRegisterWithOptions(request *ChatappMigrationRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappMigrationRegisterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappMigrationRegisterResponse
func (client *Client) ChatappMigrationRegister(request *ChatappMigrationRegisterRequest) (_result *ChatappMigrationRegisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappMigrationRegisterResponse{}
	_body, _err := client.ChatappMigrationRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappMigrationVerifiedWithOptions(request *ChatappMigrationVerifiedRequest, runtime *dara.RuntimeOptions) (_result *ChatappMigrationVerifiedResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappMigrationVerifiedResponse
func (client *Client) ChatappMigrationVerified(request *ChatappMigrationVerifiedRequest) (_result *ChatappMigrationVerifiedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappMigrationVerifiedResponse{}
	_body, _err := client.ChatappMigrationVerifiedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappPhoneNumberDeregisterWithOptions(request *ChatappPhoneNumberDeregisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappPhoneNumberDeregisterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappPhoneNumberDeregisterResponse
func (client *Client) ChatappPhoneNumberDeregister(request *ChatappPhoneNumberDeregisterRequest) (_result *ChatappPhoneNumberDeregisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappPhoneNumberDeregisterResponse{}
	_body, _err := client.ChatappPhoneNumberDeregisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappPhoneNumberRegisterWithOptions(request *ChatappPhoneNumberRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappPhoneNumberRegisterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappPhoneNumberRegisterResponse
func (client *Client) ChatappPhoneNumberRegister(request *ChatappPhoneNumberRegisterRequest) (_result *ChatappPhoneNumberRegisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappPhoneNumberRegisterResponse{}
	_body, _err := client.ChatappPhoneNumberRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappSyncPhoneNumberWithOptions(request *ChatappSyncPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ChatappSyncPhoneNumberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappSyncPhoneNumberResponse
func (client *Client) ChatappSyncPhoneNumber(request *ChatappSyncPhoneNumberRequest) (_result *ChatappSyncPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappSyncPhoneNumberResponse{}
	_body, _err := client.ChatappSyncPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatappVerifyAndRegisterWithOptions(request *ChatappVerifyAndRegisterRequest, runtime *dara.RuntimeOptions) (_result *ChatappVerifyAndRegisterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatappVerifyAndRegisterResponse
func (client *Client) ChatappVerifyAndRegister(request *ChatappVerifyAndRegisterRequest) (_result *ChatappVerifyAndRegisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatappVerifyAndRegisterResponse{}
	_body, _err := client.ChatappVerifyAndRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatFlowWithOptions(tmpReq *CreateChatFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateChatFlowRequest
//
// @return CreateChatFlowResponse
func (client *Client) CreateChatFlow(request *CreateChatFlowRequest) (_result *CreateChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatFlowResponse{}
	_body, _err := client.CreateChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatFlowByImportWithOptions(tmpReq *CreateChatFlowByImportRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowByImportResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateChatFlowByImportRequest
//
// @return CreateChatFlowByImportResponse
func (client *Client) CreateChatFlowByImport(request *CreateChatFlowByImportRequest) (_result *CreateChatFlowByImportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatFlowByImportResponse{}
	_body, _err := client.CreateChatFlowByImportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatFlowLogSettingWithOptions(request *CreateChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *CreateChatFlowLogSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChatFlowLogSettingResponse
func (client *Client) CreateChatFlowLogSetting(request *CreateChatFlowLogSettingRequest) (_result *CreateChatFlowLogSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatFlowLogSettingResponse{}
	_body, _err := client.CreateChatFlowLogSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatappMigrationInitiateWithOptions(request *CreateChatappMigrationInitiateRequest, runtime *dara.RuntimeOptions) (_result *CreateChatappMigrationInitiateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChatappMigrationInitiateResponse
func (client *Client) CreateChatappMigrationInitiate(request *CreateChatappMigrationInitiateRequest) (_result *CreateChatappMigrationInitiateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatappMigrationInitiateResponse{}
	_body, _err := client.CreateChatappMigrationInitiateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatappTemplateWithOptions(tmpReq *CreateChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateChatappTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateChatappTemplateRequest
//
// @return CreateChatappTemplateResponse
func (client *Client) CreateChatappTemplate(request *CreateChatappTemplateRequest) (_result *CreateChatappTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatappTemplateResponse{}
	_body, _err := client.CreateChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateFlowWithOptions(tmpReq *CreateFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateFlowRequest
//
// @return CreateFlowResponse
func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateFlowVersionWithOptions(tmpReq *CreateFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateFlowVersionRequest
//
// @return CreateFlowVersionResponse
func (client *Client) CreateFlowVersion(request *CreateFlowVersionRequest) (_result *CreateFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowVersionResponse{}
	_body, _err := client.CreateFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePhoneMessageQrdlWithOptions(request *CreatePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *CreatePhoneMessageQrdlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePhoneMessageQrdlResponse
func (client *Client) CreatePhoneMessageQrdl(request *CreatePhoneMessageQrdlRequest) (_result *CreatePhoneMessageQrdlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePhoneMessageQrdlResponse{}
	_body, _err := client.CreatePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChatFlowWithOptions(tmpReq *DeleteChatFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteChatFlowRequest
//
// @return DeleteChatFlowResponse
func (client *Client) DeleteChatFlow(request *DeleteChatFlowRequest) (_result *DeleteChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChatFlowResponse{}
	_body, _err := client.DeleteChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChatGroupWithOptions(request *DeleteChatGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChatGroupResponse
func (client *Client) DeleteChatGroup(request *DeleteChatGroupRequest) (_result *DeleteChatGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChatGroupResponse{}
	_body, _err := client.DeleteChatGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChatGroupInviteLinkWithOptions(request *DeleteChatGroupInviteLinkRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupInviteLinkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChatGroupInviteLinkResponse
func (client *Client) DeleteChatGroupInviteLink(request *DeleteChatGroupInviteLinkRequest) (_result *DeleteChatGroupInviteLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChatGroupInviteLinkResponse{}
	_body, _err := client.DeleteChatGroupInviteLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChatGroupParticipantsWithOptions(tmpReq *DeleteChatGroupParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatGroupParticipantsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteChatGroupParticipantsRequest
//
// @return DeleteChatGroupParticipantsResponse
func (client *Client) DeleteChatGroupParticipants(request *DeleteChatGroupParticipantsRequest) (_result *DeleteChatGroupParticipantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChatGroupParticipantsResponse{}
	_body, _err := client.DeleteChatGroupParticipantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteChatappTemplateWithOptions(request *DeleteChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteChatappTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteChatappTemplateResponse
func (client *Client) DeleteChatappTemplate(request *DeleteChatappTemplateRequest) (_result *DeleteChatappTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChatappTemplateResponse{}
	_body, _err := client.DeleteChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFlowResponse
func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFlowVersionWithOptions(tmpReq *DeleteFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteFlowVersionRequest
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersion(request *DeleteFlowVersionRequest) (_result *DeleteFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.DeleteFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePhoneMessageQrdlWithOptions(request *DeletePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *DeletePhoneMessageQrdlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePhoneMessageQrdlResponse
func (client *Client) DeletePhoneMessageQrdl(request *DeletePhoneMessageQrdlRequest) (_result *DeletePhoneMessageQrdlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePhoneMessageQrdlResponse{}
	_body, _err := client.DeletePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeprecateFlowWithOptions(request *DeprecateFlowRequest, runtime *dara.RuntimeOptions) (_result *DeprecateFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeprecateFlowResponse
func (client *Client) DeprecateFlow(request *DeprecateFlowRequest) (_result *DeprecateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeprecateFlowResponse{}
	_body, _err := client.DeprecateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EnableWhatsappROIMetricWithOptions(request *EnableWhatsappROIMetricRequest, runtime *dara.RuntimeOptions) (_result *EnableWhatsappROIMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EnableWhatsappROIMetricResponse
func (client *Client) EnableWhatsappROIMetric(request *EnableWhatsappROIMetricRequest) (_result *EnableWhatsappROIMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableWhatsappROIMetricResponse{}
	_body, _err := client.EnableWhatsappROIMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FlowBindPhoneWithOptions(tmpReq *FlowBindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowBindPhoneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FlowBindPhoneRequest
//
// @return FlowBindPhoneResponse
func (client *Client) FlowBindPhone(request *FlowBindPhoneRequest) (_result *FlowBindPhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlowBindPhoneResponse{}
	_body, _err := client.FlowBindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FlowRebindPhoneWithOptions(tmpReq *FlowRebindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowRebindPhoneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FlowRebindPhoneRequest
//
// @return FlowRebindPhoneResponse
func (client *Client) FlowRebindPhone(request *FlowRebindPhoneRequest) (_result *FlowRebindPhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlowRebindPhoneResponse{}
	_body, _err := client.FlowRebindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FlowUnbindPhoneWithOptions(tmpReq *FlowUnbindPhoneRequest, runtime *dara.RuntimeOptions) (_result *FlowUnbindPhoneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FlowUnbindPhoneRequest
//
// @return FlowUnbindPhoneResponse
func (client *Client) FlowUnbindPhone(request *FlowUnbindPhoneRequest) (_result *FlowUnbindPhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlowUnbindPhoneResponse{}
	_body, _err := client.FlowUnbindPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatFlowMetricWithOptions(tmpReq *GetChatFlowMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatFlowMetricResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetChatFlowMetricRequest
//
// @return GetChatFlowMetricResponse
func (client *Client) GetChatFlowMetric(request *GetChatFlowMetricRequest) (_result *GetChatFlowMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatFlowMetricResponse{}
	_body, _err := client.GetChatFlowMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatFlowTemplateWithOptions(request *GetChatFlowTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetChatFlowTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatFlowTemplateResponse
func (client *Client) GetChatFlowTemplate(request *GetChatFlowTemplateRequest) (_result *GetChatFlowTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatFlowTemplateResponse{}
	_body, _err := client.GetChatFlowTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatappPhoneNumberMetricWithOptions(request *GetChatappPhoneNumberMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatappPhoneNumberMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatappPhoneNumberMetricResponse
func (client *Client) GetChatappPhoneNumberMetric(request *GetChatappPhoneNumberMetricRequest) (_result *GetChatappPhoneNumberMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatappPhoneNumberMetricResponse{}
	_body, _err := client.GetChatappPhoneNumberMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatappTemplateDetailWithOptions(request *GetChatappTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetChatappTemplateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatappTemplateDetailResponse
func (client *Client) GetChatappTemplateDetail(request *GetChatappTemplateDetailRequest) (_result *GetChatappTemplateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatappTemplateDetailResponse{}
	_body, _err := client.GetChatappTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatappTemplateMetricWithOptions(request *GetChatappTemplateMetricRequest, runtime *dara.RuntimeOptions) (_result *GetChatappTemplateMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatappTemplateMetricResponse
func (client *Client) GetChatappTemplateMetric(request *GetChatappTemplateMetricRequest) (_result *GetChatappTemplateMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatappTemplateMetricResponse{}
	_body, _err := client.GetChatappTemplateMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatappUploadAuthorizationWithOptions(request *GetChatappUploadAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *GetChatappUploadAuthorizationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatappUploadAuthorizationResponse
func (client *Client) GetChatappUploadAuthorization(request *GetChatappUploadAuthorizationRequest) (_result *GetChatappUploadAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatappUploadAuthorizationResponse{}
	_body, _err := client.GetChatappUploadAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatappVerifyCodeWithOptions(request *GetChatappVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *GetChatappVerifyCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatappVerifyCodeResponse
func (client *Client) GetChatappVerifyCode(request *GetChatappVerifyCodeRequest) (_result *GetChatappVerifyCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatappVerifyCodeResponse{}
	_body, _err := client.GetChatappVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCommerceSettingWithOptions(request *GetCommerceSettingRequest, runtime *dara.RuntimeOptions) (_result *GetCommerceSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCommerceSettingResponse
func (client *Client) GetCommerceSetting(request *GetCommerceSettingRequest) (_result *GetCommerceSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCommerceSettingResponse{}
	_body, _err := client.GetCommerceSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetConversationalAutomationWithOptions(request *GetConversationalAutomationRequest, runtime *dara.RuntimeOptions) (_result *GetConversationalAutomationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetConversationalAutomationResponse
func (client *Client) GetConversationalAutomation(request *GetConversationalAutomationRequest) (_result *GetConversationalAutomationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConversationalAutomationResponse{}
	_body, _err := client.GetConversationalAutomationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFlowWithOptions(request *GetFlowRequest, runtime *dara.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFlowResponse
func (client *Client) GetFlow(request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFlowJSONAssestWithOptions(request *GetFlowJSONAssestRequest, runtime *dara.RuntimeOptions) (_result *GetFlowJSONAssestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFlowJSONAssestResponse
func (client *Client) GetFlowJSONAssest(request *GetFlowJSONAssestRequest) (_result *GetFlowJSONAssestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFlowJSONAssestResponse{}
	_body, _err := client.GetFlowJSONAssestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFlowPreviewUrlWithOptions(request *GetFlowPreviewUrlRequest, runtime *dara.RuntimeOptions) (_result *GetFlowPreviewUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFlowPreviewUrlResponse
func (client *Client) GetFlowPreviewUrl(request *GetFlowPreviewUrlRequest) (_result *GetFlowPreviewUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFlowPreviewUrlResponse{}
	_body, _err := client.GetFlowPreviewUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMigrationVerifyCodeWithOptions(request *GetMigrationVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *GetMigrationVerifyCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetMigrationVerifyCodeResponse
func (client *Client) GetMigrationVerifyCode(request *GetMigrationVerifyCodeRequest) (_result *GetMigrationVerifyCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMigrationVerifyCodeResponse{}
	_body, _err := client.GetMigrationVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPermissionByCodeWithOptions(tmpReq *GetPermissionByCodeRequest, runtime *dara.RuntimeOptions) (_result *GetPermissionByCodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetPermissionByCodeRequest
//
// @return GetPermissionByCodeResponse
func (client *Client) GetPermissionByCode(request *GetPermissionByCodeRequest) (_result *GetPermissionByCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPermissionByCodeResponse{}
	_body, _err := client.GetPermissionByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPhoneEncryptionPublicKeyWithOptions(request *GetPhoneEncryptionPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *GetPhoneEncryptionPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPhoneEncryptionPublicKeyResponse
func (client *Client) GetPhoneEncryptionPublicKey(request *GetPhoneEncryptionPublicKeyRequest) (_result *GetPhoneEncryptionPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhoneEncryptionPublicKeyResponse{}
	_body, _err := client.GetPhoneEncryptionPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPhoneNumberVerificationStatusWithOptions(request *GetPhoneNumberVerificationStatusRequest, runtime *dara.RuntimeOptions) (_result *GetPhoneNumberVerificationStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPhoneNumberVerificationStatusResponse
func (client *Client) GetPhoneNumberVerificationStatus(request *GetPhoneNumberVerificationStatusRequest) (_result *GetPhoneNumberVerificationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhoneNumberVerificationStatusResponse{}
	_body, _err := client.GetPhoneNumberVerificationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPreValidatePhoneIdWithOptions(request *GetPreValidatePhoneIdRequest, runtime *dara.RuntimeOptions) (_result *GetPreValidatePhoneIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPreValidatePhoneIdResponse
func (client *Client) GetPreValidatePhoneId(request *GetPreValidatePhoneIdRequest) (_result *GetPreValidatePhoneIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPreValidatePhoneIdResponse{}
	_body, _err := client.GetPreValidatePhoneIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWhatsappConnectionCatalogWithOptions(request *GetWhatsappConnectionCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetWhatsappConnectionCatalogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWhatsappConnectionCatalogResponse
func (client *Client) GetWhatsappConnectionCatalog(request *GetWhatsappConnectionCatalogRequest) (_result *GetWhatsappConnectionCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWhatsappConnectionCatalogResponse{}
	_body, _err := client.GetWhatsappConnectionCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWhatsappHealthStatusWithOptions(request *GetWhatsappHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *GetWhatsappHealthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWhatsappHealthStatusResponse
func (client *Client) GetWhatsappHealthStatus(request *GetWhatsappHealthStatusRequest) (_result *GetWhatsappHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWhatsappHealthStatusResponse{}
	_body, _err := client.GetWhatsappHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) IsvGetAppIdWithOptions(request *IsvGetAppIdRequest, runtime *dara.RuntimeOptions) (_result *IsvGetAppIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return IsvGetAppIdResponse
func (client *Client) IsvGetAppId(request *IsvGetAppIdRequest) (_result *IsvGetAppIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IsvGetAppIdResponse{}
	_body, _err := client.IsvGetAppIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListBindingRelationsForFlowVersionWithOptions(request *ListBindingRelationsForFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ListBindingRelationsForFlowVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListBindingRelationsForFlowVersionResponse
func (client *Client) ListBindingRelationsForFlowVersion(request *ListBindingRelationsForFlowVersionRequest) (_result *ListBindingRelationsForFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBindingRelationsForFlowVersionResponse{}
	_body, _err := client.ListBindingRelationsForFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatFlowWithOptions(tmpReq *ListChatFlowRequest, runtime *dara.RuntimeOptions) (_result *ListChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListChatFlowRequest
//
// @return ListChatFlowResponse
func (client *Client) ListChatFlow(request *ListChatFlowRequest) (_result *ListChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatFlowResponse{}
	_body, _err := client.ListChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatFlowTemplateWithOptions(request *ListChatFlowTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListChatFlowTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListChatFlowTemplateResponse
func (client *Client) ListChatFlowTemplate(request *ListChatFlowTemplateRequest) (_result *ListChatFlowTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatFlowTemplateResponse{}
	_body, _err := client.ListChatFlowTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatGroupWithOptions(tmpReq *ListChatGroupRequest, runtime *dara.RuntimeOptions) (_result *ListChatGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListChatGroupRequest
//
// @return ListChatGroupResponse
func (client *Client) ListChatGroup(request *ListChatGroupRequest) (_result *ListChatGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatGroupResponse{}
	_body, _err := client.ListChatGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatGroupParticipantsWithOptions(tmpReq *ListChatGroupParticipantsRequest, runtime *dara.RuntimeOptions) (_result *ListChatGroupParticipantsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListChatGroupParticipantsRequest
//
// @return ListChatGroupParticipantsResponse
func (client *Client) ListChatGroupParticipants(request *ListChatGroupParticipantsRequest) (_result *ListChatGroupParticipantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatGroupParticipantsResponse{}
	_body, _err := client.ListChatGroupParticipantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatappMessageWithOptions(tmpReq *ListChatappMessageRequest, runtime *dara.RuntimeOptions) (_result *ListChatappMessageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListChatappMessageRequest
//
// @return ListChatappMessageResponse
func (client *Client) ListChatappMessage(request *ListChatappMessageRequest) (_result *ListChatappMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatappMessageResponse{}
	_body, _err := client.ListChatappMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListChatappTemplateWithOptions(tmpReq *ListChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListChatappTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListChatappTemplateRequest
//
// @return ListChatappTemplateResponse
func (client *Client) ListChatappTemplate(request *ListChatappTemplateRequest) (_result *ListChatappTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatappTemplateResponse{}
	_body, _err := client.ListChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFlowWithOptions(tmpReq *ListFlowRequest, runtime *dara.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFlowRequest
//
// @return ListFlowResponse
func (client *Client) ListFlow(request *ListFlowRequest) (_result *ListFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowResponse{}
	_body, _err := client.ListFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFlowNodePrototypeV2WithOptions(request *ListFlowNodePrototypeV2Request, runtime *dara.RuntimeOptions) (_result *ListFlowNodePrototypeV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListFlowNodePrototypeV2Response
func (client *Client) ListFlowNodePrototypeV2(request *ListFlowNodePrototypeV2Request) (_result *ListFlowNodePrototypeV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowNodePrototypeV2Response{}
	_body, _err := client.ListFlowNodePrototypeV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFlowVersionWithOptions(tmpReq *ListFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ListFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListFlowVersionRequest
//
// @return ListFlowVersionResponse
func (client *Client) ListFlowVersion(request *ListFlowVersionRequest) (_result *ListFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowVersionResponse{}
	_body, _err := client.ListFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPhoneMessageQrdlWithOptions(request *ListPhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *ListPhoneMessageQrdlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPhoneMessageQrdlResponse
func (client *Client) ListPhoneMessageQrdl(request *ListPhoneMessageQrdlRequest) (_result *ListPhoneMessageQrdlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPhoneMessageQrdlResponse{}
	_body, _err := client.ListPhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProductWithOptions(request *ListProductRequest, runtime *dara.RuntimeOptions) (_result *ListProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProductResponse
func (client *Client) ListProduct(request *ListProductRequest) (_result *ListProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductResponse{}
	_body, _err := client.ListProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProductCatalogWithOptions(request *ListProductCatalogRequest, runtime *dara.RuntimeOptions) (_result *ListProductCatalogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProductCatalogResponse
func (client *Client) ListProductCatalog(request *ListProductCatalogRequest) (_result *ListProductCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductCatalogResponse{}
	_body, _err := client.ListProductCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyChatappTemplateWithOptions(tmpReq *ModifyChatappTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyChatappTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ModifyChatappTemplateRequest
//
// @return ModifyChatappTemplateResponse
func (client *Client) ModifyChatappTemplate(request *ModifyChatappTemplateRequest) (_result *ModifyChatappTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyChatappTemplateResponse{}
	_body, _err := client.ModifyChatappTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyChatappTemplatePropertiesWithOptions(request *ModifyChatappTemplatePropertiesRequest, runtime *dara.RuntimeOptions) (_result *ModifyChatappTemplatePropertiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyChatappTemplatePropertiesResponse
func (client *Client) ModifyChatappTemplateProperties(request *ModifyChatappTemplatePropertiesRequest) (_result *ModifyChatappTemplatePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyChatappTemplatePropertiesResponse{}
	_body, _err := client.ModifyChatappTemplatePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyFlowWithOptions(tmpReq *ModifyFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ModifyFlowRequest
//
// @return ModifyFlowResponse
func (client *Client) ModifyFlow(request *ModifyFlowRequest) (_result *ModifyFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFlowResponse{}
	_body, _err := client.ModifyFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyPhoneBusinessProfileWithOptions(tmpReq *ModifyPhoneBusinessProfileRequest, runtime *dara.RuntimeOptions) (_result *ModifyPhoneBusinessProfileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ModifyPhoneBusinessProfileRequest
//
// @return ModifyPhoneBusinessProfileResponse
func (client *Client) ModifyPhoneBusinessProfile(request *ModifyPhoneBusinessProfileRequest) (_result *ModifyPhoneBusinessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPhoneBusinessProfileResponse{}
	_body, _err := client.ModifyPhoneBusinessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OfflineFlowVersionWithOptions(tmpReq *OfflineFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *OfflineFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OfflineFlowVersionRequest
//
// @return OfflineFlowVersionResponse
func (client *Client) OfflineFlowVersion(request *OfflineFlowVersionRequest) (_result *OfflineFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineFlowVersionResponse{}
	_body, _err := client.OfflineFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OnlineFlowVersionWithOptions(tmpReq *OnlineFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *OnlineFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - OnlineFlowVersionRequest
//
// @return OnlineFlowVersionResponse
func (client *Client) OnlineFlowVersion(request *OnlineFlowVersionRequest) (_result *OnlineFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OnlineFlowVersionResponse{}
	_body, _err := client.OnlineFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PublishFlowWithOptions(request *PublishFlowRequest, runtime *dara.RuntimeOptions) (_result *PublishFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PublishFlowResponse
func (client *Client) PublishFlow(request *PublishFlowRequest) (_result *PublishFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishFlowResponse{}
	_body, _err := client.PublishFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryChatappBindWabaWithOptions(request *QueryChatappBindWabaRequest, runtime *dara.RuntimeOptions) (_result *QueryChatappBindWabaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryChatappBindWabaResponse
func (client *Client) QueryChatappBindWaba(request *QueryChatappBindWabaRequest) (_result *QueryChatappBindWabaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryChatappBindWabaResponse{}
	_body, _err := client.QueryChatappBindWabaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryChatappPhoneNumbersWithOptions(request *QueryChatappPhoneNumbersRequest, runtime *dara.RuntimeOptions) (_result *QueryChatappPhoneNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryChatappPhoneNumbersResponse
func (client *Client) QueryChatappPhoneNumbers(request *QueryChatappPhoneNumbersRequest) (_result *QueryChatappPhoneNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryChatappPhoneNumbersResponse{}
	_body, _err := client.QueryChatappPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPhoneBusinessProfileWithOptions(request *QueryPhoneBusinessProfileRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneBusinessProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPhoneBusinessProfileResponse
func (client *Client) QueryPhoneBusinessProfile(request *QueryPhoneBusinessProfileRequest) (_result *QueryPhoneBusinessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPhoneBusinessProfileResponse{}
	_body, _err := client.QueryPhoneBusinessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryWabaBusinessInfoWithOptions(request *QueryWabaBusinessInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryWabaBusinessInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryWabaBusinessInfoResponse
func (client *Client) QueryWabaBusinessInfo(request *QueryWabaBusinessInfoRequest) (_result *QueryWabaBusinessInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWabaBusinessInfoResponse{}
	_body, _err := client.QueryWabaBusinessInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadChatFlowWithOptions(tmpReq *ReadChatFlowRequest, runtime *dara.RuntimeOptions) (_result *ReadChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ReadChatFlowRequest
//
// @return ReadChatFlowResponse
func (client *Client) ReadChatFlow(request *ReadChatFlowRequest) (_result *ReadChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadChatFlowResponse{}
	_body, _err := client.ReadChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadChatFlowLogSettingWithOptions(request *ReadChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *ReadChatFlowLogSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadChatFlowLogSettingResponse
func (client *Client) ReadChatFlowLogSetting(request *ReadChatFlowLogSettingRequest) (_result *ReadChatFlowLogSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadChatFlowLogSettingResponse{}
	_body, _err := client.ReadChatFlowLogSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadFlowVersionWithOptions(tmpReq *ReadFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *ReadFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ReadFlowVersionRequest
//
// @return ReadFlowVersionResponse
func (client *Client) ReadFlowVersion(request *ReadFlowVersionRequest) (_result *ReadFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadFlowVersionResponse{}
	_body, _err := client.ReadFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendChatappMassMessageWithOptions(tmpReq *SendChatappMassMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatappMassMessageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SendChatappMassMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SenderList) {
		request.SenderListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SenderList, dara.String("SenderList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		body["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.CustSpaceId) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.FallBackContent) {
		body["FallBackContent"] = request.FallBackContent
	}

	if !dara.IsNil(request.FallBackDuration) {
		body["FallBackDuration"] = request.FallBackDuration
	}

	if !dara.IsNil(request.FallBackId) {
		body["FallBackId"] = request.FallBackId
	}

	if !dara.IsNil(request.FallBackRule) {
		body["FallBackRule"] = request.FallBackRule
	}

	if !dara.IsNil(request.From) {
		body["From"] = request.From
	}

	if !dara.IsNil(request.IsvCode) {
		body["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.SenderListShrink) {
		body["SenderList"] = request.SenderListShrink
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.Ttl) {
		body["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SendChatappMassMessageRequest
//
// @return SendChatappMassMessageResponse
func (client *Client) SendChatappMassMessage(request *SendChatappMassMessageRequest) (_result *SendChatappMassMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendChatappMassMessageResponse{}
	_body, _err := client.SendChatappMassMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendChatappMessageWithOptions(tmpReq *SendChatappMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatappMessageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.PayloadShrink) {
		query["Payload"] = request.PayloadShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelType) {
		body["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ContextMessageId) {
		body["ContextMessageId"] = request.ContextMessageId
	}

	if !dara.IsNil(request.CustSpaceId) {
		body["CustSpaceId"] = request.CustSpaceId
	}

	if !dara.IsNil(request.CustWabaId) {
		body["CustWabaId"] = request.CustWabaId
	}

	if !dara.IsNil(request.FallBackContent) {
		body["FallBackContent"] = request.FallBackContent
	}

	if !dara.IsNil(request.FallBackDuration) {
		body["FallBackDuration"] = request.FallBackDuration
	}

	if !dara.IsNil(request.FallBackId) {
		body["FallBackId"] = request.FallBackId
	}

	if !dara.IsNil(request.FallBackRule) {
		body["FallBackRule"] = request.FallBackRule
	}

	if !dara.IsNil(request.FlowActionShrink) {
		body["FlowAction"] = request.FlowActionShrink
	}

	if !dara.IsNil(request.From) {
		body["From"] = request.From
	}

	if !dara.IsNil(request.IsvCode) {
		body["IsvCode"] = request.IsvCode
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.Language) {
		body["Language"] = request.Language
	}

	if !dara.IsNil(request.MessageType) {
		body["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.ProductActionShrink) {
		body["ProductAction"] = request.ProductActionShrink
	}

	if !dara.IsNil(request.RecipientType) {
		body["RecipientType"] = request.RecipientType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateParamsShrink) {
		body["TemplateParams"] = request.TemplateParamsShrink
	}

	if !dara.IsNil(request.To) {
		body["To"] = request.To
	}

	if !dara.IsNil(request.TrackingData) {
		body["TrackingData"] = request.TrackingData
	}

	if !dara.IsNil(request.Ttl) {
		body["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SendChatappMessageRequest
//
// @return SendChatappMessageResponse
func (client *Client) SendChatappMessage(request *SendChatappMessageRequest) (_result *SendChatappMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendChatappMessageResponse{}
	_body, _err := client.SendChatappMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits the agreement information for independent software vendor (ISV) customers.
//
// Description:
//
//	  You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
//		- After you call the [GetChatappUploadAuthorization](~~GetChatappUploadAuthorization~~) operation to obtain the authentication information for uploading the file to Object Storage Service (OSS), you can use the authentication information to upload the file to the OSS server. To upload the file, you can call the SDK provided by OSS. When you upload the file, set the value of the key to the value of `Dir + "/" + file name`, such as C200293990209/isvTerms.pdf. The value of Dir is obtained from the [GetChatappUploadAuthorization](~~GetChatappUploadAuthorization~~) operation. The value of IsvTerms is obtained from the PutObject operation.
//
// @param request - SubmitIsvCustomerTermsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIsvCustomerTermsResponse
func (client *Client) SubmitIsvCustomerTermsWithOptions(request *SubmitIsvCustomerTermsRequest, runtime *dara.RuntimeOptions) (_result *SubmitIsvCustomerTermsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessDesc) {
		query["BusinessDesc"] = request.BusinessDesc
	}

	if !dara.IsNil(request.ContactMail) {
		query["ContactMail"] = request.ContactMail
	}

	if !dara.IsNil(request.CountryId) {
		query["CountryId"] = request.CountryId
	}

	if !dara.IsNil(request.CustName) {
		query["CustName"] = request.CustName
	}

	if !dara.IsNil(request.CustSpaceId) {
		query["CustSpaceId"] = request.CustSpaceId
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
		Action:      dara.String("SubmitIsvCustomerTerms"),
		Version:     dara.String("2020-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIsvCustomerTermsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits the agreement information for independent software vendor (ISV) customers.
//
// Description:
//
//	  You can call this operation up to five times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
//		- After you call the [GetChatappUploadAuthorization](~~GetChatappUploadAuthorization~~) operation to obtain the authentication information for uploading the file to Object Storage Service (OSS), you can use the authentication information to upload the file to the OSS server. To upload the file, you can call the SDK provided by OSS. When you upload the file, set the value of the key to the value of `Dir + "/" + file name`, such as C200293990209/isvTerms.pdf. The value of Dir is obtained from the [GetChatappUploadAuthorization](~~GetChatappUploadAuthorization~~) operation. The value of IsvTerms is obtained from the PutObject operation.
//
// @param request - SubmitIsvCustomerTermsRequest
//
// @return SubmitIsvCustomerTermsResponse
func (client *Client) SubmitIsvCustomerTerms(request *SubmitIsvCustomerTermsRequest) (_result *SubmitIsvCustomerTermsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIsvCustomerTermsResponse{}
	_body, _err := client.SubmitIsvCustomerTermsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TriggerChatFlowWithOptions(tmpReq *TriggerChatFlowRequest, runtime *dara.RuntimeOptions) (_result *TriggerChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - TriggerChatFlowRequest
//
// @return TriggerChatFlowResponse
func (client *Client) TriggerChatFlow(request *TriggerChatFlowRequest) (_result *TriggerChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerChatFlowResponse{}
	_body, _err := client.TriggerChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAccountWebhookWithOptions(request *UpdateAccountWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountWebhookResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAccountWebhookResponse
func (client *Client) UpdateAccountWebhook(request *UpdateAccountWebhookRequest) (_result *UpdateAccountWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccountWebhookResponse{}
	_body, _err := client.UpdateAccountWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChatFlowWithOptions(tmpReq *UpdateChatFlowRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatFlowResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateChatFlowRequest
//
// @return UpdateChatFlowResponse
func (client *Client) UpdateChatFlow(request *UpdateChatFlowRequest) (_result *UpdateChatFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChatFlowResponse{}
	_body, _err := client.UpdateChatFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChatFlowLogSettingWithOptions(request *UpdateChatFlowLogSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatFlowLogSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateChatFlowLogSettingResponse
func (client *Client) UpdateChatFlowLogSetting(request *UpdateChatFlowLogSettingRequest) (_result *UpdateChatFlowLogSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChatFlowLogSettingResponse{}
	_body, _err := client.UpdateChatFlowLogSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateChatGroupWithOptions(request *UpdateChatGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateChatGroupResponse
func (client *Client) UpdateChatGroup(request *UpdateChatGroupRequest) (_result *UpdateChatGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChatGroupResponse{}
	_body, _err := client.UpdateChatGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCommerceSettingWithOptions(request *UpdateCommerceSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateCommerceSettingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCommerceSettingResponse
func (client *Client) UpdateCommerceSetting(request *UpdateCommerceSettingRequest) (_result *UpdateCommerceSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCommerceSettingResponse{}
	_body, _err := client.UpdateCommerceSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateConversationalAutomationWithOptions(tmpReq *UpdateConversationalAutomationRequest, runtime *dara.RuntimeOptions) (_result *UpdateConversationalAutomationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateConversationalAutomationRequest
//
// @return UpdateConversationalAutomationResponse
func (client *Client) UpdateConversationalAutomation(request *UpdateConversationalAutomationRequest) (_result *UpdateConversationalAutomationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConversationalAutomationResponse{}
	_body, _err := client.UpdateConversationalAutomationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFlowJSONAssetWithOptions(request *UpdateFlowJSONAssetRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowJSONAssetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateFlowJSONAssetResponse
func (client *Client) UpdateFlowJSONAsset(request *UpdateFlowJSONAssetRequest) (_result *UpdateFlowJSONAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlowJSONAssetResponse{}
	_body, _err := client.UpdateFlowJSONAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateFlowVersionWithOptions(tmpReq *UpdateFlowVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowVersionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateFlowVersionRequest
//
// @return UpdateFlowVersionResponse
func (client *Client) UpdateFlowVersion(request *UpdateFlowVersionRequest) (_result *UpdateFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlowVersionResponse{}
	_body, _err := client.UpdateFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePhoneEncryptionPublicKeyWithOptions(request *UpdatePhoneEncryptionPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneEncryptionPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePhoneEncryptionPublicKeyResponse
func (client *Client) UpdatePhoneEncryptionPublicKey(request *UpdatePhoneEncryptionPublicKeyRequest) (_result *UpdatePhoneEncryptionPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePhoneEncryptionPublicKeyResponse{}
	_body, _err := client.UpdatePhoneEncryptionPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePhoneMessageQrdlWithOptions(request *UpdatePhoneMessageQrdlRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneMessageQrdlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePhoneMessageQrdlResponse
func (client *Client) UpdatePhoneMessageQrdl(request *UpdatePhoneMessageQrdlRequest) (_result *UpdatePhoneMessageQrdlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePhoneMessageQrdlResponse{}
	_body, _err := client.UpdatePhoneMessageQrdlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePhoneWebhookWithOptions(request *UpdatePhoneWebhookRequest, runtime *dara.RuntimeOptions) (_result *UpdatePhoneWebhookResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePhoneWebhookResponse
func (client *Client) UpdatePhoneWebhook(request *UpdatePhoneWebhookRequest) (_result *UpdatePhoneWebhookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePhoneWebhookResponse{}
	_body, _err := client.UpdatePhoneWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
