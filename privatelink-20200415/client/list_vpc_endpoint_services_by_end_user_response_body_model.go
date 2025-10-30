// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesByEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetRequestId() *string
	SetServices(v []*ListVpcEndpointServicesByEndUserResponseBodyServices) *ListVpcEndpointServicesByEndUserResponseBody
	GetServices() []*ListVpcEndpointServicesByEndUserResponseBodyServices
	SetTotalCount(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetTotalCount() *string
}

type ListVpcEndpointServicesByEndUserResponseBody struct {
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
	// The information about endpoint services.
	Services []*ListVpcEndpointServicesByEndUserResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetServices() []*ListVpcEndpointServicesByEndUserResponseBodyServices {
	return s.Services
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetMaxResults(v int32) *ListVpcEndpointServicesByEndUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetNextToken(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetRequestId(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetServices(v []*ListVpcEndpointServicesByEndUserResponseBodyServices) *ListVpcEndpointServicesByEndUserResponseBody {
	s.Services = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetTotalCount(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) Validate() error {
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

type ListVpcEndpointServicesByEndUserResponseBodyServices struct {
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
	// The payer. Valid values:
	//
	// 	- **Endpoint**: the service consumer
	//
	// 	- **EndpointService**: the service provider
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The domain name of the endpoint service that can be associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the endpoint service that can be associated with the endpoint.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service that can be associated with the endpoint.
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
	// Indicates whether IPv6 is enabled. Valid values:
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
	// Only **Interface*	- is returned, which indicates an interface endpoint. You can specify **CLB*	- and **ALB*	- instances as service resources.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags.
	Tags                []*ListVpcEndpointServicesByEndUserResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneAffinityEnabled *bool                                                       `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The zones of the endpoint service that can be associated with the endpoint.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetPayer() *string {
	return s.Payer
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceDomain() *string {
	return s.ServiceDomain
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetTags() []*ListVpcEndpointServicesByEndUserResponseBodyServicesTags {
	return s.Tags
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetZones() []*string {
	return s.Zones
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetAddressIpVersion(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.AddressIpVersion = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetPayer(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Payer = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetResourceGroupId(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceDomain(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceDomain = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceId(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceName(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceResourceType(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceResourceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceSupportIPv6(v bool) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceType(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetTags(v []*ListVpcEndpointServicesByEndUserResponseBodyServicesTags) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetZoneAffinityEnabled(v bool) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetZones(v []*string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.Zones = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) Validate() error {
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

type ListVpcEndpointServicesByEndUserResponseBodyServicesTags struct {
	// The key of the tag.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServicesTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) SetKey(v string) *ListVpcEndpointServicesByEndUserResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) SetValue(v string) *ListVpcEndpointServicesByEndUserResponseBodyServicesTags {
	s.Value = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServicesTags) Validate() error {
	return dara.Validate(s)
}
