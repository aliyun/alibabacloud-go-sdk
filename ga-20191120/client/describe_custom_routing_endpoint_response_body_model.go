// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetAcceleratorId() *string
	SetEndpoint(v string) *DescribeCustomRoutingEndpointResponseBody
	GetEndpoint() *string
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetEndpointGroupId() *string
	SetEndpointId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetEndpointId() *string
	SetListenerId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetListenerId() *string
	SetRequestId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeCustomRoutingEndpointResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeCustomRoutingEndpointResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointResponseBody
	GetServiceManagedInfos() []*DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos
	SetState(v string) *DescribeCustomRoutingEndpointResponseBody
	GetState() *string
	SetTrafficToEndpointPolicy(v string) *DescribeCustomRoutingEndpointResponseBody
	GetTrafficToEndpointPolicy() *string
	SetType(v string) *DescribeCustomRoutingEndpointResponseBody
	GetType() *string
}

type DescribeCustomRoutingEndpointResponseBody struct {
	// The ID of the GA instance with which the endpoint is associated.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The name of the endpoint (vSwitch).
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp1dmlohjjz4kqaun****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the listener with which the endpoint is associated.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// String	04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID to which the managed instance belongs.
	//
	// >  Valid only when the ServiceManaged parameter is True.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Is it a managed instance. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// A list of action policies that users can execute on this managed instance.
	ServiceManagedInfos []*DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the endpoint.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The access policy of traffic for the specified endpoint. Valid values:
	//
	// 	- **AllowAll**: allows all traffic to the endpoint.
	//
	// 	- **DenyAll**: denies all traffic to the endpoint.
	//
	// 	- **AllowCustom**: allows traffic only to specified destinations.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The backend service type of the endpoint.
	//
	// Set the value to **PrivateSubNet**, which indicates private CIDR blocks.
	//
	// example:
	//
	// PrivateSubNet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCustomRoutingEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetServiceManagedInfos() []*DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *DescribeCustomRoutingEndpointResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpoint(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpointId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetServiceId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetServiceManaged(v bool) *DescribeCustomRoutingEndpointResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetState(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.State = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetTrafficToEndpointPolicy(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetType(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) Validate() error {
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

type DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos struct {
	// Managed policy action name, Valid values:
	//
	// - Create
	//
	// - Update
	//
	// - Delete
	//
	// - Associate
	//
	// - UserUnmanaged
	//
	// - CreateChild
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Sub resource type, Valid values:
	//
	// - Listener
	//
	// - IpSet
	//
	// - EndpointGroup
	//
	// - ForwardingRule
	//
	// - Endpoint
	//
	// - EndpointGroupDestination
	//
	// - EndpointPolicy
	//
	// >Only valid when the Action parameter is CreateChild.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Is the managed policy action managed, Valid values:
	//
	// - true: The managed policy action is managed, and users do not have permission to perform the operation specified in the Action on the managed instance.
	//
	// - false: The managed policy action is not managed, and users have permission to perform the operation specified in the Action on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) SetAction(v string) *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
