// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *CreateVpcEndpointServiceResponseBody
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceResponseBody
	GetAutoAcceptEnabled() *bool
	SetCreateTime(v string) *CreateVpcEndpointServiceResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *CreateVpcEndpointServiceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateVpcEndpointServiceResponseBody
	GetResourceGroupId() *string
	SetServiceBusinessStatus(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceBusinessStatus() *string
	SetServiceDescription(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceDescription() *string
	SetServiceDomain(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceDomain() *string
	SetServiceId(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceName() *string
	SetServiceStatus(v string) *CreateVpcEndpointServiceResponseBody
	GetServiceStatus() *string
	SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceResponseBody
	GetServiceSupportIPv6() *bool
	SetSupportedRegionSet(v []*CreateVpcEndpointServiceResponseBodySupportedRegionSet) *CreateVpcEndpointServiceResponseBody
	GetSupportedRegionSet() []*CreateVpcEndpointServiceResponseBodySupportedRegionSet
	SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceResponseBody
	GetZoneAffinityEnabled() *bool
}

type CreateVpcEndpointServiceResponseBody struct {
	// The IP version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Indicates whether endpoint connection requests are automatically accepted. Valid values:
	//
	// - **true**: automatically accepts endpoint connection requests.
	//
	// - **false**: does not automatically accept endpoint connection requests.
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The time when the endpoint service was created.
	//
	// example:
	//
	// 2022-01-02T19:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The business status of the endpoint service. Valid values:
	//
	// - **Normal**: normal.
	//
	// - **FinancialLocked**: locked due to overdue payments.
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
	// The state of the endpoint service. Valid values:
	//
	// - **Creating**: being created.
	//
	// - **Pending**: being modified.
	//
	// - **Active**: available.
	//
	// - **Deleting**: being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Deprecated
	//
	// Indicates whether the endpoint service supports IPv6. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The list of regions in which the endpoint service is available. Service consumers can initiate endpoint connections from the regions in the list.
	SupportedRegionSet []*CreateVpcEndpointServiceResponseBodySupportedRegionSet `json:"SupportedRegionSet,omitempty" xml:"SupportedRegionSet,omitempty" type:"Repeated"`
	// Indicates whether zone affinity is enabled for endpoint domain name resolution. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s CreateVpcEndpointServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateVpcEndpointServiceResponseBody) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *CreateVpcEndpointServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateVpcEndpointServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcEndpointServiceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceBusinessStatus() *string {
	return s.ServiceBusinessStatus
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceDomain() *string {
	return s.ServiceDomain
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *CreateVpcEndpointServiceResponseBody) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *CreateVpcEndpointServiceResponseBody) GetSupportedRegionSet() []*CreateVpcEndpointServiceResponseBodySupportedRegionSet {
	return s.SupportedRegionSet
}

func (s *CreateVpcEndpointServiceResponseBody) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *CreateVpcEndpointServiceResponseBody) SetAddressIpVersion(v string) *CreateVpcEndpointServiceResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceResponseBody {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetCreateTime(v string) *CreateVpcEndpointServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetRequestId(v string) *CreateVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetResourceGroupId(v string) *CreateVpcEndpointServiceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceBusinessStatus(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceDescription(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceDomain(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceDomain = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceId(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceName(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceStatus(v string) *CreateVpcEndpointServiceResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceResponseBody {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetSupportedRegionSet(v []*CreateVpcEndpointServiceResponseBodySupportedRegionSet) *CreateVpcEndpointServiceResponseBody {
	s.SupportedRegionSet = v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBody) Validate() error {
	if s.SupportedRegionSet != nil {
		for _, item := range s.SupportedRegionSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpcEndpointServiceResponseBodySupportedRegionSet struct {
	// The business status of the region in which the endpoint service is available. Valid values:
	//
	// - **Normal**: normal.
	//
	// - **FinancialLocked**: locked due to overdue payments.
	//
	// example:
	//
	// Normal
	RegionBusinessStatus *string `json:"RegionBusinessStatus,omitempty" xml:"RegionBusinessStatus,omitempty"`
	// The state of the region in which the endpoint service is available. Valid values:
	//
	// - **Pending**: changing.
	//
	// - **Available**: available.
	//
	// - **Deleting**: being deleted.
	//
	// - **Failed**: failed.
	//
	// - **Closed**: closed.
	//
	// example:
	//
	// Available
	RegionServiceStatus *string `json:"RegionServiceStatus,omitempty" xml:"RegionServiceStatus,omitempty"`
	// Deprecated
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of a region in which the endpoint service is available.
	//
	// example:
	//
	// cn-hangzhou
	SupportedRegionId *string `json:"SupportedRegionId,omitempty" xml:"SupportedRegionId,omitempty"`
}

func (s CreateVpcEndpointServiceResponseBodySupportedRegionSet) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceResponseBodySupportedRegionSet) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) GetRegionBusinessStatus() *string {
	return s.RegionBusinessStatus
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) GetRegionServiceStatus() *string {
	return s.RegionServiceStatus
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) GetSupportedRegionId() *string {
	return s.SupportedRegionId
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) SetRegionBusinessStatus(v string) *CreateVpcEndpointServiceResponseBodySupportedRegionSet {
	s.RegionBusinessStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) SetRegionServiceStatus(v string) *CreateVpcEndpointServiceResponseBodySupportedRegionSet {
	s.RegionServiceStatus = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) SetServiceRegionId(v string) *CreateVpcEndpointServiceResponseBodySupportedRegionSet {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) SetSupportedRegionId(v string) *CreateVpcEndpointServiceResponseBodySupportedRegionSet {
	s.SupportedRegionId = &v
	return s
}

func (s *CreateVpcEndpointServiceResponseBodySupportedRegionSet) Validate() error {
	return dara.Validate(s)
}
