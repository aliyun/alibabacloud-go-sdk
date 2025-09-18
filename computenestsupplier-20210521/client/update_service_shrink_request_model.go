// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMetadata(v string) *UpdateServiceShrinkRequest
	GetAlarmMetadata() *string
	SetApprovalType(v string) *UpdateServiceShrinkRequest
	GetApprovalType() *string
	SetBuildParameters(v string) *UpdateServiceShrinkRequest
	GetBuildParameters() *string
	SetClientToken(v string) *UpdateServiceShrinkRequest
	GetClientToken() *string
	SetCommodityShrink(v string) *UpdateServiceShrinkRequest
	GetCommodityShrink() *string
	SetComplianceMetadataShrink(v string) *UpdateServiceShrinkRequest
	GetComplianceMetadataShrink() *string
	SetDeployMetadata(v string) *UpdateServiceShrinkRequest
	GetDeployMetadata() *string
	SetDeployType(v string) *UpdateServiceShrinkRequest
	GetDeployType() *string
	SetDryRun(v bool) *UpdateServiceShrinkRequest
	GetDryRun() *bool
	SetDuration(v int64) *UpdateServiceShrinkRequest
	GetDuration() *int64
	SetIsSupportOperated(v bool) *UpdateServiceShrinkRequest
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *UpdateServiceShrinkRequest
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *UpdateServiceShrinkRequest
	GetLogMetadata() *string
	SetOperationMetadata(v string) *UpdateServiceShrinkRequest
	GetOperationMetadata() *string
	SetPolicyNames(v string) *UpdateServiceShrinkRequest
	GetPolicyNames() *string
	SetRegionId(v string) *UpdateServiceShrinkRequest
	GetRegionId() *string
	SetResellable(v bool) *UpdateServiceShrinkRequest
	GetResellable() *bool
	SetServiceId(v string) *UpdateServiceShrinkRequest
	GetServiceId() *string
	SetServiceInfo(v []*UpdateServiceShrinkRequestServiceInfo) *UpdateServiceShrinkRequest
	GetServiceInfo() []*UpdateServiceShrinkRequestServiceInfo
	SetServiceLocaleConfigs(v []*UpdateServiceShrinkRequestServiceLocaleConfigs) *UpdateServiceShrinkRequest
	GetServiceLocaleConfigs() []*UpdateServiceShrinkRequestServiceLocaleConfigs
	SetServiceType(v string) *UpdateServiceShrinkRequest
	GetServiceType() *string
	SetServiceVersion(v string) *UpdateServiceShrinkRequest
	GetServiceVersion() *string
	SetShareType(v string) *UpdateServiceShrinkRequest
	GetShareType() *string
	SetTenantType(v string) *UpdateServiceShrinkRequest
	GetTenantType() *string
	SetTrialDuration(v int32) *UpdateServiceShrinkRequest
	GetTrialDuration() *int32
	SetUpdateOptionShrink(v string) *UpdateServiceShrinkRequest
	GetUpdateOptionShrink() *string
	SetUpgradeMetadata(v string) *UpdateServiceShrinkRequest
	GetUpgradeMetadata() *string
	SetVersionName(v string) *UpdateServiceShrinkRequest
	GetVersionName() *string
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
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The Parameters to build service parameters.
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
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
	ServiceInfo          []*UpdateServiceShrinkRequestServiceInfo          `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	ServiceLocaleConfigs []*UpdateServiceShrinkRequestServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
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
	return dara.Prettify(s)
}

func (s UpdateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequest) GetAlarmMetadata() *string {
	return s.AlarmMetadata
}

func (s *UpdateServiceShrinkRequest) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *UpdateServiceShrinkRequest) GetBuildParameters() *string {
	return s.BuildParameters
}

func (s *UpdateServiceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceShrinkRequest) GetCommodityShrink() *string {
	return s.CommodityShrink
}

func (s *UpdateServiceShrinkRequest) GetComplianceMetadataShrink() *string {
	return s.ComplianceMetadataShrink
}

func (s *UpdateServiceShrinkRequest) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *UpdateServiceShrinkRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateServiceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceShrinkRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *UpdateServiceShrinkRequest) GetIsSupportOperated() *bool {
	return s.IsSupportOperated
}

func (s *UpdateServiceShrinkRequest) GetLicenseMetadata() *string {
	return s.LicenseMetadata
}

func (s *UpdateServiceShrinkRequest) GetLogMetadata() *string {
	return s.LogMetadata
}

func (s *UpdateServiceShrinkRequest) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *UpdateServiceShrinkRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *UpdateServiceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceShrinkRequest) GetResellable() *bool {
	return s.Resellable
}

func (s *UpdateServiceShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateServiceShrinkRequest) GetServiceInfo() []*UpdateServiceShrinkRequestServiceInfo {
	return s.ServiceInfo
}

func (s *UpdateServiceShrinkRequest) GetServiceLocaleConfigs() []*UpdateServiceShrinkRequestServiceLocaleConfigs {
	return s.ServiceLocaleConfigs
}

func (s *UpdateServiceShrinkRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *UpdateServiceShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *UpdateServiceShrinkRequest) GetShareType() *string {
	return s.ShareType
}

func (s *UpdateServiceShrinkRequest) GetTenantType() *string {
	return s.TenantType
}

func (s *UpdateServiceShrinkRequest) GetTrialDuration() *int32 {
	return s.TrialDuration
}

func (s *UpdateServiceShrinkRequest) GetUpdateOptionShrink() *string {
	return s.UpdateOptionShrink
}

func (s *UpdateServiceShrinkRequest) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *UpdateServiceShrinkRequest) GetVersionName() *string {
	return s.VersionName
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

func (s *UpdateServiceShrinkRequest) SetServiceLocaleConfigs(v []*UpdateServiceShrinkRequestServiceLocaleConfigs) *UpdateServiceShrinkRequest {
	s.ServiceLocaleConfigs = v
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

func (s *UpdateServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetAgreements() []*UpdateServiceShrinkRequestServiceInfoAgreements {
	return s.Agreements
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetImage() *string {
	return s.Image
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetLocale() *string {
	return s.Locale
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetLongDescriptionUrl() *string {
	return s.LongDescriptionUrl
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetName() *string {
	return s.Name
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *UpdateServiceShrinkRequestServiceInfo) GetSoftwares() []*UpdateServiceShrinkRequestServiceInfoSoftwares {
	return s.Softwares
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

func (s *UpdateServiceShrinkRequestServiceInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) GetName() *string {
	return s.Name
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) GetUrl() *string {
	return s.Url
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetName(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceShrinkRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoAgreements) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) GetName() *string {
	return s.Name
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) GetVersion() *string {
	return s.Version
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) SetName(v string) *UpdateServiceShrinkRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) SetVersion(v string) *UpdateServiceShrinkRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceInfoSoftwares) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceShrinkRequestServiceLocaleConfigs struct {
	EnValue       *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	ZhValue       *string `json:"ZhValue,omitempty" xml:"ZhValue,omitempty"`
}

func (s UpdateServiceShrinkRequestServiceLocaleConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceShrinkRequestServiceLocaleConfigs) GoString() string {
	return s.String()
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) GetEnValue() *string {
	return s.EnValue
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) GetZhValue() *string {
	return s.ZhValue
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) SetEnValue(v string) *UpdateServiceShrinkRequestServiceLocaleConfigs {
	s.EnValue = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) SetOriginalValue(v string) *UpdateServiceShrinkRequestServiceLocaleConfigs {
	s.OriginalValue = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) SetZhValue(v string) *UpdateServiceShrinkRequestServiceLocaleConfigs {
	s.ZhValue = &v
	return s
}

func (s *UpdateServiceShrinkRequestServiceLocaleConfigs) Validate() error {
	return dara.Validate(s)
}
