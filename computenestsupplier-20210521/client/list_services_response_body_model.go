// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServicesResponseBody
	GetRequestId() *string
	SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody
	GetServices() []*ListServicesResponseBodyServices
	SetTotalCount(v int32) *ListServicesResponseBody
	GetTotalCount() *int32
}

type ListServicesResponseBody struct {
	// The number of entries returned on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
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
	// The list of services.
	Services []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries that meet the filter criteria.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServicesResponseBody) GetServices() []*ListServicesResponseBodyServices {
	return s.Services
}

func (s *ListServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListServicesResponseBody) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServicesResponseBodyServices struct {
	// The approval type for service usage requests. Valid values:
	//
	// - Manual: Manual approval.
	//
	// - AutoPass: Automatic approval.
	//
	// example:
	//
	// AutoPass
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The artifact ID.
	//
	// example:
	//
	// artifact-21ca53ac16a643****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The artifact version.
	//
	// example:
	//
	// draft
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The information about the service build.
	//
	// example:
	//
	// { "RepoUrl": "https://github.com/user/example.git", "Brancn": "main"}
	BuildInfo *string `json:"BuildInfo,omitempty" xml:"BuildInfo,omitempty"`
	// The service category.
	//
	// example:
	//
	// OpenSource
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity specifications.
	Commodity *ListServicesResponseBodyServicesCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the service was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service is the default version. Valid values:
	//
	// - false: The service is not the default version.
	//
	// - true: The service is the default version.
	//
	// example:
	//
	// true
	DefaultVersion *bool `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The deployment type. Valid values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - spi: The service is deployed by calling an SPI.
	//
	// - operation: The service is deployed using Alibaba Cloud Managed Services.
	//
	// - container: The service is deployed using a container.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Indicates whether the service has a beta version. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	HasBeta *bool `json:"HasBeta,omitempty" xml:"HasBeta,omitempty"`
	// Indicates whether the service has a draft version. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	HasDraft *bool `json:"HasDraft,omitempty" xml:"HasDraft,omitempty"`
	// The latest version of the source service for distribution.
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
	// The artifact association type. Valid values:
	//
	// - ServiceDeployment: Service deployment.
	//
	// - ServiceUpgrade: Service upgrade.
	//
	// example:
	//
	// ServiceDeployment
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The distribution authorization status of the service. Valid values:
	//
	// - CanApply: You can apply for authorization.
	//
	// - Applied: An application has been submitted.
	//
	// - Approved: The application is approved.
	//
	// example:
	//
	// CanApply
	ResellApplyStatus *string `json:"ResellApplyStatus,omitempty" xml:"ResellApplyStatus,omitempty"`
	// The ID of the distributed service.
	//
	// example:
	//
	// service-70a3b15bb62643****
	ResellServiceId *string `json:"ResellServiceId,omitempty" xml:"ResellServiceId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz5b555****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the service is discoverable. Valid values:
	//
	// - INVISIBLE: Not discoverable.
	//
	// - DISCOVERABLE: Discoverable.
	//
	// example:
	//
	// INVISIBLE
	ServiceDiscoverable *string `json:"ServiceDiscoverable,omitempty" xml:"ServiceDiscoverable,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-70a3b15bb62643****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service information.
	ServiceInfos []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The multi-language configurations of the service.
	ServiceLocaleConfigs []*ListServicesResponseBodyServicesServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// - private: The service is deployed in the user\\"s account.
	//
	// - managed: The service is deployed in the service provider\\"s account.
	//
	// - operation: It is an Alibaba Cloud Managed Service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The sharing type. Valid values:
	//
	// - Public: Public. Official and trial deployments are not restricted.
	//
	// - Restricted: Restricted. Official and trial deployments are restricted.
	//
	// - OnlyFormalRestricted: Only official deployments are restricted.
	//
	// - OnlyTrailRestricted: Only trial deployments are restricted.
	//
	// - Hidden: Hidden. The service is not visible and you cannot request deployment permissions.
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
	// The ID of the source service for distribution.
	//
	// example:
	//
	// service-70a3b15bb62643****
	SourceServiceId *string `json:"SourceServiceId,omitempty" xml:"SourceServiceId,omitempty"`
	// The version of the source service for distribution.
	//
	// example:
	//
	// 1
	SourceServiceVersion *string `json:"SourceServiceVersion,omitempty" xml:"SourceServiceVersion,omitempty"`
	// The name of the source service provider for distribution.
	//
	// example:
	//
	// SourceSupplier
	SourceSupplierName *string `json:"SourceSupplierName,omitempty" xml:"SourceSupplierName,omitempty"`
	// The service status. Valid values:
	//
	// - Draft: The service is in the Draft state.
	//
	// - Submitted: The service is submitted for review. Modifications are not allowed.
	//
	// - Approved: The service is approved. Modifications are not allowed. The service can be published.
	//
	// - Launching: The service is being published.
	//
	// - Online: The service is published.
	//
	// - Offline: The service is unpublished.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Company A Ltd.
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service tags.
	Tags []*ListServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant type for the managed service. Valid values:
	//
	// - SingleTenant: Single-tenant.
	//
	// - MultiTenant: Multi-tenant.
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial policy. Valid values:
	//
	// - Trial: The service supports a trial.
	//
	// - NotTrial: The service does not support a trial.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The time when the service was last updated.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// v2.0.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Indicates whether the service is a virtual Internet service. Valid values:
	//
	// - false: No.
	//
	// - true: Yes.
	//
	// example:
	//
	// false
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *ListServicesResponseBodyServices) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListServicesResponseBodyServices) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *ListServicesResponseBodyServices) GetBuildInfo() *string {
	return s.BuildInfo
}

func (s *ListServicesResponseBodyServices) GetCategories() *string {
	return s.Categories
}

func (s *ListServicesResponseBodyServices) GetCommodity() *ListServicesResponseBodyServicesCommodity {
	return s.Commodity
}

func (s *ListServicesResponseBodyServices) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListServicesResponseBodyServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServicesResponseBodyServices) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *ListServicesResponseBodyServices) GetDeployType() *string {
	return s.DeployType
}

func (s *ListServicesResponseBodyServices) GetHasBeta() *bool {
	return s.HasBeta
}

func (s *ListServicesResponseBodyServices) GetHasDraft() *bool {
	return s.HasDraft
}

func (s *ListServicesResponseBodyServices) GetLatestResellSourceServiceVersion() *string {
	return s.LatestResellSourceServiceVersion
}

func (s *ListServicesResponseBodyServices) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListServicesResponseBodyServices) GetRelationType() *string {
	return s.RelationType
}

func (s *ListServicesResponseBodyServices) GetResellApplyStatus() *string {
	return s.ResellApplyStatus
}

func (s *ListServicesResponseBodyServices) GetResellServiceId() *string {
	return s.ResellServiceId
}

func (s *ListServicesResponseBodyServices) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesResponseBodyServices) GetServiceDiscoverable() *string {
	return s.ServiceDiscoverable
}

func (s *ListServicesResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServicesResponseBodyServices) GetServiceInfos() []*ListServicesResponseBodyServicesServiceInfos {
	return s.ServiceInfos
}

func (s *ListServicesResponseBodyServices) GetServiceLocaleConfigs() []*ListServicesResponseBodyServicesServiceLocaleConfigs {
	return s.ServiceLocaleConfigs
}

func (s *ListServicesResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesResponseBodyServices) GetShareType() *string {
	return s.ShareType
}

func (s *ListServicesResponseBodyServices) GetSourceImage() *string {
	return s.SourceImage
}

func (s *ListServicesResponseBodyServices) GetSourceServiceId() *string {
	return s.SourceServiceId
}

func (s *ListServicesResponseBodyServices) GetSourceServiceVersion() *string {
	return s.SourceServiceVersion
}

func (s *ListServicesResponseBodyServices) GetSourceSupplierName() *string {
	return s.SourceSupplierName
}

func (s *ListServicesResponseBodyServices) GetStatus() *string {
	return s.Status
}

func (s *ListServicesResponseBodyServices) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListServicesResponseBodyServices) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *ListServicesResponseBodyServices) GetTags() []*ListServicesResponseBodyServicesTags {
	return s.Tags
}

func (s *ListServicesResponseBodyServices) GetTenantType() *string {
	return s.TenantType
}

func (s *ListServicesResponseBodyServices) GetTrialType() *string {
	return s.TrialType
}

func (s *ListServicesResponseBodyServices) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListServicesResponseBodyServices) GetVersion() *string {
	return s.Version
}

func (s *ListServicesResponseBodyServices) GetVersionName() *string {
	return s.VersionName
}

func (s *ListServicesResponseBodyServices) GetVirtualInternetService() *string {
	return s.VirtualInternetService
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

func (s *ListServicesResponseBodyServices) SetServiceLocaleConfigs(v []*ListServicesResponseBodyServicesServiceLocaleConfigs) *ListServicesResponseBodyServices {
	s.ServiceLocaleConfigs = v
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

func (s *ListServicesResponseBodyServices) Validate() error {
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
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

type ListServicesResponseBodyServicesCommodity struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The metadata of the SaaS Boost configuration.
	//
	// example:
	//
	// {
	//
	//      //Enable/disable SaaS Boost binding
	//
	//     "Enabled":true/false,default is false
	//
	//     //Public access URL
	//
	//     "PublicAccessUrl":"https://example.com"
	//
	// }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The type. Valid values:
	//
	// - Marketplace: Alibaba Cloud Marketplace.
	//
	// - Css: Lingxiao.
	//
	// - SaasBoost: SaaS Boost.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServicesResponseBodyServicesCommodity) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServicesCommodity) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesCommodity) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListServicesResponseBodyServicesCommodity) GetSaasBoostMetadata() *string {
	return s.SaasBoostMetadata
}

func (s *ListServicesResponseBodyServicesCommodity) GetType() *string {
	return s.Type
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

func (s *ListServicesResponseBodyServicesCommodity) Validate() error {
	return dara.Validate(s)
}

type ListServicesResponseBodyServicesServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.example.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
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
	// The service name.
	//
	// example:
	//
	// Database B
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A brief description of the service.
	//
	// example:
	//
	// B is an open-source distributed relational database independently designed and developed by Company A.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfos) GetImage() *string {
	return s.Image
}

func (s *ListServicesResponseBodyServicesServiceInfos) GetLocale() *string {
	return s.Locale
}

func (s *ListServicesResponseBodyServicesServiceInfos) GetName() *string {
	return s.Name
}

func (s *ListServicesResponseBodyServicesServiceInfos) GetShortDescription() *string {
	return s.ShortDescription
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

func (s *ListServicesResponseBodyServicesServiceInfos) Validate() error {
	return dara.Validate(s)
}

type ListServicesResponseBodyServicesServiceLocaleConfigs struct {
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

func (s ListServicesResponseBodyServicesServiceLocaleConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceLocaleConfigs) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) GetEnValue() *string {
	return s.EnValue
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) GetZhValue() *string {
	return s.ZhValue
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) SetEnValue(v string) *ListServicesResponseBodyServicesServiceLocaleConfigs {
	s.EnValue = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) SetOriginalValue(v string) *ListServicesResponseBodyServicesServiceLocaleConfigs {
	s.OriginalValue = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) SetZhValue(v string) *ListServicesResponseBodyServicesServiceLocaleConfigs {
	s.ZhValue = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceLocaleConfigs) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesTags) GetKey() *string {
	return s.Key
}

func (s *ListServicesResponseBodyServicesTags) GetValue() *string {
	return s.Value
}

func (s *ListServicesResponseBodyServicesTags) SetKey(v string) *ListServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListServicesResponseBodyServicesTags) SetValue(v string) *ListServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

func (s *ListServicesResponseBodyServicesTags) Validate() error {
	return dara.Validate(s)
}
