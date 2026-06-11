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
	// A pagination token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no further results exist.
	//
	// - If **NextToken*	- has a value, use it as the starting token for the next query.
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
	// The collection of endpoint services.
	Services []*ListVpcEndpointServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of records.
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
	// The IP address version. Valid values:
	//
	// - **IPv4**: IPv4 type.
	//
	// - **DualStack**: Dual-stack type.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connections. Valid values:
	//
	// - **true**: Automatically accept endpoint connections.
	//
	// - **false**: Do not automatically accept endpoint connections.
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth. Unit: Mbps.
	//
	// example:
	//
	// 1024
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time when the endpoint service was created.
	//
	// example:
	//
	// 2021-09-24T17:15:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The maximum bandwidth of the endpoint connection. Unit: Mbps.
	//
	// example:
	//
	// 1024
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The minimum bandwidth of the endpoint connection. Unit: Mbps.
	//
	// example:
	//
	// 100
	MinBandwidth *int32 `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	// The payer. Valid values:
	//
	// - **Endpoint**: the service consumer.
	//
	// - **EndpointService**: the service provider.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region where the endpoint service is deployed.
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
	// The business status of the endpoint service. Valid values:
	//
	// - **Normal**: The endpoint service is running as expected.
	//
	// - **FinancialLocked**: The endpoint service is locked due to an overdue payment.
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
	// - **slb**: The service resource is a CLB instance.
	//
	// - **alb**: The service resource is an ALB instance.
	//
	// - **nlb**: The service resource is an NLB instance.
	//
	// - **gwlb**: The service resource is a GWLB instance.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The status of the endpoint service. Valid values:
	//
	// - **Creating**: The endpoint service is being created.
	//
	// - **Pending**: The endpoint service is being modified.
	//
	// - **Active**: The endpoint service is available.
	//
	// - **Deleting**: The endpoint service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Deprecated
	//
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint service. Valid values:
	//
	// - **Interface**: an interface endpoint. You can add CLB, ALB, and NLB instances as service resources.
	//
	// - **GatewayLoadBalancer**: a Gateway Load Balancer endpoint. You can add GWLB instances as service resources.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of regions supported by the endpoint service. Service consumers can initiate endpoint connections from these regions.
	SupportedRegionSet []*ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet `json:"SupportedRegionSet,omitempty" xml:"SupportedRegionSet,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*ListVpcEndpointServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Specifies whether zonal affinity is enabled. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
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
	// The business status of the region supported by the endpoint service. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: Locked due to an overdue payment.
	//
	// example:
	//
	// Normal
	RegionBusinessStatus *string `json:"RegionBusinessStatus,omitempty" xml:"RegionBusinessStatus,omitempty"`
	// The status of the region supported by the endpoint service. Valid values:
	//
	// - **Pending**: The status is being updated.
	//
	// - **Available**: The region is available.
	//
	// - **Deleting**: The region is being deleted.
	//
	// - **Failed**: The operation failed.
	//
	// - **Closed**: The region is closed.
	//
	// example:
	//
	// Available
	RegionServiceStatus *string `json:"RegionServiceStatus,omitempty" xml:"RegionServiceStatus,omitempty"`
	// Deprecated
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the region supported by the endpoint service.
	//
	// example:
	//
	// cn-hangzhou
	SupportedRegionId *string `json:"SupportedRegionId,omitempty" xml:"SupportedRegionId,omitempty"`
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

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) GetSupportedRegionId() *string {
	return s.SupportedRegionId
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

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) SetSupportedRegionId(v string) *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet {
	s.SupportedRegionId = &v
	return s
}

func (s *ListVpcEndpointServicesResponseBodyServicesSupportedRegionSet) Validate() error {
	return dara.Validate(s)
}

type ListVpcEndpointServicesResponseBodyServicesTags struct {
	// The tag key of the instance.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance.
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
