// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnections(v []*ListVpcEndpointConnectionsResponseBodyConnections) *ListVpcEndpointConnectionsResponseBody
	GetConnections() []*ListVpcEndpointConnectionsResponseBodyConnections
	SetMaxResults(v int32) *ListVpcEndpointConnectionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointConnectionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListVpcEndpointConnectionsResponseBody
	GetTotalCount() *string
}

type ListVpcEndpointConnectionsResponseBody struct {
	// The endpoint connections.
	Connections []*ListVpcEndpointConnectionsResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
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
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBody) GetConnections() []*ListVpcEndpointConnectionsResponseBodyConnections {
	return s.Connections
}

func (s *ListVpcEndpointConnectionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointConnectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointConnectionsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListVpcEndpointConnectionsResponseBody) SetConnections(v []*ListVpcEndpointConnectionsResponseBodyConnections) *ListVpcEndpointConnectionsResponseBody {
	s.Connections = v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetMaxResults(v int32) *ListVpcEndpointConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetNextToken(v string) *ListVpcEndpointConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetRequestId(v string) *ListVpcEndpointConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) SetTotalCount(v string) *ListVpcEndpointConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBody) Validate() error {
	if s.Connections != nil {
		for _, item := range s.Connections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointConnectionsResponseBodyConnections struct {
	// The bandwidth of the endpoint connection. Valid values: **1024 to 10240**. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
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
	// 	- **ServiceDeleted**: The corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the Alibaba Cloud account to which the endpoint belongs.
	//
	// example:
	//
	// 25346073170691****
	EndpointOwnerId  *int64  `json:"EndpointOwnerId,omitempty" xml:"EndpointOwnerId,omitempty"`
	EndpointRegionId *string `json:"EndpointRegionId,omitempty" xml:"EndpointRegionId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	EndpointVpcId *string `json:"EndpointVpcId,omitempty" xml:"EndpointVpcId,omitempty"`
	// The time when the endpoint connection was modified.
	//
	// example:
	//
	// 2021-09-28T10:34:46Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the resource group to which the endpoint belongs.
	//
	// example:
	//
	// rg-acfmw353z35v***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The endpoint service ID.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId          *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceRegionId    *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	TrafficControlMode *string `json:"TrafficControlMode,omitempty" xml:"TrafficControlMode,omitempty"`
	// The zones.
	Zones []*ListVpcEndpointConnectionsResponseBodyConnectionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointConnectionsResponseBodyConnections) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetEndpointOwnerId() *int64 {
	return s.EndpointOwnerId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetEndpointRegionId() *string {
	return s.EndpointRegionId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetEndpointVpcId() *string {
	return s.EndpointVpcId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetResourceOwner() *bool {
	return s.ResourceOwner
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetTrafficControlMode() *string {
	return s.TrafficControlMode
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) GetZones() []*ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	return s.Zones
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetBandwidth(v int32) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.Bandwidth = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetConnectionStatus(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ConnectionStatus = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointOwnerId(v int64) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointOwnerId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointRegionId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointRegionId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetEndpointVpcId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.EndpointVpcId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetModifiedTime(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ModifiedTime = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetResourceGroupId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetResourceOwner(v bool) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ResourceOwner = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetServiceId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetServiceRegionId(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.ServiceRegionId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetTrafficControlMode(v string) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.TrafficControlMode = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) SetZones(v []*ListVpcEndpointConnectionsResponseBodyConnectionsZones) *ListVpcEndpointConnectionsResponseBodyConnections {
	s.Zones = v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnections) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointConnectionsResponseBodyConnectionsZones struct {
	// The endpoint ENI ID.
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The ID of the replaced endpoint ENI in smooth migration scenarios.
	//
	// example:
	//
	// eni-hp32lk0pyv6o94hs****
	ReplacedEniId *string `json:"ReplacedEniId,omitempty" xml:"ReplacedEniId,omitempty"`
	// The ID of the replaced service resource in smooth migration scenarios.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ReplacedResourceId *string `json:"ReplacedResourceId,omitempty" xml:"ReplacedResourceId,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the vSwitch to which the endpoint belongs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The domain name of the zone.
	//
	// example:
	//
	// ep-hp34jaz8ykl0exwt****-cn-huhehaote.epsrv-hp3vpx8yqxblby3i****.cn-huhehaote-b.privatelink.aliyuncs.com
	ZoneDomain *string `json:"ZoneDomain,omitempty" xml:"ZoneDomain,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The state of the zone. Valid values:
	//
	// 	- **Creating**: The zone is being created.
	//
	// 	- **Wait**: The zone is to be connected.
	//
	// 	- **Connected**: The zone is connected.
	//
	// 	- **Deleting**: The zone is being deleted.
	//
	// 	- **Disconnecting**: The zone is being disconnected.
	//
	// 	- **Disconnected**: The zone is disconnected.
	//
	// 	- **Connecting**: The zone is being connected.
	//
	// 	- **Migrating**: The zone is being migrated.
	//
	// 	- **Migrated**: The zone is migrated.
	//
	// example:
	//
	// Connected
	ZoneStatus *string `json:"ZoneStatus,omitempty" xml:"ZoneStatus,omitempty"`
}

func (s ListVpcEndpointConnectionsResponseBodyConnectionsZones) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponseBodyConnectionsZones) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetEniId() *string {
	return s.EniId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetReplacedEniId() *string {
	return s.ReplacedEniId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetReplacedResourceId() *string {
	return s.ReplacedResourceId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetZoneDomain() *string {
	return s.ZoneDomain
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) GetZoneStatus() *string {
	return s.ZoneStatus
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetEniId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.EniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetReplacedEniId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ReplacedEniId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetReplacedResourceId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ReplacedResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetResourceId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetVSwitchId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneDomain(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneDomain = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneId(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneId = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) SetZoneStatus(v string) *ListVpcEndpointConnectionsResponseBodyConnectionsZones {
	s.ZoneStatus = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponseBodyConnectionsZones) Validate() error {
	return dara.Validate(s)
}
