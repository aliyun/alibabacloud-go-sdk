// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDestinations(v []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) *ListCustomRoutingEndpointGroupDestinationsResponseBody
	GetDestinations() []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations
	SetPageNumber(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingEndpointGroupDestinationsResponseBody struct {
	// The destination configurations of the endpoint group.
	Destinations []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Repeated"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paging query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s ListCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) GetDestinations() []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	return s.Destinations
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinations(v []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.Destinations = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) Validate() error {
	if s.Destinations != nil {
		for _, item := range s.Destinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations struct {
	// The instance ID of the Alibaba Cloud Global Accelerator (GA) instance to which the endpoint group destination configuration belongs.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the endpoint group destination configuration.
	//
	// example:
	//
	// dst-123abc****
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The ID of the endpoint group to which the destination configuration belongs.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The start port of the backend service of the endpoint group.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The ID of the listener to which the endpoint group destination configuration belongs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The Protocol Type of the backend service of the endpoint group.
	//
	// - **TCP**: TCP protocol.
	//
	// - **UDP**: UDP protocol.
	//
	// - **TCP,UDP**: TCP and UDP protocols.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The ID of the service to which the managed instance belongs.
	//
	// > This parameter is valid only when **ServiceManaged*	- is set to **True**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// - true: The instance is managed.
	//
	// - false: The instance is not managed.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The list of action policies that the user can execute on the managed instance.
	//
	// > This parameter is valid only when **ServiceManaged*	- is set to **True**.
	//
	// > - When the instance is in the managed state, user operations on the instance are restricted, and certain operations are prohibited.
	ServiceManagedInfos []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The end port of the backend service of the endpoint group.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetDestinationId() *string {
	return s.DestinationId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetFromPort() *int32 {
	return s.FromPort
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetServiceManagedInfos() []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GetToPort() *int32 {
	return s.ToPort
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetDestinationId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.DestinationId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetFromPort(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetListenerId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetProtocols(v []*string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetServiceId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ServiceId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetServiceManaged(v bool) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ServiceManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetServiceManagedInfos(v []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetToPort(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ToPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) Validate() error {
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

type ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos struct {
	// The name of the managed policy action. Valid values:
	//
	// - **Create**: Create an instance.
	//
	// - **Update**: Update the current instance.
	//
	// - **Delete**: Delete the current instance.
	//
	// - **Associate**: Reference or be referenced by the current instance.
	//
	// - **UserUnmanaged**: Unmanage the instance.
	//
	// - **CreateChild**: Create a child resource under the current instance.
	//
	// example:
	//
	// Create
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// - **Listener**: listener resource.
	//
	// - **IpSet**: acceleration region resource.
	//
	// - **EndpointGroup**: endpoint group resource.
	//
	// - **ForwardingRule**: forwarding rule resource.
	//
	// - **Endpoint**: endpoint resource.
	//
	// - **EndpointGroupDestination**: protocol mapping resource of the endpoint group under a custom routing listener.
	//
	// - **EndpointPolicy**: endpoint traffic policy resource under a custom routing listener.
	//
	// > This parameter is valid only when **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the managed policy action is managed. Valid values:
	//
	// - **true**: The managed policy action is managed. The user cannot perform the action specified by Action on the managed instance.
	//
	// - **false**: The managed policy action is not managed. The user can perform the action specified by Action on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) SetAction(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) SetChildType(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) SetIsManaged(v bool) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinationsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
