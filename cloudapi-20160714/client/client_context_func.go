// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Unpublishes a specified API from a specified runtime environment.
//
// Description:
//
//	  This operation is intended for API providers and is the opposite of DeployApi.
//
//		- An API can be unpublished from a specified runtime environment in under 5 seconds.
//
//		- An unpublished API cannot be called in the specified runtime environment.
//
// @param request - AbolishApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishApiResponse
func (client *Client) AbolishApiWithContext(ctx context.Context, request *AbolishApiRequest, runtime *dara.RuntimeOptions) (_result *AbolishApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbolishApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbolishApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Adds an IP address entry to the access control polocy of an instance.
//
// @param request - AddAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAccessControlListEntryResponse
func (client *Client) AddAccessControlListEntryWithContext(ctx context.Context, request *AddAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *AddAccessControlListEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAccessControlListEntry"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAccessControlListEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a policy to an existing ACL.
//
// Description:
//
// When you call this operation, note that:
//
//   - This operation is intended for API providers.
//
//   - An added policy immediately takes effect on all APIs that are bound to the access control list (ACL).
//
//   - A maximum of 100 policies can be added to an ACL.
//
// @param request - AddIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpControlPolicyItemResponse
func (client *Client) AddIpControlPolicyItemWithContext(ctx context.Context, request *AddIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *AddIpControlPolicyItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CidrIp) {
		query["CidrIp"] = request.CidrIp
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpControlPolicyItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom special policy to a specified throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the input SpecialKey already exists, the previous configuration is overwritten. Use caution when calling this operation.
//
//		- Special throttling policies must be added to an existing throttling policy, and can take effect on all the APIs to which the throttling policy is bound.
//
// @param request - AddTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrafficSpecialControlResponse
func (client *Client) AddTrafficSpecialControlWithContext(ctx context.Context, request *AddTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *AddTrafficSpecialControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecialKey) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !dara.IsNil(request.SpecialType) {
		query["SpecialType"] = request.SpecialType
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficValue) {
		query["TrafficValue"] = request.TrafficValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTrafficSpecialControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an internal domain name resolution with a dedicated instance.
//
// Description:
//
// An internal domain name resolution can be associated only with a dedicated instance, not with a shared instance or shared instance cluster.
//
// @param tmpReq - AssociateInstanceWithPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateInstanceWithPrivateDNSResponse
func (client *Client) AssociateInstanceWithPrivateDNSWithContext(ctx context.Context, tmpReq *AssociateInstanceWithPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *AssociateInstanceWithPrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssociateInstanceWithPrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntranetDomains) {
		request.IntranetDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntranetDomains, dara.String("IntranetDomains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomainsShrink) {
		body["IntranetDomains"] = request.IntranetDomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateInstanceWithPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches APIs to an API product. If the API product does not exist, the system automatically creates the API product.
//
// @param request - AttachApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApiProductResponse
func (client *Client) AttachApiProductWithContext(ctx context.Context, request *AttachApiProductRequest, runtime *dara.RuntimeOptions) (_result *AttachApiProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.Apis) {
		query["Apis"] = request.Apis
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachApiProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach plugin to API group.
//
// @param request - AttachGroupPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachGroupPluginResponse
func (client *Client) AttachGroupPluginWithContext(ctx context.Context, request *AttachGroupPluginRequest, runtime *dara.RuntimeOptions) (_result *AttachGroupPluginResponse, _err error) {
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

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachGroupPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachGroupPluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a plug-in to an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can only bind plug-ins to published APIs.
//
//		- The plug-in takes effect immediately after it is bound to an API.
//
//		- If you bind a different plug-in to an API, this plug-in takes effect immediately.
//
// @param request - AttachPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPluginResponse
func (client *Client) AttachPluginWithContext(ctx context.Context, request *AttachPluginRequest, runtime *dara.RuntimeOptions) (_result *AttachPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes multiple published APIs at a time.
//
// @param request - BatchAbolishApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAbolishApisResponse
func (client *Client) BatchAbolishApisWithContext(ctx context.Context, request *BatchAbolishApisRequest, runtime *dara.RuntimeOptions) (_result *BatchAbolishApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAbolishApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAbolishApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes multiple APIs at a time.
//
// @param request - BatchDeployApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeployApisResponse
func (client *Client) BatchDeployApisWithContext(ctx context.Context, request *BatchDeployApisRequest, runtime *dara.RuntimeOptions) (_result *BatchDeployApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeployApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeployApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Creates an Access Control List (ACL). Each user is allowed to create five ACLs in each region.
//
// @param request - CreateAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessControlListResponse
func (client *Client) CreateAccessControlListWithContext(ctx context.Context, request *CreateAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessControlListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessControlList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessControlListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of an API must be unique within an API group.
//
//		- A request path must be unique within an API group.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiResponse
func (client *Client) CreateApiWithContext(ctx context.Context, request *CreateApiRequest, runtime *dara.RuntimeOptions) (_result *CreateApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendEnable) {
		query["BackendEnable"] = request.BackendEnable
	}

	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpenIdConnectConfig) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !dara.IsNil(request.RequestConfig) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !dara.IsNil(request.ResultBodyModel) {
		query["ResultBodyModel"] = request.ResultBodyModel
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceConfig) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WebSocketApiType) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConstantParameters) {
		body["ConstantParameters"] = request.ConstantParameters
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		body["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		body["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.RequestParameters) {
		body["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.ResultDescriptions) {
		body["ResultDescriptions"] = request.ResultDescriptions
	}

	if !dara.IsNil(request.ResultSample) {
		body["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		body["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.SystemParameters) {
		body["SystemParameters"] = request.SystemParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建API分组
//
// @param request - CreateApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiGroupResponse
func (client *Client) CreateApiGroupWithContext(ctx context.Context, request *CreateApiGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateApiGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a variable to an environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - CreateApiStageVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiStageVariableResponse
func (client *Client) CreateApiStageVariableWithContext(ctx context.Context, request *CreateApiStageVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateApiStageVariableResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageId) {
		query["StageId"] = request.StageId
	}

	if !dara.IsNil(request.StageRouteModel) {
		query["StageRouteModel"] = request.StageRouteModel
	}

	if !dara.IsNil(request.SupportRoute) {
		query["SupportRoute"] = request.SupportRoute
	}

	if !dara.IsNil(request.VariableName) {
		query["VariableName"] = request.VariableName
	}

	if !dara.IsNil(request.VariableValue) {
		query["VariableValue"] = request.VariableValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiStageVariable"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiStageVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application for calling APIs in API Gateway.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- Each application has a key-value pair which is used for identity verification when you call an API.
//
//		- An application must be authorized to call an API.
//
//		- Each application has only one key-value pair, which can be reset if the pair is leaked.
//
//		- A maximum of 1,000 applications can be created for each Alibaba Cloud account.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppSecret) {
		query["AppSecret"] = request.AppSecret
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AppCode to an application.
//
// @param request - CreateAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppCodeResponse
func (client *Client) CreateAppCodeWithContext(ctx context.Context, request *CreateAppCodeRequest, runtime *dara.RuntimeOptions) (_result *CreateAppCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AppKey and AppSecret pair to an application.
//
// @param request - CreateAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppKeyResponse
func (client *Client) CreateAppKeyWithContext(ctx context.Context, request *CreateAppKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppSecret) {
		query["AppSecret"] = request.AppSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppKey"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backend service in API Gateway.
//
// @param request - CreateBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackendResponse
func (client *Client) CreateBackendWithContext(ctx context.Context, request *CreateBackendRequest, runtime *dara.RuntimeOptions) (_result *CreateBackendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.CreateEventBridgeServiceLinkedRole) {
		query["CreateEventBridgeServiceLinkedRole"] = request.CreateEventBridgeServiceLinkedRole
	}

	if !dara.IsNil(request.CreateSlr) {
		query["CreateSlr"] = request.CreateSlr
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建后端服务在环境上的配置
//
// @param request - CreateBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackendModelResponse
func (client *Client) CreateBackendModelWithContext(ctx context.Context, request *CreateBackendModelRequest, runtime *dara.RuntimeOptions) (_result *CreateBackendModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelData) {
		query["BackendModelData"] = request.BackendModelData
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackendModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom dataset.
//
// @param request - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, request *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		query["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义数据集条目
//
// @param request - CreateDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetItemResponse
func (client *Client) CreateDatasetItemWithContext(ctx context.Context, request *CreateDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API Gateway instance.
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
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.InstanceCidr) {
		query["InstanceCidr"] = request.InstanceCidr
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserVpcId) {
		query["UserVpcId"] = request.UserVpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneVSwitchSecurityGroup) {
		query["ZoneVSwitchSecurityGroup"] = request.ZoneVSwitchSecurityGroup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2016-07-14"),
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
// 创建内网域名
//
// @param request - CreateIntranetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntranetDomainResponse
func (client *Client) CreateIntranetDomainWithContext(ctx context.Context, request *CreateIntranetDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateIntranetDomainResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntranetDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntranetDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a region.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An ACL must be bound to an API to take effect. After an ACL is bound to an API, the ACL takes effect on the API immediately.
//
//		- You can add policies to an ACL when you create the ACL.
//
//		- If an ACL does not have any policy, the ACL is ineffective.
//
// @param request - CreateIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpControlResponse
func (client *Client) CreateIpControlWithContext(ctx context.Context, request *CreateIpControlRequest, runtime *dara.RuntimeOptions) (_result *CreateIpControlResponse, _err error) {
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

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.IpControlPolicys) {
		query["IpControlPolicys"] = request.IpControlPolicys
	}

	if !dara.IsNil(request.IpControlType) {
		query["IpControlType"] = request.IpControlType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Simple Log Service configuration for an API.
//
// @param request - CreateLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogConfigResponse
func (client *Client) CreateLogConfigWithContext(ctx context.Context, request *CreateLogConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSlr) {
		query["CreateSlr"] = request.CreateSlr
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlsLogStore) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model for an API group.
//
// Description:
//
//	  For more information about the model definition, see [JSON Schema Draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04?spm=a2c4g.11186623.2.10.2e977ff7p4BpQd).
//
//		- JSON Schema supports only element attributes of the Object type.
//
// @param request - CreateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables CloudMonitor alerting for a specified API group.
//
// @param request - CreateMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroupWithContext(ctx context.Context, request *CreateMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["Auth"] = request.Auth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RawMonitorGroupId) {
		query["RawMonitorGroupId"] = request.RawMonitorGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The number of plug-ins of the same type that each user can create is limited. Different limits apply to different plug-in types.
//
//		- The plug-in definitions for advanced features are restricted.
//
//		- Plug-ins must be bound to APIs to take effect. After a plug-in is bound, it takes effect on that API immediately.
//
// @param request - CreatePluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginResponse
func (client *Client) CreatePluginWithContext(ctx context.Context, request *CreatePluginRequest, runtime *dara.RuntimeOptions) (_result *CreatePluginResponse, _err error) {
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

	if !dara.IsNil(request.PluginData) {
		query["PluginData"] = request.PluginData
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal domain name resolution and adds a resolution record.
//
// Description:
//
// An internal domain name resolution of the virtual private cloud (VPC) type can be bound only to traditional dedicated instances. An internal domain name resolution of the A type can be bound only to VPC integration dedicated instances.
//
// @param tmpReq - CreatePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateDNSResponse
func (client *Client) CreatePrivateDNSWithContext(ctx context.Context, tmpReq *CreatePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Records) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, dara.String("Records"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RecordsShrink) {
		body["Records"] = request.RecordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The API operation only creates a key policy. You must call the binding operation to bind the key to an API.
//
//		- After the key is bound to the API, requests sent from API Gateway to the backend service contain signature strings. You can specify whether your backend service verifies these signature strings.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSignatureResponse
func (client *Client) CreateSignatureWithContext(ctx context.Context, request *CreateSignatureRequest, runtime *dara.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureKey) {
		query["SignatureKey"] = request.SignatureKey
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	if !dara.IsNil(request.SignatureSecret) {
		query["SignatureSecret"] = request.SignatureSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Throttling policies must be bound to APIs to take effect. After a policy is bound to an API, it goes into effect on that API immediately.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlResponse
func (client *Client) CreateTrafficControlWithContext(ctx context.Context, request *CreateTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDefault) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !dara.IsNil(request.AppDefault) {
		query["AppDefault"] = request.AppDefault
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !dara.IsNil(request.TrafficControlUnit) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !dara.IsNil(request.UserDefault) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Deletes an access control policy.
//
// @param request - DeleteAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessControlListResponse
func (client *Client) DeleteAccessControlListWithContext(ctx context.Context, request *DeleteAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessControlListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessControlList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessControlListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all custom special policies of a specified throttling policy.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DeleteAllTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllTrafficSpecialControlResponse
func (client *Client) DeleteAllTrafficSpecialControlWithContext(ctx context.Context, request *DeleteAllTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllTrafficSpecialControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllTrafficSpecialControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the definition of a specified API.
//
// Description:
//
//	  This operation is intended for API providers and cannot be undone after it is complete.
//
//		- An API that is running in the runtime environment must be unpublished before you can delete the API.****
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiResponse
func (client *Client) DeleteApiWithContext(ctx context.Context, request *DeleteApiRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An API group that contains APIs cannot be deleted. To delete the API group, you must first delete its APIs.
//
//		- After an API group is deleted, the second-level domain name bound to the API group is automatically invalidated.
//
//		- If the specified API group does not exist, a success response is returned.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiGroupResponse
func (client *Client) DeleteApiGroupWithContext(ctx context.Context, request *DeleteApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiGroupResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API product. Deleting an API product causes the association between APIs and the deleted API product to be deleted as well. Exercise caution when you delete an API product. If any API in the API product is associated with an application, the API product fails to be deleted.
//
// @param request - DeleteApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiProductResponse
func (client *Client) DeleteApiProductWithContext(ctx context.Context, request *DeleteApiProductRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified variable in a specified environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DeleteApiStageVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiStageVariableResponse
func (client *Client) DeleteApiStageVariableWithContext(ctx context.Context, request *DeleteApiStageVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiStageVariableResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageId) {
		query["StageId"] = request.StageId
	}

	if !dara.IsNil(request.VariableName) {
		query["VariableName"] = request.VariableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiStageVariable"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiStageVariableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- After an application is deleted, the application and its API authorization cannot be restored.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the AppCode of an application.
//
// @param request - DeleteAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppCodeResponse
func (client *Client) DeleteAppCodeWithContext(ctx context.Context, request *DeleteAppCodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the AppKey and AppSecret of an application.
//
// @param request - DeleteAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppKeyResponse
func (client *Client) DeleteAppKeyWithContext(ctx context.Context, request *DeleteAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppKey"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backend service.
//
// @param request - DeleteBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackendResponse
func (client *Client) DeleteBackendWithContext(ctx context.Context, request *DeleteBackendRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the definition of a backend service in an environment. After the definition is deleted, the API that uses the backend service and is published to this environment will be unpublished.
//
// @param request - DeleteBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackendModelResponse
func (client *Client) DeleteBackendModelWithContext(ctx context.Context, request *DeleteBackendModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackendModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelId) {
		query["BackendModelId"] = request.BackendModelId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackendModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义数据集
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data entry from a custom dataset.
//
// @param request - DeleteDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetItemResponse
func (client *Client) DeleteDatasetItemWithContext(ctx context.Context, request *DeleteDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a custom domain name from an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the specified domain name does not exist, a successful response will still appear.
//
//		- Unbinding a domain name from an API group will affect access to the APIs in the group. Exercise caution when using this operation.
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the SSL certificate of a specified domain name. This operation is intended for API providers. If the SSL certificate does not exist, a success response is still returned. If the specified API group does not exist, the InvalidGroupId.NotFound error is returned. Access over HTTPS is not supported after the SSL certificate is deleted. Exercise caution when using this API operation.
//
// @param request - DeleteDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainCertificateResponse
func (client *Client) DeleteDomainCertificateWithContext(ctx context.Context, request *DeleteDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainCertificateResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainCertificate"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API Gateway instance.
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2016-07-14"),
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
// Deletes an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the ACL is bound to an API, you must unbind the ACL from the API before you can delete the ACL. Otherwise, an error is returned.
//
//		- If you call this operation on an ACL that does not exist, a success message is returned.
//
// @param request - DeleteIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpControlResponse
func (client *Client) DeleteIpControlWithContext(ctx context.Context, request *DeleteIpControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified log configuration.
//
// @param request - DeleteLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogConfigResponse
func (client *Client) DeleteLogConfigWithContext(ctx context.Context, request *DeleteLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @param request - DeleteModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
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

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CloudMonitor application group corresponding to an API group.
//
// @param request - DeleteMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroupWithContext(ctx context.Context, request *DeleteMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupResponse, _err error) {
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

	if !dara.IsNil(request.RawMonitorGroupId) {
		query["RawMonitorGroupId"] = request.RawMonitorGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must first unbind the plug-in from the API. Otherwise, an error is reported when you delete the plug-in.
//
// @param request - DeletePluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginResponse
func (client *Client) DeletePluginWithContext(ctx context.Context, request *DeletePluginRequest, runtime *dara.RuntimeOptions) (_result *DeletePluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal domain name resolution.
//
// @param request - DeletePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDNSResponse
func (client *Client) DeletePrivateDNSWithContext(ctx context.Context, request *DeletePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation deletes an existing backend signature key.
//
//		- You cannot delete a key that is bound to an API. To delete the key, you must unbind it first.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSignatureResponse
func (client *Client) DeleteSignatureWithContext(ctx context.Context, request *DeleteSignatureRequest, runtime *dara.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom throttling policy and the special throttling rules in the policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the throttling policy you want to delete is bound to APIs, you need to unbind the policy first. Otherwise, an error is reported when you delete the policy.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlResponse
func (client *Client) DeleteTrafficControlWithContext(ctx context.Context, request *DeleteTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom special throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can obtain the input parameters required in this operation by calling other APIs.
//
// @param request - DeleteTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficSpecialControlResponse
func (client *Client) DeleteTrafficSpecialControlWithContext(ctx context.Context, request *DeleteTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficSpecialControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecialKey) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !dara.IsNil(request.SpecialType) {
		query["SpecialType"] = request.SpecialType
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficSpecialControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an API to an environment.
//
// Description:
//
//	  This operation is intended for API providers. Only the API that you have defined and published to a runtime environment can be called.
//
//		- An API is published to a cluster in under 5 seconds.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeployApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApiResponse
func (client *Client) DeployApiWithContext(ctx context.Context, request *DeployApiRequest, runtime *dara.RuntimeOptions) (_result *DeployApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量下线API任务
//
// @param request - DescribeAbolishApiTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAbolishApiTaskResponse
func (client *Client) DescribeAbolishApiTaskWithContext(ctx context.Context, request *DescribeAbolishApiTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeAbolishApiTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAbolishApiTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAbolishApiTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries the details of an access control policy.
//
// @param request - DescribeAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListAttributeResponse
func (client *Client) DescribeAccessControlListAttributeWithContext(ctx context.Context, request *DescribeAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlListAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessControlListAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries access control policies.
//
// @param request - DescribeAccessControlListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListsResponse
func (client *Client) DescribeAccessControlListsWithContext(ctx context.Context, request *DescribeAccessControlListsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlLists"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessControlListsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the definition of an API.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiResponse
func (client *Client) DescribeApiWithContext(ctx context.Context, request *DescribeApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the documentation of an API.
//
// Description:
//
//	  For API callers, the specified API must be a public or authorized private API that has been published to a runtime environment.
//
//		- When you call this operation as an API caller, the service information, parameter definitions, and other details of the API you specify are returned.
//
//		- When you call this operation as an API provider, the definition of the specified API running in the specified runtime environment is returned. The returned definition takes effect in the runtime environment, and may be different from the definition of the API you modify.
//
//		- Before you call this operation as an API provider, ensure that the API to be queried is a public one or that your application has been authorized to call the API, because authentication on API callers is required.
//
// @param request - DescribeApiDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiDocResponse
func (client *Client) DescribeApiDocWithContext(ctx context.Context, request *DescribeApiDocRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiDocResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiDoc"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiDocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query details about an API group, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupResponse
func (client *Client) DescribeApiGroupWithContext(ctx context.Context, request *DescribeApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the VPC whitelist that is allowed to access an API group.
//
// @param request - DescribeApiGroupVpcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupVpcWhitelistResponse
func (client *Client) DescribeApiGroupVpcWhitelistWithContext(ctx context.Context, request *DescribeApiGroupVpcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupVpcWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroupVpcWhitelist"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries existing API groups and their basic information.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupsResponse
func (client *Client) DescribeApiGroupsWithContext(ctx context.Context, request *DescribeApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical versions of a specified API.
//
// Description:
//
//	  This operation is intended for API providers. Only APIs that have been published have historical version records.
//
//		- This operation allows you to obtain the historical versions of an API. This operation is always called by other operations.
//
// @param request - DescribeApiHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiHistoriesResponse
func (client *Client) DescribeApiHistoriesWithContext(ctx context.Context, request *DescribeApiHistoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiHistoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiHistories"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiHistoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified historical version of a specified API definition.
//
// Description:
//
// Queries the details of a specified historical version of a specified API definition.
//
//   - This API is intended for API providers.
//
//   - API Gateway records the time and definition of an API every time the API is published. You can use the version number obtained from other operations to query definition details at a certain publication.
//
// @param request - DescribeApiHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiHistoryResponse
func (client *Client) DescribeApiHistoryWithContext(ctx context.Context, request *DescribeApiHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryVersion) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiHistory"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) that are bound to all the APIs in an API group in a specified environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- If an optional parameter is not specified, all results are returned on separate pages.
//
// ·
//
// @param request - DescribeApiIpControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiIpControlsResponse
func (client *Client) DescribeApiIpControlsWithContext(ctx context.Context, request *DescribeApiIpControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiIpControlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiIpControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiIpControlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the response time statistics of an API.
//
// Description:
//
// You can call this operation to query the latency metrics in milliseconds for a specified API.
//
//   - This API is intended for API providers.
//
//   - Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiLatencyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiLatencyDataResponse
func (client *Client) DescribeApiLatencyDataWithContext(ctx context.Context, request *DescribeApiLatencyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiLatencyDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiLatencyData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiLatencyDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud Marketplace attributes of an API.
//
// @param request - DescribeApiMarketAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiMarketAttributesResponse
func (client *Client) DescribeApiMarketAttributesWithContext(ctx context.Context, request *DescribeApiMarketAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiMarketAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiMarketAttributes"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiMarketAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attached APIs of an API product.
//
// @param request - DescribeApiProductApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiProductApisResponse
func (client *Client) DescribeApiProductApisWithContext(ctx context.Context, request *DescribeApiProductApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiProductApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiProductApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiProductApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API products by application.
//
// @param request - DescribeApiProductsByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiProductsByAppResponse
func (client *Client) DescribeApiProductsByAppWithContext(ctx context.Context, request *DescribeApiProductsByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiProductsByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiProductsByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiProductsByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the QPS statistics of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiQpsDataResponse
func (client *Client) DescribeApiQpsDataWithContext(ctx context.Context, request *DescribeApiQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiQpsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiQpsData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiQpsDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to the APIs of a specified API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiSignaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiSignaturesResponse
func (client *Client) DescribeApiSignaturesWithContext(ctx context.Context, request *DescribeApiSignaturesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiSignaturesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiSignatures"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiSignaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling policies bound to all members of an API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiTrafficControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiTrafficControlsResponse
func (client *Client) DescribeApiTrafficControlsWithContext(ctx context.Context, request *DescribeApiTrafficControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiTrafficControlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiTrafficControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiTrafficControlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the traffic of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiTrafficDataResponse
func (client *Client) DescribeApiTrafficDataWithContext(ctx context.Context, request *DescribeApiTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiTrafficDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiTrafficData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiTrafficDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of APIs that are being defined.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation returns a list of all APIs that are being defined. The basic information about these APIs is also returned in the list.
//
//		- This operation returns all APIs that are being edited, regardless of their environments. The returned definitions may be different from the definitions in the environments.
//
// @param request - DescribeApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisResponse
func (client *Client) DescribeApisWithContext(ctx context.Context, request *DescribeApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiMethod) {
		query["ApiMethod"] = request.ApiMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiPath) {
		query["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.CatalogId) {
		query["CatalogId"] = request.CatalogId
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnDeployed) {
		query["UnDeployed"] = request.UnDeployed
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs with which an application is associated.
//
// @param request - DescribeApisByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByAppResponse
func (client *Client) DescribeApisByAppWithContext(ctx context.Context, request *DescribeApisByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiUid) {
		query["ApiUid"] = request.ApiUid
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs in the draft or published state that are created by using a specified backend service.
//
// @param request - DescribeApisByBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByBackendResponse
func (client *Client) DescribeApisByBackendWithContext(ctx context.Context, request *DescribeApisByBackendRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByBackendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByBackendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that are bound to an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByIpControlResponse
func (client *Client) DescribeApisByIpControlWithContext(ctx context.Context, request *DescribeApisByIpControlRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByIpControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByIpControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified backend signature key is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The results are returned on separate pages. You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisBySignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisBySignatureResponse
func (client *Client) DescribeApisBySignatureWithContext(ctx context.Context, request *DescribeApisBySignatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisBySignatureResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisBySignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisBySignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified throttling policy is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByTrafficControlResponse
func (client *Client) DescribeApisByTrafficControlWithContext(ctx context.Context, request *DescribeApisByTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByTrafficControlResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByTrafficControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that are associated with a virtual private cloud (VPC) access authorization in a region.
//
// @param request - DescribeApisByVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByVpcAccessResponse
func (client *Client) DescribeApisByVpcAccessWithContext(ctx context.Context, request *DescribeApisByVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByVpcAccessResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcName) {
		query["VpcName"] = request.VpcName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByVpcAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs by application. The environment information is also returned.
//
// @param request - DescribeApisWithStageNameIntegratedByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisWithStageNameIntegratedByAppResponse
func (client *Client) DescribeApisWithStageNameIntegratedByAppWithContext(ctx context.Context, request *DescribeApisWithStageNameIntegratedByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisWithStageNameIntegratedByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiUid) {
		query["ApiUid"] = request.ApiUid
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisWithStageNameIntegratedByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisWithStageNameIntegratedByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the apps that can be authorized.
//
// @param request - DescribeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppResponse
func (client *Client) DescribeAppWithContext(ctx context.Context, request *DescribeAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries apps and their basic information.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppId is optional.
//
// @param request - DescribeAppAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAttributesResponse
func (client *Client) DescribeAppAttributesWithContext(ctx context.Context, request *DescribeAppAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAttributes"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the key-related information of an application.
//
// @param request - DescribeAppSecuritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppSecuritiesResponse
func (client *Client) DescribeAppSecuritiesWithContext(ctx context.Context, request *DescribeAppSecuritiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppSecuritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppSecurities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppSecuritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This key is used for authentication when an API call is made.
//
// Description:
//
//	This operation is intended for API callers.
//
// @param request - DescribeAppSecurityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppSecurityResponse
func (client *Client) DescribeAppSecurityWithContext(ctx context.Context, request *DescribeAppSecurityRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppSecurityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppSecurity"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppSecurityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the apps of a user. App information is returned only to the app owner.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- API providers can use the app IDs or their Alibaba Cloud accounts to query app information.
//
//		- Each provider can call this operation for a maximum of 200 times every day in a region.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithContext(ctx context.Context, request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppOwner) {
		query["AppOwner"] = request.AppOwner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authorized applications by API product.
//
// @param request - DescribeAppsByApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsByApiProductResponse
func (client *Client) DescribeAppsByApiProductWithContext(ctx context.Context, request *DescribeAppsByApiProductRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsByApiProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppsByApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsByApiProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorized APIs of a specified APP.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The specified application can call all APIs included in the responses.
//
// @param request - DescribeAuthorizedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthorizedApisResponse
func (client *Client) DescribeAuthorizedApisWithContext(ctx context.Context, request *DescribeAuthorizedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthorizedApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthorizedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthorizedApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current apps.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- All applications included in the responses have access to the specified API.
//
// @param request - DescribeAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthorizedAppsResponse
func (client *Client) DescribeAuthorizedAppsWithContext(ctx context.Context, request *DescribeAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthorizedAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppOwnerId) {
		query["AppOwnerId"] = request.AppOwnerId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthorizedApps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthorizedAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a backend service and its URL configured for each environment.
//
// @param request - DescribeBackendInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackendInfoResponse
func (client *Client) DescribeBackendInfoWithContext(ctx context.Context, request *DescribeBackendInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackendInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackendInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackendInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backend services. You can filter backend services by backend service name and backend service type.
//
// @param request - DescribeBackendListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackendListResponse
func (client *Client) DescribeBackendListWithContext(ctx context.Context, request *DescribeBackendListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackendListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackendList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackendListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a single dataset.
//
// @param request - DescribeDatasetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetInfoResponse
func (client *Client) DescribeDatasetInfoWithContext(ctx context.Context, request *DescribeDatasetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a data entry in a custom dataset.
//
// @param request - DescribeDatasetItemInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetItemInfoResponse
func (client *Client) DescribeDatasetItemInfoWithContext(ctx context.Context, request *DescribeDatasetItemInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetItemInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetItemInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetItemInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data entries of a custom dataset.
//
// @param request - DescribeDatasetItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetItemListResponse
func (client *Client) DescribeDatasetItemListWithContext(ctx context.Context, request *DescribeDatasetItemListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetItemListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemIds) {
		query["DatasetItemIds"] = request.DatasetItemIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetItemList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetItemListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom datasets.
//
// @param request - DescribeDatasetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetListResponse
func (client *Client) DescribeDatasetListWithContext(ctx context.Context, request *DescribeDatasetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetIds) {
		query["DatasetIds"] = request.DatasetIds
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of an asynchronous API publishing task.
//
// @param request - DescribeDeployApiTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployApiTaskResponse
func (client *Client) DescribeDeployApiTaskWithContext(ctx context.Context, request *DescribeDeployApiTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployApiTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployApiTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployApiTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the definition of an API that takes effect in an environment. The definition may differ from the definition being edited.
//
// @param request - DescribeDeployedApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployedApiResponse
func (client *Client) DescribeDeployedApiWithContext(ctx context.Context, request *DescribeDeployedApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployedApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployedApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployedApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that have been published to a specified environment.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeDeployedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployedApisResponse
func (client *Client) DescribeDeployedApisWithContext(ctx context.Context, request *DescribeDeployedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployedApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiMethod) {
		query["ApiMethod"] = request.ApiMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiPath) {
		query["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployedApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about a bound custom domain name, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// @param request - DescribeDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResponse
func (client *Client) DescribeDomainWithContext(ctx context.Context, request *DescribeDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average latency of an API group in an environment.
//
// @param request - DescribeGroupLatencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupLatencyResponse
func (client *Client) DescribeGroupLatencyWithContext(ctx context.Context, request *DescribeGroupLatencyRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupLatencyResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupLatency"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupLatencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of requests directed to an API group within a period of time.
//
// @param request - DescribeGroupQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupQpsResponse
func (client *Client) DescribeGroupQpsWithContext(ctx context.Context, request *DescribeGroupQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupQpsResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupQps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupQpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic of an API group.
//
// @param request - DescribeGroupTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupTrafficResponse
func (client *Client) DescribeGroupTrafficWithContext(ctx context.Context, request *DescribeGroupTrafficRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupTrafficResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupTraffic"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupTrafficResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical versions of an API.
//
// @param request - DescribeHistoryApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryApisResponse
func (client *Client) DescribeHistoryApisWithContext(ctx context.Context, request *DescribeHistoryApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an OAS API import task.
//
// @param request - DescribeImportOASTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportOASTaskResponse
func (client *Client) DescribeImportOASTaskWithContext(ctx context.Context, request *DescribeImportOASTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportOASTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImportOASTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportOASTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a dedicated instance cluster.
//
// @param request - DescribeInstanceClusterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceClusterInfoResponse
func (client *Client) DescribeInstanceClusterInfoWithContext(ctx context.Context, request *DescribeInstanceClusterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceClusterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceClusterInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceClusterInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries dedicated instance clusters.
//
// @param request - DescribeInstanceClusterListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceClusterListResponse
func (client *Client) DescribeInstanceClusterListWithContext(ctx context.Context, request *DescribeInstanceClusterListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceClusterListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClusterId) {
		query["InstanceClusterId"] = request.InstanceClusterId
	}

	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceClusterList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceClusterListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of lost connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceDropConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDropConnectionsResponse
func (client *Client) DescribeInstanceDropConnectionsWithContext(ctx context.Context, request *DescribeInstanceDropConnectionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDropConnectionsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDropConnections"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDropConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of dropped packets within a period of time.
//
// @param request - DescribeInstanceDropPacketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDropPacketResponse
func (client *Client) DescribeInstanceDropPacketWithContext(ctx context.Context, request *DescribeInstanceDropPacketRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDropPacketResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDropPacket"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDropPacketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceHttpCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceHttpCodeResponse
func (client *Client) DescribeInstanceHttpCodeWithContext(ctx context.Context, request *DescribeInstanceHttpCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceHttpCodeResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceHttpCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceHttpCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average latency of a dedicated instance over a period of time.
//
// @param request - DescribeInstanceLatencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceLatencyResponse
func (client *Client) DescribeInstanceLatencyWithContext(ctx context.Context, request *DescribeInstanceLatencyRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceLatencyResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceLatency"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceLatencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of new connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceNewConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceNewConnectionsResponse
func (client *Client) DescribeInstanceNewConnectionsWithContext(ctx context.Context, request *DescribeInstanceNewConnectionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceNewConnectionsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceNewConnections"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceNewConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of inbound and outbound packets of a dedicated instance within a period of time.
//
// @param request - DescribeInstancePacketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancePacketsResponse
func (client *Client) DescribeInstancePacketsWithContext(ctx context.Context, request *DescribeInstancePacketsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancePacketsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstancePackets"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancePacketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceQpsResponse
func (client *Client) DescribeInstanceQpsWithContext(ctx context.Context, request *DescribeInstanceQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceQpsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceQps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceQpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of concurrent connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceSlbConnectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSlbConnectResponse
func (client *Client) DescribeInstanceSlbConnectWithContext(ctx context.Context, request *DescribeInstanceSlbConnectRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSlbConnectResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSlbConnect"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSlbConnectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request traffic and response traffic of a dedicated instance within a period of time.
//
// @param request - DescribeInstanceTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTrafficResponse
func (client *Client) DescribeInstanceTrafficWithContext(ctx context.Context, request *DescribeInstanceTrafficRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTrafficResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceTraffic"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceTrafficResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of instances in a region. The instances include shared instances and dedicated instances.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableTagAuthorization) {
		query["EnableTagAuthorization"] = request.EnableTagAuthorization
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rule entries of an IP address-based traffic control policy.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can filter the query results by policy ID.
//
// @param request - DescribeIpControlPolicyItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpControlPolicyItemsResponse
func (client *Client) DescribeIpControlPolicyItemsWithContext(ctx context.Context, request *DescribeIpControlPolicyItemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpControlPolicyItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyItemId) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpControlPolicyItems"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpControlPolicyItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom access control lists (ACLs) on separate pages.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation is used to query the ACLs in a region. Region is a system parameter.
//
//		- You can filter the query results by ACL ID, name, or type.
//
//		- This operation cannot be used to query specific policies. If you want to query specific policies, call the [DescribeIpControlPolicyItems](~~DescribeIpControlPolicyItems~~) operation.
//
// @param request - DescribeIpControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpControlsResponse
func (client *Client) DescribeIpControlsWithContext(ctx context.Context, request *DescribeIpControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpControlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.IpControlType) {
		query["IpControlType"] = request.IpControlType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpControlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日志配置
//
// @param request - DescribeLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogConfigResponse
func (client *Client) DescribeLogConfigWithContext(ctx context.Context, request *DescribeLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of remaining ordered relationships for a purchaser.
//
// @param request - DescribeMarketRemainsQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMarketRemainsQuotaResponse
func (client *Client) DescribeMarketRemainsQuotaWithContext(ctx context.Context, request *DescribeMarketRemainsQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarketRemainsQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMarketRemainsQuota"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMarketRemainsQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the created models of an API group.
//
// Description:
//
//	Fuzzy queries are supported.
//
// @param request - DescribeModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelsResponse
func (client *Client) DescribeModelsWithContext(ctx context.Context, request *DescribeModelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelsResponse, _err error) {
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

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
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
		Action:      dara.String("DescribeModels"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified plug-in is bound.
//
// @param request - DescribePluginApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginApisResponse
func (client *Client) DescribePluginApisWithContext(ctx context.Context, request *DescribePluginApisRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of groups bound to a plugin based on the plugin ID
//
// @param request - DescribePluginGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginGroupsResponse
func (client *Client) DescribePluginGroupsWithContext(ctx context.Context, request *DescribePluginGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginGroupsResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询插件列表
//
// @param request - DescribePluginSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginSchemasResponse
func (client *Client) DescribePluginSchemasWithContext(ctx context.Context, request *DescribePluginSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginSchemasResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginSchemas"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询插件模板
//
// @param request - DescribePluginTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginTemplatesResponse
func (client *Client) DescribePluginTemplatesWithContext(ctx context.Context, request *DescribePluginTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginTemplates"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API Gateway plug-ins and the details of the plug-ins.
//
// Description:
//
//	  This operation supports pagination.
//
//		- This operation allows you to query plug-ins by business type.
//
//		- This operation allows you to query plug-ins by ID.
//
//		- This operation allows you to query plug-ins by name.
//
// @param request - DescribePluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsResponse
func (client *Client) DescribePluginsWithContext(ctx context.Context, request *DescribePluginsRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsResponse, _err error) {
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

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlugins"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the plug-ins that are bound to a running API in an environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation supports pagination.
//
// @param request - DescribePluginsByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsByApiResponse
func (client *Client) DescribePluginsByApiWithContext(ctx context.Context, request *DescribePluginsByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsByApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginsByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsByApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Plugins Bound to API Group
//
// @param request - DescribePluginsByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsByGroupResponse
func (client *Client) DescribePluginsByGroupWithContext(ctx context.Context, request *DescribePluginsByGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsByGroupResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginsByGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsByGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an API group purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApiGroupResponse
func (client *Client) DescribePurchasedApiGroupWithContext(ctx context.Context, request *DescribePurchasedApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApiGroupResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the API groups purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApiGroupsResponse
func (client *Client) DescribePurchasedApiGroupsWithContext(ctx context.Context, request *DescribePurchasedApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApiGroupsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApiGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApiGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs that are purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApisResponse
func (client *Client) DescribePurchasedApisWithContext(ctx context.Context, request *DescribePurchasedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are supported by API Gateway.
//
// Description:
//
// This operation queries regions in which API Gateway is available.
//
//   - This operation is intended for API providers and callers.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backend signature keys.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This operation is used to query the backend signature keys in a Region. Region is a system parameter.
//
// @param request - DescribeSignaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSignaturesResponse
func (client *Client) DescribeSignaturesWithContext(ctx context.Context, request *DescribeSignaturesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSignaturesResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSignatures"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSignaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to a specified API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeSignaturesByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSignaturesByApiResponse
func (client *Client) DescribeSignaturesByApiWithContext(ctx context.Context, request *DescribeSignaturesByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeSignaturesByApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSignaturesByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSignaturesByApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of API Gateway resources in a region.
//
// @param request - DescribeSummaryDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSummaryDataResponse
func (client *Client) DescribeSummaryDataWithContext(ctx context.Context, request *DescribeSummaryDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSummaryDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSummaryData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSummaryDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the common parameters supported by the system.
//
// Description:
//
//	  This API is intended for API callers.
//
//		- The response of this API contains the system parameters that are optional in API definitions.
//
// @param request - DescribeSystemParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemParametersResponse
func (client *Client) DescribeSystemParametersWithContext(ctx context.Context, request *DescribeSystemParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemParameters"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom throttling policies and their details. Conditional queries are supported.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API can be used to query all existing throttling policies (including special throttling policies) and their details.
//
//		- You can specify query conditions. For example, you can query the throttling policies bound to a specified API or in a specified environment.
//
// @param request - DescribeTrafficControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficControlsResponse
func (client *Client) DescribeTrafficControlsWithContext(ctx context.Context, request *DescribeTrafficControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficControlsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrafficControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficControlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling policy that is bound to a specific API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeTrafficControlsByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficControlsByApiResponse
func (client *Client) DescribeTrafficControlsByApiWithContext(ctx context.Context, request *DescribeTrafficControlsByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficControlsByApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrafficControlsByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficControlsByApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量更新API后端元定结果
//
// @param request - DescribeUpdateBackendTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpdateBackendTaskResponse
func (client *Client) DescribeUpdateBackendTaskWithContext(ctx context.Context, request *DescribeUpdateBackendTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpdateBackendTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpdateBackendTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpdateBackendTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询更新VPC授权的任务
//
// @param request - DescribeUpdateVpcInfoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpdateVpcInfoTaskResponse
func (client *Client) DescribeUpdateVpcInfoTaskWithContext(ctx context.Context, request *DescribeUpdateVpcInfoTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpdateVpcInfoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpdateVpcInfoTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpdateVpcInfoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries VPC access authorizations.
//
// @param request - DescribeVpcAccessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcAccessesResponse
func (client *Client) DescribeVpcAccessesWithContext(ctx context.Context, request *DescribeVpcAccessesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcAccessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccurateQuery) {
		query["AccurateQuery"] = request.AccurateQuery
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcAccessId) {
		query["VpcAccessId"] = request.VpcAccessId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcAccesses"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcAccessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries zones in a region.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches APIs from an API product.
//
// @param request - DetachApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachApiProductResponse
func (client *Client) DetachApiProductWithContext(ctx context.Context, request *DetachApiProductRequest, runtime *dara.RuntimeOptions) (_result *DetachApiProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.Apis) {
		query["Apis"] = request.Apis
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachApiProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unbind group plugin
//
// @param request - DetachGroupPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGroupPluginResponse
func (client *Client) DetachGroupPluginWithContext(ctx context.Context, request *DetachGroupPluginRequest, runtime *dara.RuntimeOptions) (_result *DetachGroupPluginResponse, _err error) {
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

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachGroupPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachGroupPluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑插件
//
// @param request - DetachPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPluginResponse
func (client *Client) DetachPluginWithContext(ctx context.Context, request *DetachPluginRequest, runtime *dara.RuntimeOptions) (_result *DetachPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Disables access control on an instance.
//
// @param request - DisableInstanceAccessControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstanceAccessControlResponse
func (client *Client) DisableInstanceAccessControlWithContext(ctx context.Context, request *DisableInstanceAccessControlRequest, runtime *dara.RuntimeOptions) (_result *DisableInstanceAccessControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInstanceAccessControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstanceAccessControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an internal domain name resolution from a dedicated instance.
//
// @param tmpReq - DissociateInstanceWithPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateInstanceWithPrivateDNSResponse
func (client *Client) DissociateInstanceWithPrivateDNSWithContext(ctx context.Context, tmpReq *DissociateInstanceWithPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *DissociateInstanceWithPrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DissociateInstanceWithPrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntranetDomains) {
		request.IntranetDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntranetDomains, dara.String("IntranetDomains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomainsShrink) {
		body["IntranetDomains"] = request.IntranetDomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateInstanceWithPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the syntax before Swagger-compliant data is imported.
//
// @param tmpReq - DryRunSwaggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DryRunSwaggerResponse
func (client *Client) DryRunSwaggerWithContext(ctx context.Context, tmpReq *DryRunSwaggerRequest, runtime *dara.RuntimeOptions) (_result *DryRunSwaggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DryRunSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalCondition) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, dara.String("GlobalCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.GlobalConditionShrink) {
		query["GlobalCondition"] = request.GlobalConditionShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DryRunSwagger"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DryRunSwaggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Specifies the access control policy of an instance.
//
// @param request - EnableInstanceAccessControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstanceAccessControlResponse
func (client *Client) EnableInstanceAccessControlWithContext(ctx context.Context, request *EnableInstanceAccessControlRequest, runtime *dara.RuntimeOptions) (_result *EnableInstanceAccessControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInstanceAccessControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstanceAccessControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出OAS
//
// @param tmpReq - ExportOASRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportOASResponse
func (client *Client) ExportOASWithContext(ctx context.Context, tmpReq *ExportOASRequest, runtime *dara.RuntimeOptions) (_result *ExportOASResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportOASShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiIdList) {
		request.ApiIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiIdList, dara.String("ApiIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIdListShrink) {
		query["ApiIdList"] = request.ApiIdListShrink
	}

	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OasVersion) {
		query["OasVersion"] = request.OasVersion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.WithXExtensions) {
		query["WithXExtensions"] = request.WithXExtensions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportOAS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportOASResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports APIs based on the OAS standard.
//
// @param request - ImportOASRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportOASResponse
func (client *Client) ImportOASWithContext(ctx context.Context, request *ImportOASRequest, runtime *dara.RuntimeOptions) (_result *ImportOASResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IgnoreWarning) {
		query["IgnoreWarning"] = request.IgnoreWarning
	}

	if !dara.IsNil(request.OASVersion) {
		query["OASVersion"] = request.OASVersion
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.RequestMode) {
		query["RequestMode"] = request.RequestMode
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SkipDryRun) {
		query["SkipDryRun"] = request.SkipDryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportOAS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportOASResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API by importing Swagger-compliant data.
//
// Description:
//
//	  Alibaba Cloud supports extensions based on Swagger 2.0.
//
//		- Alibaba Cloud supports Swagger configuration files in JSON and YAML formats.
//
// @param tmpReq - ImportSwaggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportSwaggerResponse
func (client *Client) ImportSwaggerWithContext(ctx context.Context, tmpReq *ImportSwaggerRequest, runtime *dara.RuntimeOptions) (_result *ImportSwaggerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalCondition) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, dara.String("GlobalCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GlobalConditionShrink) {
		query["GlobalCondition"] = request.GlobalConditionShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportSwagger"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportSwaggerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries internal domain name resolutions by domain name or resolution type.
//
// @param request - ListPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateDNSResponse
func (client *Client) ListPrivateDNSWithContext(ctx context.Context, request *ListPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the visible resource tags.
//
// Description:
//
//	  The Tag.N.Key and Tag.N.Value parameters constitute a key-value pair.
//
//		- ResourceId.N must meet all the key-value pairs that are entered. If you enter multiple key-value pairs, resources that contain the specified key-value pairs are returned.
//
//		- This operation is used to query resource tags based on conditions. If no relationship matches the conditions, an empty list is returned.
//
//		- You can query both user tags and visible system tags.
//
//		- In addition to the required parameters, you can also specify ResourceId.N to query the visible resource tags of a specified resource in a region.
//
//		- You can also specify Tag.N.Key to query the visible keys of a specified key in a region.
//
//		- At least one of ResourceId.N, Tag.N.Key, and Tag.N.Value exists.
//
//		- You can query tags of the same type or different types in a single operation.
//
//		- You can query all your user types and visible system tags.
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
		Version:     dara.String("2016-07-14"),
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
// Modifies the definition of an API.
//
// Description:
//
// *This operation is intended for API providers.**
//
//   - This API operation requires a full update. Updates of partial parameters are not supported.
//
//   - When you modify an API name, make sure that the name of each API within the same group is unique.
//
//   - When you modify the request path, make sure that each request path within the same group is unique.
//
//   - The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiResponse
func (client *Client) ModifyApiWithContext(ctx context.Context, request *ModifyApiRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendEnable) {
		query["BackendEnable"] = request.BackendEnable
	}

	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpenIdConnectConfig) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !dara.IsNil(request.RequestConfig) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !dara.IsNil(request.ResultBodyModel) {
		query["ResultBodyModel"] = request.ResultBodyModel
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceConfig) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WebSocketApiType) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConstantParameters) {
		body["ConstantParameters"] = request.ConstantParameters
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		body["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		body["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.RequestParameters) {
		body["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.ResultDescriptions) {
		body["ResultDescriptions"] = request.ResultDescriptions
	}

	if !dara.IsNil(request.ResultSample) {
		body["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		body["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.SystemParameters) {
		body["SystemParameters"] = request.SystemParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the draft definition of an API. This operation is different from the ModifyApi operation. This operation does not require all information about the API. You need to only specify the parameters that you want to modify. For example, if you want to change the authentication method of the API from Anonymous to APP, you specify APP as the value of AuthType and do not need to configure other parameters.
//
// @param request - ModifyApiConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiConfigurationResponse
func (client *Client) ModifyApiConfigurationWithContext(ctx context.Context, request *ModifyApiConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BodyFormat) {
		query["BodyFormat"] = request.BodyFormat
	}

	if !dara.IsNil(request.BodyModel) {
		query["BodyModel"] = request.BodyModel
	}

	if !dara.IsNil(request.ContentTypeCategory) {
		query["ContentTypeCategory"] = request.ContentTypeCategory
	}

	if !dara.IsNil(request.ContentTypeValue) {
		query["ContentTypeValue"] = request.ContentTypeValue
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		query["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		query["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.FunctionComputeConfig) {
		query["FunctionComputeConfig"] = request.FunctionComputeConfig
	}

	if !dara.IsNil(request.HttpConfig) {
		query["HttpConfig"] = request.HttpConfig
	}

	if !dara.IsNil(request.MockConfig) {
		query["MockConfig"] = request.MockConfig
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.OssConfig) {
		query["OssConfig"] = request.OssConfig
	}

	if !dara.IsNil(request.PostBodyDescription) {
		query["PostBodyDescription"] = request.PostBodyDescription
	}

	if !dara.IsNil(request.RequestHttpMethod) {
		query["RequestHttpMethod"] = request.RequestHttpMethod
	}

	if !dara.IsNil(request.RequestMode) {
		query["RequestMode"] = request.RequestMode
	}

	if !dara.IsNil(request.RequestParameters) {
		query["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.RequestPath) {
		query["RequestPath"] = request.RequestPath
	}

	if !dara.IsNil(request.RequestProtocol) {
		query["RequestProtocol"] = request.RequestProtocol
	}

	if !dara.IsNil(request.ResultSample) {
		query["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		query["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.ServiceProtocol) {
		query["ServiceProtocol"] = request.ServiceProtocol
	}

	if !dara.IsNil(request.ServiceTimeout) {
		query["ServiceTimeout"] = request.ServiceTimeout
	}

	if !dara.IsNil(request.UseBackendService) {
		query["UseBackendService"] = request.UseBackendService
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VpcConfig) {
		query["VpcConfig"] = request.VpcConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiConfiguration"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or basepath of an existing API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupResponse
func (client *Client) ModifyApiGroupWithContext(ctx context.Context, request *ModifyApiGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
	}

	if !dara.IsNil(request.CompatibleFlags) {
		query["CompatibleFlags"] = request.CompatibleFlags
	}

	if !dara.IsNil(request.CustomAppCodeConfig) {
		query["CustomAppCodeConfig"] = request.CustomAppCodeConfig
	}

	if !dara.IsNil(request.CustomTraceConfig) {
		query["CustomTraceConfig"] = request.CustomTraceConfig
	}

	if !dara.IsNil(request.CustomerConfigs) {
		query["CustomerConfigs"] = request.CustomerConfigs
	}

	if !dara.IsNil(request.DefaultDomain) {
		query["DefaultDomain"] = request.DefaultDomain
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FilterAppCodeForBackend) {
		query["FilterAppCodeForBackend"] = request.FilterAppCodeForBackend
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PassthroughHeaders) {
		query["PassthroughHeaders"] = request.PassthroughHeaders
	}

	if !dara.IsNil(request.RpcPattern) {
		query["RpcPattern"] = request.RpcPattern
	}

	if !dara.IsNil(request.RpsLimitForServerless) {
		query["RpsLimitForServerless"] = request.RpsLimitForServerless
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SupportSSE) {
		query["SupportSSE"] = request.SupportSSE
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserLogConfig) {
		query["UserLogConfig"] = request.UserLogConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更分组实例
//
// @param request - ModifyApiGroupInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupInstanceResponse
func (client *Client) ModifyApiGroupInstanceWithContext(ctx context.Context, request *ModifyApiGroupInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupInstance"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network policy of an API group.
//
// @param request - ModifyApiGroupNetworkPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupNetworkPolicyResponse
func (client *Client) ModifyApiGroupNetworkPolicyWithContext(ctx context.Context, request *ModifyApiGroupNetworkPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupNetworkPolicyResponse, _err error) {
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

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.InnerDomainEnable) {
		query["InnerDomainEnable"] = request.InnerDomainEnable
	}

	if !dara.IsNil(request.InternetEnable) {
		query["InternetEnable"] = request.InternetEnable
	}

	if !dara.IsNil(request.InternetIPV6Enable) {
		query["InternetIPV6Enable"] = request.InternetIPV6Enable
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIntranetEnable) {
		query["VpcIntranetEnable"] = request.VpcIntranetEnable
	}

	if !dara.IsNil(request.VpcSlbIntranetEnable) {
		query["VpcSlbIntranetEnable"] = request.VpcSlbIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupNetworkPolicy"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupNetworkPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the VPC whitelist of an API group.
//
// @param request - ModifyApiGroupVpcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupVpcWhitelistResponse
func (client *Client) ModifyApiGroupVpcWhitelistWithContext(ctx context.Context, request *ModifyApiGroupVpcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupVpcWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIds) {
		query["VpcIds"] = request.VpcIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupVpcWhitelist"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a specified app.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppName or Description can be modified. If these parameters are not specified, no modifications are made and the operation will directly return a successful response.********
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改后端服务
//
// @param request - ModifyBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackendResponse
func (client *Client) ModifyBackendWithContext(ctx context.Context, request *ModifyBackendRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改后端服务在环境上的定义
//
// @param request - ModifyBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackendModelResponse
func (client *Client) ModifyBackendModelWithContext(ctx context.Context, request *ModifyBackendModelRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackendModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelData) {
		query["BackendModelData"] = request.BackendModelData
	}

	if !dara.IsNil(request.BackendModelId) {
		query["BackendModelId"] = request.BackendModelId
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackendModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a custom dataset.
//
// @param request - ModifyDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatasetResponse
func (client *Client) ModifyDatasetWithContext(ctx context.Context, request *ModifyDatasetRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the expiration time and description of a data entry in a custom dataset.
//
// @param request - ModifyDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatasetItemResponse
func (client *Client) ModifyDatasetItemWithContext(ctx context.Context, request *ModifyDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatasetItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatasetItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of an API Gateway instance.
//
// @param tmpReq - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithContext(ctx context.Context, tmpReq *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ToConnectVpcIpBlock) {
		request.ToConnectVpcIpBlockShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToConnectVpcIpBlock, dara.String("ToConnectVpcIpBlock"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteVpcIpBlock) {
		query["DeleteVpcIpBlock"] = request.DeleteVpcIpBlock
	}

	if !dara.IsNil(request.EgressIpv6Enable) {
		query["EgressIpv6Enable"] = request.EgressIpv6Enable
	}

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.IPV6Enabled) {
		query["IPV6Enabled"] = request.IPV6Enabled
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntranetSegments) {
		query["IntranetSegments"] = request.IntranetSegments
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	if !dara.IsNil(request.ToConnectVpcIpBlockShrink) {
		query["ToConnectVpcIpBlock"] = request.ToConnectVpcIpBlockShrink
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcSlbIntranetEnable) {
		query["VpcSlbIntranetEnable"] = request.VpcSlbIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the configurations of an API Gateway instance.
//
// @param request - ModifyInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpecWithContext(ctx context.Context, request *ModifyInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.ModifyAction) {
		query["ModifyAction"] = request.ModifyAction
	}

	if !dara.IsNil(request.SkipWaitSwitch) {
		query["SkipWaitSwitch"] = request.SkipWaitSwitch
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceSpec"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify instance client VPC config.
//
// @param request - ModifyInstanceVpcAttributeForConsoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVpcAttributeForConsoleResponse
func (client *Client) ModifyInstanceVpcAttributeForConsoleWithContext(ctx context.Context, request *ModifyInstanceVpcAttributeForConsoleRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVpcAttributeForConsoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteVpcAccess) {
		query["DeleteVpcAccess"] = request.DeleteVpcAccess
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		query["VpcOwnerId"] = request.VpcOwnerId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceVpcAttributeForConsole"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceVpcAttributeForConsoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the VPC domain name policy of an API group.
//
// @param request - ModifyIntranetDomainPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIntranetDomainPolicyResponse
func (client *Client) ModifyIntranetDomainPolicyWithContext(ctx context.Context, request *ModifyIntranetDomainPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyIntranetDomainPolicyResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIntranetEnable) {
		query["VpcIntranetEnable"] = request.VpcIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIntranetDomainPolicy"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIntranetDomainPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation allows you to modify only the name and description of an ACL. You cannot modify the type of the ACL.
//
// @param request - ModifyIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpControlResponse
func (client *Client) ModifyIpControlWithContext(ctx context.Context, request *ModifyIpControlRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpControlResponse, _err error) {
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

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a policy in an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The modification immediately takes effect on all the APIs that are bound to the policy.
//
//		- This operation causes a full modification of the content of a policy.
//
// @param request - ModifyIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpControlPolicyItemResponse
func (client *Client) ModifyIpControlPolicyItemWithContext(ctx context.Context, request *ModifyIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpControlPolicyItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CidrIp) {
		query["CidrIp"] = request.CidrIp
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PolicyItemId) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpControlPolicyItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改日志配置
//
// @param request - ModifyLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLogConfigResponse
func (client *Client) ModifyLogConfigWithContext(ctx context.Context, request *ModifyLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlsLogStore) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the model of an API group.
//
// @param request - ModifyModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyModelResponse
func (client *Client) ModifyModelWithContext(ctx context.Context, request *ModifyModelRequest, runtime *dara.RuntimeOptions) (_result *ModifyModelResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.NewModelName) {
		query["NewModelName"] = request.NewModelName
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of the plug-in must be unique.
//
// @param request - ModifyPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPluginResponse
func (client *Client) ModifyPluginWithContext(ctx context.Context, request *ModifyPluginRequest, runtime *dara.RuntimeOptions) (_result *ModifyPluginResponse, _err error) {
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

	if !dara.IsNil(request.PluginData) {
		query["PluginData"] = request.PluginData
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation modifies the name, Key value, and Secret value of an existing signature key.
//
//		- Note that the modification takes effect immediately. If the key has been bound to an API, you must adjust the backend signature verification based on the new key accordingly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifySignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySignatureResponse
func (client *Client) ModifySignatureWithContext(ctx context.Context, request *ModifySignatureRequest, runtime *dara.RuntimeOptions) (_result *ModifySignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.SignatureKey) {
		query["SignatureKey"] = request.SignatureKey
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	if !dara.IsNil(request.SignatureSecret) {
		query["SignatureSecret"] = request.SignatureSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The modifications take effect on the bound APIs instantly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrafficControlResponse
func (client *Client) ModifyTrafficControlWithContext(ctx context.Context, request *ModifyTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrafficControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDefault) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !dara.IsNil(request.AppDefault) {
		query["AppDefault"] = request.AppDefault
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !dara.IsNil(request.TrafficControlUnit) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !dara.IsNil(request.UserDefault) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTrafficControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a virtual private cloud (VPC) authorization and updates the metadata of the API associated with the VPC authorization.
//
// @param request - ModifyVpcAccessAndUpdateApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcAccessAndUpdateApisResponse
func (client *Client) ModifyVpcAccessAndUpdateApisWithContext(ctx context.Context, request *ModifyVpcAccessAndUpdateApisRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcAccessAndUpdateApisResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Refresh) {
		query["Refresh"] = request.Refresh
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcTargetHostName) {
		query["VpcTargetHostName"] = request.VpcTargetHostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcAccessAndUpdateApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcAccessAndUpdateApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request logs of a user.
//
// @param request - QueryRequestLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRequestLogsResponse
func (client *Client) QueryRequestLogsWithContext(ctx context.Context, request *QueryRequestLogsRequest, runtime *dara.RuntimeOptions) (_result *QueryRequestLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestLogId) {
		query["RequestLogId"] = request.RequestLogId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRequestLogs"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRequestLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reactivates a custom domain name whose validity status is Abnormal.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must solve the problem that is mentioned in the domain name exception prompt before you can reactivate the domain name.
//
//		- A typical reason why a custom domain name becomes abnormal is that the domain name does not have an ICP filing or the domain name is included in a blacklist by the administration. When a custom domain name is abnormal, users cannot use it to call APIs.
//
//		- You can call this operation to reactivate the domain name to resume normal access.
//
// @param request - ReactivateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReactivateDomainResponse
func (client *Client) ReactivateDomainWithContext(ctx context.Context, request *ReactivateDomainRequest, runtime *dara.RuntimeOptions) (_result *ReactivateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReactivateDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReactivateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除访问控制策略中IP条目
//
// @param request - RemoveAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAccessControlListEntryResponse
func (client *Client) RemoveAccessControlListEntryWithContext(ctx context.Context, request *RemoveAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *RemoveAccessControlListEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAccessControlListEntry"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAccessControlListEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on API products from an application.
//
// @param tmpReq - RemoveApiProductsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApiProductsAuthoritiesResponse
func (client *Client) RemoveApiProductsAuthoritiesWithContext(ctx context.Context, tmpReq *RemoveApiProductsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveApiProductsAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveApiProductsAuthoritiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiProductIds) {
		request.ApiProductIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiProductIds, dara.String("ApiProductIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductIdsShrink) {
		query["ApiProductIds"] = request.ApiProductIdsShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApiProductsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApiProductsAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on multiple APIs from a specified application.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveApisAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApisAuthoritiesResponse
func (client *Client) RemoveApisAuthoritiesWithContext(ctx context.Context, request *RemoveApisAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveApisAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApisAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApisAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on a specified API from multiple applications. In this case, multiple applications map to a single API.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveAppsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppsAuthoritiesResponse
func (client *Client) RemoveAppsAuthoritiesWithContext(ctx context.Context, request *RemoveAppsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveAppsAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAppsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAppsAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an API from an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The unbinding takes effect immediately. After the API is unbound from the ACL, the corresponding environment does not have any IP address access control in place for the API.
//
// @param request - RemoveIpControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpControlApisResponse
func (client *Client) RemoveIpControlApisWithContext(ctx context.Context, request *RemoveIpControlApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpControlApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveIpControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpControlApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more policies from an access control list (ACL).
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - RemoveIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpControlPolicyItemResponse
func (client *Client) RemoveIpControlPolicyItemWithContext(ctx context.Context, request *RemoveIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpControlPolicyItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PolicyItemIds) {
		query["PolicyItemIds"] = request.PolicyItemIds
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpControlPolicyItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a backend signature key from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The operation takes effect immediately. The request sent from API Gateway to the backend service does not contain the signature string. The corresponding verification step can be removed from the backend.
//
// @param request - RemoveSignatureApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSignatureApisResponse
func (client *Client) RemoveSignatureApisWithContext(ctx context.Context, request *RemoveSignatureApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveSignatureApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSignatureApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSignatureApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a specified throttling policy from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to unbind a specified throttling policy from up to 100 APIs at a time.
//
// @param request - RemoveTrafficControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTrafficControlApisResponse
func (client *Client) RemoveTrafficControlApisWithContext(ctx context.Context, request *RemoveTrafficControlApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveTrafficControlApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTrafficControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTrafficControlApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC authorization without unpublishing the associated APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Revokes the permissions of API Gateway to access your VPC instance.
//
// >  Deleting an authorization affects the associated API. Before you delete the authorization, make sure that it is not used by the API.
//
// @param request - RemoveVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVpcAccessResponse
func (client *Client) RemoveVpcAccessWithContext(ctx context.Context, request *RemoveVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *RemoveVpcAccessResponse, _err error) {
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

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVpcAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除VPC授权并下线关联API
//
// @param request - RemoveVpcAccessAndAbolishApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVpcAccessAndAbolishApisResponse
func (client *Client) RemoveVpcAccessAndAbolishApisWithContext(ctx context.Context, request *RemoveVpcAccessAndAbolishApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveVpcAccessAndAbolishApisResponse, _err error) {
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

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVpcAccessAndAbolishApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVpcAccessAndAbolishApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the AppCode of an application. You can call this operation only once per minute.
//
// @param request - ResetAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAppCodeResponse
func (client *Client) ResetAppCodeWithContext(ctx context.Context, request *ResetAppCodeRequest, runtime *dara.RuntimeOptions) (_result *ResetAppCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.NewAppCode) {
		query["NewAppCode"] = request.NewAppCode
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAppCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the key of an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A new secret is automatically generated after you have called this operation. This secret cannot be customized.
//
//		- The results returned by this operation do not contain the application secret. You can obtain the secret by calling DescribeAppSecurity.
//
// @param request - ResetAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAppSecretResponse
func (client *Client) ResetAppSecretWithContext(ctx context.Context, request *ResetAppSecretRequest, runtime *dara.RuntimeOptions) (_result *ResetAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.NewAppKey) {
		query["NewAppKey"] = request.NewAppKey
	}

	if !dara.IsNil(request.NewAppSecret) {
		query["NewAppSecret"] = request.NewAppSecret
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAppSecret"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAppSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据APP生成SDK
//
// @param request - SdkGenerateByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByAppResponse
func (client *Client) SdkGenerateByAppWithContext(ctx context.Context, request *SdkGenerateByAppRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成与App关联的API的SDK
//
// @param request - SdkGenerateByAppForRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByAppForRegionResponse
func (client *Client) SdkGenerateByAppForRegionWithContext(ctx context.Context, request *SdkGenerateByAppForRegionRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByAppForRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByAppForRegion"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByAppForRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据分组生成SDK
//
// @param request - SdkGenerateByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByGroupResponse
func (client *Client) SdkGenerateByGroupWithContext(ctx context.Context, request *SdkGenerateByGroupRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByGroupResponse, _err error) {
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

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改访问控制策略的名称
//
// @param request - SetAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccessControlListAttributeResponse
func (client *Client) SetAccessControlListAttributeWithContext(ctx context.Context, request *SetAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetAccessControlListAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAccessControlListAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAccessControlListAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on API products to an application.
//
// @param tmpReq - SetApiProductsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApiProductsAuthoritiesResponse
func (client *Client) SetApiProductsAuthoritiesWithContext(ctx context.Context, tmpReq *SetApiProductsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetApiProductsAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetApiProductsAuthoritiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiProductIds) {
		request.ApiProductIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiProductIds, dara.String("ApiProductIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductIdsShrink) {
		query["ApiProductIds"] = request.ApiProductIdsShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApiProductsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApiProductsAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a specified application to call multiple APIs.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetApisAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApisAuthoritiesResponse
func (client *Client) SetApisAuthoritiesWithContext(ctx context.Context, request *SetApisAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetApisAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApisAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApisAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes multiple applications to call APIs in an API product.
//
// @param request - SetAppsAuthToApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppsAuthToApiProductResponse
func (client *Client) SetAppsAuthToApiProductWithContext(ctx context.Context, request *SetAppsAuthToApiProductRequest, runtime *dara.RuntimeOptions) (_result *SetAppsAuthToApiProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppsAuthToApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppsAuthToApiProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants access permissions on a specified API to multiple applications.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetAppsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppsAuthoritiesResponse
func (client *Client) SetAppsAuthoritiesWithContext(ctx context.Context, request *SetAppsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetAppsAuthoritiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppsAuthoritiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a custom domain name to a specified API group.
//
// @param request - SetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainResponse
func (client *Client) SetDomainWithContext(ctx context.Context, request *SetDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindStageName) {
		query["BindStageName"] = request.BindStageName
	}

	if !dara.IsNil(request.CustomDomainType) {
		query["CustomDomainType"] = request.CustomDomainType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsForce) {
		query["IsForce"] = request.IsForce
	}

	if !dara.IsNil(request.IsHttpRedirectToHttps) {
		query["IsHttpRedirectToHttps"] = request.IsHttpRedirectToHttps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an SSL certificate for a specified custom domain name.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The SSL certificate must match the custom domain name.
//
//		- After the SSL certificate is bound, HTTPS-based API services become available.
//
// @param request - SetDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainCertificateResponse
func (client *Client) SetDomainCertificateWithContext(ctx context.Context, request *SetDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaCertificateBody) {
		query["CaCertificateBody"] = request.CaCertificateBody
	}

	if !dara.IsNil(request.CertificateBody) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificatePrivateKey) {
		query["CertificatePrivateKey"] = request.CertificatePrivateKey
	}

	if !dara.IsNil(request.ClientCertSDnPassThrough) {
		query["ClientCertSDnPassThrough"] = request.ClientCertSDnPassThrough
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SslOcspCacheEnable) {
		query["SslOcspCacheEnable"] = request.SslOcspCacheEnable
	}

	if !dara.IsNil(request.SslOcspEnable) {
		query["SslOcspEnable"] = request.SslOcspEnable
	}

	if !dara.IsNil(request.SslVerifyDepth) {
		query["SslVerifyDepth"] = request.SslVerifyDepth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainCertificate"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables WebSocket for a custom domain name.
//
// @param request - SetDomainWebSocketStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainWebSocketStatusResponse
func (client *Client) SetDomainWebSocketStatusWithContext(ctx context.Context, request *SetDomainWebSocketStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainWebSocketStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionValue) {
		query["ActionValue"] = request.ActionValue
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WSSEnable) {
		query["WSSEnable"] = request.WSSEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainWebSocketStatus"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainWebSocketStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置分组授权AppCode
//
// @param request - SetGroupAuthAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGroupAuthAppCodeResponse
func (client *Client) SetGroupAuthAppCodeWithContext(ctx context.Context, request *SetGroupAuthAppCodeRequest, runtime *dara.RuntimeOptions) (_result *SetGroupAuthAppCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthAppCode) {
		query["AuthAppCode"] = request.AuthAppCode
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGroupAuthAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGroupAuthAppCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a binding relationship between specified access control lists (ACLs) and APIs.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A maximum of 100 APIs can be bound at a time.
//
// @param request - SetIpControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIpControlApisResponse
func (client *Client) SetIpControlApisWithContext(ctx context.Context, request *SetIpControlApisRequest, runtime *dara.RuntimeOptions) (_result *SetIpControlApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIpControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIpControlApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a signature key to APIs.
//
// @param request - SetSignatureApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSignatureApisResponse
func (client *Client) SetSignatureApisWithContext(ctx context.Context, request *SetSignatureApisRequest, runtime *dara.RuntimeOptions) (_result *SetSignatureApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSignatureApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSignatureApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a throttling policy to APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to bind a specific throttling policy to up to 100 APIs at a time.
//
// @param request - SetTrafficControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTrafficControlApisResponse
func (client *Client) SetTrafficControlApisWithContext(ctx context.Context, request *SetTrafficControlApisRequest, runtime *dara.RuntimeOptions) (_result *SetTrafficControlApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTrafficControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTrafficControlApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) access authorization and enables reverse access.
//
// Description:
//
// This operation is intended for API providers.
//
//   - This operation is used to authorize API Gateway to access your VPC instance.
//
// @param request - SetVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVpcAccessResponse
func (client *Client) SetVpcAccessWithContext(ctx context.Context, request *SetVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *SetVpcAccessResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcTargetHostName) {
		query["VpcTargetHostName"] = request.VpcTargetHostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVpcAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a wildcard domain name template for a bound custom domain name.
//
// @param request - SetWildcardDomainPatternsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWildcardDomainPatternsResponse
func (client *Client) SetWildcardDomainPatternsWithContext(ctx context.Context, request *SetWildcardDomainPatternsRequest, runtime *dara.RuntimeOptions) (_result *SetWildcardDomainPatternsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WildcardDomainPatterns) {
		query["WildcardDomainPatterns"] = request.WildcardDomainPatterns
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWildcardDomainPatterns"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWildcardDomainPatternsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the definition of an API in a specified runtime environment to a historical version.
//
// @param request - SwitchApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchApiResponse
func (client *Client) SwitchApiWithContext(ctx context.Context, request *SwitchApiRequest, runtime *dara.RuntimeOptions) (_result *SwitchApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryVersion) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag-resource relationship.
//
// Description:
//
//	  All tags (key-value pairs) are applied to all resources of a specified ResourceId, with each resource specified as ResourceId.N.
//
//		- Tag.N is a resource tag consisting of a key-value pair: Tag.N.Key and Tag.N.Value.
//
//		- If you call this operation to tag multiple resources simultaneously, either all or none of the resources will be tagged.
//
//		- If you specify Tag.1.Value in addition to required parameters, you must also specify Tag.1.Key. Otherwise, an InvalidParameter.TagKey error is reported. A tag that has a value must have the corresponding key, but the key can be an empty string.
//
//		- If a tag with the same key has been bound to a resource, the new tag will overwrite the existing one.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2016-07-14"),
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
// Removes tags from resources.
//
// Description:
//
//	  If you call this operation to untag multiple resources simultaneously, either all or none of the resources will be untagged.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to true, all tags bound to the specified resources will be deleted. If a resource does not have any tags, the request is not processed but a success is returned.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to false, the request is not processed but a success is returned.
//
//		- When tag keys are specified, the All parameter is invalid.
//
//		- When multiple resources and key-value pairs are specified, the specified tags bound to the resources are deleted.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2016-07-14"),
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
// Modifies an internal domain name resolution.
//
// @param tmpReq - UpdatePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateDNSResponse
func (client *Client) UpdatePrivateDNSWithContext(ctx context.Context, tmpReq *UpdatePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Records) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, dara.String("Records"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RecordsShrink) {
		body["Records"] = request.RecordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the network connectivity between an API Gateway instance and a port on an Elastic Compute Service (ECS) or Server Load Balance (SLB) instance in a virtual private cloud (VPC) access authorization.
//
// @param request - ValidateVpcConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateVpcConnectivityResponse
func (client *Client) ValidateVpcConnectivityWithContext(ctx context.Context, request *ValidateVpcConnectivityRequest, runtime *dara.RuntimeOptions) (_result *ValidateVpcConnectivityResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcAccessId) {
		query["VpcAccessId"] = request.VpcAccessId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateVpcConnectivity"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateVpcConnectivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
