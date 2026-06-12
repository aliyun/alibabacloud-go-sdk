// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMetadata(v string) *UpdateServiceRequest
	GetAlarmMetadata() *string
	SetApprovalType(v string) *UpdateServiceRequest
	GetApprovalType() *string
	SetBuildParameters(v string) *UpdateServiceRequest
	GetBuildParameters() *string
	SetClientToken(v string) *UpdateServiceRequest
	GetClientToken() *string
	SetCommodity(v *UpdateServiceRequestCommodity) *UpdateServiceRequest
	GetCommodity() *UpdateServiceRequestCommodity
	SetComplianceMetadata(v *UpdateServiceRequestComplianceMetadata) *UpdateServiceRequest
	GetComplianceMetadata() *UpdateServiceRequestComplianceMetadata
	SetDeployMetadata(v string) *UpdateServiceRequest
	GetDeployMetadata() *string
	SetDeployType(v string) *UpdateServiceRequest
	GetDeployType() *string
	SetDryRun(v bool) *UpdateServiceRequest
	GetDryRun() *bool
	SetDuration(v int64) *UpdateServiceRequest
	GetDuration() *int64
	SetIsDefault(v bool) *UpdateServiceRequest
	GetIsDefault() *bool
	SetIsSupportOperated(v bool) *UpdateServiceRequest
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *UpdateServiceRequest
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *UpdateServiceRequest
	GetLogMetadata() *string
	SetOperationMetadata(v string) *UpdateServiceRequest
	GetOperationMetadata() *string
	SetPolicyNames(v string) *UpdateServiceRequest
	GetPolicyNames() *string
	SetRegionId(v string) *UpdateServiceRequest
	GetRegionId() *string
	SetResellable(v bool) *UpdateServiceRequest
	GetResellable() *bool
	SetServiceId(v string) *UpdateServiceRequest
	GetServiceId() *string
	SetServiceInfo(v []*UpdateServiceRequestServiceInfo) *UpdateServiceRequest
	GetServiceInfo() []*UpdateServiceRequestServiceInfo
	SetServiceLocaleConfigs(v []*UpdateServiceRequestServiceLocaleConfigs) *UpdateServiceRequest
	GetServiceLocaleConfigs() []*UpdateServiceRequestServiceLocaleConfigs
	SetServiceType(v string) *UpdateServiceRequest
	GetServiceType() *string
	SetServiceVersion(v string) *UpdateServiceRequest
	GetServiceVersion() *string
	SetShareType(v string) *UpdateServiceRequest
	GetShareType() *string
	SetTenantType(v string) *UpdateServiceRequest
	GetTenantType() *string
	SetTrialDuration(v int32) *UpdateServiceRequest
	GetTrialDuration() *int32
	SetUpdateOption(v *UpdateServiceRequestUpdateOption) *UpdateServiceRequest
	GetUpdateOption() *UpdateServiceRequestUpdateOption
	SetUpgradeMetadata(v string) *UpdateServiceRequest
	GetUpgradeMetadata() *string
	SetVersionName(v string) *UpdateServiceRequest
	GetVersionName() *string
}

type UpdateServiceRequest struct {
	// The alert configurations for the service.
	//
	// > This configuration takes effect only after you configure an alert-related access policy for **PolicyNames**.
	//
	// example:
	//
	// {\\"CmsTemplateId\\":1162921,\\"TemplateUrl\\":\\"https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1760465342xxxxxx/template/c072ef50-6c03-4d9c-8f0e-d1c440xxxxxx.json\\"}
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The approval type for service usage requests. Valid values:
	//
	// - Manual: The request is manually approved.
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
	// A client token to ensure that the request is idempotent. You can use a client to generate the token. Make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity information.
	Commodity *UpdateServiceRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The compliance check metadata.
	ComplianceMetadata *UpdateServiceRequestComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// The information about the service deployment configuration. The data format varies based on the deployment type. The value is a JSON string.
	//
	// example:
	//
	// {\\"EstimateTime\\":null,\\"SupplierDeployMetadata\\":{\\"DeployTimeout\\":7200},\\"EnableVnc\\":false}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type. Valid values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - spi: The service is deployed by calling an SPI.
	//
	// - operation: The service is an O\\&M service.
	//
	// - container: The service is deployed using containers.
	//
	// - pkg: The service is a package service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Specifies whether to perform a dry run for the request. A dry run checks the permissions and the instance status. Valid values:
	//
	// - true: sends the request but does not update the service.
	//
	// - false: sends the request. If the check is successful, the service is updated.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The O\\&M duration. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration  *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IsDefault *bool  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether to enable O\\&M. Default value: false. Valid values:
	//
	// - true: enables O\\&M.
	//
	// - false: disables O\\&M.
	//
	// > This parameter is required when **ServiceType*	- is set to **private**.
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
	// The O\\&M configuration.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"Custom_Image_Ecs\\":{\\"EnablePrometheus\\":false}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The policy name. The name of a single policy can be up to 128 characters in length. If you specify multiple policies, separate them with commas (,). Only O\\&M-related policies are supported.
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
	// Specifies whether to enable distribution. Valid values:
	//
	// - false: Distribution is not enabled.
	//
	// - true: Distribution is enabled.
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
	// The multilingual configurations of the service.
	ServiceLocaleConfigs []*UpdateServiceRequestServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// - private: The service instance is deployed in the user account.
	//
	// - managed: The service instance is deployed in the service provider account.
	//
	// - operation: The service instance is an O\\&M instance.
	//
	// - poc: The service instance is a trial instance.
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
	// The sharing type. Valid values:
	//
	// - Public: The service is public. Formal and trial deployments are not restricted.
	//
	// - Restricted: The service is restricted. Formal and trial deployments are restricted.
	//
	// - OnlyFormalRestricted: Only formal deployments are restricted.
	//
	// - OnlyTrailRestricted: Only trial deployments are restricted.
	//
	// - Hidden: The service is hidden. You cannot view the service or request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tenant type. Valid values:
	//
	// - SingleTenant: The service is single-tenant.
	//
	// - MultiTenant: The service is multi-tenant.
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
	TrialDuration *int32 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The update options.
	UpdateOption *UpdateServiceRequestUpdateOption `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty" type:"Struct"`
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

func (s UpdateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) GetAlarmMetadata() *string {
	return s.AlarmMetadata
}

func (s *UpdateServiceRequest) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *UpdateServiceRequest) GetBuildParameters() *string {
	return s.BuildParameters
}

func (s *UpdateServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceRequest) GetCommodity() *UpdateServiceRequestCommodity {
	return s.Commodity
}

func (s *UpdateServiceRequest) GetComplianceMetadata() *UpdateServiceRequestComplianceMetadata {
	return s.ComplianceMetadata
}

func (s *UpdateServiceRequest) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *UpdateServiceRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *UpdateServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *UpdateServiceRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateServiceRequest) GetIsSupportOperated() *bool {
	return s.IsSupportOperated
}

func (s *UpdateServiceRequest) GetLicenseMetadata() *string {
	return s.LicenseMetadata
}

func (s *UpdateServiceRequest) GetLogMetadata() *string {
	return s.LogMetadata
}

func (s *UpdateServiceRequest) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *UpdateServiceRequest) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *UpdateServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceRequest) GetResellable() *bool {
	return s.Resellable
}

func (s *UpdateServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateServiceRequest) GetServiceInfo() []*UpdateServiceRequestServiceInfo {
	return s.ServiceInfo
}

func (s *UpdateServiceRequest) GetServiceLocaleConfigs() []*UpdateServiceRequestServiceLocaleConfigs {
	return s.ServiceLocaleConfigs
}

func (s *UpdateServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *UpdateServiceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *UpdateServiceRequest) GetShareType() *string {
	return s.ShareType
}

func (s *UpdateServiceRequest) GetTenantType() *string {
	return s.TenantType
}

func (s *UpdateServiceRequest) GetTrialDuration() *int32 {
	return s.TrialDuration
}

func (s *UpdateServiceRequest) GetUpdateOption() *UpdateServiceRequestUpdateOption {
	return s.UpdateOption
}

func (s *UpdateServiceRequest) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *UpdateServiceRequest) GetVersionName() *string {
	return s.VersionName
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

func (s *UpdateServiceRequest) SetIsDefault(v bool) *UpdateServiceRequest {
	s.IsDefault = &v
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

func (s *UpdateServiceRequest) SetServiceLocaleConfigs(v []*UpdateServiceRequestServiceLocaleConfigs) *UpdateServiceRequest {
	s.ServiceLocaleConfigs = v
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

func (s *UpdateServiceRequest) Validate() error {
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
	if s.ServiceInfo != nil {
		for _, item := range s.ServiceInfo {
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
	if s.UpdateOption != nil {
		if err := s.UpdateOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceRequestCommodity struct {
	// This parameter is not available.
	ComponentsMappings []*UpdateServiceRequestCommodityComponentsMappings `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
	// The configuration information of the metering item. This parameter is used in the pay-as-you-go scenario of Alibaba Cloud Marketplace.
	MeteringEntityExtraInfos []*UpdateServiceRequestCommodityMeteringEntityExtraInfos `json:"MeteringEntityExtraInfos,omitempty" xml:"MeteringEntityExtraInfos,omitempty" type:"Repeated"`
	// The mapping between templates or packages and metering dimensions. This parameter is used in the pay-as-you-go scenario of Alibaba Cloud Marketplace.
	MeteringEntityMappings []*UpdateServiceRequestCommodityMeteringEntityMappings `json:"MeteringEntityMappings,omitempty" xml:"MeteringEntityMappings,omitempty" type:"Repeated"`
	// The configuration of Software as a Service (SaaS) Boost.
	//
	// example:
	//
	// {}
	SaasBoostConfig *string `json:"SaasBoostConfig,omitempty" xml:"SaasBoostConfig,omitempty"`
	// The mapping between commodity specifications and templates or packages. This parameter is used in the subscription scenario of Alibaba Cloud Marketplace.
	SpecificationMappings []*UpdateServiceRequestCommoditySpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodity) GetComponentsMappings() []*UpdateServiceRequestCommodityComponentsMappings {
	return s.ComponentsMappings
}

func (s *UpdateServiceRequestCommodity) GetMeteringEntityExtraInfos() []*UpdateServiceRequestCommodityMeteringEntityExtraInfos {
	return s.MeteringEntityExtraInfos
}

func (s *UpdateServiceRequestCommodity) GetMeteringEntityMappings() []*UpdateServiceRequestCommodityMeteringEntityMappings {
	return s.MeteringEntityMappings
}

func (s *UpdateServiceRequestCommodity) GetSaasBoostConfig() *string {
	return s.SaasBoostConfig
}

func (s *UpdateServiceRequestCommodity) GetSpecificationMappings() []*UpdateServiceRequestCommoditySpecificationMappings {
	return s.SpecificationMappings
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

func (s *UpdateServiceRequestCommodity) Validate() error {
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

type UpdateServiceRequestCommodityComponentsMappings struct {
	// This parameter is not available.
	//
	// example:
	//
	// This parameter is not publicly available.
	Mappings map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	// This parameter is not available.
	//
	// example:
	//
	// This parameter is not publicly available.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommodityComponentsMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestCommodityComponentsMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityComponentsMappings) GetMappings() map[string]*string {
	return s.Mappings
}

func (s *UpdateServiceRequestCommodityComponentsMappings) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateServiceRequestCommodityComponentsMappings) SetMappings(v map[string]*string) *UpdateServiceRequestCommodityComponentsMappings {
	s.Mappings = v
	return s
}

func (s *UpdateServiceRequestCommodityComponentsMappings) SetTemplateName(v string) *UpdateServiceRequestCommodityComponentsMappings {
	s.TemplateName = &v
	return s
}

func (s *UpdateServiceRequestCommodityComponentsMappings) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestCommodityMeteringEntityExtraInfos struct {
	// The metering item ID.
	//
	// example:
	//
	// cmgj0006xxxx-Memory-1
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The metric name. This parameter is required when Type is set to ComputeNestBill or ComputeNestPrometheus.
	//
	// example:
	//
	// VirtualCpu/ecs.InstanceType
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The Prometheus statement.
	//
	// example:
	//
	// avg_over_time(sum(rate(container_cpu_usage_seconds_total{namespace=~"ALIYUN::StackName"}[2m]))[1h:10s])
	Promql *string `json:"Promql,omitempty" xml:"Promql,omitempty"`
	// The type. Valid values:
	//
	// - **Custom**
	//
	// - **ComputeNestBill**
	//
	// - **ComputeNestPrometheus**
	//
	// - **ComputeNestTime**
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateServiceRequestCommodityMeteringEntityExtraInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestCommodityMeteringEntityExtraInfos) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) GetMetricName() *string {
	return s.MetricName
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) GetPromql() *string {
	return s.Promql
}

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) GetType() *string {
	return s.Type
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

func (s *UpdateServiceRequestCommodityMeteringEntityExtraInfos) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestCommodityMeteringEntityMappings struct {
	// The metering item ID.
	EntityIds []*string `json:"EntityIds,omitempty" xml:"EntityIds,omitempty" type:"Repeated"`
	// The package name.
	//
	// example:
	//
	// 低配版
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommodityMeteringEntityMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestCommodityMeteringEntityMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) GetEntityIds() []*string {
	return s.EntityIds
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) GetTemplateName() *string {
	return s.TemplateName
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

func (s *UpdateServiceRequestCommodityMeteringEntityMappings) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestCommoditySpecificationMappings struct {
	// The specification code.
	//
	// example:
	//
	// yuncode5767800001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Basic edition
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateServiceRequestCommoditySpecificationMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestCommoditySpecificationMappings) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) GetSpecificationCode() *string {
	return s.SpecificationCode
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *UpdateServiceRequestCommoditySpecificationMappings) GetTemplateName() *string {
	return s.TemplateName
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

func (s *UpdateServiceRequestCommoditySpecificationMappings) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestComplianceMetadata struct {
	// The selected compliance package.
	CompliancePacks []*string `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestComplianceMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestComplianceMetadata) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestComplianceMetadata) GetCompliancePacks() []*string {
	return s.CompliancePacks
}

func (s *UpdateServiceRequestComplianceMetadata) SetCompliancePacks(v []*string) *UpdateServiceRequestComplianceMetadata {
	s.CompliancePacks = v
	return s
}

func (s *UpdateServiceRequestComplianceMetadata) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestServiceInfo struct {
	// The information about the service agreements.
	Agreements []*UpdateServiceRequestServiceInfoAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
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
	// Database B
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B is an open-source distributed relational database independently designed and developed by Company A.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The information about the software used in the service.
	Softwares []*UpdateServiceRequestServiceInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s UpdateServiceRequestServiceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestServiceInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfo) GetAgreements() []*UpdateServiceRequestServiceInfoAgreements {
	return s.Agreements
}

func (s *UpdateServiceRequestServiceInfo) GetImage() *string {
	return s.Image
}

func (s *UpdateServiceRequestServiceInfo) GetLocale() *string {
	return s.Locale
}

func (s *UpdateServiceRequestServiceInfo) GetLongDescriptionUrl() *string {
	return s.LongDescriptionUrl
}

func (s *UpdateServiceRequestServiceInfo) GetName() *string {
	return s.Name
}

func (s *UpdateServiceRequestServiceInfo) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *UpdateServiceRequestServiceInfo) GetSoftwares() []*UpdateServiceRequestServiceInfoSoftwares {
	return s.Softwares
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

func (s *UpdateServiceRequestServiceInfo) Validate() error {
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

type UpdateServiceRequestServiceInfoAgreements struct {
	// The name of the agreement document.
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

func (s UpdateServiceRequestServiceInfoAgreements) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestServiceInfoAgreements) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfoAgreements) GetName() *string {
	return s.Name
}

func (s *UpdateServiceRequestServiceInfoAgreements) GetUrl() *string {
	return s.Url
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetName(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoAgreements) SetUrl(v string) *UpdateServiceRequestServiceInfoAgreements {
	s.Url = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoAgreements) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestServiceInfoSoftwares struct {
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

func (s UpdateServiceRequestServiceInfoSoftwares) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestServiceInfoSoftwares) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceInfoSoftwares) GetName() *string {
	return s.Name
}

func (s *UpdateServiceRequestServiceInfoSoftwares) GetVersion() *string {
	return s.Version
}

func (s *UpdateServiceRequestServiceInfoSoftwares) SetName(v string) *UpdateServiceRequestServiceInfoSoftwares {
	s.Name = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoSoftwares) SetVersion(v string) *UpdateServiceRequestServiceInfoSoftwares {
	s.Version = &v
	return s
}

func (s *UpdateServiceRequestServiceInfoSoftwares) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestServiceLocaleConfigs struct {
	// The English value of the business information.
	//
	// example:
	//
	// Service Name
	EnValue *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	// The raw data value of the business information.
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

func (s UpdateServiceRequestServiceLocaleConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestServiceLocaleConfigs) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestServiceLocaleConfigs) GetEnValue() *string {
	return s.EnValue
}

func (s *UpdateServiceRequestServiceLocaleConfigs) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *UpdateServiceRequestServiceLocaleConfigs) GetZhValue() *string {
	return s.ZhValue
}

func (s *UpdateServiceRequestServiceLocaleConfigs) SetEnValue(v string) *UpdateServiceRequestServiceLocaleConfigs {
	s.EnValue = &v
	return s
}

func (s *UpdateServiceRequestServiceLocaleConfigs) SetOriginalValue(v string) *UpdateServiceRequestServiceLocaleConfigs {
	s.OriginalValue = &v
	return s
}

func (s *UpdateServiceRequestServiceLocaleConfigs) SetZhValue(v string) *UpdateServiceRequestServiceLocaleConfigs {
	s.ZhValue = &v
	return s
}

func (s *UpdateServiceRequestServiceLocaleConfigs) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceRequestUpdateOption struct {
	// Specifies whether to update the deployment file.
	//
	// example:
	//
	// true
	UpdateArtifact *bool `json:"UpdateArtifact,omitempty" xml:"UpdateArtifact,omitempty"`
	// The update option. Valid values:
	//
	// - CODE: code.
	//
	// - PARAMETERS: parameters.
	//
	// example:
	//
	// PARAMETERS
	UpdateFrom *string `json:"UpdateFrom,omitempty" xml:"UpdateFrom,omitempty"`
}

func (s UpdateServiceRequestUpdateOption) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequestUpdateOption) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequestUpdateOption) GetUpdateArtifact() *bool {
	return s.UpdateArtifact
}

func (s *UpdateServiceRequestUpdateOption) GetUpdateFrom() *string {
	return s.UpdateFrom
}

func (s *UpdateServiceRequestUpdateOption) SetUpdateArtifact(v bool) *UpdateServiceRequestUpdateOption {
	s.UpdateArtifact = &v
	return s
}

func (s *UpdateServiceRequestUpdateOption) SetUpdateFrom(v string) *UpdateServiceRequestUpdateOption {
	s.UpdateFrom = &v
	return s
}

func (s *UpdateServiceRequestUpdateOption) Validate() error {
	return dara.Validate(s)
}
