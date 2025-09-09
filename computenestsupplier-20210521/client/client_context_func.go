// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a shared account of a service.
//
// @param request - AddServiceSharedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServiceSharedAccountsResponse
func (client *Client) AddServiceSharedAccountsWithContext(ctx context.Context, request *AddServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *AddServiceSharedAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SharedAccounts) {
		query["SharedAccounts"] = request.SharedAccounts
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddServiceSharedAccounts"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddServiceSharedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Merchant uses service request
//
// @param request - ApproveServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveServiceUsageResponse
func (client *Client) ApproveServiceUsageWithContext(ctx context.Context, request *ApproveServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *ApproveServiceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAliUid) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveServiceUsage"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveServiceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel service registration.
//
// Description:
//
// Only service registration in the Submitted status can be canceled.
//
// @param request - CancelServiceRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelServiceRegistrationResponse
func (client *Client) CancelServiceRegistrationWithContext(ctx context.Context, request *CancelServiceRegistrationRequest, runtime *dara.RuntimeOptions) (_result *CancelServiceRegistrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegistrationId) {
		query["RegistrationId"] = request.RegistrationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelServiceRegistration"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelServiceRegistrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Redeploys a service instance after the service instance failed to be deployed.
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstanceWithContext(ctx context.Context, request *ContinueDeployServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueDeployServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a deployment package.
//
// @param tmpReq - CreateArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactResponse
func (client *Client) CreateArtifactWithContext(ctx context.Context, tmpReq *CreateArtifactRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ArtifactBuildProperty) {
		request.ArtifactBuildPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactBuildProperty, dara.String("ArtifactBuildProperty"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ArtifactProperty) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, dara.String("ArtifactProperty"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactBuildPropertyShrink) {
		query["ArtifactBuildProperty"] = request.ArtifactBuildPropertyShrink
	}

	if !dara.IsNil(request.ArtifactBuildType) {
		query["ArtifactBuildType"] = request.ArtifactBuildType
	}

	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactPropertyShrink) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SupportRegionIds) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateArtifact"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service.
//
// @param tmpReq - CreateServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithContext(ctx context.Context, tmpReq *CreateServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComplianceMetadata) {
		request.ComplianceMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComplianceMetadata, dara.String("ComplianceMetadata"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMetadata) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !dara.IsNil(request.ApprovalType) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !dara.IsNil(request.BuildParameters) {
		query["BuildParameters"] = request.BuildParameters
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ComplianceMetadataShrink) {
		query["ComplianceMetadata"] = request.ComplianceMetadataShrink
	}

	if !dara.IsNil(request.DeployMetadata) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.IsSupportOperated) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !dara.IsNil(request.LicenseMetadata) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !dara.IsNil(request.LogMetadata) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !dara.IsNil(request.OperationMetadata) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !dara.IsNil(request.PolicyNames) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resellable) {
		query["Resellable"] = request.Resellable
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInfo) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.SourceServiceId) {
		query["SourceServiceId"] = request.SourceServiceId
	}

	if !dara.IsNil(request.SourceServiceVersion) {
		query["SourceServiceVersion"] = request.SourceServiceVersion
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TenantType) {
		query["TenantType"] = request.TenantType
	}

	if !dara.IsNil(request.TrialDuration) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !dara.IsNil(request.UpgradeMetadata) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and deploys a service instance.
//
// @param tmpReq - CreateServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstanceWithContext(ctx context.Context, tmpReq *CreateServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.SpecificationName) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create service test case.
//
// @param request - CreateServiceTestCaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceTestCaseResponse
func (client *Client) CreateServiceTestCaseWithContext(ctx context.Context, request *CreateServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceTestCaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TestCaseName) {
		query["TestCaseName"] = request.TestCaseName
	}

	if !dara.IsNil(request.TestConfig) {
		query["TestConfig"] = request.TestConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceTestCase"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceTestCaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create service test task
//
// @param request - CreateServiceTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceTestTaskResponse
func (client *Client) CreateServiceTestTaskWithContext(ctx context.Context, request *CreateServiceTestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceTestTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskRegionId) {
		query["TaskRegionId"] = request.TaskRegionId
	}

	if !dara.IsNil(request.TestCaseIds) {
		query["TestCaseIds"] = request.TestCaseIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceTestTask"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceTestTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create  Service resell application.
//
// @param request - CreateServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceUsageResponse
func (client *Client) CreateServiceUsageWithContext(ctx context.Context, request *CreateServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceUsage"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册成为服务商
//
// @param request - CreateSupplierRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSupplierRegistrationResponse
func (client *Client) CreateSupplierRegistrationWithContext(ctx context.Context, request *CreateSupplierRegistrationRequest, runtime *dara.RuntimeOptions) (_result *CreateSupplierRegistrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactEmail) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactNumber) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !dara.IsNil(request.ContactPerson) {
		query["ContactPerson"] = request.ContactPerson
	}

	if !dara.IsNil(request.ContactPersonTitle) {
		query["ContactPersonTitle"] = request.ContactPersonTitle
	}

	if !dara.IsNil(request.EnableResellerMode) {
		query["EnableResellerMode"] = request.EnableResellerMode
	}

	if !dara.IsNil(request.ProductAnnualRevenue) {
		query["ProductAnnualRevenue"] = request.ProductAnnualRevenue
	}

	if !dara.IsNil(request.ProductBusiness) {
		query["ProductBusiness"] = request.ProductBusiness
	}

	if !dara.IsNil(request.ProductDeliveryTypes) {
		query["ProductDeliveryTypes"] = request.ProductDeliveryTypes
	}

	if !dara.IsNil(request.ProductPublishTime) {
		query["ProductPublishTime"] = request.ProductPublishTime
	}

	if !dara.IsNil(request.ProductSellTypes) {
		query["ProductSellTypes"] = request.ProductSellTypes
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellBusinessDesc) {
		query["ResellBusinessDesc"] = request.ResellBusinessDesc
	}

	if !dara.IsNil(request.Suggestion) {
		query["Suggestion"] = request.Suggestion
	}

	if !dara.IsNil(request.SupplierDesc) {
		query["SupplierDesc"] = request.SupplierDesc
	}

	if !dara.IsNil(request.SupplierLogo) {
		query["SupplierLogo"] = request.SupplierLogo
	}

	if !dara.IsNil(request.SupplierName) {
		query["SupplierName"] = request.SupplierName
	}

	if !dara.IsNil(request.SupplierNameEn) {
		query["SupplierNameEn"] = request.SupplierNameEn
	}

	if !dara.IsNil(request.SupplierUrl) {
		query["SupplierUrl"] = request.SupplierUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSupplierRegistration"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSupplierRegistrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Acr镜像仓库
//
// @param request - DeleteAcrImageRepositoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAcrImageRepositoriesResponse
func (client *Client) DeleteAcrImageRepositoriesWithContext(ctx context.Context, request *DeleteAcrImageRepositoriesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcrImageRepositoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAcrImageRepositories"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAcrImageRepositoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Acr容器镜像版本
//
// @param request - DeleteAcrImageTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAcrImageTagsResponse
func (client *Client) DeleteAcrImageTagsWithContext(ctx context.Context, request *DeleteAcrImageTagsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcrImageTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAcrImageTags"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAcrImageTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an artifact.
//
// @param request - DeleteArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifactWithContext(ctx context.Context, request *DeleteArtifactRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactVersion) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteArtifact"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @param request - DeleteServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithContext(ctx context.Context, request *DeleteServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service instance.
//
// @param request - DeleteServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithContext(ctx context.Context, request *DeleteServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceInstances"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete service test configuration
//
// @param request - DeleteServiceTestCaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceTestCaseResponse
func (client *Client) DeleteServiceTestCaseWithContext(ctx context.Context, request *DeleteServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceTestCaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TestCaseId) {
		query["TestCaseId"] = request.TestCaseId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceTestCase"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceTestCaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a service instance.
//
// @param request - DeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstanceWithContext(ctx context.Context, request *DeployServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成默认服务测试配置
//
// @param request - GenerateDefaultServiceTestConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDefaultServiceTestConfigResponse
func (client *Client) GenerateDefaultServiceTestConfigWithContext(ctx context.Context, request *GenerateDefaultServiceTestConfigRequest, runtime *dara.RuntimeOptions) (_result *GenerateDefaultServiceTestConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDefaultServiceTestConfig"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDefaultServiceTestConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成并校验服务创建stack所需要
//
// @param request - GenerateServicePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateServicePolicyResponse
func (client *Client) GenerateServicePolicyWithContext(ctx context.Context, request *GenerateServicePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateServicePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationTypes) {
		query["OperationTypes"] = request.OperationTypes
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TrialType) {
		query["TrialType"] = request.TrialType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateServicePolicy"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateServicePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a deployment package.
//
// @param request - GetArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactResponse
func (client *Client) GetArtifactWithContext(ctx context.Context, request *GetArtifactRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactName) {
		query["ArtifactName"] = request.ArtifactName
	}

	if !dara.IsNil(request.ArtifactVersion) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArtifact"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the credentials that are required to upload a deployment package.
//
// @param request - GetArtifactRepositoryCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactRepositoryCredentialsResponse
func (client *Client) GetArtifactRepositoryCredentialsWithContext(ctx context.Context, request *GetArtifactRepositoryCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.DeployRegionId) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArtifactRepositoryCredentials"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据地域参数获取地域可用区列表
//
// @param request - GetNetworkAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkAvailableZonesResponse
func (client *Client) GetNetworkAvailableZonesWithContext(ctx context.Context, request *GetNetworkAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkAvailableZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsPoc) {
		body["IsPoc"] = request.IsPoc
	}

	if !dara.IsNil(request.NetworkRegionId) {
		body["NetworkRegionId"] = request.NetworkRegionId
	}

	if !dara.IsNil(request.PrivateVpcConnectionMode) {
		body["PrivateVpcConnectionMode"] = request.PrivateVpcConnectionMode
	}

	if !dara.IsNil(request.ServiceId) {
		body["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceEndpointServiceType) {
		body["ServiceInstanceEndpointServiceType"] = request.ServiceInstanceEndpointServiceType
	}

	if !dara.IsNil(request.ServiceRegionId) {
		body["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.ServiceVersion) {
		body["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkAvailableZones"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkAvailableZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service.
//
// @param request - GetServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithContext(ctx context.Context, request *GetServiceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterAliUid) {
		query["FilterAliUid"] = request.FilterAliUid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.SharedAccountType) {
		query["SharedAccountType"] = request.SharedAccountType
	}

	if !dara.IsNil(request.ShowDetail) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated price for creating a service instance.
//
// @param tmpReq - GetServiceEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCostWithContext(ctx context.Context, tmpReq *GetServiceEstimateCostRequest, runtime *dara.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceEstimateCostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Commodity) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, dara.String("Commodity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommodityShrink) {
		query["Commodity"] = request.CommodityShrink
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.SpecificationName) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceEstimateCost"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service instance.
//
// @param request - GetServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstanceWithContext(ctx context.Context, request *GetServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the activation status and Resource Access Management (RAM) roles of the cloud services required by a service.
//
// @param tmpReq - GetServiceProvisionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisionsWithContext(ctx context.Context, tmpReq *GetServiceProvisionsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceProvisionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceProvisions"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get service registration detail.
//
// @param request - GetServiceRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceRegistrationResponse
func (client *Client) GetServiceRegistrationWithContext(ctx context.Context, request *GetServiceRegistrationRequest, runtime *dara.RuntimeOptions) (_result *GetServiceRegistrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegistrationId) {
		query["RegistrationId"] = request.RegistrationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceRegistration"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceRegistrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务模板存在的规范问题
//
// @param request - GetServiceTemplateCriterionIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateCriterionIssuesResponse
func (client *Client) GetServiceTemplateCriterionIssuesWithContext(ctx context.Context, request *GetServiceTemplateCriterionIssuesRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTemplateCriterionIssuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceTemplateCriterionIssues"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceTemplateCriterionIssuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the valid values of parameters in a Resource Orchestration Service (ROS) template.
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithContext(ctx context.Context, request *GetServiceTemplateParameterConstraintsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeployRegionId) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !dara.IsNil(request.EnablePrivateVpcConnection) {
		query["EnablePrivateVpcConnection"] = request.EnablePrivateVpcConnection
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceTemplateParameterConstraints"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the execution status of Cases in the service test task
//
// @param request - GetServiceTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTestTaskResponse
func (client *Client) GetServiceTestTaskWithContext(ctx context.Context, request *GetServiceTestTaskRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTestTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceTestTask"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceTestTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get service provider information
//
// @param request - GetSupplierInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplierInformationResponse
func (client *Client) GetSupplierInformationWithContext(ctx context.Context, request *GetSupplierInformationRequest, runtime *dara.RuntimeOptions) (_result *GetSupplierInformationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupplierInformation"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupplierInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the AccessKey pair of uploaded files.
//
// @param request - GetUploadCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadCredentialsResponse
func (client *Client) GetUploadCredentialsWithContext(ctx context.Context, request *GetUploadCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetUploadCredentialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadCredentials"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Launch service
//
// Description:
//
// 需要上线的服务必须为已通过审核的服务。
//
// @param request - LaunchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchServiceResponse
func (client *Client) LaunchServiceWithContext(ctx context.Context, request *LaunchServiceRequest, runtime *dara.RuntimeOptions) (_result *LaunchServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Categories) {
		query["Categories"] = request.Categories
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Recommend) {
		query["Recommend"] = request.Recommend
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LaunchService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LaunchServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of images uploaded to Container Registry.
//
// @param request - ListAcrImageRepositoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcrImageRepositoriesResponse
func (client *Client) ListAcrImageRepositoriesWithContext(ctx context.Context, request *ListAcrImageRepositoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAcrImageRepositoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RepoName) {
		query["RepoName"] = request.RepoName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAcrImageRepositories"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAcrImageRepositoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of images that are uploaded to the image repository.
//
// @param request - ListAcrImageTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAcrImageTagsResponse
func (client *Client) ListAcrImageTagsWithContext(ctx context.Context, request *ListAcrImageTagsRequest, runtime *dara.RuntimeOptions) (_result *ListAcrImageTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactType) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAcrImageTags"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAcrImageTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询部署实例、升级以及应用的日志
//
// @param request - ListArtifactBuildLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactBuildLogsResponse
func (client *Client) ListArtifactBuildLogsWithContext(ctx context.Context, request *ListArtifactBuildLogsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactBuildLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactVersion) {
		query["ArtifactVersion"] = request.ArtifactVersion
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListArtifactBuildLogs"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactBuildLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the list of artifact security risks
//
// @param request - ListArtifactRisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactRisksResponse
func (client *Client) ListArtifactRisksWithContext(ctx context.Context, request *ListArtifactRisksRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactRisksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactVersion) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListArtifactRisks"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactRisksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version information about a deployment package.
//
// @param tmpReq - ListArtifactVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactVersionsResponse
func (client *Client) ListArtifactVersionsWithContext(ctx context.Context, tmpReq *ListArtifactVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactVersionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListArtifactVersionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filters) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, dara.String("Filters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.FiltersShrink) {
		query["Filters"] = request.FiltersShrink
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
		Action:      dara.String("ListArtifactVersions"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of deployment packages.
//
// @param request - ListArtifactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactsResponse
func (client *Client) ListArtifactsWithContext(ctx context.Context, request *ListArtifactsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListArtifacts"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paginated query of distributor information list
//
// @param request - ListResellersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResellersResponse
func (client *Client) ListResellersWithContext(ctx context.Context, request *ListResellersRequest, runtime *dara.RuntimeOptions) (_result *ListResellersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResellers"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResellersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询部署实例、升级以及应用的日志
//
// @param request - ListServiceBuildLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceBuildLogsResponse
func (client *Client) ListServiceBuildLogsWithContext(ctx context.Context, request *ListServiceBuildLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceBuildLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceBuildLogs"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceBuildLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Display service instance bill
//
// @param request - ListServiceInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBillWithContext(ctx context.Context, request *ListServiceInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceBillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.BillingDate) {
		query["BillingDate"] = request.BillingDate
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceBill"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceBillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query service instance deployment details
//
// @param request - ListServiceInstanceDeployDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceDeployDetailsResponse
func (client *Client) ListServiceInstanceDeployDetailsWithContext(ctx context.Context, request *ListServiceInstanceDeployDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceDeployDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CycleTimeZone) {
		query["CycleTimeZone"] = request.CycleTimeZone
	}

	if !dara.IsNil(request.CycleType) {
		query["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceDeployDetails"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceDeployDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// “Query logs at various levels, including service instance application, instance, and resource.”
//
// @param request - ListServiceInstanceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogsWithContext(ctx context.Context, request *ListServiceInstanceLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.LogSource) {
		query["LogSource"] = request.LogSource
	}

	if !dara.IsNil(request.Logstore) {
		query["Logstore"] = request.Logstore
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceLogs"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务实例资源
//
// @param request - ListServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResourcesWithContext(ctx context.Context, request *ListServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceInstanceResourceType) {
		query["ServiceInstanceResourceType"] = request.ServiceInstanceResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceResources"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View the upgrade history of a service instance
//
// @param request - ListServiceInstanceUpgradeHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceUpgradeHistoryResponse
func (client *Client) ListServiceInstanceUpgradeHistoryWithContext(ctx context.Context, request *ListServiceInstanceUpgradeHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceUpgradeHistory"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceUpgradeHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of service instances.
//
// @param request - ListServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithContext(ctx context.Context, request *ListServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowDeleted) {
		query["ShowDeleted"] = request.ShowDeleted
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstances"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query service registrations.
//
// @param request - ListServiceRegistrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceRegistrationsResponse
func (client *Client) ListServiceRegistrationsWithContext(ctx context.Context, request *ListServiceRegistrationsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceRegistrationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceRegistrations"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceRegistrationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用ListServiceSharedAccounts查看服务共享账号列表。
//
// @param request - ListServiceSharedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceSharedAccountsResponse
func (client *Client) ListServiceSharedAccountsWithContext(ctx context.Context, request *ListServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceSharedAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceSharedAccounts"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceSharedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Service test case list
//
// @param request - ListServiceTestCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTestCasesResponse
func (client *Client) ListServiceTestCasesWithContext(ctx context.Context, request *ListServiceTestCasesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestCasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceTestCases"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceTestCasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get service test real-time logs
//
// @param request - ListServiceTestTaskLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTestTaskLogsResponse
func (client *Client) ListServiceTestTaskLogsWithContext(ctx context.Context, request *ListServiceTestTaskLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestTaskLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceTestTaskLogs"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceTestTaskLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务执行列表
//
// @param request - ListServiceTestTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTestTasksResponse
func (client *Client) ListServiceTestTasksWithContext(ctx context.Context, request *ListServiceTestTasksRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceTestTasks"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceTestTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications for using a service.
//
// @param request - ListServiceUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsagesWithContext(ctx context.Context, request *ListServiceUsagesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceUsages"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of services.
//
// @param request - ListServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithContext(ctx context.Context, request *ListServicesRequest, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllVersions) {
		query["AllVersions"] = request.AllVersions
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of service provider onboarding reviews
//
// @param request - ListSupplierRegistrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupplierRegistrationsResponse
func (client *Client) ListSupplierRegistrationsWithContext(ctx context.Context, request *ListSupplierRegistrationsRequest, runtime *dara.RuntimeOptions) (_result *ListSupplierRegistrationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupplierRegistrations"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupplierRegistrationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query tag key list
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithContext(ctx context.Context, request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2021-05-21"),
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
// # Query resource tags
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Version:     dara.String("2021-05-21"),
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
// # Query tag value list
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithContext(ctx context.Context, request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2021-05-21"),
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
// Modifies the resource information about a service instance.
//
// @param request - ModifyServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceInstanceResourcesResponse
func (client *Client) ModifyServiceInstanceResourcesWithContext(ctx context.Context, request *ModifyServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceInstanceResourcesAction) {
		query["ServiceInstanceResourcesAction"] = request.ServiceInstanceResourcesAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyServiceInstanceResources"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyServiceInstanceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pre-release service
//
// @param request - PreLaunchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreLaunchServiceResponse
func (client *Client) PreLaunchServiceWithContext(ctx context.Context, request *PreLaunchServiceRequest, runtime *dara.RuntimeOptions) (_result *PreLaunchServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreLaunchService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreLaunchServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes metering data of an Alibaba Cloud Marketplace commodity.
//
// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithContext(ctx context.Context, request *PushMeteringDataRequest, runtime *dara.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Metering) {
		query["Metering"] = request.Metering
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMeteringData"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an artifact.
//
// @param request - RegisterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterServiceResponse
func (client *Client) RegisterServiceWithContext(ctx context.Context, request *RegisterServiceRequest, runtime *dara.RuntimeOptions) (_result *RegisterServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reject service usage.
//
// @param request - RejectServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectServiceUsageResponse
func (client *Client) RejectServiceUsageWithContext(ctx context.Context, request *RejectServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *RejectServiceUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAliUid) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectServiceUsage"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectServiceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an artifact.
//
// @param request - ReleaseArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseArtifactResponse
func (client *Client) ReleaseArtifactWithContext(ctx context.Context, request *ReleaseArtifactRequest, runtime *dara.RuntimeOptions) (_result *ReleaseArtifactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseArtifact"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove  service shared account.
//
// @param request - RemoveServiceSharedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveServiceSharedAccountsResponse
func (client *Client) RemoveServiceSharedAccountsWithContext(ctx context.Context, request *RemoveServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *RemoveServiceSharedAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAliUids) {
		query["UserAliUids"] = request.UserAliUids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveServiceSharedAccounts"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveServiceSharedAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed, call the RestartServiceInstance interface to restart the service instance.
//
// @param request - RestartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstanceWithContext(ctx context.Context, request *RestartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Rollback Service Instance
//
// @param request - RollbackServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackServiceInstanceResponse
func (client *Client) RollbackServiceInstanceWithContext(ctx context.Context, request *RollbackServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RollbackServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance status is Stopped (Stopped) or StartFailed (Startup failed), the StartServiceInstance interface is invoked to start the service instance.
//
// @param request - StartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstanceWithContext(ctx context.Context, request *StartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed and StopFailed, call the StopServiceInstance interface to stop the service instance.
//
// @param request - StopServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstanceWithContext(ctx context.Context, request *StopServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag a resource
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("TagResources"),
		Version:     dara.String("2021-05-21"),
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
// # Unbind resource from tag
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithContext(ctx context.Context, request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a deployment package.
//
// @param tmpReq - UpdateArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactResponse
func (client *Client) UpdateArtifactWithContext(ctx context.Context, tmpReq *UpdateArtifactRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ArtifactBuildProperty) {
		request.ArtifactBuildPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactBuildProperty, dara.String("ArtifactBuildProperty"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ArtifactProperty) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, dara.String("ArtifactProperty"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ArtifactBuildPropertyShrink) {
		query["ArtifactBuildProperty"] = request.ArtifactBuildPropertyShrink
	}

	if !dara.IsNil(request.ArtifactId) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !dara.IsNil(request.ArtifactPropertyShrink) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PermissionType) {
		query["PermissionType"] = request.PermissionType
	}

	if !dara.IsNil(request.SupportRegionIds) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateArtifact"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a service.
//
// @param tmpReq - UpdateServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithContext(ctx context.Context, tmpReq *UpdateServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Commodity) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, dara.String("Commodity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ComplianceMetadata) {
		request.ComplianceMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComplianceMetadata, dara.String("ComplianceMetadata"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateOption) {
		request.UpdateOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateOption, dara.String("UpdateOption"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMetadata) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !dara.IsNil(request.ApprovalType) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !dara.IsNil(request.BuildParameters) {
		query["BuildParameters"] = request.BuildParameters
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommodityShrink) {
		query["Commodity"] = request.CommodityShrink
	}

	if !dara.IsNil(request.ComplianceMetadataShrink) {
		query["ComplianceMetadata"] = request.ComplianceMetadataShrink
	}

	if !dara.IsNil(request.DeployMetadata) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.IsSupportOperated) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !dara.IsNil(request.LicenseMetadata) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !dara.IsNil(request.LogMetadata) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !dara.IsNil(request.OperationMetadata) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !dara.IsNil(request.PolicyNames) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resellable) {
		query["Resellable"] = request.Resellable
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInfo) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.TenantType) {
		query["TenantType"] = request.TenantType
	}

	if !dara.IsNil(request.TrialDuration) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !dara.IsNil(request.UpdateOptionShrink) {
		query["UpdateOption"] = request.UpdateOptionShrink
	}

	if !dara.IsNil(request.UpgradeMetadata) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of a service instance.
//
// @param tmpReq - UpdateServiceInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceAttributeResponse
func (client *Client) UpdateServiceInstanceAttributeWithContext(ctx context.Context, tmpReq *UpdateServiceInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LicenseData) {
		request.LicenseDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LicenseData, dara.String("LicenseData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LicenseDataShrink) {
		query["LicenseData"] = request.LicenseDataShrink
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceInstanceAttribute"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a service instance.
//
// @param tmpReq - UpdateServiceInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpecWithContext(ctx context.Context, tmpReq *UpdateServiceInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableUserPrometheus) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !dara.IsNil(request.OperationName) {
		query["OperationName"] = request.OperationName
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.PredefinedParametersName) {
		query["PredefinedParametersName"] = request.PredefinedParametersName
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceInstanceSpec"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Service Test Case
//
// @param request - UpdateServiceTestCaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceTestCaseResponse
func (client *Client) UpdateServiceTestCaseWithContext(ctx context.Context, request *UpdateServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceTestCaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TestCaseId) {
		query["TestCaseId"] = request.TestCaseId
	}

	if !dara.IsNil(request.TestCaseName) {
		query["TestCaseName"] = request.TestCaseName
	}

	if !dara.IsNil(request.TestConfig) {
		query["TestConfig"] = request.TestConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceTestCase"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceTestCaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Service Sharing Permissions
//
// @param request - UpdateSharedAccountPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSharedAccountPermissionResponse
func (client *Client) UpdateSharedAccountPermissionWithContext(ctx context.Context, request *UpdateSharedAccountPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSharedAccountPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserAliUid) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSharedAccountPermission"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSharedAccountPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the information of supplier.
//
// @param request - UpdateSupplierInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSupplierInformationResponse
func (client *Client) UpdateSupplierInformationWithContext(ctx context.Context, request *UpdateSupplierInformationRequest, runtime *dara.RuntimeOptions) (_result *UpdateSupplierInformationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverySettings) {
		query["DeliverySettings"] = request.DeliverySettings
	}

	if !dara.IsNil(request.OperationIp) {
		query["OperationIp"] = request.OperationIp
	}

	if !dara.IsNil(request.OperationMfaPresent) {
		query["OperationMfaPresent"] = request.OperationMfaPresent
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SupplierDesc) {
		query["SupplierDesc"] = request.SupplierDesc
	}

	if !dara.IsNil(request.SupplierLogo) {
		query["SupplierLogo"] = request.SupplierLogo
	}

	if !dara.IsNil(request.SupplierUrl) {
		query["SupplierUrl"] = request.SupplierUrl
	}

	if !dara.IsNil(request.SupportContacts) {
		query["SupportContacts"] = request.SupportContacts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSupplierInformation"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSupplierInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a service instance.
//
// @param tmpReq - UpgradeServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstanceWithContext(ctx context.Context, tmpReq *UpgradeServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpgradeServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeServiceInstance"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Withdraw service version.
//
// @param request - WithdrawServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawServiceResponse
func (client *Client) WithdrawServiceWithContext(ctx context.Context, request *WithdrawServiceRequest, runtime *dara.RuntimeOptions) (_result *WithdrawServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WithdrawService"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
