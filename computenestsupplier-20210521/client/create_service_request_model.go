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
	// The alert configurations for the service.
	//
	// > This configuration takes effect only after an alert-related access policy is configured in \\`PolicyNames\\`.
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
	// The approval type for service usage requests. Valid values:
	//
	// - Manual: The request requires manual approval.
	//
	// - AutoPass: The request is automatically approved.
	//
	// example:
	//
	// Manual
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The parameters for building the service.
	//
	// example:
	//
	// { "ServiceTemplateId": "st-xxxxx"}
	BuildParameters *string `json:"BuildParameters,omitempty" xml:"BuildParameters,omitempty"`
	// A client token used to ensure the idempotence of the request. Generate a unique value for this parameter from your client. \\`ClientToken\\` supports only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The compliance check metadata.
	ComplianceMetadata *CreateServiceRequestComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The deployment configuration of the service. The information stored varies by deployment type. Different deployment types have different data formats. The data is stored in a JSON string.
	//
	// example:
	//
	// {"TemplateConfigs":[{"Name":"Template 1","Url":"oss://computenest-test/template.json?RegionId=cn-beijing","PredefinedParameters":[{"Name":"Basic","Parameters":{"InstanceType":"ecs.g5.large","DataDiskSize":40}},{"Name":"Advanced","Parameters":{"InstanceType":"ecs.g5.large","DataDiskSize":200}}]}]}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type. Valid values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - ack: The service is deployed using ACK.
	//
	// - spi: The service is deployed by invoking an SPI.
	//
	// - operation: The service is an O\\&M service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform a dry run to check the request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The O\\&M duration. Unit: seconds.
	//
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable O\\&M. Default value: false. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// > This parameter is required when \\`ServiceType\\` is set to \\`private\\`.
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
	// The application log configuration.
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
	// The O\\&M configuration.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name of a single policy can be up to 128 characters in length. Separate multiple names with commas (,). Only policies related to O\\&M parameters are supported.
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
	// Specifies whether the service can be distributed. Valid values:
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
	// - private: The service instance is deployed in the user\\"s account.
	//
	// - managed: The service instance is managed in the service provider\\"s account.
	//
	// - operation: An O\\&M service instance.
	//
	// - poc: A trial service instance.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The sharing type. Valid values:
	//
	// - Public: The service is public. Full and trial deployments are not restricted.
	//
	// - Restricted: The service is restricted. Full and trial deployments are restricted.
	//
	// - OnlyFormalRestricted: Only full deployments are restricted.
	//
	// - OnlyTrailRestricted: Only trial deployments are restricted.
	//
	// - Hidden: The service is hidden. It is not visible and you cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The ID of the source service for distribution.
	//
	// example:
	//
	// service-70a3b15bb62643xxxxxx
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The version of the source service for distribution.
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The custom tags.
	Tag []*CreateServiceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The tenant type. Valid values:
	//
	// - SingleTenant: Single-tenant.
	//
	// - MultiTenant: Multitenant.
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: days. The maximum trial duration is 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The upgrade metadata.
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
	if s.ComplianceMetadata != nil {
		if err := s.ComplianceMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceInfo != nil {
		for _, item := range s.ServiceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceRequestComplianceMetadata struct {
	// The selected compliance packages.
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
	// The information about the service agreements.
	Agreements []*CreateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The detailed description of the service.
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
	// A brief description of the service.
	//
	// example:
	//
	// TiDB is an open-source distributed relational database designed and developed by Company A.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The information about the software used in the service.
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

type CreateServiceRequestServiceInfoAgreements struct {
	// The name of the agreement.
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
