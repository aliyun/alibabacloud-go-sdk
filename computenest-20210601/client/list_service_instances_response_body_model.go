// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceInstancesResponseBody
	GetRequestId() *string
	SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody
	GetServiceInstances() []*ListServiceInstancesResponseBodyServiceInstances
	SetTotalCount(v int64) *ListServiceInstancesResponseBody
	GetTotalCount() *int64
}

type ListServiceInstancesResponseBody struct {
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
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the service instances.
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	// The total number of results.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstancesResponseBody) GetServiceInstances() []*ListServiceInstancesResponseBodyServiceInstances {
	return s.ServiceInstances
}

func (s *ListServiceInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v int32) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int64) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstancesResponseBody) Validate() error {
	if s.ServiceInstances != nil {
		for _, item := range s.ServiceInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	// The business status of the service instance. Possible values:
	//
	// - Normal: The service instance is running as expected.
	//
	// - Renewing: The service instance is being renewed.
	//
	// - RenewFoiled: The renewal of the service instance failed.
	//
	// - Expired: The service instance has expired.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance has the Alibaba Cloud Managed Services feature. Possible values:
	//
	// - true: The service instance has the Alibaba Cloud Managed Services feature.
	//
	// - false: The service instance does not have the Alibaba Cloud Managed Services feature.
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// The expiration time of the service instance.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime           *string                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GrantedPermission *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission `json:"GrantedPermission,omitempty" xml:"GrantedPermission,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud Marketplace instance.
	//
	// example:
	//
	// 5827****
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the service instance that is managed by Alibaba Cloud.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The end time of the Alibaba Cloud Managed Services.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The start time of the Alibaba Cloud Managed Services.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2306175xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The output information of the service instance.
	//
	// example:
	//
	// {"managementUrl": "http://xx.xx"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameter information of the service instance.
	//
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method. Valid values:
	//
	// - Permanent: The service is purchased permanently.
	//
	// - Subscription: The service is a subscription instance from Alibaba Cloud Marketplace.
	//
	// - PayAsYouGo: The service is a pay-as-you-go instance from Alibaba Cloud Marketplace.
	//
	// - CustomFixTime: The service has a custom fixed duration.
	//
	// example:
	//
	// Subscription
	PayType     *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The deployment progress of the service instance. Unit: %.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz6scpcxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resources.
	//
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The service.
	Service *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service type. Possible values:
	//
	// - private: The service is deployed in the user\\"s account.
	//
	// - managed: The service is hosted in the service provider\\"s account.
	//
	// - operation: The service is an Alibaba Cloud Managed Service.
	//
	// - poc: The service instance is a trial instance.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source from which the service instance was created.
	//
	// example:
	//
	// Supplier
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the service instance. Possible values:
	//
	// - Created: The service instance is created.
	//
	// - Deploying: The service instance is being deployed.
	//
	// - DeployedFailed: The service instance failed to be deployed.
	//
	// - Deployed: The service instance is deployed.
	//
	// - Upgrading: The service instance is being upgraded.
	//
	// - Deleting: The service instance is being deleted.
	//
	// - Deleted: The service instance is deleted.
	//
	// - DeletedFailed: The service instance failed to be deleted.
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment information of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// Indicates whether the trial instance can be converted to a regular instance.
	SupportTrialToPrivate *bool `json:"SupportTrialToPrivate,omitempty" xml:"SupportTrialToPrivate,omitempty"`
	// The custom tags.
	Tags []*ListServiceInstancesResponseBodyServiceInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetBizStatus() *string {
	return s.BizStatus
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetEnableInstanceOps() *bool {
	return s.EnableInstanceOps
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetEndTime() *string {
	return s.EndTime
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetGrantedPermission() *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission {
	return s.GrantedPermission
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetMarketInstanceId() *string {
	return s.MarketInstanceId
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetName() *string {
	return s.Name
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetOperatedServiceInstanceId() *string {
	return s.OperatedServiceInstanceId
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetOperationEndTime() *string {
	return s.OperationEndTime
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetOperationStartTime() *string {
	return s.OperationStartTime
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetOrderId() *string {
	return s.OrderId
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetOutputs() *string {
	return s.Outputs
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetParameters() *string {
	return s.Parameters
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetPayType() *string {
	return s.PayType
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetProgress() *int64 {
	return s.Progress
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetResources() *string {
	return s.Resources
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetService() *ListServiceInstancesResponseBodyServiceInstancesService {
	return s.Service
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetSource() *string {
	return s.Source
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetStatus() *string {
	return s.Status
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetSupportTrialToPrivate() *bool {
	return s.SupportTrialToPrivate
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetTags() []*ListServiceInstancesResponseBodyServiceInstancesTags {
	return s.Tags
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListServiceInstancesResponseBodyServiceInstances) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetBizStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.BizStatus = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetGrantedPermission(v *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) *ListServiceInstancesResponseBodyServiceInstances {
	s.GrantedPermission = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetMarketInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.MarketInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOrderId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OrderId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOutputs(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Outputs = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPolicyNames(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PolicyNames = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResourceGroupId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResources(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Resources = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSupportTrialToPrivate(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.SupportTrialToPrivate = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) Validate() error {
	if s.GrantedPermission != nil {
		if err := s.GrantedPermission.Validate(); err != nil {
			return err
		}
	}
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			return err
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

type ListServiceInstancesResponseBodyServiceInstancesGrantedPermission struct {
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	PolicyNames      *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) GetOperationEndTime() *string {
	return s.OperationEndTime
}

func (s *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) SetPolicyNames(v string) *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission {
	s.PolicyNames = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesGrantedPermission) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	// The commodity specifications.
	Commodity *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The deployment type. Possible values:
	//
	// - ros: The service is deployed using ROS.
	//
	// - terraform: The service is deployed using Terraform.
	//
	// - ack: The service is deployed using ACK.
	//
	// - spi: The service is deployed by calling SPI.
	//
	// - operation: The service is deployed using Alibaba Cloud Managed Services.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The publishing time.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service information.
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The service type. Possible values:
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
	// The service status.
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
	// The service provider\\"s URL.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
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
	// Version A
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetCommodity() *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	return s.Commodity
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetDeployType() *string {
	return s.DeployType
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetServiceInfos() []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	return s.ServiceInfos
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetStatus() *string {
	return s.Status
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetVersion() *string {
	return s.Version
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) GetVersionName() *string {
	return s.VersionName
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetCommodity(v *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Commodity = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) Validate() error {
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
	return nil
}

type ListServiceInstancesResponseBodyServiceInstancesServiceCommodity struct {
	// The metadata of the Software as a Service (SaaS) Boost configuration.
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
	// The type. Possible values:
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

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) GetSaasBoostMetadata() *string {
	return s.SaasBoostMetadata
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) GetType() *string {
	return s.Type
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetSaasBoostMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetType(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.Type = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
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
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GetImage() *string {
	return s.Image
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GetLocale() *string {
	return s.Locale
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GetName() *string {
	return s.Name
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
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

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) Validate() error {
	return dara.Validate(s)
}
