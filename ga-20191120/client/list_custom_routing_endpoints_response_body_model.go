// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ListCustomRoutingEndpointsResponseBodyEndpoints) *ListCustomRoutingEndpointsResponseBody
	GetEndpoints() []*ListCustomRoutingEndpointsResponseBodyEndpoints
	SetPageNumber(v int32) *ListCustomRoutingEndpointsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomRoutingEndpointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingEndpointsResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingEndpointsResponseBody struct {
	// The information about the endpoints.
	Endpoints []*ListCustomRoutingEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponseBody) GetEndpoints() []*ListCustomRoutingEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *ListCustomRoutingEndpointsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingEndpointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingEndpointsResponseBody) SetEndpoints(v []*ListCustomRoutingEndpointsResponseBodyEndpoints) *ListCustomRoutingEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) Validate() error {
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

type ListCustomRoutingEndpointsResponseBodyEndpoints struct {
	// The ID of the GA instance with which the endpoint is associated.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The name of the vSwitch that is specified as an endpoint.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// epg-bp16jdc00bhe97sr5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the listener to which the endpoint belongs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the service that manages the GA instance.
	//
	// >  This parameter is valid only if **ServiceManaged*	- is set to **True**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the GA instance is managed. Valid values:
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
	// >  This parameter is valid only if **ServiceManaged*	- is set to **True**.
	//
	// 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The access policy of traffic that is destined for the endpoint. Valid values:
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
	// Only **PrivateSubNet*	- may be returned, which indicates a private CIDR block.
	//
	// example:
	//
	// PrivateSubNet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomRoutingEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetServiceManagedInfos() []*ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) GetType() *string {
	return s.Type
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetAcceleratorId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpoint(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpointGroupId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetListenerId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetServiceId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.ServiceId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetServiceManaged(v bool) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.ServiceManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetServiceManagedInfos(v []*ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetTrafficToEndpointPolicy(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetType(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.Type = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) Validate() error {
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

type ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos struct {
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
	// 	- **Listener**: listener
	//
	// 	- **IpSet**: acceleration region
	//
	// 	- **EndpointGroup**: endpoint group
	//
	// 	- **ForwardingRule**: forwarding rule
	//
	// 	- **Endpoint**: endpoint
	//
	// 	- **EndpointGroupDestination**: protocol mapping of an endpoint group associated with a custom routing listener
	//
	// 	- **EndpointPolicy**: traffic policy of an endpoint associated with a custom routing listener
	//
	// >  This parameter is valid only if **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed resource.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed resource.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) SetAction(v string) *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) SetChildType(v string) *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) SetIsManaged(v bool) *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpointsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
