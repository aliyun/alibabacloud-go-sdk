// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServicesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointServicesResponseBody
	GetRequestId() *string
	SetServices(v []*ListVpcEndpointServicesResponseBodyServices) *ListVpcEndpointServicesResponseBody
	GetServices() []*ListVpcEndpointServicesResponseBodyServices
	SetTotalCount(v int32) *ListVpcEndpointServicesResponseBody
	GetTotalCount() *int32
}

type ListVpcEndpointServicesResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The endpoint services.
	Services []*ListVpcEndpointServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointServicesResponseBody) GetServices() []*ListVpcEndpointServicesResponseBodyServices {
	return s.Services
}

func (s *ListVpcEndpointServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpcEndpointServicesResponseBody) SetMaxResults(v int32) *ListVpcEndpointServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetNextToken(v string) *ListVpcEndpointServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetRequestId(v string) *ListVpcEndpointServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetServices(v []*ListVpcEndpointServicesResponseBodyServices) *ListVpcEndpointServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) SetTotalCount(v int32) *ListVpcEndpointServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBody) Validate() error {
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

type ListVpcEndpointServicesResponseBodyServices struct {
	// The protocol. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Indicates whether endpoint connection requests are automatically accepted. Valid values:
	//
	// 	- **true**: Endpoint connection requests are automatically accepted.
	//
	// 	- **false**: Endpoint connection requests are not automatically accepted.
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time when the endpoint service was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-09-24T17:15:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The maximum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The minimum bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	MinBandwidth *int32 `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	// The payer. Valid values:
	//
	// 	- **Endpoint**: service consumer
	//
	// 	- **EndpointService**: service provider
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region ID of the endpoint service.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service state of the endpoint service. Valid values:
	//
	// 	- **Normal**: The endpoint service runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint service is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The domain name of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the endpoint service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: Network Load Balancer (NLB) instance
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The state of the endpoint service. Valid values:
	//
	// 	- **Creating**: The endpoint service is being created.
	//
	// 	- **Pending**: The endpoint service is being modified.
	//
	// 	- **Active**: The endpoint service is available.
	//
	// 	- **Deleting**: The endpoint service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Deprecated
	//
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint service.
	//
	// 	- Only **Interface*	- may be returned. You can specify CLB, ALB, and NLB instances as the service resources of the endpoint service.
	//
	// example:
	//
	// Interface
	ServiceType        *string                                                          `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupportedRegionSet []*ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet `json:"SupportedRegionSet,omitempty" xml:"SupportedRegionSet,omitempty" type:"Repeated"`
	// The tags added to the resource.
	Tags []*ListVpcEndpointServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether zone affinity is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s ListVpcEndpointServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetConnectBandwidth() *int32 {
	return s.ConnectBandwidth
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetMaxBandwidth() *int32 {
	return s.MaxBandwidth
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetMinBandwidth() *int32 {
	return s.MinBandwidth
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetPayer() *string {
	return s.Payer
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceBusinessStatus() *string {
	return s.ServiceBusinessStatus
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceDomain() *string {
	return s.ServiceDomain
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetSupportedRegionSet() []*ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet {
	return s.SupportedRegionSet
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetTags() []*ListVpcEndpointServicesResponseBodyServicesTags {
	return s.Tags
}

func (s *ListVpcEndpointServicesResponseBodyServices) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetAddressIpVersion(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.AddressIpVersion = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetAutoAcceptEnabled(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetConnectBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.ConnectBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetCreateTime(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetMaxBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.MaxBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetMinBandwidth(v int32) *ListVpcEndpointServicesResponseBodyServices {
	s.MinBandwidth = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetPayer(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.Payer = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetRegionId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetResourceGroupId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceBusinessStatus(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceDescription(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceDescription = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceDomain(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceDomain = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceId(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceName(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceResourceType(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceStatus(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceSupportIPv6(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetSupportedRegionSet(v []*ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) *ListVpcEndpointServicesResponseBodyServices {
	s.SupportedRegionSet = v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetTags(v []*ListVpcEndpointServicesResponseBodyServicesTags) *ListVpcEndpointServicesResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesResponseBodyServices {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServices) Validate() error {
	if s.SupportedRegionSet != nil {
		for _, item := range s.SupportedRegionSet {
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

type ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet struct {
	RegionBusinessStatus *string `json:"RegionBusinessStatus,omitempty" xml:"RegionBusinessStatus,omitempty"`
	RegionServiceStatus  *string `json:"RegionServiceStatus,omitempty" xml:"RegionServiceStatus,omitempty"`
	ServiceRegionId      *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) GetRegionBusinessStatus() *string {
	return s.RegionBusinessStatus
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) GetRegionServiceStatus() *string {
	return s.RegionServiceStatus
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) SetRegionBusinessStatus(v string) *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet {
	s.RegionBusinessStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) SetRegionServiceStatus(v string) *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet {
	s.RegionServiceStatus = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) SetServiceRegionId(v string) *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet {
	s.ServiceRegionId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) Validate() error {
	return dara.Validate(s)
}

type ListVpcEndpointServicesResponseBodyServicesTags struct {
	// The key of the tag added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesResponseBodyServicesTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) SetKey(v string) *ListVpcEndpointServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) SetValue(v string) *ListVpcEndpointServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesTags) Validate() error {
	return dara.Validate(s)
}
