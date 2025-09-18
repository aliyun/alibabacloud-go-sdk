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
	SecretKey       *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
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
	ServiceInfos         []*GetServiceResponseBodyServiceInfos         `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceLocaleConfigs []*GetServiceResponseBodyServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	// The configurations of the billable items.
	MeteringEntityExtraInfos []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The billable items that are associated with the package.
	MeteringEntityMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	// The mappings between the service specifications and the template or package.
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
	return dara.Validate(s)
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
	// The compliance package is selected.
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
	return dara.Validate(s)
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
	EnValue       *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	ZhValue       *string `json:"ZhValue,omitempty" xml:"ZhValue,omitempty"`
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
