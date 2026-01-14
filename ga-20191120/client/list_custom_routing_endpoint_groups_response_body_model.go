// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroups(v []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) *ListCustomRoutingEndpointGroupsResponseBody
	GetEndpointGroups() []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups
	SetPageNumber(v int32) *ListCustomRoutingEndpointGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomRoutingEndpointGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomRoutingEndpointGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomRoutingEndpointGroupsResponseBody
	GetTotalCount() *int32
}

type ListCustomRoutingEndpointGroupsResponseBody struct {
	// The configuration information about the endpoint group.
	EndpointGroups []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) GetEndpointGroups() []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	return s.EndpointGroups
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetEndpointGroups(v []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) *ListCustomRoutingEndpointGroupsResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) Validate() error {
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

type ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The description of the endpoint group.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoint group IP addresses.
	EndpointGroupIpList []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	// The ID of the region where the endpoint group is created.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The endpoint group IP addresses to be confirmed after the GA instance is upgraded.
	EndpointGroupUnconfirmedIpList []*string `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	// The ID of the custom routing listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the endpoint group.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter takes effect only if **ServiceManaged*	- is set to **True**.
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
	// >
	//
	// 	- This parameter takes effect only if **ServiceManaged*	- is set to **True**.
	//
	// 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the endpoint group. Valid values:
	//
	// 	- **init**
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **deleting**
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetDescription() *string {
	return s.Description
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupIpList() []*string {
	return s.EndpointGroupIpList
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupUnconfirmedIpList() []*string {
	return s.EndpointGroupUnconfirmedIpList
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetName() *string {
	return s.Name
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetServiceManagedInfos() []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GetState() *string {
	return s.State
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetDescription(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.Description = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupIpList(v []*string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupIpList = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupRegion(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupUnconfirmedIpList(v []*string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetListenerId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetName(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.Name = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetServiceId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetServiceManaged(v bool) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetServiceManagedInfos(v []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetState(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.State = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) Validate() error {
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

type ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos struct {
	// The name of the action on the managed instance. Valid values:
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
	// >  This parameter takes effect only if **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed.
	//
	// 	- **true**: Users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: Users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetAction(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetChildType(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetIsManaged(v bool) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
