// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMetadata(v string) *CreateServiceRequest
	GetAlarmMetadata() *string
	SetApprovalType(v string) *CreateServiceRequest
	GetApprovalType() *string
	SetBuildParameters(v string) *CreateServiceRequest
	GetBuildParameters() *string
	SetClientToken(v string) *CreateServiceRequest
	GetClientToken() *string
	SetComplianceMetadata(v *CreateServiceRequestComplianceMetadata) *CreateServiceRequest
	GetComplianceMetadata() *CreateServiceRequestComplianceMetadata
	SetDeployMetadata(v string) *CreateServiceRequest
	GetDeployMetadata() *string
	SetDeployType(v string) *CreateServiceRequest
	GetDeployType() *string
	SetDryRun(v bool) *CreateServiceRequest
	GetDryRun() *bool
	SetDuration(v int64) *CreateServiceRequest
	GetDuration() *int64
	SetIsSupportOperated(v bool) *CreateServiceRequest
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *CreateServiceRequest
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *CreateServiceRequest
	GetLogMetadata() *string
	SetOperationMetadata(v string) *CreateServiceRequest
	GetOperationMetadata() *string
	SetPolicyNames(v string) *CreateServiceRequest
	GetPolicyNames() *string
	SetRegionId(v string) *CreateServiceRequest
	GetRegionId() *string
	SetResellable(v bool) *CreateServiceRequest
	GetResellable() *bool
	SetResourceGroupId(v string) *CreateServiceRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceRequest
	GetServiceId() *string
	SetServiceInfo(v []*CreateServiceRequestServiceInfo) *CreateServiceRequest
	GetServiceInfo() []*CreateServiceRequestServiceInfo
	SetServiceType(v string) *CreateServiceRequest
	GetServiceType() *string
	SetShareType(v string) *CreateServiceRequest
	GetShareType() *string
	SetSourceServiceId(v string) *CreateServiceRequest
	GetSourceServiceId() *string
	SetSourceServiceVersion(v string) *CreateServiceRequest
	GetSourceServiceVersion() *string
	SetTag(v []*CreateServiceRequestTag) *CreateServiceRequest
	GetTag() []*CreateServiceRequestTag
	SetTenantType(v string) *CreateServiceRequest
	GetTenantType() *string
	SetTrialDuration(v int64) *CreateServiceRequest
	GetTrialDuration() *int64
	SetUpgradeMetadata(v string) *CreateServiceRequest
	GetUpgradeMetadata() *string
	SetVersionName(v string) *CreateServiceRequest
	GetVersionName() *string
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
	return dara.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) GetAlarmMetadata() *string {
	return s.AlarmMetadata
}

func (s *CreateServiceRequest) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *CreateServiceRequest) GetBuildParameters() *string {
	return s.BuildParameters
}

func (s *CreateServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceRequest) GetComplianceMetadata() *CreateServiceRequestComplianceMetadata {
	return s.ComplianceMetadata
}

func (s *CreateServiceRequest) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *CreateServiceRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *CreateServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateServiceRequest) GetIsSupportOperated() *bool {
	return s.IsSupportOperated
}

func (s *CreateServiceRequest) GetLicenseMetadata() *string {
	return s.LicenseMetadata
}

func (s *CreateServiceRequest) GetLogMetadata() *string {
	return s.LogMetadata
}

func (s *CreateServiceRequest) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *CreateServiceRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *CreateServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceRequest) GetResellable() *bool {
	return s.Resellable
}

func (s *CreateServiceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceRequest) GetServiceInfo() []*CreateServiceRequestServiceInfo {
	return s.ServiceInfo
}

func (s *CreateServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateServiceRequest) GetShareType() *string {
	return s.ShareType
}

func (s *CreateServiceRequest) GetSourceServiceId() *string {
	return s.SourceServiceId
}

func (s *CreateServiceRequest) GetSourceServiceVersion() *string {
	return s.SourceServiceVersion
}

func (s *CreateServiceRequest) GetTag() []*CreateServiceRequestTag {
	return s.Tag
}

func (s *CreateServiceRequest) GetTenantType() *string {
	return s.TenantType
}

func (s *CreateServiceRequest) GetTrialDuration() *int64 {
	return s.TrialDuration
}

func (s *CreateServiceRequest) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *CreateServiceRequest) GetVersionName() *string {
	return s.VersionName
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

func (s *CreateServiceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateServiceRequestComplianceMetadata struct {
	// The compliance package selected.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s CreateServiceRequestComplianceMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequestComplianceMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestComplianceMetadata) GetCompliancePacks() []*string {
	return s.CompliancePacks
}

func (s *CreateServiceRequestComplianceMetadata) SetCompliancePacks(v []*string) *CreateServiceRequestComplianceMetadata {
	s.CompliancePacks = v
	return s
}

func (s *CreateServiceRequestComplianceMetadata) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfo) GetAgreements() []*CreateServiceRequestServiceInfoAgreements {
	return s.Agreements
}

func (s *CreateServiceRequestServiceInfo) GetImage() *string {
	return s.Image
}

func (s *CreateServiceRequestServiceInfo) GetLocale() *string {
	return s.Locale
}

func (s *CreateServiceRequestServiceInfo) GetLongDescriptionUrl() *string {
	return s.LongDescriptionUrl
}

func (s *CreateServiceRequestServiceInfo) GetName() *string {
	return s.Name
}

func (s *CreateServiceRequestServiceInfo) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *CreateServiceRequestServiceInfo) GetSoftwares() []*CreateServiceRequestServiceInfoSoftwares {
	return s.Softwares
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

func (s *CreateServiceRequestServiceInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfoAgreements) GetName() *string {
	return s.Name
}

func (s *CreateServiceRequestServiceInfoAgreements) GetUrl() *string {
	return s.Url
}

func (s *CreateServiceRequestServiceInfoAgreements) SetName(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfoAgreements) SetUrl(v string) *CreateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

func (s *CreateServiceRequestServiceInfoAgreements) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestServiceInfoSoftwares) GetName() *string {
	return s.Name
}

func (s *CreateServiceRequestServiceInfoSoftwares) GetVersion() *string {
	return s.Version
}

func (s *CreateServiceRequestServiceInfoSoftwares) SetName(v string) *CreateServiceRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *CreateServiceRequestServiceInfoSoftwares) SetVersion(v string) *CreateServiceRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

func (s *CreateServiceRequestServiceInfoSoftwares) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateServiceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceRequestTag) SetKey(v string) *CreateServiceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceRequestTag) SetValue(v string) *CreateServiceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceRequestTag) Validate() error {
	return dara.Validate(s)
}
