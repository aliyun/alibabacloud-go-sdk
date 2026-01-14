// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointTrafficPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody
	GetPageSize() *int32
	SetPolicies(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) *ListCustomRoutingEndpointTrafficPoliciesResponseBody
	GetPolicies() []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of traffic policies.
	Policies []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) GetPolicies() []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicies(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies struct {
	// The ID of the GA instance to which the endpoint belongs.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The IP addresses of the traffic policies.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint to which the traffic policy belongs.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the custom routing listener to which the endpoint belongs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the traffic policy.
	//
	// example:
	//
	// ply-bp1dmlohjjz4kqaun****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The port range of the traffic policy.
	PortRanges []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
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
	// 	- **true**: The GA instance is managed.
	//
	// 	- **false**: The GA instance is not managed.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that users can perform on the managed instance.
	//
	// > 	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// >	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetAddress() *string {
	return s.Address
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetPortRanges() []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges {
	return s.PortRanges
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GetServiceManagedInfos() []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetAcceleratorId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetAddress(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.Address = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetEndpointGroupId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetEndpointId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetListenerId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetPortRanges(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.PortRanges = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetServiceId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.ServiceId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetServiceManaged(v bool) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.ServiceManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetServiceManagedInfos(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) Validate() error {
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

type ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges struct {
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

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) SetFromPort(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) SetToPort(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges {
	s.ToPort = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) Validate() error {
	return dara.Validate(s)
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos struct {
	// The name of the action on the managed instance. Valid values:
	//
	// 	- **Create**: Create an instance.
	//
	// 	- **Update**: Update the current instance.
	//
	// 	- **Delete**: Delete the current instance.
	//
	// 	- **Associate**: Reference the current instance.
	//
	// 	- **UserUnmanaged**: Unmanage the instance.
	//
	// 	- **CreateChild**: Create a child resource in the current instance.
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
	// 	- **EndpointGroupDestination**: protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy**: traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter is returned only if the value of **Action*	- is **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
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

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) SetAction(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) SetChildType(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) SetIsManaged(v bool) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
