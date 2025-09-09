// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMetadata(v string) *CreateServiceShrinkRequest
	GetAlarmMetadata() *string
	SetApprovalType(v string) *CreateServiceShrinkRequest
	GetApprovalType() *string
	SetBuildParameters(v string) *CreateServiceShrinkRequest
	GetBuildParameters() *string
	SetClientToken(v string) *CreateServiceShrinkRequest
	GetClientToken() *string
	SetComplianceMetadataShrink(v string) *CreateServiceShrinkRequest
	GetComplianceMetadataShrink() *string
	SetDeployMetadata(v string) *CreateServiceShrinkRequest
	GetDeployMetadata() *string
	SetDeployType(v string) *CreateServiceShrinkRequest
	GetDeployType() *string
	SetDryRun(v bool) *CreateServiceShrinkRequest
	GetDryRun() *bool
	SetDuration(v int64) *CreateServiceShrinkRequest
	GetDuration() *int64
	SetIsSupportOperated(v bool) *CreateServiceShrinkRequest
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *CreateServiceShrinkRequest
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *CreateServiceShrinkRequest
	GetLogMetadata() *string
	SetOperationMetadata(v string) *CreateServiceShrinkRequest
	GetOperationMetadata() *string
	SetPolicyNames(v string) *CreateServiceShrinkRequest
	GetPolicyNames() *string
	SetRegionId(v string) *CreateServiceShrinkRequest
	GetRegionId() *string
	SetResellable(v bool) *CreateServiceShrinkRequest
	GetResellable() *bool
	SetResourceGroupId(v string) *CreateServiceShrinkRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceShrinkRequest
	GetServiceId() *string
	SetServiceInfo(v []*CreateServiceShrinkRequestServiceInfo) *CreateServiceShrinkRequest
	GetServiceInfo() []*CreateServiceShrinkRequestServiceInfo
	SetServiceType(v string) *CreateServiceShrinkRequest
	GetServiceType() *string
	SetShareType(v string) *CreateServiceShrinkRequest
	GetShareType() *string
	SetSourceServiceId(v string) *CreateServiceShrinkRequest
	GetSourceServiceId() *string
	SetSourceServiceVersion(v string) *CreateServiceShrinkRequest
	GetSourceServiceVersion() *string
	SetTag(v []*CreateServiceShrinkRequestTag) *CreateServiceShrinkRequest
	GetTag() []*CreateServiceShrinkRequestTag
	SetTenantType(v string) *CreateServiceShrinkRequest
	GetTenantType() *string
	SetTrialDuration(v int64) *CreateServiceShrinkRequest
	GetTrialDuration() *int64
	SetUpgradeMetadata(v string) *CreateServiceShrinkRequest
	GetUpgradeMetadata() *string
	SetVersionName(v string) *CreateServiceShrinkRequest
	GetVersionName() *string
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
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequest) GetAlarmMetadata() *string {
	return s.AlarmMetadata
}

func (s *CreateServiceShrinkRequest) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *CreateServiceShrinkRequest) GetBuildParameters() *string {
	return s.BuildParameters
}

func (s *CreateServiceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceShrinkRequest) GetComplianceMetadataShrink() *string {
	return s.ComplianceMetadataShrink
}

func (s *CreateServiceShrinkRequest) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *CreateServiceShrinkRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateServiceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceShrinkRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateServiceShrinkRequest) GetIsSupportOperated() *bool {
	return s.IsSupportOperated
}

func (s *CreateServiceShrinkRequest) GetLicenseMetadata() *string {
	return s.LicenseMetadata
}

func (s *CreateServiceShrinkRequest) GetLogMetadata() *string {
	return s.LogMetadata
}

func (s *CreateServiceShrinkRequest) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *CreateServiceShrinkRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *CreateServiceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceShrinkRequest) GetResellable() *bool {
	return s.Resellable
}

func (s *CreateServiceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceShrinkRequest) GetServiceInfo() []*CreateServiceShrinkRequestServiceInfo {
	return s.ServiceInfo
}

func (s *CreateServiceShrinkRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateServiceShrinkRequest) GetShareType() *string {
	return s.ShareType
}

func (s *CreateServiceShrinkRequest) GetSourceServiceId() *string {
	return s.SourceServiceId
}

func (s *CreateServiceShrinkRequest) GetSourceServiceVersion() *string {
	return s.SourceServiceVersion
}

func (s *CreateServiceShrinkRequest) GetTag() []*CreateServiceShrinkRequestTag {
	return s.Tag
}

func (s *CreateServiceShrinkRequest) GetTenantType() *string {
	return s.TenantType
}

func (s *CreateServiceShrinkRequest) GetTrialDuration() *int64 {
	return s.TrialDuration
}

func (s *CreateServiceShrinkRequest) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *CreateServiceShrinkRequest) GetVersionName() *string {
	return s.VersionName
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

func (s *CreateServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfo) GetAgreements() []*CreateServiceShrinkRequestServiceInfoAgreements {
	return s.Agreements
}

func (s *CreateServiceShrinkRequestServiceInfo) GetImage() *string {
	return s.Image
}

func (s *CreateServiceShrinkRequestServiceInfo) GetLocale() *string {
	return s.Locale
}

func (s *CreateServiceShrinkRequestServiceInfo) GetLongDescriptionUrl() *string {
	return s.LongDescriptionUrl
}

func (s *CreateServiceShrinkRequestServiceInfo) GetName() *string {
	return s.Name
}

func (s *CreateServiceShrinkRequestServiceInfo) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *CreateServiceShrinkRequestServiceInfo) GetSoftwares() []*CreateServiceShrinkRequestServiceInfoSoftwares {
	return s.Softwares
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

func (s *CreateServiceShrinkRequestServiceInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) GetName() *string {
	return s.Name
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) GetUrl() *string {
	return s.Url
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) SetName(v string) *CreateServiceShrinkRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) SetUrl(v string) *CreateServiceShrinkRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoAgreements) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) GetName() *string {
	return s.Name
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) GetVersion() *string {
	return s.Version
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) SetName(v string) *CreateServiceShrinkRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) SetVersion(v string) *CreateServiceShrinkRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

func (s *CreateServiceShrinkRequestServiceInfoSoftwares) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceShrinkRequestTag) SetKey(v string) *CreateServiceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceShrinkRequestTag) SetValue(v string) *CreateServiceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
