// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Accepts an invitation.
//
// Description:
//
// After an invited Alibaba Cloud account joins a resource directory, it becomes a member of the resource directory. By default, the name of the invited Alibaba Cloud account is used as the display name of the account in the resource directory.
//
// @param request - AcceptHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptHandshakeResponse
func (client *Client) AcceptHandshakeWithContext(ctx context.Context, request *AcceptHandshakeRequest, runtime *dara.RuntimeOptions) (_result *AcceptHandshakeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HandshakeId) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptHandshake"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a contact.
//
// @param request - AddMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMessageContactResponse
func (client *Client) AddMessageContactWithContext(ctx context.Context, request *AddMessageContactRequest, runtime *dara.RuntimeOptions) (_result *AddMessageContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EmailAddress) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !dara.IsNil(request.MessageTypes) {
		query["MessageTypes"] = request.MessageTypes
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a contact to a resource directory, folder, or member.
//
// @param request - AssociateMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateMembersResponse
func (client *Client) AssociateMembersWithContext(ctx context.Context, request *AssociateMembersRequest, runtime *dara.RuntimeOptions) (_result *AssociateMembersResponse, _err error) {
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

	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateMembers"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an access control policy.
//
// Description:
//
// After you attach a custom access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
//
// By default, the system access control policy FullAliyunAccess is attached to each folder and member.
//
// The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
//
// A maximum of 10 access control policies can be attached to a folder or member.
//
// @param request - AttachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachControlPolicyResponse
func (client *Client) AttachControlPolicyWithContext(ctx context.Context, request *AttachControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *AttachControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a mobile phone number to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - BindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSecureMobilePhoneResponse
func (client *Client) BindSecureMobilePhoneWithContext(ctx context.Context, request *BindSecureMobilePhoneRequest, runtime *dara.RuntimeOptions) (_result *BindSecureMobilePhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.SecureMobilePhone) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	if !dara.IsNil(request.VerificationCode) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSecureMobilePhone"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the email address change of a member.
//
// @param request - CancelChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelChangeAccountEmailResponse
func (client *Client) CancelChangeAccountEmailWithContext(ctx context.Context, request *CancelChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *CancelChangeAccountEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelChangeAccountEmail"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an invitation.
//
// @param request - CancelHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelHandshakeResponse
func (client *Client) CancelHandshakeWithContext(ctx context.Context, request *CancelHandshakeRequest, runtime *dara.RuntimeOptions) (_result *CancelHandshakeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HandshakeId) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelHandshake"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the update of the mobile phone number or email address of a contact.
//
// @param request - CancelMessageContactUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelMessageContactUpdateResponse
func (client *Client) CancelMessageContactUpdateWithContext(ctx context.Context, request *CancelMessageContactUpdateRequest, runtime *dara.RuntimeOptions) (_result *CancelMessageContactUpdateResponse, _err error) {
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

	if !dara.IsNil(request.EmailAddress) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelMessageContactUpdate"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelMessageContactUpdateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the email address that is bound to a member.
//
// @param request - ChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAccountEmailResponse
func (client *Client) ChangeAccountEmailWithContext(ctx context.Context, request *ChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ChangeAccountEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAccountEmail"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a member deletion check.
//
// Description:
//
// Before you delete a member, you must call this API operation to check whether the member can be deleted.
//
// @param request - CheckAccountDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountDeleteResponse
func (client *Client) CheckAccountDeleteWithContext(ctx context.Context, request *CheckAccountDeleteRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountDelete"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom access control policy.
//
// @param request - CreateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateControlPolicyResponse
func (client *Client) CreateControlPolicyWithContext(ctx context.Context, request *CreateControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EffectScope) {
		query["EffectScope"] = request.EffectScope
	}

	if !dara.IsNil(request.PolicyDocument) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a folder.
//
// Description:
//
// A maximum of five levels of folders can be created under the Root folder.
//
// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderName) {
		query["FolderName"] = request.FolderName
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a member of the resource account type.
//
// Description:
//
// A member serves as a container for resources and is also an organizational unit in a resource directory. A member indicates a project or application. The resources of different members are isolated.
//
// This topic provides an example on how to call the API operation to create a member in the `fd-r23M55****` folder. The display name of the member is `Dev`, and the prefix for the Alibaba Cloud account name of the member is `alice`.
//
// @param request - CreateResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceAccountResponse
func (client *Client) CreateResourceAccountWithContext(ctx context.Context, request *CreateResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountNamePrefix) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.PayerAccountId) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	if !dara.IsNil(request.ResellAccountType) {
		query["ResellAccountType"] = request.ResellAccountType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects an invitation.
//
// @param request - DeclineHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeclineHandshakeResponse
func (client *Client) DeclineHandshakeWithContext(ctx context.Context, request *DeclineHandshakeRequest, runtime *dara.RuntimeOptions) (_result *DeclineHandshakeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HandshakeId) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeclineHandshake"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a member of the resource account type.
//
// Description:
//
// Before you delete a member, we recommend that you call the [CheckAccountDelete](~~CheckAccountDelete~~) and [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operations to check whether the member meets deletion requirements. You can call the DeleteAccount operation to delete only members that meet the deletion requirements.
//
// After you submit a deletion request for a member, you can call the [GetAccountDeletionStatus](~~GetAccountDeletionStatus~~) operation to query the deletion status of the member. After a member is deleted, the resources and data within the member are deleted, and you can no longer use the member to log on to the Alibaba Cloud Management Console. In addition, the member cannot be recovered. Proceed with caution. For more information about how to delete a member, see [Delete a member of the resource account type](https://help.aliyun.com/document_detail/446078.html).
//
// @param tmpReq - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithContext(ctx context.Context, tmpReq *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AbandonableCheckId) {
		request.AbandonableCheckIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AbandonableCheckId, dara.String("AbandonableCheckId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AbandonableCheckIdShrink) {
		query["AbandonableCheckId"] = request.AbandonableCheckIdShrink
	}

	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom access control policy.
//
// Description:
//
// If you want to delete a custom access control policy that is attached to folders or members, you must call the [DetachControlPolicy](~~DetachControlPolicy~~) operation to detach the policy before you delete it.
//
// @param request - DeleteControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicyWithContext(ctx context.Context, request *DeleteControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// Description:
//
// Before you delete a folder, you must make sure that the folder does not contain members or subfolders.
//
// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a contact.
//
// @param request - DeleteMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessageContactResponse
func (client *Client) DeleteMessageContactWithContext(ctx context.Context, request *DeleteMessageContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteMessageContactResponse, _err error) {
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

	if !dara.IsNil(request.RetainContactInMembers) {
		query["RetainContactInMembers"] = request.RetainContactInMembers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a delegated administrator account for a trusted service.
//
// Description:
//
// If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
//
// @param request - DeregisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterDelegatedAdministratorResponse
func (client *Client) DeregisterDelegatedAdministratorWithContext(ctx context.Context, request *DeregisterDelegatedAdministratorRequest, runtime *dara.RuntimeOptions) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ServicePrincipal) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeregisterDelegatedAdministrator"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches an access control policy.
//
// Description:
//
// After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
//
// Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
//
// @param request - DetachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachControlPolicyResponse
func (client *Client) DetachControlPolicyWithContext(ctx context.Context, request *DetachControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DetachControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a contact from a resource directory, folder, or member.
//
// @param request - DisassociateMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateMembersResponse
func (client *Client) DisassociateMembersWithContext(ctx context.Context, request *DisassociateMembersRequest, runtime *dara.RuntimeOptions) (_result *DisassociateMembersResponse, _err error) {
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

	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateMembers"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a resource directory.
//
// Description:
//
// You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// @param request - EnableResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceDirectoryResponse
func (client *Client) EnableResourceDirectoryWithContext(ctx context.Context, request *EnableResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *EnableResourceDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableMode) {
		query["EnableMode"] = request.EnableMode
	}

	if !dara.IsNil(request.MAName) {
		query["MAName"] = request.MAName
	}

	if !dara.IsNil(request.MASecureMobilePhone) {
		query["MASecureMobilePhone"] = request.MASecureMobilePhone
	}

	if !dara.IsNil(request.VerificationCode) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableResourceDirectory"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a member.
//
// @param request - GetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountResponse
func (client *Client) GetAccountWithContext(ctx context.Context, request *GetAccountRequest, runtime *dara.RuntimeOptions) (_result *GetAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a member deletion check.
//
// Description:
//
// After you call the [CheckAccountDelete](~~CheckAccountDelete~~) operation to perform a member deletion check, you can call the [GetAccountDeletionCheckResult](~~GetAccountDeletionCheckResult~~) operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
//
// @param request - GetAccountDeletionCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionCheckResultResponse
func (client *Client) GetAccountDeletionCheckResultWithContext(ctx context.Context, request *GetAccountDeletionCheckResultRequest, runtime *dara.RuntimeOptions) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountDeletionCheckResult"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a member.
//
// @param request - GetAccountDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionStatusResponse
func (client *Client) GetAccountDeletionStatusWithContext(ctx context.Context, request *GetAccountDeletionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAccountDeletionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountDeletionStatus"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access control policy.
//
// @param request - GetControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetControlPolicyResponse
func (client *Client) GetControlPolicyWithContext(ctx context.Context, request *GetControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetControlPolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithContext(ctx context.Context, request *GetFolderRequest, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an invitation.
//
// @param request - GetHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHandshakeResponse
func (client *Client) GetHandshakeWithContext(ctx context.Context, request *GetHandshakeRequest, runtime *dara.RuntimeOptions) (_result *GetHandshakeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HandshakeId) {
		query["HandshakeId"] = request.HandshakeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHandshake"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHandshakeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a contact.
//
// @param request - GetMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageContactResponse
func (client *Client) GetMessageContactWithContext(ctx context.Context, request *GetMessageContactRequest, runtime *dara.RuntimeOptions) (_result *GetMessageContactResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a contact.
//
// @param request - GetMessageContactDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageContactDeletionStatusResponse
func (client *Client) GetMessageContactDeletionStatusWithContext(ctx context.Context, request *GetMessageContactDeletionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMessageContactDeletionStatusResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageContactDeletionStatus"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageContactDeletionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a billing account.
//
// @param request - GetPayerForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPayerForAccountResponse
func (client *Client) GetPayerForAccountWithContext(ctx context.Context, request *GetPayerForAccountRequest, runtime *dara.RuntimeOptions) (_result *GetPayerForAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPayerForAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites an account to join a resource directory.
//
// @param request - InviteAccountToResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteAccountToResourceDirectoryResponse
func (client *Client) InviteAccountToResourceDirectoryWithContext(ctx context.Context, request *InviteAccountToResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetEntity) {
		query["TargetEntity"] = request.TargetEntity
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InviteAccountToResourceDirectory"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of members in a resource directory.
//
// Description:
//
// You can use only the management account of a resource directory or a delegated administrator account of a trusted service to call this operation.
//
// @param request - ListAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsResponse
func (client *Client) ListAccountsWithContext(ctx context.Context, request *ListAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccounts"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of members in a folder.
//
// @param request - ListAccountsForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsForParentResponse
func (client *Client) ListAccountsForParentWithContext(ctx context.Context, request *ListAccountsForParentRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsForParentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccountsForParent"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent folders of a specified folder.
//
// @param request - ListAncestorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAncestorsResponse
func (client *Client) ListAncestorsWithContext(ctx context.Context, request *ListAncestorsRequest, runtime *dara.RuntimeOptions) (_result *ListAncestorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildId) {
		query["ChildId"] = request.ChildId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAncestors"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAncestorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of members who have the permission to query member information in a resource directory.
//
// Description:
//
// The permission to query member information refers to the permission to call the [GetAccount](~~GetAccount~~) API operation.
//
// @param request - ListAuthorizedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedAccountsResponse
func (client *Client) ListAuthorizedAccountsWithContext(ctx context.Context, request *ListAuthorizedAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedAccounts"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of folders that have permissions to query subfolder information in a resource directory.
//
// Description:
//
// The permissions to query subfolder information refer to the permissions to call the [ListAccountsForParent](~~ListAccountsForParent~~) and [ListFoldersForParent](~~ListFoldersForParent~~) API operations.
//
// @param request - ListAuthorizedFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedFoldersResponse
func (client *Client) ListAuthorizedFoldersWithContext(ctx context.Context, request *ListAuthorizedFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedFoldersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedFolders"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedFoldersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control policies.
//
// @param request - ListControlPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPoliciesResponse
func (client *Client) ListControlPoliciesWithContext(ctx context.Context, request *ListControlPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListControlPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListControlPolicies"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control policies that are attached to a folder or member.
//
// @param request - ListControlPolicyAttachmentsForTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPolicyAttachmentsForTargetResponse
func (client *Client) ListControlPolicyAttachmentsForTargetWithContext(ctx context.Context, request *ListControlPolicyAttachmentsForTargetRequest, runtime *dara.RuntimeOptions) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
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

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListControlPolicyAttachmentsForTarget"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries delegated administrator accounts.
//
// @param request - ListDelegatedAdministratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedAdministratorsResponse
func (client *Client) ListDelegatedAdministratorsWithContext(ctx context.Context, request *ListDelegatedAdministratorsRequest, runtime *dara.RuntimeOptions) (_result *ListDelegatedAdministratorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServicePrincipal) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDelegatedAdministrators"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trusted services for which a member is specified as a delegated administrator account.
//
// @param request - ListDelegatedServicesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedServicesForAccountResponse
func (client *Client) ListDelegatedServicesForAccountWithContext(ctx context.Context, request *ListDelegatedServicesForAccountRequest, runtime *dara.RuntimeOptions) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDelegatedServicesForAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of all subfolders of a folder.
//
// Description:
//
// You can call this API operation to view the information of only the first-level subfolders of a folder.
//
// @param request - ListFoldersForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersForParentResponse
func (client *Client) ListFoldersForParentWithContext(ctx context.Context, request *ListFoldersForParentRequest, runtime *dara.RuntimeOptions) (_result *ListFoldersForParentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFoldersForParent"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invitations that are associated with an account.
//
// @param request - ListHandshakesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForAccountResponse
func (client *Client) ListHandshakesForAccountWithContext(ctx context.Context, request *ListHandshakesForAccountRequest, runtime *dara.RuntimeOptions) (_result *ListHandshakesForAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHandshakesForAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries invitations in a resource directory.
//
// @param request - ListHandshakesForResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForResourceDirectoryResponse
func (client *Client) ListHandshakesForResourceDirectoryWithContext(ctx context.Context, request *ListHandshakesForResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHandshakesForResourceDirectory"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mobile phone number or email address to be verified for a contact.
//
// @param request - ListMessageContactVerificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageContactVerificationsResponse
func (client *Client) ListMessageContactVerificationsWithContext(ctx context.Context, request *ListMessageContactVerificationsRequest, runtime *dara.RuntimeOptions) (_result *ListMessageContactVerificationsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageContactVerifications"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageContactVerificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries contacts.
//
// @param request - ListMessageContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessageContactsResponse
func (client *Client) ListMessageContactsWithContext(ctx context.Context, request *ListMessageContactsRequest, runtime *dara.RuntimeOptions) (_result *ListMessageContactsResponse, _err error) {
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

	if !dara.IsNil(request.Member) {
		query["Member"] = request.Member
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessageContacts"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessageContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyFilter) {
		query["KeyFilter"] = request.KeyFilter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to the members in a resource directory.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of a tag key.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.ValueFilter) {
		query["ValueFilter"] = request.ValueFilter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the objects to which an access control policy is attached.
//
// @param request - ListTargetAttachmentsForControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTargetAttachmentsForControlPolicyResponse
func (client *Client) ListTargetAttachmentsForControlPolicyWithContext(ctx context.Context, request *ListTargetAttachmentsForControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTargetAttachmentsForControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trusted services that are enabled within a management account or delegated administrator account.
//
// Description:
//
// Only a management account or delegated administrator account can be used to call this operation.
//
// @param request - ListTrustedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrustedServiceStatusResponse
func (client *Client) ListTrustedServiceStatusWithContext(ctx context.Context, request *ListTrustedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *ListTrustedServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminAccountId) {
		query["AdminAccountId"] = request.AdminAccountId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrustedServiceStatus"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a member from a folder to another.
//
// @param request - MoveAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveAccountResponse
func (client *Client) MoveAccountWithContext(ctx context.Context, request *MoveAccountRequest, runtime *dara.RuntimeOptions) (_result *MoveAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DestinationFolderId) {
		query["DestinationFolderId"] = request.DestinationFolderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a management account or member can be used as a main financial account.
//
// @param request - PrecheckForConsolidatedBillingAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckForConsolidatedBillingAccountResponse
func (client *Client) PrecheckForConsolidatedBillingAccountWithContext(ctx context.Context, request *PrecheckForConsolidatedBillingAccountRequest, runtime *dara.RuntimeOptions) (_result *PrecheckForConsolidatedBillingAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingAccountId) {
		query["BillingAccountId"] = request.BillingAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrecheckForConsolidatedBillingAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrecheckForConsolidatedBillingAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a member in a resource directory as a delegated administrator account of a trusted service.
//
// Description:
//
// The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory. When you call this operation, you must take note of the following limits:
//
//   - Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
//   - Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
//
//   - The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
//
// @param request - RegisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDelegatedAdministratorResponse
func (client *Client) RegisterDelegatedAdministratorWithContext(ctx context.Context, request *RegisterDelegatedAdministratorRequest, runtime *dara.RuntimeOptions) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.ServicePrincipal) {
		query["ServicePrincipal"] = request.ServicePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDelegatedAdministrator"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member of the cloud account type.
//
// @param request - RemoveCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveCloudAccountResponse
func (client *Client) RemoveCloudAccountWithContext(ctx context.Context, request *RemoveCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *RemoveCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveCloudAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends a verification email for the email address change of a member.
//
// @param request - RetryChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryChangeAccountEmailResponse
func (client *Client) RetryChangeAccountEmailWithContext(ctx context.Context, request *RetryChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *RetryChangeAccountEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryChangeAccountEmail"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends verification information to the email address of a contact.
//
// @param request - SendEmailVerificationForMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendEmailVerificationForMessageContactResponse
func (client *Client) SendEmailVerificationForMessageContactWithContext(ctx context.Context, request *SendEmailVerificationForMessageContactRequest, runtime *dara.RuntimeOptions) (_result *SendEmailVerificationForMessageContactResponse, _err error) {
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

	if !dara.IsNil(request.EmailAddress) {
		query["EmailAddress"] = request.EmailAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendEmailVerificationForMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendEmailVerificationForMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends verification information to the mobile phone number of a contact.
//
// @param request - SendPhoneVerificationForMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendPhoneVerificationForMessageContactResponse
func (client *Client) SendPhoneVerificationForMessageContactWithContext(ctx context.Context, request *SendPhoneVerificationForMessageContactRequest, runtime *dara.RuntimeOptions) (_result *SendPhoneVerificationForMessageContactResponse, _err error) {
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

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendPhoneVerificationForMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendPhoneVerificationForMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number that you want to bind to a member of the resource account type in a resource directory for security purposes.
//
// Description:
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// @param request - SendVerificationCodeForBindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForBindSecureMobilePhoneResponse
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithContext(ctx context.Context, request *SendVerificationCodeForBindSecureMobilePhoneRequest, runtime *dara.RuntimeOptions) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.SecureMobilePhone) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendVerificationCodeForBindSecureMobilePhone"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to the mobile phone number bound to a newly created account when you use the account to enable a resource directory.
//
// Description:
//
// Each Alibaba Cloud account can be used to send a maximum of 100 verification codes per day.
//
// @param request - SendVerificationCodeForEnableRDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForEnableRDResponse
func (client *Client) SendVerificationCodeForEnableRDWithContext(ctx context.Context, request *SendVerificationCodeForEnableRDRequest, runtime *dara.RuntimeOptions) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecureMobilePhone) {
		query["SecureMobilePhone"] = request.SecureMobilePhone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendVerificationCodeForEnableRD"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the member deletion feature.
//
// Description:
//
// Members of the resource account type can be deleted only after the member deletion feature is enabled.
//
// @param request - SetMemberDeletionPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMemberDeletionPermissionResponse
func (client *Client) SetMemberDeletionPermissionWithContext(ctx context.Context, request *SetMemberDeletionPermissionRequest, runtime *dara.RuntimeOptions) (_result *SetMemberDeletionPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetMemberDeletionPermission"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the Member Display Name Synchronization feature.
//
// @param request - SetMemberDisplayNameSyncStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMemberDisplayNameSyncStatusResponse
func (client *Client) SetMemberDisplayNameSyncStatusWithContext(ctx context.Context, request *SetMemberDisplayNameSyncStatusRequest, runtime *dara.RuntimeOptions) (_result *SetMemberDisplayNameSyncStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetMemberDisplayNameSyncStatus"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetMemberDisplayNameSyncStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to the members in a resource directory.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from the members in a resource directory.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the display name of a member, or switches the type of a member.
//
// Description:
//
//	  To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
//
//		- Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
//		- Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// @param request - UpdateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountResponse
func (client *Client) UpdateAccountWithContext(ctx context.Context, request *UpdateAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.NewAccountType) {
		query["NewAccountType"] = request.NewAccountType
	}

	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom access control policy.
//
// @param request - UpdateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPolicyResponse
func (client *Client) UpdateControlPolicyWithContext(ctx context.Context, request *UpdateControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewPolicyDocument) {
		query["NewPolicyDocument"] = request.NewPolicyDocument
	}

	if !dara.IsNil(request.NewPolicyName) {
		query["NewPolicyName"] = request.NewPolicyName
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateControlPolicy"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a folder.
//
// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.NewFolderName) {
		query["NewFolderName"] = request.NewFolderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a contact.
//
// @param request - UpdateMessageContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageContactResponse
func (client *Client) UpdateMessageContactWithContext(ctx context.Context, request *UpdateMessageContactRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageContactResponse, _err error) {
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

	if !dara.IsNil(request.EmailAddress) {
		query["EmailAddress"] = request.EmailAddress
	}

	if !dara.IsNil(request.MessageTypes) {
		query["MessageTypes"] = request.MessageTypes
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageContact"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the billing account of a member.
//
// @param request - UpdatePayerForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePayerForAccountResponse
func (client *Client) UpdatePayerForAccountWithContext(ctx context.Context, request *UpdatePayerForAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdatePayerForAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.PayerAccountId) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePayerForAccount"),
		Version:     dara.String("2022-04-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePayerForAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
