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
	// The number of entries returned per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI41
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of services.
	Services []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of results.
	//
	// example:
	//
	// 1
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
	// The category of the data disk. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud: basic disk.
	//
	// example:
	//
	// cloud_ssd
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity specifications.
	Commodity *ListServicesResponseBodyServicesCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The deployment source of the service. Valid values:
	//
	// - NoWhere: The service has no deployment source.
	//
	// - Marketplace: The service is deployed from Alibaba Cloud Marketplace.
	//
	// - ComputeNest: The service is deployed from Compute Nest.
	//
	// example:
	//
	// ComputeNest
	DeployFrom *string `json:"DeployFrom,omitempty" xml:"DeployFrom,omitempty"`
	// The deployment type. Valid values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - spi: The service is deployed by calling SPI.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2022-01-21T10:35:44Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The service score.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-6b5d632edd394dxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service information.
	ServiceInfos []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The multi-language configurations of the service.
	ServiceLocaleConfigs []*ListServicesResponseBodyServicesServiceLocaleConfigs `json:"ServiceLocaleConfigs,omitempty" xml:"ServiceLocaleConfigs,omitempty" type:"Repeated"`
	// The URL of the product page.
	//
	// example:
	//
	// http://example1.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The service type. Valid values:
	//
	// - private: The service is deployed in the user\\"s account.
	//
	// - managed: The service is hosted in the service provider\\"s account.
	//
	// - operation: The service is an Alibaba Cloud Managed Service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service status. Valid values:
	//
	// - Draft: The service is in the draft state.
	//
	// - Submitted: The service is submitted for review. You cannot modify the service.
	//
	// - Approved: The service is approved. You cannot modify the service, but you can publish it.
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
	// The service provider name.
	//
	// example:
	//
	// Company A Ltd.
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The English name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierNameEng *string `json:"SupplierNameEng,omitempty" xml:"SupplierNameEng,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 1436322xxxxx
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The service provider\\"s URL.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service tags.
	Tags []*ListServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant type. Valid values:
	//
	// - SingleTenant: single-tenant.
	//
	// - MultiTenant: multi-tenant.
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
	TrialDuration *string `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The trial type. Valid values:
	//
	// - Trial: The service supports trial.
	//
	// - NotTrial: The service does not support trial.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 4
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

func (s *ListServicesResponseBodyServices) GetCategories() *string {
	return s.Categories
}

func (s *ListServicesResponseBodyServices) GetCommodity() *ListServicesResponseBodyServicesCommodity {
	return s.Commodity
}

func (s *ListServicesResponseBodyServices) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListServicesResponseBodyServices) GetDeployFrom() *string {
	return s.DeployFrom
}

func (s *ListServicesResponseBodyServices) GetDeployType() *string {
	return s.DeployType
}

func (s *ListServicesResponseBodyServices) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListServicesResponseBodyServices) GetScore() *int32 {
	return s.Score
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

func (s *ListServicesResponseBodyServices) GetServiceProductUrl() *string {
	return s.ServiceProductUrl
}

func (s *ListServicesResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesResponseBodyServices) GetStatus() *string {
	return s.Status
}

func (s *ListServicesResponseBodyServices) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListServicesResponseBodyServices) GetSupplierNameEng() *string {
	return s.SupplierNameEng
}

func (s *ListServicesResponseBodyServices) GetSupplierUid() *int64 {
	return s.SupplierUid
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

func (s *ListServicesResponseBodyServices) GetTrialDuration() *string {
	return s.TrialDuration
}

func (s *ListServicesResponseBodyServices) GetTrialType() *string {
	return s.TrialType
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

func (s *ListServicesResponseBodyServices) SetDeployFrom(v string) *ListServicesResponseBodyServices {
	s.DeployFrom = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployType(v string) *ListServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPublishTime(v string) *ListServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetScore(v int32) *ListServicesResponseBodyServices {
	s.Score = &v
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

func (s *ListServicesResponseBodyServices) SetServiceProductUrl(v string) *ListServicesResponseBodyServices {
	s.ServiceProductUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
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

func (s *ListServicesResponseBodyServices) SetSupplierNameEng(v string) *ListServicesResponseBodyServices {
	s.SupplierNameEng = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUid(v int64) *ListServicesResponseBodyServices {
	s.SupplierUid = &v
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

func (s *ListServicesResponseBodyServices) SetTrialDuration(v string) *ListServicesResponseBodyServices {
	s.TrialDuration = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTrialType(v string) *ListServicesResponseBodyServices {
	s.TrialType = &v
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
	// The commodity specification code.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The deployment page.
	//
	// example:
	//
	// Order: Order page
	//
	// Detail: Details page
	DeployPage *string `json:"DeployPage,omitempty" xml:"DeployPage,omitempty"`
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

func (s *ListServicesResponseBodyServicesCommodity) GetDeployPage() *string {
	return s.DeployPage
}

func (s *ListServicesResponseBodyServicesCommodity) SetCommodityCode(v string) *ListServicesResponseBodyServicesCommodity {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetDeployPage(v string) *ListServicesResponseBodyServicesCommodity {
	s.DeployPage = &v
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
	// http://img.tidb.oss.url
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
	// The service overview.
	//
	// example:
	//
	// B is an open-source distributed relational database independently designed and developed by Company A.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The software information of the service.
	Softwares []*ListServicesResponseBodyServicesServiceInfosSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
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

func (s *ListServicesResponseBodyServicesServiceInfos) GetSoftwares() []*ListServicesResponseBodyServicesServiceInfosSoftwares {
	return s.Softwares
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

func (s *ListServicesResponseBodyServicesServiceInfos) SetSoftwares(v []*ListServicesResponseBodyServicesServiceInfosSoftwares) *ListServicesResponseBodyServicesServiceInfos {
	s.Softwares = v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) Validate() error {
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

type ListServicesResponseBodyServicesServiceInfosSoftwares struct {
	// The software name.
	//
	// example:
	//
	// wordpress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 6.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfosSoftwares) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfosSoftwares) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) GetName() *string {
	return s.Name
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) GetVersion() *string {
	return s.Version
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) SetName(v string) *ListServicesResponseBodyServicesServiceInfosSoftwares {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) SetVersion(v string) *ListServicesResponseBodyServicesServiceInfosSoftwares {
	s.Version = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) Validate() error {
	return dara.Validate(s)
}

type ListServicesResponseBodyServicesServiceLocaleConfigs struct {
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
