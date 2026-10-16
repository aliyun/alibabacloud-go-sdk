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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("resourcemanager"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Accepts an invitation by the invited Alibaba Cloud account by calling AcceptHandshake.
//
// Description:
//
// After the invited Alibaba Cloud account successfully joins a resource directory, the account becomes a member of the resource directory. The resource directory uses the name of the invited Alibaba Cloud account as the display name and account name of the member by default.
//
// This topic provides an example on how the invited Alibaba Cloud account `177242285274****` accepts an invitation from the resource directory `rd-3G****`. The handshake ID is `h-Ih8IuPfvV0t0****`.
//
// @param request - AcceptHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptHandshakeResponse
func (client *Client) AcceptHandshakeWithOptions(request *AcceptHandshakeRequest, runtime *dara.RuntimeOptions) (_result *AcceptHandshakeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts an invitation by the invited Alibaba Cloud account by calling AcceptHandshake.
//
// Description:
//
// After the invited Alibaba Cloud account successfully joins a resource directory, the account becomes a member of the resource directory. The resource directory uses the name of the invited Alibaba Cloud account as the display name and account name of the member by default.
//
// This topic provides an example on how the invited Alibaba Cloud account `177242285274****` accepts an invitation from the resource directory `rd-3G****`. The handshake ID is `h-Ih8IuPfvV0t0****`.
//
// @param request - AcceptHandshakeRequest
//
// @return AcceptHandshakeResponse
func (client *Client) AcceptHandshake(request *AcceptHandshakeRequest) (_result *AcceptHandshakeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptHandshakeResponse{}
	_body, _err := client.AcceptHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a control policy to a specified target.
//
// Description:
//
// After the control policy is attached, the resource folder or member immediately becomes subject to the control policy. Make sure that the result of the attach operation is as expected to avoid affecting your business.
//
// The system attaches the system policy FullAliyunAccess to resource folders and members by default.
//
// A control policy takes effect on the entire node to which it is attached. A control policy attached to a parent resource folder also takes effect on its child resource folders and members.
//
// Each node (resource folder or member) can have a maximum of 10 control policies attached.
//
// This topic provides an example on how to attach the custom policy `cp-jExXAqIYkwHN****` to the resource folder `fd-ZDNPiT****`.
//
// @param request - AttachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachControlPolicyResponse
func (client *Client) AttachControlPolicyWithOptions(request *AttachControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *AttachControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a control policy to a specified target.
//
// Description:
//
// After the control policy is attached, the resource folder or member immediately becomes subject to the control policy. Make sure that the result of the attach operation is as expected to avoid affecting your business.
//
// The system attaches the system policy FullAliyunAccess to resource folders and members by default.
//
// A control policy takes effect on the entire node to which it is attached. A control policy attached to a parent resource folder also takes effect on its child resource folders and members.
//
// Each node (resource folder or member) can have a maximum of 10 control policies attached.
//
// This topic provides an example on how to attach the custom policy `cp-jExXAqIYkwHN****` to the resource folder `fd-ZDNPiT****`.
//
// @param request - AttachControlPolicyRequest
//
// @return AttachControlPolicyResponse
func (client *Client) AttachControlPolicy(request *AttachControlPolicyRequest) (_result *AttachControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachControlPolicyResponse{}
	_body, _err := client.AttachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes AttachPolicy to add an access policy to an authorization object (a Resource Access Management (RAM) user, RAM user group, or RAM role), which is authorization. After the authorization is complete, the authorized object obtains access permissions on resources within the specified resource group or Alibaba Cloud account.
//
// @param request - AttachPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicyResponse
func (client *Client) AttachPolicyWithOptions(request *AttachPolicyRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes AttachPolicy to add an access policy to an authorization object (a Resource Access Management (RAM) user, RAM user group, or RAM role), which is authorization. After the authorization is complete, the authorized object obtains access permissions on resources within the specified resource group or Alibaba Cloud account.
//
// @param request - AttachPolicyRequest
//
// @return AttachPolicyResponse
func (client *Client) AttachPolicy(request *AttachPolicyRequest) (_result *AttachPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachPolicyResponse{}
	_body, _err := client.AttachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a secure phone number for a member of the resource account type.
//
// Description:
//
// This API operation is used only to set a secure phone number for a resource account for the first time. It cannot be used to modify an existing secure phone number.
//
// To ensure that the system can record the specific author of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource folder management permission (AliyunResourceDirectoryFullAccess) to invoke this API operation.
//
// This topic provides an example of how to set a secure phone number for the member `138660628348****`.
//
// @param request - BindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSecureMobilePhoneResponse
func (client *Client) BindSecureMobilePhoneWithOptions(request *BindSecureMobilePhoneRequest, runtime *dara.RuntimeOptions) (_result *BindSecureMobilePhoneResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a secure phone number for a member of the resource account type.
//
// Description:
//
// This API operation is used only to set a secure phone number for a resource account for the first time. It cannot be used to modify an existing secure phone number.
//
// To ensure that the system can record the specific author of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource folder management permission (AliyunResourceDirectoryFullAccess) to invoke this API operation.
//
// This topic provides an example of how to set a secure phone number for the member `138660628348****`.
//
// @param request - BindSecureMobilePhoneRequest
//
// @return BindSecureMobilePhoneResponse
func (client *Client) BindSecureMobilePhone(request *BindSecureMobilePhoneRequest) (_result *BindSecureMobilePhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindSecureMobilePhoneResponse{}
	_body, _err := client.BindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the modification of a member account email address.
//
// @param request - CancelChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelChangeAccountEmailResponse
func (client *Client) CancelChangeAccountEmailWithOptions(request *CancelChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *CancelChangeAccountEmailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the modification of a member account email address.
//
// @param request - CancelChangeAccountEmailRequest
//
// @return CancelChangeAccountEmailResponse
func (client *Client) CancelChangeAccountEmail(request *CancelChangeAccountEmailRequest) (_result *CancelChangeAccountEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelChangeAccountEmailResponse{}
	_body, _err := client.CancelChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the creation of a member of the cloud account type in a resource directory.
//
// @param request - CancelCreateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCreateCloudAccountResponse
func (client *Client) CancelCreateCloudAccountWithOptions(request *CancelCreateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *CancelCreateCloudAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the creation of a member of the cloud account type in a resource directory.
//
// @param request - CancelCreateCloudAccountRequest
//
// @return CancelCreateCloudAccountResponse
func (client *Client) CancelCreateCloudAccount(request *CancelCreateCloudAccountRequest) (_result *CancelCreateCloudAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCreateCloudAccountResponse{}
	_body, _err := client.CancelCreateCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a handshake.
//
// Description:
//
// This topic provides an example on how to cancel the handshake whose ID is `h-ycm4rp****`.
//
// @param request - CancelHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelHandshakeResponse
func (client *Client) CancelHandshakeWithOptions(request *CancelHandshakeRequest, runtime *dara.RuntimeOptions) (_result *CancelHandshakeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a handshake.
//
// Description:
//
// This topic provides an example on how to cancel the handshake whose ID is `h-ycm4rp****`.
//
// @param request - CancelHandshakeRequest
//
// @return CancelHandshakeResponse
func (client *Client) CancelHandshake(request *CancelHandshakeRequest) (_result *CancelHandshakeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelHandshakeResponse{}
	_body, _err := client.CancelHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the upgrade of a resource directory resource account to an Alibaba Cloud account.
//
// @param request - CancelPromoteResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPromoteResourceAccountResponse
func (client *Client) CancelPromoteResourceAccountWithOptions(request *CancelPromoteResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *CancelPromoteResourceAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the upgrade of a resource directory resource account to an Alibaba Cloud account.
//
// @param request - CancelPromoteResourceAccountRequest
//
// @return CancelPromoteResourceAccountResponse
func (client *Client) CancelPromoteResourceAccount(request *CancelPromoteResourceAccountRequest) (_result *CancelPromoteResourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelPromoteResourceAccountResponse{}
	_body, _err := client.CancelPromoteResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the email address of a member account by calling ChangeAccountEmail.
//
// @param request - ChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAccountEmailResponse
func (client *Client) ChangeAccountEmailWithOptions(request *ChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ChangeAccountEmailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the email address of a member account by calling ChangeAccountEmail.
//
// @param request - ChangeAccountEmailRequest
//
// @return ChangeAccountEmailResponse
func (client *Client) ChangeAccountEmail(request *ChangeAccountEmailRequest) (_result *ChangeAccountEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeAccountEmailResponse{}
	_body, _err := client.ChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls CheckAccountDelete to perform a member account deletion check task.
//
// Description:
//
// Before deleting a member account, call this API to perform a pre-deletion condition check on the target member account.
//
// @param request - CheckAccountDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountDeleteResponse
func (client *Client) CheckAccountDeleteWithOptions(request *CheckAccountDeleteRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountDeleteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls CheckAccountDelete to perform a member account deletion check task.
//
// Description:
//
// Before deleting a member account, call this API to perform a pre-deletion condition check on the target member account.
//
// @param request - CheckAccountDeleteRequest
//
// @return CheckAccountDeleteResponse
func (client *Client) CheckAccountDelete(request *CheckAccountDeleteRequest) (_result *CheckAccountDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAccountDeleteResponse{}
	_body, _err := client.CheckAccountDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an automatic allocation rule by calling CreateAutoGroupingRule. Two types are supported: custom allocation rules and attached-resource allocation rules.
//
// Description:
//
// You can create up to 10 custom allocation rules, and each custom allocation rule supports up to 10 rule content entries.
//
// @param request - CreateAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoGroupingRuleResponse
func (client *Client) CreateAutoGroupingRuleWithOptions(request *CreateAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoGroupingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic allocation rule by calling CreateAutoGroupingRule. Two types are supported: custom allocation rules and attached-resource allocation rules.
//
// Description:
//
// You can create up to 10 custom allocation rules, and each custom allocation rule supports up to 10 rule content entries.
//
// @param request - CreateAutoGroupingRuleRequest
//
// @return CreateAutoGroupingRuleResponse
func (client *Client) CreateAutoGroupingRule(request *CreateAutoGroupingRuleRequest) (_result *CreateAutoGroupingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoGroupingRuleResponse{}
	_body, _err := client.CreateAutoGroupingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a member of the cloud account type in a resource directory.
//
// Description:
//
// A resource directory contains two types of members: resource accounts and cloud accounts.
//
// - Resource account (recommended): A resource account serves only as a resource container and fully depends on the resource directory. It is a more secure account type that is easier to create. To create a resource account type member, see [CreateResourceAccount](https://help.aliyun.com/document_detail/159392.html).
//
// > A resource account can also be upgraded to a cloud account. For more information, see [PromoteResourceAccount](https://help.aliyun.com/document_detail/159395.html).
//
// - Cloud account: A cloud account has all the features of an Alibaba Cloud account, including root logon access.
//
// @param request - CreateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudAccountResponse
func (client *Client) CreateCloudAccountWithOptions(request *CreateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a member of the cloud account type in a resource directory.
//
// Description:
//
// A resource directory contains two types of members: resource accounts and cloud accounts.
//
// - Resource account (recommended): A resource account serves only as a resource container and fully depends on the resource directory. It is a more secure account type that is easier to create. To create a resource account type member, see [CreateResourceAccount](https://help.aliyun.com/document_detail/159392.html).
//
// > A resource account can also be upgraded to a cloud account. For more information, see [PromoteResourceAccount](https://help.aliyun.com/document_detail/159395.html).
//
// - Cloud account: A cloud account has all the features of an Alibaba Cloud account, including root logon access.
//
// @param request - CreateCloudAccountRequest
//
// @return CreateCloudAccountResponse
func (client *Client) CreateCloudAccount(request *CreateCloudAccountRequest) (_result *CreateCloudAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.CreateCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls CreateControlPolicy to create a custom control policy.
//
// Description:
//
// This topic provides an example on how to create a control policy named `ExampleControlPolicy` that prohibits modifications to the ResourceDirectoryAccountAccessRole role and its permissions.
//
// @param request - CreateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateControlPolicyResponse
func (client *Client) CreateControlPolicyWithOptions(request *CreateControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls CreateControlPolicy to create a custom control policy.
//
// Description:
//
// This topic provides an example on how to create a control policy named `ExampleControlPolicy` that prohibits modifications to the ResourceDirectoryAccountAccessRole role and its permissions.
//
// @param request - CreateControlPolicyRequest
//
// @return CreateControlPolicyResponse
func (client *Client) CreateControlPolicy(request *CreateControlPolicyRequest) (_result *CreateControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateControlPolicyResponse{}
	_body, _err := client.CreateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource folder by calling CreateFolder.
//
// Description:
//
// > You can create a maximum of five levels of resource folders.
//
// This topic provides an example of how to create a resource folder named `rdFolder` under the Root folder.
//
// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource folder by calling CreateFolder.
//
// Description:
//
// > You can create a maximum of five levels of resource folders.
//
// This topic provides an example of how to create a resource folder named `rdFolder` under the Root folder.
//
// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access policy.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access policy.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a version of an access policy by calling CreatePolicyVersion.
//
// @param request - CreatePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyVersionResponse
func (client *Client) CreatePolicyVersionWithOptions(request *CreatePolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version of an access policy by calling CreatePolicyVersion.
//
// @param request - CreatePolicyVersionRequest
//
// @return CreatePolicyVersionResponse
func (client *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (_result *CreatePolicyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyVersionResponse{}
	_body, _err := client.CreatePolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a member of the resource account type by calling CreateResourceAccount.
//
// Description:
//
// In a resource directory, a member serves as a resource container and is a resource grouping unit. A member typically represents a project or application. Resources in each member are physically isolated from resources in other members.
//
// @param request - CreateResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceAccountResponse
func (client *Client) CreateResourceAccountWithOptions(request *CreateResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a member of the resource account type by calling CreateResourceAccount.
//
// Description:
//
// In a resource directory, a member serves as a resource container and is a resource grouping unit. A member typically represents a project or application. Resources in each member are physically isolated from resources in other members.
//
// @param request - CreateResourceAccountRequest
//
// @return CreateResourceAccountResponse
func (client *Client) CreateResourceAccount(request *CreateResourceAccountRequest) (_result *CreateResourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceAccountResponse{}
	_body, _err := client.CreateResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// Description:
//
// > You can create a maximum of 30 resource groups within an Alibaba Cloud account.
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// > You can create a maximum of 30 resource groups within an Alibaba Cloud account.
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a role by calling CreateRole.
//
// @param request - CreateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a role by calling CreateRole.
//
// @param request - CreateRoleRequest
//
// @return CreateRoleResponse
func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Declines a handshake invitation by the invited Alibaba Cloud account.
//
// @param request - DeclineHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeclineHandshakeResponse
func (client *Client) DeclineHandshakeWithOptions(request *DeclineHandshakeRequest, runtime *dara.RuntimeOptions) (_result *DeclineHandshakeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Declines a handshake invitation by the invited Alibaba Cloud account.
//
// @param request - DeclineHandshakeRequest
//
// @return DeclineHandshakeResponse
func (client *Client) DeclineHandshake(request *DeclineHandshakeRequest) (_result *DeclineHandshakeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeclineHandshakeResponse{}
	_body, _err := client.DeclineHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a member of the resource account type by calling DeleteAccount.
//
// Description:
//
// Before deleting a member, call [CheckAccountDelete](https://help.aliyun.com/document_detail/448542.html) and [GetAccountDeletionCheckResult](https://help.aliyun.com/document_detail/448775.html) to perform a pre-deletion check on the member. Only members that meet the deletion conditions can be successfully deleted by calling DeleteAccount.
//
// After you successfully submit a member deletion request, you can call [GetAccountDeletionStatus](https://help.aliyun.com/document_detail/449001.html) to query the deletion status. Once a member is deleted, all resources and data in the member are deleted. You can no longer log on to or use the member, and the data cannot be recovered. Proceed with caution. For more information about member deletion, see [Delete a member of the resource account type](https://help.aliyun.com/document_detail/446078.html).
//
// @param tmpReq - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(tmpReq *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a member of the resource account type by calling DeleteAccount.
//
// Description:
//
// Before deleting a member, call [CheckAccountDelete](https://help.aliyun.com/document_detail/448542.html) and [GetAccountDeletionCheckResult](https://help.aliyun.com/document_detail/448775.html) to perform a pre-deletion check on the member. Only members that meet the deletion conditions can be successfully deleted by calling DeleteAccount.
//
// After you successfully submit a member deletion request, you can call [GetAccountDeletionStatus](https://help.aliyun.com/document_detail/449001.html) to query the deletion status. Once a member is deleted, all resources and data in the member are deleted. You can no longer log on to or use the member, and the data cannot be recovered. Proceed with caution. For more information about member deletion, see [Delete a member of the resource account type](https://help.aliyun.com/document_detail/446078.html).
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an automatic group transfer rule.
//
// @param request - DeleteAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoGroupingRuleResponse
func (client *Client) DeleteAutoGroupingRuleWithOptions(request *DeleteAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoGroupingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic group transfer rule.
//
// @param request - DeleteAutoGroupingRuleRequest
//
// @return DeleteAutoGroupingRuleResponse
func (client *Client) DeleteAutoGroupingRule(request *DeleteAutoGroupingRuleRequest) (_result *DeleteAutoGroupingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoGroupingRuleResponse{}
	_body, _err := client.DeleteAutoGroupingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom control policy.
//
// Description:
//
// To delete a custom control policy that is attached to a resource folder or member accounts, invoke [DetachControlPolicy](https://help.aliyun.com/document_detail/208331.html) to detach the policy first.
//
// This topic provides an example on how to delete the custom control policy cp-SImPt8GCEwiq****.
//
// @param request - DeleteControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom control policy.
//
// Description:
//
// To delete a custom control policy that is attached to a resource folder or member accounts, invoke [DetachControlPolicy](https://help.aliyun.com/document_detail/208331.html) to detach the policy first.
//
// This topic provides an example on how to delete the custom control policy cp-SImPt8GCEwiq****.
//
// @param request - DeleteControlPolicyRequest
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a folder.
//
// Description:
//
// > Before deleting a folder, make sure that the folder does not contain any member accounts or subfolders.
//
// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// > Before deleting a folder, make sure that the folder does not contain any member accounts or subfolders.
//
// @param request - DeleteFolderRequest
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified access policy.
//
// Description:
//
// > - Before deleting an access policy, delete all non-default versions of the policy. For information about how to delete a policy version, see [DeletePolicyVersion](https://help.aliyun.com/document_detail/159041.html).
//
// > - Before deleting an access policy, make sure that the policy is not referenced, meaning it is not attached to any Resource Access Management (RAM) user, user group, or RAM role. For information about how to remove an authorization, see [DetachPolicy](https://help.aliyun.com/document_detail/159168.html).
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified access policy.
//
// Description:
//
// > - Before deleting an access policy, delete all non-default versions of the policy. For information about how to delete a policy version, see [DeletePolicyVersion](https://help.aliyun.com/document_detail/159041.html).
//
// > - Before deleting an access policy, make sure that the policy is not referenced, meaning it is not attached to any Resource Access Management (RAM) user, user group, or RAM role. For information about how to remove an authorization, see [DetachPolicy](https://help.aliyun.com/document_detail/159168.html).
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a version of an access policy.
//
// Description:
//
// > The default version of an access policy cannot be deleted.
//
// @param request - DeletePolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyVersionResponse
func (client *Client) DeletePolicyVersionWithOptions(request *DeletePolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a version of an access policy.
//
// Description:
//
// > The default version of an access policy cannot be deleted.
//
// @param request - DeletePolicyVersionRequest
//
// @return DeletePolicyVersionResponse
func (client *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (_result *DeletePolicyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyVersionResponse{}
	_body, _err := client.DeletePolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// Description:
//
// > Before you delete a resource group, delete all resources in the resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// > Before you delete a resource group, delete all resources in the resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a role.
//
// @param request - DeleteRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a role.
//
// @param request - DeleteRoleRequest
//
// @return DeleteRoleResponse
func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteServiceLinkedRoleWithOptions(request *DeleteServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceLinkedRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteServiceLinkedRoleResponse
func (client *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (_result *DeleteServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServiceLinkedRoleResponse{}
	_body, _err := client.DeleteServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeregisterDelegatedAdministrator operation to remove the delegated administrator account of a trusted service.
//
// Description:
//
// > If the delegated administrator account being removed has historical management tasks in the trusted service, the removal may affect the normal use of the trusted service. Carefully consider the potential impact before performing the removal.
//
// @param request - DeregisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterDelegatedAdministratorResponse
func (client *Client) DeregisterDelegatedAdministratorWithOptions(request *DeregisterDelegatedAdministratorRequest, runtime *dara.RuntimeOptions) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeregisterDelegatedAdministrator operation to remove the delegated administrator account of a trusted service.
//
// Description:
//
// > If the delegated administrator account being removed has historical management tasks in the trusted service, the removal may affect the normal use of the trusted service. Carefully consider the potential impact before performing the removal.
//
// @param request - DeregisterDelegatedAdministratorRequest
//
// @return DeregisterDelegatedAdministratorResponse
func (client *Client) DeregisterDelegatedAdministrator(request *DeregisterDelegatedAdministratorRequest) (_result *DeregisterDelegatedAdministratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeregisterDelegatedAdministratorResponse{}
	_body, _err := client.DeregisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Closes a resource directory by calling DestroyResourceDirectory. This operation is irreversible. Proceed with caution.
//
// Description:
//
// Before closing a resource directory, make sure that you have completed the following operations:
//
// - All members in the resource directory have been removed. For information about how to remove a member, see [RemoveCloudAccount](https://help.aliyun.com/document_detail/159431.html).
//
// - All folders except the Root folder in the resource directory have been deleted. For information about how to delete a folder, see [DeleteFolder](https://help.aliyun.com/document_detail/159432.html).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DestroyResourceDirectoryResponse
func (client *Client) DestroyResourceDirectoryWithOptions(runtime *dara.RuntimeOptions) (_result *DestroyResourceDirectoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DestroyResourceDirectory"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes a resource directory by calling DestroyResourceDirectory. This operation is irreversible. Proceed with caution.
//
// Description:
//
// Before closing a resource directory, make sure that you have completed the following operations:
//
// - All members in the resource directory have been removed. For information about how to remove a member, see [RemoveCloudAccount](https://help.aliyun.com/document_detail/159431.html).
//
// - All folders except the Root folder in the resource directory have been deleted. For information about how to delete a folder, see [DeleteFolder](https://help.aliyun.com/document_detail/159432.html).
//
// @return DestroyResourceDirectoryResponse
func (client *Client) DestroyResourceDirectory() (_result *DestroyResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DestroyResourceDirectoryResponse{}
	_body, _err := client.DestroyResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a control policy from a resource folder or member.
//
// Description:
//
// After the control policy is detached, the resource folder or member that was previously associated with the policy immediately loses the governance enforced by the policy. Make sure that the result of the detach operation is as expected to avoid impact on your business.
//
// Both system policies and custom control policies can be detached. However, the last control policy associated with a resource folder or member cannot be detached.
//
// This topic provides an example on how to detach the custom policy `cp-jExXAqIYkwHN****` from the resource folder `fd-ZDNPiT****`.
//
// @param request - DetachControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachControlPolicyResponse
func (client *Client) DetachControlPolicyWithOptions(request *DetachControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DetachControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a control policy from a resource folder or member.
//
// Description:
//
// After the control policy is detached, the resource folder or member that was previously associated with the policy immediately loses the governance enforced by the policy. Make sure that the result of the detach operation is as expected to avoid impact on your business.
//
// Both system policies and custom control policies can be detached. However, the last control policy associated with a resource folder or member cannot be detached.
//
// This topic provides an example on how to detach the custom policy `cp-jExXAqIYkwHN****` from the resource folder `fd-ZDNPiT****`.
//
// @param request - DetachControlPolicyRequest
//
// @return DetachControlPolicyResponse
func (client *Client) DetachControlPolicy(request *DetachControlPolicyRequest) (_result *DetachControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachControlPolicyResponse{}
	_body, _err := client.DetachControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches an access policy from a specified identity. After the policy is detached, the identity loses the permissions to manage resources in the specified resource group or Alibaba Cloud account.
//
// @param request - DetachPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicyResponse
func (client *Client) DetachPolicyWithOptions(request *DetachPolicyRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches an access policy from a specified identity. After the policy is detached, the identity loses the permissions to manage resources in the specified resource group or Alibaba Cloud account.
//
// @param request - DetachPolicyRequest
//
// @return DetachPolicyResponse
func (client *Client) DetachPolicy(request *DetachPolicyRequest) (_result *DetachPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachPolicyResponse{}
	_body, _err := client.DetachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the attached-resource allocation feature for associated resources.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAssociatedTransferResponse
func (client *Client) DisableAssociatedTransferWithOptions(runtime *dara.RuntimeOptions) (_result *DisableAssociatedTransferResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAssociatedTransfer"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAssociatedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the attached-resource allocation feature for associated resources.
//
// @return DisableAssociatedTransferResponse
func (client *Client) DisableAssociatedTransfer() (_result *DisableAssociatedTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAssociatedTransferResponse{}
	_body, _err := client.DisableAssociatedTransferWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables automatic resource group transfer. After this feature is disabled, the custom condition rules and attached-resource transfer rules that have been configured are deleted. The resource groups of existing resources are not affected. If needed, you can re-enable this feature 1 minute after disabling it.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoGroupingResponse
func (client *Client) DisableAutoGroupingWithOptions(runtime *dara.RuntimeOptions) (_result *DisableAutoGroupingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAutoGrouping"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAutoGroupingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables automatic resource group transfer. After this feature is disabled, the custom condition rules and attached-resource transfer rules that have been configured are deleted. The resource groups of existing resources are not affected. If needed, you can re-enable this feature 1 minute after disabling it.
//
// @return DisableAutoGroupingResponse
func (client *Client) DisableAutoGrouping() (_result *DisableAutoGroupingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAutoGroupingResponse{}
	_body, _err := client.DisableAutoGroupingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the control policy feature.
//
// Description:
//
// After you disable the control policy feature, all control policies that are attached to resource folders and member accounts are automatically detached. However, the control policies are not deleted. They cannot be attached to any target nodes.
//
// > Disabling the control policy feature affects the permissions of resource folders and member accounts in the entire resource directory. Proceed with caution.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableControlPolicyResponse
func (client *Client) DisableControlPolicyWithOptions(runtime *dara.RuntimeOptions) (_result *DisableControlPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableControlPolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the control policy feature.
//
// Description:
//
// After you disable the control policy feature, all control policies that are attached to resource folders and member accounts are automatically detached. However, the control policies are not deleted. They cannot be attached to any target nodes.
//
// > Disabling the control policy feature affects the permissions of resource folders and member accounts in the entire resource directory. Proceed with caution.
//
// @return DisableControlPolicyResponse
func (client *Client) DisableControlPolicy() (_result *DisableControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableControlPolicyResponse{}
	_body, _err := client.DisableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a resource group event notification that is currently enabled. If the notification is already disabled, the error code `UnsupportedOperation.RepeatDisableResourceGroupNotification` is returned. You can call `EnableResourceGroupNotification` to enable the notification first.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableResourceGroupNotificationResponse
func (client *Client) DisableResourceGroupNotificationWithOptions(runtime *dara.RuntimeOptions) (_result *DisableResourceGroupNotificationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableResourceGroupNotification"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableResourceGroupNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a resource group event notification that is currently enabled. If the notification is already disabled, the error code `UnsupportedOperation.RepeatDisableResourceGroupNotification` is returned. You can call `EnableResourceGroupNotification` to enable the notification first.
//
// @return DisableResourceGroupNotificationResponse
func (client *Client) DisableResourceGroupNotification() (_result *DisableResourceGroupNotificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableResourceGroupNotificationResponse{}
	_body, _err := client.DisableResourceGroupNotificationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the attached-resource allocation feature for resource group transfers.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAssociatedTransferResponse
func (client *Client) EnableAssociatedTransferWithOptions(runtime *dara.RuntimeOptions) (_result *EnableAssociatedTransferResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAssociatedTransfer"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAssociatedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the attached-resource allocation feature for resource group transfers.
//
// @return EnableAssociatedTransferResponse
func (client *Client) EnableAssociatedTransfer() (_result *EnableAssociatedTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAssociatedTransferResponse{}
	_body, _err := client.EnableAssociatedTransferWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables automatic grouping. After automatic grouping is enabled, you can create, update, delete, and query automatic grouping rules.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoGroupingResponse
func (client *Client) EnableAutoGroupingWithOptions(runtime *dara.RuntimeOptions) (_result *EnableAutoGroupingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAutoGrouping"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAutoGroupingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables automatic grouping. After automatic grouping is enabled, you can create, update, delete, and query automatic grouping rules.
//
// @return EnableAutoGroupingResponse
func (client *Client) EnableAutoGrouping() (_result *EnableAutoGroupingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAutoGroupingResponse{}
	_body, _err := client.EnableAutoGroupingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the control policy feature by calling EnableControlPolicy.
//
// Description:
//
// A resource directory control policy is an access control policy based on the resource structure (resource folders or member accounts). It allows you to centrally manage the access permissions border for resource access at each level of the resource directory and establish overall or localized access control principles for your enterprise. Control policies define only the access permissions border and do not actually grant permissions. You must also set permissions by using Resource Access Management (RAM) in a member account before the corresponding identity can have access permissions to resources. For more information, see [Overview of control policies](https://help.aliyun.com/document_detail/178671.html).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableControlPolicyResponse
func (client *Client) EnableControlPolicyWithOptions(runtime *dara.RuntimeOptions) (_result *EnableControlPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableControlPolicy"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the control policy feature by calling EnableControlPolicy.
//
// Description:
//
// A resource directory control policy is an access control policy based on the resource structure (resource folders or member accounts). It allows you to centrally manage the access permissions border for resource access at each level of the resource directory and establish overall or localized access control principles for your enterprise. Control policies define only the access permissions border and do not actually grant permissions. You must also set permissions by using Resource Access Management (RAM) in a member account before the corresponding identity can have access permissions to resources. For more information, see [Overview of control policies](https://help.aliyun.com/document_detail/178671.html).
//
// @return EnableControlPolicyResponse
func (client *Client) EnableControlPolicy() (_result *EnableControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableControlPolicyResponse{}
	_body, _err := client.EnableControlPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a resource directory.
//
// Description:
//
// You can enable a resource directory by using the current logon account or by creating a new management account. For more information, refer to [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// This topic provides an example of how to enable a resource directory by using the current logon account.
//
// @param request - EnableResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceDirectoryResponse
func (client *Client) EnableResourceDirectoryWithOptions(request *EnableResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *EnableResourceDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can enable a resource directory by using the current logon account or by creating a new management account. For more information, refer to [Enable a resource directory](https://help.aliyun.com/document_detail/111215.html).
//
// This topic provides an example of how to enable a resource directory by using the current logon account.
//
// @param request - EnableResourceDirectoryRequest
//
// @return EnableResourceDirectoryResponse
func (client *Client) EnableResourceDirectory(request *EnableResourceDirectoryRequest) (_result *EnableResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableResourceDirectoryResponse{}
	_body, _err := client.EnableResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables resource group event notifications by calling EnableResourceGroupNotification.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceGroupNotificationResponse
func (client *Client) EnableResourceGroupNotificationWithOptions(runtime *dara.RuntimeOptions) (_result *EnableResourceGroupNotificationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableResourceGroupNotification"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableResourceGroupNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables resource group event notifications by calling EnableResourceGroupNotification.
//
// @return EnableResourceGroupNotificationResponse
func (client *Client) EnableResourceGroupNotification() (_result *EnableResourceGroupNotificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableResourceGroupNotificationResponse{}
	_body, _err := client.EnableResourceGroupNotificationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls GetAccount to retrieve member account information.
//
// @param request - GetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountResponse
func (client *Client) GetAccountWithOptions(request *GetAccountRequest, runtime *dara.RuntimeOptions) (_result *GetAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls GetAccount to retrieve member account information.
//
// @param request - GetAccountRequest
//
// @return GetAccountResponse
func (client *Client) GetAccount(request *GetAccountRequest) (_result *GetAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountResponse{}
	_body, _err := client.GetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a member deletion check by calling GetAccountDeletionCheckResult.
//
// Description:
//
// After you call [CheckAccountDelete](https://help.aliyun.com/document_detail/448542.html) to perform a member deletion check task, you can call GetAccountDeletionCheckResult to query the result of the member deletion check. If the deletion check requirements are met, the member account can be deleted. Otherwise, manually handle the non-compliant items until all requirements are met.
//
// This topic provides an example on how to query the deletion check result of the member account `179855839641****` that you want to delete. The response shows that the member account does not meet the deletion requirements.
//
// @param request - GetAccountDeletionCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionCheckResultResponse
func (client *Client) GetAccountDeletionCheckResultWithOptions(request *GetAccountDeletionCheckResultRequest, runtime *dara.RuntimeOptions) (_result *GetAccountDeletionCheckResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a member deletion check by calling GetAccountDeletionCheckResult.
//
// Description:
//
// After you call [CheckAccountDelete](https://help.aliyun.com/document_detail/448542.html) to perform a member deletion check task, you can call GetAccountDeletionCheckResult to query the result of the member deletion check. If the deletion check requirements are met, the member account can be deleted. Otherwise, manually handle the non-compliant items until all requirements are met.
//
// This topic provides an example on how to query the deletion check result of the member account `179855839641****` that you want to delete. The response shows that the member account does not meet the deletion requirements.
//
// @param request - GetAccountDeletionCheckResultRequest
//
// @return GetAccountDeletionCheckResultResponse
func (client *Client) GetAccountDeletionCheckResult(request *GetAccountDeletionCheckResultRequest) (_result *GetAccountDeletionCheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountDeletionCheckResultResponse{}
	_body, _err := client.GetAccountDeletionCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a member account.
//
// Description:
//
// This topic provides an example on how to query the deletion status of the member account `169946124551****`. The response shows that the member account has been successfully deleted.
//
// @param request - GetAccountDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountDeletionStatusResponse
func (client *Client) GetAccountDeletionStatusWithOptions(request *GetAccountDeletionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetAccountDeletionStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deletion status of a member account.
//
// Description:
//
// This topic provides an example on how to query the deletion status of the member account `169946124551****`. The response shows that the member account has been successfully deleted.
//
// @param request - GetAccountDeletionStatusRequest
//
// @return GetAccountDeletionStatusResponse
func (client *Client) GetAccountDeletionStatus(request *GetAccountDeletionStatusRequest) (_result *GetAccountDeletionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountDeletionStatusResponse{}
	_body, _err := client.GetAccountDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an automatic resource group transfer rule.
//
// @param request - GetAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoGroupingRuleResponse
func (client *Client) GetAutoGroupingRuleWithOptions(request *GetAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAutoGroupingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an automatic resource group transfer rule.
//
// @param request - GetAutoGroupingRuleRequest
//
// @return GetAutoGroupingRuleResponse
func (client *Client) GetAutoGroupingRule(request *GetAutoGroupingRuleRequest) (_result *GetAutoGroupingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoGroupingRuleResponse{}
	_body, _err := client.GetAutoGroupingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of automatic resource group transfer settings by calling GetAutoGroupingStatus.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoGroupingStatusResponse
func (client *Client) GetAutoGroupingStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetAutoGroupingStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoGroupingStatus"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoGroupingStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of automatic resource group transfer settings by calling GetAutoGroupingStatus.
//
// @return GetAutoGroupingStatusResponse
func (client *Client) GetAutoGroupingStatus() (_result *GetAutoGroupingStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoGroupingStatusResponse{}
	_body, _err := client.GetAutoGroupingStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a control policy.
//
// Description:
//
// This topic provides an example on how to query the details of the control policy whose ID is `cp-SImPt8GCEwiq****`.
//
// @param request - GetControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetControlPolicyResponse
func (client *Client) GetControlPolicyWithOptions(request *GetControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a control policy.
//
// Description:
//
// This topic provides an example on how to query the details of the control policy whose ID is `cp-SImPt8GCEwiq****`.
//
// @param request - GetControlPolicyRequest
//
// @return GetControlPolicyResponse
func (client *Client) GetControlPolicy(request *GetControlPolicyRequest) (_result *GetControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetControlPolicyResponse{}
	_body, _err := client.GetControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the enabling status of the control policy feature.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetControlPolicyEnablementStatusResponse
func (client *Client) GetControlPolicyEnablementStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetControlPolicyEnablementStatus"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the enabling status of the control policy feature.
//
// @return GetControlPolicyEnablementStatusResponse
func (client *Client) GetControlPolicyEnablementStatus() (_result *GetControlPolicyEnablementStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetControlPolicyEnablementStatusResponse{}
	_body, _err := client.GetControlPolicyEnablementStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls GetFolder to retrieve the information about a folder.
//
// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls GetFolder to retrieve the information about a folder.
//
// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves handshake information by calling GetHandshake.
//
// Description:
//
// This topic provides an example on how to retrieve the information about the handshake whose ID is `h-ycm4rp****`.
//
// @param request - GetHandshakeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHandshakeResponse
func (client *Client) GetHandshakeWithOptions(request *GetHandshakeRequest, runtime *dara.RuntimeOptions) (_result *GetHandshakeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves handshake information by calling GetHandshake.
//
// Description:
//
// This topic provides an example on how to retrieve the information about the handshake whose ID is `h-ycm4rp****`.
//
// @param request - GetHandshakeRequest
//
// @return GetHandshakeResponse
func (client *Client) GetHandshake(request *GetHandshakeRequest) (_result *GetHandshakeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHandshakeResponse{}
	_body, _err := client.GetHandshakeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the billing account information of a member account by calling GetPayerForAccount.
//
// @param request - GetPayerForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPayerForAccountResponse
func (client *Client) GetPayerForAccountWithOptions(request *GetPayerForAccountRequest, runtime *dara.RuntimeOptions) (_result *GetPayerForAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the billing account information of a member account by calling GetPayerForAccount.
//
// @param request - GetPayerForAccountRequest
//
// @return GetPayerForAccountResponse
func (client *Client) GetPayerForAccount(request *GetPayerForAccountRequest) (_result *GetPayerForAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPayerForAccountResponse{}
	_body, _err := client.GetPayerForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about a specified access policy.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a specified access policy.
//
// @param request - GetPolicyRequest
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a version of an access policy.
//
// @param request - GetPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyVersionResponse
func (client *Client) GetPolicyVersionWithOptions(request *GetPolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a version of an access policy.
//
// @param request - GetPolicyVersionRequest
//
// @return GetPolicyVersionResponse
func (client *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (_result *GetPolicyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyVersionResponse{}
	_body, _err := client.GetPolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource directory information. When called by a management account, returns the details of the resource directory created by the management account. When called by a member, returns the details of the resource directory to which the member belongs.
//
// Description:
//
// This topic provides an example of how to use a management account to query the details of the resource directory created by the management account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceDirectoryResponse
func (client *Client) GetResourceDirectoryWithOptions(runtime *dara.RuntimeOptions) (_result *GetResourceDirectoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceDirectory"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource directory information. When called by a management account, returns the details of the resource directory created by the management account. When called by a member, returns the details of the resource directory to which the member belongs.
//
// Description:
//
// This topic provides an example of how to use a management account to query the details of the resource directory created by the management account.
//
// @return GetResourceDirectoryResponse
func (client *Client) GetResourceDirectory() (_result *GetResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceDirectoryResponse{}
	_body, _err := client.GetResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group.
//
// @param request - GetResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithOptions(request *GetResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group.
//
// @param request - GetResourceGroupRequest
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroup(request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the administrator settings of a resource group.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupAdminSettingResponse
func (client *Client) GetResourceGroupAdminSettingWithOptions(runtime *dara.RuntimeOptions) (_result *GetResourceGroupAdminSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupAdminSetting"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupAdminSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the administrator settings of a resource group.
//
// @return GetResourceGroupAdminSettingResponse
func (client *Client) GetResourceGroupAdminSetting() (_result *GetResourceGroupAdminSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceGroupAdminSettingResponse{}
	_body, _err := client.GetResourceGroupAdminSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the notification settings of a resource group by calling GetResourceGroupNotificationSetting.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupNotificationSettingResponse
func (client *Client) GetResourceGroupNotificationSettingWithOptions(runtime *dara.RuntimeOptions) (_result *GetResourceGroupNotificationSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupNotificationSetting"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupNotificationSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the notification settings of a resource group by calling GetResourceGroupNotificationSetting.
//
// @return GetResourceGroupNotificationSettingResponse
func (client *Client) GetResourceGroupNotificationSetting() (_result *GetResourceGroupNotificationSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceGroupNotificationSettingResponse{}
	_body, _err := client.GetResourceGroupNotificationSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource count in visible resource groups.
//
// @param request - GetResourceGroupResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResourceCountsResponse
func (client *Client) GetResourceGroupResourceCountsWithOptions(request *GetResourceGroupResourceCountsRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResourceCountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource count in visible resource groups.
//
// @param request - GetResourceGroupResourceCountsRequest
//
// @return GetResourceGroupResourceCountsResponse
func (client *Client) GetResourceGroupResourceCounts(request *GetResourceGroupResourceCountsRequest) (_result *GetResourceGroupResourceCountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceGroupResourceCountsResponse{}
	_body, _err := client.GetResourceGroupResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a role by calling GetRole.
//
// @param request - GetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithOptions(request *GetRoleRequest, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a role by calling GetRole.
//
// @param request - GetRoleRequest
//
// @return GetRoleResponse
func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a task that deletes a service-linked role.
//
// @param request - GetServiceLinkedRoleDeletionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLinkedRoleDeletionStatusResponse
func (client *Client) GetServiceLinkedRoleDeletionStatusWithOptions(request *GetServiceLinkedRoleDeletionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a task that deletes a service-linked role.
//
// @param request - GetServiceLinkedRoleDeletionStatusRequest
//
// @return GetServiceLinkedRoleDeletionStatusResponse
func (client *Client) GetServiceLinkedRoleDeletionStatus(request *GetServiceLinkedRoleDeletionStatusRequest) (_result *GetServiceLinkedRoleDeletionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceLinkedRoleDeletionStatusResponse{}
	_body, _err := client.GetServiceLinkedRoleDeletionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls InitResourceDirectory to enable a resource directory. After you enable a resource directory, the system creates a Root folder and sets the current account as the management account. The management account has all permissions to manage the resource directory.
//
// Description:
//
// > - Use an account that has passed enterprise real-name verification to enable a resource directory. An account that has passed only individual real-name verification cannot enable a resource directory.
//
// > - Use the management account only as the administrator of the resource directory. Do not purchase cloud service resources with the management account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitResourceDirectoryResponse
func (client *Client) InitResourceDirectoryWithOptions(runtime *dara.RuntimeOptions) (_result *InitResourceDirectoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("InitResourceDirectory"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls InitResourceDirectory to enable a resource directory. After you enable a resource directory, the system creates a Root folder and sets the current account as the management account. The management account has all permissions to manage the resource directory.
//
// Description:
//
// > - Use an account that has passed enterprise real-name verification to enable a resource directory. An account that has passed only individual real-name verification cannot enable a resource directory.
//
// > - Use the management account only as the administrator of the resource directory. Do not purchase cloud service resources with the management account.
//
// @return InitResourceDirectoryResponse
func (client *Client) InitResourceDirectory() (_result *InitResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitResourceDirectoryResponse{}
	_body, _err := client.InitResourceDirectoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invites an external account to join a resource directory.
//
// Description:
//
// This topic provides an example on how to invite the external account `someone@example.com` to join a resource directory.
//
// @param request - InviteAccountToResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteAccountToResourceDirectoryResponse
func (client *Client) InviteAccountToResourceDirectoryWithOptions(request *InviteAccountToResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites an external account to join a resource directory.
//
// Description:
//
// This topic provides an example on how to invite the external account `someone@example.com` to join a resource directory.
//
// @param request - InviteAccountToResourceDirectoryRequest
//
// @return InviteAccountToResourceDirectoryResponse
func (client *Client) InviteAccountToResourceDirectory(request *InviteAccountToResourceDirectoryRequest) (_result *InviteAccountToResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InviteAccountToResourceDirectoryResponse{}
	_body, _err := client.InviteAccountToResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of all member accounts in a resource directory by calling ListAccounts.
//
// Description:
//
// The management account of the resource directory and the delegated administrator accounts of trusted services can call this operation.
//
// @param request - ListAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsResponse
func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of all member accounts in a resource directory by calling ListAccounts.
//
// Description:
//
// The management account of the resource directory and the delegated administrator accounts of trusted services can call this operation.
//
// @param request - ListAccountsRequest
//
// @return ListAccountsResponse
func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of member accounts in a resource folder by calling ListAccountsForParent.
//
// @param request - ListAccountsForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountsForParentResponse
func (client *Client) ListAccountsForParentWithOptions(request *ListAccountsForParentRequest, runtime *dara.RuntimeOptions) (_result *ListAccountsForParentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of member accounts in a resource folder by calling ListAccountsForParent.
//
// @param request - ListAccountsForParentRequest
//
// @return ListAccountsForParentResponse
func (client *Client) ListAccountsForParent(request *ListAccountsForParentRequest) (_result *ListAccountsForParentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountsForParentResponse{}
	_body, _err := client.ListAccountsForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all parent folders of a specified folder. The results are displayed in order from top to bottom.
//
// @param request - ListAncestorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAncestorsResponse
func (client *Client) ListAncestorsWithOptions(request *ListAncestorsRequest, runtime *dara.RuntimeOptions) (_result *ListAncestorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all parent folders of a specified folder. The results are displayed in order from top to bottom.
//
// @param request - ListAncestorsRequest
//
// @return ListAncestorsResponse
func (client *Client) ListAncestors(request *ListAncestorsRequest) (_result *ListAncestorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAncestorsResponse{}
	_body, _err := client.ListAncestorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the settings of the associated resource group transfer feature.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssociatedTransferSettingResponse
func (client *Client) ListAssociatedTransferSettingWithOptions(runtime *dara.RuntimeOptions) (_result *ListAssociatedTransferSettingResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListAssociatedTransferSetting"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAssociatedTransferSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the settings of the associated resource group transfer feature.
//
// @return ListAssociatedTransferSettingResponse
func (client *Client) ListAssociatedTransferSetting() (_result *ListAssociatedTransferSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAssociatedTransferSettingResponse{}
	_body, _err := client.ListAssociatedTransferSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists auto-grouping remediation records.
//
// @param request - ListAutoGroupingRemediationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoGroupingRemediationsResponse
func (client *Client) ListAutoGroupingRemediationsWithOptions(request *ListAutoGroupingRemediationsRequest, runtime *dara.RuntimeOptions) (_result *ListAutoGroupingRemediationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists auto-grouping remediation records.
//
// @param request - ListAutoGroupingRemediationsRequest
//
// @return ListAutoGroupingRemediationsResponse
func (client *Client) ListAutoGroupingRemediations(request *ListAutoGroupingRemediationsRequest) (_result *ListAutoGroupingRemediationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoGroupingRemediationsResponse{}
	_body, _err := client.ListAutoGroupingRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of automatic group transfer rules.
//
// @param request - ListAutoGroupingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoGroupingRulesResponse
func (client *Client) ListAutoGroupingRulesWithOptions(request *ListAutoGroupingRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAutoGroupingRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of automatic group transfer rules.
//
// @param request - ListAutoGroupingRulesRequest
//
// @return ListAutoGroupingRulesResponse
func (client *Client) ListAutoGroupingRules(request *ListAutoGroupingRulesRequest) (_result *ListAutoGroupingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoGroupingRulesResponse{}
	_body, _err := client.ListAutoGroupingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of control policies.
//
// @param request - ListControlPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPoliciesResponse
func (client *Client) ListControlPoliciesWithOptions(request *ListControlPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListControlPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of control policies.
//
// @param request - ListControlPoliciesRequest
//
// @return ListControlPoliciesResponse
func (client *Client) ListControlPolicies(request *ListControlPoliciesRequest) (_result *ListControlPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListControlPoliciesResponse{}
	_body, _err := client.ListControlPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of control policies attached to a resource folder or member account.
//
// Description:
//
// This topic provides an example on how to query the list of control policies attached to the resource folder `fd-ZDNPiT****`.
//
// @param request - ListControlPolicyAttachmentsForTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListControlPolicyAttachmentsForTargetResponse
func (client *Client) ListControlPolicyAttachmentsForTargetWithOptions(request *ListControlPolicyAttachmentsForTargetRequest, runtime *dara.RuntimeOptions) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of control policies attached to a resource folder or member account.
//
// Description:
//
// This topic provides an example on how to query the list of control policies attached to the resource folder `fd-ZDNPiT****`.
//
// @param request - ListControlPolicyAttachmentsForTargetRequest
//
// @return ListControlPolicyAttachmentsForTargetResponse
func (client *Client) ListControlPolicyAttachmentsForTarget(request *ListControlPolicyAttachmentsForTargetRequest) (_result *ListControlPolicyAttachmentsForTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListControlPolicyAttachmentsForTargetResponse{}
	_body, _err := client.ListControlPolicyAttachmentsForTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of delegated administrator accounts for a trusted service.
//
// @param request - ListDelegatedAdministratorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedAdministratorsResponse
func (client *Client) ListDelegatedAdministratorsWithOptions(request *ListDelegatedAdministratorsRequest, runtime *dara.RuntimeOptions) (_result *ListDelegatedAdministratorsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of delegated administrator accounts for a trusted service.
//
// @param request - ListDelegatedAdministratorsRequest
//
// @return ListDelegatedAdministratorsResponse
func (client *Client) ListDelegatedAdministrators(request *ListDelegatedAdministratorsRequest) (_result *ListDelegatedAdministratorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDelegatedAdministratorsResponse{}
	_body, _err := client.ListDelegatedAdministratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trusted services for which a specified member is set as a delegated administrator account.
//
// @param request - ListDelegatedServicesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDelegatedServicesForAccountResponse
func (client *Client) ListDelegatedServicesForAccountWithOptions(request *ListDelegatedServicesForAccountRequest, runtime *dara.RuntimeOptions) (_result *ListDelegatedServicesForAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trusted services for which a specified member is set as a delegated administrator account.
//
// @param request - ListDelegatedServicesForAccountRequest
//
// @return ListDelegatedServicesForAccountResponse
func (client *Client) ListDelegatedServicesForAccount(request *ListDelegatedServicesForAccountRequest) (_result *ListDelegatedServicesForAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDelegatedServicesForAccountResponse{}
	_body, _err := client.ListDelegatedServicesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subfolders of a specified folder.
//
// Description:
//
// > You can query only the subfolders at the next level of the specified folder. You cannot query subfolders at deeper levels.
//
// @param request - ListFoldersForParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersForParentResponse
func (client *Client) ListFoldersForParentWithOptions(request *ListFoldersForParentRequest, runtime *dara.RuntimeOptions) (_result *ListFoldersForParentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subfolders of a specified folder.
//
// Description:
//
// > You can query only the subfolders at the next level of the specified folder. You cannot query subfolders at deeper levels.
//
// @param request - ListFoldersForParentRequest
//
// @return ListFoldersForParentResponse
func (client *Client) ListFoldersForParent(request *ListFoldersForParentRequest) (_result *ListFoldersForParentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFoldersForParentResponse{}
	_body, _err := client.ListFoldersForParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the invitation records under the current account.
//
// Description:
//
// This topic provides an example on how to query the invitation records of the current management account `172841235500****`.
//
// @param request - ListHandshakesForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForAccountResponse
func (client *Client) ListHandshakesForAccountWithOptions(request *ListHandshakesForAccountRequest, runtime *dara.RuntimeOptions) (_result *ListHandshakesForAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invitation records under the current account.
//
// Description:
//
// This topic provides an example on how to query the invitation records of the current management account `172841235500****`.
//
// @param request - ListHandshakesForAccountRequest
//
// @return ListHandshakesForAccountResponse
func (client *Client) ListHandshakesForAccount(request *ListHandshakesForAccountRequest) (_result *ListHandshakesForAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHandshakesForAccountResponse{}
	_body, _err := client.ListHandshakesForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all invitations in a resource directory.
//
// @param request - ListHandshakesForResourceDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHandshakesForResourceDirectoryResponse
func (client *Client) ListHandshakesForResourceDirectoryWithOptions(request *ListHandshakesForResourceDirectoryRequest, runtime *dara.RuntimeOptions) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all invitations in a resource directory.
//
// @param request - ListHandshakesForResourceDirectoryRequest
//
// @return ListHandshakesForResourceDirectoryResponse
func (client *Client) ListHandshakesForResourceDirectory(request *ListHandshakesForResourceDirectoryRequest) (_result *ListHandshakesForResourceDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHandshakesForResourceDirectoryResponse{}
	_body, _err := client.ListHandshakesForResourceDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ListPolicies to query a list of access policies.
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ListPolicies to query a list of access policies.
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of policy attachments by calling ListPolicyAttachments.
//
// Description:
//
// You can query the following information:
//
// - All authorization records under an Alibaba Cloud account or a resource group.
//
// - All access policies granted to a Resource Access Management (RAM) user, RAM user group, or RAM role.
//
// - RAM users, RAM user groups, and RAM roles that have a specific permission under an Alibaba Cloud account or a resource group.
//
// @param request - ListPolicyAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyAttachmentsResponse
func (client *Client) ListPolicyAttachmentsWithOptions(request *ListPolicyAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of policy attachments by calling ListPolicyAttachments.
//
// Description:
//
// You can query the following information:
//
// - All authorization records under an Alibaba Cloud account or a resource group.
//
// - All access policies granted to a Resource Access Management (RAM) user, RAM user group, or RAM role.
//
// - RAM users, RAM user groups, and RAM roles that have a specific permission under an Alibaba Cloud account or a resource group.
//
// @param request - ListPolicyAttachmentsRequest
//
// @return ListPolicyAttachmentsResponse
func (client *Client) ListPolicyAttachments(request *ListPolicyAttachmentsRequest) (_result *ListPolicyAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicyAttachmentsResponse{}
	_body, _err := client.ListPolicyAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of an access policy.
//
// @param request - ListPolicyVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyVersionsResponse
func (client *Client) ListPolicyVersionsWithOptions(request *ListPolicyVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of an access policy.
//
// @param request - ListPolicyVersionsRequest
//
// @return ListPolicyVersionsResponse
func (client *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (_result *ListPolicyVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicyVersionsResponse{}
	_body, _err := client.ListPolicyVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the capability items of resource groups.
//
// @param request - ListResourceGroupCapabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupCapabilityResponse
func (client *Client) ListResourceGroupCapabilityWithOptions(request *ListResourceGroupCapabilityRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupCapabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.SupportResourceGroupEvent) {
		query["SupportResourceGroupEvent"] = request.SupportResourceGroupEvent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroupCapability"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupCapabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the capability items of resource groups.
//
// @param request - ListResourceGroupCapabilityRequest
//
// @return ListResourceGroupCapabilityResponse
func (client *Client) ListResourceGroupCapability(request *ListResourceGroupCapabilityRequest) (_result *ListResourceGroupCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupCapabilityResponse{}
	_body, _err := client.ListResourceGroupCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of resource groups.
//
// Description:
//
// You can call this operation to query all resource groups under the current account. You can also filter resource groups by resource group status, resource group ID, resource group identifier, or resource group display name.
//
// @param request - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of resource groups.
//
// Description:
//
// You can call this operation to query all resource groups under the current account. You can also filter resource groups by resource group status, resource group ID, resource group identifier, or resource group display name.
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists resource groups along with user authorization details.
//
// @param request - ListResourceGroupsWithAuthDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsWithAuthDetailsResponse
func (client *Client) ListResourceGroupsWithAuthDetailsWithOptions(request *ListResourceGroupsWithAuthDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsWithAuthDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists resource groups along with user authorization details.
//
// @param request - ListResourceGroupsWithAuthDetailsRequest
//
// @return ListResourceGroupsWithAuthDetailsResponse
func (client *Client) ListResourceGroupsWithAuthDetails(request *ListResourceGroupsWithAuthDetailsRequest) (_result *ListResourceGroupsWithAuthDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupsWithAuthDetailsResponse{}
	_body, _err := client.ListResourceGroupsWithAuthDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources in the resource groups accessible to the current account.
//
// Description:
//
// > This API can be called by assuming a RAM role, but session policies for the role are not supported.
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in the resource groups accessible to the current account.
//
// Description:
//
// > This API can be called by assuming a RAM role, but session policies for the role are not supported.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of roles by calling ListRoles.
//
// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of roles by calling ListRoles.
//
// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag key list of resource directory member accounts by calling ListTagKeys.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag key list of resource directory member accounts by calling ListTagKeys.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a resource group or a resource directory member.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a resource group or a resource directory member.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of a specified tag key by calling ListTagValues.
//
// Description:
//
// ### Usage notes
//
// This topic provides an example on how to query the tag values of the tag key k1. The response shows that the tag value corresponding to the tag key k1 is v1.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of a specified tag key by calling ListTagValues.
//
// Description:
//
// ### Usage notes
//
// This topic provides an example on how to query the tag values of the tag key k1. The response shows that the tag value corresponding to the tag key k1 is v1.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the nodes to which a control policy is attached.
//
// @param request - ListTargetAttachmentsForControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTargetAttachmentsForControlPolicyResponse
func (client *Client) ListTargetAttachmentsForControlPolicyWithOptions(request *ListTargetAttachmentsForControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the nodes to which a control policy is attached.
//
// @param request - ListTargetAttachmentsForControlPolicyRequest
//
// @return ListTargetAttachmentsForControlPolicyResponse
func (client *Client) ListTargetAttachmentsForControlPolicy(request *ListTargetAttachmentsForControlPolicyRequest) (_result *ListTargetAttachmentsForControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTargetAttachmentsForControlPolicyResponse{}
	_body, _err := client.ListTargetAttachmentsForControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of trusted services enabled by the management account or a delegated administrator account.
//
// Description:
//
// > This operation can be called only by the management account or a delegated administrator account.
//
// @param request - ListTrustedServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrustedServiceStatusResponse
func (client *Client) ListTrustedServiceStatusWithOptions(request *ListTrustedServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *ListTrustedServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of trusted services enabled by the management account or a delegated administrator account.
//
// Description:
//
// > This operation can be called only by the management account or a delegated administrator account.
//
// @param request - ListTrustedServiceStatusRequest
//
// @return ListTrustedServiceStatusResponse
func (client *Client) ListTrustedServiceStatus(request *ListTrustedServiceStatusRequest) (_result *ListTrustedServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTrustedServiceStatusResponse{}
	_body, _err := client.ListTrustedServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource group events.
//
// @param request - LookupResourceGroupEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupResourceGroupEventsResponse
func (client *Client) LookupResourceGroupEventsWithOptions(request *LookupResourceGroupEventsRequest, runtime *dara.RuntimeOptions) (_result *LookupResourceGroupEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCategory) {
		query["EventCategory"] = request.EventCategory
	}

	if !dara.IsNil(request.LookupAttributes) {
		query["LookupAttributes"] = request.LookupAttributes
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupDisplayName) {
		query["ResourceGroupDisplayName"] = request.ResourceGroupDisplayName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupResourceGroupEvents"),
		Version:     dara.String("2020-03-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupResourceGroupEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource group events.
//
// @param request - LookupResourceGroupEventsRequest
//
// @return LookupResourceGroupEventsResponse
func (client *Client) LookupResourceGroupEvents(request *LookupResourceGroupEventsRequest) (_result *LookupResourceGroupEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LookupResourceGroupEventsResponse{}
	_body, _err := client.LookupResourceGroupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a member to another folder.
//
// @param request - MoveAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveAccountResponse
func (client *Client) MoveAccountWithOptions(request *MoveAccountRequest, runtime *dara.RuntimeOptions) (_result *MoveAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a member to another folder.
//
// @param request - MoveAccountRequest
//
// @return MoveAccountResponse
func (client *Client) MoveAccount(request *MoveAccountRequest) (_result *MoveAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveAccountResponse{}
	_body, _err := client.MoveAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls MoveResources to move resources from one resource group to another.
//
// Description:
//
// Supports batch transfer of resources across regions, cloud services, and resource groups.
//
// For cloud services that support resource group transfer, see the **API*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
//
// @param request - MoveResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourcesResponse
func (client *Client) MoveResourcesWithOptions(request *MoveResourcesRequest, runtime *dara.RuntimeOptions) (_result *MoveResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls MoveResources to move resources from one resource group to another.
//
// Description:
//
// Supports batch transfer of resources across regions, cloud services, and resource groups.
//
// For cloud services that support resource group transfer, see the **API*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
//
// @param request - MoveResourcesRequest
//
// @return MoveResourcesResponse
func (client *Client) MoveResources(request *MoveResourcesRequest) (_result *MoveResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourcesResponse{}
	_body, _err := client.MoveResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades a resource account to an Alibaba Cloud account. This is an asynchronous operation.
//
// @param request - PromoteResourceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromoteResourceAccountResponse
func (client *Client) PromoteResourceAccountWithOptions(request *PromoteResourceAccountRequest, runtime *dara.RuntimeOptions) (_result *PromoteResourceAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a resource account to an Alibaba Cloud account. This is an asynchronous operation.
//
// @param request - PromoteResourceAccountRequest
//
// @return PromoteResourceAccountResponse
func (client *Client) PromoteResourceAccount(request *PromoteResourceAccountRequest) (_result *PromoteResourceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PromoteResourceAccountResponse{}
	_body, _err := client.PromoteResourceAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a member account in a resource directory as a delegated administrator account for a trusted service.
//
// Description:
//
// After the settings are complete, the delegated administrator account is granted authorization to access the organizational information of the resource directory, which provides visibility into the organization members and structure. The delegated administrator account executes service-related management operations on behalf of the management account in the specified trusted service.
//
// Limits:
//
// - Only some trusted services support the delegated administrator feature. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
// - Only the management account of the resource directory and authorized Resource Access Management (RAM) users and RAM roles under the management account can invoke this API operation.
//
// - The maximum number of delegated administrator accounts that can be added is defined by each trusted service.
//
// @param request - RegisterDelegatedAdministratorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDelegatedAdministratorResponse
func (client *Client) RegisterDelegatedAdministratorWithOptions(request *RegisterDelegatedAdministratorRequest, runtime *dara.RuntimeOptions) (_result *RegisterDelegatedAdministratorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a member account in a resource directory as a delegated administrator account for a trusted service.
//
// Description:
//
// After the settings are complete, the delegated administrator account is granted authorization to access the organizational information of the resource directory, which provides visibility into the organization members and structure. The delegated administrator account executes service-related management operations on behalf of the management account in the specified trusted service.
//
// Limits:
//
// - Only some trusted services support the delegated administrator feature. For more information, see [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
//
// - Only the management account of the resource directory and authorized Resource Access Management (RAM) users and RAM roles under the management account can invoke this API operation.
//
// - The maximum number of delegated administrator accounts that can be added is defined by each trusted service.
//
// @param request - RegisterDelegatedAdministratorRequest
//
// @return RegisterDelegatedAdministratorResponse
func (client *Client) RegisterDelegatedAdministrator(request *RegisterDelegatedAdministratorRequest) (_result *RegisterDelegatedAdministratorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterDelegatedAdministratorResponse{}
	_body, _err := client.RegisterDelegatedAdministratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a member of the cloud account type. After the member is removed, it exists as an independent Alibaba Cloud account and is no longer managed by the management account of the resource directory.
//
// @param request - RemoveCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveCloudAccountResponse
func (client *Client) RemoveCloudAccountWithOptions(request *RemoveCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *RemoveCloudAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member of the cloud account type. After the member is removed, it exists as an independent Alibaba Cloud account and is no longer managed by the management account of the resource directory.
//
// @param request - RemoveCloudAccountRequest
//
// @return RemoveCloudAccountResponse
func (client *Client) RemoveCloudAccount(request *RemoveCloudAccountRequest) (_result *RemoveCloudAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveCloudAccountResponse{}
	_body, _err := client.RemoveCloudAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for creating a cloud account.
//
// @param request - ResendCreateCloudAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendCreateCloudAccountEmailResponse
func (client *Client) ResendCreateCloudAccountEmailWithOptions(request *ResendCreateCloudAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for creating a cloud account.
//
// @param request - ResendCreateCloudAccountEmailRequest
//
// @return ResendCreateCloudAccountEmailResponse
func (client *Client) ResendCreateCloudAccountEmail(request *ResendCreateCloudAccountEmailRequest) (_result *ResendCreateCloudAccountEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResendCreateCloudAccountEmailResponse{}
	_body, _err := client.ResendCreateCloudAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for upgrading a resource account to a cloud account.
//
// @param request - ResendPromoteResourceAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendPromoteResourceAccountEmailResponse
func (client *Client) ResendPromoteResourceAccountEmailWithOptions(request *ResendPromoteResourceAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for upgrading a resource account to a cloud account.
//
// @param request - ResendPromoteResourceAccountEmailRequest
//
// @return ResendPromoteResourceAccountEmailResponse
func (client *Client) ResendPromoteResourceAccountEmail(request *ResendPromoteResourceAccountEmailRequest) (_result *ResendPromoteResourceAccountEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResendPromoteResourceAccountEmailResponse{}
	_body, _err := client.ResendPromoteResourceAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for changing the email address of a member account.
//
// @param request - RetryChangeAccountEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryChangeAccountEmailResponse
func (client *Client) RetryChangeAccountEmailWithOptions(request *RetryChangeAccountEmailRequest, runtime *dara.RuntimeOptions) (_result *RetryChangeAccountEmailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends the confirmation email for changing the email address of a member account.
//
// @param request - RetryChangeAccountEmailRequest
//
// @return RetryChangeAccountEmailResponse
func (client *Client) RetryChangeAccountEmail(request *RetryChangeAccountEmailRequest) (_result *RetryChangeAccountEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryChangeAccountEmailResponse{}
	_body, _err := client.RetryChangeAccountEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a verification code to a secure phone number when you set the secure phone number for a member of the resource account type by calling SendVerificationCodeForBindSecureMobilePhone.
//
// Description:
//
// To ensure that the system can record the specific operator of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource directory management permission (AliyunResourceDirectoryFullAccess) to call this API operation.
//
// @param request - SendVerificationCodeForBindSecureMobilePhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForBindSecureMobilePhoneResponse
func (client *Client) SendVerificationCodeForBindSecureMobilePhoneWithOptions(request *SendVerificationCodeForBindSecureMobilePhoneRequest, runtime *dara.RuntimeOptions) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to a secure phone number when you set the secure phone number for a member of the resource account type by calling SendVerificationCodeForBindSecureMobilePhone.
//
// Description:
//
// To ensure that the system can record the specific operator of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource directory management permission (AliyunResourceDirectoryFullAccess) to call this API operation.
//
// @param request - SendVerificationCodeForBindSecureMobilePhoneRequest
//
// @return SendVerificationCodeForBindSecureMobilePhoneResponse
func (client *Client) SendVerificationCodeForBindSecureMobilePhone(request *SendVerificationCodeForBindSecureMobilePhoneRequest) (_result *SendVerificationCodeForBindSecureMobilePhoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendVerificationCodeForBindSecureMobilePhoneResponse{}
	_body, _err := client.SendVerificationCodeForBindSecureMobilePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a verification code to the secure phone number of a new management account when you call SendVerificationCodeForEnableRD to enable a resource directory by creating a new management account.
//
// Description:
//
// Each Alibaba Cloud account can send up to 100 verification codes per day.
//
// @param request - SendVerificationCodeForEnableRDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationCodeForEnableRDResponse
func (client *Client) SendVerificationCodeForEnableRDWithOptions(request *SendVerificationCodeForEnableRDRequest, runtime *dara.RuntimeOptions) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a verification code to the secure phone number of a new management account when you call SendVerificationCodeForEnableRD to enable a resource directory by creating a new management account.
//
// Description:
//
// Each Alibaba Cloud account can send up to 100 verification codes per day.
//
// @param request - SendVerificationCodeForEnableRDRequest
//
// @return SendVerificationCodeForEnableRDResponse
func (client *Client) SendVerificationCodeForEnableRD(request *SendVerificationCodeForEnableRDRequest) (_result *SendVerificationCodeForEnableRDResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendVerificationCodeForEnableRDResponse{}
	_body, _err := client.SendVerificationCodeForEnableRDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the default version of an access policy.
//
// @param request - SetDefaultPolicyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultPolicyVersionResponse
func (client *Client) SetDefaultPolicyVersionWithOptions(request *SetDefaultPolicyVersionRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultPolicyVersionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the default version of an access policy.
//
// @param request - SetDefaultPolicyVersionRequest
//
// @return SetDefaultPolicyVersionResponse
func (client *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (_result *SetDefaultPolicyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultPolicyVersionResponse{}
	_body, _err := client.SetDefaultPolicyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the member deletion permission by calling SetMemberDeletionPermission.
//
// Description:
//
// You can delete members of the resource account type only after the member deletion permission is enabled.
//
// @param request - SetMemberDeletionPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMemberDeletionPermissionResponse
func (client *Client) SetMemberDeletionPermissionWithOptions(request *SetMemberDeletionPermissionRequest, runtime *dara.RuntimeOptions) (_result *SetMemberDeletionPermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the member deletion permission by calling SetMemberDeletionPermission.
//
// Description:
//
// You can delete members of the resource account type only after the member deletion permission is enabled.
//
// @param request - SetMemberDeletionPermissionRequest
//
// @return SetMemberDeletionPermissionResponse
func (client *Client) SetMemberDeletionPermission(request *SetMemberDeletionPermissionRequest) (_result *SetMemberDeletionPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetMemberDeletionPermissionResponse{}
	_body, _err := client.SetMemberDeletionPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches labels to resource groups or member accounts of a resource directory.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches labels to resource groups or member accounts of a resource directory.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds tags from resource groups or resource directory members.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds tags from resource groups or resource directory members.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a member or switches the member type by calling UpdateAccount.
//
// Description:
//
// ### Before you begin
//
// - To ensure that the system can record the specific operator of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource folder management permission (AliyunResourceDirectoryFullAccess) to invoke this API operation.
//
// - For the conditions that must be met to switch a resource account to a cloud account, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
// - For the conditions that must be met to switch a cloud account to a resource account, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// @param request - UpdateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountResponse
func (client *Client) UpdateAccountWithOptions(request *UpdateAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a member or switches the member type by calling UpdateAccount.
//
// Description:
//
// ### Before you begin
//
// - To ensure that the system can record the specific operator of management operations, use a Resource Access Management (RAM) user or RAM role under the management account that has the resource folder management permission (AliyunResourceDirectoryFullAccess) to invoke this API operation.
//
// - For the conditions that must be met to switch a resource account to a cloud account, see [Switch a resource account to a cloud account](https://help.aliyun.com/document_detail/111233.html).
//
// - For the conditions that must be met to switch a cloud account to a resource account, see [Switch a cloud account to a resource account](https://help.aliyun.com/document_detail/209980.html).
//
// @param request - UpdateAccountRequest
//
// @return UpdateAccountResponse
func (client *Client) UpdateAccount(request *UpdateAccountRequest) (_result *UpdateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccountResponse{}
	_body, _err := client.UpdateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the settings of the associated resource group transfer feature by calling UpdateAssociatedTransferSetting.
//
// Description:
//
// For cloud resources that support the attached-resource allocation feature, see [Attached-resource allocation](https://help.aliyun.com/document_detail/606232.html).
//
// @param request - UpdateAssociatedTransferSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssociatedTransferSettingResponse
func (client *Client) UpdateAssociatedTransferSettingWithOptions(request *UpdateAssociatedTransferSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateAssociatedTransferSettingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the settings of the associated resource group transfer feature by calling UpdateAssociatedTransferSetting.
//
// Description:
//
// For cloud resources that support the attached-resource allocation feature, see [Attached-resource allocation](https://help.aliyun.com/document_detail/606232.html).
//
// @param request - UpdateAssociatedTransferSettingRequest
//
// @return UpdateAssociatedTransferSettingResponse
func (client *Client) UpdateAssociatedTransferSetting(request *UpdateAssociatedTransferSettingRequest) (_result *UpdateAssociatedTransferSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAssociatedTransferSettingResponse{}
	_body, _err := client.UpdateAssociatedTransferSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls UpdateAutoGroupingConfig to update the automatic group transfer settings. Currently, only the setting for whether to enable existing associated resources to follow the group transfer task is supported.
//
// @param request - UpdateAutoGroupingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoGroupingConfigResponse
func (client *Client) UpdateAutoGroupingConfigWithOptions(request *UpdateAutoGroupingConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoGroupingConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls UpdateAutoGroupingConfig to update the automatic group transfer settings. Currently, only the setting for whether to enable existing associated resources to follow the group transfer task is supported.
//
// @param request - UpdateAutoGroupingConfigRequest
//
// @return UpdateAutoGroupingConfigResponse
func (client *Client) UpdateAutoGroupingConfig(request *UpdateAutoGroupingConfigRequest) (_result *UpdateAutoGroupingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoGroupingConfigResponse{}
	_body, _err := client.UpdateAutoGroupingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an automatic resource group transfer rule by calling UpdateAutoGroupingRule.
//
// @param request - UpdateAutoGroupingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoGroupingRuleResponse
func (client *Client) UpdateAutoGroupingRuleWithOptions(request *UpdateAutoGroupingRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoGroupingRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an automatic resource group transfer rule by calling UpdateAutoGroupingRule.
//
// @param request - UpdateAutoGroupingRuleRequest
//
// @return UpdateAutoGroupingRuleResponse
func (client *Client) UpdateAutoGroupingRule(request *UpdateAutoGroupingRuleRequest) (_result *UpdateAutoGroupingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoGroupingRuleResponse{}
	_body, _err := client.UpdateAutoGroupingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a custom control policy by calling UpdateControlPolicy.
//
// Description:
//
// This topic provides an example on how to modify the name of the control policy `cp-jExXAqIYkwHN****` to `NewControlPolicy`.
//
// @param request - UpdateControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateControlPolicyResponse
func (client *Client) UpdateControlPolicyWithOptions(request *UpdateControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom control policy by calling UpdateControlPolicy.
//
// Description:
//
// This topic provides an example on how to modify the name of the control policy `cp-jExXAqIYkwHN****` to `NewControlPolicy`.
//
// @param request - UpdateControlPolicyRequest
//
// @return UpdateControlPolicyResponse
func (client *Client) UpdateControlPolicy(request *UpdateControlPolicyRequest) (_result *UpdateControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateControlPolicyResponse{}
	_body, _err := client.UpdateControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a folder by calling UpdateFolder.
//
// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a folder by calling UpdateFolder.
//
// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the display name of a resource group.
//
// @param request - UpdateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithOptions(request *UpdateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the display name of a resource group.
//
// @param request - UpdateResourceGroupRequest
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the administrator settings of a resource group by calling UpdateResourceGroupAdminSetting.
//
// @param request - UpdateResourceGroupAdminSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupAdminSettingResponse
func (client *Client) UpdateResourceGroupAdminSettingWithOptions(request *UpdateResourceGroupAdminSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupAdminSettingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the administrator settings of a resource group by calling UpdateResourceGroupAdminSetting.
//
// @param request - UpdateResourceGroupAdminSettingRequest
//
// @return UpdateResourceGroupAdminSettingResponse
func (client *Client) UpdateResourceGroupAdminSetting(request *UpdateResourceGroupAdminSettingRequest) (_result *UpdateResourceGroupAdminSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceGroupAdminSettingResponse{}
	_body, _err := client.UpdateResourceGroupAdminSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a RAM role.
//
// @param request - UpdateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a RAM role.
//
// @param request - UpdateRoleRequest
//
// @return UpdateRoleResponse
func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
