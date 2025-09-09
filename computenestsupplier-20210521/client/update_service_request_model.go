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

func (s *UpdateServiceRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	// The compliance pack.
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
	return dara.Validate(s)
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
