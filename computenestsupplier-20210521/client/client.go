// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddServiceSharedAccountsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-63b8a060e9d54cxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The shared account and permissions of the service.
	//
	// This parameter is required.
	SharedAccounts []*AddServiceSharedAccountsRequestSharedAccounts `json:"SharedAccounts,omitempty" xml:"SharedAccounts,omitempty" type:"Repeated"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSharedAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequest) SetClientToken(v string) *AddServiceSharedAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetRegionId(v string) *AddServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetServiceId(v string) *AddServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetSharedAccounts(v []*AddServiceSharedAccountsRequestSharedAccounts) *AddServiceSharedAccountsRequest {
	s.SharedAccounts = v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetType(v string) *AddServiceSharedAccountsRequest {
	s.Type = &v
	return s
}

type AddServiceSharedAccountsRequestSharedAccounts struct {
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserAliUid *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s AddServiceSharedAccountsRequestSharedAccounts) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsRequestSharedAccounts) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetPermission(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.Permission = &v
	return s
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetUserAliUid(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.UserAliUid = &v
	return s
}

type AddServiceSharedAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E2815213-EA4F-5759-8EA1-56DD051BB3FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServiceSharedAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponseBody) SetRequestId(v string) *AddServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

type AddServiceSharedAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddServiceSharedAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *AddServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetStatusCode(v int32) *AddServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetBody(v *AddServiceSharedAccountsResponseBody) *AddServiceSharedAccountsResponse {
	s.Body = v
	return s
}

type ApproveServiceUsageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Approval comments.
	//
	// example:
	//
	// Welcome to TuGraph. Your application has been approved. Should you have any questions, please feel free to email us at tugraph@service.alipay.com or call our contact number 400-903-0809. Contact person: Yuansu.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// ServiceSharedAccountType，
	//
	// example:
	//
	// SharedAccount
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// User ali uid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s ApproveServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageRequest) SetClientToken(v string) *ApproveServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetComments(v string) *ApproveServiceUsageRequest {
	s.Comments = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetRegionId(v string) *ApproveServiceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetServiceId(v string) *ApproveServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetType(v int32) *ApproveServiceUsageRequest {
	s.Type = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetUserAliUid(v int64) *ApproveServiceUsageRequest {
	s.UserAliUid = &v
	return s
}

type ApproveServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageResponseBody) SetRequestId(v string) *ApproveServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type ApproveServiceUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageResponse) SetHeaders(v map[string]*string) *ApproveServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *ApproveServiceUsageResponse) SetStatusCode(v int32) *ApproveServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveServiceUsageResponse) SetBody(v *ApproveServiceUsageResponseBody) *ApproveServiceUsageResponse {
	s.Body = v
	return s
}

type CancelServiceRegistrationRequest struct {
	// Client token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service registration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sr-540930183f93xxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
}

func (s CancelServiceRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationRequest) SetClientToken(v string) *CancelServiceRegistrationRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetRegionId(v string) *CancelServiceRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetRegistrationId(v string) *CancelServiceRegistrationRequest {
	s.RegistrationId = &v
	return s
}

type CancelServiceRegistrationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// C4A145D8-6F6C-532A-9001-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelServiceRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponseBody) SetRequestId(v string) *CancelServiceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type CancelServiceRegistrationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelServiceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelServiceRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponse) SetHeaders(v map[string]*string) *CancelServiceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CancelServiceRegistrationResponse) SetStatusCode(v int32) *CancelServiceRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelServiceRegistrationResponse) SetBody(v *CancelServiceRegistrationResponseBody) *CancelServiceRegistrationResponse {
	s.Body = v
	return s
}

type ContinueDeployServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetDryRun(v bool) *ContinueDeployServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBody struct {
	// The dry run result.
	DryRunResult *ContinueDeployServiceInstanceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 82DF27ED-E538-5AC0-A11C-39334A873189
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBodyDryRunResult struct {
	// The parameters that can be modified. The operation that is performed to modify the parameters does not cause a validation error.
	//
	// >  This parameter is returned only if DryRun is set to true.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions. The new values of the parameters determine whether the operation that is performed to modify the parameters causes a validation error.
	//
	// >  This parameter is returned only if DryRun is set to true.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified. The operation that is performed to modify the parameters causes a validation error.
	//
	// >  This parameter is returned only if DryRun is set to true.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

type ContinueDeployServiceInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponse) SetHeaders(v map[string]*string) *ContinueDeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetStatusCode(v int32) *ContinueDeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateArtifactRequest struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildProperty *CreateArtifactRequestArtifactBuildProperty `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty" type:"Struct"`
	// The type of the artifact build task. Valid values:
	//
	// - EcsImage: Build ECS (Elastic Container Service) image.
	//
	// - Dockerfile: Build container image based on Dockerfile.
	//
	// - Buildpacks: Build container image based on Buildpacks.
	//
	// - ContainerImage: Rebuild container image by renaming an existing container image.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment object.
	ArtifactProperty *CreateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	// The type of the deployment package. Valid values:
	//
	// 	- EcsImage: Elastic Compute Service (ECS) image.
	//
	// 	- AcrImage: container image.
	//
	// 	- File: Object Storage Service (OSS) object.
	//
	// 	- Script: script.
	//
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Test artifact
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The supported regions.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The custom tags.
	Tag []*CreateArtifactRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequest) SetArtifactBuildProperty(v *CreateArtifactRequestArtifactBuildProperty) *CreateArtifactRequest {
	s.ArtifactBuildProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactBuildType(v string) *CreateArtifactRequest {
	s.ArtifactBuildType = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactId(v string) *CreateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactProperty(v *CreateArtifactRequestArtifactProperty) *CreateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactType(v string) *CreateArtifactRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactRequest) SetClientToken(v string) *CreateArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateArtifactRequest) SetDescription(v string) *CreateArtifactRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactRequest) SetName(v string) *CreateArtifactRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactRequest) SetResourceGroupId(v string) *CreateArtifactRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactRequest) SetSupportRegionIds(v []*string) *CreateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactRequest) SetTag(v []*CreateArtifactRequestTag) *CreateArtifactRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactRequest) SetVersionName(v string) *CreateArtifactRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactRequestArtifactBuildProperty struct {
	// The build arguments used during the image build process.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	BuildArgs []*CreateArtifactRequestArtifactBuildPropertyBuildArgs `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The address of the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile or Buildpacks type.
	CodeRepo *CreateArtifactRequestArtifactBuildPropertyCodeRepo `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty" type:"Struct"`
	// The command content.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// echo "start run command"
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances.
	//
	// 	- RunShellScript: shell command, applicable to Linux instances.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The relative path to the Dockerfile within the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	//
	// example:
	//
	// ./file/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// The region ID where the source mirror image is located.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The pull location of the source container image. This is used for the command docker pull ${SourceContainerImage}.
	//
	// >  This parameter is available only if the ArtifactBuildType is ContainerImage type.
	//
	// example:
	//
	// pytorch/pytorch:2.5.1-cuda12.4-cudnn9-devel
	SourceContainerImage *string `json:"SourceContainerImage,omitempty" xml:"SourceContainerImage,omitempty"`
	// The source image id. Supported Types:
	//
	// - Image ID: Pass the Image ID of the Ecs image directly.
	//
	// - OOS Common Parameter Name: Obtain the corresponding Image ID automatically by using the OOS common parameter name.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// Image ID：m-t4nhenrdc38pe4*****
	//
	// ubuntu_22_04_x64_20G_alibase_20240926.vhd
	//
	// OOS Common Parameter Name：aliyun/services/computenest/images/aliyun_3_2104_python_3_11
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildProperty) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetBuildArgs(v []*CreateArtifactRequestArtifactBuildPropertyBuildArgs) *CreateArtifactRequestArtifactBuildProperty {
	s.BuildArgs = v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCodeRepo(v *CreateArtifactRequestArtifactBuildPropertyCodeRepo) *CreateArtifactRequestArtifactBuildProperty {
	s.CodeRepo = v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCommandContent(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.CommandContent = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCommandType(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.CommandType = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetDockerfilePath(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.DockerfilePath = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetRegionId(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetSourceContainerImage(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.SourceContainerImage = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetSourceImageId(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.SourceImageId = &v
	return s
}

type CreateArtifactRequestArtifactBuildPropertyBuildArgs struct {
	// The name of a specific build argument.
	//
	// example:
	//
	// ENV
	ArgumentName *string `json:"ArgumentName,omitempty" xml:"ArgumentName,omitempty"`
	// The value of a specific build argument.
	//
	// example:
	//
	// nginx:latest
	ArgumentValue *string `json:"ArgumentValue,omitempty" xml:"ArgumentValue,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildPropertyBuildArgs) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildPropertyBuildArgs) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentName(v string) *CreateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentName = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentValue(v string) *CreateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentValue = &v
	return s
}

type CreateArtifactRequestArtifactBuildPropertyCodeRepo struct {
	// The name of the branch in the code repository.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The endpoint.
	//
	// The URL address used to access the privately deployed GitLab instance.
	//
	// example:
	//
	// http://121.40.25.0
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// 455231
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The owner of the code repository.
	//
	// >  This parameter is available only if the git repository is private.
	//
	// example:
	//
	// aliyun-computenest
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The platform type. Valid values:
	//
	// - github
	//
	// - gitee
	//
	// - gitlab
	//
	// - codeup
	//
	// example:
	//
	// github
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The repository ID.
	//
	// example:
	//
	// 103
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// aliyun-computenest/quickstart-Lobexxx
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildPropertyCodeRepo) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildPropertyCodeRepo) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetBranch(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Branch = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetEndpoint(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Endpoint = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetOrgId(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.OrgId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetOwner(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Owner = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetPlatform(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Platform = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoId(v int64) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoName(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoName = &v
	return s
}

type CreateArtifactRequestArtifactProperty struct {
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity version of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// The image ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// m-0xij191j9cuev6xxxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// crr-d8o1nponyc2t1gcg
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `Public`: a public repository.
	//
	// 	- `Private`: a private repository.
	//
	// You can specify the RepoType or Summary parameter. The RepoType parameter is optional.
	//
	// example:
	//
	// Public
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The version tag of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The object URL of the deployment package.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityCode(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetImageId(v string) *CreateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRegionId(v string) *CreateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoId(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoName(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoType(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoType = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetTag(v string) *CreateArtifactRequestArtifactProperty {
	s.Tag = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetUrl(v string) *CreateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type CreateArtifactRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateArtifactRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestTag) SetKey(v string) *CreateArtifactRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactRequestTag) SetValue(v string) *CreateArtifactRequestTag {
	s.Value = &v
	return s
}

type CreateArtifactShrinkRequest struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the artifact build task. Valid values:
	//
	// - EcsImage: Build ECS (Elastic Container Service) image.
	//
	// - Dockerfile: Build container image based on Dockerfile.
	//
	// - Buildpacks: Build container image based on Buildpacks.
	//
	// - ContainerImage: Rebuild container image by renaming an existing container image.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment object.
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package. Valid values:
	//
	// 	- EcsImage: Elastic Compute Service (ECS) image.
	//
	// 	- AcrImage: container image.
	//
	// 	- File: Object Storage Service (OSS) object.
	//
	// 	- Script: script.
	//
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Test artifact
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The supported regions.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The custom tags.
	Tag []*CreateArtifactShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequest) SetArtifactBuildPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactBuildPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactBuildType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactBuildType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactId(v string) *CreateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetClientToken(v string) *CreateArtifactShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetDescription(v string) *CreateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetName(v string) *CreateArtifactShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetResourceGroupId(v string) *CreateArtifactShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *CreateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetTag(v []*CreateArtifactShrinkRequestTag) *CreateArtifactShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetVersionName(v string) *CreateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateArtifactShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequestTag) SetKey(v string) *CreateArtifactShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactShrinkRequestTag) SetValue(v string) *CreateArtifactShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateArtifactResponseBody struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment object.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Test artifact
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-11-11T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version of the deployment package.
	//
	// example:
	//
	// 1
	MaxVersion *int64 `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// [
	//
	// 			"cn-beijing",
	//
	// 			"cn-hangzhou",
	//
	// 			"cn-shanghai"
	//
	// 		]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponseBody) SetArtifactBuildProperty(v string) *CreateArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactBuildType(v string) *CreateArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactId(v string) *CreateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactProperty(v string) *CreateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactType(v string) *CreateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactVersion(v string) *CreateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetDescription(v string) *CreateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *CreateArtifactResponseBody) SetGmtModified(v string) *CreateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateArtifactResponseBody) SetMaxVersion(v int64) *CreateArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetName(v string) *CreateArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *CreateArtifactResponseBody) SetRequestId(v string) *CreateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetStatus(v string) *CreateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *CreateArtifactResponseBody) SetStatusDetail(v string) *CreateArtifactResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *CreateArtifactResponseBody) SetSupportRegionIds(v string) *CreateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *CreateArtifactResponseBody) SetVersionName(v string) *CreateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type CreateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponse) SetHeaders(v map[string]*string) *CreateArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactResponse) SetStatusCode(v int32) *CreateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactResponse) SetBody(v *CreateArtifactResponseBody) *CreateArtifactResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// {
	//
	//   "TemplateUrl": "http://template.file.url",
	//
	//   // 应用分组级别告警元数据
	//
	//   "ApplicationGroups": [
	//
	//     {
	//
	//       "Name": "applicationGroup1",
	//
	//       "TemplateUrl": "url1"
	//
	//     },
	//
	//     {
	//
	//       "Name": "applicationGroup2",
	//
	//       "TemplateUrl": "url2"
	//
	//     }
	//
	//   ]
	//
	// }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application. Valid values:
	//
	// 	- Manual: The application is manually approved.
	//
	// 	- AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The parameters for building the service
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Compliance check metadata.
	ComplianceMetadata *CreateServiceRequestComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"TemplateConfigs\\":[{\\"Name\\":\\"模板1\\",\\"Url\\":\\"oss://computenest-test/template"
	//
	//             + ".json?RegionId=cn-beijing\\",\\"PredefinedParameters\\":[{\\"Name\\":\\"低配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":40}},{\\"Name\\":\\"高配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":200}}]}]}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service.
	//
	// 	- false: performs a dry run for the request, and create a service if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable the hosted O\\&M feature for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is required if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {\\"RetentionDays\\":3}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// { "Logstores": [ { "LogstoreName": "access-log", "LogPath": "/home/admin/app/logs", # This parameter is not required for containers. Configure the parameter in the YAML file. "FilePattern": "access.log\\*" # This parameter is not required for containers. Configure the parameter in the YAML file. } ] }
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Whether resell is supported.
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25refu7r3opq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo []*CreateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The source service ID for resell。
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The source service version for resell。
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The custom tags.
	Tag []*CreateServiceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The version name.
	//
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetAlarmMetadata(v string) *CreateServiceRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetApprovalType(v string) *CreateServiceRequest {
	s.ApprovalType = &v
	return s
}

func (s *CreateServiceRequest) SetBuildParameters(v string) *CreateServiceRequest {
	s.BuildParameters = &v
	return s
}

func (s *CreateServiceRequest) SetClientToken(v string) *CreateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceRequest) SetComplianceMetadata(v *CreateServiceRequestComplianceMetadata) *CreateServiceRequest {
	s.ComplianceMetadata = v
	return s
}

func (s *CreateServiceRequest) SetDeployMetadata(v string) *CreateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetDeployType(v string) *CreateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *CreateServiceRequest) SetDryRun(v bool) *CreateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceRequest) SetDuration(v int64) *CreateServiceRequest {
	s.Duration = &v
	return s
}

func (s *CreateServiceRequest) SetIsSupportOperated(v bool) *CreateServiceRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *CreateServiceRequest) SetLicenseMetadata(v string) *CreateServiceRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetLogMetadata(v string) *CreateServiceRequest {
	s.LogMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetOperationMetadata(v string) *CreateServiceRequest {
	s.OperationMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetPolicyNames(v string) *CreateServiceRequest {
	s.PolicyNames = &v
	return s
}

func (s *CreateServiceRequest) SetRegionId(v string) *CreateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceRequest) SetResellable(v bool) *CreateServiceRequest {
	s.Resellable = &v
	return s
}

func (s *CreateServiceRequest) SetResourceGroupId(v string) *CreateServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceId(v string) *CreateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetServiceInfo(v []*CreateServiceRequestServiceInfo) *CreateServiceRequest {
	s.ServiceInfo = v
	return s
}

func (s *CreateServiceRequest) SetServiceType(v string) *CreateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateServiceRequest) SetShareType(v string) *CreateServiceRequest {
	s.ShareType = &v
	return s
}

func (s *CreateServiceRequest) SetSourceServiceId(v string) *CreateServiceRequest {
	s.SourceServiceId = &v
	return s
}

func (s *CreateServiceRequest) SetSourceServiceVersion(v string) *CreateServiceRequest {
	s.SourceServiceVersion = &v
	return s
}

func (s *CreateServiceRequest) SetTag(v []*CreateServiceRequestTag) *CreateServiceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceRequest) SetTenantType(v string) *CreateServiceRequest {
	s.TenantType = &v
	return s
}

func (s *CreateServiceRequest) SetTrialDuration(v int64) *CreateServiceRequest {
	s.TrialDuration = &v
	return s
}

func (s *CreateServiceRequest) SetUpgradeMetadata(v string) *CreateServiceRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *CreateServiceRequest) SetVersionName(v string) *CreateServiceRequest {
	s.VersionName = &v
	return s
}

type CreateServiceRequestComplianceMetadata struct {
	// The compliance package selected.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s CreateServiceRequestComplianceMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestComplianceMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestComplianceMetadata) SetCompliancePacks(v []*string) *CreateServiceRequestComplianceMetadata {
	s.CompliancePacks = v
	return s
}

type CreateServiceRequestServiceInfo struct {
	// Protocol document information about the service.
	Agreements []*CreateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The URL of the detailed description of the service.
	//
	// example:
	//
	// http://description.tidb.oss.url
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TiDB Database
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// TiDB是A公司自主设计、研发的开源分布式关系型数据库。
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
	Softwares []*CreateServiceRequestServiceInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s CreateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfo) SetAgreements(v []*CreateServiceRequestServiceInfoAgreements) *CreateServiceRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetImage(v string) *CreateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetLocale(v string) *CreateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetLongDescriptionUrl(v string) *CreateServiceRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetName(v string) *CreateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetShortDescription(v string) *CreateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *CreateServiceRequestServiceInfo) SetSoftwares(v []*CreateServiceRequestServiceInfoSoftwares) *CreateServiceRequestServiceInfo {
	s.Softwares = v
	return s
}

type CreateServiceRequestServiceInfoAgreements struct {
	// Protocol name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol url.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateServiceRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfoAgreements) SetName(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfoAgreements) SetUrl(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type CreateServiceRequestServiceInfoSoftwares struct {
	// The name of the software.
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServiceRequestServiceInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfoSoftwares) SetName(v string) *CreateServiceRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfoSoftwares) SetVersion(v string) *CreateServiceRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

type CreateServiceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Web
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestTag) SetKey(v string) *CreateServiceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceRequestTag) SetValue(v string) *CreateServiceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceShrinkRequest struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// {
	//
	//   "TemplateUrl": "http://template.file.url",
	//
	//   // 应用分组级别告警元数据
	//
	//   "ApplicationGroups": [
	//
	//     {
	//
	//       "Name": "applicationGroup1",
	//
	//       "TemplateUrl": "url1"
	//
	//     },
	//
	//     {
	//
	//       "Name": "applicationGroup2",
	//
	//       "TemplateUrl": "url2"
	//
	//     }
	//
	//   ]
	//
	// }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application. Valid values:
	//
	// 	- Manual: The application is manually approved.
	//
	// 	- AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The parameters for building the service
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Compliance check metadata.
	ComplianceMetadataShrink *string `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty"`
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"TemplateConfigs\\":[{\\"Name\\":\\"模板1\\",\\"Url\\":\\"oss://computenest-test/template"
	//
	//             + ".json?RegionId=cn-beijing\\",\\"PredefinedParameters\\":[{\\"Name\\":\\"低配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":40}},{\\"Name\\":\\"高配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":200}}]}]}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service.
	//
	// 	- false: performs a dry run for the request, and create a service if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable the hosted O\\&M feature for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is required if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {\\"RetentionDays\\":3}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// { "Logstores": [ { "LogstoreName": "access-log", "LogPath": "/home/admin/app/logs", # This parameter is not required for containers. Configure the parameter in the YAML file. "FilePattern": "access.log\\*" # This parameter is not required for containers. Configure the parameter in the YAML file. } ] }
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Whether resell is supported.
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25refu7r3opq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo []*CreateServiceShrinkRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The source service ID for resell。
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The source service version for resell。
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The custom tags.
	Tag []*CreateServiceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The version name.
	//
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequest) SetAlarmMetadata(v string) *CreateServiceShrinkRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetApprovalType(v string) *CreateServiceShrinkRequest {
	s.ApprovalType = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetBuildParameters(v string) *CreateServiceShrinkRequest {
	s.BuildParameters = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetClientToken(v string) *CreateServiceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetComplianceMetadataShrink(v string) *CreateServiceShrinkRequest {
	s.ComplianceMetadataShrink = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDeployMetadata(v string) *CreateServiceShrinkRequest {
	s.DeployMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDeployType(v string) *CreateServiceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDryRun(v bool) *CreateServiceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetDuration(v int64) *CreateServiceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetIsSupportOperated(v bool) *CreateServiceShrinkRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetLicenseMetadata(v string) *CreateServiceShrinkRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetLogMetadata(v string) *CreateServiceShrinkRequest {
	s.LogMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetOperationMetadata(v string) *CreateServiceShrinkRequest {
	s.OperationMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetPolicyNames(v string) *CreateServiceShrinkRequest {
	s.PolicyNames = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetRegionId(v string) *CreateServiceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetResellable(v bool) *CreateServiceShrinkRequest {
	s.Resellable = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetResourceGroupId(v string) *CreateServiceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetServiceId(v string) *CreateServiceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetServiceInfo(v []*CreateServiceShrinkRequestServiceInfo) *CreateServiceShrinkRequest {
	s.ServiceInfo = v
	return s
}

func (s *CreateServiceShrinkRequest) SetServiceType(v string) *CreateServiceShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetShareType(v string) *CreateServiceShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetSourceServiceId(v string) *CreateServiceShrinkRequest {
	s.SourceServiceId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetSourceServiceVersion(v string) *CreateServiceShrinkRequest {
	s.SourceServiceVersion = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetTag(v []*CreateServiceShrinkRequestTag) *CreateServiceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceShrinkRequest) SetTenantType(v string) *CreateServiceShrinkRequest {
	s.TenantType = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetTrialDuration(v int64) *CreateServiceShrinkRequest {
	s.TrialDuration = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetUpgradeMetadata(v string) *CreateServiceShrinkRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetVersionName(v string) *CreateServiceShrinkRequest {
	s.VersionName = &v
	return s
}

type CreateServiceShrinkRequestServiceInfo struct {
	// Protocol document information about the service.
	Agreements []*CreateServiceShrinkRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The URL of the detailed description of the service.
	//
	// example:
	//
	// http://description.tidb.oss.url
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TiDB Database
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// TiDB是A公司自主设计、研发的开源分布式关系型数据库。
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
	Softwares []*CreateServiceShrinkRequestServiceInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s CreateServiceShrinkRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfo) SetAgreements(v []*CreateServiceShrinkRequestServiceInfoAgreements) *CreateServiceShrinkRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetImage(v string) *CreateServiceShrinkRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetLocale(v string) *CreateServiceShrinkRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetLongDescriptionUrl(v string) *CreateServiceShrinkRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetName(v string) *CreateServiceShrinkRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetShortDescription(v string) *CreateServiceShrinkRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfo) SetSoftwares(v []*CreateServiceShrinkRequestServiceInfoSoftwares) *CreateServiceShrinkRequestServiceInfo {
	s.Softwares = v
	return s
}

type CreateServiceShrinkRequestServiceInfoAgreements struct {
	// Protocol name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol url.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateServiceShrinkRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) SetName(v string) *CreateServiceShrinkRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) SetUrl(v string) *CreateServiceShrinkRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type CreateServiceShrinkRequestServiceInfoSoftwares struct {
	// The name of the software.
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServiceShrinkRequestServiceInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) SetName(v string) *CreateServiceShrinkRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) SetVersion(v string) *CreateServiceShrinkRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

type CreateServiceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Web
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestTag) SetKey(v string) *CreateServiceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceShrinkRequestTag) SetValue(v string) *CreateServiceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceResponseBody struct {
	// The dry run result.
	DryRunResult *CreateServiceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The status of the service.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetDryRunResult(v *CreateServiceResponseBodyDryRunResult) *CreateServiceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceResponseBody) SetVersion(v string) *CreateServiceResponseBody {
	s.Version = &v
	return s
}

type CreateServiceResponseBodyDryRunResult struct {
	// The required ram policy for deploying role.
	RolePolicy *CreateServiceResponseBodyDryRunResultRolePolicy `json:"RolePolicy,omitempty" xml:"RolePolicy,omitempty" type:"Struct"`
}

func (s CreateServiceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResult) SetRolePolicy(v *CreateServiceResponseBodyDryRunResultRolePolicy) *CreateServiceResponseBodyDryRunResult {
	s.RolePolicy = v
	return s
}

type CreateServiceResponseBodyDryRunResultRolePolicy struct {
	// The missing ram policy for deploying role.
	MissingPolicy []*CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The required ram policy for deploying role.
	//
	// example:
	//
	// {
	//
	// 	"Statement": [{
	//
	// 		"Action": ["oos:CancelExecutions", "oos:DeleteExecutions", "oos:GetTemplate", "oos:ListExecutions", "oos:ListTemplates", "oos:NotifyExecution", "oos:StartExecution"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ram:PassRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ros:CreateStack", "ros:GetStack", "ros:UpdateStack", "ros:ListStackEvents", "ros:ListStackResources", "ros:ListStackResources", "ros:DeleteStack", "ram:GetRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}],
	//
	// 	"Version": "1"
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateServiceResponseBodyDryRunResultRolePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResultRolePolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) SetMissingPolicy(v []*CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) *CreateServiceResponseBodyDryRunResultRolePolicy {
	s.MissingPolicy = v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) SetPolicy(v string) *CreateServiceResponseBodyDryRunResultRolePolicy {
	s.Policy = &v
	return s
}

type CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy struct {
	// The Actions.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// Resource in ram policy.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The service name in ram policy.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetAction(v []*string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Action = v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetResource(v string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetServiceName(v string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.ServiceName = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when the service instance was released.
	//
	// >  This parameter is available only for the service instances that are managed by service providers.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-12T03:13:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the service instance. The value must meet the following requirements:
	//
	// 	- The name cannot exceed 64 characters in length.
	//
	// 	- It can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that are specified for service instance deployment.
	//
	// >  If you want to specify the region in which the service instance is deployed, you must specify the information in Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// 	- ap-southeast-1: Singapore
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The custom tags.
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The template name. You must specify a template name if the service supports multiple templates.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEndTime(v string) *CreateServiceInstanceRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetUserId(v string) *CreateServiceInstanceRequest {
	s.UserId = &v
	return s
}

type CreateServiceInstanceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when the service instance was released.
	//
	// >  This parameter is available only for the service instances that are managed by service providers.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-12T03:13:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the service instance. The value must meet the following requirements:
	//
	// 	- The name cannot exceed 64 characters in length.
	//
	// 	- It can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that are specified for service instance deployment.
	//
	// >  If you want to specify the region in which the service instance is deployed, you must specify the information in Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// 	- ap-southeast-1: Singapore
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The custom tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The template name. You must specify a template name if the service supports multiple templates.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEndTime(v string) *CreateServiceInstanceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetUserId(v string) *CreateServiceInstanceShrinkRequest {
	s.UserId = &v
	return s
}

type CreateServiceInstanceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceTestCaseRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Service Test case name.
	//
	// This parameter is required.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// The service test config
	//
	// This parameter is required.
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s CreateServiceTestCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseRequest) SetRegionId(v string) *CreateServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetServiceId(v string) *CreateServiceTestCaseRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetServiceVersion(v string) *CreateServiceTestCaseRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTemplateName(v string) *CreateServiceTestCaseRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTestCaseName(v string) *CreateServiceTestCaseRequest {
	s.TestCaseName = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTestConfig(v string) *CreateServiceTestCaseRequest {
	s.TestConfig = &v
	return s
}

type CreateServiceTestCaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The test case Id
	//
	// example:
	//
	// stc-5ba03a6a9a2746be8739
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
}

func (s CreateServiceTestCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseResponseBody) SetRequestId(v string) *CreateServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTestCaseResponseBody) SetTestCaseId(v string) *CreateServiceTestCaseResponseBody {
	s.TestCaseId = &v
	return s
}

type CreateServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTestCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseResponse) SetHeaders(v map[string]*string) *CreateServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTestCaseResponse) SetStatusCode(v int32) *CreateServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTestCaseResponse) SetBody(v *CreateServiceTestCaseResponseBody) *CreateServiceTestCaseResponse {
	s.Body = v
	return s
}

type CreateServiceTestTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// nametest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The Task Execution Region
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
	// The service test case ids.
	//
	// This parameter is required.
	TestCaseIds []*string `json:"TestCaseIds,omitempty" xml:"TestCaseIds,omitempty" type:"Repeated"`
}

func (s CreateServiceTestTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskRequest) SetRegionId(v string) *CreateServiceTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTaskName(v string) *CreateServiceTestTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTaskRegionId(v string) *CreateServiceTestTaskRequest {
	s.TaskRegionId = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTestCaseIds(v []*string) *CreateServiceTestTaskRequest {
	s.TestCaseIds = v
	return s
}

type CreateServiceTestTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-568c2c5a687a409b977e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateServiceTestTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskResponseBody) SetRequestId(v string) *CreateServiceTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTestTaskResponseBody) SetTaskId(v string) *CreateServiceTestTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateServiceTestTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTestTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskResponse) SetHeaders(v map[string]*string) *CreateServiceTestTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTestTaskResponse) SetStatusCode(v int32) *CreateServiceTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTestTaskResponse) SetBody(v *CreateServiceTestTaskResponseBody) *CreateServiceTestTaskResponse {
	s.Body = v
	return s
}

type CreateServiceUsageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageRequest) SetClientToken(v string) *CreateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageRequest) SetRegionId(v string) *CreateServiceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceUsageRequest) SetServiceId(v string) *CreateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

type CreateServiceUsageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponseBody) SetRequestId(v string) *CreateServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponse) SetHeaders(v map[string]*string) *CreateServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceUsageResponse) SetStatusCode(v int32) *CreateServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceUsageResponse) SetBody(v *CreateServiceUsageResponseBody) *CreateServiceUsageResponse {
	s.Body = v
	return s
}

type CreateSupplierRegistrationRequest struct {
	// Contact email
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx@xxx.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// Contact number
	//
	// This parameter is required.
	//
	// example:
	//
	// 186xxxxxxxxx
	ContactNumber *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	// Contact person
	//
	// This parameter is required.
	//
	// example:
	//
	// Mike
	ContactPerson *string `json:"ContactPerson,omitempty" xml:"ContactPerson,omitempty"`
	// Contact person tiltle
	//
	// This parameter is required.
	//
	// example:
	//
	// CTO
	ContactPersonTitle *string `json:"ContactPersonTitle,omitempty" xml:"ContactPersonTitle,omitempty"`
	// Whether to enable the resell mode
	//
	// example:
	//
	// true
	EnableResellerMode *bool `json:"EnableResellerMode,omitempty" xml:"EnableResellerMode,omitempty"`
	// Annual product revenue
	//
	// example:
	//
	// 1000
	ProductAnnualRevenue *string `json:"ProductAnnualRevenue,omitempty" xml:"ProductAnnualRevenue,omitempty"`
	// The business of product
	//
	// example:
	//
	// AI
	ProductBusiness *string `json:"ProductBusiness,omitempty" xml:"ProductBusiness,omitempty"`
	// Product delivery type
	//
	// This parameter is required.
	ProductDeliveryTypes []*string `json:"ProductDeliveryTypes,omitempty" xml:"ProductDeliveryTypes,omitempty" type:"Repeated"`
	// The publish time of product
	//
	// example:
	//
	// 2020.10.10
	ProductPublishTime *string `json:"ProductPublishTime,omitempty" xml:"ProductPublishTime,omitempty"`
	// Product sell type
	//
	// This parameter is required.
	ProductSellTypes []*string `json:"ProductSellTypes,omitempty" xml:"ProductSellTypes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of resell business.
	//
	// example:
	//
	// empty
	ResellBusinessDesc *string `json:"ResellBusinessDesc,omitempty" xml:"ResellBusinessDesc,omitempty"`
	// The demands of service providers.
	//
	// example:
	//
	// empty
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The description of service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/34DB-4F4C-9373-003AA060****.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The english name of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierNameEn *string `json:"SupplierNameEn,omitempty" xml:"SupplierNameEn,omitempty"`
	// The URL of the service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s CreateSupplierRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSupplierRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationRequest) SetContactEmail(v string) *CreateSupplierRegistrationRequest {
	s.ContactEmail = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactNumber(v string) *CreateSupplierRegistrationRequest {
	s.ContactNumber = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactPerson(v string) *CreateSupplierRegistrationRequest {
	s.ContactPerson = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetContactPersonTitle(v string) *CreateSupplierRegistrationRequest {
	s.ContactPersonTitle = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetEnableResellerMode(v bool) *CreateSupplierRegistrationRequest {
	s.EnableResellerMode = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductAnnualRevenue(v string) *CreateSupplierRegistrationRequest {
	s.ProductAnnualRevenue = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductBusiness(v string) *CreateSupplierRegistrationRequest {
	s.ProductBusiness = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductDeliveryTypes(v []*string) *CreateSupplierRegistrationRequest {
	s.ProductDeliveryTypes = v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductPublishTime(v string) *CreateSupplierRegistrationRequest {
	s.ProductPublishTime = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetProductSellTypes(v []*string) *CreateSupplierRegistrationRequest {
	s.ProductSellTypes = v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetRegionId(v string) *CreateSupplierRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetResellBusinessDesc(v string) *CreateSupplierRegistrationRequest {
	s.ResellBusinessDesc = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSuggestion(v string) *CreateSupplierRegistrationRequest {
	s.Suggestion = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierDesc(v string) *CreateSupplierRegistrationRequest {
	s.SupplierDesc = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierLogo(v string) *CreateSupplierRegistrationRequest {
	s.SupplierLogo = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierName(v string) *CreateSupplierRegistrationRequest {
	s.SupplierName = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierNameEn(v string) *CreateSupplierRegistrationRequest {
	s.SupplierNameEn = &v
	return s
}

func (s *CreateSupplierRegistrationRequest) SetSupplierUrl(v string) *CreateSupplierRegistrationRequest {
	s.SupplierUrl = &v
	return s
}

type CreateSupplierRegistrationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C4A145D8-xxxx-xxxx-xxxx-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupplierRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSupplierRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationResponseBody) SetRequestId(v string) *CreateSupplierRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type CreateSupplierRegistrationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSupplierRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSupplierRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSupplierRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationResponse) SetHeaders(v map[string]*string) *CreateSupplierRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CreateSupplierRegistrationResponse) SetStatusCode(v int32) *CreateSupplierRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSupplierRegistrationResponse) SetBody(v *CreateSupplierRegistrationResponseBody) *CreateSupplierRegistrationResponse {
	s.Body = v
	return s
}

type DeleteArtifactRequest struct {
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The version of the artifact.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactRequest) SetArtifactId(v string) *DeleteArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *DeleteArtifactRequest) SetArtifactVersion(v string) *DeleteArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

func (s *DeleteArtifactRequest) SetClientToken(v string) *DeleteArtifactRequest {
	s.ClientToken = &v
	return s
}

type DeleteArtifactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponseBody) SetRequestId(v string) *DeleteArtifactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponse) SetHeaders(v map[string]*string) *DeleteArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactResponse) SetStatusCode(v int32) *DeleteArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactResponse) SetBody(v *DeleteArtifactResponseBody) *DeleteArtifactResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetClientToken(v string) *DeleteServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceRequest) SetRegionId(v string) *DeleteServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceId(v string) *DeleteServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteServiceRequest) SetServiceVersion(v string) *DeleteServiceRequest {
	s.ServiceVersion = &v
	return s
}

type DeleteServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the service instances.
	//
	// This parameter is required.
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DB140E67-D75F-5585-946E-41D8DC8F4E00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type DeleteServiceTestCaseRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service test case id.
	//
	// This parameter is required.
	//
	// example:
	//
	// stc-0b2a3ad7e1de4c299eec
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
}

func (s DeleteServiceTestCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseRequest) SetRegionId(v string) *DeleteServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceTestCaseRequest) SetTestCaseId(v string) *DeleteServiceTestCaseRequest {
	s.TestCaseId = &v
	return s
}

type DeleteServiceTestCaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceTestCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseResponseBody) SetRequestId(v string) *DeleteServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceTestCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseResponse) SetHeaders(v map[string]*string) *DeleteServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceTestCaseResponse) SetStatusCode(v int32) *DeleteServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceTestCaseResponse) SetBody(v *DeleteServiceTestCaseResponseBody) *DeleteServiceTestCaseResponse {
	s.Body = v
	return s
}

type DeployServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s DeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceRequest) SetClientToken(v string) *DeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetRegionId(v string) *DeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetServiceInstanceId(v string) *DeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type DeployServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B8A6AEA6-0D8F-589A-A7FF-B44FD842456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponseBody) SetRequestId(v string) *DeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeployServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponse) SetHeaders(v map[string]*string) *DeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceInstanceResponse) SetStatusCode(v int32) *DeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployServiceInstanceResponse) SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse {
	s.Body = v
	return s
}

type GenerateDefaultServiceTestConfigRequest struct {
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-41ad58439b4b4bf8ae73
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// test-1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GenerateDefaultServiceTestConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigRequest) SetServiceId(v string) *GenerateDefaultServiceTestConfigRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigRequest) SetServiceVersion(v string) *GenerateDefaultServiceTestConfigRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigRequest) SetTemplateName(v string) *GenerateDefaultServiceTestConfigRequest {
	s.TemplateName = &v
	return s
}

type GenerateDefaultServiceTestConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service test config
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s GenerateDefaultServiceTestConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigResponseBody) SetRequestId(v string) *GenerateDefaultServiceTestConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponseBody) SetTestConfig(v string) *GenerateDefaultServiceTestConfigResponseBody {
	s.TestConfig = &v
	return s
}

type GenerateDefaultServiceTestConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDefaultServiceTestConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDefaultServiceTestConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigResponse) SetHeaders(v map[string]*string) *GenerateDefaultServiceTestConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponse) SetStatusCode(v int32) *GenerateDefaultServiceTestConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponse) SetBody(v *GenerateDefaultServiceTestConfigResponseBody) *GenerateDefaultServiceTestConfigResponse {
	s.Body = v
	return s
}

type GenerateServicePolicyRequest struct {
	// The type of operation N for which you want to generate the policy information.
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-6c20f0f8085645xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GenerateServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyRequest) SetOperationTypes(v []*string) *GenerateServicePolicyRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateServicePolicyRequest) SetRegionId(v string) *GenerateServicePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceId(v string) *GenerateServicePolicyRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceVersion(v string) *GenerateServicePolicyRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTemplateName(v string) *GenerateServicePolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTrialType(v string) *GenerateServicePolicyRequest {
	s.TrialType = &v
	return s
}

type GenerateServicePolicyResponseBody struct {
	// The policies that are missing.
	MissingPolicy []*GenerateServicePolicyResponseBodyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The RAM policy.
	//
	// example:
	//
	// {Statement": [{ "Action": ["oos:*"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:DescribeInstances"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:RunInstance"], "Effect": "Allow", "Resource": "*"}], "Version": "1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5040BE9E-8DA2-5C9D-9B70-0EE6027A14BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBody) SetMissingPolicy(v []*GenerateServicePolicyResponseBodyMissingPolicy) *GenerateServicePolicyResponseBody {
	s.MissingPolicy = v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetPolicy(v string) *GenerateServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetRequestId(v string) *GenerateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateServicePolicyResponseBodyMissingPolicy struct {
	// Operations on specific resources.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The specific objects authorized. An asterisk (*) denotes all resources.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetAction(v []*string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Action = v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetResource(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetServiceName(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.ServiceName = &v
	return s
}

type GenerateServicePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponse) SetHeaders(v map[string]*string) *GenerateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateServicePolicyResponse) SetStatusCode(v int32) *GenerateServicePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateServicePolicyResponse) SetBody(v *GenerateServicePolicyResponseBody) *GenerateServicePolicyResponse {
	s.Body = v
	return s
}

type GetArtifactRequest struct {
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// gpu-test
	ArtifactName *string `json:"ArtifactName,omitempty" xml:"ArtifactName,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 1
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) SetArtifactId(v string) *GetArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactName(v string) *GetArtifactRequest {
	s.ArtifactName = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactVersion(v string) *GetArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

type GetArtifactResponseBody struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version of the deployment package.
	//
	// example:
	//
	// 2
	MaxVersion *int64 `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact They can only change from Automatic to Public. Options:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The distribution progress of the deployment package.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// ["cn-hangzhou","cn-beijing"]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The tags of the deployment package.
	Tags []*GetArtifactResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBody) SetArtifactBuildProperty(v string) *GetArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactBuildType(v string) *GetArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactId(v string) *GetArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactProperty(v string) *GetArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactType(v string) *GetArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactVersion(v string) *GetArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetDescription(v string) *GetArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *GetArtifactResponseBody) SetGmtModified(v string) *GetArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetArtifactResponseBody) SetMaxVersion(v int64) *GetArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetName(v string) *GetArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *GetArtifactResponseBody) SetPermissionType(v string) *GetArtifactResponseBody {
	s.PermissionType = &v
	return s
}

func (s *GetArtifactResponseBody) SetProgress(v string) *GetArtifactResponseBody {
	s.Progress = &v
	return s
}

func (s *GetArtifactResponseBody) SetRequestId(v string) *GetArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactResponseBody) SetResourceGroupId(v string) *GetArtifactResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatus(v string) *GetArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatusDetail(v string) *GetArtifactResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetArtifactResponseBody) SetSupportRegionIds(v string) *GetArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *GetArtifactResponseBody) SetTags(v []*GetArtifactResponseBodyTags) *GetArtifactResponseBody {
	s.Tags = v
	return s
}

func (s *GetArtifactResponseBody) SetVersionName(v string) *GetArtifactResponseBody {
	s.VersionName = &v
	return s
}

type GetArtifactResponseBodyTags struct {
	// The tag key of the deployment package.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the deployment package.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetArtifactResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBodyTags) SetKey(v string) *GetArtifactResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetArtifactResponseBodyTags) SetValue(v string) *GetArtifactResponseBodyTags {
	s.Value = &v
	return s
}

type GetArtifactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactResponse) SetHeaders(v map[string]*string) *GetArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactResponse) SetStatusCode(v int32) *GetArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactResponse) SetBody(v *GetArtifactResponseBody) *GetArtifactResponse {
	s.Body = v
	return s
}

type GetArtifactRepositoryCredentialsRequest struct {
	// The type of the deployment package. Valid values:
	//
	// 	- File: Object Storage Service (OSS) object.
	//
	// 	- AcrImage: container image.
	//
	// This parameter is required.
	//
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsRequest) SetArtifactType(v string) *GetArtifactRepositoryCredentialsRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsRequest) SetDeployRegionId(v string) *GetArtifactRepositoryCredentialsRequest {
	s.DeployRegionId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBody struct {
	// The information about the resources that can be uploaded.
	AvailableResources []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	// The credentials.
	Credentials *GetArtifactRepositoryCredentialsResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// The time when the credentials expired.
	//
	// example:
	//
	// 1526549792000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetAvailableResources(v []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources) *GetArtifactRepositoryCredentialsResponseBody {
	s.AvailableResources = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetCredentials(v *GetArtifactRepositoryCredentialsResponseBodyCredentials) *GetArtifactRepositoryCredentialsResponseBody {
	s.Credentials = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetExpireDate(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetRequestId(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.RequestId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyAvailableResources struct {
	// The path.
	//
	// example:
	//
	// "/xxx/"
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID.
	//
	// example:
	//
	// oss-cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// computenest-artifacts-draft-cn-hangzhou
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetPath(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.Path = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRegionId(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RegionId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRepositoryName(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RepositoryName = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyCredentials struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.xxx
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// xxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The password.
	//
	// example:
	//
	// eyJ0aW1lIjoiMTUyNjU0OTc5:0705733****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The Security Token Service (STS) token.
	//
	// example:
	//
	// xxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The username.
	//
	// example:
	//
	// xxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeyId(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeySecret(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetPassword(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Password = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetSecurityToken(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetUsername(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Username = &v
	return s
}

type GetArtifactRepositoryCredentialsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactRepositoryCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponse) SetHeaders(v map[string]*string) *GetArtifactRepositoryCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetStatusCode(v int32) *GetArtifactRepositoryCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetBody(v *GetArtifactRepositoryCredentialsResponseBody) *GetArtifactRepositoryCredentialsResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	// Specifies whether to filter information based on Alibaba Cloud account IDs.
	//
	// example:
	//
	// false
	FilterAliUid *bool `json:"FilterAliUid,omitempty" xml:"FilterAliUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-4ee86df83fd948******
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The Service Instance Id.
	//
	// example:
	//
	// si-85b1exxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The Service Name.
	//
	// example:
	//
	// Wordpress社区版
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Resell: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	SharedAccountType *string `json:"SharedAccountType,omitempty" xml:"SharedAccountType,omitempty"`
	// The information that you want to query.
	ShowDetail []*string `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty" type:"Repeated"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetFilterAliUid(v bool) *GetServiceRequest {
	s.FilterAliUid = &v
	return s
}

func (s *GetServiceRequest) SetRegionId(v string) *GetServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v string) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceInstanceId(v string) *GetServiceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceName(v string) *GetServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceRequest) SetServiceVersion(v string) *GetServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRequest) SetSharedAccountType(v string) *GetServiceRequest {
	s.SharedAccountType = &v
	return s
}

func (s *GetServiceRequest) SetShowDetail(v []*string) *GetServiceRequest {
	s.ShowDetail = v
	return s
}

type GetServiceResponseBody struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// {
	//
	//   "TemplateUrl": "http://template.file.url",
	//
	//   // 应用分组级别告警元数据
	//
	//   "ApplicationGroups": [
	//
	//     {
	//
	//       "Name": "applicationGroup1",
	//
	//       "TemplateUrl": "url1"
	//
	//     },
	//
	//     {
	//
	//       "Name": "applicationGroup2",
	//
	//       "TemplateUrl": "url2"
	//
	//     }
	//
	//   ]
	//
	// }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application. Valid values:
	//
	// 	- Manual: The application is manually approved.
	//
	// 	- AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The information of build service information.
	//
	// example:
	//
	// { "RepoUrl": "https://github.com/user/repo.git", "Brancn": "main"}
	BuildInfo *string `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	// The parameters for building the service
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The category of the service.
	//
	// example:
	//
	// DevOps
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity details.
	Commodity *GetServiceResponseBodyCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Compliance check metadata.
	ComplianceMetadata *GetServiceResponseBodyComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The time when the service was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The binding configurations of the commodity module.
	//
	// example:
	//
	// componesConfigs
	CrossRegionConnectionStatus *string `json:"CrossRegionConnectionStatus,omitempty" xml:"CrossRegionConnectionStatus,omitempty"`
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"TemplateConfigs\\":[{\\"Name\\":\\"模板1\\",\\"Url\\":\\"oss://computenest-test/template"
	//
	//             + ".json?RegionId=cn-beijing\\",\\"PredefinedParameters\\":[{\\"Name\\":\\"低配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":40}},{\\"Name\\":\\"高配版\\","
	//
	//             + "\\"Parameters\\":{\\"InstanceType\\":\\"ecs.g5.large\\",\\"DataDiskSize\\":200}}]}]}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// 	- container: The service is deployed by using a container.
	//
	// 	- pkg: The service is deployed by using a package.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The report source.
	EntitySource map[string]*string `json:"EntitySource,omitempty" xml:"EntitySource,omitempty"`
	// Indicates whether the hosted O\\&M feature is enabled for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is returned if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {"renewType":"MONTHLY"}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// { "Logstores": [ { "LogstoreName": "access-log", "LogPath": "/home/admin/app/logs", # This parameter is not required for containers. Configure the parameter in the YAML file. "FilePattern": "access.log\\*" # This parameter is not required for containers. Configure the parameter in the YAML file. } ] }
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The source for which fees are generated. Valid values:
	//
	// 	- None: No fees are generated.
	//
	// 	- Marketplace: Fees are generated for Alibaba Cloud Marketplace.
	//
	// 	- Custom: The custom fees.
	//
	// example:
	//
	// None
	PayFromType *string `json:"PayFromType,omitempty" xml:"PayFromType,omitempty"`
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The registration ID.
	//
	// example:
	//
	// sr-04056c2ab4b94bxxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the distribution is supported. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzuqyxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The URL of the service audit file.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1690707531xxxxxx/service-document/be3382cd-xxxx-xxxx-xxxx-f8707ec12879.docx
	ServiceAuditDocumentUrl *string `json:"ServiceAuditDocumentUrl,omitempty" xml:"ServiceAuditDocumentUrl,omitempty"`
	// Indicates whether the service is visible. Valid values:
	//
	// 	- INVISIBLE
	//
	// 	- DISCOVERABLE
	//
	// example:
	//
	// DISCOVERABLE
	ServiceDiscoverable *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// Service document information.
	ServiceDocumentInfos []*GetServiceResponseBodyServiceDocumentInfos `json:"ServiceDocumentInfos,omitempty" xml:"ServiceDocumentInfos,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// http://example2.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The share status of the instance.
	//
	// > This parameter is discontinued.
	//
	// example:
	//
	// This parameter is discontinued.
	ShareTypeStatus *string `json:"ShareTypeStatus,omitempty" xml:"ShareTypeStatus,omitempty"`
	// The ID of the distribution source service.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The version of the distribution source service.
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The name of the distribution source service provider.
	//
	// example:
	//
	// SourceSupplier
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// The statistics.
	Statistic *GetServiceResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	// The status of the service. Valid values:
	//
	// 	- Draft: The service is a draft.
	//
	// 	- Submitted: The service is submitted for review. You cannot modify services in this state.
	//
	// 	- Approved: The service is approved. You cannot modify services in this state. You can publish services in this state.
	//
	// 	- Launching: The service is being published.
	//
	// 	- Online: The service is published.
	//
	// 	- Offline: The service is unpublished.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the service status.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider.
	SupportContacts []*GetServiceResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
	// The service tags.
	Tags []*GetServiceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The status of the test. Valid values:
	//
	// 	- `CONFIG_IS_NULL`: No test configurations exist.
	//
	// 	- `SERVICE_TEST_SUCCEED`: The service passed the test.
	//
	// 	- `SERVICE_TSET_DOING`: The service does not pass the test.
	//
	// example:
	//
	// SERVICE_TEST_SUCCEED
	TestStatus *string `json:"TestStatus,omitempty" xml:"TestStatus,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The time when the service was updated.
	//
	// example:
	//
	// 2021-05-22T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Indicates whether the service is a virtual Internet service. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
	// The ID of the virtual Internet service.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	VirtualInternetServiceId *string `json:"VirtualInternetServiceId,omitempty" xml:"VirtualInternetServiceId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetAlarmMetadata(v string) *GetServiceResponseBody {
	s.AlarmMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetApprovalType(v string) *GetServiceResponseBody {
	s.ApprovalType = &v
	return s
}

func (s *GetServiceResponseBody) SetBuildInfo(v string) *GetServiceResponseBody {
	s.BuildInfo = &v
	return s
}

func (s *GetServiceResponseBody) SetBuildParameters(v string) *GetServiceResponseBody {
	s.BuildParameters = &v
	return s
}

func (s *GetServiceResponseBody) SetCategories(v string) *GetServiceResponseBody {
	s.Categories = &v
	return s
}

func (s *GetServiceResponseBody) SetCommodity(v *GetServiceResponseBodyCommodity) *GetServiceResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceResponseBody) SetComplianceMetadata(v *GetServiceResponseBodyComplianceMetadata) *GetServiceResponseBody {
	s.ComplianceMetadata = v
	return s
}

func (s *GetServiceResponseBody) SetCreateTime(v string) *GetServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceResponseBody) SetCrossRegionConnectionStatus(v string) *GetServiceResponseBody {
	s.CrossRegionConnectionStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployMetadata(v string) *GetServiceResponseBody {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployType(v string) *GetServiceResponseBody {
	s.DeployType = &v
	return s
}

func (s *GetServiceResponseBody) SetDuration(v int64) *GetServiceResponseBody {
	s.Duration = &v
	return s
}

func (s *GetServiceResponseBody) SetEntitySource(v map[string]*string) *GetServiceResponseBody {
	s.EntitySource = v
	return s
}

func (s *GetServiceResponseBody) SetIsSupportOperated(v bool) *GetServiceResponseBody {
	s.IsSupportOperated = &v
	return s
}

func (s *GetServiceResponseBody) SetLicenseMetadata(v string) *GetServiceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetLogMetadata(v string) *GetServiceResponseBody {
	s.LogMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetOperationMetadata(v string) *GetServiceResponseBody {
	s.OperationMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetPayFromType(v string) *GetServiceResponseBody {
	s.PayFromType = &v
	return s
}

func (s *GetServiceResponseBody) SetPermission(v string) *GetServiceResponseBody {
	s.Permission = &v
	return s
}

func (s *GetServiceResponseBody) SetPolicyNames(v string) *GetServiceResponseBody {
	s.PolicyNames = &v
	return s
}

func (s *GetServiceResponseBody) SetProgress(v int64) *GetServiceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceResponseBody) SetPublishTime(v string) *GetServiceResponseBody {
	s.PublishTime = &v
	return s
}

func (s *GetServiceResponseBody) SetRegistrationId(v string) *GetServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetResellable(v bool) *GetServiceResponseBody {
	s.Resellable = &v
	return s
}

func (s *GetServiceResponseBody) SetResourceGroupId(v string) *GetServiceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceAuditDocumentUrl(v string) *GetServiceResponseBody {
	s.ServiceAuditDocumentUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceDiscoverable(v string) *GetServiceResponseBody {
	s.ServiceDiscoverable = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceDocumentInfos(v []*GetServiceResponseBodyServiceDocumentInfos) *GetServiceResponseBody {
	s.ServiceDocumentInfos = v
	return s
}

func (s *GetServiceResponseBody) SetServiceId(v string) *GetServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceInfos(v []*GetServiceResponseBodyServiceInfos) *GetServiceResponseBody {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceResponseBody) SetServiceProductUrl(v string) *GetServiceResponseBody {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceType(v string) *GetServiceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceResponseBody) SetShareType(v string) *GetServiceResponseBody {
	s.ShareType = &v
	return s
}

func (s *GetServiceResponseBody) SetShareTypeStatus(v string) *GetServiceResponseBody {
	s.ShareTypeStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceServiceId(v string) *GetServiceResponseBody {
	s.SourceServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceServiceVersion(v string) *GetServiceResponseBody {
	s.SourceServiceVersion = &v
	return s
}

func (s *GetServiceResponseBody) SetSourceSupplierName(v string) *GetServiceResponseBody {
	s.SourceSupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetStatistic(v *GetServiceResponseBodyStatistic) *GetServiceResponseBody {
	s.Statistic = v
	return s
}

func (s *GetServiceResponseBody) SetStatus(v string) *GetServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceResponseBody) SetStatusDetail(v string) *GetServiceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierName(v string) *GetServiceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUrl(v string) *GetServiceResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetSupportContacts(v []*GetServiceResponseBodySupportContacts) *GetServiceResponseBody {
	s.SupportContacts = v
	return s
}

func (s *GetServiceResponseBody) SetTags(v []*GetServiceResponseBodyTags) *GetServiceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceResponseBody) SetTenantType(v string) *GetServiceResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetServiceResponseBody) SetTestStatus(v string) *GetServiceResponseBody {
	s.TestStatus = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialDuration(v int64) *GetServiceResponseBody {
	s.TrialDuration = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialType(v string) *GetServiceResponseBody {
	s.TrialType = &v
	return s
}

func (s *GetServiceResponseBody) SetUpdateTime(v string) *GetServiceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceResponseBody) SetUpgradeMetadata(v string) *GetServiceResponseBody {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetVersion(v string) *GetServiceResponseBody {
	s.Version = &v
	return s
}

func (s *GetServiceResponseBody) SetVersionName(v string) *GetServiceResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetServiceResponseBody) SetVirtualInternetService(v string) *GetServiceResponseBody {
	s.VirtualInternetService = &v
	return s
}

func (s *GetServiceResponseBody) SetVirtualInternetServiceId(v string) *GetServiceResponseBody {
	s.VirtualInternetServiceId = &v
	return s
}

type GetServiceResponseBodyCommodity struct {
	// The billing method of the service. Valid values:
	//
	// 	- **PREPAY*	- (default): subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity modules.
	Components []*string `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The configuration metadata related to Lingxiao.
	CssMetadata *GetServiceResponseBodyCommodityCssMetadata `json:"CssMetadata,omitempty" xml:"CssMetadata,omitempty" type:"Struct"`
	// The metadata of Alibaba Cloud Marketplace.
	MarketplaceMetadata *GetServiceResponseBodyCommodityMarketplaceMetadata `json:"MarketplaceMetadata,omitempty" xml:"MarketplaceMetadata,omitempty" type:"Struct"`
	// The information about the billable item.
	MeteringEntities []*GetServiceResponseBodyCommodityMeteringEntities `json:"MeteringEntities,omitempty" xml:"MeteringEntities,omitempty" type:"Repeated"`
	// The configuration metadata related to Saas Boost.
	//
	// example:
	//
	// { "Enabled":false // The public endpoint of the SaaS Boost instance. "PublicAccessUrl":"https://example.com" }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The specification details of the service in Alibaba Cloud Marketplace.
	Specifications []*GetServiceResponseBodyCommoditySpecifications `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- marketplace: Alibaba Cloud Marketplace.
	//
	// 	- Css: Lingxiao.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodity) SetChargeType(v string) *GetServiceResponseBodyCommodity {
	s.ChargeType = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCommodityCode(v string) *GetServiceResponseBodyCommodity {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetComponents(v []*string) *GetServiceResponseBodyCommodity {
	s.Components = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCssMetadata(v *GetServiceResponseBodyCommodityCssMetadata) *GetServiceResponseBodyCommodity {
	s.CssMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMarketplaceMetadata(v *GetServiceResponseBodyCommodityMarketplaceMetadata) *GetServiceResponseBodyCommodity {
	s.MarketplaceMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMeteringEntities(v []*GetServiceResponseBodyCommodityMeteringEntities) *GetServiceResponseBodyCommodity {
	s.MeteringEntities = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSaasBoostMetadata(v string) *GetServiceResponseBodyCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSpecifications(v []*GetServiceResponseBodyCommoditySpecifications) *GetServiceResponseBodyCommodity {
	s.Specifications = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetType(v string) *GetServiceResponseBodyCommodity {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadata struct {
	// The mapping information about the billing items.
	ComponentsMappings []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
	// Metering item configuration information.
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The binding relationship between package and measurement dimension.
	MeteringEntityMappings []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityCssMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetComponentsMappings(v []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.ComponentsMappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetMeteringEntityExtraInfos(v []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) *GetServiceResponseBodyCommodityCssMetadata {
	s.MeteringEntityExtraInfos = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetMeteringEntityMappings(v []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.MeteringEntityMappings = v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataComponentsMappings struct {
	// The mappings.
	Mappings map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetMappings(v map[string]*string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.Mappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos struct {
	// The ID of the entity.
	//
	// example:
	//
	// cmgj0048****-Frequency-1
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// Name of a measurement indicator.
	//
	// example:
	//
	// AvgMemory
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Custom PromQL.
	//
	// example:
	//
	// avg_over_time(count(kube_pod_info{namespace=\\"default\\"})[1h:1m])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// Measurement indicators.
	//
	// example:
	//
	// ComputeNestPrometheus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetEntityId(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetMetricName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.MetricName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetPromql(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.Promql = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) SetType(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings struct {
	// The ID of the entity.
	//
	// example:
	//
	// cmgj0015****-Frequency-1
	EntityIds *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	// The package name.
	//
	// example:
	//
	// Pay-as-you-go package
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetEntityIds(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.EntityIds = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadata struct {
	// The configurations of the billable items.
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The billable items that are associated with the package.
	MeteringEntityMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	// The mappings between the service specifications and the template or package.
	SpecificationMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetMeteringEntityExtraInfos(v []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.MeteringEntityExtraInfos = v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetMeteringEntityMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.MeteringEntityMappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetSpecificationMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.SpecificationMappings = v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos struct {
	// The ID of the billable item.
	//
	// example:
	//
	// cmgjxxxxxxxx-NetworkOut-2
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metric name.
	//
	// example:
	//
	// NetworkLantency
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The custom prometheus statement.
	//
	// example:
	//
	// avg_over_time(count(kube_pod_info{namespace=\\"default\\"})[1h:1m])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// The metric.
	//
	// example:
	//
	// AvgPod
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetEntityId(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetMetricName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.MetricName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetPromql(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.Promql = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) SetType(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings struct {
	// The ID of the billable item.
	//
	// example:
	//
	// cmgjxxxxxxxx-NetworkOut-2
	EntityIds *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// Pay-as-you-go Package
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetEntityIds(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.EntityIds = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings struct {
	// The specification code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// Pay-as-you-go
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationCode(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.TemplateName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetTrialType(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.TrialType = &v
	return s
}

type GetServiceResponseBodyCommodityMeteringEntities struct {
	// The ID of the billable item.
	//
	// example:
	//
	// cmgjxxxxxxxx-NetworkOut
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the billable item.
	//
	// example:
	//
	// spring-boot-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetServiceResponseBodyCommodityMeteringEntities) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMeteringEntities) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetEntityId(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetName(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.Name = &v
	return s
}

type GetServiceResponseBodyCommoditySpecifications struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00xxxx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The specification name.
	//
	// example:
	//
	// specifications1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subscription duration. Unit: week or year.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommoditySpecifications) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommoditySpecifications) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetCode(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Code = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetName(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetTimes(v []*string) *GetServiceResponseBodyCommoditySpecifications {
	s.Times = v
	return s
}

type GetServiceResponseBodyComplianceMetadata struct {
	// The compliance package is selected.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyComplianceMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyComplianceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyComplianceMetadata) SetCompliancePacks(v []*string) *GetServiceResponseBodyComplianceMetadata {
	s.CompliancePacks = v
	return s
}

type GetServiceResponseBodyServiceDocumentInfos struct {
	// The URL that is used to access the document.
	//
	// example:
	//
	// http://docurl
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The language of the return data. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The template name.
	//
	// example:
	//
	// Default Template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyServiceDocumentInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceDocumentInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetDocumentUrl(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.DocumentUrl = &v
	return s
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetLocale(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetTemplateName(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyServiceInfos struct {
	// The agreement information about the service.
	Agreements []*GetServiceResponseBodyServiceInfosAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The URL of the detailed description of the service.
	//
	// example:
	//
	// https://example.com
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// The service name.
	//
	// example:
	//
	// WordPress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the information about the software in the service.
	Softwares []*GetServiceResponseBodyServiceInfosSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfos) SetAgreements(v []*GetServiceResponseBodyServiceInfosAgreements) *GetServiceResponseBodyServiceInfos {
	s.Agreements = v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetImage(v string) *GetServiceResponseBodyServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetLocale(v string) *GetServiceResponseBodyServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetLongDescriptionUrl(v string) *GetServiceResponseBodyServiceInfos {
	s.LongDescriptionUrl = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetName(v string) *GetServiceResponseBodyServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetShortDescription(v string) *GetServiceResponseBodyServiceInfos {
	s.ShortDescription = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetSoftwares(v []*GetServiceResponseBodyServiceInfosSoftwares) *GetServiceResponseBodyServiceInfos {
	s.Softwares = v
	return s
}

type GetServiceResponseBodyServiceInfosAgreements struct {
	// The agreement name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The agreement URL.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosAgreements) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosAgreements) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetName(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetUrl(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Url = &v
	return s
}

type GetServiceResponseBodyServiceInfosSoftwares struct {
	// The name of the software
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosSoftwares) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosSoftwares) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetName(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetVersion(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Version = &v
	return s
}

type GetServiceResponseBodyStatistic struct {
	// The total number of service instances that belong to the service. The service instances that are deleted are counted.
	//
	// example:
	//
	// 75
	AccumulativeInstanceCount *int32 `json:"AccumulativeInstanceCount,omitempty" xml:"AccumulativeInstanceCount,omitempty"`
	// The total amount consumed for trial service instances. Unit: CNY.
	//
	// example:
	//
	// 80.35
	AccumulativePocAmount *float64 `json:"AccumulativePocAmount,omitempty" xml:"AccumulativePocAmount,omitempty"`
	// The total number of users who use the service. The historical users are counted.
	//
	// example:
	//
	// 60
	AccumulativeUserCount *int32 `json:"AccumulativeUserCount,omitempty" xml:"AccumulativeUserCount,omitempty"`
	// The average amount consumed for trial service instances per instance. Unit: CNY.
	//
	// example:
	//
	// 40.17
	AveragePocAmount *float64 `json:"AveragePocAmount,omitempty" xml:"AveragePocAmount,omitempty"`
	// The average duration for which trial service instances are in use. Unit: Hour.
	//
	// example:
	//
	// 1
	AveragePocDuration *float64 `json:"AveragePocDuration,omitempty" xml:"AveragePocDuration,omitempty"`
	// The average amount consumed for trial service instances per a period of time. Unit: CNY.
	//
	// example:
	//
	// 167.9
	AveragePocUnitAmount *float64 `json:"AveragePocUnitAmount,omitempty" xml:"AveragePocUnitAmount,omitempty"`
	// The number of online service instances. It means the number of service instances that are successfully deployed.
	//
	// example:
	//
	// 20
	DeployedServiceInstanceCount *int32 `json:"DeployedServiceInstanceCount,omitempty" xml:"DeployedServiceInstanceCount,omitempty"`
	// The number of online users. It means the number of users who successfully deployed the service instances.
	//
	// example:
	//
	// 10
	DeployedUserCount *int32 `json:"DeployedUserCount,omitempty" xml:"DeployedUserCount,omitempty"`
	// The number of service applications that are in the Submitted state.
	//
	// example:
	//
	// 10
	SubmittedUsageCount *int32 `json:"SubmittedUsageCount,omitempty" xml:"SubmittedUsageCount,omitempty"`
}

func (s GetServiceResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativeInstanceCount(v int32) *GetServiceResponseBodyStatistic {
	s.AccumulativeInstanceCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativePocAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AccumulativePocAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAccumulativeUserCount(v int32) *GetServiceResponseBodyStatistic {
	s.AccumulativeUserCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocDuration(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocDuration = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetAveragePocUnitAmount(v float64) *GetServiceResponseBodyStatistic {
	s.AveragePocUnitAmount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetDeployedServiceInstanceCount(v int32) *GetServiceResponseBodyStatistic {
	s.DeployedServiceInstanceCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetDeployedUserCount(v int32) *GetServiceResponseBodyStatistic {
	s.DeployedUserCount = &v
	return s
}

func (s *GetServiceResponseBodyStatistic) SetSubmittedUsageCount(v int32) *GetServiceResponseBodyStatistic {
	s.SubmittedUsageCount = &v
	return s
}

type GetServiceResponseBodySupportContacts struct {
	// The type of Contact information.
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of Contact information.
	//
	// example:
	//
	// supplier@test.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodySupportContacts) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodySupportContacts) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodySupportContacts) SetType(v string) *GetServiceResponseBodySupportContacts {
	s.Type = &v
	return s
}

func (s *GetServiceResponseBodySupportContacts) SetValue(v string) *GetServiceResponseBodySupportContacts {
	s.Value = &v
	return s
}

type GetServiceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyTags) SetKey(v string) *GetServiceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceResponseBodyTags) SetValue(v string) *GetServiceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetServiceEstimateCostRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	Commodity *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-16fbd358d75e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequest) SetClientToken(v string) *GetServiceEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetParameters(v map[string]interface{}) *GetServiceEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetRegionId(v string) *GetServiceEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceVersion(v string) *GetServiceEstimateCostRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetSpecificationName(v string) *GetServiceEstimateCostRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetTemplateName(v string) *GetServiceEstimateCostRequest {
	s.TemplateName = &v
	return s
}

type GetServiceEstimateCostRequestCommodity struct {
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int32 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// 	- Day
	//
	// example:
	//
	// Month
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s GetServiceEstimateCostRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequestCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriod(v int32) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriodUnit(v string) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

type GetServiceEstimateCostShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-16fbd358d75e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceEstimateCostShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostShrinkRequest) SetClientToken(v string) *GetServiceEstimateCostShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.CommodityShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetParametersShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetRegionId(v string) *GetServiceEstimateCostShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceVersion(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetSpecificationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTemplateName(v string) *GetServiceEstimateCostShrinkRequest {
	s.TemplateName = &v
	return s
}

type GetServiceEstimateCostResponseBody struct {
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	//
	// example:
	//
	// {\\"PayPeriodUnit\\":Month,\\"PayPeriod\\":1}
	Commodity map[string]interface{} `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	//
	// example:
	//
	// {
	//
	//       "ECSInstances":{
	//
	//         "Type":"ALIYUN::ECS::InstanceGroup",
	//
	//         "Success":true,
	//
	//         "Result":{
	//
	//           "Order":{
	//
	//             "Currency":"CNY",
	//
	//             "RuleIds":[
	//
	//               1752723
	//
	//             ],
	//
	//             "DetailInfos":{
	//
	//               "ResourcePriceModel":[
	//
	//                 {
	//
	//                   "Resource":"bandwidth",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"image",
	//
	//                   "TradeAmount":0.0,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.0,
	//
	//                   "DiscountAmount":0.0
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"instanceType",
	//
	//                   "TradeAmount":0.006966,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.45,
	//
	//                   "DiscountAmount":0.443034
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"systemDisk",
	//
	//                   "TradeAmount":0.000867,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.056,
	//
	//                   "DiscountAmount":0.055133
	//
	//                 },
	//
	//                 {
	//
	//                   "Resource":"dataDisk",
	//
	//                   "TradeAmount":0.002167,
	//
	//                   "SubRuleIds":[],
	//
	//                   "OriginalAmount":0.14,
	//
	//                   "DiscountAmount":0.137833
	//
	//                 }
	//
	//               ]
	//
	//             }
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetServiceEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponseBody) SetCommodity(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetRequestId(v string) *GetServiceEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Resources = v
	return s
}

type GetServiceEstimateCostResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEstimateCostResponse) SetStatusCode(v int32) *GetServiceEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEstimateCostResponse) SetBody(v *GetServiceEstimateCostResponseBody) *GetServiceEstimateCostResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	// The business status of the service instance. Valid values:
	//
	// 	- Normal: The service instance is normal.
	//
	// 	- Renewing: The service instance is being renewed.
	//
	// 	- RenewFailed: The service instance failed to be renewed.
	//
	// 	- Expired: The service instance expired.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Indicates whether the Prometheus monitoring feature is enabled on the user side.
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The time when the service instance expires.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// https://g.xxxxxx.aliyun.com/d/xxxxxxxx-cn-mariadb/mysql-xxxxxxxx-quickstart-and-dashboard?orgId=355401&refresh=10s
	GrafanaDashBoardUrl *string `json:"GrafanaDashBoardUrl,omitempty" xml:"GrafanaDashBoardUrl,omitempty"`
	// Indicates whether the hosted O\\&M feature is enabled for the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {"renewType":"MONTHLY"}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network configurations.
	//
	// >  This parameter is discontinued.
	NetworkConfig *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// The ID of the service instance that is used to implement hosted O\\&M.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The end of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// Operate extra info
	//
	// example:
	//
	// 6cc5ce3590ed7f5529e85fec
	OperationExtraInfo *string `json:"OperationExtraInfo,omitempty" xml:"OperationExtraInfo,omitempty"`
	// The beginning of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The outputs returned from creating the service instance.
	//
	// 	- If the service is deployed by using a ROS template, all output fields of the template are returned.
	//
	// 	- If the service is deployed by calling an SPI operation, the output fields of the service provider and for the Compute Nest additional features are returned.
	//
	// example:
	//
	// {
	//
	//       "InstanceIds": [
	//
	//             "i-hp38ofxl0dsyfi7z****"
	//
	//       ]
	//
	// }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method of the service. Valid values:
	//
	// 	- Permanent: Once you purchase the service, you can use it permanently.
	//
	// 	- Subscription: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a subscription basis.
	//
	// 	- PayAsYouGo: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a pay-as-you-go basis.
	//
	// 	- CustomFixTime: You are charged for the service based on a custom duration fixed by the service provider.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package 1
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The logon URL for the accounts in the resource directory corresponding to the service instance.
	//
	// example:
	//
	// https://signin.aliyun.com/switchRole.htm?parentAlias=1307782318217247.onaliyun.com&roleName=ResourceDirectoryAccountAccessRole
	RdAccountLoginUrl *string `json:"RdAccountLoginUrl,omitempty" xml:"RdAccountLoginUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resources.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "StackId": "stack-xxx"
	//
	//       }
	//
	// ]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The information about the service to which the service instance belongs.
	Service *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source of the service instance. Valid values:
	//
	// 	- User: Compute Nest customer
	//
	// 	- Market: Alibaba Cloud Marketplace
	//
	// 	- Supplier: Compute Nest service provider
	//
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The deployment state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment state of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 1964460391538545
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The custom tags.
	Tags []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the service instance was updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// example:
	//
	// 1234567
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetBizStatus(v string) *GetServiceInstanceResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody {
	s.EnableUserPrometheus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody {
	s.GrafanaDashBoardUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseMetadata(v string) *GetServiceInstanceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationExtraInfo(v string) *GetServiceInstanceResponseBody {
	s.OperationExtraInfo = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody {
	s.PredefinedParameterName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRdAccountLoginUrl(v string) *GetServiceInstanceResponseBody {
	s.RdAccountLoginUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResourceGroupId(v string) *GetServiceInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	// The ID of the endpoint for the private connection.
	//
	// >  This parameter is discontinued.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the endpoint service for the private connection.
	//
	// >  This parameter is discontinued.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	// The information about private connections.
	PrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	// The information about the reverse private connection.
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	// The network configurations, which are mainly used for the private connection.
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// The ID of the endpoint for the private connection.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the endpoint service for the private connection.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	// The bandwidth limit for the private connection established based on the private network interconnection mode of Compute Nest.
	//
	// example:
	//
	// 1536Mbps
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of the endpoints for private connections.
	EndpointIps []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// The status of the Ingress endpoint. Valid values:
	//
	// 	- Ready: The Ingress endpoint is connected.
	//
	// 	- Pending: The Ingress endpoint is being connected.
	//
	// 	- Failed: The Ingress endpoint fails to be connected.
	//
	// 	- Deleted: The Ingress endpoint is deleted.
	//
	// 	- Deleting: The Ingress endpoint is being deleted.
	//
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// The status of the network service. Valid values:
	//
	// 	- Ready: The network service is connected.
	//
	// 	- Pending: The network service is being connected.
	//
	// 	- Failed: The network service fails to be connected.
	//
	// 	- Deleted: The network service is deleted.
	//
	// 	- Deleting: The network service is being deleted.
	//
	// example:
	//
	// Ready
	NetworkServiceStatus *string `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	// The names of the security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The names of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetConnectBandwidth(v int32) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetDomainName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.DomainName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	// The ID of the endpoint for the reverse private connection.
	//
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the endpoint service for the reverse private connection.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The URL of the service documentation.
	//
	// example:
	//
	// http://example.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-9c8a3522528b4fe8****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the service. Valid values:
	//
	// Draft: The registration request of the service is pending to be submitted.
	//
	// Submitted: The registration request of the service is submitted.
	//
	// Approved: The registration request of the service is approved.
	//
	// Online: The service is published.
	//
	// Offline: The service is unpublished.
	//
	// Deleted: The service is deleted.
	//
	// Launching: The service is being published.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The upgradable service Info.
	UpgradableServiceInfos []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos `json:"UpgradableServiceInfos,omitempty" xml:"UpgradableServiceInfos,omitempty" type:"Repeated"`
	// The service versions that can be updated.
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	// The service version.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// 20231127020225
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceInfos(v []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The service name.
	//
	// example:
	//
	// WordPress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponseBodyServiceUpgradableServiceInfos struct {
	// The upgradable service version.
	//
	// example:
	//
	// 4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name of an upgradable service version.
	//
	// example:
	//
	// Init version
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersion(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersionName(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type GetServiceProvisionsRequest struct {
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-20b8a396048346xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceProvisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequest) SetParameters(v map[string]interface{}) *GetServiceProvisionsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsRequest) SetRegionId(v string) *GetServiceProvisionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceId(v string) *GetServiceProvisionsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceVersion(v string) *GetServiceProvisionsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateName(v string) *GetServiceProvisionsRequest {
	s.TemplateName = &v
	return s
}

type GetServiceProvisionsShrinkRequest struct {
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-20b8a396048346xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceProvisionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsShrinkRequest) SetParametersShrink(v string) *GetServiceProvisionsShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetRegionId(v string) *GetServiceProvisionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceId(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceVersion(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetTemplateName(v string) *GetServiceProvisionsShrinkRequest {
	s.TemplateName = &v
	return s
}

type GetServiceProvisionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DB1FA13E-1087-5654-84D5-58A0ACAD1B18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the cloud services.
	ServiceProvisions []*GetServiceProvisionsResponseBodyServiceProvisions `json:"ServiceProvisions,omitempty" xml:"ServiceProvisions,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBody) SetRequestId(v string) *GetServiceProvisionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceProvisionsResponseBody) SetServiceProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisions) *GetServiceProvisionsResponseBody {
	s.ServiceProvisions = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisions struct {
	// Indicates whether automatic activation for the service is defined in the template. Valid values:
	//
	// 	- true: Automatic activation for the service is defined in the template.
	//
	// 	- false: Manual activation for the service is defined in the template.
	//
	// example:
	//
	// true
	AutoEnableService   *bool                                                                   `json:"AutoEnableService,omitempty" xml:"AutoEnableService,omitempty"`
	CommodityProvisions []*GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions `json:"CommodityProvisions,omitempty" xml:"CommodityProvisions,omitempty" type:"Repeated"`
	// The URL that points to the activation page of the service.
	//
	// > This parameter is returned if Status is set to Disabled.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?commodityCode=sls
	EnableURL *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	// The information about the RAM roles of the cloud service. If this parameter is empty, no RAM roles is associated with the service.
	RoleProvision *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision `json:"RoleProvision,omitempty" xml:"RoleProvision,omitempty" type:"Struct"`
	// The name of the cloud service.
	//
	// example:
	//
	// db
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The activation status of the cloud service. Valid values:
	//
	// - Enabled: The cloud service is activated.
	//
	// - EnabledByDefault: The cloud service is activated by default.
	//
	// - Disabled: The cloud service is not activated.
	//
	// - Unknown: The activation status of the cloud service is unknown.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the service is in the Disabled or Unknown state.
	//
	// > This parameter is returned if Status is set to Disabled or Unknown.
	//
	// example:
	//
	// resource(Eip) import info missing
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetAutoEnableService(v bool) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.AutoEnableService = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetCommodityProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.CommodityProvisions = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetRoleProvision(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.RoleProvision = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetServiceName(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.Status = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatusReason(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.StatusReason = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	EnableURL     *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetCommodityCode(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisionsCommodityProvisions {
	s.Status = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision struct {
	// The authorization URL of the RAM role.
	//
	// > This parameter is returned if Created is set to false.
	//
	// example:
	//
	// https://ram.console.aliyun.com/role/authorization?request={"Services":[{"Service":"CS","Roles":[{"RoleName":"AliyunCSManagedVKRole","TemplateId":"AliyunCSManagedVKRole"},{"RoleName":"AliyunCSDefaultRole","TemplateId":"Default"}]}],"ReturnUrl":"https://cs.console.aliyun.com/"}
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	// The RAM roles.
	Roles []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetAuthorizationURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.AuthorizationURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetRoles(v []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.Roles = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles struct {
	// The information about the API operation that is used to create the RAM role.
	ApiForCreation *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation `json:"ApiForCreation,omitempty" xml:"ApiForCreation,omitempty" type:"Struct"`
	// Indicates whether the RAM role is created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Created *bool `json:"Created,omitempty" xml:"Created,omitempty"`
	// The purpose for which the RAM role is used. Default value: Default. A value of Default indicates that the RAM role is the default role of the service.
	//
	// example:
	//
	// Default
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// AliyunServiceRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetApiForCreation(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.ApiForCreation = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetCreated(v bool) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Created = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetFunction(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Function = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetRoleName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.RoleName = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation struct {
	// The name of the API operation.
	//
	// example:
	//
	// GetServiceProvisions
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the Alibaba Cloud service to which the API operation belongs.
	//
	// example:
	//
	// ComputeNest
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The type of the API operation. Valid values:
	//
	// 	- Open: public
	//
	// 	- Inner: private
	//
	// example:
	//
	// Open
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// The parameters of the API operation. ${Variable name} indicates a dynamic parameter.
	//
	// example:
	//
	// { "ServiceLinkedRole": "AliyunServiceRoleForRdsPgsqlOnEcs", "RegionId": "${RegionId}" }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiProductId(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiProductId = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiType(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiType = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetParameters(v map[string]interface{}) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.Parameters = v
	return s
}

type GetServiceProvisionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceProvisionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceProvisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponse) SetHeaders(v map[string]*string) *GetServiceProvisionsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceProvisionsResponse) SetStatusCode(v int32) *GetServiceProvisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceProvisionsResponse) SetBody(v *GetServiceProvisionsResponseBody) *GetServiceProvisionsResponse {
	s.Body = v
	return s
}

type GetServiceRegistrationRequest struct {
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service registration ID.
	//
	// example:
	//
	// sr-1b4aabc1c9eb4109****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
}

func (s GetServiceRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationRequest) SetRegionId(v string) *GetServiceRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRegistrationRequest) SetRegistrationId(v string) *GetServiceRegistrationRequest {
	s.RegistrationId = &v
	return s
}

type GetServiceRegistrationResponseBody struct {
	// Comment from reviewer.
	//
	// example:
	//
	// comment message
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The details of service audit.
	Detail *GetServiceRegistrationResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// Finish time.
	//
	// example:
	//
	// 2024-12-07T11:05:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Service registration ID.
	//
	// example:
	//
	// sr-1b4aabc1c9eb4109****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c2d118c9193e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo *GetServiceRegistrationResponseBodyServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Struct"`
	// The service version.
	//
	// example:
	//
	// beta
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The status of service registration. Valid values:
	//
	// 	- Submitted
	//
	// 	- Approved
	//
	// 	- Rejected
	//
	// 	- Canceled
	//
	// 	- Executed
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Submit time.
	//
	// example:
	//
	// 2024-12-07T11:05:50Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s GetServiceRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBody) SetComment(v string) *GetServiceRegistrationResponseBody {
	s.Comment = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetDetail(v *GetServiceRegistrationResponseBodyDetail) *GetServiceRegistrationResponseBody {
	s.Detail = v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetFinishTime(v string) *GetServiceRegistrationResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetRegistrationId(v string) *GetServiceRegistrationResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetRequestId(v string) *GetServiceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceId(v string) *GetServiceRegistrationResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceInfo(v *GetServiceRegistrationResponseBodyServiceInfo) *GetServiceRegistrationResponseBody {
	s.ServiceInfo = v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetServiceVersion(v string) *GetServiceRegistrationResponseBody {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetStatus(v string) *GetServiceRegistrationResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceRegistrationResponseBody) SetSubmitTime(v string) *GetServiceRegistrationResponseBody {
	s.SubmitTime = &v
	return s
}

type GetServiceRegistrationResponseBodyDetail struct {
	// Whether risk exists.
	//
	// example:
	//
	// true
	AtRisk *bool `json:"AtRisk,omitempty" xml:"AtRisk,omitempty"`
	// Whether service is associated with artifact.
	//
	// example:
	//
	// true
	HasRelatedArtifact *bool `json:"HasRelatedArtifact,omitempty" xml:"HasRelatedArtifact,omitempty"`
	// The reports.
	//
	// example:
	//
	// { "template1":"https://compute-nest-security-audit-bucket.oss-cn-hangzhou.aliyuncs.com/report" }
	Reports *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	// The url of template diff file.
	//
	// example:
	//
	// https://compute-nest-template-diff-bucket.oss-cn-hangzhou.aliyuncs.com/service-abc/diff
	TemplateDiffUrl *string `json:"TemplateDiffUrl,omitempty" xml:"TemplateDiffUrl,omitempty"`
}

func (s GetServiceRegistrationResponseBodyDetail) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRegistrationResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBodyDetail) SetAtRisk(v bool) *GetServiceRegistrationResponseBodyDetail {
	s.AtRisk = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetHasRelatedArtifact(v bool) *GetServiceRegistrationResponseBodyDetail {
	s.HasRelatedArtifact = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetReports(v string) *GetServiceRegistrationResponseBodyDetail {
	s.Reports = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyDetail) SetTemplateDiffUrl(v string) *GetServiceRegistrationResponseBodyDetail {
	s.TemplateDiffUrl = &v
	return s
}

type GetServiceRegistrationResponseBodyServiceInfo struct {
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The version name.
	//
	// example:
	//
	// v1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceRegistrationResponseBodyServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRegistrationResponseBodyServiceInfo) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetServiceType(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.ServiceType = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetTrialType(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.TrialType = &v
	return s
}

func (s *GetServiceRegistrationResponseBodyServiceInfo) SetVersionName(v string) *GetServiceRegistrationResponseBodyServiceInfo {
	s.VersionName = &v
	return s
}

type GetServiceRegistrationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponse) SetHeaders(v map[string]*string) *GetServiceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *GetServiceRegistrationResponse) SetStatusCode(v int32) *GetServiceRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceRegistrationResponse) SetBody(v *GetServiceRegistrationResponseBody) *GetServiceRegistrationResponse {
	s.Body = v
	return s
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region in which the service instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// Specifies whether to enable the private connection feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// The parameters in the template.
	Parameters []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-1c11f365190c44xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// The parameter name.
	//
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// PostPaid
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBody struct {
	// The constraint families.
	FamilyConstraints []*string `json:"FamilyConstraints,omitempty" xml:"FamilyConstraints,omitempty" type:"Repeated"`
	// The parameters in the template.
	ParameterConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C81C0732-DEBC-559C-B563-7EB2BEB21088
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.FamilyConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The names of the associated parameters.
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// The behavior of the parameter. Valid values:
	//
	// 	- NoLimit: The value of this parameter is not limited.
	//
	// 	- NotSupport: The value of this parameter cannot be queried.
	//
	// 	- QueryError: The query failed.
	//
	// >  If AllowedValues is not returned, Behavior and BehaviorReason are returned.
	//
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason why the behavior of the parameter is returned.
	//
	// example:
	//
	// none
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	// The original constraint information.
	OriginalConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// The name of the parameter.
	//
	// example:
	//
	// PayType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The property name.
	//
	// example:
	//
	// lnch_Source
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The resource name.
	//
	// example:
	//
	// i-8vb0smn1lf6g77md****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// serviceinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

type GetServiceTestTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetServiceTestTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTestTaskRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskRequest) SetRegionId(v string) *GetServiceTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTestTaskRequest) SetTaskId(v string) *GetServiceTestTaskRequest {
	s.TaskId = &v
	return s
}

type GetServiceTestTaskResponseBody struct {
	// The execution details.
	ExecutionDetails []*GetServiceTestTaskResponseBodyExecutionDetails `json:"ExecutionDetails,omitempty" xml:"ExecutionDetails,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service test task. Valid values:
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Failure
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// nametest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task execution region.
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
}

func (s GetServiceTestTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponseBody) SetExecutionDetails(v []*GetServiceTestTaskResponseBodyExecutionDetails) *GetServiceTestTaskResponseBody {
	s.ExecutionDetails = v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetRequestId(v string) *GetServiceTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetStatus(v string) *GetServiceTestTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetTaskName(v string) *GetServiceTestTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetServiceTestTaskResponseBody) SetTaskRegionId(v string) *GetServiceTestTaskResponseBody {
	s.TaskRegionId = &v
	return s
}

type GetServiceTestTaskResponseBodyExecutionDetails struct {
	// The service test case name.
	//
	// example:
	//
	// case1
	CaseName *string `json:"CaseName,omitempty" xml:"CaseName,omitempty"`
	// The execution report
	//
	// example:
	//
	// -----------------------------------------------------------------------------
	//
	// Region: cn-qingdao
	//
	// StackName: iact3-default-cn-qingd
	//
	// StackId: 009d2991-f494-d
	//
	// *****************************************************************************
	ExecutionReport *string `json:"ExecutionReport,omitempty" xml:"ExecutionReport,omitempty"`
	// The sub task status.
	//
	// example:
	//
	// Runing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub task id.
	//
	// example:
	//
	// stt-xxxx
	SubTaskId *string `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
}

func (s GetServiceTestTaskResponseBodyExecutionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTestTaskResponseBodyExecutionDetails) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetCaseName(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.CaseName = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetExecutionReport(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.ExecutionReport = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetStatus(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.Status = &v
	return s
}

func (s *GetServiceTestTaskResponseBodyExecutionDetails) SetSubTaskId(v string) *GetServiceTestTaskResponseBodyExecutionDetails {
	s.SubTaskId = &v
	return s
}

type GetServiceTestTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTestTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTestTaskResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponse) SetHeaders(v map[string]*string) *GetServiceTestTaskResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTestTaskResponse) SetStatusCode(v int32) *GetServiceTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTestTaskResponse) SetBody(v *GetServiceTestTaskResponseBody) *GetServiceTestTaskResponse {
	s.Body = v
	return s
}

type GetSupplierInformationRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupplierInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationRequest) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationRequest) SetRegionId(v string) *GetSupplierInformationRequest {
	s.RegionId = &v
	return s
}

type GetSupplierInformationResponseBody struct {
	// Acr container namespace
	//
	// example:
	//
	// computenest
	AcrNamespace *string `json:"AcrNamespace,omitempty" xml:"AcrNamespace,omitempty"`
	// The delivery settings.
	DeliverySettings *GetSupplierInformationResponseBodyDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// Whether to enable reseller
	//
	// example:
	//
	// true
	EnableReseller *bool `json:"EnableReseller,omitempty" xml:"EnableReseller,omitempty"`
	// The Ip of the operation.
	//
	// example:
	//
	// 10.xxx.xxx.xxx/101
	OperationIp *string `json:"OperationIp,omitempty" xml:"OperationIp,omitempty"`
	// The MFA of the operation.
	//
	// example:
	//
	// true
	OperationMfaPresent *bool `json:"OperationMfaPresent,omitempty" xml:"OperationMfaPresent,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/34DB-4F4C-9373-003AA060****.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*GetSupplierInformationResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
}

func (s GetSupplierInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBody) SetAcrNamespace(v string) *GetSupplierInformationResponseBody {
	s.AcrNamespace = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetDeliverySettings(v *GetSupplierInformationResponseBodyDeliverySettings) *GetSupplierInformationResponseBody {
	s.DeliverySettings = v
	return s
}

func (s *GetSupplierInformationResponseBody) SetEnableReseller(v bool) *GetSupplierInformationResponseBody {
	s.EnableReseller = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetOperationIp(v string) *GetSupplierInformationResponseBody {
	s.OperationIp = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetOperationMfaPresent(v bool) *GetSupplierInformationResponseBody {
	s.OperationMfaPresent = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetRequestId(v string) *GetSupplierInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierDesc(v string) *GetSupplierInformationResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierLogo(v string) *GetSupplierInformationResponseBody {
	s.SupplierLogo = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierName(v string) *GetSupplierInformationResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupplierUrl(v string) *GetSupplierInformationResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetSupplierInformationResponseBody) SetSupportContacts(v []*GetSupplierInformationResponseBodySupportContacts) *GetSupplierInformationResponseBody {
	s.SupportContacts = v
	return s
}

type GetSupplierInformationResponseBodyDeliverySettings struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// mybucket
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Indicates whether screencast delivery to Object Storage Service (OSS) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// path1/path2/
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s GetSupplierInformationResponseBodyDeliverySettings) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponseBodyDeliverySettings) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssBucketName(v string) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssEnabled(v bool) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssExpirationDays(v int64) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *GetSupplierInformationResponseBodyDeliverySettings) SetOssPath(v string) *GetSupplierInformationResponseBodyDeliverySettings {
	s.OssPath = &v
	return s
}

type GetSupplierInformationResponseBodySupportContacts struct {
	// The type of contact information.
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of contact information.
	//
	// example:
	//
	// supplier@example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSupplierInformationResponseBodySupportContacts) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponseBodySupportContacts) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponseBodySupportContacts) SetType(v string) *GetSupplierInformationResponseBodySupportContacts {
	s.Type = &v
	return s
}

func (s *GetSupplierInformationResponseBodySupportContacts) SetValue(v string) *GetSupplierInformationResponseBodySupportContacts {
	s.Value = &v
	return s
}

type GetSupplierInformationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplierInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplierInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupplierInformationResponse) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationResponse) SetHeaders(v map[string]*string) *GetSupplierInformationResponse {
	s.Headers = v
	return s
}

func (s *GetSupplierInformationResponse) SetStatusCode(v int32) *GetSupplierInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplierInformationResponse) SetBody(v *GetSupplierInformationResponseBody) *GetSupplierInformationResponse {
	s.Body = v
	return s
}

type GetUploadCredentialsRequest struct {
	// The file name to upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// template.yaml
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Specifies whether the file is publicly accessible. Valid values: **public*	- or **private**. The default value is **private**.
	//
	// example:
	//
	// public
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetUploadCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsRequest) SetFileName(v string) *GetUploadCredentialsRequest {
	s.FileName = &v
	return s
}

func (s *GetUploadCredentialsRequest) SetVisibility(v string) *GetUploadCredentialsRequest {
	s.Visibility = &v
	return s
}

type GetUploadCredentialsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetUploadCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FCC3321E-D518-1BC4-861E-588E9D4DAFB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates the request was successful. A value of false indicates the request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBody) SetCode(v string) *GetUploadCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetData(v *GetUploadCredentialsResponseBodyData) *GetUploadCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetHttpStatusCode(v int32) *GetUploadCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetMessage(v string) *GetUploadCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetRequestId(v string) *GetUploadCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetSuccess(v bool) *GetUploadCredentialsResponseBody {
	s.Success = &v
	return s
}

type GetUploadCredentialsResponseBodyData struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NUCe19W1FKaHAYAhe********
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// 8LQGp59mY23pcXeTdcvSA1cUQZBeD92sFrXi********
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// service-info-private
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The time when the AccessKey pair expires.
	//
	// example:
	//
	// 2023-05-18T12:27:59Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// 221514575922756034/cn-hangzhou/d57c62fbd508xxxxxxxx.json
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAISzQN1q6Ft5B2yfSjIr5b2LouNuu5n/KOjQ3/wjGUHYdlagYGdmzz2IH1Le3NrBO8esfgymGFU6v8dlo1dYLQeHhadQI5cs80HtFqLSNaE65LswPlZ2M2ISETPJzfV9pCK
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetUploadCredentialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeyId(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeySecret(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetBucketName(v string) *GetUploadCredentialsResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetExpireDate(v string) *GetUploadCredentialsResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetKey(v string) *GetUploadCredentialsResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetRegionId(v string) *GetUploadCredentialsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetSecurityToken(v string) *GetUploadCredentialsResponseBodyData {
	s.SecurityToken = &v
	return s
}

type GetUploadCredentialsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponse) SetHeaders(v map[string]*string) *GetUploadCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadCredentialsResponse) SetStatusCode(v int32) *GetUploadCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponse) SetBody(v *GetUploadCredentialsResponseBody) *GetUploadCredentialsResponse {
	s.Body = v
	return s
}

type LaunchServiceRequest struct {
	// The categories of the service.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Whether to set the recommended service publishing to the service directory.
	//
	// example:
	//
	// false
	Recommend *bool `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s LaunchServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceRequest) GoString() string {
	return s.String()
}

func (s *LaunchServiceRequest) SetCategories(v []*string) *LaunchServiceRequest {
	s.Categories = v
	return s
}

func (s *LaunchServiceRequest) SetClientToken(v string) *LaunchServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *LaunchServiceRequest) SetRecommend(v bool) *LaunchServiceRequest {
	s.Recommend = &v
	return s
}

func (s *LaunchServiceRequest) SetRegionId(v string) *LaunchServiceRequest {
	s.RegionId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceId(v string) *LaunchServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *LaunchServiceRequest) SetServiceVersion(v string) *LaunchServiceRequest {
	s.ServiceVersion = &v
	return s
}

type LaunchServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The mode of the service online. Valid Type
	//
	// - PublishNewVersion: Launch new version
	//
	// - PublishOfflineVersion:  The offline version is online again.
	//
	// - UpdateLatestVersion: Update the latest version online
	//
	// example:
	//
	// PublishNewVersion
	ServiceLaunchResultType *string `json:"ServiceLaunchResultType,omitempty" xml:"ServiceLaunchResultType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s LaunchServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponseBody) SetRequestId(v string) *LaunchServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchServiceResponseBody) SetServiceLaunchResultType(v string) *LaunchServiceResponseBody {
	s.ServiceLaunchResultType = &v
	return s
}

func (s *LaunchServiceResponseBody) SetVersion(v string) *LaunchServiceResponseBody {
	s.Version = &v
	return s
}

type LaunchServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchServiceResponse) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponse) SetHeaders(v map[string]*string) *LaunchServiceResponse {
	s.Headers = v
	return s
}

func (s *LaunchServiceResponse) SetStatusCode(v int32) *LaunchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchServiceResponse) SetBody(v *LaunchServiceResponseBody) *LaunchServiceResponse {
	s.Body = v
	return s
}

type ListAcrImageRepositoriesRequest struct {
	// The type of the artifact. Default value: AcrImage. Valid values:
	//
	// 	- HelmChart: Helm chart image.
	//
	// 	- AcrImage: container image.
	//
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s ListAcrImageRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesRequest) SetArtifactType(v string) *ListAcrImageRepositoriesRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetMaxResults(v int32) *ListAcrImageRepositoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetNextToken(v string) *ListAcrImageRepositoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesRequest) SetRepoName(v string) *ListAcrImageRepositoriesRequest {
	s.RepoName = &v
	return s
}

type ListAcrImageRepositoriesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The images.
	Repositories []*ListAcrImageRepositoriesResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C4A145D8-6F6C-532A-9001-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBody) SetMaxResults(v int32) *ListAcrImageRepositoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetNextToken(v string) *ListAcrImageRepositoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRepositories(v []*ListAcrImageRepositoriesResponseBodyRepositories) *ListAcrImageRepositoriesResponseBody {
	s.Repositories = v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetRequestId(v string) *ListAcrImageRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBody) SetTotalCount(v int32) *ListAcrImageRepositoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAcrImageRepositoriesResponseBodyRepositories struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the image was modified.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The namespace of the repository
	//
	// example:
	//
	// computenest
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The image repo ID.
	//
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The image repo name.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `Private`: a private repository
	//
	// 	- `Public`: a public repository
	//
	// example:
	//
	// Private
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetCreateTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetModifiedTime(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetNamespace(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.Namespace = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoId(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoName(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListAcrImageRepositoriesResponseBodyRepositories) SetRepoType(v string) *ListAcrImageRepositoriesResponseBodyRepositories {
	s.RepoType = &v
	return s
}

type ListAcrImageRepositoriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponse) SetHeaders(v map[string]*string) *ListAcrImageRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetStatusCode(v int32) *ListAcrImageRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetBody(v *ListAcrImageRepositoriesResponseBody) *ListAcrImageRepositoriesResponse {
	s.Body = v
	return s
}

type ListAcrImageTagsRequest struct {
	// The type of the deployment package. Default value: AcrImage. Valid values:
	//
	// 	- HelmChart: Helm chart image.
	//
	// 	- AcrImage: container image.
	//
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The image ID.
	//
	// example:
	//
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListAcrImageTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsRequest) SetArtifactType(v string) *ListAcrImageTagsRequest {
	s.ArtifactType = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetMaxResults(v int32) *ListAcrImageTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetNextToken(v string) *ListAcrImageTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsRequest) SetRepoId(v string) *ListAcrImageTagsRequest {
	s.RepoId = &v
	return s
}

type ListAcrImageTagsResponseBody struct {
	// The list of images.
	Images []*ListAcrImageTagsResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ey14..
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEF343B9-1A15-5789-BE88-7B36190F5BF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBody) SetImages(v []*ListAcrImageTagsResponseBodyImages) *ListAcrImageTagsResponseBody {
	s.Images = v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetMaxResults(v int32) *ListAcrImageTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetNextToken(v string) *ListAcrImageTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetRequestId(v string) *ListAcrImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetTotalCount(v int32) *ListAcrImageTagsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAcrImageTagsResponseBodyImages struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The image size. Unit: bytes.
	//
	// example:
	//
	// 188394616
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was modified.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The image version.
	//
	// example:
	//
	// 5.7.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAcrImageTagsResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBodyImages) SetCreateTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetImageSize(v string) *ListAcrImageTagsResponseBodyImages {
	s.ImageSize = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetModifiedTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetTag(v string) *ListAcrImageTagsResponseBodyImages {
	s.Tag = &v
	return s
}

type ListAcrImageTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAcrImageTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponse) SetHeaders(v map[string]*string) *ListAcrImageTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageTagsResponse) SetStatusCode(v int32) *ListAcrImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageTagsResponse) SetBody(v *ListAcrImageTagsResponseBody) *ListAcrImageTagsResponse {
	s.Body = v
	return s
}

type ListArtifactRisksRequest struct {
	// Artifact ID.
	//
	// example:
	//
	// artifact-3fd95cdfdf0d4b1fa00c
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// Artifact version.
	//
	// example:
	//
	// 1
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s ListArtifactRisksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactRisksRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksRequest) SetArtifactId(v string) *ListArtifactRisksRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactRisksRequest) SetArtifactVersion(v string) *ListArtifactRisksRequest {
	s.ArtifactVersion = &v
	return s
}

type ListArtifactRisksResponseBody struct {
	// List of artifact risks
	ArtifactRiskList []*ListArtifactRisksResponseBodyArtifactRiskList `json:"ArtifactRiskList,omitempty" xml:"ArtifactRiskList,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 52919DB1-03A0-55F5-BDD4-DB6DEBB8267A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListArtifactRisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactRisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponseBody) SetArtifactRiskList(v []*ListArtifactRisksResponseBodyArtifactRiskList) *ListArtifactRisksResponseBody {
	s.ArtifactRiskList = v
	return s
}

func (s *ListArtifactRisksResponseBody) SetRequestId(v string) *ListArtifactRisksResponseBody {
	s.RequestId = &v
	return s
}

type ListArtifactRisksResponseBodyArtifactRiskList struct {
	// CVE numbers
	//
	// example:
	//
	// CVE-2023-4911
	CveNos *string `json:"CveNos,omitempty" xml:"CveNos,omitempty"`
	// Extended information, in JSON format, to be parsed according to the risk category
	//
	// example:
	//
	// {
	//
	//   "feature": "ntpdate",
	//
	//   "version": "4.2.6",
	//
	//   "cveLocation": "/usr/lib"
	//
	// }
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Risk level:
	//
	// - high represents high
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Risk name.
	//
	// example:
	//
	// USN-3686-1: file vulnerabilities
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// Risk type. Values:
	//
	// - AcrCve  Container image system vulnerability
	//
	// - AcrSca  Container image application vulnerability
	//
	// - EcsVulnerability  ECS image vulnerability information
	//
	// - EcsAlarm  ECS image security alarm
	//
	// - EcsBaseline  ECS image baseline check
	//
	// example:
	//
	// AcrCve
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// Risk Type name
	//
	// example:
	//
	// Container System Vulner
	RiskTypeName *string `json:"RiskTypeName,omitempty" xml:"RiskTypeName,omitempty"`
	// Solution for the risk item.
	//
	// example:
	//
	// apt-get update && apt-get install ntpdate --only-upgrade
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s ListArtifactRisksResponseBodyArtifactRiskList) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactRisksResponseBodyArtifactRiskList) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetCveNos(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.CveNos = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetExtendInfo(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.ExtendInfo = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetLevel(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.Level = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskName(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskName = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskType(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskType = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskTypeName(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskTypeName = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetSolution(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.Solution = &v
	return s
}

type ListArtifactRisksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactRisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactRisksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactRisksResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponse) SetHeaders(v map[string]*string) *ListArtifactRisksResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactRisksResponse) SetStatusCode(v int32) *ListArtifactRisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactRisksResponse) SetBody(v *ListArtifactRisksResponseBody) *ListArtifactRisksResponse {
	s.Body = v
	return s
}

type ListArtifactVersionsRequest struct {
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The filter.
	Filters []*ListArtifactVersionsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequest) SetArtifactId(v string) *ListArtifactVersionsRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetFilters(v []*ListArtifactVersionsRequestFilters) *ListArtifactVersionsRequest {
	s.Filters = v
	return s
}

func (s *ListArtifactVersionsRequest) SetMaxResults(v int32) *ListArtifactVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetNextToken(v string) *ListArtifactVersionsRequest {
	s.NextToken = &v
	return s
}

type ListArtifactVersionsRequestFilters struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// **Status**：The artifact status
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactVersionsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequestFilters) SetName(v string) *ListArtifactVersionsRequestFilters {
	s.Name = &v
	return s
}

func (s *ListArtifactVersionsRequestFilters) SetValues(v []*string) *ListArtifactVersionsRequestFilters {
	s.Values = v
	return s
}

type ListArtifactVersionsShrinkRequest struct {
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The filter.
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactVersionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsShrinkRequest) SetArtifactId(v string) *ListArtifactVersionsShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetFiltersShrink(v string) *ListArtifactVersionsShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetMaxResults(v int32) *ListArtifactVersionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsShrinkRequest) SetNextToken(v string) *ListArtifactVersionsShrinkRequest {
	s.NextToken = &v
	return s
}

type ListArtifactVersionsResponseBody struct {
	// The version information about the deployment package.
	Artifacts []*ListArtifactVersionsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBody) SetArtifacts(v []*ListArtifactVersionsResponseBodyArtifacts) *ListArtifactVersionsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetMaxResults(v int32) *ListArtifactVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetNextToken(v string) *ListArtifactVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetRequestId(v string) *ListArtifactVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetTotalCount(v int32) *ListArtifactVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactVersionsResponseBodyArtifacts struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2022-10-20T02:19:53Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The distribution result of the image.
	ImageDelivery map[string]*string `json:"ImageDelivery,omitempty" xml:"ImageDelivery,omitempty"`
	// The distribution progress of the deployment package.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result file of the security scan.
	//
	// example:
	//
	// https://compute-nest-security-audit-bucket-ap-southeast-1.oss-ap-southeast-1.aliyuncs.com/51416747xxxx/xxxx
	ResultFile *string `json:"ResultFile,omitempty" xml:"ResultFile,omitempty"`
	// The result of the security scan. Valid values:
	//
	// 	- Normal: No risks exist on the deployment package.
	//
	// 	- AtRisk: Risks exist on the deployment package.
	//
	// 	- Processing: The deployment package is being scanned.
	//
	// example:
	//
	// Normal
	SecurityAuditResult *string `json:"SecurityAuditResult,omitempty" xml:"SecurityAuditResult,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- Created: The deployment package is created.
	//
	// 	- Scanning: The deployment package is being scanned.
	//
	// 	- ScanFailed: The deployment package failed to be scanned.
	//
	// 	- Delivering: The deployment package is being distributed.
	//
	// 	- Available: The deployment package is available.
	//
	// 	- Deleted: The deployment package is deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// [
	//
	// 					"cn-beijing",
	//
	// 					"cn-hangzhou",
	//
	// 					"cn-shanghai"
	//
	// 				]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListArtifactVersionsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactBuildProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactBuildType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactBuildType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactVersion(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtCreate(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtCreate = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetImageDelivery(v map[string]*string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ImageDelivery = v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetProgress(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Progress = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetResultFile(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ResultFile = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSecurityAuditResult(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SecurityAuditResult = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatus(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatusDetail(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.StatusDetail = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSupportRegionIds(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SupportRegionIds = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetVersionName(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.VersionName = &v
	return s
}

type ListArtifactVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponse) SetHeaders(v map[string]*string) *ListArtifactVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactVersionsResponse) SetStatusCode(v int32) *ListArtifactVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactVersionsResponse) SetBody(v *ListArtifactVersionsResponseBody) *ListArtifactVersionsResponse {
	s.Body = v
	return s
}

type ListArtifactsRequest struct {
	// The filter.
	Filter []*ListArtifactsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25yfxxxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListArtifactsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) SetFilter(v []*ListArtifactsRequestFilter) *ListArtifactsRequest {
	s.Filter = v
	return s
}

func (s *ListArtifactsRequest) SetMaxResults(v int32) *ListArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsRequest) SetNextToken(v string) *ListArtifactsRequest {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsRequest) SetResourceGroupId(v string) *ListArtifactsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsRequest) SetTag(v []*ListArtifactsRequestTag) *ListArtifactsRequest {
	s.Tag = v
	return s
}

type ListArtifactsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- *Name*: The name of the deployment package. Fuzzy match is used.
	//
	// 	- ArtifactId: The ID of the deployment package.
	//
	// 	- ArtifactType: The type of the deployment package.
	//
	// example:
	//
	// ArtifactType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestFilter) SetName(v string) *ListArtifactsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListArtifactsRequestFilter) SetValues(v []*string) *ListArtifactsRequestFilter {
	s.Values = v
	return s
}

type ListArtifactsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListArtifactsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequestTag) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestTag) SetKey(v string) *ListArtifactsRequestTag {
	s.Key = &v
	return s
}

func (s *ListArtifactsRequestTag) SetValue(v string) *ListArtifactsRequestTag {
	s.Value = &v
	return s
}

type ListArtifactsResponseBody struct {
	// The information about deployment packages.
	Artifacts []*ListArtifactsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactsResponseBody) SetMaxResults(v int32) *ListArtifactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsResponseBody) SetNextToken(v string) *ListArtifactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsResponseBody) SetRequestId(v string) *ListArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactsResponseBody) SetTotalCount(v int32) *ListArtifactsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactsResponseBodyArtifacts struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// {\\"CodeRepo\\":{\\"Owner\\":\\"wenle\\",\\"Platform\\":\\"github\\",\\"Branch\\":\\"main\\",\\"RepoName\\":\\"aliyun-computenest/java-springboot-demo\\"}}
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version of the deployment package.
	//
	// example:
	//
	// 2
	MaxVersion *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact They can only change from Automatic to Public. Options:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25rexxxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- Created: The deployment package is created.
	//
	// 	- Scanning: The deployment package is being scanned.
	//
	// 	- ScanFailed: The deployment package failed to be scanned.
	//
	// 	- Delivering: The deployment package is being distributed.
	//
	// 	- Available: The deployment package is available.
	//
	// 	- Deleted: The deployment package is deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListArtifactsResponseBodyArtifactsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListArtifactsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactBuildProperty(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetDescription(v string) *ListArtifactsResponseBodyArtifacts {
	s.Description = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetMaxVersion(v string) *ListArtifactsResponseBodyArtifacts {
	s.MaxVersion = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetName(v string) *ListArtifactsResponseBodyArtifacts {
	s.Name = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetPermissionType(v string) *ListArtifactsResponseBodyArtifacts {
	s.PermissionType = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetResourceGroupId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetStatus(v string) *ListArtifactsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetTags(v []*ListArtifactsResponseBodyArtifactsTags) *ListArtifactsResponseBodyArtifacts {
	s.Tags = v
	return s
}

type ListArtifactsResponseBodyArtifactsTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListArtifactsResponseBodyArtifactsTags) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifactsTags) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetKey(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Key = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetValue(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Value = &v
	return s
}

type ListArtifactsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponse) SetHeaders(v map[string]*string) *ListArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactsResponse) SetStatusCode(v int32) *ListArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactsResponse) SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse {
	s.Body = v
	return s
}

type ListResellersRequest struct {
	// The filters.
	Filter []*ListResellersRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAWVmrOoWHbw/80lX0TWxe/s=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListResellersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResellersRequest) GoString() string {
	return s.String()
}

func (s *ListResellersRequest) SetFilter(v []*ListResellersRequestFilter) *ListResellersRequest {
	s.Filter = v
	return s
}

func (s *ListResellersRequest) SetMaxResults(v int32) *ListResellersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResellersRequest) SetNextToken(v string) *ListResellersRequest {
	s.NextToken = &v
	return s
}

func (s *ListResellersRequest) SetRegionId(v string) *ListResellersRequest {
	s.RegionId = &v
	return s
}

type ListResellersRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- ResellerUid: the uid of the distributor.
	//
	// 	- Name: the name of the distributor.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Filter value array.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListResellersRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListResellersRequestFilter) GoString() string {
	return s.String()
}

func (s *ListResellersRequestFilter) SetName(v string) *ListResellersRequestFilter {
	s.Name = &v
	return s
}

func (s *ListResellersRequestFilter) SetValue(v []*string) *ListResellersRequestFilter {
	s.Value = v
	return s
}

type ListResellersResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// distributor informations
	SupplierInformation []*ListResellersResponseBodySupplierInformation `json:"SupplierInformation,omitempty" xml:"SupplierInformation,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResellersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResellersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResellersResponseBody) SetMaxResults(v int32) *ListResellersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResellersResponseBody) SetNextToken(v string) *ListResellersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResellersResponseBody) SetRequestId(v string) *ListResellersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResellersResponseBody) SetSupplierInformation(v []*ListResellersResponseBodySupplierInformation) *ListResellersResponseBody {
	s.SupplierInformation = v
	return s
}

func (s *ListResellersResponseBody) SetTotalCount(v int32) *ListResellersResponseBody {
	s.TotalCount = &v
	return s
}

type ListResellersResponseBodySupplierInformation struct {
	// The description of distributor.
	//
	// example:
	//
	// It is a XXXX  company
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of distributor
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/31978070/service-image/d5c3b585-ff6b-4e4e-8885-xxxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the distributor
	//
	// example:
	//
	// Distributor A
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The Alibaba Cloud account ID of the distributor.
	//
	// example:
	//
	// 152xxxxxxxxxxx
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the distributor.
	//
	// example:
	//
	// http://www.aliyun.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s ListResellersResponseBodySupplierInformation) String() string {
	return tea.Prettify(s)
}

func (s ListResellersResponseBodySupplierInformation) GoString() string {
	return s.String()
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierDesc(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierDesc = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierLogo(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierLogo = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierName(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierName = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierUid(v int64) *ListResellersResponseBodySupplierInformation {
	s.SupplierUid = &v
	return s
}

func (s *ListResellersResponseBodySupplierInformation) SetSupplierUrl(v string) *ListResellersResponseBodySupplierInformation {
	s.SupplierUrl = &v
	return s
}

type ListResellersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResellersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResellersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResellersResponse) GoString() string {
	return s.String()
}

func (s *ListResellersResponse) SetHeaders(v map[string]*string) *ListResellersResponse {
	s.Headers = v
	return s
}

func (s *ListResellersResponse) SetStatusCode(v int32) *ListResellersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResellersResponse) SetBody(v *ListResellersResponseBody) *ListResellersResponse {
	s.Body = v
	return s
}

type ListServiceInstanceBillRequest struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2024-12-05
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills by month. The data queried is consistent with the data that is displayed for the specified billing cycle on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- DAILY: queries bills by day. The data queried is consistent with the data that is displayed for the specified day on the Billing Details tab of the Bill Details page in User Center.
	//
	// You must set the **BillingDate*	- parameter before you can set the Granularity parameter to DAILY.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-6121296da4f44e469519
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-0d0d7bc9accc4e2e8a8f
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListServiceInstanceBillRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillRequest) SetBillingCycle(v string) *ListServiceInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetBillingDate(v string) *ListServiceInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetGranularity(v string) *ListServiceInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetMaxResults(v int32) *ListServiceInstanceBillRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetNextToken(v string) *ListServiceInstanceBillRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceInstanceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceVersion(v string) *ListServiceInstanceBillRequest {
	s.ServiceVersion = &v
	return s
}

type ListServiceInstanceBillResponseBody struct {
	// The billing information of the backup schedule.
	Item []*ListServiceInstanceBillResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B68C6D1-0C22-53F7-997C-F95D109C31AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 65
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstanceBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBody) SetItem(v []*ListServiceInstanceBillResponseBodyItem) *ListServiceInstanceBillResponseBody {
	s.Item = v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetMaxResults(v int32) *ListServiceInstanceBillResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetNextToken(v string) *ListServiceInstanceBillResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetRequestId(v string) *ListServiceInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetTotalCount(v int64) *ListServiceInstanceBillResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstanceBillResponseBodyItem struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2024-12
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2024-12-05
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billable item.
	//
	// example:
	//
	// Bandwidth
	BillingItem *string `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	// The code of the billable item.
	//
	// example:
	//
	// disk
	BillingItemCode *string `json:"BillingItemCode,omitempty" xml:"BillingItemCode,omitempty"`
	// The currency unit.
	//
	// 	- China site: **CNY**.
	//
	// 	- International site: **USD**.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount deducted with resource plans.
	//
	// example:
	//
	// 0
	DeductedByResourcePackage *string `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-wz9k051l23j3o3xa9
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	InvoiceDiscount *string `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The unit price.
	//
	// example:
	//
	// 0.12
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The unit of the unit price.
	//
	// example:
	//
	// CNY/GB
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 0
	PretaxAmount *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The details of the service.
	//
	// example:
	//
	// sls
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The name of the cloud service or the name of the service-linked role with which the cloud service is associated.
	//
	// example:
	//
	// oss
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Service instance ID.
	//
	// example:
	//
	// si-bc5d6ac7022647d3a3bd
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The billing cycle in which the bill is split.
	//
	// example:
	//
	// 2021-06
	SplitBillingCycle *string `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The amount of resource usage.
	//
	// example:
	//
	// 1
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The unit of usage.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
}

func (s ListServiceInstanceBillResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponseBodyItem) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingDate(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItem(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItem = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItemCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItemCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetCurrency(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Currency = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetDeductedByResourcePackage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInstanceID(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InstanceID = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInvoiceDiscount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPrice(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPrice = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPriceUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPriceUnit = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxGrossAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductDetail(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductDetail = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductName(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductName = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetServiceInstanceId(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSplitBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SplitBillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSubscriptionType(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SubscriptionType = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Usage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsageUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.UsageUnit = &v
	return s
}

type ListServiceInstanceBillResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceBillResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponse) SetHeaders(v map[string]*string) *ListServiceInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceBillResponse) SetStatusCode(v int32) *ListServiceInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceBillResponse) SetBody(v *ListServiceInstanceBillResponseBody) *ListServiceInstanceBillResponse {
	s.Body = v
	return s
}

type ListServiceInstanceDeployDetailsRequest struct {
	// The time zone.
	//
	// Reference Format: "+08:00"
	//
	// Valid Range: "-12:59" to "+13:00"
	//
	// example:
	//
	// +08:00
	CycleTimeZone *string `json:"CycleTimeZone,omitempty" xml:"CycleTimeZone,omitempty"`
	// Determines the time period over which data is aggregated. If no aggregation dimension is specified, the query defaults to providing detailed, unaggregated results.
	//
	// Optional Values:
	//
	// - Year
	//
	// - Month
	//
	// - Day
	//
	// - All
	//
	// example:
	//
	// Month
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The dimension names. (Equivalent to SQL\\"s GROUP BY Clause)
	//
	// Optional Values:
	//
	// - UserId
	//
	// - ServiceId
	//
	// - ServiceVersion
	//
	// - ServiceInstanceId
	//
	// - DeploySucceeded
	//
	// - ErrorType
	//
	// - ErrorCode
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-12-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter.
	Filter []*ListServiceInstanceDeployDetailsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAZbOYA+x9UgM6xrgcMqFUjk=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2024-08-25T02:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListServiceInstanceDeployDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsRequest) SetCycleTimeZone(v string) *ListServiceInstanceDeployDetailsRequest {
	s.CycleTimeZone = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetCycleType(v string) *ListServiceInstanceDeployDetailsRequest {
	s.CycleType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetDimension(v []*string) *ListServiceInstanceDeployDetailsRequest {
	s.Dimension = v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetEndTime(v string) *ListServiceInstanceDeployDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetFilter(v []*ListServiceInstanceDeployDetailsRequestFilter) *ListServiceInstanceDeployDetailsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetMaxResults(v int32) *ListServiceInstanceDeployDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetNextToken(v string) *ListServiceInstanceDeployDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetRegionId(v string) *ListServiceInstanceDeployDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetStartTime(v string) *ListServiceInstanceDeployDetailsRequest {
	s.StartTime = &v
	return s
}

type ListServiceInstanceDeployDetailsRequestFilter struct {
	// Filter Value Names (Equivalent to SQL\\"s WHERE Clause)
	//
	// Available Options:
	//
	// - UserId
	//
	// - ServiceId
	//
	// - ServiceVersion
	//
	// - ServiceInstanceId
	//
	// - DeploySucceeded (Accepts True or False and case-insensitive)
	//
	// - ErrorType
	//
	// - ErrorCode
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceDeployDetailsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) SetName(v string) *ListServiceInstanceDeployDetailsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) SetValue(v []*string) *ListServiceInstanceDeployDetailsRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstanceDeployDetailsResponseBody struct {
	// The details of the service instance deployment.
	DeployDetails []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails `json:"DeployDetails,omitempty" xml:"DeployDetails,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0036D82E-0624-5B37-B797-C460F4B02026
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetDeployDetails(v []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails) *ListServiceInstanceDeployDetailsResponseBody {
	s.DeployDetails = v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetMaxResults(v int32) *ListServiceInstanceDeployDetailsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetNextToken(v string) *ListServiceInstanceDeployDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetRequestId(v string) *ListServiceInstanceDeployDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetTotalCount(v int32) *ListServiceInstanceDeployDetailsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstanceDeployDetailsResponseBodyDeployDetails struct {
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 4
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2024-04-10T01:58:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The period over which data is aggregated.
	//
	// example:
	//
	// Month
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The indicates whether the deployment was successful.
	//
	// example:
	//
	// False
	DeploySucceeded *string `json:"DeploySucceeded,omitempty" xml:"DeploySucceeded,omitempty"`
	// The error code.
	//
	// example:
	//
	// StackValidationFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// {code: StackValidationFailed, message: \\"Failed to continue create ROS stack 89e724e2-84e6-4517-a372-30a545ab4145: Resource [LinuxInstanceRunCommand]: i-wz91nfbh1fxtmfb0try4 are not running. Command invocation only support running instances. ErrorCode: StackValidationFailed\\", requestId: null}
	ErrorDetail *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	// The type of error that caused the deployment to fail.
	//
	// example:
	//
	// ValidationError
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c751ed91f2074af39779
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-273e8cee11d349e1803c
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The name of the service in Chinese.
	//
	// example:
	//
	// 测试服务(Test Service)
	ServiceNameChn *string `json:"ServiceNameChn,omitempty" xml:"ServiceNameChn,omitempty"`
	// The name of the service in English.
	//
	// example:
	//
	// Test Service
	ServiceNameEng *string `json:"ServiceNameEng,omitempty" xml:"ServiceNameEng,omitempty"`
	// The type of service.
	//
	// Possible values:
	//
	// - private: Deployed under the user\\"s account.
	//
	// - managed: Hosted under the service provider\\"s account.
	//
	// - operation: Managed operation service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 42
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The timestamp when the response is returned.
	//
	// example:
	//
	// 1723946641994
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The aliuid of user.
	//
	// example:
	//
	// 1591457835436382
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponseBodyDeployDetails) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCount(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Count = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCreateTime(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCycle(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Cycle = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetDeploySucceeded(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.DeploySucceeded = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorCode(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorDetail(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorDetail = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorType(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceInstanceId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceNameChn(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceNameChn = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceNameEng(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceNameEng = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceType(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceVersion(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceVersion = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetTimestamp(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Timestamp = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetUserId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.UserId = &v
	return s
}

type ListServiceInstanceDeployDetailsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceDeployDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceDeployDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponse) SetStatusCode(v int32) *ListServiceInstanceDeployDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponse) SetBody(v *ListServiceInstanceDeployDetailsResponseBody) *ListServiceInstanceDeployDetailsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceLogsRequest struct {
	// The filters.
	Filter []*ListServiceInstanceLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The log source. When this field is empty, query logs with the source set to computeNest and ros.
	//
	// Valid values:
	//
	// 	- computeNest : logs of the deployment and upgrade of the service instance.
	//
	// 	- application: logs generated by the application.
	//
	// 	- actionTrail: logs generated by ActionTrail.
	//
	// 	- compliancePack:  Logs originating from the compliance package.
	//
	// 	- ros: Logs originating from ROS.
	//
	// 	- meteringData：Logs originating from the pay-as-you-go model.
	//
	// example:
	//
	// computeNest
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The name of the Logstore to which log entries are delivered.
	//
	// It needs to be provided only when LogSource is set to Application.
	//
	// example:
	//
	// livelog
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb626435b****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// Sort Order. Possible values:
	//
	// + Ascending: Ascending order
	//
	// + Descending (default value): Descending order
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) SetFilter(v []*ListServiceInstanceLogsRequestFilter) *ListServiceInstanceLogsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogSource(v string) *ListServiceInstanceLogsRequest {
	s.LogSource = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogstore(v string) *ListServiceInstanceLogsRequest {
	s.Logstore = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v int32) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetSortOrder(v string) *ListServiceInstanceLogsRequest {
	s.SortOrder = &v
	return s
}

type ListServiceInstanceLogsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- StartTime: the start time of the log event.
	//
	// 	- EndTime: the end time of the ActionTrail event.
	//
	// 	- EventName: the name of the ActionTrail event.
	//
	// 	- ResourceName: the name of the ActionTrail resource.
	//
	// 	- ApplicationGroupName: the name of the application group.
	//
	// example:
	//
	// StartTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequestFilter) SetName(v string) *ListServiceInstanceLogsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceLogsRequestFilter) SetValue(v []*string) *ListServiceInstanceLogsRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstanceLogsResponseBody struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The logs of the service instance.
	ServiceInstancesLogs []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs `json:"ServiceInstancesLogs,omitempty" xml:"ServiceInstancesLogs,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBody) SetMaxResults(v int32) *ListServiceInstanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetNextToken(v string) *ListServiceInstanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetRequestId(v string) *ListServiceInstanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetServiceInstancesLogs(v []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs) *ListServiceInstanceLogsResponseBody {
	s.ServiceInstancesLogs = v
	return s
}

type ListServiceInstanceLogsResponseBodyServiceInstancesLogs struct {
	// Compliance package risk types. This only applies when the source is CompliancePack. . For example, data security checks within a VPC, such as VpcDataRisk
	//
	// example:
	//
	// VpcDataRisk
	CompliancePackType *string `json:"CompliancePackType,omitempty" xml:"CompliancePackType,omitempty"`
	// Specific risk rule names for the compliance package. This only applies when the source is CompliancePack. . For example, ECS instance migration out of VPC - ecs-move-out-vpc.
	//
	// example:
	//
	// vpc-ecs-move-out-vpc
	ComplianceRuleName *string `json:"ComplianceRuleName,omitempty" xml:"ComplianceRuleName,omitempty"`
	// The log content.
	//
	// example:
	//
	// Start creating service instance
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The log type. Valid values:
	//
	// 	- serviceInstance: log generated by the service instance.
	//
	// 	- resource: log generated by ROS resources.
	//
	// example:
	//
	// serviceInstance
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// si-5c6525c0589545c3****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The Resouce Type.
	//
	// example:
	//
	// ROS.Stack
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The log source.
	//
	// Valid values:
	//
	// 	- computeNest : logs of the deployment and upgrade of the service instance.
	//
	// 	- application: logs generated by the application.
	//
	// 	- actionTrail: logs generated by ActionTrail.
	//
	// 	- compliancePack:  Logs originating from the compliance package.
	//
	// 	- ros: Logs originating from ROS.
	//
	// 	- meteringData：Logs originating from the pay-as-you-go model.
	//
	// example:
	//
	// computeNest
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The deployment state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The timestamp of the service instance log.
	//
	// example:
	//
	// 2022-05-21T00:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetCompliancePackType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.CompliancePackType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetComplianceRuleName(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ComplianceRuleName = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetContent(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Content = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetLogType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.LogType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceId(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetSource(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Source = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetStatus(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetTimestamp(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Timestamp = &v
	return s
}

type ListServiceInstanceLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetStatusCode(v int32) *ListServiceInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceResourcesRequest struct {
	// The filter.
	Filters []*ListServiceInstanceResourcesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// Service Instance resource type，include AliyunResource and ContainerResource.
	//
	// example:
	//
	// AliyunResource
	ServiceInstanceResourceType *string `json:"ServiceInstanceResourceType,omitempty" xml:"ServiceInstanceResourceType,omitempty"`
	// The tags.
	Tag []*ListServiceInstanceResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequest) SetFilters(v []*ListServiceInstanceResourcesRequestFilters) *ListServiceInstanceResourcesRequest {
	s.Filters = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetMaxResults(v int32) *ListServiceInstanceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetNextToken(v string) *ListServiceInstanceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetRegionId(v string) *ListServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceResourceType(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceResourceType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest {
	s.Tag = v
	return s
}

type ListServiceInstanceResourcesRequestFilters struct {
	// Vaild values:
	//
	// - ExpireTimeStart
	//
	// - ExpireTimeEnd
	//
	// - PayType
	//
	// - ResourceARN
	//
	// example:
	//
	// ResourceARN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestFilters) SetName(v string) *ListServiceInstanceResourcesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestFilters) SetValues(v []*string) *ListServiceInstanceResourcesRequestFilters {
	s.Values = v
	return s
}

type ListServiceInstanceResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstanceResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestTag) SetKey(v string) *ListServiceInstanceResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) SetValue(v string) *ListServiceInstanceResourcesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstanceResourcesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	Resources []*ListServiceInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBody) SetMaxResults(v int32) *ListServiceInstanceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetNextToken(v string) *ListServiceInstanceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetRequestId(v string) *ListServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

type ListServiceInstanceResourcesResponseBodyResources struct {
	// The time when the service instance was created.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription
	//
	// 	- PayAsYouGo
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// RDS
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The renewal state. Valid values:
	//
	// 	- AutoRenewal
	//
	// 	- ManualRenewal
	//
	// 	- NotRenewal
	//
	// example:
	//
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The renewal period.
	//
	// example:
	//
	// 1
	RenewalPeriod *int32 `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:sag:cn-hangzhou:130920852836****:ccn/ccn-b3qf0q439sq2de****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServiceInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetCreateTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetExpireTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetPayType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductCode(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewStatus = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriod(v int32) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriod = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriodUnit(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetResourceARN(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ResourceARN = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

type ListServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetStatusCode(v int32) *ListServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type ListServiceInstanceUpgradeHistoryRequest struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query, which should be the value of the NextToken parameter from the previous API call.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetRegionId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetServiceInstanceId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceUpgradeHistoryResponseBody struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to use for the next query.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 86CAC31E-3527-562C-869F-347E931C9B25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total count of upgrade history
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of upgrade histories.
	UpgradeHistory []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory `json:"UpgradeHistory,omitempty" xml:"UpgradeHistory,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetRequestId(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetTotalCount(v int64) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetUpgradeHistory(v []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.UpgradeHistory = v
	return s
}

type ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory struct {
	// End time of the upgrade.
	//
	// example:
	//
	// 2022-04-26T09:09:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Version before the upgrade.
	//
	// example:
	//
	// 1
	FromVersion *string `json:"FromVersion,omitempty" xml:"FromVersion,omitempty"`
	// Upgrade result.
	//
	// example:
	//
	// {\\"PreUpgradeExecutionId\\":\\"exec-123\\"}
	Results *string `json:"Results,omitempty" xml:"Results,omitempty"`
	// Start time of the upgrade.
	//
	// example:
	//
	// 2022-04-26T08:09:51Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Upgrade status. Possible values:
	//
	// - upgrading: In progress.
	//
	// - UpgradeSuccessful: Upgrade successful.
	//
	// - UpgradeFailed: Upgrade failed.
	//
	// example:
	//
	// UpgradeFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Version after the upgrade.
	//
	// example:
	//
	// 3
	ToVersion *string `json:"ToVersion,omitempty" xml:"ToVersion,omitempty"`
	// Upgrade type.
	//
	// - Upgrade
	//
	// - Rollback
	//
	// example:
	//
	// Upgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Upgrade history ID.
	//
	// example:
	//
	// uh-1b21d65f75e94fa09745
	UpgradeHistoryId *string `json:"UpgradeHistoryId,omitempty" xml:"UpgradeHistoryId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetEndTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetFromVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.FromVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetResults(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Results = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStartTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.StartTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStatus(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetToVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.ToVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetType(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Type = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetUpgradeHistoryId(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.UpgradeHistoryId = &v
	return s
}

type ListServiceInstanceUpgradeHistoryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceUpgradeHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetHeaders(v map[string]*string) *ListServiceInstanceUpgradeHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetStatusCode(v int32) *ListServiceInstanceUpgradeHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetBody(v *ListServiceInstanceUpgradeHistoryResponseBody) *ListServiceInstanceUpgradeHistoryResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	// The filter.
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to display the information that the service instance is deleted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ShowDeleted *bool `json:"ShowDeleted,omitempty" xml:"ShowDeleted,omitempty"`
	// The custom tags.
	Tag []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetShowDeleted(v bool) *ListServiceInstancesRequest {
	s.ShowDeleted = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- Name: The service name. If you want to perform a fuzzy match, specify the service name in the *xxx	- format. For example, if the service name is My Service, you can set the filter value to *My	- or *Service*.
	//
	// 	- ServiceInstanceId: The ID of the service instance.
	//
	// 	- ServiceId: The service ID.
	//
	// 	- UserId: The user ID.
	//
	// 	- Version: The service version.
	//
	// 	- Status: The status of the service instance.
	//
	// 	- DeployType: The deployment type of the service.
	//
	// 	- ServiceType: The service type.
	//
	// 	- OperationStartTimeBefore: The time before the hosted O\\&M starts.
	//
	// 	- OperationStartTimeAfter: The time after the hosted O\\&M starts.
	//
	// 	- OperationEndTimeBefore: The time before the hosted O\\&M ends.
	//
	// 	- OperationEndTimeAfter: The time after the hosted O\\&M ends.
	//
	// 	- OperatedServiceInstanceId: The ID of the hosted O\\&M instance that belongs to a private service.
	//
	// 	- OperationServiceInstanceId: The ID of the hosted O\\&M service instance that belongs to a hosted O\\&M service.
	//
	// 	- EnableInstanceOps: Whether the hosted O\\&M feature is enabled for service instances.
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about service instances.
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v int32) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	// The business status of the service instance. Valid values:
	//
	// 	- Normal: The service instance is normal.
	//
	// 	- Renewing: The service instance is being renewed.
	//
	// 	- RenewFailed: The service instance failed to be renewed.
	//
	// 	- Expired: The service instance expired.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// The time when the service instance expires.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the hosted O\\&M feature is enabled for the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the service instance that is used to implement hosted O\\&M.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The end of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The beginning of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The parameters of the service instance.
	//
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method of the service. Valid values:
	//
	// 	- Permanent: Once you purchase the service, you can use it permanently.
	//
	// 	- Subscription: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a subscription basis.
	//
	// 	- PayAsYouGo: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a pay-as-you-go basis.
	//
	// 	- CustomFixTime: You are charged for the service based on a custom duration fixed by the service provider.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// si-d6ab3a63ccxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the service.
	Service *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source from which the service instance is created.
	//
	// example:
	//
	// Supplier
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The custom tags.
	Tags []*ListServiceInstancesResponseBodyServiceInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the service instance was updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// example:
	//
	// 1234567
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetBizStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.BizStatus = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetIsOperated(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.IsOperated = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResourceGroupId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUserId(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.UserId = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Indicates whether the private connection feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The name of the distribution source service provider.
	//
	// example:
	//
	// Company A
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// The status of the service.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// 版本A
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetEnablePrivateVpcConnection(v bool) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSourceSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SourceSupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The service name.
	//
	// example:
	//
	// WordPress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServiceRegistrationsRequest struct {
	// Filter.
	Filter []*ListServiceRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of items to return per page during a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListServiceRegistrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequest) SetFilter(v []*ListServiceRegistrationsRequestFilter) *ListServiceRegistrationsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceRegistrationsRequest) SetMaxResults(v int32) *ListServiceRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetNextToken(v string) *ListServiceRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetRegionId(v string) *ListServiceRegistrationsRequest {
	s.RegionId = &v
	return s
}

type ListServiceRegistrationsRequestFilter struct {
	// Name of the filter field. Allowed values:
	//
	// - ServiceId: Service ID.
	//
	// - RegistrationId: Registration ID.
	//
	// - Status: Registration status. Allowed values: Submitted, Approved, Rejected, Canceled, and Executed.
	//
	// example:
	//
	// Canceled
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceRegistrationsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequestFilter) SetName(v string) *ListServiceRegistrationsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceRegistrationsRequestFilter) SetValue(v []*string) *ListServiceRegistrationsRequestFilter {
	s.Value = v
	return s
}

type ListServiceRegistrationsResponseBody struct {
	// Number of items per page in a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service registration information.
	ServiceRegistrations []*ListServiceRegistrationsResponseBodyServiceRegistrations `json:"ServiceRegistrations,omitempty" xml:"ServiceRegistrations,omitempty" type:"Repeated"`
	// Total number of records that meet the criteria.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceRegistrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBody) SetMaxResults(v int32) *ListServiceRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetNextToken(v string) *ListServiceRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetRequestId(v string) *ListServiceRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetServiceRegistrations(v []*ListServiceRegistrationsResponseBodyServiceRegistrations) *ListServiceRegistrationsResponseBody {
	s.ServiceRegistrations = v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetTotalCount(v int32) *ListServiceRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceRegistrationsResponseBodyServiceRegistrations struct {
	// Comment.
	//
	// example:
	//
	// some info is missing
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Finish time.
	//
	// example:
	//
	// 2021-05-23T00:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Registration ID.
	//
	// example:
	//
	// sr-001
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// Service ID.
	//
	// example:
	//
	// service-f4c0026a254bxxxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Registration status. Allowed values:
	//
	// - Submitted
	//
	// - Approved
	//
	// - Rejected
	//
	// - Canceled
	//
	// - Executed
	//
	// - Executed: Executed.
	//
	// example:
	//
	// Rejected
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Submit time.
	//
	// example:
	//
	// 2021-05-22T00:00:00Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetComment(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Comment = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetFinishTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.FinishTime = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetRegistrationId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetServiceId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.ServiceId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetStatus(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Status = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetSubmitTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.SubmitTime = &v
	return s
}

type ListServiceRegistrationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceRegistrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponse) SetHeaders(v map[string]*string) *ListServiceRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRegistrationsResponse) SetStatusCode(v int32) *ListServiceRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceRegistrationsResponse) SetBody(v *ListServiceRegistrationsResponseBody) *ListServiceRegistrationsResponse {
	s.Body = v
	return s
}

type ListServiceSharedAccountsRequest struct {
	// The filters.
	Filter []*ListServiceSharedAccountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAR130adlM4fHHVSWpTca/t4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The region ID where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-e10349089de34exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListServiceSharedAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsRequest) SetFilter(v []*ListServiceSharedAccountsRequestFilter) *ListServiceSharedAccountsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetMaxResults(v int32) *ListServiceSharedAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetNextToken(v string) *ListServiceSharedAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetPermission(v string) *ListServiceSharedAccountsRequest {
	s.Permission = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetRegionId(v string) *ListServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetServiceId(v string) *ListServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

type ListServiceSharedAccountsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- Name: the name of the service.
	//
	// example:
	//
	// UserAliUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value N of the filter. Valid values of N: 1 to 10.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceSharedAccountsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceSharedAccountsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsRequestFilter) SetName(v string) *ListServiceSharedAccountsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceSharedAccountsRequestFilter) SetValue(v []*string) *ListServiceSharedAccountsRequestFilter {
	s.Value = v
	return s
}

type ListServiceSharedAccountsResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA3AE512-6D30-549A-B52D-B9042CA8D515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service shared account information.
	ShareAccount []*ListServiceSharedAccountsResponseBodyShareAccount `json:"ShareAccount,omitempty" xml:"ShareAccount,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceSharedAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponseBody) SetMaxResults(v int32) *ListServiceSharedAccountsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetNextToken(v string) *ListServiceSharedAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetRequestId(v string) *ListServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetShareAccount(v []*ListServiceSharedAccountsResponseBodyShareAccount) *ListServiceSharedAccountsResponseBody {
	s.ShareAccount = v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetTotalCount(v int32) *ListServiceSharedAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceSharedAccountsResponseBodyShareAccount struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2021-12-28T02:47:46.000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Service logo.
	//
	// example:
	//
	// logo
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// The name of the service instance. The value must meet the following requirements:
	//
	// 	- The name cannot exceed 64 characters in length.
	//
	// 	- It can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-e10349089de34exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The time when the service was updated.
	//
	// example:
	//
	// 2023-02-13T02:16:03.756Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user aliUid.
	//
	// example:
	//
	// 127383705xxxxxx
	UserAliUid *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s ListServiceSharedAccountsResponseBodyShareAccount) String() string {
	return tea.Prettify(s)
}

func (s ListServiceSharedAccountsResponseBodyShareAccount) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetCreateTime(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.CreateTime = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetLogo(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Logo = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetName(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Name = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetPermission(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Permission = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetServiceId(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.ServiceId = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetUpdateTime(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetUserAliUid(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.UserAliUid = &v
	return s
}

type ListServiceSharedAccountsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceSharedAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *ListServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceSharedAccountsResponse) SetStatusCode(v int32) *ListServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceSharedAccountsResponse) SetBody(v *ListServiceSharedAccountsResponseBody) *ListServiceSharedAccountsResponse {
	s.Body = v
	return s
}

type ListServiceTestCasesRequest struct {
	// The filters.
	Filters []*ListServiceTestCasesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListServiceTestCasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestCasesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesRequest) SetFilters(v []*ListServiceTestCasesRequestFilters) *ListServiceTestCasesRequest {
	s.Filters = v
	return s
}

func (s *ListServiceTestCasesRequest) SetMaxResults(v int32) *ListServiceTestCasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetNextToken(v string) *ListServiceTestCasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetRegionId(v string) *ListServiceTestCasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetServiceId(v string) *ListServiceTestCasesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetServiceVersion(v string) *ListServiceTestCasesRequest {
	s.ServiceVersion = &v
	return s
}

type ListServiceTestCasesRequestFilters struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// **Status**
	//
	// **TaskId**
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceTestCasesRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestCasesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesRequestFilters) SetName(v string) *ListServiceTestCasesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListServiceTestCasesRequestFilters) SetValue(v []*string) *ListServiceTestCasesRequestFilters {
	s.Value = v
	return s
}

type ListServiceTestCasesResponseBody struct {
	// The data returned.
	Data []*ListServiceTestCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CA3AE512-6D30-549A-B52D-B9042CA8D515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceTestCasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestCasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponseBody) SetData(v []*ListServiceTestCasesResponseBodyData) *ListServiceTestCasesResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetMaxResults(v int32) *ListServiceTestCasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetNextToken(v string) *ListServiceTestCasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetRequestId(v string) *ListServiceTestCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestCasesResponseBody) SetTotalCount(v int32) *ListServiceTestCasesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceTestCasesResponseBodyData struct {
	// The template name.
	//
	// example:
	//
	// test-1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The service test case id.
	//
	// example:
	//
	// stc-83fcee1383354e35b151
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
	// The service test case name.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// The service test config.
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s ListServiceTestCasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponseBodyData) SetTemplateName(v string) *ListServiceTestCasesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestCaseId(v string) *ListServiceTestCasesResponseBodyData {
	s.TestCaseId = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestCaseName(v string) *ListServiceTestCasesResponseBodyData {
	s.TestCaseName = &v
	return s
}

func (s *ListServiceTestCasesResponseBodyData) SetTestConfig(v string) *ListServiceTestCasesResponseBodyData {
	s.TestConfig = &v
	return s
}

type ListServiceTestCasesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestCasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestCasesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponse) SetHeaders(v map[string]*string) *ListServiceTestCasesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestCasesResponse) SetStatusCode(v int32) *ListServiceTestCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestCasesResponse) SetBody(v *ListServiceTestCasesResponseBody) *ListServiceTestCasesResponse {
	s.Body = v
	return s
}

type ListServiceTestTaskLogsRequest struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sort Order. Possible values:
	//
	// + Ascending: Ascending order
	//
	// + Descending (default value): Descending order
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-568c2c5a687a409b977e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListServiceTestTaskLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTaskLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsRequest) SetMaxResults(v int32) *ListServiceTestTaskLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetNextToken(v string) *ListServiceTestTaskLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetSortOrder(v string) *ListServiceTestTaskLogsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListServiceTestTaskLogsRequest) SetTaskId(v string) *ListServiceTestTaskLogsRequest {
	s.TaskId = &v
	return s
}

type ListServiceTestTaskLogsResponseBody struct {
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDEE055B-D5F4-5B92-8F21-999D408F1214
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service test task logs.
	TaskLogs []*ListServiceTestTaskLogsResponseBodyTaskLogs `json:"TaskLogs,omitempty" xml:"TaskLogs,omitempty" type:"Repeated"`
}

func (s ListServiceTestTaskLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponseBody) SetMaxResults(v int32) *ListServiceTestTaskLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetNextToken(v string) *ListServiceTestTaskLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetRequestId(v string) *ListServiceTestTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBody) SetTaskLogs(v []*ListServiceTestTaskLogsResponseBodyTaskLogs) *ListServiceTestTaskLogsResponseBody {
	s.TaskLogs = v
	return s
}

type ListServiceTestTaskLogsResponseBodyTaskLogs struct {
	// The log content.
	//
	// example:
	//
	// log content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The UTC timestamp when the response is returned.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-12-12T20:00:09Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceTestTaskLogsResponseBodyTaskLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTaskLogsResponseBodyTaskLogs) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) SetContent(v string) *ListServiceTestTaskLogsResponseBodyTaskLogs {
	s.Content = &v
	return s
}

func (s *ListServiceTestTaskLogsResponseBodyTaskLogs) SetTimestamp(v string) *ListServiceTestTaskLogsResponseBodyTaskLogs {
	s.Timestamp = &v
	return s
}

type ListServiceTestTaskLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestTaskLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestTaskLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTaskLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestTaskLogsResponse) SetHeaders(v map[string]*string) *ListServiceTestTaskLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestTaskLogsResponse) SetStatusCode(v int32) *ListServiceTestTaskLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestTaskLogsResponse) SetBody(v *ListServiceTestTaskLogsResponseBody) *ListServiceTestTaskLogsResponse {
	s.Body = v
	return s
}

type ListServiceTestTasksRequest struct {
	// The filters.
	Filter []*ListServiceTestTasksRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Number of items per page in a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-062ae8e13b394dd5b63c
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListServiceTestTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTasksRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksRequest) SetFilter(v []*ListServiceTestTasksRequestFilter) *ListServiceTestTasksRequest {
	s.Filter = v
	return s
}

func (s *ListServiceTestTasksRequest) SetMaxResults(v int32) *ListServiceTestTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetNextToken(v string) *ListServiceTestTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetRegionId(v string) *ListServiceTestTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetServiceId(v string) *ListServiceTestTasksRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetServiceVersion(v string) *ListServiceTestTasksRequest {
	s.ServiceVersion = &v
	return s
}

type ListServiceTestTasksRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- Status: the status of the task.
	//
	// 	- TaskId: the task id.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value N of the filter. Valid values of N: 1 to 10.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceTestTasksRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTasksRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksRequestFilter) SetName(v string) *ListServiceTestTasksRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceTestTasksRequestFilter) SetValue(v []*string) *ListServiceTestTasksRequestFilter {
	s.Value = v
	return s
}

type ListServiceTestTasksResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service test tasks.
	ServiceTestTasks []*ListServiceTestTasksResponseBodyServiceTestTasks `json:"ServiceTestTasks,omitempty" xml:"ServiceTestTasks,omitempty" type:"Repeated"`
}

func (s ListServiceTestTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponseBody) SetCount(v int32) *ListServiceTestTasksResponseBody {
	s.Count = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetMaxResults(v int32) *ListServiceTestTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetNextToken(v string) *ListServiceTestTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetRequestId(v string) *ListServiceTestTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetServiceTestTasks(v []*ListServiceTestTasksResponseBodyServiceTestTasks) *ListServiceTestTasksResponseBody {
	s.ServiceTestTasks = v
	return s
}

type ListServiceTestTasksResponseBodyServiceTestTasks struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2024-02-26T02:16:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// the status of service task.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// sttt-000h5nd4yrg59ucurzy1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// dadadad
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task region id.
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
}

func (s ListServiceTestTasksResponseBodyServiceTestTasks) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTasksResponseBodyServiceTestTasks) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetCreateTime(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.CreateTime = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetStatus(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.Status = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskId(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskId = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskName(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskName = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskRegionId(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskRegionId = &v
	return s
}

type ListServiceTestTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTestTasksResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponse) SetHeaders(v map[string]*string) *ListServiceTestTasksResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestTasksResponse) SetStatusCode(v int32) *ListServiceTestTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestTasksResponse) SetBody(v *ListServiceTestTasksResponseBody) *ListServiceTestTasksResponse {
	s.Body = v
	return s
}

type ListServiceUsagesRequest struct {
	// The filter.
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The role of the service provider.
	//
	// example:
	//
	// Supplier
	SupplierRole *string `json:"SupplierRole,omitempty" xml:"SupplierRole,omitempty"`
}

func (s ListServiceUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequest) SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceUsagesRequest) SetMaxResults(v int32) *ListServiceUsagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesRequest) SetNextToken(v string) *ListServiceUsagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesRequest) SetSupplierRole(v string) *ListServiceUsagesRequest {
	s.SupplierRole = &v
	return s
}

type ListServiceUsagesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- ServiceName: the service name.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value N of the filter. Valid values of N: 1 to 10.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceUsagesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequestFilter) SetName(v string) *ListServiceUsagesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesRequestFilter) SetValue(v []*string) *ListServiceUsagesRequestFilter {
	s.Value = v
	return s
}

type ListServiceUsagesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18AD0960-A9FE-1AC8-ADF8-22131Fxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service applications.
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBody) SetMaxResults(v int32) *ListServiceUsagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetNextToken(v string) *ListServiceUsagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetRequestId(v string) *ListServiceUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody {
	s.ServiceUsages = v
	return s
}

func (s *ListServiceUsagesResponseBody) SetTotalCount(v int32) *ListServiceUsagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	// The comment on the approval.
	//
	// example:
	//
	// Reject
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c9f36ec6d19b4exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// ServiceTest
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The status of the service application. Valid values:
	//
	// 	- Submitted: The application is submitted or is to be approved.
	//
	// 	- Approved: The application is approved.
	//
	// 	- Rejected: The application is rejected.
	//
	// 	- Canceled: The application is canceled.
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// SupplierAAA
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	// The user information.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s ListServiceUsagesResponseBodyServiceUsages) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsages) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetComments(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Comments = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetCreateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.CreateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceId(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceId = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetStatus(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Status = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetSupplierName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.SupplierName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUpdateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserAliUid(v int64) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserInformation(v map[string]*string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserInformation = v
	return s
}

type ListServiceUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponse) SetHeaders(v map[string]*string) *ListServiceUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceUsagesResponse) SetStatusCode(v int32) *ListServiceUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceUsagesResponse) SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// Specifies whether to return all versions of a service. Default value: false, which specifies that only the default version of a service is returned.
	//
	// example:
	//
	// false
	AllVersions *bool `json:"AllVersions,omitempty" xml:"AllVersions,omitempty"`
	// The filters.
	Filter []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The custom tags.
	Tag []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetAllVersions(v bool) *ListServicesRequest {
	s.AllVersions = &v
	return s
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetResourceGroupId(v string) *ListServicesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesRequest) SetTag(v []*ListServicesRequestTag) *ListServicesRequest {
	s.Tag = v
	return s
}

type ListServicesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- Name: the name of the service.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServicesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

type ListServicesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTag) SetKey(v string) *ListServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTag) SetValue(v string) *ListServicesRequestTag {
	s.Value = &v
	return s
}

type ListServicesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The services.
	Services []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetMaxResults(v int32) *ListServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponseBodyServices struct {
	// The approval type for applications for using the service. Valid values:
	//
	// 	- Manual: The applications are manual reviewed.
	//
	// 	- AutoPass: The applications are automatically approved.
	//
	// example:
	//
	// AutoPass
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The ID of the artifact.
	//
	// example:
	//
	// artifact-21ca53ac16a643xxxxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The version of the artifact.
	//
	// example:
	//
	// draft
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The informathon for build service.
	//
	// example:
	//
	// { "RepoUrl": "https://github.com/user/repo.git", "Brancn": "main"}
	BuildInfo *string `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	// The category of the service.
	//
	// example:
	//
	// OpenSource
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity details.
	Commodity *ListServicesResponseBodyServicesCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the service was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- spi: The service is deployed by calling the Service Provider Interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// 	- container: The service is deployed by using a container.
	//
	// *
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Indicates whether the service has a beta version. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	HasBeta *bool `json:"HasBeta,omitempty" xml:"HasBeta,omitempty"`
	// Indicates whether the service has a draft version. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	HasDraft *bool `json:"HasDraft,omitempty" xml:"HasDraft,omitempty"`
	// The latest version of the distribution source service.
	//
	// example:
	//
	// 1
	LatestResellSourceServiceVersion *string `json:"LatestResellSourceServiceVersion,omitempty" xml:"LatestResellSourceServiceVersion,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The purpose of the artifact. Valid values:
	//
	// 	- ServiceDeployment: The artifact is used to create service instances.
	//
	// 	- ServiceUpgrade: The artifact is used to upgrade service instances.
	//
	// example:
	//
	// ServiceDeployment
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The state of distribution authorization of the service. Valid values:
	//
	// 	- CanApply: Distributors can apply for distribution permissions.
	//
	// 	- Applied: The application for distribution permissions is submitted.
	//
	// 	- Approved: The application for distribution permissions is approved.
	//
	// example:
	//
	// CanApply
	ResellApplyStatus *string `json:"ResellApplyStatus,omitempty" xml:"ResellApplyStatus,omitempty"`
	// The ID of the distribution service.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	ResellServiceId *string `json:"ResellServiceId,omitempty" xml:"ResellServiceId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz5b555xxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the service is visible. Valid values:
	//
	// 	- INVISIBLE
	//
	// 	- DISCOVERABLE
	//
	// example:
	//
	// INVISIBLE
	ServiceDiscoverable *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a formal service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a formal service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a formal service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The source image.
	//
	// example:
	//
	// centos_7
	SourceImage *string `json:"SourceImage,omitempty" xml:"SourceImage,omitempty"`
	// The ID of the distribution source service.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The version of the distribution source service.
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The name of the distribution source service provider.
	//
	// example:
	//
	// SourceSupplier
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// The state of the service. Valid values:
	//
	// 	- Draft: The service is a draft.
	//
	// 	- Submitted: The service is submitted for review. You cannot modify services in this state.
	//
	// 	- Approved: The service is approved. You cannot modify services in this state. You can publish services in this state.
	//
	// 	- Launching: The service is being published.
	//
	// 	- Online: The service is published.
	//
	// 	- Offline: The service is unpublished.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service tags.
	Tags []*ListServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant type of the managed service. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The time when the service was modified.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// v2.0.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Indicates whether the service is a virtual Internet service. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetApprovalType(v string) *ListServicesResponseBodyServices {
	s.ApprovalType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetArtifactId(v string) *ListServicesResponseBodyServices {
	s.ArtifactId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetArtifactVersion(v string) *ListServicesResponseBodyServices {
	s.ArtifactVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetBuildInfo(v string) *ListServicesResponseBodyServices {
	s.BuildInfo = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCategories(v string) *ListServicesResponseBodyServices {
	s.Categories = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodity(v *ListServicesResponseBodyServicesCommodity) *ListServicesResponseBodyServices {
	s.Commodity = v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodityCode(v string) *ListServicesResponseBodyServices {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCreateTime(v string) *ListServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDefaultVersion(v bool) *ListServicesResponseBodyServices {
	s.DefaultVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployType(v string) *ListServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetHasBeta(v bool) *ListServicesResponseBodyServices {
	s.HasBeta = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetHasDraft(v bool) *ListServicesResponseBodyServices {
	s.HasDraft = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetLatestResellSourceServiceVersion(v string) *ListServicesResponseBodyServices {
	s.LatestResellSourceServiceVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPublishTime(v string) *ListServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetRelationType(v string) *ListServicesResponseBodyServices {
	s.RelationType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResellApplyStatus(v string) *ListServicesResponseBodyServices {
	s.ResellApplyStatus = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResellServiceId(v string) *ListServicesResponseBodyServices {
	s.ResellServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetResourceGroupId(v string) *ListServicesResponseBodyServices {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceDiscoverable(v string) *ListServicesResponseBodyServices {
	s.ServiceDiscoverable = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceInfos(v []*ListServicesResponseBodyServicesServiceInfos) *ListServicesResponseBodyServices {
	s.ServiceInfos = v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetShareType(v string) *ListServicesResponseBodyServices {
	s.ShareType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceImage(v string) *ListServicesResponseBodyServices {
	s.SourceImage = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceServiceId(v string) *ListServicesResponseBodyServices {
	s.SourceServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceServiceVersion(v string) *ListServicesResponseBodyServices {
	s.SourceServiceVersion = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSourceSupplierName(v string) *ListServicesResponseBodyServices {
	s.SourceSupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetStatus(v string) *ListServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierName(v string) *ListServicesResponseBodyServices {
	s.SupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUrl(v string) *ListServicesResponseBodyServices {
	s.SupplierUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTags(v []*ListServicesResponseBodyServicesTags) *ListServicesResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListServicesResponseBodyServices) SetTenantType(v string) *ListServicesResponseBodyServices {
	s.TenantType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTrialType(v string) *ListServicesResponseBodyServices {
	s.TrialType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetUpdateTime(v string) *ListServicesResponseBodyServices {
	s.UpdateTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersion(v string) *ListServicesResponseBodyServices {
	s.Version = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersionName(v string) *ListServicesResponseBodyServices {
	s.VersionName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVirtualInternetService(v string) *ListServicesResponseBodyServices {
	s.VirtualInternetService = &v
	return s
}

type ListServicesResponseBodyServicesCommodity struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The configuration metadata related to Saas Boost.
	//
	// example:
	//
	// { // Specifies whether to associate the service with the SaaS Boost commodity. Default value: false. "Enabled":true/false // The public endpoint of the SaaS Boost instance. "PublicAccessUrl":"https://example.com" }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The platform type. Valid values:
	//
	// 	- marketplace: Alibaba Cloud Marketplace.
	//
	// 	- Css: Lingxiao.
	//
	// 	- SaasBoost: Saas Boost.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServicesResponseBodyServicesCommodity) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesCommodity) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesCommodity) SetCommodityCode(v string) *ListServicesResponseBodyServicesCommodity {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetSaasBoostMetadata(v string) *ListServicesResponseBodyServicesCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetType(v string) *ListServicesResponseBodyServicesCommodity {
	s.Type = &v
	return s
}

type ListServicesResponseBodyServicesServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese.
	//
	// 	- en-US: English.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// SpringBoot-ECS容器版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetImage(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetLocale(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetName(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetShortDescription(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServicesResponseBodyServicesTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesResponseBodyServicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesTags) SetKey(v string) *ListServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListServicesResponseBodyServicesTags) SetValue(v string) *ListServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListSupplierRegistrationsRequest struct {
	// The filter.
	Filter []*ListSupplierRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Number of items per page in a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSupplierRegistrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSupplierRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsRequest) SetFilter(v []*ListSupplierRegistrationsRequestFilter) *ListSupplierRegistrationsRequest {
	s.Filter = v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetMaxResults(v int32) *ListSupplierRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetNextToken(v string) *ListSupplierRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetRegionId(v string) *ListSupplierRegistrationsRequest {
	s.RegionId = &v
	return s
}

type ListSupplierRegistrationsRequestFilter struct {
	// Name of the filter field. Allowed values:
	//
	// - SupplierUid: The aliUid of supplier.
	//
	// - SupplierName: The name of supplier.
	//
	// - RegistrationId: Registration ID.
	//
	// - Status: Registration status. Allowed values: Submitted, Approved, Rejected.
	//
	// example:
	//
	// SupplierUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Filter value.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListSupplierRegistrationsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListSupplierRegistrationsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsRequestFilter) SetName(v string) *ListSupplierRegistrationsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListSupplierRegistrationsRequestFilter) SetValue(v []*string) *ListSupplierRegistrationsRequestFilter {
	s.Value = v
	return s
}

type ListSupplierRegistrationsResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdx9kBO7qKpr9My/+XQo0oY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6CC568D-xxxx-xxxx-xxxx-08EB8E9F9F20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supplier registrations
	SupplierRegistrations []*ListSupplierRegistrationsResponseBodySupplierRegistrations `json:"SupplierRegistrations,omitempty" xml:"SupplierRegistrations,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSupplierRegistrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSupplierRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponseBody) SetMaxResults(v int32) *ListSupplierRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetNextToken(v string) *ListSupplierRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetRequestId(v string) *ListSupplierRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetSupplierRegistrations(v []*ListSupplierRegistrationsResponseBodySupplierRegistrations) *ListSupplierRegistrationsResponseBody {
	s.SupplierRegistrations = v
	return s
}

func (s *ListSupplierRegistrationsResponseBody) SetTotalCount(v int32) *ListSupplierRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSupplierRegistrationsResponseBodySupplierRegistrations struct {
	// The comment of this registration.
	//
	// example:
	//
	// empty
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Contact email
	//
	// example:
	//
	// test@163.com
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// Contact number
	//
	// example:
	//
	// 135xxxxxxxx
	ContactNumber *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	// Contact person
	//
	// example:
	//
	// Mike
	ContactPerson *string `json:"ContactPerson,omitempty" xml:"ContactPerson,omitempty"`
	// Contact person tiltle.
	//
	// example:
	//
	// CTO
	ContactPersonTitle *string `json:"ContactPersonTitle,omitempty" xml:"ContactPersonTitle,omitempty"`
	// Whether to enable the resell mode.
	//
	// example:
	//
	// false
	EnableResellerMode *bool `json:"EnableResellerMode,omitempty" xml:"EnableResellerMode,omitempty"`
	// Annual product revenue
	//
	// example:
	//
	// empty
	ProductAnnualRevenue *string `json:"ProductAnnualRevenue,omitempty" xml:"ProductAnnualRevenue,omitempty"`
	// The business of product.
	//
	// example:
	//
	// AI
	ProductBusiness *string `json:"ProductBusiness,omitempty" xml:"ProductBusiness,omitempty"`
	// Product delivery type，Valid values:
	//
	// SaaS
	//
	// License
	//
	// API
	//
	// DesktopSoftware
	//
	// Others
	//
	// example:
	//
	// SaaS
	ProductDeliveryTypes *string `json:"ProductDeliveryTypes,omitempty" xml:"ProductDeliveryTypes,omitempty"`
	// The publish time of product.
	//
	// example:
	//
	// 2024.10.24
	ProductPublishTime *string `json:"ProductPublishTime,omitempty" xml:"ProductPublishTime,omitempty"`
	// Product sell type, Valid values:
	//
	// - Direct
	//
	// - Channel
	//
	// example:
	//
	// Direct
	ProductSellTypes *string `json:"ProductSellTypes,omitempty" xml:"ProductSellTypes,omitempty"`
	// The registration ID.
	//
	// example:
	//
	// sr-xxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The description of resell business.
	//
	// example:
	//
	// empty
	ResellBusinessDesc *string `json:"ResellBusinessDesc,omitempty" xml:"ResellBusinessDesc,omitempty"`
	// The deployment state of the registration. Valid values:
	//
	// - Submitted
	//
	// - Approved
	//
	// - Rejected
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submit time of this registration.
	//
	// example:
	//
	// 2025-01-22 09:47:58
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/xxx/service-xxx/xxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Aliibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The english name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud ComputeNest
	SupplierNameEn *string `json:"SupplierNameEn,omitempty" xml:"SupplierNameEn,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 1256xxx23434
	SupplierUid *string `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// https://www.guangbao-uni.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s ListSupplierRegistrationsResponseBodySupplierRegistrations) String() string {
	return tea.Prettify(s)
}

func (s ListSupplierRegistrationsResponseBodySupplierRegistrations) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetComment(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.Comment = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactEmail(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactEmail = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactNumber(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactNumber = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactPerson(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactPerson = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetContactPersonTitle(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ContactPersonTitle = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetEnableResellerMode(v bool) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.EnableResellerMode = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductAnnualRevenue(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductAnnualRevenue = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductBusiness(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductBusiness = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductDeliveryTypes(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductDeliveryTypes = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductPublishTime(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductPublishTime = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetProductSellTypes(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ProductSellTypes = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetRegistrationId(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetResellBusinessDesc(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.ResellBusinessDesc = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetStatus(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.Status = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSubmitTime(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SubmitTime = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierDesc(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierDesc = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierLogo(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierLogo = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierName(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierName = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierNameEn(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierNameEn = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierUid(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierUid = &v
	return s
}

func (s *ListSupplierRegistrationsResponseBodySupplierRegistrations) SetSupplierUrl(v string) *ListSupplierRegistrationsResponseBodySupplierRegistrations {
	s.SupplierUrl = &v
	return s
}

type ListSupplierRegistrationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupplierRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupplierRegistrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSupplierRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsResponse) SetHeaders(v map[string]*string) *ListSupplierRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListSupplierRegistrationsResponse) SetStatusCode(v int32) *ListSupplierRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupplierRegistrationsResponse) SetBody(v *ListSupplierRegistrationsResponseBody) *ListSupplierRegistrationsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - service
	//
	// - serviceinstance
	//
	// - artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The tag keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8872ACE6-0297-54A4-8AAD-3A8623EC6C5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value:
	//
	// - service
	//
	// - serviceinstance
	//
	// - artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources and their tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// Resource ID
	//
	// example:
	//
	// service-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid value:
	//
	// - service
	//
	// - serviceinstance
	//
	// - artifact
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// usage
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfmTH5rcd4YFfob4P0uDAAc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - service
	//
	// - service instance
	//
	// - artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	// A pagination token.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0631D623-D917-1C2D-ACD6-5B3B19XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyServiceInstanceResourcesRequest struct {
	// The imported resources.
	//
	// example:
	//
	// {
	//
	//   "RegionId": "cn-hangzhou",
	//
	//   "Type": "ResourceIds",
	//
	//   "ResourceIds": {
	//
	//     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],
	//
	//     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],
	//
	//     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],
	//
	//     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]
	//
	//   }
	//
	// }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of operation performed on the service instance resource. Valid values:
	//
	// 	- Import: The resource is imported.
	//
	// 	- UnImport: The resource import is canceled.
	//
	// example:
	//
	// Import
	ServiceInstanceResourcesAction *string `json:"ServiceInstanceResourcesAction,omitempty" xml:"ServiceInstanceResourcesAction,omitempty"`
}

func (s ModifyServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesRequest) SetResources(v string) *ModifyServiceInstanceResourcesRequest {
	s.Resources = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ModifyServiceInstanceResourcesRequest) SetServiceInstanceResourcesAction(v string) *ModifyServiceInstanceResourcesRequest {
	s.ServiceInstanceResourcesAction = &v
	return s
}

type ModifyServiceInstanceResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponseBody) SetRequestId(v string) *ModifyServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ModifyServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetStatusCode(v int32) *ModifyServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceInstanceResourcesResponse) SetBody(v *ModifyServiceInstanceResourcesResponseBody) *ModifyServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type PreLaunchServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-f7024a22ea5149xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PreLaunchServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PreLaunchServiceRequest) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceRequest) SetClientToken(v string) *PreLaunchServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *PreLaunchServiceRequest) SetRegionId(v string) *PreLaunchServiceRequest {
	s.RegionId = &v
	return s
}

func (s *PreLaunchServiceRequest) SetServiceId(v string) *PreLaunchServiceRequest {
	s.ServiceId = &v
	return s
}

type PreLaunchServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreLaunchServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreLaunchServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceResponseBody) SetRequestId(v string) *PreLaunchServiceResponseBody {
	s.RequestId = &v
	return s
}

type PreLaunchServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreLaunchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreLaunchServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s PreLaunchServiceResponse) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceResponse) SetHeaders(v map[string]*string) *PreLaunchServiceResponse {
	s.Headers = v
	return s
}

func (s *PreLaunchServiceResponse) SetStatusCode(v int32) *PreLaunchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreLaunchServiceResponse) SetBody(v *PreLaunchServiceResponseBody) *PreLaunchServiceResponse {
	s.Body = v
	return s
}

type PushMeteringDataRequest struct {
	// The metering data. Parameters in the example value:
	//
	// 	- InstanceId: the ID of an instance in Alibaba Cloud Marketplace. Parameter type: STRING.
	//
	// 	- StartTime: the time when the metering operation started. Set the parameter to a UNIX timestamp. Unit: seconds. Parameter type: LONG.
	//
	// 	- EndTime: the time when the metering operation ended. Set the parameter to a UNIX timestamp. Unit: seconds. Parameter type: LONG.
	//
	// 	- Entities: the metering entities. Parameter type: LIST.
	//
	//     	- Key: the name of the metering item. Parameter type: STRING.
	//
	//         	- Frequency: the number of times the instance was used.
	//
	//         	- Period: the usage duration of the instance. Unit: seconds.
	//
	//     Note: The metering unit is second, whereas the billing unit is hour. Therefore, when bills are generated, seconds are converted to hours. For example, the usage metered from 19:00 to 20:00 is 1800 seconds and the price is USD 1 per hour. In this case, the hourly bill for 19:00 to 20:00 is calculated by using the following formula: 1800/3600 x 1 = 0.5. If the result is a decimal, only the first two decimal places are retained.
	//
	//           - Storage: The used storage space. Unit: bytes.
	//
	//            Note: The metering unit is byte, whereas the billing unit is MB. Therefore, when bills are generated, bytes are converted to megabytes. For example, the usage metered from 19:00 to 20:00 is 524,288 bytes and the price is USD 1 per MB. In this case, the hourly bill for 19:00 to 20:00 is calculated by using the following formula: 524288/1024/1024 x 1 = 0.5. If the result is a decimal, only the first two decimal places are retained.  - NetworkOut: the upstream traffic consumed. Unit: bit.
	//
	//            Note: The metering unit is bit, whereas the billing unit is Mbit. Therefore, when bills are generated, bits are converted to megabits. For example, the usage metered from 19:00 to 20:00 is 524,288 bits and the price is USD 1 per Mbit. In this case, the hourly bill for 19:00 to 20:00 is calculated by using the following formula: 524288/1024/1024 x 1 = 0.5. If the result is a decimal, only the first two decimal places are retained.  - NetworkIn: the downstream traffic consumed. Unit: bit.
	//
	//            Note: The metering unit is bit, whereas the billing unit is Mbit. Therefore, when bills are generated, bits are converted to megabits. For example, the usage metered from 19:00 to 20:00 is 524,288 bits and the price is USD 1 per Mbit. In this case, the hourly bill for 19:00 to 20:00 is calculated by using the following formula: 524288/1024/1024 x 1 = 0.5. If the result is a decimal, only the first two decimal places are retained.  - Character: the number of characters.
	//
	//           - DailyActiveUser: the number of daily active users (DAU).
	//
	//           - PeriodMin: the usage duration of the instance. Unit: minutes.  - VirtualCpu: the number of virtual CPU cores.
	//
	//     	- Value: the value of the metering item. The value is equal to or greater than 0. Parameter type: INTEGER.
	//
	// **Note**:
	//
	// 	- If bills are generated for the commodity in real time, the difference between the values of StartTime and EndTime is not limited. However, the time specified by EndTime must be later than that specified by StartTime.
	//
	// 	- If bills are generated for the commodity by billing cycle, such as by hour, by day, or by month, the difference between the values of StartTime and EndTime must be greater than 5 minutes.
	//
	// 	- In a request for pushing multiple metering data records, the values of InstanceId must indicate instances of the same commodity. You cannot push metering data of instances of multiple commodities at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"StartTime\\":1681264800,\\"EndTime\\":1681268400,\\"Entities\\":[{\\"Key\\":\\"Unit\\",\\"Value\\":\\"0\\"}]}]
	Metering *string `json:"Metering,omitempty" xml:"Metering,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s PushMeteringDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) SetMetering(v string) *PushMeteringDataRequest {
	s.Metering = &v
	return s
}

func (s *PushMeteringDataRequest) SetServiceInstanceId(v string) *PushMeteringDataRequest {
	s.ServiceInstanceId = &v
	return s
}

type PushMeteringDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMeteringDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
	s.Body = v
	return s
}

type RegisterServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-f7024a22ea5149xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s RegisterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceRequest) GoString() string {
	return s.String()
}

func (s *RegisterServiceRequest) SetClientToken(v string) *RegisterServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterServiceRequest) SetRegionId(v string) *RegisterServiceRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterServiceRequest) SetServiceId(v string) *RegisterServiceRequest {
	s.ServiceId = &v
	return s
}

type RegisterServiceResponseBody struct {
	// The registration ID.
	//
	// example:
	//
	// sr-72dd5071e90c40xxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-2713-52C8-AFFC-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponseBody) SetRegistrationId(v string) *RegisterServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterServiceResponseBody) SetRequestId(v string) *RegisterServiceResponseBody {
	s.RequestId = &v
	return s
}

type RegisterServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterServiceResponse) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponse) SetHeaders(v map[string]*string) *RegisterServiceResponse {
	s.Headers = v
	return s
}

func (s *RegisterServiceResponse) SetStatusCode(v int32) *RegisterServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterServiceResponse) SetBody(v *RegisterServiceResponseBody) *RegisterServiceResponse {
	s.Body = v
	return s
}

type RejectServiceUsageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Reject comments.
	//
	// example:
	//
	// Thanks for your application, please add your industry information.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-2117508c874c41xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// User ali uid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s RejectServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageRequest) SetClientToken(v string) *RejectServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *RejectServiceUsageRequest) SetComments(v string) *RejectServiceUsageRequest {
	s.Comments = &v
	return s
}

func (s *RejectServiceUsageRequest) SetServiceId(v string) *RejectServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *RejectServiceUsageRequest) SetType(v int32) *RejectServiceUsageRequest {
	s.Type = &v
	return s
}

func (s *RejectServiceUsageRequest) SetUserAliUid(v int64) *RejectServiceUsageRequest {
	s.UserAliUid = &v
	return s
}

type RejectServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageResponseBody) SetRequestId(v string) *RejectServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type RejectServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageResponse) SetHeaders(v map[string]*string) *RejectServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *RejectServiceUsageResponse) SetStatusCode(v int32) *RejectServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectServiceUsageResponse) SetBody(v *RejectServiceUsageResponseBody) *RejectServiceUsageResponse {
	s.Body = v
	return s
}

type ReleaseArtifactRequest struct {
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactRequest) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactRequest) SetArtifactId(v string) *ReleaseArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactRequest) SetClientToken(v string) *ReleaseArtifactRequest {
	s.ClientToken = &v
	return s
}

type ReleaseArtifactResponseBody struct {
	// The ID of the artifact.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The content of the artifact.
	//
	// example:
	//
	// "{\\"Url\\":\\"https://computenest-artifacts-draft-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/130920852836xxxx/cn-hangzhou/service-8072a04e5a134382xxxx/165095355xxxx/changes.txt\\",\\"ConfigurationMetadata\\":\\"{\\\\\\"WorkDir\\\\\\":\\\\\\"/root\\\\\\",\\\\\\"Platform\\\\\\":\\\\\\"Linux\\\\\\",\\\\\\"CommandType\\\\\\":\\\\\\"RunShellScript\\\\\\",\\\\\\"UpgradeScript\\\\\\":\\\\\\"cd /root\\\\\\\\ncp changes.txt cpchanges.txt\\\\\\\\nmv changes.txt mvchangge.txt\\\\\\"}\\"}"
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the artifact.
	//
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the artifact.
	//
	// example:
	//
	// draft
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the artifact was modified.
	//
	// example:
	//
	// 1650954178000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3818BA7D-3F50-1A44-9FF3-04A52A59XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the artifact. Valid values:
	//
	// 	- Created: The artifact is created.
	//
	// 	- Scanning: The artifact is being scanned.
	//
	// 	- ScanFailed: The artifact failed to be scanned.
	//
	// 	- Delivering: The artifact is being distributed.
	//
	// 	- Available: The artifact is available.
	//
	// 	- Deleted: The artifact is deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version name of the artifact.
	//
	// example:
	//
	// V1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ReleaseArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponseBody) SetArtifactId(v string) *ReleaseArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactProperty(v string) *ReleaseArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactType(v string) *ReleaseArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactVersion(v string) *ReleaseArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetDescription(v string) *ReleaseArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetGmtModified(v string) *ReleaseArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetRequestId(v string) *ReleaseArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetStatus(v string) *ReleaseArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetVersionName(v string) *ReleaseArtifactResponseBody {
	s.VersionName = &v
	return s
}

type ReleaseArtifactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponse) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponse) SetHeaders(v map[string]*string) *ReleaseArtifactResponse {
	s.Headers = v
	return s
}

func (s *ReleaseArtifactResponse) SetStatusCode(v int32) *ReleaseArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseArtifactResponse) SetBody(v *ReleaseArtifactResponseBody) *ReleaseArtifactResponse {
	s.Body = v
	return s
}

type RemoveServiceSharedAccountsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Whitelist accounts for service sharing.
	//
	// This parameter is required.
	UserAliUids []*int64 `json:"UserAliUids,omitempty" xml:"UserAliUids,omitempty" type:"Repeated"`
}

func (s RemoveServiceSharedAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *RemoveServiceSharedAccountsRequest) SetClientToken(v string) *RemoveServiceSharedAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetRegionId(v string) *RemoveServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetServiceId(v string) *RemoveServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetType(v string) *RemoveServiceSharedAccountsRequest {
	s.Type = &v
	return s
}

func (s *RemoveServiceSharedAccountsRequest) SetUserAliUids(v []*int64) *RemoveServiceSharedAccountsRequest {
	s.UserAliUids = v
	return s
}

type RemoveServiceSharedAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveServiceSharedAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServiceSharedAccountsResponseBody) SetRequestId(v string) *RemoveServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveServiceSharedAccountsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveServiceSharedAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *RemoveServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *RemoveServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *RemoveServiceSharedAccountsResponse) SetStatusCode(v int32) *RemoveServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveServiceSharedAccountsResponse) SetBody(v *RemoveServiceSharedAccountsResponseBody) *RemoveServiceSharedAccountsResponse {
	s.Body = v
	return s
}

type RestartServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RestartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceRequest) SetClientToken(v string) *RestartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetRegionId(v string) *RestartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetServiceInstanceId(v string) *RestartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RestartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponseBody) SetRequestId(v string) *RestartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponse) SetHeaders(v map[string]*string) *RestartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceInstanceResponse) SetStatusCode(v int32) *RestartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceInstanceResponse) SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse {
	s.Body = v
	return s
}

type RollbackServiceInstanceRequest struct {
	// Ensures idempotence of the request. Generate a value from your client to ensure it is unique across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service instance ID.
	//
	// You can obtain the service instance ID by calling [ListServiceInstances - Query Service Instance List](https://help.aliyun.com/document_detail/396200.html).
	//
	// example:
	//
	// si-3a8f9a75da074f52b969
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RollbackServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceRequest) SetClientToken(v string) *RollbackServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetRegionId(v string) *RollbackServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetServiceInstanceId(v string) *RollbackServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RollbackServiceInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// EE9EF87D-46F8-5AF6-9A65-6B034E204136
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service instance ID.
	//
	// example:
	//
	// si-5289e1d6d0c14397881d
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The deployment status of the service instance. Possible values:
	//
	// - Created: Created
	//
	// - Deploying: Deploying
	//
	// - DeployedFailed: Deployment Failed
	//
	// - Deployed: Deployed
	//
	// - Upgrading: Upgrading
	//
	// - UpgradeRollbacking: Rolling Back
	//
	// - Deleting: Deleting
	//
	// - Deleted: Deleted
	//
	// - DeletedFailed: Deletion Failed
	//
	// example:
	//
	// UpgradeRollbacking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RollbackServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponseBody) SetRequestId(v string) *RollbackServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackServiceInstanceResponseBody) SetServiceInstanceId(v string) *RollbackServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *RollbackServiceInstanceResponseBody) SetStatus(v string) *RollbackServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type RollbackServiceInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponse) SetHeaders(v map[string]*string) *RollbackServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RollbackServiceInstanceResponse) SetStatusCode(v int32) *RollbackServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackServiceInstanceResponse) SetBody(v *RollbackServiceInstanceResponseBody) *RollbackServiceInstanceResponse {
	s.Body = v
	return s
}

type StartServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceRequest) SetClientToken(v string) *StartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceInstanceRequest) SetRegionId(v string) *StartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartServiceInstanceRequest) SetServiceInstanceId(v string) *StartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2E91D771-0183-52CE-86CB-882D99B2CB77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponseBody) SetRequestId(v string) *StartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponse) SetHeaders(v map[string]*string) *StartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceInstanceResponse) SetStatusCode(v int32) *StartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceInstanceResponse) SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse {
	s.Body = v
	return s
}

type StopServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region id where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-c39ed4779cec449f****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StopServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceRequest) SetClientToken(v string) *StopServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopServiceInstanceRequest) SetRegionId(v string) *StopServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopServiceInstanceRequest) SetServiceInstanceId(v string) *StopServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StopServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 49A369EF-A302-5006-B0CE-94CED47C38CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponseBody) SetRequestId(v string) *StopServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponse) SetHeaders(v map[string]*string) *StopServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceInstanceResponse) SetStatusCode(v int32) *StopServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceInstanceResponse) SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value:
	//
	// - service
	//
	// - serviceinstance
	//
	// - artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you specify both the All and TagKey.N parameters, the All parameter does not take effect.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. valid value:
	//
	// - service
	//
	// - serviceinstance
	//
	// - artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. You can specify 1 to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpdateArtifactRequest struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildProperty *UpdateArtifactRequestArtifactBuildProperty `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty" type:"Struct"`
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	ArtifactProperty *UpdateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact. They can only change from Automatic to Public. Options:
	//
	// Public
	//
	// Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of the regions that support the deployment package.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequest) SetArtifactBuildProperty(v *UpdateArtifactRequestArtifactBuildProperty) *UpdateArtifactRequest {
	s.ArtifactBuildProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactId(v string) *UpdateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactProperty(v *UpdateArtifactRequestArtifactProperty) *UpdateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetClientToken(v string) *UpdateArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateArtifactRequest) SetDescription(v string) *UpdateArtifactRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactRequest) SetPermissionType(v string) *UpdateArtifactRequest {
	s.PermissionType = &v
	return s
}

func (s *UpdateArtifactRequest) SetSupportRegionIds(v []*string) *UpdateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactRequest) SetVersionName(v string) *UpdateArtifactRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactRequestArtifactBuildProperty struct {
	// The build arguments used during the image build process.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	BuildArgs []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The address of the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile or Buildpacks type.
	CodeRepo *UpdateArtifactRequestArtifactBuildPropertyCodeRepo `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty" type:"Struct"`
	// The command content.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// echo "start run command"
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances.
	//
	// 	- RunShellScript: shell command, applicable to Linux instances.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The relative path to the Dockerfile within the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	//
	// example:
	//
	// ./file/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// The region ID where the source mirror image is located.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The pull location of the source container image. This is used for the command docker pull ${SourceContainerImage}.
	//
	// >  This parameter is available only if the ArtifactBuildType is ContainerImage type.
	//
	// example:
	//
	// pytorch/pytorch:2.5.1-cuda12.4-cudnn9-devel
	SourceContainerImage *string `json:"SourceContainerImage,omitempty" xml:"SourceContainerImage,omitempty"`
	// The source image id. Supported Types:
	//
	// - Image ID: Pass the Image ID of the Ecs image directly.
	//
	// - OOS Common Parameter Name: Obtain the corresponding Image ID automatically by using the OOS common parameter name.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// Image ID：m-t4nhenrdc38pe4*****
	//
	// ubuntu_22_04_x64_20G_alibase_20240926.vhd
	//
	// OOS Common Parameter Name：aliyun/services/computenest/images/aliyun_3_2104_python_3_11
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildProperty) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetBuildArgs(v []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs) *UpdateArtifactRequestArtifactBuildProperty {
	s.BuildArgs = v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCodeRepo(v *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) *UpdateArtifactRequestArtifactBuildProperty {
	s.CodeRepo = v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCommandContent(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.CommandContent = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCommandType(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.CommandType = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetDockerfilePath(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.DockerfilePath = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetSourceContainerImage(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.SourceContainerImage = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetSourceImageId(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.SourceImageId = &v
	return s
}

type UpdateArtifactRequestArtifactBuildPropertyBuildArgs struct {
	// The name of a specific build argument.
	//
	// example:
	//
	// ENV
	ArgumentName *string `json:"ArgumentName,omitempty" xml:"ArgumentName,omitempty"`
	// The value of a specific build argument.
	//
	// example:
	//
	// nginx:latest
	ArgumentValue *string `json:"ArgumentValue,omitempty" xml:"ArgumentValue,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildPropertyBuildArgs) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildPropertyBuildArgs) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentName(v string) *UpdateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentName = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentValue(v string) *UpdateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentValue = &v
	return s
}

type UpdateArtifactRequestArtifactBuildPropertyCodeRepo struct {
	// The name of the branch in the code repository.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The endpoint.
	//
	// The URL address used to access the privately deployed GitLab instance.
	//
	// example:
	//
	// http://121.40.25.0
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// 455231
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The owner of the code repository.
	//
	// >  This parameter is available only if the git repository is private.
	//
	// example:
	//
	// aliyun-computenest
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The platform type. Valid values:
	//
	// - github
	//
	// - gitee
	//
	// - gitlab
	//
	// - codeup
	//
	// example:
	//
	// github
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The repository ID.
	//
	// example:
	//
	// 103
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// aliyun-computenest/quickstart-Lobexxx
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildPropertyCodeRepo) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetBranch(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Branch = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetEndpoint(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Endpoint = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetOrgId(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.OrgId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetOwner(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Owner = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetPlatform(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Platform = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoId(v int64) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoName(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoName = &v
	return s
}

type UpdateArtifactRequestArtifactProperty struct {
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity version of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// The image ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// m-0xij191j9cuev6ucxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Container Registry  repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// crr-yy4g68uhi39ttkm8
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the Container Registry repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// volcanosh/vc-webhook-manager
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The type of the repository.Valid values:
	//
	// 	- `Public`: a public repository.
	//
	// 	- `Private`: a private repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// Public
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The version tag of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The URL of the deployment package object.
	//
	//
	// > Note This parameter is available only if the deployment package is an file.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityCode(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetImageId(v string) *UpdateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoName(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoType(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoType = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetTag(v string) *UpdateArtifactRequestArtifactProperty {
	s.Tag = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetUrl(v string) *UpdateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type UpdateArtifactShrinkRequest struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact. They can only change from Automatic to Public. Options:
	//
	// Public
	//
	// Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of the regions that support the deployment package.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactShrinkRequest) SetArtifactBuildPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactBuildPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactId(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetClientToken(v string) *UpdateArtifactShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetDescription(v string) *UpdateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetPermissionType(v string) *UpdateArtifactShrinkRequest {
	s.PermissionType = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *UpdateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetVersionName(v string) *UpdateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponseBody struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- Created: The deployment package is created.
	//
	// 	- Scanning: The deployment package is being scanned.
	//
	// 	- ScanFailed: The deployment package failed to be scanned.
	//
	// 	- Delivering: The deployment package is being distributed.
	//
	// 	- Available: The deployment package is available.
	//
	// 	- Deleted: The deployment package is deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// [
	//
	// 			"cn-beijing",
	//
	// 			"cn-hangzhou",
	//
	// 			"cn-shanghai"
	//
	// 		]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponseBody) SetArtifactBuildProperty(v string) *UpdateArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactBuildType(v string) *UpdateArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactId(v string) *UpdateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactProperty(v string) *UpdateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactType(v string) *UpdateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactVersion(v string) *UpdateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetDescription(v string) *UpdateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetGmtModified(v string) *UpdateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetRequestId(v string) *UpdateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetStatus(v string) *UpdateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetStatusDetail(v string) *UpdateArtifactResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetSupportRegionIds(v string) *UpdateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetVersionName(v string) *UpdateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponse) SetHeaders(v map[string]*string) *UpdateArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactResponse) SetStatusCode(v int32) *UpdateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactResponse) SetBody(v *UpdateArtifactResponseBody) *UpdateArtifactResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// {\\"CmsTemplateId\\":1162921,\\"TemplateUrl\\":\\"https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1760465342xxxxxx/template/c072ef50-6c03-4d9c-8f0e-d1c440xxxxxx.json\\"}
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application. Valid values:
	//
	// 	- Manual: The application is manually approved.
	//
	// 	- AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType    *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity details.
	Commodity *UpdateServiceRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Compliance check metadata.
	ComplianceMetadata *UpdateServiceRequestComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The deployment configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"EstimateTime\\":null,\\"SupplierDeployMetadata\\":{\\"DeployTimeout\\":7200},\\"EnableVnc\\":false}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// terraform: The service is deployed by using Terraform.
	//
	// ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// spi: The service is deployed by calling a service provider interface (SPI).
	//
	// operation: The service is deployed by using a hosted O&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not update a service.
	//
	// 	- false: performs a dry run for the request, and update a service if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable the hosted O\\&M feature for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is required if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// Metering Item Configuration Information (Cloud Marketplace - Pay-As-You-Go Use)
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// Specifies whether to support distribution. Valid values:
	//
	// 	- false
	//
	// 	- true
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"Custom_Image_Ecs\\":{\\"EnablePrometheus\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Whether resell is supported.
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-1dda29c3eca648xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo []*UpdateServiceRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int32 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The update option.
	UpdateOption *UpdateServiceRequestUpdateOption `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty" type:"Struct"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The version name.
	//
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetAlarmMetadata(v string) *UpdateServiceRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetApprovalType(v string) *UpdateServiceRequest {
	s.ApprovalType = &v
	return s
}

func (s *UpdateServiceRequest) SetBuildParameters(v string) *UpdateServiceRequest {
	s.BuildParameters = &v
	return s
}

func (s *UpdateServiceRequest) SetClientToken(v string) *UpdateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceRequest) SetCommodity(v *UpdateServiceRequestCommodity) *UpdateServiceRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceRequest) SetComplianceMetadata(v *UpdateServiceRequestComplianceMetadata) *UpdateServiceRequest {
	s.ComplianceMetadata = v
	return s
}

func (s *UpdateServiceRequest) SetDeployMetadata(v string) *UpdateServiceRequest {
	s.DeployMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetDeployType(v string) *UpdateServiceRequest {
	s.DeployType = &v
	return s
}

func (s *UpdateServiceRequest) SetDryRun(v bool) *UpdateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceRequest) SetDuration(v int64) *UpdateServiceRequest {
	s.Duration = &v
	return s
}

func (s *UpdateServiceRequest) SetIsSupportOperated(v bool) *UpdateServiceRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *UpdateServiceRequest) SetLicenseMetadata(v string) *UpdateServiceRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetLogMetadata(v string) *UpdateServiceRequest {
	s.LogMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetOperationMetadata(v string) *UpdateServiceRequest {
	s.OperationMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetPolicyNames(v string) *UpdateServiceRequest {
	s.PolicyNames = &v
	return s
}

func (s *UpdateServiceRequest) SetRegionId(v string) *UpdateServiceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceRequest) SetResellable(v bool) *UpdateServiceRequest {
	s.Resellable = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceId(v string) *UpdateServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceInfo(v []*UpdateServiceRequestServiceInfo) *UpdateServiceRequest {
	s.ServiceInfo = v
	return s
}

func (s *UpdateServiceRequest) SetServiceType(v string) *UpdateServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateServiceRequest) SetServiceVersion(v string) *UpdateServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpdateServiceRequest) SetShareType(v string) *UpdateServiceRequest {
	s.ShareType = &v
	return s
}

func (s *UpdateServiceRequest) SetTenantType(v string) *UpdateServiceRequest {
	s.TenantType = &v
	return s
}

func (s *UpdateServiceRequest) SetTrialDuration(v int32) *UpdateServiceRequest {
	s.TrialDuration = &v
	return s
}

func (s *UpdateServiceRequest) SetUpdateOption(v *UpdateServiceRequestUpdateOption) *UpdateServiceRequest {
	s.UpdateOption = v
	return s
}

func (s *UpdateServiceRequest) SetUpgradeMetadata(v string) *UpdateServiceRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *UpdateServiceRequest) SetVersionName(v string) *UpdateServiceRequest {
	s.VersionName = &v
	return s
}

type UpdateServiceRequestCommodity struct {
	// This parameter is not available to the public.
	ComponentsMappings []*UpdateServiceRequestCommodityComponentsMappings `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
	// Metering entity extra information.
	MeteringEntityExtraInfos []*UpdateServiceRequestCommodityMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// Binding relationship between templates/specifications and metering dimensions (marketplace - PayAsYouGo)
	MeteringEntityMappings []*UpdateServiceRequestCommodityMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	// SaaS Boost configuration.
	//
	// example:
	//
	// {}
	SaasBoostConfig *string `json:"SaasBoostConfig,omitempty" xml:"SaasBoostConfig,omitempty"`
	// Product specifications and template/package mappings (Used in marketplace - subscription scenario)
	SpecificationMappings []*UpdateServiceRequestCommoditySpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodity) SetComponentsMappings(v []*UpdateServiceRequestCommodityComponentsMappings) *UpdateServiceRequestCommodity {
	s.ComponentsMappings = v
	return s
}

func (s *UpdateServiceRequestCommodity) SetMeteringEntityExtraInfos(v []*UpdateServiceRequestCommodityMeteringEntityExtraInfos) *UpdateServiceRequestCommodity {
	s.MeteringEntityExtraInfos = v
	return s
}

func (s *UpdateServiceRequestCommodity) SetMeteringEntityMappings(v []*UpdateServiceRequestCommodityMeteringEntityMappings) *UpdateServiceRequestCommodity {
	s.MeteringEntityMappings = v
	return s
}

func (s *UpdateServiceRequestCommodity) SetSaasBoostConfig(v string) *UpdateServiceRequestCommodity {
	s.SaasBoostConfig = &v
	return s
}

func (s *UpdateServiceRequestCommodity) SetSpecificationMappings(v []*UpdateServiceRequestCommoditySpecificationMappings) *UpdateServiceRequestCommodity {
	s.SpecificationMappings = v
	return s
}

type UpdateServiceRequestCommodityComponentsMappings struct {
	// This parameter is not available to the public.
	Mappings map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	// This parameter is not available to the public.
	//
	// example:
	//
	// This parameter is not available to the public.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommodityComponentsMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestCommodityComponentsMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityComponentsMappings) SetMappings(v map[string]*string) *UpdateServiceRequestCommodityComponentsMappings {
	s.Mappings = v
	return s
}

func (s *UpdateServiceRequestCommodityComponentsMappings) SetTemplateName(v string) *UpdateServiceRequestCommodityComponentsMappings {
	s.TemplateName = &v
	return s
}

type UpdateServiceRequestCommodityMeteringEntityExtraInfos struct {
	// Metering entity ID.
	//
	// example:
	//
	// cmgj0006xxxx-Memory-1
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// Metric name, required when type is ComputeNestBill or ComputeNestPrometheus.
	//
	// example:
	//
	// VirtualCpu/ecs.InstanceType
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Promql statement.
	//
	// example:
	//
	// avg_over_time(sum(rate(container_cpu_usage_seconds_total{namespace=~"ALIYUN::StackName"}[2m]))[1h:10s])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// Type. Valid values:
	//
	// - Custom
	//
	// - ComputeNestBill
	//
	// - ComputeNestPrometheus
	//
	// - ComputeNestTime
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateServiceRequestCommodityMeteringEntityExtraInfos) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestCommodityMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) SetEntityId(v string) *UpdateServiceRequestCommodityMeteringEntityExtraInfos {
	s.EntityId = &v
	return s
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) SetMetricName(v string) *UpdateServiceRequestCommodityMeteringEntityExtraInfos {
	s.MetricName = &v
	return s
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) SetPromql(v string) *UpdateServiceRequestCommodityMeteringEntityExtraInfos {
	s.Promql = &v
	return s
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) SetType(v string) *UpdateServiceRequestCommodityMeteringEntityExtraInfos {
	s.Type = &v
	return s
}

type UpdateServiceRequestCommodityMeteringEntityMappings struct {
	// Metering entity IDs.
	EntityIds []*string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty" type:"Repeated"`
	// The specification name.
	//
	// example:
	//
	// This parameter is not publicly accessible.
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// The service ID.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommodityMeteringEntityMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestCommodityMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) SetEntityIds(v []*string) *UpdateServiceRequestCommodityMeteringEntityMappings {
	s.EntityIds = v
	return s
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) SetSpecificationName(v string) *UpdateServiceRequestCommodityMeteringEntityMappings {
	s.SpecificationName = &v
	return s
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) SetTemplateName(v string) *UpdateServiceRequestCommodityMeteringEntityMappings {
	s.TemplateName = &v
	return s
}

type UpdateServiceRequestCommoditySpecificationMappings struct {
	// Specification code.
	//
	// example:
	//
	// yuncode5767800001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// Type, value：
	//
	// 	- **Custom**
	//
	// 	- **ComputeNestBill**
	//
	// 	- **ComputeNestPrometheus**
	//
	// 	- **ComputeNestTime**
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Product Specifications and Template/specification mapping Relationships (Cloud Marketplace - Subscription/Permanent Use)
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommoditySpecificationMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestCommoditySpecificationMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) SetSpecificationCode(v string) *UpdateServiceRequestCommoditySpecificationMappings {
	s.SpecificationCode = &v
	return s
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) SetSpecificationName(v string) *UpdateServiceRequestCommoditySpecificationMappings {
	s.SpecificationName = &v
	return s
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) SetTemplateName(v string) *UpdateServiceRequestCommoditySpecificationMappings {
	s.TemplateName = &v
	return s
}

type UpdateServiceRequestComplianceMetadata struct {
	// The compliance pack.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestComplianceMetadata) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestComplianceMetadata) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestComplianceMetadata) SetCompliancePacks(v []*string) *UpdateServiceRequestComplianceMetadata {
	s.CompliancePacks = v
	return s
}

type UpdateServiceRequestServiceInfo struct {
	// Protocol document information about the service.
	Agreements []*UpdateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The URL of the detailed description of the service.
	//
	// example:
	//
	// http://description.tidb.oss.url
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// The service name.
	//
	// example:
	//
	// Metric Name, filled in when Type is ComputeNestBill or ComputeNestPrometheus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// The URL of the detailed description of the service.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
	Softwares []*UpdateServiceRequestServiceInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfo) SetAgreements(v []*UpdateServiceRequestServiceInfoAgreements) *UpdateServiceRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetImage(v string) *UpdateServiceRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetLocale(v string) *UpdateServiceRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetLongDescriptionUrl(v string) *UpdateServiceRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetName(v string) *UpdateServiceRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetShortDescription(v string) *UpdateServiceRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *UpdateServiceRequestServiceInfo) SetSoftwares(v []*UpdateServiceRequestServiceInfoSoftwares) *UpdateServiceRequestServiceInfo {
	s.Softwares = v
	return s
}

type UpdateServiceRequestServiceInfoAgreements struct {
	// Protocol name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol url.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateServiceRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetName(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type UpdateServiceRequestServiceInfoSoftwares struct {
	// The name of the software.
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateServiceRequestServiceInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfoSoftwares) SetName(v string) *UpdateServiceRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoSoftwares) SetVersion(v string) *UpdateServiceRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

type UpdateServiceRequestUpdateOption struct {
	// Whether to update artifact.
	//
	// example:
	//
	// true
	UpdateArtifact *bool `json:"UpdateArtifact,omitempty" xml:"UpdateArtifact,omitempty"`
	// Update from. Valid values:
	//
	// - CODE
	//
	// - PARAMETERS
	//
	// example:
	//
	// PARAMETERS
	UpdateFrom *string `json:"UpdateFrom,omitempty" xml:"UpdateFrom,omitempty"`
}

func (s UpdateServiceRequestUpdateOption) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequestUpdateOption) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestUpdateOption) SetUpdateArtifact(v bool) *UpdateServiceRequestUpdateOption {
	s.UpdateArtifact = &v
	return s
}

func (s *UpdateServiceRequestUpdateOption) SetUpdateFrom(v string) *UpdateServiceRequestUpdateOption {
	s.UpdateFrom = &v
	return s
}

type UpdateServiceShrinkRequest struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// {\\"CmsTemplateId\\":1162921,\\"TemplateUrl\\":\\"https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1760465342xxxxxx/template/c072ef50-6c03-4d9c-8f0e-d1c440xxxxxx.json\\"}
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application. Valid values:
	//
	// 	- Manual: The application is manually approved.
	//
	// 	- AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType    *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity details.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// Compliance check metadata.
	ComplianceMetadataShrink *string `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty"`
	// The deployment configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"EstimateTime\\":null,\\"SupplierDeployMetadata\\":{\\"DeployTimeout\\":7200},\\"EnableVnc\\":false}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// terraform: The service is deployed by using Terraform.
	//
	// ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// spi: The service is deployed by calling a service provider interface (SPI).
	//
	// operation: The service is deployed by using a hosted O&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not update a service.
	//
	// 	- false: performs a dry run for the request, and update a service if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable the hosted O\\&M feature for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is required if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// Metering Item Configuration Information (Cloud Marketplace - Pay-As-You-Go Use)
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// Specifies whether to support distribution. Valid values:
	//
	// 	- false
	//
	// 	- true
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"Custom_Image_Ecs\\":{\\"EnablePrometheus\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Whether resell is supported.
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-1dda29c3eca648xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service details.
	ServiceInfo []*UpdateServiceShrinkRequestServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int32 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The update option.
	UpdateOptionShrink *string `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"xxx\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The version name.
	//
	// example:
	//
	// Draft
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequest) SetAlarmMetadata(v string) *UpdateServiceShrinkRequest {
	s.AlarmMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetApprovalType(v string) *UpdateServiceShrinkRequest {
	s.ApprovalType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetBuildParameters(v string) *UpdateServiceShrinkRequest {
	s.BuildParameters = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetClientToken(v string) *UpdateServiceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetCommodityShrink(v string) *UpdateServiceShrinkRequest {
	s.CommodityShrink = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetComplianceMetadataShrink(v string) *UpdateServiceShrinkRequest {
	s.ComplianceMetadataShrink = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDeployMetadata(v string) *UpdateServiceShrinkRequest {
	s.DeployMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDeployType(v string) *UpdateServiceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDryRun(v bool) *UpdateServiceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetDuration(v int64) *UpdateServiceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetIsSupportOperated(v bool) *UpdateServiceShrinkRequest {
	s.IsSupportOperated = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetLicenseMetadata(v string) *UpdateServiceShrinkRequest {
	s.LicenseMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetLogMetadata(v string) *UpdateServiceShrinkRequest {
	s.LogMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetOperationMetadata(v string) *UpdateServiceShrinkRequest {
	s.OperationMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetPolicyNames(v string) *UpdateServiceShrinkRequest {
	s.PolicyNames = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetRegionId(v string) *UpdateServiceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetResellable(v bool) *UpdateServiceShrinkRequest {
	s.Resellable = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceId(v string) *UpdateServiceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceInfo(v []*UpdateServiceShrinkRequestServiceInfo) *UpdateServiceShrinkRequest {
	s.ServiceInfo = v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceType(v string) *UpdateServiceShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetServiceVersion(v string) *UpdateServiceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetShareType(v string) *UpdateServiceShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetTenantType(v string) *UpdateServiceShrinkRequest {
	s.TenantType = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetTrialDuration(v int32) *UpdateServiceShrinkRequest {
	s.TrialDuration = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetUpdateOptionShrink(v string) *UpdateServiceShrinkRequest {
	s.UpdateOptionShrink = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetUpgradeMetadata(v string) *UpdateServiceShrinkRequest {
	s.UpgradeMetadata = &v
	return s
}

func (s *UpdateServiceShrinkRequest) SetVersionName(v string) *UpdateServiceShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateServiceShrinkRequestServiceInfo struct {
	// Protocol document information about the service.
	Agreements []*UpdateServiceShrinkRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The URL of the detailed description of the service.
	//
	// example:
	//
	// http://description.tidb.oss.url
	LongDescriptionUrl *string `json:"LongDescriptionUrl,omitempty" xml:"LongDescriptionUrl,omitempty"`
	// The service name.
	//
	// example:
	//
	// Metric Name, filled in when Type is ComputeNestBill or ComputeNestPrometheus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// The URL of the detailed description of the service.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
	Softwares []*UpdateServiceShrinkRequestServiceInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s UpdateServiceShrinkRequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetAgreements(v []*UpdateServiceShrinkRequestServiceInfoAgreements) *UpdateServiceShrinkRequestServiceInfo {
	s.Agreements = v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetImage(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Image = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetLocale(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Locale = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetLongDescriptionUrl(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.LongDescriptionUrl = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetName(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetShortDescription(v string) *UpdateServiceShrinkRequestServiceInfo {
	s.ShortDescription = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfo) SetSoftwares(v []*UpdateServiceShrinkRequestServiceInfoSoftwares) *UpdateServiceShrinkRequestServiceInfo {
	s.Softwares = v
	return s
}

type UpdateServiceShrinkRequestServiceInfoAgreements struct {
	// Protocol name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol url.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateServiceShrinkRequestServiceInfoAgreements) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetName(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

type UpdateServiceShrinkRequestServiceInfoSoftwares struct {
	// The name of the software.
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateServiceShrinkRequestServiceInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) SetName(v string) *UpdateServiceShrinkRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) SetVersion(v string) *UpdateServiceShrinkRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

type UpdateServiceResponseBody struct {
	// The dry run result.
	DryRunResult *UpdateServiceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetDryRunResult(v *UpdateServiceResponseBodyDryRunResult) *UpdateServiceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponseBodyDryRunResult struct {
	// The required ram policy for deploying role.
	RolePolicy *UpdateServiceResponseBodyDryRunResultRolePolicy `json:"RolePolicy,omitempty" xml:"RolePolicy,omitempty" type:"Struct"`
}

func (s UpdateServiceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResult) SetRolePolicy(v *UpdateServiceResponseBodyDryRunResultRolePolicy) *UpdateServiceResponseBodyDryRunResult {
	s.RolePolicy = v
	return s
}

type UpdateServiceResponseBodyDryRunResultRolePolicy struct {
	// The missing  ram policy for deploying role.
	MissingPolicy []*UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The required ram policy for deploying role.
	//
	// example:
	//
	// {
	//
	// 	"Statement": [{
	//
	// 		"Action": ["oos:CancelExecutions", "oos:DeleteExecutions", "oos:GetTemplate", "oos:ListExecutions", "oos:ListTemplates", "oos:NotifyExecution", "oos:StartExecution"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ram:PassRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ros:CreateStack", "ros:GetStack", "ros:UpdateStack", "ros:ListStackEvents", "ros:ListStackResources", "ros:ListStackResources", "ros:DeleteStack", "ram:GetRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}],
	//
	// 	"Version": "1"
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) SetMissingPolicy(v []*UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) *UpdateServiceResponseBodyDryRunResultRolePolicy {
	s.MissingPolicy = v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) SetPolicy(v string) *UpdateServiceResponseBodyDryRunResultRolePolicy {
	s.Policy = &v
	return s
}

type UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy struct {
	// The Actions.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The responses.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The service name.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetAction(v []*string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Action = v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetResource(v string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetServiceName(v string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.ServiceName = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceAttributeRequest struct {
	// The time when the service instance expires.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-25T02:28:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The License Data
	LicenseData *UpdateServiceInstanceAttributeRequestLicenseData `json:"LicenseData,omitempty" xml:"LicenseData,omitempty" type:"Struct"`
	// Application reason, currently used for trial application extension.
	//
	// example:
	//
	// \\"\\"
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-3df88e962cdexxxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequest) SetEndTime(v string) *UpdateServiceInstanceAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetLicenseData(v *UpdateServiceInstanceAttributeRequestLicenseData) *UpdateServiceInstanceAttributeRequest {
	s.LicenseData = v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetReason(v string) *UpdateServiceInstanceAttributeRequest {
	s.Reason = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetRegionId(v string) *UpdateServiceInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceAttributeRequestLicenseData struct {
	// The Custom Data
	//
	// example:
	//
	// {"Test"}
	CustomData *string `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
	// Mock response info.
	ResponseInfo *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo `json:"ResponseInfo,omitempty" xml:"ResponseInfo,omitempty" type:"Struct"`
}

func (s UpdateServiceInstanceAttributeRequestLicenseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequestLicenseData) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) SetCustomData(v string) *UpdateServiceInstanceAttributeRequestLicenseData {
	s.CustomData = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) SetResponseInfo(v *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) *UpdateServiceInstanceAttributeRequestLicenseData {
	s.ResponseInfo = v
	return s
}

type UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo struct {
	// Mock error code.
	//
	// example:
	//
	// EntityNotExist.Service
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Mock error message.
	//
	// example:
	//
	// The provided parameter "ServiceId" is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// if you want mock response, please open this option.
	//
	// example:
	//
	// true
	UpdateResponse *bool `json:"UpdateResponse,omitempty" xml:"UpdateResponse,omitempty"`
}

func (s UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetErrorCode(v string) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.ErrorCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetErrorMessage(v string) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetUpdateResponse(v bool) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.UpdateResponse = &v
	return s
}

type UpdateServiceInstanceAttributeShrinkRequest struct {
	// The time when the service instance expires.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-25T02:28:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The License Data
	LicenseDataShrink *string `json:"LicenseData,omitempty" xml:"LicenseData,omitempty"`
	// Application reason, currently used for trial application extension.
	//
	// example:
	//
	// \\"\\"
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-3df88e962cdexxxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetEndTime(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetLicenseDataShrink(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.LicenseDataShrink = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetReason(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.Reason = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetRegionId(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0CB2E0A9-B4DF-5C16-86AD-C511C483144B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponseBody) SetRequestId(v string) *UpdateServiceInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetStatusCode(v int32) *UpdateServiceInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetBody(v *UpdateServiceInstanceAttributeResponseBody) *UpdateServiceInstanceAttributeResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable Prometheus on the customer side. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {\\"EcsInstanceParameter\\":\\"4vCPU 8GiB\\",\\"ZoneId\\":\\"cn-heyuan-a\\",\\"SystemDiskSize\\":50,\\"DataDiskSize\\":150,\\"InternetMaxBandwidthOut\\":2,\\"RegionId\\":\\"cn-heyuan\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// package one
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable Prometheus on the customer side. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {\\"EcsInstanceParameter\\":\\"4vCPU 8GiB\\",\\"ZoneId\\":\\"cn-heyuan-a\\",\\"SystemDiskSize\\":50,\\"DataDiskSize\\":150,\\"InternetMaxBandwidthOut\\":2,\\"RegionId\\":\\"cn-heyuan\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// package one
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 2306175xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponseBody) SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse {
	s.Body = v
	return s
}

type UpdateServiceTestCaseRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service test case id.
	//
	// This parameter is required.
	//
	// example:
	//
	// stc-2deec15c20b24aaf9f16
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
	// The service test case name.
	//
	// This parameter is required.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// The service test config.
	//
	// This parameter is required.
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s UpdateServiceTestCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseRequest) SetRegionId(v string) *UpdateServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestCaseId(v string) *UpdateServiceTestCaseRequest {
	s.TestCaseId = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestCaseName(v string) *UpdateServiceTestCaseRequest {
	s.TestCaseName = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestConfig(v string) *UpdateServiceTestCaseRequest {
	s.TestConfig = &v
	return s
}

type UpdateServiceTestCaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DB1FA13E-1087-5654-84D5-58A0ACAD1B18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceTestCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseResponseBody) SetRequestId(v string) *UpdateServiceTestCaseResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceTestCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseResponse) SetHeaders(v map[string]*string) *UpdateServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceTestCaseResponse) SetStatusCode(v int32) *UpdateServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceTestCaseResponse) SetBody(v *UpdateServiceTestCaseResponseBody) *UpdateServiceTestCaseResponse {
	s.Body = v
	return s
}

type UpdateSharedAccountPermissionRequest struct {
	// Client token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Permission type. Possible values:
	//
	// - Deployable: Can be deployed.
	//
	// - Accessible: Can be accessed.
	//
	// - AccessibleIncludeBeta: Can access all versions, including Beta versions.
	//
	// - DeployableIncludeBeta: Can deploy all versions, including Beta versions.
	//
	// - Authorized: Authorized (for reselling scenarios)
	//
	// - Unauthorized: Unauthorized (for reselling scenarios)
	//
	// This parameter is required.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0d6e1d846e4c4axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service sharing type, with a default value of SharedAccount. Available options:
	//
	// - SharedAccount: Regular sharing type.
	//
	// - Reseller: Reselling sharing type.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Whitelist account for service sharing.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s UpdateSharedAccountPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSharedAccountPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionRequest) SetClientToken(v string) *UpdateSharedAccountPermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetPermission(v string) *UpdateSharedAccountPermissionRequest {
	s.Permission = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetRegionId(v string) *UpdateSharedAccountPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetServiceId(v string) *UpdateSharedAccountPermissionRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetType(v string) *UpdateSharedAccountPermissionRequest {
	s.Type = &v
	return s
}

func (s *UpdateSharedAccountPermissionRequest) SetUserAliUid(v int64) *UpdateSharedAccountPermissionRequest {
	s.UserAliUid = &v
	return s
}

type UpdateSharedAccountPermissionResponseBody struct {
	// RequestId
	//
	// example:
	//
	// C68B41B4-A646-5680-8A33-67884E3823A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSharedAccountPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSharedAccountPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionResponseBody) SetRequestId(v string) *UpdateSharedAccountPermissionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSharedAccountPermissionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSharedAccountPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSharedAccountPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSharedAccountPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionResponse) SetHeaders(v map[string]*string) *UpdateSharedAccountPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSharedAccountPermissionResponse) SetStatusCode(v int32) *UpdateSharedAccountPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSharedAccountPermissionResponse) SetBody(v *UpdateSharedAccountPermissionResponseBody) *UpdateSharedAccountPermissionResponse {
	s.Body = v
	return s
}

type UpdateSupplierInformationRequest struct {
	// The delivery settings.
	DeliverySettings *UpdateSupplierInformationRequestDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The Ip of operation.
	//
	// example:
	//
	// 192.xxx.xxx.xxx/16,192.xxx.xxx.xxx
	OperationIp *string `json:"OperationIp,omitempty" xml:"OperationIp,omitempty"`
	// The MFA of operation.
	//
	// example:
	//
	// true
	OperationMfaPresent *bool `json:"OperationMfaPresent,omitempty" xml:"OperationMfaPresent,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// http://example.aliyundoc.com/cover/34DB-4F4C-9373-003AA060****.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*UpdateSupplierInformationRequestSupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
}

func (s UpdateSupplierInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSupplierInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequest) SetDeliverySettings(v *UpdateSupplierInformationRequestDeliverySettings) *UpdateSupplierInformationRequest {
	s.DeliverySettings = v
	return s
}

func (s *UpdateSupplierInformationRequest) SetOperationIp(v string) *UpdateSupplierInformationRequest {
	s.OperationIp = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetOperationMfaPresent(v bool) *UpdateSupplierInformationRequest {
	s.OperationMfaPresent = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetRegionId(v string) *UpdateSupplierInformationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierDesc(v string) *UpdateSupplierInformationRequest {
	s.SupplierDesc = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierLogo(v string) *UpdateSupplierInformationRequest {
	s.SupplierLogo = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupplierUrl(v string) *UpdateSupplierInformationRequest {
	s.SupplierUrl = &v
	return s
}

func (s *UpdateSupplierInformationRequest) SetSupportContacts(v []*UpdateSupplierInformationRequestSupportContacts) *UpdateSupplierInformationRequest {
	s.SupportContacts = v
	return s
}

type UpdateSupplierInformationRequestDeliverySettings struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// mybucket
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to enable screencast delivery to Object Storage Service (OSS). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// path1/path2/
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s UpdateSupplierInformationRequestDeliverySettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateSupplierInformationRequestDeliverySettings) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssBucketName(v string) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssEnabled(v bool) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssExpirationDays(v int64) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *UpdateSupplierInformationRequestDeliverySettings) SetOssPath(v string) *UpdateSupplierInformationRequestDeliverySettings {
	s.OssPath = &v
	return s
}

type UpdateSupplierInformationRequestSupportContacts struct {
	// The type of  contact information
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of contact information
	//
	// example:
	//
	// supplier@example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSupplierInformationRequestSupportContacts) String() string {
	return tea.Prettify(s)
}

func (s UpdateSupplierInformationRequestSupportContacts) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationRequestSupportContacts) SetType(v string) *UpdateSupplierInformationRequestSupportContacts {
	s.Type = &v
	return s
}

func (s *UpdateSupplierInformationRequestSupportContacts) SetValue(v string) *UpdateSupplierInformationRequestSupportContacts {
	s.Value = &v
	return s
}

type UpdateSupplierInformationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s UpdateSupplierInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSupplierInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationResponseBody) SetRequestId(v string) *UpdateSupplierInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierDesc(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierName(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierName = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierUrl(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierUrl = &v
	return s
}

type UpdateSupplierInformationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSupplierInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSupplierInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSupplierInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationResponse) SetHeaders(v map[string]*string) *UpdateSupplierInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateSupplierInformationResponse) SetStatusCode(v int32) *UpdateSupplierInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSupplierInformationResponse) SetBody(v *UpdateSupplierInformationResponseBody) *UpdateSupplierInformationResponse {
	s.Body = v
	return s
}

type UpgradeServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceRequest) SetClientToken(v string) *UpgradeServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetDryRun(v string) *UpgradeServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetParameters(v map[string]interface{}) *UpgradeServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetRegionId(v string) *UpgradeServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceVersion(v string) *UpgradeServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceShrinkRequest) SetClientToken(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetDryRun(v string) *UpgradeServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetParametersShrink(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetRegionId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceVersion(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F224E002-AB0E-5FD9-A87E-54AEE56F6CAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The deployment state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters required for the upgrade.
	UpgradeRequiredParameters []*string `json:"UpgradeRequiredParameters,omitempty" xml:"UpgradeRequiredParameters,omitempty" type:"Repeated"`
}

func (s UpgradeServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponseBody) SetRequestId(v string) *UpgradeServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetServiceInstanceId(v string) *UpgradeServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetStatus(v string) *UpgradeServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetUpgradeRequiredParameters(v []*string) *UpgradeServiceInstanceResponseBody {
	s.UpgradeRequiredParameters = v
	return s
}

type UpgradeServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponse) SetHeaders(v map[string]*string) *UpgradeServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetStatusCode(v int32) *UpgradeServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetBody(v *UpgradeServiceInstanceResponseBody) *UpgradeServiceInstanceResponse {
	s.Body = v
	return s
}

type WithdrawServiceRequest struct {
	// Client token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s WithdrawServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceRequest) GoString() string {
	return s.String()
}

func (s *WithdrawServiceRequest) SetClientToken(v string) *WithdrawServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *WithdrawServiceRequest) SetRegionId(v string) *WithdrawServiceRequest {
	s.RegionId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceId(v string) *WithdrawServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceVersion(v string) *WithdrawServiceRequest {
	s.ServiceVersion = &v
	return s
}

type WithdrawServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WithdrawServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponseBody) SetRequestId(v string) *WithdrawServiceResponseBody {
	s.RequestId = &v
	return s
}

type WithdrawServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawServiceResponse) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponse) SetHeaders(v map[string]*string) *WithdrawServiceResponse {
	s.Headers = v
	return s
}

func (s *WithdrawServiceResponse) SetStatusCode(v int32) *WithdrawServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawServiceResponse) SetBody(v *WithdrawServiceResponseBody) *WithdrawServiceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenestsupplier"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
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
func (client *Client) AddServiceSharedAccountsWithOptions(request *AddServiceSharedAccountsRequest, runtime *util.RuntimeOptions) (_result *AddServiceSharedAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedAccounts)) {
		query["SharedAccounts"] = request.SharedAccounts
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServiceSharedAccounts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServiceSharedAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 商家通过服务使用请求
//
// @param request - ApproveServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveServiceUsageResponse
func (client *Client) ApproveServiceUsageWithOptions(request *ApproveServiceUsageRequest, runtime *util.RuntimeOptions) (_result *ApproveServiceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAliUid)) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveServiceUsage"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商家通过服务使用请求
//
// @param request - ApproveServiceUsageRequest
//
// @return ApproveServiceUsageResponse
func (client *Client) ApproveServiceUsage(request *ApproveServiceUsageRequest) (_result *ApproveServiceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CancelServiceRegistrationWithOptions(request *CancelServiceRegistrationRequest, runtime *util.RuntimeOptions) (_result *CancelServiceRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationId)) {
		query["RegistrationId"] = request.RegistrationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelServiceRegistration"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelServiceRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateArtifactWithOptions(tmpReq *CreateArtifactRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactBuildProperty)) {
		request.ArtifactBuildPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactBuildProperty, tea.String("ArtifactBuildProperty"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactBuildPropertyShrink)) {
		query["ArtifactBuildProperty"] = request.ArtifactBuildPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactBuildType)) {
		query["ArtifactBuildType"] = request.ArtifactBuildType
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Creates a service.
//
// @param tmpReq - CreateServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(tmpReq *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ComplianceMetadata)) {
		request.ComplianceMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComplianceMetadata, tea.String("ComplianceMetadata"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMetadata)) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovalType)) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !tea.BoolValue(util.IsUnset(request.BuildParameters)) {
		query["BuildParameters"] = request.BuildParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ComplianceMetadataShrink)) {
		query["ComplianceMetadata"] = request.ComplianceMetadataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeployMetadata)) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IsSupportOperated)) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseMetadata)) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.LogMetadata)) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyNames)) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resellable)) {
		query["Resellable"] = request.Resellable
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInfo)) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceServiceId)) {
		query["SourceServiceId"] = request.SourceServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceServiceVersion)) {
		query["SourceServiceVersion"] = request.SourceServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["TenantType"] = request.TenantType
	}

	if !tea.BoolValue(util.IsUnset(request.TrialDuration)) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeMetadata)) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateServiceTestCaseWithOptions(request *CreateServiceTestCaseRequest, runtime *util.RuntimeOptions) (_result *CreateServiceTestCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TestCaseName)) {
		query["TestCaseName"] = request.TestCaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TestConfig)) {
		query["TestConfig"] = request.TestConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceTestCase"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceTestCaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 创建服务测试任务
//
// @param request - CreateServiceTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceTestTaskResponse
func (client *Client) CreateServiceTestTaskWithOptions(request *CreateServiceTestTaskRequest, runtime *util.RuntimeOptions) (_result *CreateServiceTestTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskRegionId)) {
		query["TaskRegionId"] = request.TaskRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TestCaseIds)) {
		query["TestCaseIds"] = request.TestCaseIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceTestTask"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceTestTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务测试任务
//
// @param request - CreateServiceTestTaskRequest
//
// @return CreateServiceTestTaskResponse
func (client *Client) CreateServiceTestTask(request *CreateServiceTestTaskRequest) (_result *CreateServiceTestTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateServiceUsageWithOptions(request *CreateServiceUsageRequest, runtime *util.RuntimeOptions) (_result *CreateServiceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceUsage"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateSupplierRegistrationWithOptions(request *CreateSupplierRegistrationRequest, runtime *util.RuntimeOptions) (_result *CreateSupplierRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactEmail)) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !tea.BoolValue(util.IsUnset(request.ContactNumber)) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPerson)) {
		query["ContactPerson"] = request.ContactPerson
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPersonTitle)) {
		query["ContactPersonTitle"] = request.ContactPersonTitle
	}

	if !tea.BoolValue(util.IsUnset(request.EnableResellerMode)) {
		query["EnableResellerMode"] = request.EnableResellerMode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductAnnualRevenue)) {
		query["ProductAnnualRevenue"] = request.ProductAnnualRevenue
	}

	if !tea.BoolValue(util.IsUnset(request.ProductBusiness)) {
		query["ProductBusiness"] = request.ProductBusiness
	}

	if !tea.BoolValue(util.IsUnset(request.ProductDeliveryTypes)) {
		query["ProductDeliveryTypes"] = request.ProductDeliveryTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ProductPublishTime)) {
		query["ProductPublishTime"] = request.ProductPublishTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSellTypes)) {
		query["ProductSellTypes"] = request.ProductSellTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResellBusinessDesc)) {
		query["ResellBusinessDesc"] = request.ResellBusinessDesc
	}

	if !tea.BoolValue(util.IsUnset(request.Suggestion)) {
		query["Suggestion"] = request.Suggestion
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierDesc)) {
		query["SupplierDesc"] = request.SupplierDesc
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierLogo)) {
		query["SupplierLogo"] = request.SupplierLogo
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierName)) {
		query["SupplierName"] = request.SupplierName
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierNameEn)) {
		query["SupplierNameEn"] = request.SupplierNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierUrl)) {
		query["SupplierUrl"] = request.SupplierUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSupplierRegistration"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSupplierRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Deletes an artifact.
//
// @param request - DeleteArtifactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactResponse
func (client *Client) DeleteArtifactWithOptions(request *DeleteArtifactRequest, runtime *util.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 删除服务测试配置
//
// @param request - DeleteServiceTestCaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceTestCaseResponse
func (client *Client) DeleteServiceTestCaseWithOptions(request *DeleteServiceTestCaseRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceTestCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TestCaseId)) {
		query["TestCaseId"] = request.TestCaseId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceTestCase"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceTestCaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务测试配置
//
// @param request - DeleteServiceTestCaseRequest
//
// @return DeleteServiceTestCaseResponse
func (client *Client) DeleteServiceTestCase(request *DeleteServiceTestCaseRequest) (_result *DeleteServiceTestCaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeployServiceInstanceWithOptions(request *DeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GenerateDefaultServiceTestConfigWithOptions(request *GenerateDefaultServiceTestConfigRequest, runtime *util.RuntimeOptions) (_result *GenerateDefaultServiceTestConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDefaultServiceTestConfig"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDefaultServiceTestConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GenerateServicePolicyWithOptions(request *GenerateServicePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationTypes)) {
		query["OperationTypes"] = request.OperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateServicePolicy"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateServicePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetArtifactWithOptions(request *GetArtifactRequest, runtime *util.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactName)) {
		query["ArtifactName"] = request.ArtifactName
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetArtifactRepositoryCredentialsWithOptions(request *GetArtifactRepositoryCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactRepositoryCredentials"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Queries the information about a service.
//
// @param request - GetServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterAliUid)) {
		query["FilterAliUid"] = request.FilterAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SharedAccountType)) {
		query["SharedAccountType"] = request.SharedAccountType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetail)) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetServiceEstimateCostWithOptions(tmpReq *GetServiceEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceEstimateCostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Commodity)) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, tea.String("Commodity"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityShrink)) {
		query["Commodity"] = request.CommodityShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceEstimateCost"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 计算巢查询服务是否开通
//
// @param tmpReq - GetServiceProvisionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisionsWithOptions(tmpReq *GetServiceProvisionsRequest, runtime *util.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceProvisionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceProvisions"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 计算巢查询服务是否开通
//
// @param request - GetServiceProvisionsRequest
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisions(request *GetServiceProvisionsRequest) (_result *GetServiceProvisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetServiceRegistrationWithOptions(request *GetServiceRegistrationRequest, runtime *util.RuntimeOptions) (_result *GetServiceRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationId)) {
		query["RegistrationId"] = request.RegistrationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceRegistration"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Queries the valid values of parameters in a Resource Orchestration Service (ROS) template.
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *util.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePrivateVpcConnection)) {
		query["EnablePrivateVpcConnection"] = request.EnablePrivateVpcConnection
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTemplateParameterConstraints"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 获取服务测试任务中Cases执行情况
//
// @param request - GetServiceTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTestTaskResponse
func (client *Client) GetServiceTestTaskWithOptions(request *GetServiceTestTaskRequest, runtime *util.RuntimeOptions) (_result *GetServiceTestTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTestTask"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTestTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务测试任务中Cases执行情况
//
// @param request - GetServiceTestTaskRequest
//
// @return GetServiceTestTaskResponse
func (client *Client) GetServiceTestTask(request *GetServiceTestTaskRequest) (_result *GetServiceTestTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 获取服务商信息
//
// @param request - GetSupplierInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplierInformationResponse
func (client *Client) GetSupplierInformationWithOptions(request *GetSupplierInformationRequest, runtime *util.RuntimeOptions) (_result *GetSupplierInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSupplierInformation"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSupplierInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务商信息
//
// @param request - GetSupplierInformationRequest
//
// @return GetSupplierInformationResponse
func (client *Client) GetSupplierInformation(request *GetSupplierInformationRequest) (_result *GetSupplierInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetUploadCredentialsWithOptions(request *GetUploadCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetUploadCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadCredentials"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 上线服务
//
// @param request - LaunchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchServiceResponse
func (client *Client) LaunchServiceWithOptions(request *LaunchServiceRequest, runtime *util.RuntimeOptions) (_result *LaunchServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Categories)) {
		query["Categories"] = request.Categories
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Recommend)) {
		query["Recommend"] = request.Recommend
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LaunchService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LaunchServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线服务
//
// @param request - LaunchServiceRequest
//
// @return LaunchServiceResponse
func (client *Client) LaunchService(request *LaunchServiceRequest) (_result *LaunchServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListAcrImageRepositoriesWithOptions(request *ListAcrImageRepositoriesRequest, runtime *util.RuntimeOptions) (_result *ListAcrImageRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcrImageRepositories"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAcrImageRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListAcrImageTagsWithOptions(request *ListAcrImageTagsRequest, runtime *util.RuntimeOptions) (_result *ListAcrImageTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcrImageTags"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAcrImageTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// # Get the list of artifact security risks
//
// @param request - ListArtifactRisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactRisksResponse
func (client *Client) ListArtifactRisksWithOptions(request *ListArtifactRisksRequest, runtime *util.RuntimeOptions) (_result *ListArtifactRisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactRisks"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactRisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListArtifactVersionsWithOptions(tmpReq *ListArtifactVersionsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactVersionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListArtifactVersionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filters)) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, tea.String("Filters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.FiltersShrink)) {
		query["Filters"] = request.FiltersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactVersions"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListArtifactsWithOptions(request *ListArtifactsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifacts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// # Paginated query of distributor information list
//
// @param request - ListResellersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResellersResponse
func (client *Client) ListResellersWithOptions(request *ListResellersRequest, runtime *util.RuntimeOptions) (_result *ListResellersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResellers"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResellersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 展示服务实例账单
//
// @param request - ListServiceInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBillWithOptions(request *ListServiceInstanceBillRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingCycle)) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.BillingDate)) {
		query["BillingDate"] = request.BillingDate
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceBill"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示服务实例账单
//
// @param request - ListServiceInstanceBillRequest
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBill(request *ListServiceInstanceBillRequest) (_result *ListServiceInstanceBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 查询服务实例部署详情
//
// @param request - ListServiceInstanceDeployDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceDeployDetailsResponse
func (client *Client) ListServiceInstanceDeployDetailsWithOptions(request *ListServiceInstanceDeployDetailsRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceDeployDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CycleTimeZone)) {
		query["CycleTimeZone"] = request.CycleTimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.CycleType)) {
		query["CycleType"] = request.CycleType
	}

	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		query["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceDeployDetails"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceDeployDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务实例部署详情
//
// @param request - ListServiceInstanceDeployDetailsRequest
//
// @return ListServiceInstanceDeployDetailsResponse
func (client *Client) ListServiceInstanceDeployDetails(request *ListServiceInstanceDeployDetailsRequest) (_result *ListServiceInstanceDeployDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		query["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["Logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceLogs"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceResourceType)) {
		query["ServiceInstanceResourceType"] = request.ServiceInstanceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceInstanceUpgradeHistoryWithOptions(request *ListServiceInstanceUpgradeHistoryRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceUpgradeHistory"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceUpgradeHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDeleted)) {
		query["ShowDeleted"] = request.ShowDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceRegistrationsWithOptions(request *ListServiceRegistrationsRequest, runtime *util.RuntimeOptions) (_result *ListServiceRegistrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceRegistrations"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceRegistrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceSharedAccountsWithOptions(request *ListServiceSharedAccountsRequest, runtime *util.RuntimeOptions) (_result *ListServiceSharedAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		query["Permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceSharedAccounts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceSharedAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 服务测试用例列表
//
// @param request - ListServiceTestCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTestCasesResponse
func (client *Client) ListServiceTestCasesWithOptions(request *ListServiceTestCasesRequest, runtime *util.RuntimeOptions) (_result *ListServiceTestCasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceTestCases"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceTestCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务测试用例列表
//
// @param request - ListServiceTestCasesRequest
//
// @return ListServiceTestCasesResponse
func (client *Client) ListServiceTestCases(request *ListServiceTestCasesRequest) (_result *ListServiceTestCasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 获取服务测试实时日志
//
// @param request - ListServiceTestTaskLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceTestTaskLogsResponse
func (client *Client) ListServiceTestTaskLogsWithOptions(request *ListServiceTestTaskLogsRequest, runtime *util.RuntimeOptions) (_result *ListServiceTestTaskLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceTestTaskLogs"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceTestTaskLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务测试实时日志
//
// @param request - ListServiceTestTaskLogsRequest
//
// @return ListServiceTestTaskLogsResponse
func (client *Client) ListServiceTestTaskLogs(request *ListServiceTestTaskLogsRequest) (_result *ListServiceTestTaskLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceTestTasksWithOptions(request *ListServiceTestTasksRequest, runtime *util.RuntimeOptions) (_result *ListServiceTestTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceTestTasks"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceTestTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *util.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceUsages"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllVersions)) {
		query["AllVersions"] = request.AllVersions
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 查询服务商入职审核列表
//
// @param request - ListSupplierRegistrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupplierRegistrationsResponse
func (client *Client) ListSupplierRegistrationsWithOptions(request *ListSupplierRegistrationsRequest, runtime *util.RuntimeOptions) (_result *ListSupplierRegistrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSupplierRegistrations"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSupplierRegistrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务商入职审核列表
//
// @param request - ListSupplierRegistrationsRequest
//
// @return ListSupplierRegistrationsResponse
func (client *Client) ListSupplierRegistrations(request *ListSupplierRegistrationsRequest) (_result *ListSupplierRegistrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 查询标签键列表
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签键列表
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 查询资源标签
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源标签
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 查询标签值列表
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签值列表
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ModifyServiceInstanceResourcesWithOptions(request *ModifyServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ModifyServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceResourcesAction)) {
		query["ServiceInstanceResourcesAction"] = request.ServiceInstanceResourcesAction
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyServiceInstanceResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 预发布服务
//
// @param request - PreLaunchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreLaunchServiceResponse
func (client *Client) PreLaunchServiceWithOptions(request *PreLaunchServiceRequest, runtime *util.RuntimeOptions) (_result *PreLaunchServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreLaunchService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreLaunchServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预发布服务
//
// @param request - PreLaunchServiceRequest
//
// @return PreLaunchServiceResponse
func (client *Client) PreLaunchService(request *PreLaunchServiceRequest) (_result *PreLaunchServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *util.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Metering)) {
		query["Metering"] = request.Metering
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMeteringData"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) RegisterServiceWithOptions(request *RegisterServiceRequest, runtime *util.RuntimeOptions) (_result *RegisterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) RejectServiceUsageWithOptions(request *RejectServiceUsageRequest, runtime *util.RuntimeOptions) (_result *RejectServiceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAliUid)) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectServiceUsage"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ReleaseArtifactWithOptions(request *ReleaseArtifactRequest, runtime *util.RuntimeOptions) (_result *ReleaseArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) RemoveServiceSharedAccountsWithOptions(request *RemoveServiceSharedAccountsRequest, runtime *util.RuntimeOptions) (_result *RemoveServiceSharedAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAliUids)) {
		query["UserAliUids"] = request.UserAliUids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveServiceSharedAccounts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveServiceSharedAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) RollbackServiceInstanceWithOptions(request *RollbackServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RollbackServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 给资源打标签
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给资源打标签
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 资源解绑标签
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源解绑标签
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateArtifactWithOptions(tmpReq *UpdateArtifactRequest, runtime *util.RuntimeOptions) (_result *UpdateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactBuildProperty)) {
		request.ArtifactBuildPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactBuildProperty, tea.String("ArtifactBuildProperty"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactBuildPropertyShrink)) {
		query["ArtifactBuildProperty"] = request.ArtifactBuildPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionType)) {
		query["PermissionType"] = request.PermissionType
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateServiceWithOptions(tmpReq *UpdateServiceRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Commodity)) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, tea.String("Commodity"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ComplianceMetadata)) {
		request.ComplianceMetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComplianceMetadata, tea.String("ComplianceMetadata"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateOption)) {
		request.UpdateOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateOption, tea.String("UpdateOption"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmMetadata)) {
		query["AlarmMetadata"] = request.AlarmMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovalType)) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !tea.BoolValue(util.IsUnset(request.BuildParameters)) {
		query["BuildParameters"] = request.BuildParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityShrink)) {
		query["Commodity"] = request.CommodityShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ComplianceMetadataShrink)) {
		query["ComplianceMetadata"] = request.ComplianceMetadataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeployMetadata)) {
		query["DeployMetadata"] = request.DeployMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IsSupportOperated)) {
		query["IsSupportOperated"] = request.IsSupportOperated
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseMetadata)) {
		query["LicenseMetadata"] = request.LicenseMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.LogMetadata)) {
		query["LogMetadata"] = request.LogMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyNames)) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resellable)) {
		query["Resellable"] = request.Resellable
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInfo)) {
		query["ServiceInfo"] = request.ServiceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantType)) {
		query["TenantType"] = request.TenantType
	}

	if !tea.BoolValue(util.IsUnset(request.TrialDuration)) {
		query["TrialDuration"] = request.TrialDuration
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateOptionShrink)) {
		query["UpdateOption"] = request.UpdateOptionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeMetadata)) {
		query["UpgradeMetadata"] = request.UpgradeMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateServiceInstanceAttributeWithOptions(tmpReq *UpdateServiceInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LicenseData)) {
		request.LicenseDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LicenseData, tea.String("LicenseData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseDataShrink)) {
		query["LicenseData"] = request.LicenseDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceAttribute"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PredefinedParametersName)) {
		query["PredefinedParametersName"] = request.PredefinedParametersName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceSpec"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 修改服务测试用例
//
// @param request - UpdateServiceTestCaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceTestCaseResponse
func (client *Client) UpdateServiceTestCaseWithOptions(request *UpdateServiceTestCaseRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceTestCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TestCaseId)) {
		query["TestCaseId"] = request.TestCaseId
	}

	if !tea.BoolValue(util.IsUnset(request.TestCaseName)) {
		query["TestCaseName"] = request.TestCaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TestConfig)) {
		query["TestConfig"] = request.TestConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceTestCase"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceTestCaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改服务测试用例
//
// @param request - UpdateServiceTestCaseRequest
//
// @return UpdateServiceTestCaseResponse
func (client *Client) UpdateServiceTestCase(request *UpdateServiceTestCaseRequest) (_result *UpdateServiceTestCaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateSharedAccountPermissionWithOptions(request *UpdateSharedAccountPermissionRequest, runtime *util.RuntimeOptions) (_result *UpdateSharedAccountPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Permission)) {
		query["Permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAliUid)) {
		query["UserAliUid"] = request.UserAliUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSharedAccountPermission"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSharedAccountPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// 更新供应商全局信息
//
// @param request - UpdateSupplierInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSupplierInformationResponse
func (client *Client) UpdateSupplierInformationWithOptions(request *UpdateSupplierInformationRequest, runtime *util.RuntimeOptions) (_result *UpdateSupplierInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverySettings)) {
		query["DeliverySettings"] = request.DeliverySettings
	}

	if !tea.BoolValue(util.IsUnset(request.OperationIp)) {
		query["OperationIp"] = request.OperationIp
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMfaPresent)) {
		query["OperationMfaPresent"] = request.OperationMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierDesc)) {
		query["SupplierDesc"] = request.SupplierDesc
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierLogo)) {
		query["SupplierLogo"] = request.SupplierLogo
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierUrl)) {
		query["SupplierUrl"] = request.SupplierUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SupportContacts)) {
		query["SupportContacts"] = request.SupportContacts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSupplierInformation"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSupplierInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新供应商全局信息
//
// @param request - UpdateSupplierInformationRequest
//
// @return UpdateSupplierInformationResponse
func (client *Client) UpdateSupplierInformation(request *UpdateSupplierInformationRequest) (_result *UpdateSupplierInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpgradeServiceInstanceWithOptions(tmpReq *UpgradeServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradeServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) WithdrawServiceWithOptions(request *WithdrawServiceRequest, runtime *util.RuntimeOptions) (_result *WithdrawServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WithdrawService"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
	_result = &WithdrawServiceResponse{}
	_body, _err := client.WithdrawServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
