// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMetadata(v string) *GetServiceResponseBody
	GetAlarmMetadata() *string
	SetApprovalType(v string) *GetServiceResponseBody
	GetApprovalType() *string
	SetBuildInfo(v string) *GetServiceResponseBody
	GetBuildInfo() *string
	SetBuildParameters(v string) *GetServiceResponseBody
	GetBuildParameters() *string
	SetCategories(v string) *GetServiceResponseBody
	GetCategories() *string
	SetCommodity(v *GetServiceResponseBodyCommodity) *GetServiceResponseBody
	GetCommodity() *GetServiceResponseBodyCommodity
	SetComplianceMetadata(v *GetServiceResponseBodyComplianceMetadata) *GetServiceResponseBody
	GetComplianceMetadata() *GetServiceResponseBodyComplianceMetadata
	SetCreateTime(v string) *GetServiceResponseBody
	GetCreateTime() *string
	SetCrossRegionConnectionStatus(v string) *GetServiceResponseBody
	GetCrossRegionConnectionStatus() *string
	SetDeployMetadata(v string) *GetServiceResponseBody
	GetDeployMetadata() *string
	SetDeployType(v string) *GetServiceResponseBody
	GetDeployType() *string
	SetDuration(v int64) *GetServiceResponseBody
	GetDuration() *int64
	SetEntitySource(v map[string]*string) *GetServiceResponseBody
	GetEntitySource() map[string]*string
	SetIsSupportOperated(v bool) *GetServiceResponseBody
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *GetServiceResponseBody
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *GetServiceResponseBody
	GetLogMetadata() *string
	SetOperationMetadata(v string) *GetServiceResponseBody
	GetOperationMetadata() *string
	SetPayFromType(v string) *GetServiceResponseBody
	GetPayFromType() *string
	SetPermission(v string) *GetServiceResponseBody
	GetPermission() *string
	SetPolicyNames(v string) *GetServiceResponseBody
	GetPolicyNames() *string
	SetProgress(v int64) *GetServiceResponseBody
	GetProgress() *int64
	SetPublishTime(v string) *GetServiceResponseBody
	GetPublishTime() *string
	SetRegistrationId(v string) *GetServiceResponseBody
	GetRegistrationId() *string
	SetRequestId(v string) *GetServiceResponseBody
	GetRequestId() *string
	SetResellable(v bool) *GetServiceResponseBody
	GetResellable() *bool
	SetResourceGroupId(v string) *GetServiceResponseBody
	GetResourceGroupId() *string
	SetSecretKey(v string) *GetServiceResponseBody
	GetSecretKey() *string
	SetServiceAuditDocumentUrl(v string) *GetServiceResponseBody
	GetServiceAuditDocumentUrl() *string
	SetServiceDiscoverable(v string) *GetServiceResponseBody
	GetServiceDiscoverable() *string
	SetServiceDocumentInfos(v []*GetServiceResponseBodyServiceDocumentInfos) *GetServiceResponseBody
	GetServiceDocumentInfos() []*GetServiceResponseBodyServiceDocumentInfos
	SetServiceId(v string) *GetServiceResponseBody
	GetServiceId() *string
	SetServiceInfos(v []*GetServiceResponseBodyServiceInfos) *GetServiceResponseBody
	GetServiceInfos() []*GetServiceResponseBodyServiceInfos
	SetServiceLocaleConfigs(v []*GetServiceResponseBodyServiceLocaleConfigs) *GetServiceResponseBody
	GetServiceLocaleConfigs() []*GetServiceResponseBodyServiceLocaleConfigs
	SetServiceProductUrl(v string) *GetServiceResponseBody
	GetServiceProductUrl() *string
	SetServiceType(v string) *GetServiceResponseBody
	GetServiceType() *string
	SetShareType(v string) *GetServiceResponseBody
	GetShareType() *string
	SetShareTypeStatus(v string) *GetServiceResponseBody
	GetShareTypeStatus() *string
	SetSourceServiceId(v string) *GetServiceResponseBody
	GetSourceServiceId() *string
	SetSourceServiceVersion(v string) *GetServiceResponseBody
	GetSourceServiceVersion() *string
	SetSourceSupplierName(v string) *GetServiceResponseBody
	GetSourceSupplierName() *string
	SetStatistic(v *GetServiceResponseBodyStatistic) *GetServiceResponseBody
	GetStatistic() *GetServiceResponseBodyStatistic
	SetStatus(v string) *GetServiceResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *GetServiceResponseBody
	GetStatusDetail() *string
	SetSupplierName(v string) *GetServiceResponseBody
	GetSupplierName() *string
	SetSupplierUrl(v string) *GetServiceResponseBody
	GetSupplierUrl() *string
	SetSupportContacts(v []*GetServiceResponseBodySupportContacts) *GetServiceResponseBody
	GetSupportContacts() []*GetServiceResponseBodySupportContacts
	SetTags(v []*GetServiceResponseBodyTags) *GetServiceResponseBody
	GetTags() []*GetServiceResponseBodyTags
	SetTenantType(v string) *GetServiceResponseBody
	GetTenantType() *string
	SetTestStatus(v string) *GetServiceResponseBody
	GetTestStatus() *string
	SetTrialDuration(v int64) *GetServiceResponseBody
	GetTrialDuration() *int64
	SetTrialType(v string) *GetServiceResponseBody
	GetTrialType() *string
	SetUpdateTime(v string) *GetServiceResponseBody
	GetUpdateTime() *string
	SetUpgradeMetadata(v string) *GetServiceResponseBody
	GetUpgradeMetadata() *string
	SetVersion(v string) *GetServiceResponseBody
	GetVersion() *string
	SetVersionName(v string) *GetServiceResponseBody
	GetVersionName() *string
	SetVirtualInternetService(v string) *GetServiceResponseBody
	GetVirtualInternetService() *string
	SetVirtualInternetServiceId(v string) *GetServiceResponseBody
	GetVirtualInternetServiceId() *string
}

type GetServiceResponseBody struct {
	// The alert configurations of the service.
	//
	// example:
	//
	// {
	//
	//   "TemplateUrl": "http://template.file.url",
	//
	//   // Application group level alarm metadata
	//
	//   "ApplicationGroups": [
	//
	//     {
	//
	//       "Name": "applicationGroup1",
	//
	//       "TemplateUrl": "url1"
	//
	//     },
	//
	//     {
	//
	//       "Name": "applicationGroup2",
	//
	//       "TemplateUrl": "url2"
	//
	//     }
	//
	//   ]
	//
	// }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type of the service usage application.
	//
	// Valid values:
	//
	// - Manual: The application is manually approved.
	//
	// - AutoPass: The application is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The information about the build service.
	//
	// example:
	//
	// { "RepoUrl": "https://github.com/user/repo.git", "Brancn": "main"}
	BuildInfo *string `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	// The parameters of the build service.
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// The service category.
	//
	// example:
	//
	// DevOps
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity specifications.
	Commodity *GetServiceResponseBodyCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The compliance check metadata.
	ComplianceMetadata *GetServiceResponseBodyComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The time when the service was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The binding relationship of the commodity module.
	//
	// example:
	//
	// componesConfigs
	CrossRegionConnectionStatus *string `json:"CrossRegionConnectionStatus,omitempty" xml:"CrossRegionConnectionStatus,omitempty"`
	// The information about the service deployment configuration.
	//
	// The information varies based on the deployment type. Different deployment types use different data formats. Therefore, the information is stored as a JSON string.
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
	// The deployment type.
	//
	// Valid values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - spi: The service is deployed by calling SPI.
	//
	// - operation: The service is deployed using Alibaba Cloud Managed Services.
	//
	// - container: The service is deployed using containers.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The duration of the Alibaba Cloud Managed Services. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The source of the reported data.
	EntitySource map[string]*string `json:"EntitySource,omitempty" xml:"EntitySource,omitempty"`
	// Indicates whether Alibaba Cloud Managed Services is enabled.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {\\"PayType\\":\\"CustomFixTime\\",\\"DefaultLicenseDays\\":7,\\"CustomMetadata\\":[{\\"TemplateName\\":\\" template1\\",\\"SpecificationName\\":\\"bandwith-0\\",\\"CustomData\\":\\"1\\"}]}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The application log configurations.
	//
	// example:
	//
	// {
	//
	//   "Logstores": [
	//
	//     {
	//
	//     "LogstoreName": "access-log",
	//
	//   "LogPath": "/home/admin/app/logs", # Not required for containers. Configure in YAML
	//
	//   "FilePattern": "access.log*" # Not required for containers. Configure in YAML
	//
	//     }
	//
	//   ]
	//
	// }
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The configurations of Alibaba Cloud Managed Services.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The payment source.
	//
	// Valid values:
	//
	// - None: The service is free of charge.
	//
	// - Marketplace: The service is paid on Alibaba Cloud Marketplace.
	//
	// - Custom: The service is paid using a custom payment method.
	//
	// example:
	//
	// None
	PayFromType *string `json:"PayFromType,omitempty" xml:"PayFromType,omitempty"`
	// The permission type.
	//
	// Valid values:
	//
	// - Deployable: The service can be deployed.
	//
	// - Accessible: The service can be accessed.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The policy names.
	//
	// A policy name can be up to 128 characters in length. Multiple policy names are separated by commas (,). Only policies related to Alibaba Cloud Managed Services are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The deployment progress of the service instance. Unit: %.
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
	// sr-1b4aabc1c9eb4109****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AC8E73E-88DE-52C2-A29B-531FC13A5604
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the service can be distributed.
	//
	// Valid values:
	//
	// - false: The service cannot be distributed.
	//
	// - true: The service can be distributed.
	//
	// example:
	//
	// false
	Resellable *bool `json:"Resellable,omitempty" xml:"Resellable,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm2jfvb7b****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service key. It is used for digital signature encryption.
	//
	// example:
	//
	// 6cfc5d4649c54216****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The URL of the service review file.
	ServiceAuditDocumentUrl *string `json:"ServiceAuditDocumentUrl,omitempty" xml:"ServiceAuditDocumentUrl,omitempty"`
	// Indicates whether the service is discoverable.
	//
	// Valid values:
	//
	// - INVISIBLE: The service is not discoverable.
	//
	// - DISCOVERABLE: The service is discoverable.
	//
	// example:
	//
	// DISCOVERABLE
	ServiceDiscoverable *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// The service document information.
	ServiceDocumentInfos []*GetServiceResponseBodyServiceDocumentInfos `json:"ServiceDocumentInfos,omitempty" xml:"ServiceDocumentInfos,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// service-ca83ff3cb6b14dbc****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service information.
	ServiceInfos []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The multi-language configurations of the service.
	ServiceLocaleConfigs []*GetServiceResponseBodyServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
	// The URL of the product page.
	//
	// example:
	//
	// http://example2.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The service type.
	//
	// Valid values:
	//
	// - private: The service is deployed in the user\\"s account.
	//
	// - managed: The service is deployed in the service provider\\"s account.
	//
	// - operation: The service is an Alibaba Cloud Managed Service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The sharing type.
	//
	// Valid values:
	//
	// - Public: The service is public. Formal and trial deployments are not restricted.
	//
	// - Restricted: The service is restricted. Formal and trial deployments are restricted.
	//
	// - OnlyFormalRestricted: Only formal deployments are restricted.
	//
	// - OnlyTrailRestricted: Only trial deployments are restricted.
	//
	// - Hidden: The service is hidden. It is not visible and you cannot apply for deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The sharing status of the service. 	Notice: This parameter is deprecated.
	//
	// example:
	//
	// This parameter is deprecated.
	ShareTypeStatus *string `json:"ShareTypeStatus,omitempty" xml:"ShareTypeStatus,omitempty"`
	// The ID of the source service for distribution.
	//
	// example:
	//
	// service-70a3b15bb6264315****
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The version of the source service for distribution.
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The name of the service provider of the source service for distribution.
	//
	// example:
	//
	// SourceSupplier
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// The statistics information.
	Statistic *GetServiceResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	// The service status.
	//
	// Valid values:
	//
	// - Draft: The service is in the draft state.
	//
	// - Submitted: The service is submitted for review. You cannot modify the service.
	//
	// - Approved: The service is approved. You cannot modify the service, but you can publish the service.
	//
	// - Launching: The service is being published.
	//
	// - Online: The service is published.
	//
	// - Offline: The service is unpublished.
	//
	// - Beta: The service is in beta.
	//
	// - Creating: The service is being created.
	//
	// - CreateFailed: The service failed to be created.
	//
	// - Updating: The service is being updated.
	//
	// - UpdateFailed: The service failed to be updated.
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
	// Company A
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The contact information of the service provider.
	SupportContacts []*GetServiceResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
	// The service tags.
	Tags []*GetServiceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant type.
	//
	// Valid values:
	//
	// - SingleTenant: single-tenant.
	//
	// - MultiTenant: multitenancy.
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The test status.
	//
	// Valid values:
	//
	// - CONFIG_IS_NULL: The test configuration does not exist.
	//
	// - SERVICE_TEST_SUCCEED: The service passed the test.
	//
	// - SERVICE_TSET_DOING: The service has not passed the test.
	//
	// example:
	//
	// SERVICE_TEST_SUCCEED
	TestStatus *string `json:"TestStatus,omitempty" xml:"TestStatus,omitempty"`
	// The trial duration. Unit: days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The trial type.
	//
	// Valid values:
	//
	// - Trial: The service supports trial.
	//
	// - NotTrial: The service does not support trial.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The time when the service was last updated.
	//
	// example:
	//
	// 2021-05-22T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The upgrade metadata.
	//
	// example:
	//
	// {\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\"]}
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
	// Indicates whether the service is a virtual Internet service.
	//
	// Valid values:
	//
	// - false: No.
	//
	// - true: Yes.
	//
	// example:
	//
	// false
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
	// The virtual Internet service ID.
	//
	// example:
	//
	// service-70a3b15bb6264345****
	VirtualInternetServiceId *string `json:"VirtualInternetServiceId,omitempty" xml:"VirtualInternetServiceId,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) GetAlarmMetadata() *string {
	return s.AlarmMetadata
}

func (s *GetServiceResponseBody) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *GetServiceResponseBody) GetBuildInfo() *string {
	return s.BuildInfo
}

func (s *GetServiceResponseBody) GetBuildParameters() *string {
	return s.BuildParameters
}

func (s *GetServiceResponseBody) GetCategories() *string {
	return s.Categories
}

func (s *GetServiceResponseBody) GetCommodity() *GetServiceResponseBodyCommodity {
	return s.Commodity
}

func (s *GetServiceResponseBody) GetComplianceMetadata() *GetServiceResponseBodyComplianceMetadata {
	return s.ComplianceMetadata
}

func (s *GetServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetServiceResponseBody) GetCrossRegionConnectionStatus() *string {
	return s.CrossRegionConnectionStatus
}

func (s *GetServiceResponseBody) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *GetServiceResponseBody) GetDeployType() *string {
	return s.DeployType
}

func (s *GetServiceResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetServiceResponseBody) GetEntitySource() map[string]*string {
	return s.EntitySource
}

func (s *GetServiceResponseBody) GetIsSupportOperated() *bool {
	return s.IsSupportOperated
}

func (s *GetServiceResponseBody) GetLicenseMetadata() *string {
	return s.LicenseMetadata
}

func (s *GetServiceResponseBody) GetLogMetadata() *string {
	return s.LogMetadata
}

func (s *GetServiceResponseBody) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *GetServiceResponseBody) GetPayFromType() *string {
	return s.PayFromType
}

func (s *GetServiceResponseBody) GetPermission() *string {
	return s.Permission
}

func (s *GetServiceResponseBody) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *GetServiceResponseBody) GetProgress() *int64 {
	return s.Progress
}

func (s *GetServiceResponseBody) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetServiceResponseBody) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *GetServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceResponseBody) GetResellable() *bool {
	return s.Resellable
}

func (s *GetServiceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetServiceResponseBody) GetSecretKey() *string {
	return s.SecretKey
}

func (s *GetServiceResponseBody) GetServiceAuditDocumentUrl() *string {
	return s.ServiceAuditDocumentUrl
}

func (s *GetServiceResponseBody) GetServiceDiscoverable() *string {
	return s.ServiceDiscoverable
}

func (s *GetServiceResponseBody) GetServiceDocumentInfos() []*GetServiceResponseBodyServiceDocumentInfos {
	return s.ServiceDocumentInfos
}

func (s *GetServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceResponseBody) GetServiceInfos() []*GetServiceResponseBodyServiceInfos {
	return s.ServiceInfos
}

func (s *GetServiceResponseBody) GetServiceLocaleConfigs() []*GetServiceResponseBodyServiceLocaleConfigs {
	return s.ServiceLocaleConfigs
}

func (s *GetServiceResponseBody) GetServiceProductUrl() *string {
	return s.ServiceProductUrl
}

func (s *GetServiceResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceResponseBody) GetShareType() *string {
	return s.ShareType
}

func (s *GetServiceResponseBody) GetShareTypeStatus() *string {
	return s.ShareTypeStatus
}

func (s *GetServiceResponseBody) GetSourceServiceId() *string {
	return s.SourceServiceId
}

func (s *GetServiceResponseBody) GetSourceServiceVersion() *string {
	return s.SourceServiceVersion
}

func (s *GetServiceResponseBody) GetSourceSupplierName() *string {
	return s.SourceSupplierName
}

func (s *GetServiceResponseBody) GetStatistic() *GetServiceResponseBodyStatistic {
	return s.Statistic
}

func (s *GetServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *GetServiceResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *GetServiceResponseBody) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *GetServiceResponseBody) GetSupportContacts() []*GetServiceResponseBodySupportContacts {
	return s.SupportContacts
}

func (s *GetServiceResponseBody) GetTags() []*GetServiceResponseBodyTags {
	return s.Tags
}

func (s *GetServiceResponseBody) GetTenantType() *string {
	return s.TenantType
}

func (s *GetServiceResponseBody) GetTestStatus() *string {
	return s.TestStatus
}

func (s *GetServiceResponseBody) GetTrialDuration() *int64 {
	return s.TrialDuration
}

func (s *GetServiceResponseBody) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetServiceResponseBody) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *GetServiceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetServiceResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetServiceResponseBody) GetVirtualInternetService() *string {
	return s.VirtualInternetService
}

func (s *GetServiceResponseBody) GetVirtualInternetServiceId() *string {
	return s.VirtualInternetServiceId
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

func (s *GetServiceResponseBody) SetSecretKey(v string) *GetServiceResponseBody {
	s.SecretKey = &v
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

func (s *GetServiceResponseBody) SetServiceLocaleConfigs(v []*GetServiceResponseBodyServiceLocaleConfigs) *GetServiceResponseBody {
	s.ServiceLocaleConfigs = v
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

func (s *GetServiceResponseBody) Validate() error {
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
		}
	}
	if s.ComplianceMetadata != nil {
		if err := s.ComplianceMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceDocumentInfos != nil {
		for _, item := range s.ServiceDocumentInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceInfos != nil {
		for _, item := range s.ServiceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceLocaleConfigs != nil {
		for _, item := range s.ServiceLocaleConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistic != nil {
		if err := s.Statistic.Validate(); err != nil {
			return err
		}
	}
	if s.SupportContacts != nil {
		for _, item := range s.SupportContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyCommodity struct {
	// The billing method.
	//
	// Valid values:
	//
	// - **PREPAY*	- (default): subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code of Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity modules.
	Components []*string `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The configuration metadata of Lingxiao.
	CssMetadata *GetServiceResponseBodyCommodityCssMetadata `json:"CssMetadata,omitempty" xml:"CssMetadata,omitempty" type:"Struct"`
	// The metadata of Alibaba Cloud Marketplace.
	MarketplaceMetadata *GetServiceResponseBodyCommodityMarketplaceMetadata `json:"MarketplaceMetadata,omitempty" xml:"MarketplaceMetadata,omitempty" type:"Struct"`
	// The metering information.
	MeteringEntities []*GetServiceResponseBodyCommodityMeteringEntities `json:"MeteringEntities,omitempty" xml:"MeteringEntities,omitempty" type:"Repeated"`
	// The configuration metadata of SaaS Boost.
	//
	// example:
	//
	// {
	//
	//     "Enabled":false
	//
	//     //Public access URL
	//
	//     "PublicAccessUrl":"https://example.com"
	//
	// }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The details of the Alibaba Cloud Marketplace specifications.
	Specifications []*GetServiceResponseBodyCommoditySpecifications `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	// The type.
	//
	// Valid values:
	//
	// - Marketplace: Alibaba Cloud Marketplace.
	//
	// - Css: Lingxiao.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodity) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodity) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetServiceResponseBodyCommodity) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetServiceResponseBodyCommodity) GetComponents() []*string {
	return s.Components
}

func (s *GetServiceResponseBodyCommodity) GetCssMetadata() *GetServiceResponseBodyCommodityCssMetadata {
	return s.CssMetadata
}

func (s *GetServiceResponseBodyCommodity) GetMarketplaceMetadata() *GetServiceResponseBodyCommodityMarketplaceMetadata {
	return s.MarketplaceMetadata
}

func (s *GetServiceResponseBodyCommodity) GetMeteringEntities() []*GetServiceResponseBodyCommodityMeteringEntities {
	return s.MeteringEntities
}

func (s *GetServiceResponseBodyCommodity) GetSaasBoostMetadata() *string {
	return s.SaasBoostMetadata
}

func (s *GetServiceResponseBodyCommodity) GetSpecifications() []*GetServiceResponseBodyCommoditySpecifications {
	return s.Specifications
}

func (s *GetServiceResponseBodyCommodity) GetType() *string {
	return s.Type
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

func (s *GetServiceResponseBodyCommodity) Validate() error {
	if s.CssMetadata != nil {
		if err := s.CssMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.MarketplaceMetadata != nil {
		if err := s.MarketplaceMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.MeteringEntities != nil {
		for _, item := range s.MeteringEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Specifications != nil {
		for _, item := range s.Specifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyCommodityCssMetadata struct {
	// The billing item mapping.
	ComponentsMappings []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
	// The configuration information of the metering item.
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The binding relationship between the package and the metering dimension.
	MeteringEntityMappings []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityCssMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadata) GetComponentsMappings() []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	return s.ComponentsMappings
}

func (s *GetServiceResponseBodyCommodityCssMetadata) GetMeteringEntityExtraInfos() []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos {
	return s.MeteringEntityExtraInfos
}

func (s *GetServiceResponseBodyCommodityCssMetadata) GetMeteringEntityMappings() []*GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings {
	return s.MeteringEntityMappings
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

func (s *GetServiceResponseBodyCommodityCssMetadata) Validate() error {
	if s.ComponentsMappings != nil {
		for _, item := range s.ComponentsMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MeteringEntityExtraInfos != nil {
		for _, item := range s.MeteringEntityExtraInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MeteringEntityMappings != nil {
		for _, item := range s.MeteringEntityMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyCommodityCssMetadataComponentsMappings struct {
	// The mapping.
	Mappings map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GetMappings() map[string]*string {
	return s.Mappings
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetMappings(v map[string]*string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.Mappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.TemplateName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj0048****-Frequency-1
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metric name.
	//
	// example:
	//
	// AvgMemory
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The custom Prometheus statement.
	//
	// example:
	//
	// avg_over_time(count(kube_pod_info{namespace=\\"default\\"})[1h:1m])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// The metering metric.
	//
	// example:
	//
	// ComputeNestPrometheus
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GetEntityId() *string {
	return s.EntityId
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GetMetricName() *string {
	return s.MetricName
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GetPromql() *string {
	return s.Promql
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) GetType() *string {
	return s.Type
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

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityExtraInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj0015****-Frequency-1
	EntityIds *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	// The package name.
	//
	// example:
	//
	// 按量付费套餐
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GetEntityIds() *string {
	return s.EntityIds
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) GetTemplateName() *string {
	return s.TemplateName
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

func (s *GetServiceResponseBodyCommodityCssMetadataMeteringEntityMappings) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityMarketplaceMetadata struct {
	// The configuration information of the metering item.
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The binding relationship between the package and the metering dimension.
	MeteringEntityMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	// The relationship between commodity specifications and templates or packages.
	SpecificationMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) GetMeteringEntityExtraInfos() []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos {
	return s.MeteringEntityExtraInfos
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) GetMeteringEntityMappings() []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings {
	return s.MeteringEntityMappings
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) GetSpecificationMappings() []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	return s.SpecificationMappings
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

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) Validate() error {
	if s.MeteringEntityExtraInfos != nil {
		for _, item := range s.MeteringEntityExtraInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MeteringEntityMappings != nil {
		for _, item := range s.MeteringEntityMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SpecificationMappings != nil {
		for _, item := range s.SpecificationMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj1596****-NetworkOut-2
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the metering metric.
	//
	// example:
	//
	// NetworkLantency
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The custom Prometheus statement.
	//
	// example:
	//
	// avg_over_time(count(kube_pod_info{namespace=\\"default\\"})[1h:1m])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// The metering metric.
	//
	// example:
	//
	// AvgPod
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GetEntityId() *string {
	return s.EntityId
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GetMetricName() *string {
	return s.MetricName
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GetPromql() *string {
	return s.Promql
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) GetType() *string {
	return s.Type
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

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj1596****-NetworkOut-2
	EntityIds *string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty"`
	// The package name.
	//
	// example:
	//
	// Pay-as-you-go plan
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GetEntityIds() *string {
	return s.EntityIds
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) GetTemplateName() *string {
	return s.TemplateName
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

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings struct {
	// The commodity specification code of Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
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
	// The trial type.
	//
	// Valid values:
	//
	// - Trial: Trial is supported.
	//
	// - NotTrial: Trial is not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GetSpecificationCode() *string {
	return s.SpecificationCode
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GetTrialType() *string {
	return s.TrialType
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

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityMeteringEntities struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj5682****-NetworkOut
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the metering item property.
	//
	// example:
	//
	// spring-boot-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetServiceResponseBodyCommodityMeteringEntities) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMeteringEntities) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetEntityId(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.EntityId = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) SetName(v string) *GetServiceResponseBodyCommodityMeteringEntities {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMeteringEntities) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommoditySpecifications struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The specification name.
	//
	// example:
	//
	// specifications1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The purchasable duration. Unit: week or year.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommoditySpecifications) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommoditySpecifications) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommoditySpecifications) GetCode() *string {
	return s.Code
}

func (s *GetServiceResponseBodyCommoditySpecifications) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBodyCommoditySpecifications) GetTimes() []*string {
	return s.Times
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

func (s *GetServiceResponseBodyCommoditySpecifications) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyComplianceMetadata struct {
	// The selected compliance package.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyComplianceMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyComplianceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyComplianceMetadata) GetCompliancePacks() []*string {
	return s.CompliancePacks
}

func (s *GetServiceResponseBodyComplianceMetadata) SetCompliancePacks(v []*string) *GetServiceResponseBodyComplianceMetadata {
	s.CompliancePacks = v
	return s
}

func (s *GetServiceResponseBodyComplianceMetadata) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyServiceDocumentInfos struct {
	// The URL of the document.
	//
	// example:
	//
	// http://doc.com
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The language of the service configuration. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyServiceDocumentInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceDocumentInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceDocumentInfos) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *GetServiceResponseBodyServiceDocumentInfos) GetLocale() *string {
	return s.Locale
}

func (s *GetServiceResponseBodyServiceDocumentInfos) GetTemplateName() *string {
	return s.TemplateName
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

func (s *GetServiceResponseBodyServiceDocumentInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyServiceInfos struct {
	// The information about the service agreement.
	Agreements []*GetServiceResponseBodyServiceInfosAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f35660****.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service configuration.
	//
	// Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
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
	// Database B
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The summary of the service.
	//
	// example:
	//
	// B is an open-source distributed relational database independently designed and developed by Company A.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The information about the software used in the service.
	Softwares []*GetServiceResponseBodyServiceInfosSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfos) GetAgreements() []*GetServiceResponseBodyServiceInfosAgreements {
	return s.Agreements
}

func (s *GetServiceResponseBodyServiceInfos) GetImage() *string {
	return s.Image
}

func (s *GetServiceResponseBodyServiceInfos) GetLocale() *string {
	return s.Locale
}

func (s *GetServiceResponseBodyServiceInfos) GetLongDescriptionUrl() *string {
	return s.LongDescriptionUrl
}

func (s *GetServiceResponseBodyServiceInfos) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBodyServiceInfos) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *GetServiceResponseBodyServiceInfos) GetSoftwares() []*GetServiceResponseBodyServiceInfosSoftwares {
	return s.Softwares
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

func (s *GetServiceResponseBodyServiceInfos) Validate() error {
	if s.Agreements != nil {
		for _, item := range s.Agreements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Softwares != nil {
		for _, item := range s.Softwares {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyServiceInfosAgreements struct {
	// The agreement name.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The URL of the agreement.
	//
	// example:
	//
	// https://aliyun.com/xxxxxxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosAgreements) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosAgreements) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosAgreements) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBodyServiceInfosAgreements) GetUrl() *string {
	return s.Url
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetName(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetUrl(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Url = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosAgreements) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyServiceInfosSoftwares struct {
	// The software name.
	//
	// example:
	//
	// MySQL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 5.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosSoftwares) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosSoftwares) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) GetName() *string {
	return s.Name
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) GetVersion() *string {
	return s.Version
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetName(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetVersion(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Version = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyServiceLocaleConfigs struct {
	// The English value of the business information.
	//
	// example:
	//
	// Service Name
	EnValue *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	// The original value of the business information.
	//
	// example:
	//
	// Service Name
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	// The Chinese value of the business information.
	//
	// example:
	//
	// 服务名称
	ZhValue *string `json:"ZhValue,omitempty" xml:"ZhValue,omitempty"`
}

func (s GetServiceResponseBodyServiceLocaleConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceLocaleConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) GetEnValue() *string {
	return s.EnValue
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) GetZhValue() *string {
	return s.ZhValue
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) SetEnValue(v string) *GetServiceResponseBodyServiceLocaleConfigs {
	s.EnValue = &v
	return s
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) SetOriginalValue(v string) *GetServiceResponseBodyServiceLocaleConfigs {
	s.OriginalValue = &v
	return s
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) SetZhValue(v string) *GetServiceResponseBodyServiceLocaleConfigs {
	s.ZhValue = &v
	return s
}

func (s *GetServiceResponseBodyServiceLocaleConfigs) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyStatistic struct {
	// The total number of service instances. This includes deleted instances.
	//
	// example:
	//
	// 75
	AccumulativeInstanceCount *int32 `json:"AccumulativeInstanceCount,omitempty" xml:"AccumulativeInstanceCount,omitempty"`
	// The total consumption amount of the trial service. Unit: CNY.
	//
	// example:
	//
	// 80.35
	AccumulativePocAmount *float64 `json:"AccumulativePocAmount,omitempty" xml:"AccumulativePocAmount,omitempty"`
	// The total number of users. This includes historical users.
	//
	// example:
	//
	// 60
	AccumulativeUserCount *int32 `json:"AccumulativeUserCount,omitempty" xml:"AccumulativeUserCount,omitempty"`
	// The average consumption amount of a trial service instance. Unit: CNY.
	//
	// example:
	//
	// 40.17
	AveragePocAmount *float64 `json:"AveragePocAmount,omitempty" xml:"AveragePocAmount,omitempty"`
	// The average trial duration of a service instance. Unit: hours.
	//
	// example:
	//
	// 1
	AveragePocDuration *float64 `json:"AveragePocDuration,omitempty" xml:"AveragePocDuration,omitempty"`
	// The average consumption amount of a trial service instance per unit of time. Unit: CNY.
	//
	// example:
	//
	// 167.9
	AveragePocUnitAmount *float64 `json:"AveragePocUnitAmount,omitempty" xml:"AveragePocUnitAmount,omitempty"`
	// The number of online service instances. This indicates the number of service instances that are successfully deployed.
	//
	// example:
	//
	// 20
	DeployedServiceInstanceCount *int32 `json:"DeployedServiceInstanceCount,omitempty" xml:"DeployedServiceInstanceCount,omitempty"`
	// The number of online users. This indicates the number of users who have successfully deployed service instances.
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
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyStatistic) GetAccumulativeInstanceCount() *int32 {
	return s.AccumulativeInstanceCount
}

func (s *GetServiceResponseBodyStatistic) GetAccumulativePocAmount() *float64 {
	return s.AccumulativePocAmount
}

func (s *GetServiceResponseBodyStatistic) GetAccumulativeUserCount() *int32 {
	return s.AccumulativeUserCount
}

func (s *GetServiceResponseBodyStatistic) GetAveragePocAmount() *float64 {
	return s.AveragePocAmount
}

func (s *GetServiceResponseBodyStatistic) GetAveragePocDuration() *float64 {
	return s.AveragePocDuration
}

func (s *GetServiceResponseBodyStatistic) GetAveragePocUnitAmount() *float64 {
	return s.AveragePocUnitAmount
}

func (s *GetServiceResponseBodyStatistic) GetDeployedServiceInstanceCount() *int32 {
	return s.DeployedServiceInstanceCount
}

func (s *GetServiceResponseBodyStatistic) GetDeployedUserCount() *int32 {
	return s.DeployedUserCount
}

func (s *GetServiceResponseBodyStatistic) GetSubmittedUsageCount() *int32 {
	return s.SubmittedUsageCount
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

func (s *GetServiceResponseBodyStatistic) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodySupportContacts struct {
	// The type of the contact information.
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The contact information.
	//
	// example:
	//
	// supplier@test.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodySupportContacts) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodySupportContacts) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodySupportContacts) GetType() *string {
	return s.Type
}

func (s *GetServiceResponseBodySupportContacts) GetValue() *string {
	return s.Value
}

func (s *GetServiceResponseBodySupportContacts) SetType(v string) *GetServiceResponseBodySupportContacts {
	s.Type = &v
	return s
}

func (s *GetServiceResponseBodySupportContacts) SetValue(v string) *GetServiceResponseBodySupportContacts {
	s.Value = &v
	return s
}

func (s *GetServiceResponseBodySupportContacts) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetServiceResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetServiceResponseBodyTags) SetKey(v string) *GetServiceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceResponseBodyTags) SetValue(v string) *GetServiceResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetServiceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
