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
	// The IP address version. Valid values:
	//
	// - **IPv4**: The IPv4 address family.
	//
	// - **DualStack**: The dual stack IP address family.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connections. Valid values:
	//
	// - **true**: Endpoint connections are automatically accepted.
	//
	// - **false**: Endpoint connections are not automatically accepted.
	//
	// example:
	//
	// true
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The default maximum bandwidth of an endpoint connection, in Mbps. Valid values are **100 to 10,240**.
	//
	// example:
	//
	// 3072
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The time the service was created.
	//
	// example:
	//
	// 2020-01-02T19:11:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The maximum bandwidth of the endpoint connection, in Mbps.
	//
	// example:
	//
	// 1024
	MaxBandwidth *int32 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// The minimum bandwidth of the endpoint connection, in Mbps.
	//
	// example:
	//
	// 100
	MinBandwidth *int32 `json:"MinBandwidth,omitempty" xml:"MinBandwidth,omitempty"`
	// The party that pays for the service. Valid values:
	//
	// - **Endpoint**: The service consumer.
	//
	// - **EndpointService**: The service provider.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region where the service is deployed.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The business status of the service. Valid values:
	//
	// - **Normal**: The service is operating normally.
	//
	// - **FinancialLocked**: The endpoint service is locked due to an overdue payment.
	//
	// example:
	//
	// Normal
	ServiceBusinessStatus *string `json:"ServiceBusinessStatus,omitempty" xml:"ServiceBusinessStatus,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The domain name of the service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****.cn-huhehaote.privatelink.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service resource. Valid values:
	//
	// - **slb**: The service resource is a Classic Load Balancer (CLB).
	//
	// - **alb**: The service resource is an Application Load Balancer (ALB).
	//
	// - **nlb**: The service resource is a Network Load Balancer (NLB).
	//
	// - **gwlb**: The service resource is a Gateway Load Balancer (GWLB).
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// The status of the service. Valid values:
	//
	// - **Creating**: The service is being created.
	//
	// - **Pending**: The service is being updated.
	//
	// - **Active**: The service is available.
	//
	// - **Deleting**: The service is being deleted.
	//
	// example:
	//
	// Active
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// Deprecated
	//
	// Specifies whether the service supports IPv6. Valid values:
	//
	// - **true**: The service supports IPv6.
	//
	// - **false*	- (default): The service does not support IPv6.
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The endpoint type.
	//
	// - **Interface**: An interface endpoint. You can add Classic Load Balancer (CLB), Application Load Balancer (ALB), and Network Load Balancer (NLB) instances as service resources.
	//
	// - **GatewayLoadBalancer**: A gateway endpoint. You can add Gateway Load Balancer (GWLB) instances as service resources.
	//
	// example:
	//
	// Interface
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The regions where the service is supported. Service consumers can create endpoint connections to the service from these regions.
	SupportedRegionSet []*GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet `json:"SupportedRegionSet,omitempty" xml:"SupportedRegionSet,omitempty" type:"Repeated"`
	// Specifies whether zone affinity is enabled. Valid values:
	//
	// - **true*	- (default): Zone affinity is enabled.
	//
	// - **false**: Zone affinity is disabled.
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The zones where the service is available.
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
	// The business status of the endpoint service in the supported region. Valid values:
	//
	// - **Normal**: The service is operating normally in the supported region.
	//
	// - **FinancialLocked**: The endpoint service is locked due to an overdue payment.
	//
	// example:
	//
	// Normal
	RegionBusinessStatus *string `json:"RegionBusinessStatus,omitempty" xml:"RegionBusinessStatus,omitempty"`
	// The status of the endpoint service in the supported region. Valid values:
	//
	// - **Pending**: The supported region is being modified.
	//
	// - **Available**: The service is available in the supported region.
	//
	// - **Deleting**: The supported region is being deleted.
	//
	// - **Failed**: The service failed to be deployed in the supported region.
	//
	// - **Closed**: The endpoint service is not available in the supported region.
	//
	// example:
	//
	// Available
	RegionServiceStatus *string `json:"RegionServiceStatus,omitempty" xml:"RegionServiceStatus,omitempty"`
	// Deprecated
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the supported region.
	//
	// example:
	//
	// cn-hangzhou
	SupportedRegionId *string `json:"SupportedRegionId,omitempty" xml:"SupportedRegionId,omitempty"`
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

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) GetSupportedRegionId() *string {
	return s.SupportedRegionId
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

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) SetSupportedRegionId(v string) *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet {
	s.SupportedRegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponseBodySupportedRegionSet) Validate() error {
	return dara.Validate(s)
}
