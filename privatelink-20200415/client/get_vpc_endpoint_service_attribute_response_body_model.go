// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointServiceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody
	GetAutoAcceptEnabled() *bool
	SetConnectBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody
	GetConnectBandwidth() *int32
	SetCreateTime(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetCreateTime() *string
	SetMaxBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody
	GetMaxBandwidth() *int32
	SetMinBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody
	GetMinBandwidth() *int32
	SetPayer(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetPayer() *string
	SetRegionId(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetResourceGroupId() *string
	SetServiceBusinessStatus(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceBusinessStatus() *string
	SetServiceDescription(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceDescription() *string
	SetServiceDomain(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceDomain() *string
	SetServiceId(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceId() *string
	SetServiceName(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceName() *string
	SetServiceResourceType(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceResourceType() *string
	SetServiceStatus(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceStatus() *string
	SetServiceSupportIPv6(v bool) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceSupportIPv6() *bool
	SetServiceType(v string) *GetVpcEndpointServiceAttributeResponseBody
	GetServiceType() *string
	SetSupportedRegionSet(v []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) *GetVpcEndpointServiceAttributeResponseBody
	GetSupportedRegionSet() []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet
	SetZoneAffinityEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody
	GetZoneAffinityEnabled() *bool
	SetZones(v []*string) *GetVpcEndpointServiceAttributeResponseBody
	GetZones() []*string
}

type GetVpcEndpointServiceAttributeResponseBody struct {
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s. Valid values: **100*	- to 10240.
	//
	// example:
	//
	// 1024
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time when the endpoint service was created.
	//
	// example:
	//
	// 2020-01-02T19:11:12Z
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
	// The payer of the endpoint service. Valid values:
	//
	// 	- **Endpoint**: the service consumer.
	//
	// 	- **EndpointService**: the service provider.
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
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service status of the endpoint service. Valid values:
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
	// The endpoint service ID.
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
	// 	- **slb**: a CLB instance.
	//
	// 	- **alb**: an ALB instance.
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
	// 	- **Inactive**: The endpoint service is unavailable.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Deprecated
	//
	// Specifies whether the endpoint service supports IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The type of the endpoint.
	//
	// Only **Interface*	- is returned. The value indicates the interface endpoint. Then, you can specify ALB and CLB instances as service resources for the endpoint service.
	//
	// example:
	//
	// Interface
	ServiceType        *string                                                         `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SupportedRegionSet []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet `json:"SupportedRegionSet,omitempty" xml:"SupportedRegionSet,omitempty" type:"Repeated"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The zones to which the service resources belong.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetVpcEndpointServiceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetConnectBandwidth() *int32 {
	return s.ConnectBandwidth
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetMaxBandwidth() *int32 {
	return s.MaxBandwidth
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetMinBandwidth() *int32 {
	return s.MinBandwidth
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetPayer() *string {
	return s.Payer
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceBusinessStatus() *string {
	return s.ServiceBusinessStatus
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceDomain() *string {
	return s.ServiceDomain
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetSupportedRegionSet() []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet {
	return s.SupportedRegionSet
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *GetVpcEndpointServiceAttributeResponseBody) GetZones() []*string {
	return s.Zones
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetAddressIpVersion(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetAutoAcceptEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetConnectBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetCreateTime(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetMaxBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.MaxBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetMinBandwidth(v int32) *GetVpcEndpointServiceAttributeResponseBody {
	s.MinBandwidth = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetPayer(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.Payer = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetRegionId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetRequestId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetResourceGroupId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceBusinessStatus(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceBusinessStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceDescription(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceDescription = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceDomain(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceDomain = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceId(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceName(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceResourceType(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceResourceType = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceStatus(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceSupportIPv6(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetServiceType(v string) *GetVpcEndpointServiceAttributeResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetSupportedRegionSet(v []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) *GetVpcEndpointServiceAttributeResponseBody {
	s.SupportedRegionSet = v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetZoneAffinityEnabled(v bool) *GetVpcEndpointServiceAttributeResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) SetZones(v []*string) *GetVpcEndpointServiceAttributeResponseBody {
	s.Zones = v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBody) Validate() error {
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

type GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet struct {
	RegionBusinessStatus *string `json:"RegionBusinessStatus,omitempty" xml:"RegionBusinessStatus,omitempty"`
	RegionServiceStatus  *string `json:"RegionServiceStatus,omitempty" xml:"RegionServiceStatus,omitempty"`
	ServiceRegionId      *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) GetRegionBusinessStatus() *string {
	return s.RegionBusinessStatus
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) GetRegionServiceStatus() *string {
	return s.RegionServiceStatus
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) SetRegionBusinessStatus(v string) *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet {
	s.RegionBusinessStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) SetRegionServiceStatus(v string) *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet {
	s.RegionServiceStatus = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) SetServiceRegionId(v string) *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet {
	s.ServiceRegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) Validate() error {
	return dara.Validate(s)
}
