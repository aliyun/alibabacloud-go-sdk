// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 挂载connector的应用
//
// @param tmpReq - AttachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApplication2ConnectorResponse
func (client *Client) AttachApplication2ConnectorWithContext(ctx context.Context, tmpReq *AttachApplication2ConnectorRequest, runtime *dara.RuntimeOptions) (_result *AttachApplication2ConnectorResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AttachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationIds) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, dara.String("ApplicationIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIdsShrink) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachApplication2Connector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachApplication2ConnectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 挂载业务策略至指定审批流程
//
// @param request - AttachPolicy2ApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicy2ApprovalProcessResponse
func (client *Client) AttachPolicy2ApprovalProcessWithContext(ctx context.Context, request *AttachPolicy2ApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicy2ApprovalProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPolicy2ApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPolicy2ApprovalProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建审批流程
//
// @param tmpReq - CreateApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApprovalProcessResponse
func (client *Client) CreateApprovalProcessWithContext(ctx context.Context, tmpReq *CreateApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateApprovalProcessResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateApprovalProcessShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchSchemas) {
		request.MatchSchemasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchSchemas, dara.String("MatchSchemas"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchSchemasShrink) {
		body["MatchSchemas"] = request.MatchSchemasShrink
	}

	if !dara.IsNil(request.ProcessName) {
		body["ProcessName"] = request.ProcessName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessNodes) {
		bodyFlat["ProcessNodes"] = request.ProcessNodes
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApprovalProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义身份源用户
//
// @param request - CreateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientUserResponse
func (client *Client) CreateClientUserWithContext(ctx context.Context, request *CreateClientUserRequest, runtime *dara.RuntimeOptions) (_result *CreateClientUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	if !dara.IsNil(request.MobileNumber) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建动态路由
//
// @param request - CreateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDynamicRouteResponse
func (client *Client) CreateDynamicRouteWithContext(ctx context.Context, request *CreateDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicRouteType) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextHop) {
		body["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionIds) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建加速策略
//
// @param request - CreateEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseAcceleratePolicyResponse
func (client *Client) CreateEnterpriseAcceleratePolicyWithContext(ctx context.Context, request *CreateEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseAcceleratePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationType) {
		body["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ShowInClient) {
		body["ShowInClient"] = request.ShowInClient
	}

	if !dara.IsNil(request.UpstreamHost) {
		body["UpstreamHost"] = request.UpstreamHost
	}

	if !dara.IsNil(request.UpstreamPort) {
		body["UpstreamPort"] = request.UpstreamPort
	}

	if !dara.IsNil(request.UpstreamType) {
		body["UpstreamType"] = request.UpstreamType
	}

	if !dara.IsNil(request.UserAttributeGroup) {
		body["UserAttributeGroup"] = request.UserAttributeGroup
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建加速对象
//
// @param request - CreateEnterpriseAccelerateTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseAccelerateTargetResponse
func (client *Client) CreateEnterpriseAccelerateTargetWithContext(ctx context.Context, request *CreateEnterpriseAccelerateTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseAccelerateTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		bodyFlat["Target"] = request.Target
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseAccelerateTarget"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseAccelerateTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义身份源部门
//
// @param request - CreateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdpDepartmentResponse
func (client *Client) CreateIdpDepartmentWithContext(ctx context.Context, request *CreateIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *CreateIdpDepartmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdpDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an office application within the current Alibaba Cloud account.
//
// Description:
//
// By default, you can create a maximum of 500 office applications.
//
// @param tmpReq - CreatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessApplicationResponse
func (client *Client) CreatePrivateAccessApplicationWithContext(ctx context.Context, tmpReq *CreatePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.L7Config) {
		request.L7ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.L7Config, dara.String("L7Config"), dara.String("json"))
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Addresses) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !dara.IsNil(request.BrowserAccessStatus) {
		body["BrowserAccessStatus"] = request.BrowserAccessStatus
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.L7ConfigShrink) {
		body["L7Config"] = request.L7ConfigShrink
	}

	if !dara.IsNil(request.L7ProxyDomainAutomaticPrefix) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !dara.IsNil(request.L7ProxyDomainCustom) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PortRanges) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Private Access Policy
//
// Description:
//
// By default, up to 500 private access policies can be created.
//
// @param request - CreatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessPolicyResponse
func (client *Client) CreatePrivateAccessPolicyWithContext(ctx context.Context, request *CreatePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.CustomUserAttributes) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceAttributeAction) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !dara.IsNil(request.DeviceAttributeId) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PolicyAction) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TriggerTemplateId) {
		body["TriggerTemplateId"] = request.TriggerTemplateId
	}

	if !dara.IsNil(request.TrustedProcessGroupIds) {
		bodyFlat["TrustedProcessGroupIds"] = request.TrustedProcessGroupIds
	}

	if !dara.IsNil(request.TrustedProcessStatus) {
		body["TrustedProcessStatus"] = request.TrustedProcessStatus
	}

	if !dara.IsNil(request.TrustedSoftwareIds) {
		bodyFlat["TrustedSoftwareIds"] = request.TrustedSoftwareIds
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserGroupMode) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内网访问标签
//
// @param request - CreatePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessTagResponse
func (client *Client) CreatePrivateAccessTagWithContext(ctx context.Context, request *CreatePrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建设备注册策略
//
// @param tmpReq - CreateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRegistrationPolicyResponse
func (client *Client) CreateRegistrationPolicyWithContext(ctx context.Context, tmpReq *CreateRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateRegistrationPolicyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompanyLimitCount) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, dara.String("CompanyLimitCount"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PersonalLimitCount) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, dara.String("PersonalLimitCount"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyLimitCountShrink) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !dara.IsNil(request.CompanyLimitType) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalLimitCountShrink) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !dara.IsNil(request.PersonalLimitType) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRegistrationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户组
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数字水印暗水印透明底图
//
// @param tmpReq - CreateWmBaseImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmBaseImageResponse
func (client *Client) CreateWmBaseImageWithContext(ctx context.Context, tmpReq *CreateWmBaseImageRequest, runtime *dara.RuntimeOptions) (_result *CreateWmBaseImageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateWmBaseImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageControl) {
		request.ImageControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageControl, dara.String("ImageControl"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["comment"] = request.Comment
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.ImageControlShrink) {
		body["ImageControl"] = request.ImageControlShrink
	}

	if !dara.IsNil(request.Opacity) {
		body["Opacity"] = request.Opacity
	}

	if !dara.IsNil(request.Scale) {
		body["Scale"] = request.Scale
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmInfoUint) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmBaseImage"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmBaseImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建嵌入水印任务
//
// @param tmpReq - CreateWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmEmbedTaskResponse
func (client *Client) CreateWmEmbedTaskWithContext(ctx context.Context, tmpReq *CreateWmEmbedTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWmEmbedTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateWmEmbedTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CsvControl) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, dara.String("CsvControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocumentControl) {
		request.DocumentControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentControl, dara.String("DocumentControl"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CsvControlShrink) {
		query["CsvControl"] = request.CsvControlShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentControlShrink) {
		body["DocumentControl"] = request.DocumentControlShrink
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Filename) {
		body["Filename"] = request.Filename
	}

	if !dara.IsNil(request.ImageEmbedJpegQuality) {
		body["ImageEmbedJpegQuality"] = request.ImageEmbedJpegQuality
	}

	if !dara.IsNil(request.ImageEmbedLevel) {
		body["ImageEmbedLevel"] = request.ImageEmbedLevel
	}

	if !dara.IsNil(request.VideoBitrate) {
		body["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.VideoIsLong) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmInfoUint) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmEmbedTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmEmbedTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a digital watermark extraction task.
//
// @param tmpReq - CreateWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmExtractTaskResponse
func (client *Client) CreateWmExtractTaskWithContext(ctx context.Context, tmpReq *CreateWmExtractTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWmExtractTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateWmExtractTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CsvControl) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, dara.String("CsvControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ImageExtractParamsOpenApi) {
		request.ImageExtractParamsOpenApiShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageExtractParamsOpenApi, dara.String("ImageExtractParamsOpenApi"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CsvControlShrink) {
		query["CsvControl"] = request.CsvControlShrink
	}

	if !dara.IsNil(request.ImageExtractParamsOpenApiShrink) {
		query["ImageExtractParamsOpenApi"] = request.ImageExtractParamsOpenApiShrink
	}

	if !dara.IsNil(request.IsClientEmbed) {
		query["IsClientEmbed"] = request.IsClientEmbed
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIsCapture) {
		body["DocumentIsCapture"] = request.DocumentIsCapture
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Filename) {
		body["Filename"] = request.Filename
	}

	if !dara.IsNil(request.VideoIsLong) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !dara.IsNil(request.VideoSpeed) {
		body["VideoSpeed"] = request.VideoSpeed
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmExtractTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmExtractTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一条字符串水印信息到数字水印信息的映射记录
//
// @param request - CreateWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmInfoMappingResponse
func (client *Client) CreateWmInfoMappingWithContext(ctx context.Context, request *CreateWmInfoMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateWmInfoMappingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmInfoMapping"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmInfoMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除审批流程
//
// @param request - DeleteApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApprovalProcessesResponse
func (client *Client) DeleteApprovalProcessesWithContext(ctx context.Context, request *DeleteApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *DeleteApprovalProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessIds) {
		bodyFlat["ProcessIds"] = request.ProcessIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApprovalProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义身份源指定用户
//
// @param request - DeleteClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientUserResponse
func (client *Client) DeleteClientUserWithContext(ctx context.Context, request *DeleteClientUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除动态路由
//
// @param request - DeleteDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicRouteResponse
func (client *Client) DeleteDynamicRouteWithContext(ctx context.Context, request *DeleteDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicRouteId) {
		query["DynamicRouteId"] = request.DynamicRouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除加速策略
//
// @param request - DeleteEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseAcceleratePolicyResponse
func (client *Client) DeleteEnterpriseAcceleratePolicyWithContext(ctx context.Context, request *DeleteEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseAcceleratePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除加速对象
//
// @param request - DeleteEnterpriseAccelerateTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseAccelerateTargetResponse
func (client *Client) DeleteEnterpriseAccelerateTargetWithContext(ctx context.Context, request *DeleteEnterpriseAccelerateTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseAccelerateTargetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		bodyFlat["Target"] = request.Target
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterpriseAccelerateTarget"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseAccelerateTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定自定义身份源部门
//
// @param request - DeleteIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdpDepartmentResponse
func (client *Client) DeleteIdpDepartmentWithContext(ctx context.Context, request *DeleteIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdpDepartmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdpDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteOtpConfig
//
// @param request - DeleteOtpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOtpConfigResponse
func (client *Client) DeleteOtpConfigWithContext(ctx context.Context, request *DeleteOtpConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteOtpConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOtpConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOtpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问应用
//
// @param request - DeletePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessApplicationResponse
func (client *Client) DeletePrivateAccessApplicationWithContext(ctx context.Context, request *DeletePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问策略
//
// @param request - DeletePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessPolicyResponse
func (client *Client) DeletePrivateAccessPolicyWithContext(ctx context.Context, request *DeletePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问标签
//
// @param request - DeletePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessTagResponse
func (client *Client) DeletePrivateAccessTagWithContext(ctx context.Context, request *DeletePrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TagId) {
		body["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备注册策略
//
// @param request - DeleteRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistrationPoliciesResponse
func (client *Client) DeleteRegistrationPoliciesWithContext(ctx context.Context, request *DeleteRegistrationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRegistrationPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIds) {
		bodyFlat["PolicyIds"] = request.PolicyIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistrationPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistrationPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除用户非在线设备
//
// @param request - DeleteUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDevicesResponse
func (client *Client) DeleteUserDevicesWithContext(ctx context.Context, request *DeleteUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户组
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		body["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载connector的应用
//
// @param tmpReq - DetachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachApplication2ConnectorResponse
func (client *Client) DetachApplication2ConnectorWithContext(ctx context.Context, tmpReq *DetachApplication2ConnectorRequest, runtime *dara.RuntimeOptions) (_result *DetachApplication2ConnectorResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DetachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationIds) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, dara.String("ApplicationIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIdsShrink) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachApplication2Connector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachApplication2ConnectorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑业务策略与审批流程
//
// @param request - DetachPolicy2ApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicy2ApprovalProcessResponse
func (client *Client) DetachPolicy2ApprovalProcessWithContext(ctx context.Context, request *DetachPolicy2ApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicy2ApprovalProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPolicy2ApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPolicy2ApprovalProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用加速策略
//
// @param request - DisableEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEnterpriseAcceleratePolicyResponse
func (client *Client) DisableEnterpriseAcceleratePolicyWithContext(ctx context.Context, request *DisableEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableEnterpriseAcceleratePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用加速策略
//
// @param request - EnableEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEnterpriseAcceleratePolicyResponse
func (client *Client) EnableEnterpriseAcceleratePolicyWithContext(ctx context.Context, request *EnableEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableEnterpriseAcceleratePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ExportUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportUserDevicesResponse
func (client *Client) ExportUserDevicesWithContext(ctx context.Context, request *ExportUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *ExportUserDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AppStatuses) {
		bodyFlat["AppStatuses"] = request.AppStatuses
	}

	if !dara.IsNil(request.Department) {
		body["Department"] = request.Department
	}

	if !dara.IsNil(request.DeviceBelong) {
		body["DeviceBelong"] = request.DeviceBelong
	}

	if !dara.IsNil(request.DeviceStatuses) {
		bodyFlat["DeviceStatuses"] = request.DeviceStatuses
	}

	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !dara.IsNil(request.DeviceTypes) {
		bodyFlat["DeviceTypes"] = request.DeviceTypes
	}

	if !dara.IsNil(request.DlpStatuses) {
		bodyFlat["DlpStatuses"] = request.DlpStatuses
	}

	if !dara.IsNil(request.Hostname) {
		body["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.IaStatuses) {
		bodyFlat["IaStatuses"] = request.IaStatuses
	}

	if !dara.IsNil(request.Mac) {
		body["Mac"] = request.Mac
	}

	if !dara.IsNil(request.NacStatuses) {
		bodyFlat["NacStatuses"] = request.NacStatuses
	}

	if !dara.IsNil(request.PaStatuses) {
		bodyFlat["PaStatuses"] = request.PaStatuses
	}

	if !dara.IsNil(request.SaseUserId) {
		body["SaseUserId"] = request.SaseUserId
	}

	if !dara.IsNil(request.SharingStatus) {
		body["SharingStatus"] = request.SharingStatus
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportUserDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批
//
// @param request - GetApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalResponse
func (client *Client) GetApprovalWithContext(ctx context.Context, request *GetApprovalRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApproval"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批流程
//
// @param request - GetApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalProcessResponse
func (client *Client) GetApprovalProcessWithContext(ctx context.Context, request *GetApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批动态模板
//
// @param request - GetApprovalSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalSchemaResponse
func (client *Client) GetApprovalSchemaWithContext(ctx context.Context, request *GetApprovalSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalSchemaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApprovalSchema"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源指定用户
//
// @param request - GetClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientUserResponse
func (client *Client) GetClientUserWithContext(ctx context.Context, request *GetClientUserRequest, runtime *dara.RuntimeOptions) (_result *GetClientUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动态路由详情
//
// @param request - GetDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDynamicRouteResponse
func (client *Client) GetDynamicRouteWithContext(ctx context.Context, request *GetDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *GetDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份源配置详情
//
// @param request - GetIdpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdpConfigResponse
func (client *Client) GetIdpConfigWithContext(ctx context.Context, request *GetIdpConfigRequest, runtime *dara.RuntimeOptions) (_result *GetIdpConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdpConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the office applications that belong to the current Alibaba Cloud account.
//
// @param request - GetPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessApplicationResponse
func (client *Client) GetPrivateAccessApplicationWithContext(ctx context.Context, request *GetPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetPrivateAccessApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Intranet Access Policy Details
//
// @param request - GetPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessPolicyResponse
func (client *Client) GetPrivateAccessPolicyWithContext(ctx context.Context, request *GetPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备注册策略详情
//
// @param request - GetRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistrationPolicyResponse
func (client *Client) GetRegistrationPolicyWithContext(ctx context.Context, request *GetRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetRegistrationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegistrationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户设备详情
//
// @param request - GetUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeviceResponse
func (client *Client) GetUserDeviceWithContext(ctx context.Context, request *GetUserDeviceRequest, runtime *dara.RuntimeOptions) (_result *GetUserDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserDevice"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户组详情
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithContext(ctx context.Context, request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询嵌入水印任务
//
// @param request - GetWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmEmbedTaskResponse
func (client *Client) GetWmEmbedTaskWithContext(ctx context.Context, request *GetWmEmbedTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWmEmbedTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWmEmbedTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWmEmbedTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文件水印提取任务详情
//
// @param request - GetWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmExtractTaskResponse
func (client *Client) GetWmExtractTaskWithContext(ctx context.Context, request *GetWmExtractTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWmExtractTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWmExtractTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWmExtractTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量导入加速对象异步任务
//
// @param request - ImportEnterpriseAccelerateTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportEnterpriseAccelerateTargetsResponse
func (client *Client) ImportEnterpriseAccelerateTargetsWithContext(ctx context.Context, request *ImportEnterpriseAccelerateTargetsRequest, runtime *dara.RuntimeOptions) (_result *ImportEnterpriseAccelerateTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportEnterpriseAccelerateTargets"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的应用
//
// @param request - ListApplicationsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessPolicyResponse
func (client *Client) ListApplicationsForPrivateAccessPolicyWithContext(ctx context.Context, request *ListApplicationsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForPrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的应用
//
// @param request - ListApplicationsForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessTagResponse
func (client *Client) ListApplicationsForPrivateAccessTagWithContext(ctx context.Context, request *ListApplicationsForPrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForPrivateAccessTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForPrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询审批流程
//
// @param request - ListApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalProcessesResponse
func (client *Client) ListApprovalProcessesWithContext(ctx context.Context, request *ListApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批渲染模板关联的流程
//
// @param request - ListApprovalProcessesForApprovalSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalProcessesForApprovalSchemasResponse
func (client *Client) ListApprovalProcessesForApprovalSchemasWithContext(ctx context.Context, request *ListApprovalProcessesForApprovalSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalProcessesForApprovalSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalProcessesForApprovalSchemas"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalProcessesForApprovalSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询审批动态模板
//
// @param request - ListApprovalSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalSchemasResponse
func (client *Client) ListApprovalSchemasWithContext(ctx context.Context, request *ListApprovalSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalSchemas"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批流程关联的渲染模板
//
// @param request - ListApprovalSchemasForApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalSchemasForApprovalProcessesResponse
func (client *Client) ListApprovalSchemasForApprovalProcessesWithContext(ctx context.Context, request *ListApprovalSchemasForApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalSchemasForApprovalProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalSchemasForApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalSchemasForApprovalProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询审批
//
// @param request - ListApprovalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalsResponse
func (client *Client) ListApprovalsWithContext(ctx context.Context, request *ListApprovalsRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovals"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源用户
//
// @param request - ListClientUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientUsersResponse
func (client *Client) ListClientUsersWithContext(ctx context.Context, request *ListClientUsersRequest, runtime *dara.RuntimeOptions) (_result *ListClientUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientUsers"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询connector
//
// @param request - ListConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectorsResponse
func (client *Client) ListConnectorsWithContext(ctx context.Context, request *ListConnectorsRequest, runtime *dara.RuntimeOptions) (_result *ListConnectorsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnectors"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectorsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Query Dynamic Policy Disposal Processes
//
// @param request - ListDynamicDisposalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicDisposalProcessesResponse
func (client *Client) ListDynamicDisposalProcessesWithContext(ctx context.Context, request *ListDynamicDisposalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicDisposalProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicDisposalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicDisposalProcessesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由
//
// @param request - ListDynamicRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicRoutesResponse
func (client *Client) ListDynamicRoutesWithContext(ctx context.Context, request *ListDynamicRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicRoutesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicRoutes"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询加速策略日志列表
//
// @param request - ListEnterpriseAccelerateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAccelerateLogsResponse
func (client *Client) ListEnterpriseAccelerateLogsWithContext(ctx context.Context, request *ListEnterpriseAccelerateLogsRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAccelerateLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAccelerateLogs"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAccelerateLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询加速策略列表
//
// @param request - ListEnterpriseAcceleratePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAcceleratePoliciesResponse
func (client *Client) ListEnterpriseAcceleratePoliciesWithContext(ctx context.Context, request *ListEnterpriseAcceleratePoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAcceleratePoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAcceleratePolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAcceleratePoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询加速对象列表
//
// @param request - ListEnterpriseAccelerateTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAccelerateTargetsResponse
func (client *Client) ListEnterpriseAccelerateTargetsWithContext(ctx context.Context, request *ListEnterpriseAccelerateTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAccelerateTargetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAccelerateTargets"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询超额注册申请列表
//
// @param request - ListExcessiveDeviceRegistrationApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExcessiveDeviceRegistrationApplicationsResponse
func (client *Client) ListExcessiveDeviceRegistrationApplicationsWithContext(ctx context.Context, request *ListExcessiveDeviceRegistrationApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListExcessiveDeviceRegistrationApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExcessiveDeviceRegistrationApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExcessiveDeviceRegistrationApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询IDP配置
//
// @param request - ListIdpConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpConfigsResponse
func (client *Client) ListIdpConfigsWithContext(ctx context.Context, request *ListIdpConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListIdpConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdpConfigs"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdpConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源部门
//
// @param request - ListIdpDepartmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpDepartmentsResponse
func (client *Client) ListIdpDepartmentsWithContext(ctx context.Context, request *ListIdpDepartmentsRequest, runtime *dara.RuntimeOptions) (_result *ListIdpDepartmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdpDepartments"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdpDepartmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 入网用户列表
//
// @param request - ListNacUserCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacUserCertResponse
func (client *Client) ListNacUserCertWithContext(ctx context.Context, request *ListNacUserCertRequest, runtime *dara.RuntimeOptions) (_result *ListNacUserCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNacUserCert"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNacUserCertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的策略
//
// @param request - ListPolicesForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessApplicationResponse
func (client *Client) ListPolicesForPrivateAccessApplicationWithContext(ctx context.Context, request *ListPolicesForPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForPrivateAccessApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的策略
//
// @param request - ListPolicesForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessTagResponse
func (client *Client) ListPolicesForPrivateAccessTagWithContext(ctx context.Context, request *ListPolicesForPrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForPrivateAccessTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForPrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForPrivateAccessTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户组的策略
//
// @param request - ListPolicesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForUserGroupResponse
func (client *Client) ListPolicesForUserGroupWithContext(ctx context.Context, request *ListPolicesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pop节点流量统计
//
// @param request - ListPopTrafficStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPopTrafficStatisticsResponse
func (client *Client) ListPopTrafficStatisticsWithContext(ctx context.Context, request *ListPopTrafficStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListPopTrafficStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPopTrafficStatistics"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPopTrafficStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用
//
// @param request - ListPrivateAccessApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsResponse
func (client *Client) ListPrivateAccessApplicationsWithContext(ctx context.Context, request *ListPrivateAccessApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessModes) {
		query["AccessModes"] = request.AccessModes
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问应用
//
// @param request - ListPrivateAccessApplicationsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsForDynamicRouteResponse
func (client *Client) ListPrivateAccessApplicationsForDynamicRouteWithContext(ctx context.Context, request *ListPrivateAccessApplicationsForDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessApplicationsForDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessApplicationsForDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessApplicationsForDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the private access policies within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessPolicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessPolicesResponse
func (client *Client) ListPrivateAccessPolicesWithContext(ctx context.Context, request *ListPrivateAccessPolicesRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessPolicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessPolices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessPolicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all internal access tags within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsResponse
func (client *Client) ListPrivateAccessTagsWithContext(ctx context.Context, request *ListPrivateAccessTagsRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessTags"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问标签
//
// @param request - ListPrivateAccessTagsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsForDynamicRouteResponse
func (client *Client) ListPrivateAccessTagsForDynamicRouteWithContext(ctx context.Context, request *ListPrivateAccessTagsForDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessTagsForDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessTagsForDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessTagsForDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户设备注册策略列表
//
// @param request - ListRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesResponse
func (client *Client) ListRegistrationPoliciesWithContext(ctx context.Context, request *ListRegistrationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListRegistrationPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistrationPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistrationPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户组相关的设备注册策略
//
// @param request - ListRegistrationPoliciesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesForUserGroupResponse
func (client *Client) ListRegistrationPoliciesForUserGroupWithContext(ctx context.Context, request *ListRegistrationPoliciesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListRegistrationPoliciesForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistrationPoliciesForUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistrationPoliciesForUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询终端安装软件列表
//
// @param request - ListSoftwareForUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwareForUserDeviceResponse
func (client *Client) ListSoftwareForUserDeviceWithContext(ctx context.Context, request *ListSoftwareForUserDeviceRequest, runtime *dara.RuntimeOptions) (_result *ListSoftwareForUserDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSoftwareForUserDevice"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSoftwareForUserDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的标签
//
// @param request - ListTagsForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessApplicationResponse
func (client *Client) ListTagsForPrivateAccessApplicationWithContext(ctx context.Context, request *ListTagsForPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListTagsForPrivateAccessApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagsForPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的标签
//
// @param request - ListTagsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessPolicyResponse
func (client *Client) ListTagsForPrivateAccessPolicyWithContext(ctx context.Context, request *ListTagsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListTagsForPrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询卸载申请列表
//
// @param request - ListUninstallApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUninstallApplicationsResponse
func (client *Client) ListUninstallApplicationsWithContext(ctx context.Context, request *ListUninstallApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListUninstallApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUninstallApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUninstallApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询用户应用权限
//
// @param request - ListUserApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserApplicationsResponse
func (client *Client) ListUserApplicationsWithContext(ctx context.Context, request *ListUserApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListUserApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ListUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDevicesResponse
func (client *Client) ListUserDevicesWithContext(ctx context.Context, request *ListUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListUserDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户组
//
// @param request - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithContext(ctx context.Context, request *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的用户组
//
// @param request - ListUserGroupsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForPrivateAccessPolicyResponse
func (client *Client) ListUserGroupsForPrivateAccessPolicyWithContext(ctx context.Context, request *ListUserGroupsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsForPrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备注册策略相关用户组
//
// @param request - ListUserGroupsForRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForRegistrationPolicyResponse
func (client *Client) ListUserGroupsForRegistrationPolicyWithContext(ctx context.Context, request *ListUserGroupsForRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsForRegistrationPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsForRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsForRegistrationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List query of user zero trust policies
//
// @param request - ListUserPrivateAccessPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPrivateAccessPoliciesResponse
func (client *Client) ListUserPrivateAccessPoliciesWithContext(ctx context.Context, request *ListUserPrivateAccessPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListUserPrivateAccessPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPrivateAccessPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPrivateAccessPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询登陆用户
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据数字水印信息查询字符串水印信息
//
// @param request - LookupWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupWmInfoMappingResponse
func (client *Client) LookupWmInfoMappingWithContext(ctx context.Context, request *LookupWmInfoMappingRequest, runtime *dara.RuntimeOptions) (_result *LookupWmInfoMappingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupWmInfoMapping"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupWmInfoMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改加速策略
//
// @param request - ModifyEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEnterpriseAcceleratePolicyResponse
func (client *Client) ModifyEnterpriseAcceleratePolicyWithContext(ctx context.Context, request *ModifyEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyEnterpriseAcceleratePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationType) {
		body["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OnTls) {
		body["OnTls"] = request.OnTls
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ShowInClient) {
		body["ShowInClient"] = request.ShowInClient
	}

	if !dara.IsNil(request.UpstreamHost) {
		body["UpstreamHost"] = request.UpstreamHost
	}

	if !dara.IsNil(request.UpstreamPort) {
		body["UpstreamPort"] = request.UpstreamPort
	}

	if !dara.IsNil(request.UpstreamType) {
		body["UpstreamType"] = request.UpstreamType
	}

	if !dara.IsNil(request.UserAttributeGroup) {
		body["UserAttributeGroup"] = request.UserAttributeGroup
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RevokeUserSession is deprecated
//
// Summary:
//
// 吊销用户登录会话
//
// @param request - RevokeUserSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeUserSessionResponse
func (client *Client) RevokeUserSessionWithContext(ctx context.Context, request *RevokeUserSessionRequest, runtime *dara.RuntimeOptions) (_result *RevokeUserSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIds) {
		query["ExternalIds"] = request.ExternalIds
	}

	if !dara.IsNil(request.IdpId) {
		query["IdpId"] = request.IdpId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeUserSession"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeUserSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新审批流程
//
// @param tmpReq - UpdateApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApprovalProcessResponse
func (client *Client) UpdateApprovalProcessWithContext(ctx context.Context, tmpReq *UpdateApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *UpdateApprovalProcessResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateApprovalProcessShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchSchemas) {
		request.MatchSchemasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchSchemas, dara.String("MatchSchemas"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchSchemasShrink) {
		body["MatchSchemas"] = request.MatchSchemasShrink
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProcessName) {
		body["ProcessName"] = request.ProcessName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessNodes) {
		bodyFlat["ProcessNodes"] = request.ProcessNodes
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApprovalProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改审批状态
//
// @param request - UpdateApprovalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApprovalStatusResponse
func (client *Client) UpdateApprovalStatusWithContext(ctx context.Context, request *UpdateApprovalStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateApprovalStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalId) {
		query["ApprovalId"] = request.ApprovalId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApprovalStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApprovalStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自启动与防卸载策略配置
//
// @param tmpReq - UpdateBootAndAntiUninstallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBootAndAntiUninstallPolicyResponse
func (client *Client) UpdateBootAndAntiUninstallPolicyWithContext(ctx context.Context, tmpReq *UpdateBootAndAntiUninstallPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateBootAndAntiUninstallPolicyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBootAndAntiUninstallPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BlockContent) {
		request.BlockContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BlockContent, dara.String("BlockContent"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowReport) {
		body["AllowReport"] = request.AllowReport
	}

	if !dara.IsNil(request.BlockContentShrink) {
		body["BlockContent"] = request.BlockContentShrink
	}

	if !dara.IsNil(request.IsAntiUninstall) {
		body["IsAntiUninstall"] = request.IsAntiUninstall
	}

	if !dara.IsNil(request.IsBoot) {
		body["IsBoot"] = request.IsBoot
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.WhitelistUsers) {
		bodyFlat["WhitelistUsers"] = request.WhitelistUsers
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBootAndAntiUninstallPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBootAndAntiUninstallPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户
//
// @param request - UpdateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserResponse
func (client *Client) UpdateClientUserWithContext(ctx context.Context, request *UpdateClientUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MobileNumber) {
		query["MobileNumber"] = request.MobileNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户密码
//
// @param request - UpdateClientUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserPasswordResponse
func (client *Client) UpdateClientUserPasswordWithContext(ctx context.Context, request *UpdateClientUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUserPassword"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户启用状态
//
// @param request - UpdateClientUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserStatusResponse
func (client *Client) UpdateClientUserStatusWithContext(ctx context.Context, request *UpdateClientUserStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUserStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改动态路由
//
// @param request - UpdateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicRouteResponse
func (client *Client) UpdateDynamicRouteWithContext(ctx context.Context, request *UpdateDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateDynamicRouteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicRouteId) {
		body["DynamicRouteId"] = request.DynamicRouteId
	}

	if !dara.IsNil(request.DynamicRouteType) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextHop) {
		body["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionIds) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDynamicRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新超额注册申请状态
//
// @param request - UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatusWithContext(ctx context.Context, request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExcessiveDeviceRegistrationApplicationsStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改指定自定义身份源部门
//
// @param request - UpdateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdpDepartmentResponse
func (client *Client) UpdateIdpDepartmentWithContext(ctx context.Context, request *UpdateIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdpDepartmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdpDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新NAC User 状态
//
// @param request - UpdateNacUserCertStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacUserCertStatusResponse
func (client *Client) UpdateNacUserCertStatusWithContext(ctx context.Context, request *UpdateNacUserCertStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacUserCertStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.IdList) {
		bodyFlat["IdList"] = request.IdList
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacUserCertStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacUserCertStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the office applications of the current Alibaba Cloud account.
//
// @param tmpReq - UpdatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessApplicationResponse
func (client *Client) UpdatePrivateAccessApplicationWithContext(ctx context.Context, tmpReq *UpdatePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.L7Config) {
		request.L7ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.L7Config, dara.String("L7Config"), dara.String("json"))
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Addresses) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.L7ConfigShrink) {
		body["L7Config"] = request.L7ConfigShrink
	}

	if !dara.IsNil(request.L7ProxyDomainAutomaticPrefix) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !dara.IsNil(request.L7ProxyDomainCustom) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !dara.IsNil(request.L7ProxyDomainPrivate) {
		body["L7ProxyDomainPrivate"] = request.L7ProxyDomainPrivate
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.PortRanges) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Private Access Policy
//
// @param request - UpdatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessPolicyResponse
func (client *Client) UpdatePrivateAccessPolicyWithContext(ctx context.Context, request *UpdatePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.CustomUserAttributes) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceAttributeAction) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !dara.IsNil(request.DeviceAttributeId) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.PolicyAction) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TriggerTemplateId) {
		body["TriggerTemplateId"] = request.TriggerTemplateId
	}

	if !dara.IsNil(request.TrustedProcessGroupIds) {
		bodyFlat["TrustedProcessGroupIds"] = request.TrustedProcessGroupIds
	}

	if !dara.IsNil(request.TrustedProcessStatus) {
		body["TrustedProcessStatus"] = request.TrustedProcessStatus
	}

	if !dara.IsNil(request.TrustedSoftwareIds) {
		bodyFlat["TrustedSoftwareIds"] = request.TrustedSoftwareIds
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserGroupMode) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改设备注册策略
//
// @param tmpReq - UpdateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRegistrationPolicyResponse
func (client *Client) UpdateRegistrationPolicyWithContext(ctx context.Context, tmpReq *UpdateRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRegistrationPolicyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompanyLimitCount) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, dara.String("CompanyLimitCount"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PersonalLimitCount) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, dara.String("PersonalLimitCount"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyLimitCountShrink) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !dara.IsNil(request.CompanyLimitType) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalLimitCountShrink) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !dara.IsNil(request.PersonalLimitType) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRegistrationPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改卸载申请状态
//
// @param request - UpdateUninstallApplicationsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUninstallApplicationsStatusResponse
func (client *Client) UpdateUninstallApplicationsStatusWithContext(ctx context.Context, request *UpdateUninstallApplicationsStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUninstallApplicationsStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUninstallApplicationsStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUninstallApplicationsStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新用户设备共享状态
//
// @param request - UpdateUserDevicesSharingStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesSharingStatusResponse
func (client *Client) UpdateUserDevicesSharingStatusWithContext(ctx context.Context, request *UpdateUserDevicesSharingStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDevicesSharingStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !dara.IsNil(request.SharingStatus) {
		body["SharingStatus"] = request.SharingStatus
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDevicesSharingStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDevicesSharingStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新用户设备状态
//
// @param request - UpdateUserDevicesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesStatusResponse
func (client *Client) UpdateUserDevicesStatusWithContext(ctx context.Context, request *UpdateUserDevicesStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDevicesStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceAction) {
		body["DeviceAction"] = request.DeviceAction
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDevicesStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDevicesStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户组
//
// @param request - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithContext(ctx context.Context, request *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.UserGroupId) {
		body["UserGroupId"] = request.UserGroupId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改登陆用户状态
//
// @param request - UpdateUsersStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUsersStatusResponse
func (client *Client) UpdateUsersStatusWithContext(ctx context.Context, request *UpdateUsersStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUsersStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SaseUserIds) {
		query["SaseUserIds"] = request.SaseUserIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUsersStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUsersStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
