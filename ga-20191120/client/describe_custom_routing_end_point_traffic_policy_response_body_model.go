// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndPointTrafficPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetAcceleratorId() *string
	SetAddress(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetAddress() *string
	SetEndpoint(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetEndpoint() *string
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetEndpointGroupId() *string
	SetEndpointId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetEndpointId() *string
	SetListenerId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetListenerId() *string
	SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetPolicyId() *string
	SetPortRanges(v []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetPortRanges() []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges
	SetRequestId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetServiceManagedInfos() []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos
	SetState(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
	GetState() *string
}

type DescribeCustomRoutingEndPointTrafficPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the traffic policy.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the endpoint to which the traffic policy belongs.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the listener to which the endpoint belongs.
	//
	// example:
	//
	// epg-bp1bpn0kn908w4nb****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// ep-2zewuzypq5e6r3pfh****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the GA instance to which the endpoint belongs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the vSwitch to which the traffic policy belongs.
	//
	// example:
	//
	// ply-bptest2****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The IP address of the traffic policy.
	PortRanges []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The ID of the endpoint to which the traffic destination belongs.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that users can perform on the managed instance.
	//
	// > 	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// > 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the traffic destination.
	//
	// - init: being initialized.
	//
	// - active: running as expected.
	//
	// - updating: being updated.
	//
	// - deleting: being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetPortRanges() []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges {
	return s.PortRanges
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetServiceManagedInfos() []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetAddress(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpoint(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpointId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetPortRanges(v []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.PortRanges = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetServiceId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetServiceManaged(v bool) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetServiceManagedInfos(v []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetState(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.State = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceManagedInfos != nil {
		for _, item := range s.ServiceManagedInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges struct {
	// The first port of the port range used by the traffic destination to process requests.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the port range used by the traffic destination to process requests.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) SetFromPort(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges {
	s.FromPort = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) SetToPort(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges {
	s.ToPort = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos struct {
	// The name of the action performed on the managed instance. Valid values:
	//
	// 	- **Create**
	//
	// 	- **Update**
	//
	// 	- **Delete**
	//
	// 	- **Associate**
	//
	// 	- **UserUnmanaged**
	//
	// 	- **CreateChild**
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// 	- **Listener**: listener.
	//
	// 	- **IpSet**: acceleration region.
	//
	// 	- **EndpointGroup**: endpoint group.
	//
	// 	- **ForwardingRule**: forwarding rule.
	//
	// 	- **Endpoint**: endpoint.
	//
	// 	- **EndpointGroupDestination**: protocol mapping of an endpoint group that is associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: traffic policy of an endpoint that is associated with a custom routing listener.
	//
	// >  This parameter takes effect only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed.
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) SetAction(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
