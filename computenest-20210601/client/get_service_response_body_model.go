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
	SetCategories(v string) *GetServiceResponseBody
	GetCategories() *string
	SetCommodity(v *GetServiceResponseBodyCommodity) *GetServiceResponseBody
	GetCommodity() *GetServiceResponseBodyCommodity
	SetComplianceMetadata(v *GetServiceResponseBodyComplianceMetadata) *GetServiceResponseBody
	GetComplianceMetadata() *GetServiceResponseBodyComplianceMetadata
	SetDeployFrom(v string) *GetServiceResponseBody
	GetDeployFrom() *string
	SetDeployMetadata(v string) *GetServiceResponseBody
	GetDeployMetadata() *string
	SetDeployType(v string) *GetServiceResponseBody
	GetDeployType() *string
	SetDuration(v int64) *GetServiceResponseBody
	GetDuration() *int64
	SetInstanceRoleInfos(v []*GetServiceResponseBodyInstanceRoleInfos) *GetServiceResponseBody
	GetInstanceRoleInfos() []*GetServiceResponseBodyInstanceRoleInfos
	SetIsSupportOperated(v bool) *GetServiceResponseBody
	GetIsSupportOperated() *bool
	SetLicenseMetadata(v string) *GetServiceResponseBody
	GetLicenseMetadata() *string
	SetLogMetadata(v string) *GetServiceResponseBody
	GetLogMetadata() *string
	SetOperationMetadata(v string) *GetServiceResponseBody
	GetOperationMetadata() *string
	SetPermission(v string) *GetServiceResponseBody
	GetPermission() *string
	SetPolicyNames(v string) *GetServiceResponseBody
	GetPolicyNames() *string
	SetPublishTime(v string) *GetServiceResponseBody
	GetPublishTime() *string
	SetRequestId(v string) *GetServiceResponseBody
	GetRequestId() *string
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
	SetStatus(v string) *GetServiceResponseBody
	GetStatus() *string
	SetSupplierDesc(v string) *GetServiceResponseBody
	GetSupplierDesc() *string
	SetSupplierLogo(v string) *GetServiceResponseBody
	GetSupplierLogo() *string
	SetSupplierName(v string) *GetServiceResponseBody
	GetSupplierName() *string
	SetSupplierUid(v int64) *GetServiceResponseBody
	GetSupplierUid() *int64
	SetSupplierUrl(v string) *GetServiceResponseBody
	GetSupplierUrl() *string
	SetSupportContacts(v []*GetServiceResponseBodySupportContacts) *GetServiceResponseBody
	GetSupportContacts() []*GetServiceResponseBodySupportContacts
	SetTags(v []*GetServiceResponseBodyTags) *GetServiceResponseBody
	GetTags() []*GetServiceResponseBodyTags
	SetTenantType(v string) *GetServiceResponseBody
	GetTenantType() *string
	SetTrialDuration(v int64) *GetServiceResponseBody
	GetTrialDuration() *int64
	SetTrialType(v string) *GetServiceResponseBody
	GetTrialType() *string
	SetVersion(v string) *GetServiceResponseBody
	GetVersion() *string
	SetVersionName(v string) *GetServiceResponseBody
	GetVersionName() *string
}

type GetServiceResponseBody struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// { "TemplateUrl": "http://template.file.url", "ApplicationGroups": [ { "Name": "applicationGroup1", "TemplateUrl": "url1" }, { "Name": "applicationGroup2", "TemplateUrl": "url2" } ] }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The categories of the Flow.
	//
	// example:
	//
	// AI
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace.
	Commodity *GetServiceResponseBodyCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Compliance check metadata.
	ComplianceMetadata *GetServiceResponseBodyComplianceMetadata `json:"ComplianceMetadata,omitempty" xml:"ComplianceMetadata,omitempty" type:"Struct"`
	// Service deployment approach, Valid values：
	//
	// - NoWhere
	//
	// - Marketplace
	//
	// - ComputeNest
	//
	// example:
	//
	// Marketplace
	DeployFrom *string `json:"DeployFrom,omitempty" xml:"DeployFrom,omitempty"`
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"TemplateUrl\\": \\"http://tidbRosFile\\"}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
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
	// Information about the ram role created in the service template.
	InstanceRoleInfos []*GetServiceResponseBodyInstanceRoleInfos `json:"InstanceRoleInfos,omitempty" xml:"InstanceRoleInfos,omitempty" type:"Repeated"`
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
	// {\\"PayType\\":\\"CustomFixTime\\",\\"DefaultLicenseDays\\":7,\\"CustomMetadata\\":[{\\"TemplateName\\":\\"ECS\\",\\"SpecificationName\\":\\"bandwith-0\\",\\"CustomData\\":\\"1\\"}]}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// {\\"Logstores\\":[]}
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The operation metadata.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
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
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service document information.
	ServiceDocumentInfos []*GetServiceResponseBodyServiceDocumentInfos `json:"ServiceDocumentInfos,omitempty" xml:"ServiceDocumentInfos,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos         []*GetServiceResponseBodyServiceInfos         `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceLocaleConfigs []*GetServiceResponseBodyServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// http://example1.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// - private: The service is a private service and is deployed within the account of a customer.
	//
	// - managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// - operation: The service is a hosted O&M service.
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
	// The deploy status of the service. Valid values:
	//
	// - Draft
	//
	// - Beta
	//
	// - Submitted
	//
	// - Approved
	//
	// - Launching
	//
	// - Online
	//
	// - Offline
	//
	// - Creating
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Computing Nest Community service
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/xxx/service-image/xxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 158927391332*****
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*GetServiceResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
	// The tags.
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

func (s *GetServiceResponseBody) GetCategories() *string {
	return s.Categories
}

func (s *GetServiceResponseBody) GetCommodity() *GetServiceResponseBodyCommodity {
	return s.Commodity
}

func (s *GetServiceResponseBody) GetComplianceMetadata() *GetServiceResponseBodyComplianceMetadata {
	return s.ComplianceMetadata
}

func (s *GetServiceResponseBody) GetDeployFrom() *string {
	return s.DeployFrom
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

func (s *GetServiceResponseBody) GetInstanceRoleInfos() []*GetServiceResponseBodyInstanceRoleInfos {
	return s.InstanceRoleInfos
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

func (s *GetServiceResponseBody) GetPermission() *string {
	return s.Permission
}

func (s *GetServiceResponseBody) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *GetServiceResponseBody) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetServiceResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *GetServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceResponseBody) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *GetServiceResponseBody) GetSupplierLogo() *string {
	return s.SupplierLogo
}

func (s *GetServiceResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *GetServiceResponseBody) GetSupplierUid() *int64 {
	return s.SupplierUid
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

func (s *GetServiceResponseBody) GetTrialDuration() *int64 {
	return s.TrialDuration
}

func (s *GetServiceResponseBody) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetServiceResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetServiceResponseBody) SetAlarmMetadata(v string) *GetServiceResponseBody {
	s.AlarmMetadata = &v
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

func (s *GetServiceResponseBody) SetDeployFrom(v string) *GetServiceResponseBody {
	s.DeployFrom = &v
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

func (s *GetServiceResponseBody) SetInstanceRoleInfos(v []*GetServiceResponseBodyInstanceRoleInfos) *GetServiceResponseBody {
	s.InstanceRoleInfos = v
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

func (s *GetServiceResponseBody) SetPermission(v string) *GetServiceResponseBody {
	s.Permission = &v
	return s
}

func (s *GetServiceResponseBody) SetPolicyNames(v string) *GetServiceResponseBody {
	s.PolicyNames = &v
	return s
}

func (s *GetServiceResponseBody) SetPublishTime(v string) *GetServiceResponseBody {
	s.PublishTime = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
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

func (s *GetServiceResponseBody) SetStatus(v string) *GetServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierDesc(v string) *GetServiceResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierLogo(v string) *GetServiceResponseBody {
	s.SupplierLogo = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierName(v string) *GetServiceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUid(v int64) *GetServiceResponseBody {
	s.SupplierUid = &v
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

func (s *GetServiceResponseBody) SetTrialDuration(v int64) *GetServiceResponseBody {
	s.TrialDuration = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialType(v string) *GetServiceResponseBody {
	s.TrialType = &v
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
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The configuration metadata related to Lingxiao.
	CssMetadata *GetServiceResponseBodyCommodityCssMetadata `json:"CssMetadata,omitempty" xml:"CssMetadata,omitempty" type:"Struct"`
	// The deploy page.
	//
	// example:
	//
	// Order： Order page
	//
	// Detail： Detail page
	DeployPage *string `json:"DeployPage,omitempty" xml:"DeployPage,omitempty"`
	// The metadata of Alibaba Cloud Marketplace.
	MarketplaceMetadata *GetServiceResponseBodyCommodityMarketplaceMetadata `json:"MarketplaceMetadata,omitempty" xml:"MarketplaceMetadata,omitempty" type:"Struct"`
	// The order time.
	OrderTime map[string][]*string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	// The configuration metadata related to Saas Boost.
	//
	// example:
	//
	// {
	//
	//     "Enabled":false    "PublicAccessUrl":"https://example.com"
	//
	// }
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

func (s *GetServiceResponseBodyCommodity) GetCssMetadata() *GetServiceResponseBodyCommodityCssMetadata {
	return s.CssMetadata
}

func (s *GetServiceResponseBodyCommodity) GetDeployPage() *string {
	return s.DeployPage
}

func (s *GetServiceResponseBodyCommodity) GetMarketplaceMetadata() *GetServiceResponseBodyCommodityMarketplaceMetadata {
	return s.MarketplaceMetadata
}

func (s *GetServiceResponseBodyCommodity) GetOrderTime() map[string][]*string {
	return s.OrderTime
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

func (s *GetServiceResponseBodyCommodity) SetCssMetadata(v *GetServiceResponseBodyCommodityCssMetadata) *GetServiceResponseBodyCommodity {
	s.CssMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetDeployPage(v string) *GetServiceResponseBodyCommodity {
	s.DeployPage = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMarketplaceMetadata(v *GetServiceResponseBodyCommodityMarketplaceMetadata) *GetServiceResponseBodyCommodity {
	s.MarketplaceMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetOrderTime(v map[string][]*string) *GetServiceResponseBodyCommodity {
	s.OrderTime = v
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

func (s *GetServiceResponseBodyCommodityCssMetadata) SetComponentsMappings(v []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.ComponentsMappings = v
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
	// Template one.
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

type GetServiceResponseBodyCommodityMarketplaceMetadata struct {
	// The mappings between the service specifications and the template or package.
	SpecificationMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) GetSpecificationMappings() []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	return s.SpecificationMappings
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetSpecificationMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.SpecificationMappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings struct {
	// The specification code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package one.
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) Validate() error {
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
	// The compliance pack list.
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

type GetServiceResponseBodyInstanceRoleInfos struct {
	// The content of the policy.
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The information of the RAM entity.
	Principals []*string `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The ram role name.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyInstanceRoleInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyInstanceRoleInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyInstanceRoleInfos) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetServiceResponseBodyInstanceRoleInfos) GetPrincipals() []*string {
	return s.Principals
}

func (s *GetServiceResponseBodyInstanceRoleInfos) GetRoleName() *string {
	return s.RoleName
}

func (s *GetServiceResponseBodyInstanceRoleInfos) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetPolicyDocument(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.PolicyDocument = &v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetPrincipals(v []*string) *GetServiceResponseBodyInstanceRoleInfos {
	s.Principals = v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetRoleName(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.RoleName = &v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetTemplateName(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.TemplateName = &v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceResponseBodyServiceDocumentInfos struct {
	// The URL that is used to access the document.
	//
	// example:
	//
	// https://help.aliyun.com/zh/compute-nest/use-cases/deploy-an-sd-painting-service-instance?spm=a2c4g.11186623.0.i2
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The language that you use for the query. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
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
	// The service name.
	//
	// example:
	//
	// Service document information.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// Docker Community Edition (CE) is a free version of the Docker project, aimed at developers, enthusiasts, and individuals and organizations who want to use container technology.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
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
	// User agreement
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The agreement URL.
	//
	// example:
	//
	// https://url
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
	// The name of the Software.
	//
	// example:
	//
	// wordpress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 6.0.1
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

type GetServiceResponseBodySupportContacts struct {
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
