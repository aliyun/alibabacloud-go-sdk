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
	client.Endpoint, _err = client.GetEndpoint(dara.String("computenestsupplier"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a shared account of a service.
//
// @param request - AddServiceSharedAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServiceSharedAccountsResponse
func (client *Client) AddServiceSharedAccountsWithOptions(request *AddServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *AddServiceSharedAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a shared account of a service.
//
// @param request - AddServiceSharedAccountsRequest
//
// @return AddServiceSharedAccountsResponse
func (client *Client) AddServiceSharedAccounts(request *AddServiceSharedAccountsRequest) (_result *AddServiceSharedAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddServiceSharedAccountsResponse{}
	_body, _err := client.AddServiceSharedAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ApproveServiceUsageWithOptions(request *ApproveServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *ApproveServiceUsageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ApproveServiceUsageResponse
func (client *Client) ApproveServiceUsage(request *ApproveServiceUsageRequest) (_result *ApproveServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApproveServiceUsageResponse{}
	_body, _err := client.ApproveServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CancelServiceRegistrationWithOptions(request *CancelServiceRegistrationRequest, runtime *dara.RuntimeOptions) (_result *CancelServiceRegistrationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CancelServiceRegistrationResponse
func (client *Client) CancelServiceRegistration(request *CancelServiceRegistrationRequest) (_result *CancelServiceRegistrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelServiceRegistrationResponse{}
	_body, _err := client.CancelServiceRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstance(request *ContinueDeployServiceInstanceRequest) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.ContinueDeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateArtifactWithOptions(tmpReq *CreateArtifactRequest, runtime *dara.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateArtifactRequest
//
// @return CreateArtifactResponse
func (client *Client) CreateArtifact(request *CreateArtifactRequest) (_result *CreateArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CreateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建运维公告
//
// @param tmpReq - CreateOpsNoticeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpsNoticeResponse
func (client *Client) CreateOpsNoticeWithOptions(tmpReq *CreateOpsNoticeRequest, runtime *dara.RuntimeOptions) (_result *CreateOpsNoticeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateOpsNoticeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributesShrink) {
		query["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
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

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Solutions) {
		query["Solutions"] = request.Solutions
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOpsNotice"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOpsNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建运维公告
//
// @param request - CreateOpsNoticeRequest
//
// @return CreateOpsNoticeResponse
func (client *Client) CreateOpsNotice(request *CreateOpsNoticeRequest) (_result *CreateOpsNoticeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOpsNoticeResponse{}
	_body, _err := client.CreateOpsNoticeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceWithOptions(tmpReq *CreateServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateServiceInstanceRequest
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (_result *CreateServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CreateServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceTestCaseWithOptions(request *CreateServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceTestCaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateServiceTestCaseResponse
func (client *Client) CreateServiceTestCase(request *CreateServiceTestCaseRequest) (_result *CreateServiceTestCaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceTestCaseResponse{}
	_body, _err := client.CreateServiceTestCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceTestTaskWithOptions(request *CreateServiceTestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceTestTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateServiceTestTaskResponse
func (client *Client) CreateServiceTestTask(request *CreateServiceTestTaskRequest) (_result *CreateServiceTestTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceTestTaskResponse{}
	_body, _err := client.CreateServiceTestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateServiceUsageWithOptions(request *CreateServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceUsageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateServiceUsageResponse
func (client *Client) CreateServiceUsage(request *CreateServiceUsageRequest) (_result *CreateServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceUsageResponse{}
	_body, _err := client.CreateServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSupplierRegistrationWithOptions(request *CreateSupplierRegistrationRequest, runtime *dara.RuntimeOptions) (_result *CreateSupplierRegistrationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSupplierRegistrationResponse
func (client *Client) CreateSupplierRegistration(request *CreateSupplierRegistrationRequest) (_result *CreateSupplierRegistrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSupplierRegistrationResponse{}
	_body, _err := client.CreateSupplierRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAcrImageRepositoriesWithOptions(request *DeleteAcrImageRepositoriesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcrImageRepositoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAcrImageRepositoriesResponse
func (client *Client) DeleteAcrImageRepositories(request *DeleteAcrImageRepositoriesRequest) (_result *DeleteAcrImageRepositoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAcrImageRepositoriesResponse{}
	_body, _err := client.DeleteAcrImageRepositoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAcrImageTagsWithOptions(request *DeleteAcrImageTagsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAcrImageTagsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAcrImageTagsResponse
func (client *Client) DeleteAcrImageTags(request *DeleteAcrImageTagsRequest) (_result *DeleteAcrImageTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAcrImageTagsResponse{}
	_body, _err := client.DeleteAcrImageTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteArtifactWithOptions(request *DeleteArtifactRequest, runtime *dara.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifact(request *DeleteArtifactRequest) (_result *DeleteArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.DeleteArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteServiceResponse
func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteServiceTestCaseWithOptions(request *DeleteServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceTestCaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteServiceTestCaseResponse
func (client *Client) DeleteServiceTestCase(request *DeleteServiceTestCaseRequest) (_result *DeleteServiceTestCaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServiceTestCaseResponse{}
	_body, _err := client.DeleteServiceTestCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeployServiceInstanceWithOptions(request *DeployServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstance(request *DeployServiceInstanceRequest) (_result *DeployServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.DeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateDefaultServiceTestConfigWithOptions(request *GenerateDefaultServiceTestConfigRequest, runtime *dara.RuntimeOptions) (_result *GenerateDefaultServiceTestConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateDefaultServiceTestConfigResponse
func (client *Client) GenerateDefaultServiceTestConfig(request *GenerateDefaultServiceTestConfigRequest) (_result *GenerateDefaultServiceTestConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateDefaultServiceTestConfigResponse{}
	_body, _err := client.GenerateDefaultServiceTestConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateServicePolicyWithOptions(request *GenerateServicePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateServicePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateServicePolicyResponse
func (client *Client) GenerateServicePolicy(request *GenerateServicePolicyRequest) (_result *GenerateServicePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateServicePolicyResponse{}
	_body, _err := client.GenerateServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactWithOptions(request *GetArtifactRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactResponse
func (client *Client) GetArtifact(request *GetArtifactRequest) (_result *GetArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactResponse{}
	_body, _err := client.GetArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetArtifactRepositoryCredentialsWithOptions(request *GetArtifactRepositoryCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetArtifactRepositoryCredentialsResponse
func (client *Client) GetArtifactRepositoryCredentials(request *GetArtifactRepositoryCredentialsRequest) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.GetArtifactRepositoryCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetNetworkAvailableZonesWithOptions(request *GetNetworkAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkAvailableZonesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetNetworkAvailableZonesResponse
func (client *Client) GetNetworkAvailableZones(request *GetNetworkAvailableZonesRequest) (_result *GetNetworkAvailableZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkAvailableZonesResponse{}
	_body, _err := client.GetNetworkAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看运维公告详情
//
// @param request - GetOpsNoticeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpsNoticeResponse
func (client *Client) GetOpsNoticeWithOptions(request *GetOpsNoticeRequest, runtime *dara.RuntimeOptions) (_result *GetOpsNoticeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NoticeId) {
		query["NoticeId"] = request.NoticeId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpsNotice"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpsNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看运维公告详情
//
// @param request - GetOpsNoticeRequest
//
// @return GetOpsNoticeResponse
func (client *Client) GetOpsNotice(request *GetOpsNoticeRequest) (_result *GetOpsNoticeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpsNoticeResponse{}
	_body, _err := client.GetOpsNoticeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceResponse
func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceEstimateCostWithOptions(tmpReq *GetServiceEstimateCostRequest, runtime *dara.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetServiceEstimateCostRequest
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCost(request *GetServiceEstimateCostRequest) (_result *GetServiceEstimateCostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.GetServiceEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceProvisionsWithOptions(tmpReq *GetServiceProvisionsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetServiceProvisionsRequest
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisions(request *GetServiceProvisionsRequest) (_result *GetServiceProvisionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.GetServiceProvisionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceRegistrationWithOptions(request *GetServiceRegistrationRequest, runtime *dara.RuntimeOptions) (_result *GetServiceRegistrationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceRegistrationResponse
func (client *Client) GetServiceRegistration(request *GetServiceRegistrationRequest) (_result *GetServiceRegistrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceRegistrationResponse{}
	_body, _err := client.GetServiceRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceTemplateCriterionIssuesWithOptions(request *GetServiceTemplateCriterionIssuesRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTemplateCriterionIssuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceTemplateCriterionIssuesResponse
func (client *Client) GetServiceTemplateCriterionIssues(request *GetServiceTemplateCriterionIssuesRequest) (_result *GetServiceTemplateCriterionIssuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceTemplateCriterionIssuesResponse{}
	_body, _err := client.GetServiceTemplateCriterionIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraints(request *GetServiceTemplateParameterConstraintsRequest) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.GetServiceTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetServiceTestTaskWithOptions(request *GetServiceTestTaskRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTestTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetServiceTestTaskResponse
func (client *Client) GetServiceTestTask(request *GetServiceTestTaskRequest) (_result *GetServiceTestTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceTestTaskResponse{}
	_body, _err := client.GetServiceTestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSupplierInformationWithOptions(request *GetSupplierInformationRequest, runtime *dara.RuntimeOptions) (_result *GetSupplierInformationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSupplierInformationResponse
func (client *Client) GetSupplierInformation(request *GetSupplierInformationRequest) (_result *GetSupplierInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSupplierInformationResponse{}
	_body, _err := client.GetSupplierInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUploadCredentialsWithOptions(request *GetUploadCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetUploadCredentialsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetUploadCredentialsResponse
func (client *Client) GetUploadCredentials(request *GetUploadCredentialsRequest) (_result *GetUploadCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadCredentialsResponse{}
	_body, _err := client.GetUploadCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) LaunchServiceWithOptions(request *LaunchServiceRequest, runtime *dara.RuntimeOptions) (_result *LaunchServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return LaunchServiceResponse
func (client *Client) LaunchService(request *LaunchServiceRequest) (_result *LaunchServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LaunchServiceResponse{}
	_body, _err := client.LaunchServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAcrImageRepositoriesWithOptions(request *ListAcrImageRepositoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAcrImageRepositoriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAcrImageRepositoriesResponse
func (client *Client) ListAcrImageRepositories(request *ListAcrImageRepositoriesRequest) (_result *ListAcrImageRepositoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAcrImageRepositoriesResponse{}
	_body, _err := client.ListAcrImageRepositoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAcrImageTagsWithOptions(request *ListAcrImageTagsRequest, runtime *dara.RuntimeOptions) (_result *ListAcrImageTagsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAcrImageTagsResponse
func (client *Client) ListAcrImageTags(request *ListAcrImageTagsRequest) (_result *ListAcrImageTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAcrImageTagsResponse{}
	_body, _err := client.ListAcrImageTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactBuildLogsWithOptions(request *ListArtifactBuildLogsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactBuildLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactBuildLogsResponse
func (client *Client) ListArtifactBuildLogs(request *ListArtifactBuildLogsRequest) (_result *ListArtifactBuildLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactBuildLogsResponse{}
	_body, _err := client.ListArtifactBuildLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactRisksWithOptions(request *ListArtifactRisksRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactRisksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactRisksResponse
func (client *Client) ListArtifactRisks(request *ListArtifactRisksRequest) (_result *ListArtifactRisksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactRisksResponse{}
	_body, _err := client.ListArtifactRisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactVersionsWithOptions(tmpReq *ListArtifactVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactVersionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListArtifactVersionsRequest
//
// @return ListArtifactVersionsResponse
func (client *Client) ListArtifactVersions(request *ListArtifactVersionsRequest) (_result *ListArtifactVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.ListArtifactVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListArtifactsWithOptions(request *ListArtifactsRequest, runtime *dara.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListArtifactsResponse
func (client *Client) ListArtifacts(request *ListArtifactsRequest) (_result *ListArtifactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListArtifactsResponse{}
	_body, _err := client.ListArtifactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询运维公告
//
// @param request - ListOpsNoticesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpsNoticesResponse
func (client *Client) ListOpsNoticesWithOptions(request *ListOpsNoticesRequest, runtime *dara.RuntimeOptions) (_result *ListOpsNoticesResponse, _err error) {
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
		Action:      dara.String("ListOpsNotices"),
		Version:     dara.String("2021-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpsNoticesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询运维公告
//
// @param request - ListOpsNoticesRequest
//
// @return ListOpsNoticesResponse
func (client *Client) ListOpsNotices(request *ListOpsNoticesRequest) (_result *ListOpsNoticesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOpsNoticesResponse{}
	_body, _err := client.ListOpsNoticesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListResellersWithOptions(request *ListResellersRequest, runtime *dara.RuntimeOptions) (_result *ListResellersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListResellersResponse
func (client *Client) ListResellers(request *ListResellersRequest) (_result *ListResellersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResellersResponse{}
	_body, _err := client.ListResellersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceBuildLogsWithOptions(request *ListServiceBuildLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceBuildLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceBuildLogsResponse
func (client *Client) ListServiceBuildLogs(request *ListServiceBuildLogsRequest) (_result *ListServiceBuildLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceBuildLogsResponse{}
	_body, _err := client.ListServiceBuildLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstanceBillWithOptions(request *ListServiceInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceBillResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBill(request *ListServiceInstanceBillRequest) (_result *ListServiceInstanceBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceBillResponse{}
	_body, _err := client.ListServiceInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstanceDeployDetailsWithOptions(request *ListServiceInstanceDeployDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceDeployDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstanceDeployDetailsResponse
func (client *Client) ListServiceInstanceDeployDetails(request *ListServiceInstanceDeployDetailsRequest) (_result *ListServiceInstanceDeployDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceDeployDetailsResponse{}
	_body, _err := client.ListServiceInstanceDeployDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogs(request *ListServiceInstanceLogsRequest) (_result *ListServiceInstanceLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.ListServiceInstanceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResources(request *ListServiceInstanceResourcesRequest) (_result *ListServiceInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.ListServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstanceUpgradeHistoryWithOptions(request *ListServiceInstanceUpgradeHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstanceUpgradeHistoryResponse
func (client *Client) ListServiceInstanceUpgradeHistory(request *ListServiceInstanceUpgradeHistoryRequest) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceUpgradeHistoryResponse{}
	_body, _err := client.ListServiceInstanceUpgradeHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceRegistrationsWithOptions(request *ListServiceRegistrationsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceRegistrationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceRegistrationsResponse
func (client *Client) ListServiceRegistrations(request *ListServiceRegistrationsRequest) (_result *ListServiceRegistrationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceRegistrationsResponse{}
	_body, _err := client.ListServiceRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceSharedAccountsWithOptions(request *ListServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceSharedAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceSharedAccountsResponse
func (client *Client) ListServiceSharedAccounts(request *ListServiceSharedAccountsRequest) (_result *ListServiceSharedAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceSharedAccountsResponse{}
	_body, _err := client.ListServiceSharedAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceTestCasesWithOptions(request *ListServiceTestCasesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestCasesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceTestCasesResponse
func (client *Client) ListServiceTestCases(request *ListServiceTestCasesRequest) (_result *ListServiceTestCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceTestCasesResponse{}
	_body, _err := client.ListServiceTestCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceTestTaskLogsWithOptions(request *ListServiceTestTaskLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestTaskLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceTestTaskLogsResponse
func (client *Client) ListServiceTestTaskLogs(request *ListServiceTestTaskLogsRequest) (_result *ListServiceTestTaskLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceTestTaskLogsResponse{}
	_body, _err := client.ListServiceTestTaskLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceTestTasksWithOptions(request *ListServiceTestTasksRequest, runtime *dara.RuntimeOptions) (_result *ListServiceTestTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceTestTasksResponse
func (client *Client) ListServiceTestTasks(request *ListServiceTestTasksRequest) (_result *ListServiceTestTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceTestTasksResponse{}
	_body, _err := client.ListServiceTestTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsages(request *ListServiceUsagesRequest) (_result *ListServiceUsagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.ListServiceUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSupplierRegistrationsWithOptions(request *ListSupplierRegistrationsRequest, runtime *dara.RuntimeOptions) (_result *ListSupplierRegistrationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSupplierRegistrationsResponse
func (client *Client) ListSupplierRegistrations(request *ListSupplierRegistrationsRequest) (_result *ListSupplierRegistrationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupplierRegistrationsResponse{}
	_body, _err := client.ListSupplierRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// # Query resource tags
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// # Query tag value list
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Modifies the resource information about a service instance.
//
// @param request - ModifyServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceInstanceResourcesResponse
func (client *Client) ModifyServiceInstanceResourcesWithOptions(request *ModifyServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyServiceInstanceResourcesResponse
func (client *Client) ModifyServiceInstanceResources(request *ModifyServiceInstanceResourcesRequest) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyServiceInstanceResourcesResponse{}
	_body, _err := client.ModifyServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PreLaunchServiceWithOptions(request *PreLaunchServiceRequest, runtime *dara.RuntimeOptions) (_result *PreLaunchServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PreLaunchServiceResponse
func (client *Client) PreLaunchService(request *PreLaunchServiceRequest) (_result *PreLaunchServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreLaunchServiceResponse{}
	_body, _err := client.PreLaunchServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *dara.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PushMeteringDataResponse
func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RegisterServiceWithOptions(request *RegisterServiceRequest, runtime *dara.RuntimeOptions) (_result *RegisterServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RegisterServiceResponse
func (client *Client) RegisterService(request *RegisterServiceRequest) (_result *RegisterServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterServiceResponse{}
	_body, _err := client.RegisterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RejectServiceUsageWithOptions(request *RejectServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *RejectServiceUsageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RejectServiceUsageResponse
func (client *Client) RejectServiceUsage(request *RejectServiceUsageRequest) (_result *RejectServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectServiceUsageResponse{}
	_body, _err := client.RejectServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReleaseArtifactWithOptions(request *ReleaseArtifactRequest, runtime *dara.RuntimeOptions) (_result *ReleaseArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReleaseArtifactResponse
func (client *Client) ReleaseArtifact(request *ReleaseArtifactRequest) (_result *ReleaseArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.ReleaseArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveServiceSharedAccountsWithOptions(request *RemoveServiceSharedAccountsRequest, runtime *dara.RuntimeOptions) (_result *RemoveServiceSharedAccountsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveServiceSharedAccountsResponse
func (client *Client) RemoveServiceSharedAccounts(request *RemoveServiceSharedAccountsRequest) (_result *RemoveServiceSharedAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveServiceSharedAccountsResponse{}
	_body, _err := client.RemoveServiceSharedAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstance(request *RestartServiceInstanceRequest) (_result *RestartServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.RestartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RollbackServiceInstanceWithOptions(request *RollbackServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RollbackServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RollbackServiceInstanceResponse
func (client *Client) RollbackServiceInstance(request *RollbackServiceInstanceRequest) (_result *RollbackServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackServiceInstanceResponse{}
	_body, _err := client.RollbackServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstance(request *StartServiceInstanceRequest) (_result *StartServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.StartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstance(request *StopServiceInstanceRequest) (_result *StopServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.StopServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// # Unbind resource from tag
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateArtifactWithOptions(tmpReq *UpdateArtifactRequest, runtime *dara.RuntimeOptions) (_result *UpdateArtifactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateArtifactRequest
//
// @return UpdateArtifactResponse
func (client *Client) UpdateArtifact(request *UpdateArtifactRequest) (_result *UpdateArtifactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.UpdateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateServiceWithOptions(tmpReq *UpdateServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
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

	if !dara.IsNil(request.ServiceLocaleConfigs) {
		query["ServiceLocaleConfigs"] = request.ServiceLocaleConfigs
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateServiceRequest
//
// @return UpdateServiceResponse
func (client *Client) UpdateService(request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateServiceInstanceAttributeWithOptions(tmpReq *UpdateServiceInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateServiceInstanceAttributeRequest
//
// @return UpdateServiceInstanceAttributeResponse
func (client *Client) UpdateServiceInstanceAttribute(request *UpdateServiceInstanceAttributeRequest) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceInstanceAttributeResponse{}
	_body, _err := client.UpdateServiceInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpdateServiceInstanceSpecRequest
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpec(request *UpdateServiceInstanceSpecRequest) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.UpdateServiceInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateServiceTestCaseWithOptions(request *UpdateServiceTestCaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceTestCaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateServiceTestCaseResponse
func (client *Client) UpdateServiceTestCase(request *UpdateServiceTestCaseRequest) (_result *UpdateServiceTestCaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceTestCaseResponse{}
	_body, _err := client.UpdateServiceTestCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateSharedAccountPermissionWithOptions(request *UpdateSharedAccountPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSharedAccountPermissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateSharedAccountPermissionResponse
func (client *Client) UpdateSharedAccountPermission(request *UpdateSharedAccountPermissionRequest) (_result *UpdateSharedAccountPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSharedAccountPermissionResponse{}
	_body, _err := client.UpdateSharedAccountPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateSupplierInformationWithOptions(request *UpdateSupplierInformationRequest, runtime *dara.RuntimeOptions) (_result *UpdateSupplierInformationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateSupplierInformationResponse
func (client *Client) UpdateSupplierInformation(request *UpdateSupplierInformationRequest) (_result *UpdateSupplierInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSupplierInformationResponse{}
	_body, _err := client.UpdateSupplierInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpgradeServiceInstanceWithOptions(tmpReq *UpgradeServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - UpgradeServiceInstanceRequest
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstance(request *UpgradeServiceInstanceRequest) (_result *UpgradeServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.UpgradeServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) WithdrawServiceWithOptions(request *WithdrawServiceRequest, runtime *dara.RuntimeOptions) (_result *WithdrawServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return WithdrawServiceResponse
func (client *Client) WithdrawService(request *WithdrawServiceRequest) (_result *WithdrawServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.WithdrawServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
