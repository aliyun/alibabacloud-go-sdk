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
	client.Endpoint, _err = client.GetEndpoint(dara.String("computenest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels a service usage request.
//
// @param request - CancelServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelServiceUsageResponse
func (client *Client) CancelServiceUsageWithOptions(request *CancelServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *CancelServiceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NeedDelete) {
		query["NeedDelete"] = request.NeedDelete
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
		Action:      dara.String("CancelServiceUsage"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a service usage request.
//
// @param request - CancelServiceUsageRequest
//
// @return CancelServiceUsageResponse
func (client *Client) CancelServiceUsage(request *CancelServiceUsageRequest) (_result *CancelServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelServiceUsageResponse{}
	_body, _err := client.CancelServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group for a cloud resource.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group for a cloud resource.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a precheck to verify that a service instance can be deployed.
//
// @param request - CheckServiceDeployableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceDeployableResponse
func (client *Client) CheckServiceDeployableWithOptions(request *CheckServiceDeployableRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceDeployableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PostPaidAmount) {
		query["PostPaidAmount"] = request.PostPaidAmount
	}

	if !dara.IsNil(request.PrePaidAmount) {
		query["PrePaidAmount"] = request.PrePaidAmount
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
		Action:      dara.String("CheckServiceDeployable"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceDeployableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a precheck to verify that a service instance can be deployed.
//
// @param request - CheckServiceDeployableRequest
//
// @return CheckServiceDeployableResponse
func (client *Client) CheckServiceDeployable(request *CheckServiceDeployableRequest) (_result *CheckServiceDeployableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckServiceDeployableResponse{}
	_body, _err := client.CheckServiceDeployableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If a service instance fails to deploy, you can call the ContinueDeployServiceInstance operation to resume the deployment.
//
// Description:
//
// If a private service instance that uses the ROS deployment type fails to deploy, you can call the ContinueDeployServiceInstance operation to resume the deployment.
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
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
		Version:     dara.String("2021-06-01"),
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
// If a service instance fails to deploy, you can call the ContinueDeployServiceInstance operation to resume the deployment.
//
// Description:
//
// If a private service instance that uses the ROS deployment type fails to deploy, you can call the ContinueDeployServiceInstance operation to resume the deployment.
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
// Creates a backup for a Compute Nest service instance.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and configured the corresponding parameters when the service was created.
//
// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
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

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup for a Compute Nest service instance.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and configured the corresponding parameters when the service was created.
//
// @param request - CreateBackupRequest
//
// @return CreateBackupResponse
func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a restore job.
//
// Description:
//
// ### Prerequisites
//
// Before you call this operation, ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when the service was created.
//
// @param request - CreateRestoreTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestoreTaskResponse
func (client *Client) CreateRestoreTaskWithOptions(request *CreateRestoreTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRestoreTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRestoreTask"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a restore job.
//
// Description:
//
// ### Prerequisites
//
// Before you call this operation, ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when the service was created.
//
// @param request - CreateRestoreTaskRequest
//
// @return CreateRestoreTaskResponse
func (client *Client) CreateRestoreTask(request *CreateRestoreTaskRequest) (_result *CreateRestoreTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CreateRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and deploys a service instance in Compute Nest by specifying parameters such as a region ID, a service instance name, and a service ID.
//
// @param tmpReq - CreateServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableInstanceOps) {
		query["EnableInstanceOps"] = request.EnableInstanceOps
	}

	if !dara.IsNil(request.EnableUserPrometheus) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OperationMetadata) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceAutoPay) {
		query["ResourceAutoPay"] = request.ResourceAutoPay
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !dara.IsNil(request.SpecificationCode) {
		query["SpecificationCode"] = request.SpecificationCode
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

	if !dara.IsNil(request.TrialType) {
		query["TrialType"] = request.TrialType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceInstance"),
		Version:     dara.String("2021-06-01"),
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
// Creates and deploys a service instance in Compute Nest by specifying parameters such as a region ID, a service instance name, and a service ID.
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
// Creates a service usage request.
//
// @param tmpReq - CreateServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceUsageResponse
func (client *Client) CreateServiceUsageWithOptions(tmpReq *CreateServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateServiceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInformation) {
		request.UserInformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInformation, dara.String("UserInformation"), dara.String("json"))
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

	if !dara.IsNil(request.UserInformationShrink) {
		query["UserInformation"] = request.UserInformationShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceUsage"),
		Version:     dara.String("2021-06-01"),
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
// Creates a service usage request.
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
// Creates a Skill.
//
// Description:
//
// ### Prerequisites
//
// Before you begin, ensure that the service provider has enabled the modification feature and configured its parameters during service creation.
//
// @param request - CreateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillResponse
func (client *Client) CreateSkillWithOptions(request *CreateSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.SkillDescription) {
		query["SkillDescription"] = request.SkillDescription
	}

	if !dara.IsNil(request.SkillLabels) {
		query["SkillLabels"] = request.SkillLabels
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillSpaceId) {
		query["SkillSpaceId"] = request.SkillSpaceId
	}

	if !dara.IsNil(request.SourceSkillId) {
		query["SourceSkillId"] = request.SourceSkillId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkill"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Skill.
//
// Description:
//
// ### Prerequisites
//
// Before you begin, ensure that the service provider has enabled the modification feature and configured its parameters during service creation.
//
// @param request - CreateSkillRequest
//
// @return CreateSkillResponse
func (client *Client) CreateSkill(request *CreateSkillRequest) (_result *CreateSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillResponse{}
	_body, _err := client.CreateSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Skill file detection task.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modification feature and configured its parameters during service creation.
//
// @param request - CreateSkillFileDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillFileDetectResponse
func (client *Client) CreateSkillFileDetectWithOptions(request *CreateSkillFileDetectRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillFileDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillFileDetect"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillFileDetectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Skill file detection task.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modification feature and configured its parameters during service creation.
//
// @param request - CreateSkillFileDetectRequest
//
// @return CreateSkillFileDetectResponse
func (client *Client) CreateSkillFileDetect(request *CreateSkillFileDetectRequest) (_result *CreateSkillFileDetectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillFileDetectResponse{}
	_body, _err := client.CreateSkillFileDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a SkillHub config.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modify configuration feature and set the required parameters.
//
// @param request - CreateSkillHubConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillHubConfigResponse
func (client *Client) CreateSkillHubConfigWithOptions(request *CreateSkillHubConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillHubConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssRegionId) {
		query["OssRegionId"] = request.OssRegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillHubConfig"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillHubConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a SkillHub config.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modify configuration feature and set the required parameters.
//
// @param request - CreateSkillHubConfigRequest
//
// @return CreateSkillHubConfigResponse
func (client *Client) CreateSkillHubConfig(request *CreateSkillHubConfigRequest) (_result *CreateSkillHubConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillHubConfigResponse{}
	_body, _err := client.CreateSkillHubConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a skill space.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the modify configuration feature and set the required parameters during service creation.
//
// @param request - CreateSkillSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillSpaceResponse
func (client *Client) CreateSkillSpaceWithOptions(request *CreateSkillSpaceRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillSpaceDescription) {
		query["SkillSpaceDescription"] = request.SkillSpaceDescription
	}

	if !dara.IsNil(request.SkillSpaceName) {
		query["SkillSpaceName"] = request.SkillSpaceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillSpace"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a skill space.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the modify configuration feature and set the required parameters during service creation.
//
// @param request - CreateSkillSpaceRequest
//
// @return CreateSkillSpaceResponse
func (client *Client) CreateSkillSpace(request *CreateSkillSpaceRequest) (_result *CreateSkillSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillSpaceResponse{}
	_body, _err := client.CreateSkillSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Compute Nest instance backup.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and configured the upgrade/downgrade parameters when the service was created.
//
// @param request - DeleteBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Compute Nest instance backup.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and configured the upgrade/downgrade parameters when the service was created.
//
// @param request - DeleteBackupRequest
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified service instances in a region.
//
// @param request - DeleteServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Deletes the specified service instances in a region.
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
// Deletes a skill.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the modification feature and configured the modification parameters when creating the service.
//
// @param request - DeleteSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkillWithOptions(request *DeleteSkillRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a skill.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the modification feature and configured the modification parameters when creating the service.
//
// @param request - DeleteSkillRequest
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkill(request *DeleteSkillRequest) (_result *DeleteSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillResponse{}
	_body, _err := client.DeleteSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a SkillSpace.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the modify configuration feature and set the required parameters during service creation.
//
// @param request - DeleteSkillSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillSpaceResponse
func (client *Client) DeleteSkillSpaceWithOptions(request *DeleteSkillSpaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillSpaceId) {
		query["SkillSpaceId"] = request.SkillSpaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkillSpace"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a SkillSpace.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the modify configuration feature and set the required parameters during service creation.
//
// @param request - DeleteSkillSpaceRequest
//
// @return DeleteSkillSpaceResponse
func (client *Client) DeleteSkillSpace(request *DeleteSkillSpaceRequest) (_result *DeleteSkillSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillSpaceResponse{}
	_body, _err := client.DeleteSkillSpaceWithOptions(request, runtime)
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Queries the available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check for missing access policies before you create a service instance.
//
// @param tmpReq - GenerateServicePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateServicePolicyResponse
func (client *Client) GenerateServicePolicyWithOptions(tmpReq *GenerateServicePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateServicePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateServicePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationTypes) {
		query["OperationTypes"] = request.OperationTypes
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
		Version:     dara.String("2021-06-01"),
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
// Check for missing access policies before you create a service instance.
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
// Retrieves the details of a backup.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when the service was created.
//
// @param request - GetBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupResponse
func (client *Client) GetBackupWithOptions(request *GetBackupRequest, runtime *dara.RuntimeOptions) (_result *GetBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBackup"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a backup.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when the service was created.
//
// @param request - GetBackupRequest
//
// @return GetBackupResponse
func (client *Client) GetBackup(request *GetBackupRequest) (_result *GetBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBackupResponse{}
	_body, _err := client.GetBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns a list of available zones for a specified region.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and set the related parameters when the service was created.
//
// @param request - GetNetworkAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkAvailableZonesResponse
func (client *Client) GetNetworkAvailableZonesWithOptions(request *GetNetworkAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkAvailableZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Returns a list of available zones for a specified region.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade/downgrade feature and set the related parameters when the service was created.
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
// Retrieves the details of a specified service.
//
// @param request - GetServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.ShowDetails) {
		query["ShowDetails"] = request.ShowDetails
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2021-06-01"),
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
// Retrieves the details of a specified service.
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
// Queries the estimated cost of creating a service instance.
//
// @param tmpReq - GetServiceEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCostWithOptions(tmpReq *GetServiceEstimateCostRequest, runtime *dara.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.OperationName) {
		query["OperationName"] = request.OperationName
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

	if !dara.IsNil(request.TrialType) {
		query["TrialType"] = request.TrialType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceEstimateCost"),
		Version:     dara.String("2021-06-01"),
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
// Queries the estimated cost of creating a service instance.
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
// Queries the details of a service instance by its region ID, service instance ID, or Alibaba Cloud Marketplace instance ID. The returned information includes the service status, template name, and a list of resources.
//
// @param request - GetServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarketInstanceId) {
		query["MarketInstanceId"] = request.MarketInstanceId
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
		Action:      dara.String("GetServiceInstance"),
		Version:     dara.String("2021-06-01"),
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
// Queries the details of a service instance by its region ID, service instance ID, or Alibaba Cloud Marketplace instance ID. The returned information includes the service status, template name, and a list of resources.
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
// Queries the estimated renewal cost of subscription resources in a privately deployed service instance. You can query the cost for all resources associated with a service instance ID, or for a specific set of resources.
//
// @param request - GetServiceInstanceSubscriptionEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceSubscriptionEstimateCostResponse
func (client *Client) GetServiceInstanceSubscriptionEstimateCostWithOptions(request *GetServiceInstanceSubscriptionEstimateCostRequest, runtime *dara.RuntimeOptions) (_result *GetServiceInstanceSubscriptionEstimateCostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourcePeriod) {
		query["ResourcePeriod"] = request.ResourcePeriod
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceInstanceSubscriptionEstimateCost"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceInstanceSubscriptionEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated renewal cost of subscription resources in a privately deployed service instance. You can query the cost for all resources associated with a service instance ID, or for a specific set of resources.
//
// @param request - GetServiceInstanceSubscriptionEstimateCostRequest
//
// @return GetServiceInstanceSubscriptionEstimateCostResponse
func (client *Client) GetServiceInstanceSubscriptionEstimateCost(request *GetServiceInstanceSubscriptionEstimateCostRequest) (_result *GetServiceInstanceSubscriptionEstimateCostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceInstanceSubscriptionEstimateCostResponse{}
	_body, _err := client.GetServiceInstanceSubscriptionEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the provisions of a service, such as its activation status and service roles.
//
// @param tmpReq - GetServiceProvisionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisionsWithOptions(tmpReq *GetServiceProvisionsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.TrialType) {
		query["TrialType"] = request.TrialType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceProvisions"),
		Version:     dara.String("2021-06-01"),
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
// Retrieves the provisions of a service, such as its activation status and service roles.
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
// Queries the parameter constraints for a Resource Orchestration Service (ROS) template in a specified region using the template name and service ID.
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.SpecificationName) {
		query["SpecificationName"] = request.SpecificationName
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
		Action:      dara.String("GetServiceTemplateParameterConstraints"),
		Version:     dara.String("2021-06-01"),
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
// Queries the parameter constraints for a Resource Orchestration Service (ROS) template in a specified region using the template name and service ID.
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
// Queries the details of a Skill.
//
// Description:
//
// ### Prerequisites
//
// Before calling this operation, ensure that the service provider enabled the modification feature and configured the required parameters during service creation.
//
// @param request - GetSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillResponse
func (client *Client) GetSkillWithOptions(request *GetSkillRequest, runtime *dara.RuntimeOptions) (_result *GetSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkill"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Skill.
//
// Description:
//
// ### Prerequisites
//
// Before calling this operation, ensure that the service provider enabled the modification feature and configured the required parameters during service creation.
//
// @param request - GetSkillRequest
//
// @return GetSkillResponse
func (client *Client) GetSkill(request *GetSkillRequest) (_result *GetSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillResponse{}
	_body, _err := client.GetSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the result of a Skill file detection task.
//
// Description:
//
// ### Prerequisites
//
// This operation is available only if the service provider enabled the modification feature and configured the relevant parameters during service creation.
//
// @param request - GetSkillFileDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillFileDetectResultResponse
func (client *Client) GetSkillFileDetectResultWithOptions(request *GetSkillFileDetectResultRequest, runtime *dara.RuntimeOptions) (_result *GetSkillFileDetectResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HashKey) {
		query["HashKey"] = request.HashKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillFileDetectResult"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillFileDetectResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a Skill file detection task.
//
// Description:
//
// ### Prerequisites
//
// This operation is available only if the service provider enabled the modification feature and configured the relevant parameters during service creation.
//
// @param request - GetSkillFileDetectResultRequest
//
// @return GetSkillFileDetectResultResponse
func (client *Client) GetSkillFileDetectResult(request *GetSkillFileDetectResultRequest) (_result *GetSkillFileDetectResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillFileDetectResultResponse{}
	_body, _err := client.GetSkillFileDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the SkillHub configuration.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the modification feature and configured its parameters when creating the service.
//
// @param request - GetSkillHubConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillHubConfigResponse
func (client *Client) GetSkillHubConfigWithOptions(request *GetSkillHubConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSkillHubConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillHubConfig"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillHubConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the SkillHub configuration.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the modification feature and configured its parameters when creating the service.
//
// @param request - GetSkillHubConfigRequest
//
// @return GetSkillHubConfigResponse
func (client *Client) GetSkillHubConfig(request *GetSkillHubConfigRequest) (_result *GetSkillHubConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillHubConfigResponse{}
	_body, _err := client.GetSkillHubConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a SkillSpace.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modification feature and configured the modification parameters for the service.
//
// @param request - GetSkillSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillSpaceResponse
func (client *Client) GetSkillSpaceWithOptions(request *GetSkillSpaceRequest, runtime *dara.RuntimeOptions) (_result *GetSkillSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillSpaceId) {
		query["SkillSpaceId"] = request.SkillSpaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillSpace"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a SkillSpace.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the modification feature and configured the modification parameters for the service.
//
// @param request - GetSkillSpaceRequest
//
// @return GetSkillSpaceResponse
func (client *Client) GetSkillSpace(request *GetSkillSpaceRequest) (_result *GetSkillSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillSpaceResponse{}
	_body, _err := client.GetSkillSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user information.
//
// @param request - GetUserInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInformationResponse
func (client *Client) GetUserInformationWithOptions(request *GetUserInformationRequest, runtime *dara.RuntimeOptions) (_result *GetUserInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInformation"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user information.
//
// @param request - GetUserInformationRequest
//
// @return GetUserInformationResponse
func (client *Client) GetUserInformation(request *GetUserInformationRequest) (_result *GetUserInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserInformationResponse{}
	_body, _err := client.GetUserInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of backups for Compute Nest service instances.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured the related parameters when the service was created.
//
// @param request - ListBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupsResponse
func (client *Client) ListBackupsWithOptions(request *ListBackupsRequest, runtime *dara.RuntimeOptions) (_result *ListBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBackups"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of backups for Compute Nest service instances.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured the related parameters when the service was created.
//
// @param request - ListBackupsRequest
//
// @return ListBackupsResponse
func (client *Client) ListBackups(request *ListBackupsRequest) (_result *ListBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBackupsResponse{}
	_body, _err := client.ListBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of access policies.
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
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2021-06-01"),
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
// Retrieves a list of access policies.
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
// Lists public skills.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the change specification feature and configured the relevant parameters during service creation.
//
// @param request - ListPublicSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicSkillsResponse
func (client *Client) ListPublicSkillsWithOptions(request *ListPublicSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListPublicSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedDownloadUrl) {
		query["NeedDownloadUrl"] = request.NeedDownloadUrl
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublicSkills"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublicSkillsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists public skills.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the change specification feature and configured the relevant parameters during service creation.
//
// @param request - ListPublicSkillsRequest
//
// @return ListPublicSkillsResponse
func (client *Client) ListPublicSkills(request *ListPublicSkillsRequest) (_result *ListPublicSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPublicSkillsResponse{}
	_body, _err := client.ListPublicSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of restore jobs.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when creating the service.
//
// @param request - ListRestoreTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRestoreTasksResponse
func (client *Client) ListRestoreTasksWithOptions(request *ListRestoreTasksRequest, runtime *dara.RuntimeOptions) (_result *ListRestoreTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRestoreTasks"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRestoreTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of restore jobs.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the upgrade and downgrade feature and configured its parameters when creating the service.
//
// @param request - ListRestoreTasksRequest
//
// @return ListRestoreTasksResponse
func (client *Client) ListRestoreTasks(request *ListRestoreTasksRequest) (_result *ListRestoreTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRestoreTasksResponse{}
	_body, _err := client.ListRestoreTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the available service categories.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceCategoriesResponse
func (client *Client) ListServiceCategoriesWithOptions(runtime *dara.RuntimeOptions) (_result *ListServiceCategoriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceCategories"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the available service categories.
//
// @return ListServiceCategoriesResponse
func (client *Client) ListServiceCategories() (_result *ListServiceCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceCategoriesResponse{}
	_body, _err := client.ListServiceCategoriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListServiceInstanceBill operation to query bills for service instances.
//
// @param request - ListServiceInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBillWithOptions(request *ListServiceInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceBill"),
		Version:     dara.String("2021-06-01"),
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
// Call the ListServiceInstanceBill operation to query bills for service instances.
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
// Retrieves the deployment and upgrade logs of a service instance.
//
// @param request - ListServiceInstanceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Retrieves the deployment and upgrade logs of a service instance.
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
// Queries the resources of a service instance based on parameters such as the service instance ID, Alibaba Cloud Resource Name (ARN), and region ID.
//
// @param request - ListServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Queries the resources of a service instance based on parameters such as the service instance ID, Alibaba Cloud Resource Name (ARN), and region ID.
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
// You can call ListServiceInstanceUpgradeHistory to query the upgrade history of a service instance.
//
// @param request - ListServiceInstanceUpgradeHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceUpgradeHistoryResponse
func (client *Client) ListServiceInstanceUpgradeHistoryWithOptions(request *ListServiceInstanceUpgradeHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
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
		Version:     dara.String("2021-06-01"),
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
// You can call ListServiceInstanceUpgradeHistory to query the upgrade history of a service instance.
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
// Queries service instances for your account in a specified region.
//
// @param request - ListServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstances"),
		Version:     dara.String("2021-06-01"),
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
// Queries service instances for your account in a specified region.
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
// Queries your service usages.
//
// @param request - ListServiceUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *dara.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListServiceUsages"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries your service usages.
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
// Queries information about your services.
//
// @param request - ListServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.FuzzyKeyword) {
		query["FuzzyKeyword"] = request.FuzzyKeyword
	}

	if !dara.IsNil(request.InUsed) {
		query["InUsed"] = request.InUsed
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderByType) {
		query["OrderByType"] = request.OrderByType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceAccessType) {
		query["ServiceAccessType"] = request.ServiceAccessType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2021-06-01"),
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
// Queries information about your services.
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
// Lists the files of a specified Skill.
//
// Description:
//
// ### Prerequisites
//
// Before you call this operation, ensure that the service provider has enabled the modification feature and configured the modification parameters when creating the service.
//
// @param request - ListSkillFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillFilesResponse
func (client *Client) ListSkillFilesWithOptions(request *ListSkillFilesRequest, runtime *dara.RuntimeOptions) (_result *ListSkillFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillFiles"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the files of a specified Skill.
//
// Description:
//
// ### Prerequisites
//
// Before you call this operation, ensure that the service provider has enabled the modification feature and configured the modification parameters when creating the service.
//
// @param request - ListSkillFilesRequest
//
// @return ListSkillFilesResponse
func (client *Client) ListSkillFiles(request *ListSkillFilesRequest) (_result *ListSkillFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillFilesResponse{}
	_body, _err := client.ListSkillFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of SkillSpaces.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the upgrade feature and configured its parameters when creating the service.
//
// @param request - ListSkillSpacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillSpacesResponse
func (client *Client) ListSkillSpacesWithOptions(request *ListSkillSpacesRequest, runtime *dara.RuntimeOptions) (_result *ListSkillSpacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillSpaces"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillSpacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of SkillSpaces.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider enabled the upgrade feature and configured its parameters when creating the service.
//
// @param request - ListSkillSpacesRequest
//
// @return ListSkillSpacesResponse
func (client *Client) ListSkillSpaces(request *ListSkillSpacesRequest) (_result *ListSkillSpacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillSpacesResponse{}
	_body, _err := client.ListSkillSpacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of Skills.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled modification and configured the modification parameters when creating the service.
//
// @param request - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(request *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NeedDownloadUrl) {
		query["NeedDownloadUrl"] = request.NeedDownloadUrl
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of Skills.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled modification and configured the modification parameters when creating the service.
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation lists existing tag keys.
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
		Version:     dara.String("2021-06-01"),
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
// This operation lists existing tag keys.
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
// Queries a list of tagged resources.
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
		Version:     dara.String("2021-06-01"),
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
// Queries a list of tagged resources.
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
// Lists the tag values for a specified tag key.
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
		Version:     dara.String("2021-06-01"),
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
// Lists the tag values for a specified tag key.
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
// Renews a service instance.
//
// @param request - RenewServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewServiceInstanceResponse
func (client *Client) RenewServiceInstanceWithOptions(request *RenewServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewServiceInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewServiceInstance"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a service instance.
//
// @param request - RenewServiceInstanceRequest
//
// @return RenewServiceInstanceResponse
func (client *Client) RenewServiceInstance(request *RenewServiceInstanceRequest) (_result *RenewServiceInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewServiceInstanceResponse{}
	_body, _err := client.RenewServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews the subscription resources for a private deployment service instance. This operation supports two renewal methods: renewing all subscription resources for a service instance or renewing only specific resources. You can use only one method per request.
//
// @param request - RenewServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewServiceInstanceResourcesResponse
func (client *Client) RenewServiceInstanceResourcesWithOptions(request *RenewServiceInstanceResourcesRequest, runtime *dara.RuntimeOptions) (_result *RenewServiceInstanceResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourcePeriod) {
		query["ResourcePeriod"] = request.ResourcePeriod
	}

	if !dara.IsNil(request.ServiceInstanceId) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewServiceInstanceResources"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews the subscription resources for a private deployment service instance. This operation supports two renewal methods: renewing all subscription resources for a service instance or renewing only specific resources. You can use only one method per request.
//
// @param request - RenewServiceInstanceResourcesRequest
//
// @return RenewServiceInstanceResourcesResponse
func (client *Client) RenewServiceInstanceResources(request *RenewServiceInstanceResourcesRequest) (_result *RenewServiceInstanceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewServiceInstanceResourcesResponse{}
	_body, _err := client.RenewServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a service instance in the Deployed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the restart O\\&M operation when the service was created.
//
// @param request - RestartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Restarts a service instance in the Deployed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider enabled the restart O\\&M operation when the service was created.
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
// Rolls back an upgraded service instance.
//
// @param request - RollbackServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackServiceInstanceResponse
func (client *Client) RollbackServiceInstanceWithOptions(request *RollbackServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *RollbackServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Rolls back an upgraded service instance.
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
// Call the StartServiceInstance operation to start a service instance that is in the Stopped or StartFailed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider configured the start and stop O\\&M actions when creating the service.
//
// @param request - StartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Call the StartServiceInstance operation to start a service instance that is in the Stopped or StartFailed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider configured the start and stop O\\&M actions when creating the service.
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
// Call the StopServiceInstance operation to shut down a service instance that is in the Deployed or StopFailed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider configured the Operations and Maintenance (O\\&M) operations for startup and shutdown when the service was created.
//
// @param request - StopServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Call the StopServiceInstance operation to shut down a service instance that is in the Deployed or StopFailed state.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider configured the Operations and Maintenance (O\\&M) operations for startup and shutdown when the service was created.
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
// You can call TagResources to create and attach tags to multiple resources.
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
		Version:     dara.String("2021-06-01"),
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
// You can call TagResources to create and attach tags to multiple resources.
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
// Detaches tags from specified resources.
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
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
		Version:     dara.String("2021-06-01"),
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
// Detaches tags from specified resources.
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
// Updates the attributes of a service instance.
//
// @param request - UpdateServiceInstanceAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceAttributesResponse
func (client *Client) UpdateServiceInstanceAttributesWithOptions(request *UpdateServiceInstanceAttributesRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableOperation) {
		query["EnableOperation"] = request.EnableOperation
	}

	if !dara.IsNil(request.GrantedPermission) {
		query["GrantedPermission"] = request.GrantedPermission
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
		Action:      dara.String("UpdateServiceInstanceAttributes"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceInstanceAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a service instance.
//
// @param request - UpdateServiceInstanceAttributesRequest
//
// @return UpdateServiceInstanceAttributesResponse
func (client *Client) UpdateServiceInstanceAttributes(request *UpdateServiceInstanceAttributesRequest) (_result *UpdateServiceInstanceAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceInstanceAttributesResponse{}
	_body, _err := client.UpdateServiceInstanceAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateServiceInstanceSpec operation to upgrade or downgrade a service instance if its parameters or package settings do not meet your requirements.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the upgrade/downgrade feature and configured the related parameters when the service was created.
//
// @param tmpReq - UpdateServiceInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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
		Version:     dara.String("2021-06-01"),
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
// You can call the UpdateServiceInstanceSpec operation to upgrade or downgrade a service instance if its parameters or package settings do not meet your requirements.
//
// Description:
//
// ### Prerequisites
//
// Ensure that the service provider has enabled the upgrade/downgrade feature and configured the related parameters when the service was created.
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
// Updates the information about a service usage request.
//
// @param tmpReq - UpdateServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceUsageResponse
func (client *Client) UpdateServiceUsageWithOptions(tmpReq *UpdateServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateServiceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInformation) {
		request.UserInformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInformation, dara.String("UserInformation"), dara.String("json"))
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

	if !dara.IsNil(request.UserInformationShrink) {
		query["UserInformation"] = request.UserInformationShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceUsage"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a service usage request.
//
// @param request - UpdateServiceUsageRequest
//
// @return UpdateServiceUsageResponse
func (client *Client) UpdateServiceUsage(request *UpdateServiceUsageRequest) (_result *UpdateServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceUsageResponse{}
	_body, _err := client.UpdateServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a skill.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the option to modify configurations and configured the relevant parameters during service creation.
//
// @param request - UpdateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkillWithOptions(request *UpdateSkillRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
	}

	if !dara.IsNil(request.SkillDescription) {
		query["SkillDescription"] = request.SkillDescription
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	if !dara.IsNil(request.SkillLabels) {
		query["SkillLabels"] = request.SkillLabels
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SourceSkillId) {
		query["SourceSkillId"] = request.SourceSkillId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkill"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a skill.
//
// Description:
//
// ### Prerequisites
//
// Ensure the service provider has enabled the option to modify configurations and configured the relevant parameters during service creation.
//
// @param request - UpdateSkillRequest
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkill(request *UpdateSkillRequest) (_result *UpdateSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillResponse{}
	_body, _err := client.UpdateSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The UpdateUserInformation operation updates user information.
//
// @param request - UpdateUserInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserInformationResponse
func (client *Client) UpdateUserInformationWithOptions(request *UpdateUserInformationRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverySettings) {
		query["DeliverySettings"] = request.DeliverySettings
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserInformation"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The UpdateUserInformation operation updates user information.
//
// @param request - UpdateUserInformationRequest
//
// @return UpdateUserInformationResponse
func (client *Client) UpdateUserInformation(request *UpdateUserInformationRequest) (_result *UpdateUserInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserInformationResponse{}
	_body, _err := client.UpdateUserInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades a service instance to a new version.
//
// @param tmpReq - UpgradeServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstanceWithOptions(tmpReq *UpgradeServiceInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2021-06-01"),
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
// Upgrades a service instance to a new version.
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
// Validates the name of a service instance.
//
// @param request - ValidateServiceInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateServiceInstanceNameResponse
func (client *Client) ValidateServiceInstanceNameWithOptions(request *ValidateServiceInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *ValidateServiceInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IsTrial) {
		query["IsTrial"] = request.IsTrial
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceInstanceName) {
		query["ServiceInstanceName"] = request.ServiceInstanceName
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
		Action:      dara.String("ValidateServiceInstanceName"),
		Version:     dara.String("2021-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateServiceInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the name of a service instance.
//
// @param request - ValidateServiceInstanceNameRequest
//
// @return ValidateServiceInstanceNameResponse
func (client *Client) ValidateServiceInstanceName(request *ValidateServiceInstanceNameRequest) (_result *ValidateServiceInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateServiceInstanceNameResponse{}
	_body, _err := client.ValidateServiceInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
