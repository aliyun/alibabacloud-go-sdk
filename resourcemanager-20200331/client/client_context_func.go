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
// This topic provides an example on how to call the API operation to accept the invitation `h-Ih8IuPfvV0t0****` that is initiated to invite the Alibaba Cloud account `177242285274****` to join the resource directory `rd-3G****`.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// After you attach an access control policy, the operations performed on resources by using members are limited by the policy. Make sure that the attached policy meets your expectations. Otherwise, your business may be affected.
//
// By default, the system access control policy FullAliyunAccess is attached to each folder and member.
//
// The access control policy that is attached to a folder also applies to all its subfolders and all members in the subfolders.
//
// A maximum of 10 access control policies can be attached to a folder or member.
//
// This topic provides an example on how to call the API operation to attach the custom access control policy `cp-jExXAqIYkwHN****` to the folder `fd-ZDNPiT****`.
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
		Version:     dara.String("2020-03-31"),
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
// Attaches a permission policy to an object, which can be a RAM user, RAM user group, or RAM role. After you attach a permission policy to an object, the object has the operation permissions on the resources in a specific resource group or within a specific Alibaba Cloud account.
//
// Description:
//
// In this example, the policy `AdministratorAccess` is attached to the RAM user `alice@demo.onaliyun.com` and takes effect only for resources in the `rg-9gLOoK****` resource group.
//
// @param request - AttachPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicyResponse
func (client *Client) AttachPolicyWithContext(ctx context.Context, request *AttachPolicyRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PrincipalName) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置安全手机号
//
// Description:
//
// You can call this API operation only to bind a mobile phone number to a member of the resource account type. You cannot call this API operation to change the mobile phone number that is bound to a member of the resource account type.
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this API operation.
//
// This topic provides an example on how to call the API operation to bind a mobile phone number to the member `138660628348****` for security purposes.
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
		Version:     dara.String("2020-03-31"),
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
// 取消修改邮箱
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
		Version:     dara.String("2020-03-31"),
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
// 取消创建云账号类型的成员
//
// @param request - CancelCreateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCreateCloudAccountResponse
func (client *Client) CancelCreateCloudAccountWithContext(ctx context.Context, request *CancelCreateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *CancelCreateCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCreateCloudAccount"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCreateCloudAccountResponse{}
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
// Description:
//
// This topic provides an example on how to call the API operation to cancel the invitation whose ID is `h-ycm4rp****`.
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
		Version:     dara.String("2020-03-31"),
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
// 取消升级资源账号
//
// @param request - CancelPromoteResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPromoteResourceAccountResponse
func (client *Client) CancelPromoteResourceAccountWithContext(ctx context.Context, request *CancelPromoteResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *CancelPromoteResourceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPromoteResourceAccount"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPromoteResourceAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成员账号设置安全邮箱
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
		Version:     dara.String("2020-03-31"),
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
// This topic provides an example on how to call the API operation to perform a deletion check on the member whose ID is `179855839641****`.
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
		Version:     dara.String("2020-03-31"),
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
// Creates a transfer rule. Custom transfer rules and transfer rules for associated resources are supported.
//
// Description:
//
// You can create up to 10 custom transfer rules. Each custom transfer rule can contain up to 10 content records.
//
// @param request - CreateAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoGroupingRuleResponse
func (client *Client) CreateAutoGroupingRuleWithContext(ctx context.Context, request *CreateAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoGroupingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		query["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		query["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		query["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceTypesScope) {
		query["ExcludeResourceTypesScope"] = request.ExcludeResourceTypesScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		query["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		query["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		query["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceTypesScope) {
		query["ResourceTypesScope"] = request.ResourceTypesScope
	}

	if !dara.IsNil(request.RuleContents) {
		query["RuleContents"] = request.RuleContents
	}

	if !dara.IsNil(request.RuleDesc) {
		query["RuleDesc"] = request.RuleDesc
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoGroupingRule"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoGroupingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建云账号
//
// Description:
//
// A resource directory supports two types of member accounts: resource accounts and cloud accounts.
//
//   - Resource account (recommended): A resource account is only used as a resource container and fully depends on a resource directory. Such member accounts are secure and easy-to-create. For more information about how to create a resource account, see [CreateResourceAccount](https://help.aliyun.com/document_detail/159392.html).
//
// >  A resource account can be upgraded to a cloud account. For more information, see [PromoteResourceAccount](https://help.aliyun.com/document_detail/159395.html) .
//
//   - Cloud account: A cloud account has all the features of an Alibaba Cloud account, including root permissions.
//
// @param request - CreateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudAccountResponse
func (client *Client) CreateCloudAccountWithContext(ctx context.Context, request *CreateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.ParentFolderId) {
		query["ParentFolderId"] = request.ParentFolderId
	}

	if !dara.IsNil(request.PayerAccountId) {
		query["PayerAccountId"] = request.PayerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudAccount"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// This topic provides an example on how to call the API operation to create a custom access control policy named `ExampleControlPolicy`. This access control policy is used to prohibit modifications to the ResourceDirectoryAccountAccessRole role and the permissions of the role.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateControlPolicy"),
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// >  A maximum of five levels of folders can be created under the root folder.
//
// In this example, a folder named `rdFolder` is created under the root folder.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2020-03-31"),
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
// Creates a permission policy.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyDocument) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version for a permission policy.
//
// @param request - CreatePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyVersionResponse
func (client *Client) CreatePolicyVersionWithContext(ctx context.Context, request *CreatePolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyDocument) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.SetAsDefault) {
		query["SetAsDefault"] = request.SetAsDefault
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyVersion"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyVersionResponse{}
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
		Version:     dara.String("2020-03-31"),
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
// Creates a resource group.
//
// Description:
//
// > A maximum of 30 resource groups can be created within an Alibaba Cloud account.
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithContext(ctx context.Context, request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a RAM role.
//
// @param request - CreateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssumeRolePolicyDocument) {
		query["AssumeRolePolicyDocument"] = request.AssumeRolePolicyDocument
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MaxSessionDuration) {
		query["MaxSessionDuration"] = request.MaxSessionDuration
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithContext(ctx context.Context, request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomSuffix) {
		query["CustomSuffix"] = request.CustomSuffix
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
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
		Version:     dara.String("2020-03-31"),
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
// 账号一键删除
//
// Description:
//
// The ID of the member that you want to delete.
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
		Version:     dara.String("2020-03-31"),
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
// Deletes a transfer rule.
//
// @param request - DeleteAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoGroupingRuleResponse
func (client *Client) DeleteAutoGroupingRuleWithContext(ctx context.Context, request *DeleteAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoGroupingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoGroupingRule"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoGroupingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除管控策略
//
// Description:
//
// If you want to delete a custom control policy that is attached to folders or member accounts, you must call the [DetachControlPolicy](https://help.aliyun.com/document_detail/208331.html) operation to detach the policy before you delete it.
//
// In this example, the custom control policy `cp-SImPt8GCEwiq****` is deleted.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// >  Before you delete a folder, make sure that the folder does not contain any member accounts or child folders.
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
		Version:     dara.String("2020-03-31"),
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
// Deletes a permission policy.
//
// Description:
//
// >
//
//   - Before you delete a permission policy, you must delete its all non-default versions. For information about how to delete a policy version, see [DeletePolicyVersion](https://help.aliyun.com/document_detail/159041.html).
//
//   - Before you delete a permission policy, you must make sure that the policy is not attached to a RAM user, a RAM user group, or a RAM role. For information about how to detach a policy, see [DetachPolicy](https://help.aliyun.com/document_detail/159168.html).
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a version of a permission policy.
//
// Description:
//
// >  The default version of a policy cannot be deleted.
//
// @param request - DeletePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyVersionResponse
func (client *Client) DeletePolicyVersionWithContext(ctx context.Context, request *DeletePolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyVersion"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// Description:
//
// >  Before you delete a resource group, you must delete all the resources in it.
//
// In this example, the resource group whose ID is `rg-9gLOoK****` is deleted.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a RAM role.
//
// @param request - DeleteRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service-linked role.
//
// @param request - DeleteServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceLinkedRoleResponse
func (client *Client) DeleteServiceLinkedRoleWithContext(ctx context.Context, request *DeleteServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceLinkedRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注销代理管理员
//
// Description:
//
// >  If the delegated administrator account that you want to remove has historical management tasks in the related trusted service, the trusted service may be affected after the delegated administrator account is removed. Therefore, proceed with caution.
//
// This topic provides an example on how to call the API operation to remove the delegated administrator account `181761095690****` for Cloud Firewall.
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
		Version:     dara.String("2020-03-31"),
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
// 解绑管控策略
//
// Description:
//
// After you detach an access control policy, the operations performed on resources by using members are not limited by the policy. Make sure that the detached policy meets your expectations. Otherwise, your business may be affected.
//
// Both the system and custom access control policies can be detached. If an object has only one access control policy attached, the policy cannot be detached.
//
// This topic provides an example on how to call the API operation to detach the custom control policy `cp-jExXAqIYkwHN****` from the folder `fd-ZDNPiT****`.
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
		Version:     dara.String("2020-03-31"),
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
// Detaches a permission policy from an object. After you detach a policy from an object, the object does not have the operation permissions on the current resource group or the resources within the current account.
//
// @param request - DetachPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicyResponse
func (client *Client) DetachPolicyWithContext(ctx context.Context, request *DetachPolicyRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PrincipalName) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启RD
//
// Description:
//
// You can use the current account or a newly created account to enable a resource directory. For more information, see [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// In this example, the current account is used to enable a resource directory.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// This topic provides an example on how to call the API operation to query the information of the member whose Alibaba Cloud account ID is `181761095690****`.
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
		Version:     dara.String("2020-03-31"),
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
// After you call the [CheckAccountDelete](https://help.aliyun.com/document_detail/448542.html) operation to perform a member deletion check, you can call the GetAccountDeletionCheckResult operation to query the check result. If the check result shows that the member meets deletion requirements, you can delete the member. Otherwise, you need to first modify the items that do not meet requirements.
//
// This topic provides an example on how to call the API operation to query the result of the deletion check for the member whose ID is `179855839641****`. The response shows that the member does not meet deletion requirements.
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
		Version:     dara.String("2020-03-31"),
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
// 获取账号删除状态
//
// Description:
//
// This topic provides an example on how to call the API operation to query the deletion status of the member whose Alibaba Cloud account ID is `169946124551****`. The response shows that the member is deleted.
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
		Version:     dara.String("2020-03-31"),
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
// Queries the information about a transfer rule.
//
// @param request - GetAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoGroupingRuleResponse
func (client *Client) GetAutoGroupingRuleWithContext(ctx context.Context, request *GetAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAutoGroupingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoGroupingRule"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoGroupingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// This topic provides an example on how to call the API operation to query the details of the access control policy whose ID is `cp-SImPt8GCEwiq****`.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// In this example, the information of the folder `fd-Jyl5U7****` is queried.
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
		Version:     dara.String("2020-03-31"),
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
// Description:
//
// In this example, the information of the invitation whose ID is `h-ycm4rp****` is queried.
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
		Version:     dara.String("2020-03-31"),
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
		Version:     dara.String("2020-03-31"),
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
// Queries the information about a permission policy.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithContext(ctx context.Context, request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a version of a permission policy.
//
// @param request - GetPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyVersionResponse
func (client *Client) GetPolicyVersionWithContext(ctx context.Context, request *GetPolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyVersion"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a resource group.
//
// @param request - GetResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithContext(ctx context.Context, request *GetResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroup"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of resources in a visible resource group.
//
// @param request - GetResourceGroupResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResourceCountsResponse
func (client *Client) GetResourceGroupResourceCountsWithContext(ctx context.Context, request *GetResourceGroupResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResourceCountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupByKey) {
		query["GroupByKey"] = request.GroupByKey
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupResourceCounts"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupResourceCountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a RAM role.
//
// @param request - GetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
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

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the task that is used to delete a service-linked role.
//
// @param request - GetServiceLinkedRoleDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLinkedRoleDeletionStatusResponse
func (client *Client) GetServiceLinkedRoleDeletionStatusWithContext(ctx context.Context, request *GetServiceLinkedRoleDeletionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionTaskId) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceLinkedRoleDeletionStatus"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceLinkedRoleDeletionStatusResponse{}
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
// Description:
//
// This topic provides an example on how to call the API operation to invite the account `someone@example.com` to join a resource directory.
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
		Version:     dara.String("2020-03-31"),
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
// Queries all the members in a resource directory.
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccounts"),
		Version:     dara.String("2020-03-31"),
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
		Version:     dara.String("2020-03-31"),
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
		Version:     dara.String("2020-03-31"),
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
// 列出自动分组修正记录
//
// @param request - ListAutoGroupingRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoGroupingRemediationsResponse
func (client *Client) ListAutoGroupingRemediationsWithContext(ctx context.Context, request *ListAutoGroupingRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListAutoGroupingRemediationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EarliestRemediationTime) {
		query["EarliestRemediationTime"] = request.EarliestRemediationTime
	}

	if !dara.IsNil(request.LatestRemediationTime) {
		query["LatestRemediationTime"] = request.LatestRemediationTime
	}

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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.TargetResourceGroupId) {
		query["TargetResourceGroupId"] = request.TargetResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoGroupingRemediations"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoGroupingRemediationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of transfer rules.
//
// @param request - ListAutoGroupingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoGroupingRulesResponse
func (client *Client) ListAutoGroupingRulesWithContext(ctx context.Context, request *ListAutoGroupingRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAutoGroupingRulesResponse, _err error) {
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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoGroupingRules"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoGroupingRulesResponse{}
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
// Description:
//
// This topic provides an example on how to call the API operation to query the system access control policies within a resource directory. The response shows that the resource directory has only one system access control policy. The policy is named `FullAliyunAccess`.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListControlPolicies"),
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// This topic provides an example on how to call the API operation to query the access control policies that are attached to the folder `fd-ZDNPiT****`.
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
		Version:     dara.String("2020-03-31"),
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
// 列出所有的代理管理员
//
// Description:
//
// This topic provides an example on how to call the API operation to query all delegated administrator accounts in a resource directory. The response shows that two delegated administrator accounts for Cloud Firewall exist in the resource directory.
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
		Version:     dara.String("2020-03-31"),
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
// 查看指定账号被设置为哪些可信服务的委派管理员
//
// Description:
//
// This topic provides an example on how to call the API operation to query the trusted services for which the member `138660628348****` is specified as a delegated administrator account. The response shows that the member is specified as a delegated administrator account of Cloud Firewall.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDelegatedServicesForAccount"),
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// >  You can view the information of only the first-level subfolders of a folder.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFoldersForParent"),
		Version:     dara.String("2020-03-31"),
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
// Description:
//
// This topic provides an example on how to call the API operation to query the invitations that are associated with the management account `172841235500****`. The response shows that two invitations are associated with the management account.
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
		Version:     dara.String("2020-03-31"),
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
		Version:     dara.String("2020-03-31"),
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
// Queries a list of permission policies.
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries policy attachment records.
//
// Description:
//
// You can view the following information:
//
//   - Policy attachment records within an Alibaba Cloud account or a resource group
//
//   - Permission policies attached to RAM users, RAM user groups, or RAM roles
//
//   - RAM users, RAM user groups, or RAM roles to which permission policies are attached within an Alibaba Cloud account or a resource group
//
// @param request - ListPolicyAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyAttachmentsResponse
func (client *Client) ListPolicyAttachmentsWithContext(ctx context.Context, request *ListPolicyAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PrincipalName) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyAttachments"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of versions of a policy.
//
// @param request - ListPolicyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyVersionsResponse
func (client *Client) ListPolicyVersionsWithContext(ctx context.Context, request *ListPolicyVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyVersions"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// You can call this API operation to query all resource groups within the current account. You can also call this API operation to query a specific resource group based on the status, ID, identifier, or display name of the resource group.
//
// This topic provides an example on how to call the API operation to query the basic information about the resource groups `rg-1hSBH2****` and `rg-9gLOoK****` within the current account.
//
// @param request - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithContext(ctx context.Context, request *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出资源组与用户授权信息
//
// @param request - ListResourceGroupsWithAuthDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsWithAuthDetailsResponse
func (client *Client) ListResourceGroupsWithAuthDetailsWithContext(ctx context.Context, request *ListResourceGroupsWithAuthDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsWithAuthDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroupsWithAuthDetails"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsWithAuthDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources that can be accessed by the current account in resource groups.
//
// Description:
//
// >  You can use a RAM role that is not associated with a session policy to call this API operation.
//
// This topic provides an example on how to call the API operation to query resources that can be accessed by the current account in resource groups. The response shows that the current account can access only the Elastic Compute Service (ECS) instance `i-23v38****` in the resource group `rg-uPJpP****`.
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of RAM roles.
//
// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出所有的Tag key
//
// Description:
//
// This topic provides an example on how to call the API operation to query tag keys. The response shows that the custom tag key team exists.
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
		Version:     dara.String("2020-03-31"),
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
// Queries the tags that are added to resource groups or the members in a resource directory.
//
// Description:
//
// This topic provides an example on how to call the API operation to query the tags that are added to the resource group with an ID of `rg-aekz6bre2uq****`. The response shows that only the `k1:v1` tag is added to the resource group.
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
		Version:     dara.String("2020-03-31"),
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
// 列出所有的Tag values
//
// Description:
//
// This topic provides an example on how to call the API operation to query the tag values of the tag key k1. The response shows that the tag value of the tag key k1 is v1.
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
		Version:     dara.String("2020-03-31"),
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
// Queries the objects to which a specific control policy is attached.
//
// Description:
//
// In this example, the folders or member accounts to which the control policy `cp-jExXAqIYkwHN****` is attached are queried. The returned result shows that the control policy is attached to the folder `fd-ZDNPiT****`.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// >  Only an enterprise management account or delegated administrator account can be used to call this operation.
//
// In this example, the trusted services that are enabled within an enterprise management account are queried. The returned result shows that the trusted services Cloud Config and ActionTrail are enabled within the enterprise management account.
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
		Version:     dara.String("2020-03-31"),
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
// 移动账号
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
		Version:     dara.String("2020-03-31"),
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
// Moves resources from one resource group to another. You can move multiple resources that reside in different regions, are used by different Alibaba Cloud services, or belong to different resource groups.
//
// Description:
//
// For more information about Alibaba Cloud services whose resources can be moved between resource groups, see the **Supported by the API*	- column in [Alibaba Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
//
// In this example, two virtual private clouds (VPCs) `vpc-bp1sig0mjktx5ewx1****` and `vpc-bp1visxm225pv49dz****` that reside in different regions and belong to different resource groups are moved to the resource group `rg-aekzmeobk5w****`.
//
// @param request - MoveResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourcesResponse
func (client *Client) MoveResourcesWithContext(ctx context.Context, request *MoveResourcesRequest, runtime *dara.RuntimeOptions) (_result *MoveResourcesResponse, _err error) {
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

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResources"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级资源账号
//
// @param request - PromoteResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromoteResourceAccountResponse
func (client *Client) PromoteResourceAccountWithContext(ctx context.Context, request *PromoteResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *PromoteResourceAccountResponse, _err error) {
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
		Action:      dara.String("PromoteResourceAccount"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PromoteResourceAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// The delegated administrator account can be used to access the information of the resource directory and view the structure and members of the resource directory. The delegated administrator account can also be used to perform service-related management operations in the trusted service on behalf of the management account of the resource directory.
//
// When you call this operation, you must take note of the following limits:
//
//   - Only some trusted services support delegated administrator accounts. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
//   - Only the management account of a resource directory or an authorized RAM user or RAM role of the management account can be used to call this operation.
//
//   - The number of delegated administrator accounts that are allowed for a trusted service is defined by the trusted service.
//
// This topic provides an example on how to call the API operation to specify the member `181761095690****` as a delegated administrator account of Cloud Firewall.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// This topic provides an example on how to call the API operation to remove the member `177242285274****` from a resource directory.
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
		Version:     dara.String("2020-03-31"),
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
// 重新发送创建云账号的邮箱验证
//
// @param request - ResendCreateCloudAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendCreateCloudAccountEmailResponse
func (client *Client) ResendCreateCloudAccountEmailWithContext(ctx context.Context, request *ResendCreateCloudAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendCreateCloudAccountEmail"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendCreateCloudAccountEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新发送升级资源账号的邮箱验证
//
// @param request - ResendPromoteResourceAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendPromoteResourceAccountEmailResponse
func (client *Client) ResendPromoteResourceAccountEmailWithContext(ctx context.Context, request *ResendPromoteResourceAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendPromoteResourceAccountEmail"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendPromoteResourceAccountEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新发送确认邮件
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
		Version:     dara.String("2020-03-31"),
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
// 发送绑定安全手机验证码
//
// Description:
//
// To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
//
// In this example, a verification code is sent to the mobile phone number that you want to bind to the resource account `138660628348****`.
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
		Version:     dara.String("2020-03-31"),
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
// 发送开启资源目录短信
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
		Version:     dara.String("2020-03-31"),
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
// Sets a default version for a permission policy.
//
// @param request - SetDefaultPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultPolicyVersionResponse
func (client *Client) SetDefaultPolicyVersionWithContext(ctx context.Context, request *SetDefaultPolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultPolicyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultPolicyVersion"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultPolicyVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启或关闭成员删除许可
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
		Version:     dara.String("2020-03-31"),
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
// Adds tags to resource groups or the members in a resource directory.
//
// Description:
//
// This topic provides an example on how to call the API operation to add the tag `k1:v1` to the resource group with an ID of `rg-aekz6bre2uq****`.
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
		Version:     dara.String("2020-03-31"),
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
// Removes tags from resource groups or the members in a resource directory.
//
// Description:
//
// This topic provides an example on how to call the API operation to remove the tag whose tag key is `k1` from the resource group whose ID is `rg-aek2dpwyrfr****`.
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
//	  To ensure that the system can record the operators of management operations, you must use a RAM user or RAM role to which the AliyunResourceDirectoryFullAccess policy is attached within the management account of your resource directory to call this operation.
//
//		- Before you switch the type of a member from resource account to cloud account, make sure that specific conditions are met. For more information about the conditions, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
//		- Before you switch the type of a member from cloud account to resource account, make sure that specific conditions are met. For more information about the conditions, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// This example provides an example on how to call the API operation to change the display name of the member `12323344****` to `admin`.
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
		Version:     dara.String("2020-03-31"),
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
// Updates the settings of the Transfer Associated Resources feature.
//
// Description:
//
// For information about the resources that support the Transfer Associated Resources feature, see [Use the Transfer Associated Resources feature](https://help.aliyun.com/document_detail/606232.html).
//
// @param request - UpdateAssociatedTransferSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssociatedTransferSettingResponse
func (client *Client) UpdateAssociatedTransferSettingWithContext(ctx context.Context, request *UpdateAssociatedTransferSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateAssociatedTransferSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableExistingResourcesTransfer) {
		query["EnableExistingResourcesTransfer"] = request.EnableExistingResourcesTransfer
	}

	if !dara.IsNil(request.RuleSettings) {
		query["RuleSettings"] = request.RuleSettings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAssociatedTransferSetting"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAssociatedTransferSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of the Automatic Resource Transfer feature. You can update only the configuration of the Transfer Existing Associated Resources feature.
//
// @param request - UpdateAutoGroupingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoGroupingConfigResponse
func (client *Client) UpdateAutoGroupingConfigWithContext(ctx context.Context, request *UpdateAutoGroupingConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoGroupingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableExistingResourcesTransfer) {
		query["EnableExistingResourcesTransfer"] = request.EnableExistingResourcesTransfer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoGroupingConfig"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoGroupingConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a transfer rule.
//
// @param request - UpdateAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoGroupingRuleResponse
func (client *Client) UpdateAutoGroupingRuleWithContext(ctx context.Context, request *UpdateAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoGroupingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeRegionIdsScope) {
		query["ExcludeRegionIdsScope"] = request.ExcludeRegionIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceGroupIdsScope) {
		query["ExcludeResourceGroupIdsScope"] = request.ExcludeResourceGroupIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceIdsScope) {
		query["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !dara.IsNil(request.ExcludeResourceTypesScope) {
		query["ExcludeResourceTypesScope"] = request.ExcludeResourceTypesScope
	}

	if !dara.IsNil(request.RegionIdsScope) {
		query["RegionIdsScope"] = request.RegionIdsScope
	}

	if !dara.IsNil(request.ResourceGroupIdsScope) {
		query["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !dara.IsNil(request.ResourceIdsScope) {
		query["ResourceIdsScope"] = request.ResourceIdsScope
	}

	if !dara.IsNil(request.ResourceTypesScope) {
		query["ResourceTypesScope"] = request.ResourceTypesScope
	}

	if !dara.IsNil(request.RuleContents) {
		query["RuleContents"] = request.RuleContents
	}

	if !dara.IsNil(request.RuleDesc) {
		query["RuleDesc"] = request.RuleDesc
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoGroupingRule"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoGroupingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// In this example, the name of the access control policy whose ID is `cp-jExXAqIYkwHN****` is changed to `NewControlPolicy`.
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
		Version:     dara.String("2020-03-31"),
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
		Version:     dara.String("2020-03-31"),
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

// Description:
//
// In this example, the display name of the resource group `rg-9gLOoK****` is changed to `project`.
//
// @param request - UpdateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithContext(ctx context.Context, request *UpdateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroup"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源组管理员配置
//
// @param request - UpdateResourceGroupAdminSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupAdminSettingResponse
func (client *Client) UpdateResourceGroupAdminSettingWithContext(ctx context.Context, request *UpdateResourceGroupAdminSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupAdminSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatorAsAdmin) {
		query["CreatorAsAdmin"] = request.CreatorAsAdmin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroupAdminSetting"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupAdminSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Resource Access Management (RAM) role.
//
// Description:
//
// In this example, the description of the RAM role `ECSAdmin` is updated to `ECS administrator`.
//
// @param request - UpdateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithContext(ctx context.Context, request *UpdateRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewAssumeRolePolicyDocument) {
		query["NewAssumeRolePolicyDocument"] = request.NewAssumeRolePolicyDocument
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewMaxSessionDuration) {
		query["NewMaxSessionDuration"] = request.NewMaxSessionDuration
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRole"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
