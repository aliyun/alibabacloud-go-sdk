// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetAcceleratorId() *string
	SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetDestinationId() *string
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetEndpointGroupId() *string
	SetFromPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetFromPort() *int32
	SetListenerId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetListenerId() *string
	SetProtocols(v []*string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetProtocols() []*string
	SetRequestId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetServiceManagedInfos() []*DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos
	SetState(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetState() *string
	SetToPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
	GetToPort() *int32
}

type DescribeCustomRoutingEndpointGroupDestinationsResponseBody struct {
	// The ID of the Global Accelerator (GA) instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the endpoint group mapping configuration.
	//
	// example:
	//
	// dst-123abc****
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The start port of the backend service port range of the endpoint group.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The backend service protocol of the endpoint group.
	//
	// 	- **TCP**: TCP
	//
	// 	- **UDP**: UDP
	//
	// 	- **TCP,UDP**: TCP and UDP
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
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
	ServiceManagedInfos []*DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the endpoint group mapping configuration.
	//
	// 	- **init**: being initialized.
	//
	// 	- **active**: normal.
	//
	// 	- **updating**: being updated.
	//
	// 	- **deleting**: being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The end port of the backend service port range of the endpoint group.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetDestinationId() *string {
	return s.DestinationId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetFromPort() *int32 {
	return s.FromPort
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetProtocols() []*string {
	return s.Protocols
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetServiceManagedInfos() []*DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GetToPort() *int32 {
	return s.ToPort
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.DestinationId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetFromPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.FromPort = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetProtocols(v []*string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.Protocols = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetServiceId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetServiceManaged(v bool) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetState(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.State = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetToPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ToPort = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) Validate() error {
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

type DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos struct {
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

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) SetAction(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
