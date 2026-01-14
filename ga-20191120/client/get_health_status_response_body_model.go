// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroups(v []*GetHealthStatusResponseBodyEndpointGroups) *GetHealthStatusResponseBody
	GetEndpointGroups() []*GetHealthStatusResponseBodyEndpointGroups
	SetHealthStatus(v string) *GetHealthStatusResponseBody
	GetHealthStatus() *string
	SetListenerId(v string) *GetHealthStatusResponseBody
	GetListenerId() *string
	SetRequestId(v string) *GetHealthStatusResponseBody
	GetRequestId() *string
}

type GetHealthStatusResponseBody struct {
	// The information about the endpoint groups.
	EndpointGroups []*GetHealthStatusResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	// The health status of endpoints and endpoint groups. Valid values:
	//
	// 	- **normal**
	//
	// 	- **abnormal**
	//
	// 	- **partiallyAbnormal**
	//
	// example:
	//
	// normal
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBody) GetEndpointGroups() []*GetHealthStatusResponseBodyEndpointGroups {
	return s.EndpointGroups
}

func (s *GetHealthStatusResponseBody) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *GetHealthStatusResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHealthStatusResponseBody) SetEndpointGroups(v []*GetHealthStatusResponseBodyEndpointGroups) *GetHealthStatusResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *GetHealthStatusResponseBody) SetHealthStatus(v string) *GetHealthStatusResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusResponseBody) SetListenerId(v string) *GetHealthStatusResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetHealthStatusResponseBody) SetRequestId(v string) *GetHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHealthStatusResponseBody) Validate() error {
	if s.EndpointGroups != nil {
		for _, item := range s.EndpointGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHealthStatusResponseBodyEndpointGroups struct {
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// 	- **default:*	- the default endpoint group.
	//
	// 	- **virtual:*	- a virtual endpoint group.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The information about the endpoints.
	Endpoints []*GetHealthStatusResponseBodyEndpointGroupsEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The IDs of the forwarding rules.
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	// The health status of the endpoint group. Valid values:
	//
	// 	- **init:*	- The endpoint group is being initialized.
	//
	// 	- **normal:*	- The endpoint group is normal.
	//
	// 	- **abnormal:*	- The endpoint group is abnormal.
	//
	// 	- **partiallyAbnormal:*	- The endpoint group is partially abnormal.
	//
	// example:
	//
	// normal
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
}

func (s GetHealthStatusResponseBodyEndpointGroups) String() string {
	return dara.Prettify(s)
}

func (s GetHealthStatusResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBodyEndpointGroups) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *GetHealthStatusResponseBodyEndpointGroups) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *GetHealthStatusResponseBodyEndpointGroups) GetEndpoints() []*GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	return s.Endpoints
}

func (s *GetHealthStatusResponseBodyEndpointGroups) GetForwardingRuleIds() []*string {
	return s.ForwardingRuleIds
}

func (s *GetHealthStatusResponseBodyEndpointGroups) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpointGroupId(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpointGroupType(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.EndpointGroupType = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpoints(v []*GetHealthStatusResponseBodyEndpointGroupsEndpoints) *GetHealthStatusResponseBodyEndpointGroups {
	s.Endpoints = v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetForwardingRuleIds(v []*string) *GetHealthStatusResponseBodyEndpointGroups {
	s.ForwardingRuleIds = v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetHealthStatus(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHealthStatusResponseBodyEndpointGroupsEndpoints struct {
	// The IP address of the endpoint.
	//
	// example:
	//
	// 47.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The health check details of the endpoint.
	//
	// >  This parameter is unavailable.
	//
	// example:
	//
	// null
	HealthDetail *string `json:"HealthDetail,omitempty" xml:"HealthDetail,omitempty"`
	// The health status of the endpoint. Valid values:
	//
	// 	- **init:*	- The endpoint is being initialized.
	//
	// 	- **normal:*	- The endpoint is normal.
	//
	// 	- **abnormal:*	- The endpoint is abnormal.
	//
	// example:
	//
	// normal
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The port that is used to connect to the endpoint.
	//
	// example:
	//
	// 80
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Domain:*	- a custom domain name.
	//
	// 	- **Ip:*	- a custom IP address.
	//
	// 	- **PublicIp:*	- a public IP address provided by Alibaba Cloud.
	//
	// 	- **ECS:*	- an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB:*	- a Classic Load Balancer (CLB) instance.
	//
	// 	- **ALB:*	- an Application Load Balancer (ALB) instance.
	//
	// 	- **OSS:*	- an Object Storage Service (OSS) bucket.
	//
	// 	- **ENI:*	- an elastic network interface (ENI).
	//
	// 	- **NLB:*	- a Network Load Balancer (NLB) instance.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHealthStatusResponseBodyEndpointGroupsEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetHealthStatusResponseBodyEndpointGroupsEndpoints) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetAddress() *string {
	return s.Address
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetHealthDetail() *string {
	return s.HealthDetail
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetPort() *int64 {
	return s.Port
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) GetType() *string {
	return s.Type
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetAddress(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Address = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetEndpointId(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.EndpointId = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetHealthDetail(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.HealthDetail = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetHealthStatus(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetPort(v int64) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Port = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetType(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Type = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) Validate() error {
	return dara.Validate(s)
}
