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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudsso"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a Security Assertion Markup Language (SAML) signing certificate.
//
// Description:
//
// You can add up to two SAML signing certificates.
//
// This topic provides an example on how to add a SAML signing certificate to the directory `d-00fc2p61****`.
//
// @param request - AddExternalSAMLIdPCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddExternalSAMLIdPCertificateResponse
func (client *Client) AddExternalSAMLIdPCertificateWithOptions(request *AddExternalSAMLIdPCertificateRequest, runtime *dara.RuntimeOptions) (_result *AddExternalSAMLIdPCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.X509Certificate) {
		query["X509Certificate"] = request.X509Certificate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddExternalSAMLIdPCertificate"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddExternalSAMLIdPCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Security Assertion Markup Language (SAML) signing certificate.
//
// Description:
//
// You can add up to two SAML signing certificates.
//
// This topic provides an example on how to add a SAML signing certificate to the directory `d-00fc2p61****`.
//
// @param request - AddExternalSAMLIdPCertificateRequest
//
// @return AddExternalSAMLIdPCertificateResponse
func (client *Client) AddExternalSAMLIdPCertificate(request *AddExternalSAMLIdPCertificateRequest) (_result *AddExternalSAMLIdPCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddExternalSAMLIdPCertificateResponse{}
	_body, _err := client.AddExternalSAMLIdPCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a policy to an access configuration.
//
// Description:
//
// This topic provides an example on how to add the system policy `AliyunECSFullAccess` to the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - AddPermissionPolicyToAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPermissionPolicyToAccessConfigurationResponse
func (client *Client) AddPermissionPolicyToAccessConfigurationWithOptions(request *AddPermissionPolicyToAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *AddPermissionPolicyToAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.InlinePolicyDocument) {
		query["InlinePolicyDocument"] = request.InlinePolicyDocument
	}

	if !dara.IsNil(request.PermissionPolicyName) {
		query["PermissionPolicyName"] = request.PermissionPolicyName
	}

	if !dara.IsNil(request.PermissionPolicyType) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPermissionPolicyToAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPermissionPolicyToAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a policy to an access configuration.
//
// Description:
//
// This topic provides an example on how to add the system policy `AliyunECSFullAccess` to the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - AddPermissionPolicyToAccessConfigurationRequest
//
// @return AddPermissionPolicyToAccessConfigurationResponse
func (client *Client) AddPermissionPolicyToAccessConfiguration(request *AddPermissionPolicyToAccessConfigurationRequest) (_result *AddPermissionPolicyToAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPermissionPolicyToAccessConfigurationResponse{}
	_body, _err := client.AddPermissionPolicyToAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a group.
//
// Description:
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot add a user to a group that is synchronized by using SCIM.
//
// This topic provides an example on how to add the user `u-00q8wbq42wiltcrk****` to the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - AddUserToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a group.
//
// Description:
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot add a user to a group that is synchronized by using SCIM.
//
// This topic provides an example on how to add the user `u-00q8wbq42wiltcrk****` to the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - AddUserToGroupRequest
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (_result *AddUserToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.AddUserToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the configurations of a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// If single sign-on (SSO) logon is disabled, you can clear the configurations of a SAML IdP. If SSO logon is enabled, you cannot clear the configurations.
//
// This topic provides an example on how to clear the configurations of the SAML IdP within the directory `d-00fc2p61****`.
//
// @param request - ClearExternalSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearExternalSAMLIdentityProviderResponse
func (client *Client) ClearExternalSAMLIdentityProviderWithOptions(request *ClearExternalSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *ClearExternalSAMLIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearExternalSAMLIdentityProvider"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the configurations of a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// If single sign-on (SSO) logon is disabled, you can clear the configurations of a SAML IdP. If SSO logon is enabled, you cannot clear the configurations.
//
// This topic provides an example on how to clear the configurations of the SAML IdP within the directory `d-00fc2p61****`.
//
// @param request - ClearExternalSAMLIdentityProviderRequest
//
// @return ClearExternalSAMLIdentityProviderResponse
func (client *Client) ClearExternalSAMLIdentityProvider(request *ClearExternalSAMLIdentityProviderRequest) (_result *ClearExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearExternalSAMLIdentityProviderResponse{}
	_body, _err := client.ClearExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns access permissions on an account in your resource directory to a user or a group by using an access configuration.
//
// Description:
//
// When you call this operation, an asynchronous task is created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// For more information about how to assign permissions on an account in your resource directory, see [Overview of multi-account authorization](https://help.aliyun.com/document_detail/266726.html).
//
// This topic provides an example on how to assign access permissions on the account `114240524784****` in your resource directory to the CloudSSO user `u-00q8wbq42wiltcrk****` by using the access configuration `ac-00jhtfl8thteu6uj****`. After the call is successful, the CloudSSO user can access resources within the account in the resource directory.
//
// @param request - CreateAccessAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessAssignmentResponse
func (client *Client) CreateAccessAssignmentWithOptions(request *CreateAccessAssignmentRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessAssignmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessAssignment"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessAssignmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns access permissions on an account in your resource directory to a user or a group by using an access configuration.
//
// Description:
//
// When you call this operation, an asynchronous task is created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// For more information about how to assign permissions on an account in your resource directory, see [Overview of multi-account authorization](https://help.aliyun.com/document_detail/266726.html).
//
// This topic provides an example on how to assign access permissions on the account `114240524784****` in your resource directory to the CloudSSO user `u-00q8wbq42wiltcrk****` by using the access configuration `ac-00jhtfl8thteu6uj****`. After the call is successful, the CloudSSO user can access resources within the account in the resource directory.
//
// @param request - CreateAccessAssignmentRequest
//
// @return CreateAccessAssignmentResponse
func (client *Client) CreateAccessAssignment(request *CreateAccessAssignmentRequest) (_result *CreateAccessAssignmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessAssignmentResponse{}
	_body, _err := client.CreateAccessAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access configuration.
//
// Description:
//
// For more information about access configurations, see [Access configuration overview](https://help.aliyun.com/document_detail/266737.html).
//
// This topic provides an example on how to create an access configuration named `ECS-Admin`.
//
// @param request - CreateAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessConfigurationResponse
func (client *Client) CreateAccessConfigurationWithOptions(request *CreateAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationName) {
		query["AccessConfigurationName"] = request.AccessConfigurationName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.RelayState) {
		query["RelayState"] = request.RelayState
	}

	if !dara.IsNil(request.SessionDuration) {
		query["SessionDuration"] = request.SessionDuration
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access configuration.
//
// Description:
//
// For more information about access configurations, see [Access configuration overview](https://help.aliyun.com/document_detail/266737.html).
//
// This topic provides an example on how to create an access configuration named `ECS-Admin`.
//
// @param request - CreateAccessConfigurationRequest
//
// @return CreateAccessConfigurationResponse
func (client *Client) CreateAccessConfiguration(request *CreateAccessConfigurationRequest) (_result *CreateAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessConfigurationResponse{}
	_body, _err := client.CreateAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a directory.
//
// Description:
//
// ### [](#)Operation description
//
// A directory is a CloudSSO instance. Before you can use CloudSSO, you must create a directory. The directory is used to manage all CloudSSO resources.
//
// To create a directory, you must select a region. Alibaba Cloud stores data in the directory only in the region that you select. However, you can deploy Alibaba Cloud resources including Elastic Compute Service (ECS) instances and ApsaraDB RDS instances in other regions. You can also use your cloud account for logons and access the Alibaba Cloud resources in other regions. You can select a region to create a directory based on your security compliance requirements and the geographic location of specific users. If you do not have strict security compliance requirements, we recommend that you select a region that is the closest to the geographical location of the specific users. This way, access to cloud resources is accelerated. You can create the CloudSSO directory in the China (Shanghai), China (Hong Kong), US (Silicon Valley), or Germany (Frankfurt) region.
//
// This topic provides an example on how to create a directory named `example` in the China (Shanghai) region.
//
// ### [](#)Limits
//
//   - You can create only one directory for a management account.
//
//   - If you want to change the region of a directory, you must delete the directory and then create a directory in a different region.
//
// @param request - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(request *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryName) {
		query["DirectoryName"] = request.DirectoryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDirectory"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory.
//
// Description:
//
// ### [](#)Operation description
//
// A directory is a CloudSSO instance. Before you can use CloudSSO, you must create a directory. The directory is used to manage all CloudSSO resources.
//
// To create a directory, you must select a region. Alibaba Cloud stores data in the directory only in the region that you select. However, you can deploy Alibaba Cloud resources including Elastic Compute Service (ECS) instances and ApsaraDB RDS instances in other regions. You can also use your cloud account for logons and access the Alibaba Cloud resources in other regions. You can select a region to create a directory based on your security compliance requirements and the geographic location of specific users. If you do not have strict security compliance requirements, we recommend that you select a region that is the closest to the geographical location of the specific users. This way, access to cloud resources is accelerated. You can create the CloudSSO directory in the China (Shanghai), China (Hong Kong), US (Silicon Valley), or Germany (Frankfurt) region.
//
// This topic provides an example on how to create a directory named `example` in the China (Shanghai) region.
//
// ### [](#)Limits
//
//   - You can create only one directory for a management account.
//
//   - If you want to change the region of a directory, you must delete the directory and then create a directory in a different region.
//
// @param request - CreateDirectoryRequest
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// Description:
//
// This topic provides an example on how to create a group named `TestGroup`.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
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

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// Description:
//
// This topic provides an example on how to create a group named `TestGroup`.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// SCIM credentials are required for SCIM synchronization. You can create up to two SCIM credentials.
//
// This topic provides an example on how to create a SCIM credential within the directory `d-00fc2p61****`.
//
// @param request - CreateSCIMServerCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSCIMServerCredentialResponse
func (client *Client) CreateSCIMServerCredentialWithOptions(request *CreateSCIMServerCredentialRequest, runtime *dara.RuntimeOptions) (_result *CreateSCIMServerCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSCIMServerCredential"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSCIMServerCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// SCIM credentials are required for SCIM synchronization. You can create up to two SCIM credentials.
//
// This topic provides an example on how to create a SCIM credential within the directory `d-00fc2p61****`.
//
// @param request - CreateSCIMServerCredentialRequest
//
// @return CreateSCIMServerCredentialResponse
func (client *Client) CreateSCIMServerCredential(request *CreateSCIMServerCredentialRequest) (_result *CreateSCIMServerCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSCIMServerCredentialResponse{}
	_body, _err := client.CreateSCIMServerCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user by calling CreateUser.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
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

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.FirstName) {
		query["FirstName"] = request.FirstName
	}

	if !dara.IsNil(request.LastName) {
		query["LastName"] = request.LastName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user by calling CreateUser.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user provisioning.
//
// Description:
//
// You can create a RAM user provisioning for a member in your resource directory to create a RAM user that has the same username as a CloudSSO user. This way, the CloudSSO user can access the resources of the member as the RAM user.
//
// @param request - CreateUserProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserProvisioningResponse
func (client *Client) CreateUserProvisioningWithOptions(request *CreateUserProvisioningRequest, runtime *dara.RuntimeOptions) (_result *CreateUserProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionStrategy) {
		query["DeletionStrategy"] = request.DeletionStrategy
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DuplicationStrategy) {
		query["DuplicationStrategy"] = request.DuplicationStrategy
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserProvisioning"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user provisioning.
//
// Description:
//
// You can create a RAM user provisioning for a member in your resource directory to create a RAM user that has the same username as a CloudSSO user. This way, the CloudSSO user can access the resources of the member as the RAM user.
//
// @param request - CreateUserProvisioningRequest
//
// @return CreateUserProvisioningResponse
func (client *Client) CreateUserProvisioning(request *CreateUserProvisioningRequest) (_result *CreateUserProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserProvisioningResponse{}
	_body, _err := client.CreateUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the access permissions on an account in a resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to remove the access permissions on the account `114240524784****` in the resource directory from the CloudSSO user `u-00q8wbq42wiltcrk****`. The access permissions are assigned by using the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - DeleteAccessAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessAssignmentResponse
func (client *Client) DeleteAccessAssignmentWithOptions(request *DeleteAccessAssignmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessAssignmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DeprovisionStrategy) {
		query["DeprovisionStrategy"] = request.DeprovisionStrategy
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessAssignment"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessAssignmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the access permissions on an account in a resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to remove the access permissions on the account `114240524784****` in the resource directory from the CloudSSO user `u-00q8wbq42wiltcrk****`. The access permissions are assigned by using the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - DeleteAccessAssignmentRequest
//
// @return DeleteAccessAssignmentResponse
func (client *Client) DeleteAccessAssignment(request *DeleteAccessAssignmentRequest) (_result *DeleteAccessAssignmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessAssignmentResponse{}
	_body, _err := client.DeleteAccessAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access configuration.
//
// Description:
//
// ### [](#)Prerequisites
//
// The access configuration that you want to delete is de-provisioned from the accounts in your resource directory. For more information, see [DeprovisionAccessConfiguration](https://help.aliyun.com/document_detail/338352.html).
//
// ### [](#)Operation description
//
// This topic provides an example on how to delete the access configuration whose ID is `ac-001j9mcm3k7335bc****`.
//
// @param request - DeleteAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessConfigurationResponse
func (client *Client) DeleteAccessConfigurationWithOptions(request *DeleteAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.ForceRemovePermissionPolicies) {
		query["ForceRemovePermissionPolicies"] = request.ForceRemovePermissionPolicies
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access configuration.
//
// Description:
//
// ### [](#)Prerequisites
//
// The access configuration that you want to delete is de-provisioned from the accounts in your resource directory. For more information, see [DeprovisionAccessConfiguration](https://help.aliyun.com/document_detail/338352.html).
//
// ### [](#)Operation description
//
// This topic provides an example on how to delete the access configuration whose ID is `ac-001j9mcm3k7335bc****`.
//
// @param request - DeleteAccessConfigurationRequest
//
// @return DeleteAccessConfigurationResponse
func (client *Client) DeleteAccessConfiguration(request *DeleteAccessConfigurationRequest) (_result *DeleteAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessConfigurationResponse{}
	_body, _err := client.DeleteAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a directory.
//
// Description:
//
// ### [](#)Prerequisites
//
// No resources are contained in the directory that you want to delete.
//
//   - Access permissions on the accounts in your resource directory are removed from all users and groups. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
//   - Users are deleted. For more information, see [DeleteUser](https://help.aliyun.com/document_detail/341671.html).
//
//   - Groups are deleted. For more information, see [DeleteGroup](https://help.aliyun.com/document_detail/341821.html).
//
//   - Access configurations are deleted. For more information, see [DeleteAccessConfiguration](https://help.aliyun.com/document_detail/336907.html).
//
//   - System for Cross-domain Identity Management (SCIM) credentials are deleted. For more information, see [DeleteSCIMServerCredential](https://help.aliyun.com/document_detail/341842.html).
//
//   - Single sign-on (SSO) logon configurations are deleted. For more information, see [ClearExternalSAMLIdentityProvider](https://help.aliyun.com/document_detail/341573.html).
//
// ### [](#)Operation description
//
// This topic provides an example on how to delete a directory whose ID is `d-00fc2p61****`.
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDirectory"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a directory.
//
// Description:
//
// ### [](#)Prerequisites
//
// No resources are contained in the directory that you want to delete.
//
//   - Access permissions on the accounts in your resource directory are removed from all users and groups. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
//   - Users are deleted. For more information, see [DeleteUser](https://help.aliyun.com/document_detail/341671.html).
//
//   - Groups are deleted. For more information, see [DeleteGroup](https://help.aliyun.com/document_detail/341821.html).
//
//   - Access configurations are deleted. For more information, see [DeleteAccessConfiguration](https://help.aliyun.com/document_detail/336907.html).
//
//   - System for Cross-domain Identity Management (SCIM) credentials are deleted. For more information, see [DeleteSCIMServerCredential](https://help.aliyun.com/document_detail/341842.html).
//
//   - Single sign-on (SSO) logon configurations are deleted. For more information, see [ClearExternalSAMLIdentityProvider](https://help.aliyun.com/document_detail/341573.html).
//
// ### [](#)Operation description
//
// This topic provides an example on how to delete a directory whose ID is `d-00fc2p61****`.
//
// @param request - DeleteDirectoryRequest
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a group.
//
// Description:
//
// ### [](#)Prerequisites
//
// The group that you want to delete is not associated with the following resources. If the group is associated with the resources, the deletion fails.
//
// - Users: You must remove users from the group. For more information, see [RemoveUserFromGroup](https://help.aliyun.com/document_detail/335116.html).
//
// - Access permissions: You must remove the access permissions on the accounts in your resource directory from the group. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
// ### [](#)Operation description
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a group that is synchronized by using SCIM.
//
// This topic provides an example on how to delete the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a group.
//
// Description:
//
// ### [](#)Prerequisites
//
// The group that you want to delete is not associated with the following resources. If the group is associated with the resources, the deletion fails.
//
// - Users: You must remove users from the group. For more information, see [RemoveUserFromGroup](https://help.aliyun.com/document_detail/335116.html).
//
// - Access permissions: You must remove the access permissions on the accounts in your resource directory from the group. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
// ### [](#)Operation description
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a group that is synchronized by using SCIM.
//
// This topic provides an example on how to delete the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the MFA device of a user.
//
// Description:
//
// This topic provides an example on how to delete the MFA device `mfa-00ujhet8pycljj7j****` that is attached to the user `u-00q8wbq42wiltcrk****`.
//
// @param request - DeleteMFADeviceForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMFADeviceForUserResponse
func (client *Client) DeleteMFADeviceForUserWithOptions(request *DeleteMFADeviceForUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteMFADeviceForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MFADeviceId) {
		query["MFADeviceId"] = request.MFADeviceId
	}

	if !dara.IsNil(request.MfaType) {
		query["MfaType"] = request.MfaType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMFADeviceForUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMFADeviceForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the MFA device of a user.
//
// Description:
//
// This topic provides an example on how to delete the MFA device `mfa-00ujhet8pycljj7j****` that is attached to the user `u-00q8wbq42wiltcrk****`.
//
// @param request - DeleteMFADeviceForUserRequest
//
// @return DeleteMFADeviceForUserResponse
func (client *Client) DeleteMFADeviceForUser(request *DeleteMFADeviceForUserRequest) (_result *DeleteMFADeviceForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMFADeviceForUserResponse{}
	_body, _err := client.DeleteMFADeviceForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// After a SCIM credential is deleted, the synchronization task that uses the SCIM 
//
// This topic provides an example on how to delete the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`.
//
// @param request - DeleteSCIMServerCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSCIMServerCredentialResponse
func (client *Client) DeleteSCIMServerCredentialWithOptions(request *DeleteSCIMServerCredentialRequest, runtime *dara.RuntimeOptions) (_result *DeleteSCIMServerCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSCIMServerCredential"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSCIMServerCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// After a SCIM credential is deleted, the synchronization task that uses the SCIM credential fails.
//
// This topic provides an example on how to delete the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`.
//
// @param request - DeleteSCIMServerCredentialRequest
//
// @return DeleteSCIMServerCredentialResponse
func (client *Client) DeleteSCIMServerCredential(request *DeleteSCIMServerCredentialRequest) (_result *DeleteSCIMServerCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSCIMServerCredentialResponse{}
	_body, _err := client.DeleteSCIMServerCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you delete a user, make sure that the user is not associated with the following resources. Otherwise, the deletion fails.
//
//   - Multi-factor authentication (MFA) devices: You must delete the MFA devices bound to the user. For more information, see [DeleteMFADeviceForUser](https://help.aliyun.com/document_detail/341675.html).
//
//   - Access permissions: You must remove the access permissions on the accounts in your resource directory from the user. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
//   - Groups: You must remove the user from groups. For more information, see [RemoveUserFromGroup](https://help.aliyun.com/document_detail/335116.html).
//
// ### [](#)Precautions
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a user that is synchronized by using SCIM.
//
// This topic provides an example on how to delete the user whose ID is `u-00q8wbq42wiltcrk****`.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you delete a user, make sure that the user is not associated with the following resources. Otherwise, the deletion fails.
//
//   - Multi-factor authentication (MFA) devices: You must delete the MFA devices bound to the user. For more information, see [DeleteMFADeviceForUser](https://help.aliyun.com/document_detail/341675.html).
//
//   - Access permissions: You must remove the access permissions on the accounts in your resource directory from the user. For more information, see [DeleteAccessAssignment](https://help.aliyun.com/document_detail/338350.html).
//
//   - Groups: You must remove the user from groups. For more information, see [RemoveUserFromGroup](https://help.aliyun.com/document_detail/335116.html).
//
// ### [](#)Precautions
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot delete a user that is synchronized by using SCIM.
//
// This topic provides an example on how to delete the user whose ID is `u-00q8wbq42wiltcrk****`.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user provisioning.
//
// @param request - DeleteUserProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserProvisioningResponse
func (client *Client) DeleteUserProvisioningWithOptions(request *DeleteUserProvisioningRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionStrategy) {
		query["DeletionStrategy"] = request.DeletionStrategy
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserProvisioning"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user provisioning.
//
// @param request - DeleteUserProvisioningRequest
//
// @return DeleteUserProvisioningResponse
func (client *Client) DeleteUserProvisioning(request *DeleteUserProvisioningRequest) (_result *DeleteUserProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserProvisioningResponse{}
	_body, _err := client.DeleteUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user provisioning event.
//
// @param request - DeleteUserProvisioningEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserProvisioningEventResponse
func (client *Client) DeleteUserProvisioningEventWithOptions(request *DeleteUserProvisioningEventRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserProvisioningEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserProvisioningEvent"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user provisioning event.
//
// @param request - DeleteUserProvisioningEventRequest
//
// @return DeleteUserProvisioningEventResponse
func (client *Client) DeleteUserProvisioningEvent(request *DeleteUserProvisioningEventRequest) (_result *DeleteUserProvisioningEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserProvisioningEventResponse{}
	_body, _err := client.DeleteUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// De-provisions an access configuration from an account in your resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to de-provision the access configuration `ac-00jhtfl8thteu6uj****` from the account `114240524784****` in your resource directory.
//
// @param request - DeprovisionAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionAccessConfigurationResponse
func (client *Client) DeprovisionAccessConfigurationWithOptions(request *DeprovisionAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeprovisionAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeprovisionAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeprovisionAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// De-provisions an access configuration from an account in your resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to de-provision the access configuration `ac-00jhtfl8thteu6uj****` from the account `114240524784****` in your resource directory.
//
// @param request - DeprovisionAccessConfigurationRequest
//
// @return DeprovisionAccessConfigurationResponse
func (client *Client) DeprovisionAccessConfiguration(request *DeprovisionAccessConfigurationRequest) (_result *DeprovisionAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeprovisionAccessConfigurationResponse{}
	_body, _err := client.DeprovisionAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the delegated administrator account of CloudSSO.
//
// @param request - DisableDelegateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDelegateAccountResponse
func (client *Client) DisableDelegateAccountWithOptions(request *DisableDelegateAccountRequest, runtime *dara.RuntimeOptions) (_result *DisableDelegateAccountResponse, _err error) {
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
		Action:      dara.String("DisableDelegateAccount"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDelegateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the delegated administrator account of CloudSSO.
//
// @param request - DisableDelegateAccountRequest
//
// @return DisableDelegateAccountResponse
func (client *Client) DisableDelegateAccount(request *DisableDelegateAccountRequest) (_result *DisableDelegateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDelegateAccountResponse{}
	_body, _err := client.DisableDelegateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables CloudSSO.
//
// Description:
//
// You can disable CloudSSO only when no directories exist in CloudSSO. After you disable CloudSSO, you can re-enable it at any time.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableServiceResponse
func (client *Client) DisableServiceWithOptions(runtime *dara.RuntimeOptions) (_result *DisableServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DisableService"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables CloudSSO.
//
// Description:
//
// You can disable CloudSSO only when no directories exist in CloudSSO. After you disable CloudSSO, you can re-enable it at any time.
//
// @return DisableServiceResponse
func (client *Client) DisableService() (_result *DisableServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableServiceResponse{}
	_body, _err := client.DisableServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the delegated administrator account of CloudSSO.
//
// Description:
//
// You can use the management account of a resource directory to specify a member of the resource directory as the delegated administrator account of CloudSSO. For more information, see [Add a delegated administrator account](https://help.aliyun.com/document_detail/208117.html).
//
// After the delegated administrator account of CloudSSO is specified, you can call this operation to enable the delegated administrator account of CloudSSO to manage CloudSSO resources.
//
// @param request - EnableDelegateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDelegateAccountResponse
func (client *Client) EnableDelegateAccountWithOptions(request *EnableDelegateAccountRequest, runtime *dara.RuntimeOptions) (_result *EnableDelegateAccountResponse, _err error) {
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
		Action:      dara.String("EnableDelegateAccount"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDelegateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the delegated administrator account of CloudSSO.
//
// Description:
//
// You can use the management account of a resource directory to specify a member of the resource directory as the delegated administrator account of CloudSSO. For more information, see [Add a delegated administrator account](https://help.aliyun.com/document_detail/208117.html).
//
// After the delegated administrator account of CloudSSO is specified, you can call this operation to enable the delegated administrator account of CloudSSO to manage CloudSSO resources.
//
// @param request - EnableDelegateAccountRequest
//
// @return EnableDelegateAccountResponse
func (client *Client) EnableDelegateAccount(request *EnableDelegateAccountRequest) (_result *EnableDelegateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDelegateAccountResponse{}
	_body, _err := client.EnableDelegateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables CloudSSO.
//
// Description:
//
// Only users under the management account of a resource directory who have the permissions to enable CloudSSO can call this operation. For more information, see [Enable CloudSSO](https://help.aliyun.com/document_detail/262819.html).
//
// By calling this operation, you agree to the [Alibaba Cloud International Website Product Terms of Service](https://www.alibabacloud.com/help/doc-detail/42416.htm).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableServiceResponse
func (client *Client) EnableServiceWithOptions(runtime *dara.RuntimeOptions) (_result *EnableServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableService"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables CloudSSO.
//
// Description:
//
// Only users under the management account of a resource directory who have the permissions to enable CloudSSO can call this operation. For more information, see [Enable CloudSSO](https://help.aliyun.com/document_detail/262819.html).
//
// By calling this operation, you agree to the [Alibaba Cloud International Website Product Terms of Service](https://www.alibabacloud.com/help/doc-detail/42416.htm).
//
// @return EnableServiceResponse
func (client *Client) EnableService() (_result *EnableServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableServiceResponse{}
	_body, _err := client.EnableServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an access configuration.
//
// Description:
//
// This topic provides an example on how to query the information about the access configuration whose ID is `ac-00ccule7tadaijxc****`.
//
// @param request - GetAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessConfigurationResponse
func (client *Client) GetAccessConfigurationWithOptions(request *GetAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an access configuration.
//
// Description:
//
// This topic provides an example on how to query the information about the access configuration whose ID is `ac-00ccule7tadaijxc****`.
//
// @param request - GetAccessConfigurationRequest
//
// @return GetAccessConfigurationResponse
func (client *Client) GetAccessConfiguration(request *GetAccessConfigurationRequest) (_result *GetAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessConfigurationResponse{}
	_body, _err := client.GetAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attribute passing settings of a specified directory to retrieve the current configuration of the SourceIdentity pass-through mode.
//
// Description:
//
// You must have the cloudsso:GetAttributePassingSetting permission to call this operation. If the directory is not explicitly configured, SourceIdentityPassing returns Disabled by default.
//
// @param request - GetAttributePassingSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttributePassingSettingResponse
func (client *Client) GetAttributePassingSettingWithOptions(request *GetAttributePassingSettingRequest, runtime *dara.RuntimeOptions) (_result *GetAttributePassingSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttributePassingSetting"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttributePassingSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attribute passing settings of a specified directory to retrieve the current configuration of the SourceIdentity pass-through mode.
//
// Description:
//
// You must have the cloudsso:GetAttributePassingSetting permission to call this operation. If the directory is not explicitly configured, SourceIdentityPassing returns Disabled by default.
//
// @param request - GetAttributePassingSettingRequest
//
// @return GetAttributePassingSettingResponse
func (client *Client) GetAttributePassingSetting(request *GetAttributePassingSettingRequest) (_result *GetAttributePassingSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAttributePassingSettingResponse{}
	_body, _err := client.GetAttributePassingSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a directory.
//
// Description:
//
// This topic provides an example on how to query information about the directory whose ID is `d-00fc2p61****`.
//
// @param request - GetDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryResponse
func (client *Client) GetDirectoryWithOptions(request *GetDirectoryRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectory"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a directory.
//
// Description:
//
// This topic provides an example on how to query information about the directory whose ID is `d-00fc2p61****`.
//
// @param request - GetDirectoryRequest
//
// @return GetDirectoryResponse
func (client *Client) GetDirectory(request *GetDirectoryRequest) (_result *GetDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDirectoryResponse{}
	_body, _err := client.GetDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a Security Assertion Markup Language (SAML) service provider (SP).
//
// Description:
//
// During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an identity provider (IdP).
//
// This topic provides an example on how to query information about the SP within the directory `d-00fc2p61****`.
//
// @param request - GetDirectorySAMLServiceProviderInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectorySAMLServiceProviderInfoResponse
func (client *Client) GetDirectorySAMLServiceProviderInfoWithOptions(request *GetDirectorySAMLServiceProviderInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDirectorySAMLServiceProviderInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectorySAMLServiceProviderInfo"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectorySAMLServiceProviderInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Security Assertion Markup Language (SAML) service provider (SP).
//
// Description:
//
// During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is an SP, and the identity management system of an enterprise is an identity provider (IdP).
//
// This topic provides an example on how to query information about the SP within the directory `d-00fc2p61****`.
//
// @param request - GetDirectorySAMLServiceProviderInfoRequest
//
// @return GetDirectorySAMLServiceProviderInfoResponse
func (client *Client) GetDirectorySAMLServiceProviderInfo(request *GetDirectorySAMLServiceProviderInfoRequest) (_result *GetDirectorySAMLServiceProviderInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDirectorySAMLServiceProviderInfoResponse{}
	_body, _err := client.GetDirectorySAMLServiceProviderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of a directory.
//
// Description:
//
// This topic provides an example on how to query the statistics of a directory whose ID is `d-00fc2p61****`. The statistics include the number of users, quota for users, number of groups, quota for groups, number of access configurations, quota for access configurations, number of access permissions that are assigned, number of system policies that can be configured for an access configuration, number of System for Cross-domain Identity Management (SCIM) credentials, number of asynchronous tasks, status of single sign-on (SSO) logon, and status of SCIM synchronization.
//
// @param request - GetDirectoryStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryStatisticsResponse
func (client *Client) GetDirectoryStatisticsWithOptions(request *GetDirectoryStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectoryStatistics"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of a directory.
//
// Description:
//
// This topic provides an example on how to query the statistics of a directory whose ID is `d-00fc2p61****`. The statistics include the number of users, quota for users, number of groups, quota for groups, number of access configurations, quota for access configurations, number of access permissions that are assigned, number of system policies that can be configured for an access configuration, number of System for Cross-domain Identity Management (SCIM) credentials, number of asynchronous tasks, status of single sign-on (SSO) logon, and status of SCIM synchronization.
//
// @param request - GetDirectoryStatisticsRequest
//
// @return GetDirectoryStatisticsResponse
func (client *Client) GetDirectoryStatistics(request *GetDirectoryStatisticsRequest) (_result *GetDirectoryStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDirectoryStatisticsResponse{}
	_body, _err := client.GetDirectoryStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// This topic provides an example on how to query the configurations of the SAML IdP within the directory `d-00fc2p61****`.
//
// @param request - GetExternalSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalSAMLIdentityProviderResponse
func (client *Client) GetExternalSAMLIdentityProviderWithOptions(request *GetExternalSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetExternalSAMLIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExternalSAMLIdentityProvider"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// This topic provides an example on how to query the configurations of the SAML IdP within the directory `d-00fc2p61****`.
//
// @param request - GetExternalSAMLIdentityProviderRequest
//
// @return GetExternalSAMLIdentityProviderResponse
func (client *Client) GetExternalSAMLIdentityProvider(request *GetExternalSAMLIdentityProviderRequest) (_result *GetExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.GetExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a group.
//
// Description:
//
// This topic provides an example on how to query information about the group `g-00jqzghi2n3o5hkh****` in the directory `d-00fc2p61****`.
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a group.
//
// Description:
//
// This topic provides an example on how to query information about the group `g-00jqzghi2n3o5hkh****` in the directory `d-00fc2p61****`.
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logon preference of CloudSSO users.
//
// @param request - GetLoginPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginPreferenceResponse
func (client *Client) GetLoginPreferenceWithOptions(request *GetLoginPreferenceRequest, runtime *dara.RuntimeOptions) (_result *GetLoginPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginPreference"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logon preference of CloudSSO users.
//
// @param request - GetLoginPreferenceRequest
//
// @return GetLoginPreferenceResponse
func (client *Client) GetLoginPreference(request *GetLoginPreferenceRequest) (_result *GetLoginPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.GetLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global multi-factor authentication (MFA) configuration.
//
// Description:
//
// When username-password logon is enabled, you can retrieve the global MFA verification policy for user logon.
//
// This topic provides an example on how to query the global MFA verification policy for CloudSSO users in the directory `u-00q8wbq42wiltcrk****`.
//
// @param request - GetMFAAuthenticationSettingInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMFAAuthenticationSettingInfoResponse
func (client *Client) GetMFAAuthenticationSettingInfoWithOptions(request *GetMFAAuthenticationSettingInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMFAAuthenticationSettingInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMFAAuthenticationSettingInfo"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMFAAuthenticationSettingInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global multi-factor authentication (MFA) configuration.
//
// Description:
//
// When username-password logon is enabled, you can retrieve the global MFA verification policy for user logon.
//
// This topic provides an example on how to query the global MFA verification policy for CloudSSO users in the directory `u-00q8wbq42wiltcrk****`.
//
// @param request - GetMFAAuthenticationSettingInfoRequest
//
// @return GetMFAAuthenticationSettingInfoResponse
func (client *Client) GetMFAAuthenticationSettingInfo(request *GetMFAAuthenticationSettingInfoRequest) (_result *GetMFAAuthenticationSettingInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMFAAuthenticationSettingInfoResponse{}
	_body, _err := client.GetMFAAuthenticationSettingInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) setting of all users.
//
// Description:
//
// >  This operation is no longer maintained and updated. You can call the [GetMFAAuthenticationSettingInfo](https://help.aliyun.com/document_detail/611286.html) operation to query more detailed information.
//
// This topic provides an example on how to query the MFA setting of the users that belong to the directory named `d-00fc2p61****`. The returned result shows that MFA is enabled for all the users.
//
// @param request - GetMFAAuthenticationSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMFAAuthenticationSettingsResponse
func (client *Client) GetMFAAuthenticationSettingsWithOptions(request *GetMFAAuthenticationSettingsRequest, runtime *dara.RuntimeOptions) (_result *GetMFAAuthenticationSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMFAAuthenticationSettings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) setting of all users.
//
// Description:
//
// >  This operation is no longer maintained and updated. You can call the [GetMFAAuthenticationSettingInfo](https://help.aliyun.com/document_detail/611286.html) operation to query more detailed information.
//
// This topic provides an example on how to query the MFA setting of the users that belong to the directory named `d-00fc2p61****`. The returned result shows that MFA is enabled for all the users.
//
// @param request - GetMFAAuthenticationSettingsRequest
//
// @return GetMFAAuthenticationSettingsResponse
func (client *Client) GetMFAAuthenticationSettings(request *GetMFAAuthenticationSettingsRequest) (_result *GetMFAAuthenticationSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMFAAuthenticationSettingsResponse{}
	_body, _err := client.GetMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether multi-factor authentication (MFA) is enabled for users.
//
// Description:
//
// This topic provides an example on how to check whether MFA is enabled for users in the directory whose ID is `d-00fc2p61****`. The returned result shows that MFA is in the Enabled state.
//
// @param request - GetMFAAuthenticationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMFAAuthenticationStatusResponse
func (client *Client) GetMFAAuthenticationStatusWithOptions(request *GetMFAAuthenticationStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMFAAuthenticationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMFAAuthenticationStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMFAAuthenticationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether multi-factor authentication (MFA) is enabled for users.
//
// Description:
//
// This topic provides an example on how to check whether MFA is enabled for users in the directory whose ID is `d-00fc2p61****`. The returned result shows that MFA is in the Enabled state.
//
// @param request - GetMFAAuthenticationStatusRequest
//
// @return GetMFAAuthenticationStatusResponse
func (client *Client) GetMFAAuthenticationStatus(request *GetMFAAuthenticationStatusRequest) (_result *GetMFAAuthenticationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMFAAuthenticationStatusResponse{}
	_body, _err := client.GetMFAAuthenticationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password policy of CloudSSO users.
//
// @param request - GetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicyWithOptions(request *GetPasswordPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordPolicy"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password policy of CloudSSO users.
//
// @param request - GetPasswordPolicyRequest
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicy(request *GetPasswordPolicyRequest) (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of System for Cross-domain Identity Management (SCIM) synchronization.
//
// Description:
//
// This topic provides an example on how to query the status of SCIM synchronization within the directory `d-00fc2p61****`. The returned result shows that SCIM synchronization is in the Enabled state.
//
// @param request - GetSCIMSynchronizationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSCIMSynchronizationStatusResponse
func (client *Client) GetSCIMSynchronizationStatusWithOptions(request *GetSCIMSynchronizationStatusRequest, runtime *dara.RuntimeOptions) (_result *GetSCIMSynchronizationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSCIMSynchronizationStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSCIMSynchronizationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of System for Cross-domain Identity Management (SCIM) synchronization.
//
// Description:
//
// This topic provides an example on how to query the status of SCIM synchronization within the directory `d-00fc2p61****`. The returned result shows that SCIM synchronization is in the Enabled state.
//
// @param request - GetSCIMSynchronizationStatusRequest
//
// @return GetSCIMSynchronizationStatusResponse
func (client *Client) GetSCIMSynchronizationStatus(request *GetSCIMSynchronizationStatusRequest) (_result *GetSCIMSynchronizationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSCIMSynchronizationStatusResponse{}
	_body, _err := client.GetSCIMSynchronizationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of CloudSSO.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceStatusResponse
func (client *Client) GetServiceStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetServiceStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of CloudSSO.
//
// @return GetServiceStatusResponse
func (client *Client) GetServiceStatus() (_result *GetServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceStatusResponse{}
	_body, _err := client.GetServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an asynchronous task.
//
// Description:
//
// This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an asynchronous task.
//
// Description:
//
// This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
//
// @param request - GetTaskRequest
//
// @return GetTaskResponse
func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an asynchronous task.
//
// Description:
//
// You can call the GetTaskStatus operation to query the status of an asynchronous task. If you want to query more information about an asynchronous task, call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation.
//
// This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
//
// @param request - GetTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an asynchronous task.
//
// Description:
//
// You can call the GetTaskStatus operation to query the status of an asynchronous task. If you want to query more information about an asynchronous task, call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation.
//
// This topic provides an example on how to query the information about the task whose ID is `t-shfqw1u1edszvxw5****`.
//
// @param request - GetTaskStatusRequest
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a specified user.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// Queries the information about a specified user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of a user in a resource directory by using the ExternalId parameter.
//
// @param tmpReq - GetUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdResponse
func (client *Client) GetUserIdWithOptions(tmpReq *GetUserIdRequest, runtime *dara.RuntimeOptions) (_result *GetUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExternalId) {
		request.ExternalIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalId, dara.String("ExternalId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.ExternalIdShrink) {
		query["ExternalId"] = request.ExternalIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserId"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// Queries the ID of a user in a resource directory by using the ExternalId parameter.
//
// @param request - GetUserIdRequest
//
// @return GetUserIdResponse
func (client *Client) GetUserId(request *GetUserIdRequest) (_result *GetUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserIdResponse{}
	_body, _err := client.GetUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) setting of a single user.
//
// Description:
//
// This topic provides an example on how to query the MFA setting of the user named `u-00q8wbq42wiltcrk****`. The returned result shows that MFA is enabled for the user.
//
// @param request - GetUserMFAAuthenticationSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMFAAuthenticationSettingsResponse
func (client *Client) GetUserMFAAuthenticationSettingsWithOptions(request *GetUserMFAAuthenticationSettingsRequest, runtime *dara.RuntimeOptions) (_result *GetUserMFAAuthenticationSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserMFAAuthenticationSettings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the multi-factor authentication (MFA) setting of a single user.
//
// Description:
//
// This topic provides an example on how to query the MFA setting of the user named `u-00q8wbq42wiltcrk****`. The returned result shows that MFA is enabled for the user.
//
// @param request - GetUserMFAAuthenticationSettingsRequest
//
// @return GetUserMFAAuthenticationSettingsResponse
func (client *Client) GetUserMFAAuthenticationSettings(request *GetUserMFAAuthenticationSettingsRequest) (_result *GetUserMFAAuthenticationSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.GetUserMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserProvisioningResponse
func (client *Client) GetUserProvisioningWithOptions(request *GetUserProvisioningRequest, runtime *dara.RuntimeOptions) (_result *GetUserProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserProvisioning"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningRequest
//
// @return GetUserProvisioningResponse
func (client *Client) GetUserProvisioning(request *GetUserProvisioningRequest) (_result *GetUserProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserProvisioningResponse{}
	_body, _err := client.GetUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global configurations of a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserProvisioningConfigurationResponse
func (client *Client) GetUserProvisioningConfigurationWithOptions(request *GetUserProvisioningConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetUserProvisioningConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserProvisioningConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserProvisioningConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global configurations of a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningConfigurationRequest
//
// @return GetUserProvisioningConfigurationResponse
func (client *Client) GetUserProvisioningConfiguration(request *GetUserProvisioningConfigurationRequest) (_result *GetUserProvisioningConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserProvisioningConfigurationResponse{}
	_body, _err := client.GetUserProvisioningConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserProvisioningEventResponse
func (client *Client) GetUserProvisioningEventWithOptions(request *GetUserProvisioningEventRequest, runtime *dara.RuntimeOptions) (_result *GetUserProvisioningEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserProvisioningEvent"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningEventRequest
//
// @return GetUserProvisioningEventResponse
func (client *Client) GetUserProvisioningEvent(request *GetUserProvisioningEventRequest) (_result *GetUserProvisioningEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserProvisioningEventResponse{}
	_body, _err := client.GetUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics of Resource Access Management (RAM) user provisioning events that are created for the member in a resource directory.
//
// @param request - GetUserProvisioningRdAccountStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserProvisioningRdAccountStatisticsResponse
func (client *Client) GetUserProvisioningRdAccountStatisticsWithOptions(request *GetUserProvisioningRdAccountStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetUserProvisioningRdAccountStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.RdMemberId) {
		query["RdMemberId"] = request.RdMemberId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserProvisioningRdAccountStatistics"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserProvisioningRdAccountStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics of Resource Access Management (RAM) user provisioning events that are created for the member in a resource directory.
//
// @param request - GetUserProvisioningRdAccountStatisticsRequest
//
// @return GetUserProvisioningRdAccountStatisticsResponse
func (client *Client) GetUserProvisioningRdAccountStatistics(request *GetUserProvisioningRdAccountStatisticsRequest) (_result *GetUserProvisioningRdAccountStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserProvisioningRdAccountStatisticsResponse{}
	_body, _err := client.GetUserProvisioningRdAccountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserProvisioningStatisticsResponse
func (client *Client) GetUserProvisioningStatisticsWithOptions(request *GetUserProvisioningStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetUserProvisioningStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserProvisioningStatistics"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserProvisioningStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of a Resource Access Management (RAM) user provisioning.
//
// @param request - GetUserProvisioningStatisticsRequest
//
// @return GetUserProvisioningStatisticsResponse
func (client *Client) GetUserProvisioningStatistics(request *GetUserProvisioningStatisticsRequest) (_result *GetUserProvisioningStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserProvisioningStatisticsResponse{}
	_body, _err := client.GetUserProvisioningStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access permissions that are assigned.
//
// Description:
//
// This topic provides an example on how to query the assigned access permissions on the account `114240524784****` in your resource directory. The returned result shows that access permissions on the account in your resource directory is assigned to one user.
//
// @param request - ListAccessAssignmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessAssignmentsResponse
func (client *Client) ListAccessAssignmentsWithOptions(request *ListAccessAssignmentsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessAssignmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessAssignments"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessAssignmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access permissions that are assigned.
//
// Description:
//
// This topic provides an example on how to query the assigned access permissions on the account `114240524784****` in your resource directory. The returned result shows that access permissions on the account in your resource directory is assigned to one user.
//
// @param request - ListAccessAssignmentsRequest
//
// @return ListAccessAssignmentsResponse
func (client *Client) ListAccessAssignments(request *ListAccessAssignmentsRequest) (_result *ListAccessAssignmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessAssignmentsResponse{}
	_body, _err := client.ListAccessAssignmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access configurations that are provisioned.
//
// Description:
//
// This topic provides an example on how to query the accounts for which the access permission `ac-00ccule7tadaijxc****` is provisioned. The returned result shows that the access configuration is provisioned for two accounts in your resource directory.
//
// @param request - ListAccessConfigurationProvisioningsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessConfigurationProvisioningsResponse
func (client *Client) ListAccessConfigurationProvisioningsWithOptions(request *ListAccessConfigurationProvisioningsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessConfigurationProvisioningsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.ProvisioningStatus) {
		query["ProvisioningStatus"] = request.ProvisioningStatus
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessConfigurationProvisionings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessConfigurationProvisioningsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access configurations that are provisioned.
//
// Description:
//
// This topic provides an example on how to query the accounts for which the access permission `ac-00ccule7tadaijxc****` is provisioned. The returned result shows that the access configuration is provisioned for two accounts in your resource directory.
//
// @param request - ListAccessConfigurationProvisioningsRequest
//
// @return ListAccessConfigurationProvisioningsResponse
func (client *Client) ListAccessConfigurationProvisionings(request *ListAccessConfigurationProvisioningsRequest) (_result *ListAccessConfigurationProvisioningsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessConfigurationProvisioningsResponse{}
	_body, _err := client.ListAccessConfigurationProvisioningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of access configurations.
//
// Description:
//
// This topic provides an example on how to query the access configurations in the directory `d-00fc2p61****`. The response shows that there are two access configurations: `VPC-Admin` and `ECS-Admin`.
//
// @param request - ListAccessConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessConfigurationsResponse
func (client *Client) ListAccessConfigurationsWithOptions(request *ListAccessConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StatusNotifications) {
		query["StatusNotifications"] = request.StatusNotifications
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessConfigurations"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of access configurations.
//
// Description:
//
// This topic provides an example on how to query the access configurations in the directory `d-00fc2p61****`. The response shows that there are two access configurations: `VPC-Admin` and `ECS-Admin`.
//
// @param request - ListAccessConfigurationsRequest
//
// @return ListAccessConfigurationsResponse
func (client *Client) ListAccessConfigurations(request *ListAccessConfigurationsRequest) (_result *ListAccessConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessConfigurationsResponse{}
	_body, _err := client.ListAccessConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries directories.
//
// Description:
//
// This topic provides an example on how to query the directories within your Alibaba Cloud account. The returned result shows that only one directory with the ID `d-00fc2p61****` is created within your Alibaba Cloud account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDirectoriesResponse
func (client *Client) ListDirectoriesWithOptions(runtime *dara.RuntimeOptions) (_result *ListDirectoriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListDirectories"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries directories.
//
// Description:
//
// This topic provides an example on how to query the directories within your Alibaba Cloud account. The returned result shows that only one directory with the ID `d-00fc2p61****` is created within your Alibaba Cloud account.
//
// @return ListDirectoriesResponse
func (client *Client) ListDirectories() (_result *ListDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDirectoriesResponse{}
	_body, _err := client.ListDirectoriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Security Assertion Markup Language (SAML) signing certificates.
//
// Description:
//
// This topic provides an example on how to query the SAML signing certificates within the directory `d-00fc2p61****`. The returned result shows that the directory contains one SAML signing certificate.
//
// @param request - ListExternalSAMLIdPCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalSAMLIdPCertificatesResponse
func (client *Client) ListExternalSAMLIdPCertificatesWithOptions(request *ListExternalSAMLIdPCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListExternalSAMLIdPCertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExternalSAMLIdPCertificates"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalSAMLIdPCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Security Assertion Markup Language (SAML) signing certificates.
//
// Description:
//
// This topic provides an example on how to query the SAML signing certificates within the directory `d-00fc2p61****`. The returned result shows that the directory contains one SAML signing certificate.
//
// @param request - ListExternalSAMLIdPCertificatesRequest
//
// @return ListExternalSAMLIdPCertificatesResponse
func (client *Client) ListExternalSAMLIdPCertificates(request *ListExternalSAMLIdPCertificatesRequest) (_result *ListExternalSAMLIdPCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExternalSAMLIdPCertificatesResponse{}
	_body, _err := client.ListExternalSAMLIdPCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the users in a group.
//
// Description:
//
// This topic provides an example on how to query the users in the group `g-00jqzghi2n3o5hkh****`. The returned result shows that the group contains the user `Alice` and the user `user1`.
//
// @param request - ListGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupMembersResponse
func (client *Client) ListGroupMembersWithOptions(request *ListGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *ListGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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
		Action:      dara.String("ListGroupMembers"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the users in a group.
//
// Description:
//
// This topic provides an example on how to query the users in the group `g-00jqzghi2n3o5hkh****`. The returned result shows that the group contains the user `Alice` and the user `user1`.
//
// @param request - ListGroupMembersRequest
//
// @return ListGroupMembersResponse
func (client *Client) ListGroupMembers(request *ListGroupMembersRequest) (_result *ListGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupMembersResponse{}
	_body, _err := client.ListGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries groups.
//
// Description:
//
// This topic provides an example on how to query the groups in the directory `d-00fc2p61****`. The returned result shows that the directory contains three groups. The groups `group1` and `group2` are synchronized from an external identity provider (IdP). The group `TestGroup` is manually created in CloudSSO.
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProvisionType) {
		query["ProvisionType"] = request.ProvisionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries groups.
//
// Description:
//
// This topic provides an example on how to query the groups in the directory `d-00fc2p61****`. The returned result shows that the directory contains three groups. The groups `group1` and `group2` are synchronized from an external identity provider (IdP). The group `TestGroup` is manually created in CloudSSO.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the groups to which a user is added.
//
// Description:
//
// This topic provides an example on how to query the groups to which the user `u-00q8wbq42wiltcrk****` is added. The returned result shows that the user is added to both the `TestGroup` and the `group1` groups.
//
// @param request - ListJoinedGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJoinedGroupsForUserResponse
func (client *Client) ListJoinedGroupsForUserWithOptions(request *ListJoinedGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListJoinedGroupsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJoinedGroupsForUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJoinedGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the groups to which a user is added.
//
// Description:
//
// This topic provides an example on how to query the groups to which the user `u-00q8wbq42wiltcrk****` is added. The returned result shows that the user is added to both the `TestGroup` and the `group1` groups.
//
// @param request - ListJoinedGroupsForUserRequest
//
// @return ListJoinedGroupsForUserResponse
func (client *Client) ListJoinedGroupsForUser(request *ListJoinedGroupsForUserRequest) (_result *ListJoinedGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJoinedGroupsForUserResponse{}
	_body, _err := client.ListJoinedGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the MFA device list of a user. Each user can have a maximum of two MFA devices.
//
// Description:
//
// This topic provides an example on how to query the MFA device list of the user `u-00q8wbq42wiltcrk****`. The response shows that the user has one MFA device named `Alice-MFA1`.
//
// @param request - ListMFADevicesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMFADevicesForUserResponse
func (client *Client) ListMFADevicesForUserWithOptions(request *ListMFADevicesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListMFADevicesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMFADevicesForUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMFADevicesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the MFA device list of a user. Each user can have a maximum of two MFA devices.
//
// Description:
//
// This topic provides an example on how to query the MFA device list of the user `u-00q8wbq42wiltcrk****`. The response shows that the user has one MFA device named `Alice-MFA1`.
//
// @param request - ListMFADevicesForUserRequest
//
// @return ListMFADevicesForUserResponse
func (client *Client) ListMFADevicesForUser(request *ListMFADevicesForUserRequest) (_result *ListMFADevicesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMFADevicesForUserResponse{}
	_body, _err := client.ListMFADevicesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the policies that are created for an access configuration.
//
// Description:
//
// This topic provides an example on how to query the policies that are created for the access configuration `ac-00jhtfl8thteu6uj****`. The returned result shows that the access configuration contains one system policy and one inline policy.
//
// @param request - ListPermissionPoliciesInAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionPoliciesInAccessConfigurationResponse
func (client *Client) ListPermissionPoliciesInAccessConfigurationWithOptions(request *ListPermissionPoliciesInAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ListPermissionPoliciesInAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.PermissionPolicyType) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionPoliciesInAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionPoliciesInAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies that are created for an access configuration.
//
// Description:
//
// This topic provides an example on how to query the policies that are created for the access configuration `ac-00jhtfl8thteu6uj****`. The returned result shows that the access configuration contains one system policy and one inline policy.
//
// @param request - ListPermissionPoliciesInAccessConfigurationRequest
//
// @return ListPermissionPoliciesInAccessConfigurationResponse
func (client *Client) ListPermissionPoliciesInAccessConfiguration(request *ListPermissionPoliciesInAccessConfigurationRequest) (_result *ListPermissionPoliciesInAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPermissionPoliciesInAccessConfigurationResponse{}
	_body, _err := client.ListPermissionPoliciesInAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Cross-domain Identity Management (SCIM) credentials.
//
// Description:
//
// This topic provides an example on how to query the SCIM credentials within the `d-00fc2p61****` directory.
//
// @param request - ListSCIMServerCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSCIMServerCredentialsResponse
func (client *Client) ListSCIMServerCredentialsWithOptions(request *ListSCIMServerCredentialsRequest, runtime *dara.RuntimeOptions) (_result *ListSCIMServerCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSCIMServerCredentials"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSCIMServerCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Cross-domain Identity Management (SCIM) credentials.
//
// Description:
//
// This topic provides an example on how to query the SCIM credentials within the `d-00fc2p61****` directory.
//
// @param request - ListSCIMServerCredentialsRequest
//
// @return ListSCIMServerCredentialsResponse
func (client *Client) ListSCIMServerCredentials(request *ListSCIMServerCredentialsRequest) (_result *ListSCIMServerCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSCIMServerCredentialsResponse{}
	_body, _err := client.ListSCIMServerCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries asynchronous tasks.
//
// Description:
//
// By default, this operation queries the tasks within the previous 24 hours. This operation allows you to query the tasks within a maximum of 7 days. You can specify the start time of the query by using `Filter`.
//
// This topic provides an example on how to query the tasks within the previous 24 hours.
//
// @param request - ListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries asynchronous tasks.
//
// Description:
//
// By default, this operation queries the tasks within the previous 24 hours. This operation allows you to query the tasks within a maximum of 7 days. You can specify the start time of the query by using `Filter`.
//
// This topic provides an example on how to query the tasks within the previous 24 hours.
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user provisioning events.
//
// @param request - ListUserProvisioningEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserProvisioningEventsResponse
func (client *Client) ListUserProvisioningEventsWithOptions(request *ListUserProvisioningEventsRequest, runtime *dara.RuntimeOptions) (_result *ListUserProvisioningEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserProvisioningEvents"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserProvisioningEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user provisioning events.
//
// @param request - ListUserProvisioningEventsRequest
//
// @return ListUserProvisioningEventsResponse
func (client *Client) ListUserProvisioningEvents(request *ListUserProvisioningEventsRequest) (_result *ListUserProvisioningEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserProvisioningEventsResponse{}
	_body, _err := client.ListUserProvisioningEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user provisionings.
//
// @param request - ListUserProvisioningsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserProvisioningsResponse
func (client *Client) ListUserProvisioningsWithOptions(request *ListUserProvisioningsRequest, runtime *dara.RuntimeOptions) (_result *ListUserProvisioningsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PrincipalId) {
		query["PrincipalId"] = request.PrincipalId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserProvisionings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserProvisioningsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user provisionings.
//
// @param request - ListUserProvisioningsRequest
//
// @return ListUserProvisioningsResponse
func (client *Client) ListUserProvisionings(request *ListUserProvisioningsRequest) (_result *ListUserProvisioningsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserProvisioningsResponse{}
	_body, _err := client.ListUserProvisioningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProvisionType) {
		query["ProvisionType"] = request.ProvisionType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provisions an access configuration for an account in your resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to provision the access configuration `ac-00jhtfl8thteu6uj****` for the account `114240524784****` in your resource directory.
//
// @param request - ProvisionAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionAccessConfigurationResponse
func (client *Client) ProvisionAccessConfigurationWithOptions(request *ProvisionAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ProvisionAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.OriginTargetId) {
		query["OriginTargetId"] = request.OriginTargetId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProvisionAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProvisionAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provisions an access configuration for an account in your resource directory.
//
// Description:
//
// When you call this operation, an asynchronous task is automatically created. You can call the [GetTask](https://help.aliyun.com/document_detail/340670.html) operation to query the progress of the task based on the value of the `TaskId` response parameter.
//
// This topic provides an example on how to provision the access configuration `ac-00jhtfl8thteu6uj****` for the account `114240524784****` in your resource directory.
//
// @param request - ProvisionAccessConfigurationRequest
//
// @return ProvisionAccessConfigurationResponse
func (client *Client) ProvisionAccessConfiguration(request *ProvisionAccessConfigurationRequest) (_result *ProvisionAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProvisionAccessConfigurationResponse{}
	_body, _err := client.ProvisionAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Security Assertion Markup Language (SAML) signing certificate.
//
// Description:
//
// This topic provides an example on how to remove the SAML signing certificate whose ID is `idp-c-00dt9gnl7fmjaw9c****`.
//
// @param request - RemoveExternalSAMLIdPCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveExternalSAMLIdPCertificateResponse
func (client *Client) RemoveExternalSAMLIdPCertificateWithOptions(request *RemoveExternalSAMLIdPCertificateRequest, runtime *dara.RuntimeOptions) (_result *RemoveExternalSAMLIdPCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveExternalSAMLIdPCertificate"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveExternalSAMLIdPCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a Security Assertion Markup Language (SAML) signing certificate.
//
// Description:
//
// This topic provides an example on how to remove the SAML signing certificate whose ID is `idp-c-00dt9gnl7fmjaw9c****`.
//
// @param request - RemoveExternalSAMLIdPCertificateRequest
//
// @return RemoveExternalSAMLIdPCertificateResponse
func (client *Client) RemoveExternalSAMLIdPCertificate(request *RemoveExternalSAMLIdPCertificateRequest) (_result *RemoveExternalSAMLIdPCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveExternalSAMLIdPCertificateResponse{}
	_body, _err := client.RemoveExternalSAMLIdPCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a policy from an access configuration.
//
// Description:
//
// After you remove an inline policy from an access configuration, the policy cannot be restored.
//
// This topic provides an example on how to remove the system policy `AliyunECSFullAccess` from the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - RemovePermissionPolicyFromAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePermissionPolicyFromAccessConfigurationResponse
func (client *Client) RemovePermissionPolicyFromAccessConfigurationWithOptions(request *RemovePermissionPolicyFromAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *RemovePermissionPolicyFromAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.PermissionPolicyName) {
		query["PermissionPolicyName"] = request.PermissionPolicyName
	}

	if !dara.IsNil(request.PermissionPolicyType) {
		query["PermissionPolicyType"] = request.PermissionPolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePermissionPolicyFromAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePermissionPolicyFromAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a policy from an access configuration.
//
// Description:
//
// After you remove an inline policy from an access configuration, the policy cannot be restored.
//
// This topic provides an example on how to remove the system policy `AliyunECSFullAccess` from the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - RemovePermissionPolicyFromAccessConfigurationRequest
//
// @return RemovePermissionPolicyFromAccessConfigurationResponse
func (client *Client) RemovePermissionPolicyFromAccessConfiguration(request *RemovePermissionPolicyFromAccessConfigurationRequest) (_result *RemovePermissionPolicyFromAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePermissionPolicyFromAccessConfigurationResponse{}
	_body, _err := client.RemovePermissionPolicyFromAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a user from a group.
//
// Description:
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot remove a user from a group that is synchronized by using SCIM.
//
// This topic provides an example on how to remove the user `u-00q8wbq42wiltcrk****` from the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - RemoveUserFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a group.
//
// Description:
//
// If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot remove a user from a group that is synchronized by using SCIM.
//
// This topic provides an example on how to remove the user `u-00q8wbq42wiltcrk****` from the group `g-00jqzghi2n3o5hkh****`.
//
// @param request - RemoveUserFromGroupRequest
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a user.
//
// Description:
//
// If a user forgets the password, the password expires, or the password poses security risks, a CloudSSO administrator can reset the password for the user.
//
// > After you enable single sign-on (SSO) logon, the password of a user cannot be reset.
//
// This topic provides an example on how to reset the password of the user `u-00q8wbq42wiltcrk****`. The new password is automatically generated by the system.
//
// @param request - ResetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GenerateRandomPassword) {
		query["GenerateRandomPassword"] = request.GenerateRandomPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RequirePasswordResetForNextLogin) {
		query["RequirePasswordResetForNextLogin"] = request.RequirePasswordResetForNextLogin
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserPassword"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a user.
//
// Description:
//
// If a user forgets the password, the password expires, or the password poses security risks, a CloudSSO administrator can reset the password for the user.
//
// > After you enable single sign-on (SSO) logon, the password of a user cannot be reset.
//
// This topic provides an example on how to reset the password of the user `u-00q8wbq42wiltcrk****`. The new password is automatically generated by the system.
//
// @param request - ResetUserPasswordRequest
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries a Resource Access Management (RAM) user provisioning event.
//
// @param request - RetryUserProvisioningEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryUserProvisioningEventResponse
func (client *Client) RetryUserProvisioningEventWithOptions(request *RetryUserProvisioningEventRequest, runtime *dara.RuntimeOptions) (_result *RetryUserProvisioningEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.DuplicationStrategy) {
		query["DuplicationStrategy"] = request.DuplicationStrategy
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryUserProvisioningEvent"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryUserProvisioningEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries a Resource Access Management (RAM) user provisioning event.
//
// @param request - RetryUserProvisioningEventRequest
//
// @return RetryUserProvisioningEventResponse
func (client *Client) RetryUserProvisioningEvent(request *RetryUserProvisioningEventRequest) (_result *RetryUserProvisioningEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryUserProvisioningEventResponse{}
	_body, _err := client.RetryUserProvisioningEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is a service provider (SP), and the identity management system of an enterprise is an IdP.
//
// You can use one of the following methods to configure a SAML IdP. You can obtain the required metadata file or parameter values from your IdP.
//
//   - Use the metadata file: You can specify the `EncodedMetadataDocument` parameter to upload the metadata file.
//
//   - Manually configure the IdP: You can manually specify the following parameters for your IdP: `EntityId`, `LoginUrl`, `WantRequestSigned`, and `X509Certificate`.
//
// If you have configured a SAML IdP, the existing configurations are replaced after you call this operation.
//
//   - If the IdP is configured by using the metadata file, all existing configurations are replaced with new configurations.
//
//   - If the IdP is manually configured, the original parameter values that are different from the new parameter values are replaced.
//
// >  If SSO logon is enabled, new configurations immediately take effect. Take note of the impacts on the production environment.
//
// This topic provides an example on how to configure an IdP by using the metadata file within the directory `d-00fc2p61****`.
//
// @param request - SetExternalSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetExternalSAMLIdentityProviderResponse
func (client *Client) SetExternalSAMLIdentityProviderWithOptions(request *SetExternalSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *SetExternalSAMLIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindingType) {
		query["BindingType"] = request.BindingType
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EncodedMetadataDocument) {
		query["EncodedMetadataDocument"] = request.EncodedMetadataDocument
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.LoginUrl) {
		query["LoginUrl"] = request.LoginUrl
	}

	if !dara.IsNil(request.SSOStatus) {
		query["SSOStatus"] = request.SSOStatus
	}

	if !dara.IsNil(request.WantRequestSigned) {
		query["WantRequestSigned"] = request.WantRequestSigned
	}

	if !dara.IsNil(request.X509Certificate) {
		query["X509Certificate"] = request.X509Certificate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetExternalSAMLIdentityProvider"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a Security Assertion Markup Language (SAML) identity provider (IdP).
//
// Description:
//
// During SAML 2.0-based single sign-on (SSO) logon, CloudSSO is a service provider (SP), and the identity management system of an enterprise is an IdP.
//
// You can use one of the following methods to configure a SAML IdP. You can obtain the required metadata file or parameter values from your IdP.
//
//   - Use the metadata file: You can specify the `EncodedMetadataDocument` parameter to upload the metadata file.
//
//   - Manually configure the IdP: You can manually specify the following parameters for your IdP: `EntityId`, `LoginUrl`, `WantRequestSigned`, and `X509Certificate`.
//
// If you have configured a SAML IdP, the existing configurations are replaced after you call this operation.
//
//   - If the IdP is configured by using the metadata file, all existing configurations are replaced with new configurations.
//
//   - If the IdP is manually configured, the original parameter values that are different from the new parameter values are replaced.
//
// >  If SSO logon is enabled, new configurations immediately take effect. Take note of the impacts on the production environment.
//
// This topic provides an example on how to configure an IdP by using the metadata file within the directory `d-00fc2p61****`.
//
// @param request - SetExternalSAMLIdentityProviderRequest
//
// @return SetExternalSAMLIdentityProviderResponse
func (client *Client) SetExternalSAMLIdentityProvider(request *SetExternalSAMLIdentityProviderRequest) (_result *SetExternalSAMLIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetExternalSAMLIdentityProviderResponse{}
	_body, _err := client.SetExternalSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the logon preference of CloudSSO users.
//
// @param request - SetLoginPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoginPreferenceResponse
func (client *Client) SetLoginPreferenceWithOptions(request *SetLoginPreferenceRequest, runtime *dara.RuntimeOptions) (_result *SetLoginPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowUserToGetCredentials) {
		query["AllowUserToGetCredentials"] = request.AllowUserToGetCredentials
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.LoginNetworkMasks) {
		query["LoginNetworkMasks"] = request.LoginNetworkMasks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoginPreference"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the logon preference of CloudSSO users.
//
// @param request - SetLoginPreferenceRequest
//
// @return SetLoginPreferenceResponse
func (client *Client) SetLoginPreference(request *SetLoginPreferenceRequest) (_result *SetLoginPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoginPreferenceResponse{}
	_body, _err := client.SetLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables multi-factor authentication (MFA) for users in a directory.
//
// Description:
//
// If a CloudSSO administrator enables username-password logon for users, CloudSSO automatically enables MFA to improve security. The administrator can call this operation to enable or disable MFA based on the business requirements.
//
// This topic provides an example on how to enable MFA for users.
//
// @param request - SetMFAAuthenticationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMFAAuthenticationStatusResponse
func (client *Client) SetMFAAuthenticationStatusWithOptions(request *SetMFAAuthenticationStatusRequest, runtime *dara.RuntimeOptions) (_result *SetMFAAuthenticationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MFAAuthenticationStatus) {
		query["MFAAuthenticationStatus"] = request.MFAAuthenticationStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetMFAAuthenticationStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetMFAAuthenticationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables multi-factor authentication (MFA) for users in a directory.
//
// Description:
//
// If a CloudSSO administrator enables username-password logon for users, CloudSSO automatically enables MFA to improve security. The administrator can call this operation to enable or disable MFA based on the business requirements.
//
// This topic provides an example on how to enable MFA for users.
//
// @param request - SetMFAAuthenticationStatusRequest
//
// @return SetMFAAuthenticationStatusResponse
func (client *Client) SetMFAAuthenticationStatus(request *SetMFAAuthenticationStatusRequest) (_result *SetMFAAuthenticationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetMFAAuthenticationStatusResponse{}
	_body, _err := client.SetMFAAuthenticationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password policy for CloudSSO users.
//
// @param request - SetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicyWithOptions(request *SetPasswordPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MaxLoginAttempts) {
		query["MaxLoginAttempts"] = request.MaxLoginAttempts
	}

	if !dara.IsNil(request.MaxPasswordAge) {
		query["MaxPasswordAge"] = request.MaxPasswordAge
	}

	if !dara.IsNil(request.MinPasswordDifferentChars) {
		query["MinPasswordDifferentChars"] = request.MinPasswordDifferentChars
	}

	if !dara.IsNil(request.MinPasswordLength) {
		query["MinPasswordLength"] = request.MinPasswordLength
	}

	if !dara.IsNil(request.PasswordNotContainUsername) {
		query["PasswordNotContainUsername"] = request.PasswordNotContainUsername
	}

	if !dara.IsNil(request.PasswordReusePrevention) {
		query["PasswordReusePrevention"] = request.PasswordReusePrevention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordPolicy"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password policy for CloudSSO users.
//
// @param request - SetPasswordPolicyRequest
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (_result *SetPasswordPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.SetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables System for Cross-domain Identity Management (SCIM) synchronization.
//
// Description:
//
// You can synchronize users or groups from an external identity provider (IdP) that supports SCIM 2.0 to CloudSSO only after SCIM synchronization is enabled. If you disable SCIM synchronization, you can no longer synchronize users or groups to CloudSSO. The following list describes the impacts after SCIM synchronization is enabled or disabled:
//
//   - After you enable SCIM synchronization, you cannot modify or delete the users or groups that are synchronized to CloudSSO by using SCIM. In addition, you cannot add users to or remove users from the groups. However, you can change the passwords of the users and enable or disable the logon of the users.
//
//   - After you disable SCIM synchronization, you can modify and delete the users and groups that are synchronized to CloudSSO by using SCIM. You can also add users to or remove users from the groups.
//
// This topic provides an example on how to enable SCIM synchronization within the directory `d-00fc2p61****`.
//
// @param request - SetSCIMSynchronizationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSCIMSynchronizationStatusResponse
func (client *Client) SetSCIMSynchronizationStatusWithOptions(request *SetSCIMSynchronizationStatusRequest, runtime *dara.RuntimeOptions) (_result *SetSCIMSynchronizationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.SCIMSynchronizationStatus) {
		query["SCIMSynchronizationStatus"] = request.SCIMSynchronizationStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSCIMSynchronizationStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSCIMSynchronizationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables System for Cross-domain Identity Management (SCIM) synchronization.
//
// Description:
//
// You can synchronize users or groups from an external identity provider (IdP) that supports SCIM 2.0 to CloudSSO only after SCIM synchronization is enabled. If you disable SCIM synchronization, you can no longer synchronize users or groups to CloudSSO. The following list describes the impacts after SCIM synchronization is enabled or disabled:
//
//   - After you enable SCIM synchronization, you cannot modify or delete the users or groups that are synchronized to CloudSSO by using SCIM. In addition, you cannot add users to or remove users from the groups. However, you can change the passwords of the users and enable or disable the logon of the users.
//
//   - After you disable SCIM synchronization, you can modify and delete the users and groups that are synchronized to CloudSSO by using SCIM. You can also add users to or remove users from the groups.
//
// This topic provides an example on how to enable SCIM synchronization within the directory `d-00fc2p61****`.
//
// @param request - SetSCIMSynchronizationStatusRequest
//
// @return SetSCIMSynchronizationStatusResponse
func (client *Client) SetSCIMSynchronizationStatus(request *SetSCIMSynchronizationStatusRequest) (_result *SetSCIMSynchronizationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSCIMSynchronizationStatusResponse{}
	_body, _err := client.SetSCIMSynchronizationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about an access configuration.
//
// Description:
//
// You can modify the `Description`, `SessionDuration`, and `RelayState` parameters for an access configuration.
//
// This topic provides an example on how to change the initial web page in the access configuration `ac-00jhtfl8thteu6uj****` to `https://cloudsso.console.aliyun.com`.
//
// @param request - UpdateAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessConfigurationResponse
func (client *Client) UpdateAccessConfigurationWithOptions(request *UpdateAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewRelayState) {
		query["NewRelayState"] = request.NewRelayState
	}

	if !dara.IsNil(request.NewSessionDuration) {
		query["NewSessionDuration"] = request.NewSessionDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about an access configuration.
//
// Description:
//
// You can modify the `Description`, `SessionDuration`, and `RelayState` parameters for an access configuration.
//
// This topic provides an example on how to change the initial web page in the access configuration `ac-00jhtfl8thteu6uj****` to `https://cloudsso.console.aliyun.com`.
//
// @param request - UpdateAccessConfigurationRequest
//
// @return UpdateAccessConfigurationResponse
func (client *Client) UpdateAccessConfiguration(request *UpdateAccessConfigurationRequest) (_result *UpdateAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccessConfigurationResponse{}
	_body, _err := client.UpdateAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attribute passing settings for a specified directory, allowing you to set the SourceIdentity pass-through mode to IdP, UserName, or Disabled.
//
// Description:
//
// You must have the cloudsso:UpdateAttributePassingSetting permission to call this operation. If the SourceIdentityPassing request parameter is not specified, the existing value is retained. If an invalid enum value is specified, the InvalidParameter.SourceIdentityPassing error is returned.
//
// @param request - UpdateAttributePassingSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAttributePassingSettingResponse
func (client *Client) UpdateAttributePassingSettingWithOptions(request *UpdateAttributePassingSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateAttributePassingSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.SourceIdentityPassing) {
		query["SourceIdentityPassing"] = request.SourceIdentityPassing
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAttributePassingSetting"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAttributePassingSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attribute passing settings for a specified directory, allowing you to set the SourceIdentity pass-through mode to IdP, UserName, or Disabled.
//
// Description:
//
// You must have the cloudsso:UpdateAttributePassingSetting permission to call this operation. If the SourceIdentityPassing request parameter is not specified, the existing value is retained. If an invalid enum value is specified, the InvalidParameter.SourceIdentityPassing error is returned.
//
// @param request - UpdateAttributePassingSettingRequest
//
// @return UpdateAttributePassingSettingResponse
func (client *Client) UpdateAttributePassingSetting(request *UpdateAttributePassingSettingRequest) (_result *UpdateAttributePassingSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAttributePassingSettingResponse{}
	_body, _err := client.UpdateAttributePassingSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a directory.
//
// Description:
//
// After you change the name of a directory, the URL that is used to log on to the CloudSSO user portal is changed. You must notify the CloudSSO users of the correct URL.
//
// This topic provides an example on how to change the name of a directory to `new-example`.
//
// @param request - UpdateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDirectoryResponse
func (client *Client) UpdateDirectoryWithOptions(request *UpdateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewDirectoryName) {
		query["NewDirectoryName"] = request.NewDirectoryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDirectory"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a directory.
//
// Description:
//
// After you change the name of a directory, the URL that is used to log on to the CloudSSO user portal is changed. You must notify the CloudSSO users of the correct URL.
//
// This topic provides an example on how to change the name of a directory to `new-example`.
//
// @param request - UpdateDirectoryRequest
//
// @return UpdateDirectoryResponse
func (client *Client) UpdateDirectory(request *UpdateDirectoryRequest) (_result *UpdateDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDirectoryResponse{}
	_body, _err := client.UpdateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a group.
//
// Description:
//
// You can modify `GroupName` and `Description` for a group.
//
// > If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify information about a group that is synchronized by using SCIM.
//
// This topic provides an example on how to modify the name of the group `g-00jqzghi2n3o5hkh****` to `NewTestGroup`.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewGroupName) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a group.
//
// Description:
//
// You can modify `GroupName` and `Description` for a group.
//
// > If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify information about a group that is synchronized by using SCIM.
//
// This topic provides an example on how to modify the name of the group `g-00jqzghi2n3o5hkh****` to `NewTestGroup`.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an inline policy that is created for an access configuration.
//
// Description:
//
// This topic provides an example on how to modify an inline policy that is created for the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - UpdateInlinePolicyForAccessConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInlinePolicyForAccessConfigurationResponse
func (client *Client) UpdateInlinePolicyForAccessConfigurationWithOptions(request *UpdateInlinePolicyForAccessConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateInlinePolicyForAccessConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessConfigurationId) {
		query["AccessConfigurationId"] = request.AccessConfigurationId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.InlinePolicyName) {
		query["InlinePolicyName"] = request.InlinePolicyName
	}

	if !dara.IsNil(request.NewInlinePolicyDocument) {
		query["NewInlinePolicyDocument"] = request.NewInlinePolicyDocument
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInlinePolicyForAccessConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInlinePolicyForAccessConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an inline policy that is created for an access configuration.
//
// Description:
//
// This topic provides an example on how to modify an inline policy that is created for the access configuration `ac-00jhtfl8thteu6uj****`.
//
// @param request - UpdateInlinePolicyForAccessConfigurationRequest
//
// @return UpdateInlinePolicyForAccessConfigurationResponse
func (client *Client) UpdateInlinePolicyForAccessConfiguration(request *UpdateInlinePolicyForAccessConfigurationRequest) (_result *UpdateInlinePolicyForAccessConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInlinePolicyForAccessConfigurationResponse{}
	_body, _err := client.UpdateInlinePolicyForAccessConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the global multi-factor authentication (MFA) settings.
//
// Description:
//
// When username-password logon is enabled, you can configure the global MFA verification policy for user logon.
//
// This topic provides an example on how to enable MFA verification for all CloudSSO users in the directory `d-00fc2p61****`.
//
// @param tmpReq - UpdateMFAAuthenticationSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMFAAuthenticationSettingsResponse
func (client *Client) UpdateMFAAuthenticationSettingsWithOptions(tmpReq *UpdateMFAAuthenticationSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateMFAAuthenticationSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMFAAuthenticationSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedVerificationTypes) {
		request.AllowedVerificationTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedVerificationTypes, dara.String("AllowedVerificationTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowedVerificationTypesShrink) {
		query["AllowedVerificationTypes"] = request.AllowedVerificationTypesShrink
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.MFAAuthenticationSettings) {
		query["MFAAuthenticationSettings"] = request.MFAAuthenticationSettings
	}

	if !dara.IsNil(request.OperationForRiskLogin) {
		query["OperationForRiskLogin"] = request.OperationForRiskLogin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMFAAuthenticationSettings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the global multi-factor authentication (MFA) settings.
//
// Description:
//
// When username-password logon is enabled, you can configure the global MFA verification policy for user logon.
//
// This topic provides an example on how to enable MFA verification for all CloudSSO users in the directory `d-00fc2p61****`.
//
// @param request - UpdateMFAAuthenticationSettingsRequest
//
// @return UpdateMFAAuthenticationSettingsResponse
func (client *Client) UpdateMFAAuthenticationSettings(request *UpdateMFAAuthenticationSettingsRequest) (_result *UpdateMFAAuthenticationSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMFAAuthenticationSettingsResponse{}
	_body, _err := client.UpdateMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// This topic provides an example on how to disable the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`. After the SCIM credential is disabled, the synchronization task that uses the SCIM credential fails. You can call this operation again to enable the SCIM credential.
//
// @param request - UpdateSCIMServerCredentialStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSCIMServerCredentialStatusResponse
func (client *Client) UpdateSCIMServerCredentialStatusWithOptions(request *UpdateSCIMServerCredentialStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateSCIMServerCredentialStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewStatus) {
		query["NewStatus"] = request.NewStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSCIMServerCredentialStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSCIMServerCredentialStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a System for Cross-domain Identity Management (SCIM) credential.
//
// Description:
//
// This topic provides an example on how to disable the SCIM credential whose ID is `scimcred-004whl0kvfwcypbi****`. After the SCIM credential is disabled, the synchronization task that uses the SCIM credential fails. You can call this operation again to enable the SCIM credential.
//
// @param request - UpdateSCIMServerCredentialStatusRequest
//
// @return UpdateSCIMServerCredentialStatusResponse
func (client *Client) UpdateSCIMServerCredentialStatus(request *UpdateSCIMServerCredentialStatusRequest) (_result *UpdateSCIMServerCredentialStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSCIMServerCredentialStatusResponse{}
	_body, _err := client.UpdateSCIMServerCredentialStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a user.
//
// Description:
//
// You can modify `FirstName`, `LastName`, `DisplayName`, `Email`, and `Description` for a user. You cannot modify `UserName` for a user.
//
// > If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify information about a user that is synchronized by using SCIM.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !dara.IsNil(request.NewEmail) {
		query["NewEmail"] = request.NewEmail
	}

	if !dara.IsNil(request.NewFirstName) {
		query["NewFirstName"] = request.NewFirstName
	}

	if !dara.IsNil(request.NewLastName) {
		query["NewLastName"] = request.NewLastName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a user.
//
// Description:
//
// You can modify `FirstName`, `LastName`, `DisplayName`, `Email`, and `Description` for a user. You cannot modify `UserName` for a user.
//
// > If System for Cross-domain Identity Management (SCIM) synchronization is enabled, you cannot modify information about a user that is synchronized by using SCIM.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the multi-factor authentication (MFA) setting of a single user.
//
// Description:
//
// If you call the [UpdateMFAAuthenticationSettings](https://help.aliyun.com/document_detail/450134.html) operation to set the MFAAuthenticationSettings parameter to `Byuser`, user-specific settings are applied. Then, you must call the UpdateUserMFAAuthenticationSettings operation to configure MFA for each user.
//
// By default, the MFAAuthenticationSettings parameter is set to `Enabled` for a new user.
//
// This topic provides an example on how to enable MFA for the user named `u-00q8wbq42wiltcrk****`.
//
// @param request - UpdateUserMFAAuthenticationSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserMFAAuthenticationSettingsResponse
func (client *Client) UpdateUserMFAAuthenticationSettingsWithOptions(request *UpdateUserMFAAuthenticationSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserMFAAuthenticationSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserMFAAuthenticationSettings) {
		query["UserMFAAuthenticationSettings"] = request.UserMFAAuthenticationSettings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserMFAAuthenticationSettings"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the multi-factor authentication (MFA) setting of a single user.
//
// Description:
//
// If you call the [UpdateMFAAuthenticationSettings](https://help.aliyun.com/document_detail/450134.html) operation to set the MFAAuthenticationSettings parameter to `Byuser`, user-specific settings are applied. Then, you must call the UpdateUserMFAAuthenticationSettings operation to configure MFA for each user.
//
// By default, the MFAAuthenticationSettings parameter is set to `Enabled` for a new user.
//
// This topic provides an example on how to enable MFA for the user named `u-00q8wbq42wiltcrk****`.
//
// @param request - UpdateUserMFAAuthenticationSettingsRequest
//
// @return UpdateUserMFAAuthenticationSettingsResponse
func (client *Client) UpdateUserMFAAuthenticationSettings(request *UpdateUserMFAAuthenticationSettingsRequest) (_result *UpdateUserMFAAuthenticationSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserMFAAuthenticationSettingsResponse{}
	_body, _err := client.UpdateUserMFAAuthenticationSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Resource Access Management (RAM) user provisioning.
//
// @param request - UpdateUserProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserProvisioningResponse
func (client *Client) UpdateUserProvisioningWithOptions(request *UpdateUserProvisioningRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewDeletionStrategy) {
		query["NewDeletionStrategy"] = request.NewDeletionStrategy
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewDuplicationStrategy) {
		query["NewDuplicationStrategy"] = request.NewDuplicationStrategy
	}

	if !dara.IsNil(request.UserProvisioningId) {
		query["UserProvisioningId"] = request.UserProvisioningId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserProvisioning"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Resource Access Management (RAM) user provisioning.
//
// @param request - UpdateUserProvisioningRequest
//
// @return UpdateUserProvisioningResponse
func (client *Client) UpdateUserProvisioning(request *UpdateUserProvisioningRequest) (_result *UpdateUserProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserProvisioningResponse{}
	_body, _err := client.UpdateUserProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the global configurations of a Resource Access Management (RAM) user provisioning.
//
// @param request - UpdateUserProvisioningConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserProvisioningConfigurationResponse
func (client *Client) UpdateUserProvisioningConfigurationWithOptions(request *UpdateUserProvisioningConfigurationRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserProvisioningConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewDefaultLandingPage) {
		query["NewDefaultLandingPage"] = request.NewDefaultLandingPage
	}

	if !dara.IsNil(request.NewSessionDuration) {
		query["NewSessionDuration"] = request.NewSessionDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserProvisioningConfiguration"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserProvisioningConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the global configurations of a Resource Access Management (RAM) user provisioning.
//
// @param request - UpdateUserProvisioningConfigurationRequest
//
// @return UpdateUserProvisioningConfigurationResponse
func (client *Client) UpdateUserProvisioningConfiguration(request *UpdateUserProvisioningConfigurationRequest) (_result *UpdateUserProvisioningConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserProvisioningConfigurationResponse{}
	_body, _err := client.UpdateUserProvisioningConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a user.
//
// Description:
//
// This topic provides an example on how to change the status of the user whose ID is `u-00q8wbq42wiltcrk****` to Disabled. Users in the Disabled state cannot log on to the CloudSSO user portal.
//
// @param request - UpdateUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserStatusResponse
func (client *Client) UpdateUserStatusWithOptions(request *UpdateUserStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.NewStatus) {
		query["NewStatus"] = request.NewStatus
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserStatus"),
		Version:     dara.String("2021-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a user.
//
// Description:
//
// This topic provides an example on how to change the status of the user whose ID is `u-00q8wbq42wiltcrk****` to Disabled. Users in the Disabled state cannot log on to the CloudSSO user portal.
//
// @param request - UpdateUserStatusRequest
//
// @return UpdateUserStatusResponse
func (client *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (_result *UpdateUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserStatusResponse{}
	_body, _err := client.UpdateUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
