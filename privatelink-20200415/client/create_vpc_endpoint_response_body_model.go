// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *CreateVpcEndpointResponseBody
	GetAddressIpVersion() *string
	SetBandwidth(v int64) *CreateVpcEndpointResponseBody
	GetBandwidth() *int64
	SetConnectionStatus(v string) *CreateVpcEndpointResponseBody
	GetConnectionStatus() *string
	SetCreateTime(v string) *CreateVpcEndpointResponseBody
	GetCreateTime() *string
	SetCrossRegionBandwidth(v int32) *CreateVpcEndpointResponseBody
	GetCrossRegionBandwidth() *int32
	SetEndpointBusinessStatus(v string) *CreateVpcEndpointResponseBody
	GetEndpointBusinessStatus() *string
	SetEndpointDescription(v string) *CreateVpcEndpointResponseBody
	GetEndpointDescription() *string
	SetEndpointDomain(v string) *CreateVpcEndpointResponseBody
	GetEndpointDomain() *string
	SetEndpointId(v string) *CreateVpcEndpointResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *CreateVpcEndpointResponseBody
	GetEndpointName() *string
	SetEndpointStatus(v string) *CreateVpcEndpointResponseBody
	GetEndpointStatus() *string
	SetRequestId(v string) *CreateVpcEndpointResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateVpcEndpointResponseBody
	GetServiceId() *string
	SetServiceName(v string) *CreateVpcEndpointResponseBody
	GetServiceName() *string
	SetServiceRegionId(v string) *CreateVpcEndpointResponseBody
	GetServiceRegionId() *string
	SetVpcId(v string) *CreateVpcEndpointResponseBody
	GetVpcId() *string
	SetZoneAffinityEnabled(v bool) *CreateVpcEndpointResponseBody
	GetZoneAffinityEnabled() *bool
}

type CreateVpcEndpointResponseBody struct {
	// The protocol. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **DualStack**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The connection is being modified.
	//
	// 	- **Connecting**: The connection is being established.
	//
	// 	- **Connected**: The connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2022-01-02T19:11:12Z
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossRegionBandwidth *int32  `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The domain name of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the VPC to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneAffinityEnabled *bool   `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s CreateVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateVpcEndpointResponseBody) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateVpcEndpointResponseBody) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *CreateVpcEndpointResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateVpcEndpointResponseBody) GetCrossRegionBandwidth() *int32 {
	return s.CrossRegionBandwidth
}

func (s *CreateVpcEndpointResponseBody) GetEndpointBusinessStatus() *string {
	return s.EndpointBusinessStatus
}

func (s *CreateVpcEndpointResponseBody) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *CreateVpcEndpointResponseBody) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *CreateVpcEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateVpcEndpointResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateVpcEndpointResponseBody) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *CreateVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcEndpointResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateVpcEndpointResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateVpcEndpointResponseBody) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *CreateVpcEndpointResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcEndpointResponseBody) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *CreateVpcEndpointResponseBody) SetAddressIpVersion(v string) *CreateVpcEndpointResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetBandwidth(v int64) *CreateVpcEndpointResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetConnectionStatus(v string) *CreateVpcEndpointResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetCreateTime(v string) *CreateVpcEndpointResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetCrossRegionBandwidth(v int32) *CreateVpcEndpointResponseBody {
	s.CrossRegionBandwidth = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointBusinessStatus(v string) *CreateVpcEndpointResponseBody {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointDescription(v string) *CreateVpcEndpointResponseBody {
	s.EndpointDescription = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointDomain(v string) *CreateVpcEndpointResponseBody {
	s.EndpointDomain = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointId(v string) *CreateVpcEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointName(v string) *CreateVpcEndpointResponseBody {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetEndpointStatus(v string) *CreateVpcEndpointResponseBody {
	s.EndpointStatus = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetRequestId(v string) *CreateVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetServiceId(v string) *CreateVpcEndpointResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetServiceName(v string) *CreateVpcEndpointResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetServiceRegionId(v string) *CreateVpcEndpointResponseBody {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetVpcId(v string) *CreateVpcEndpointResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
