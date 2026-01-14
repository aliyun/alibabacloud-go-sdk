// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpSets(v []*ListIpSetsResponseBodyIpSets) *ListIpSetsResponseBody
	GetIpSets() []*ListIpSetsResponseBodyIpSets
	SetPageNumber(v int32) *ListIpSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIpSetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListIpSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIpSetsResponseBody
	GetTotalCount() *int32
}

type ListIpSetsResponseBody struct {
	// The acceleration regions.
	IpSets []*ListIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A9B4E54C-9CCD-4002-91A9-D38C6C209192
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBody) GetIpSets() []*ListIpSetsResponseBodyIpSets {
	return s.IpSets
}

func (s *ListIpSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIpSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIpSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIpSetsResponseBody) SetIpSets(v []*ListIpSetsResponseBodyIpSets) *ListIpSetsResponseBody {
	s.IpSets = v
	return s
}

func (s *ListIpSetsResponseBody) SetPageNumber(v int32) *ListIpSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIpSetsResponseBody) SetPageSize(v int32) *ListIpSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIpSetsResponseBody) SetRequestId(v string) *ListIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpSetsResponseBody) SetTotalCount(v int32) *ListIpSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpSetsResponseBody) Validate() error {
	if s.IpSets != nil {
		for _, item := range s.IpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpSetsResponseBodyIpSets struct {
	// The ID of the acceleration region.
	//
	// example:
	//
	// cn-hangzhou
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The bandwidth that is allocated to the acceleration region. Unit: **Mbit/s**.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The accelerated IP addresses.
	IpAddressList []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// 	- **DUAL_STACK**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values:
	//
	// 	- **BGP**: BGP (Multi-ISP) lines. This is the default value.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter takes effect only if **ServiceManaged*	- is set to **true**.
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
	// > 	- This parameter takes effect only if **ServiceManaged*	- is set to **true**.
	//
	// >	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListIpSetsResponseBodyIpSetsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the acceleration region. Valid values:
	//
	// 	- **init:*	- The acceleration region is being initialized.
	//
	// 	- **active:*	- The acceleration region is in the running state.
	//
	// 	- **updating:*	- The acceleration region is being configured.
	//
	// 	- **deleting:*	- The acceleration region is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListIpSetsResponseBodyIpSets) String() string {
	return dara.Prettify(s)
}

func (s ListIpSetsResponseBodyIpSets) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBodyIpSets) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *ListIpSetsResponseBodyIpSets) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListIpSetsResponseBodyIpSets) GetIpAddressList() []*string {
	return s.IpAddressList
}

func (s *ListIpSetsResponseBodyIpSets) GetIpSetId() *string {
	return s.IpSetId
}

func (s *ListIpSetsResponseBodyIpSets) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListIpSetsResponseBodyIpSets) GetIspType() *string {
	return s.IspType
}

func (s *ListIpSetsResponseBodyIpSets) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListIpSetsResponseBodyIpSets) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListIpSetsResponseBodyIpSets) GetServiceManagedInfos() []*ListIpSetsResponseBodyIpSetsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListIpSetsResponseBodyIpSets) GetState() *string {
	return s.State
}

func (s *ListIpSetsResponseBodyIpSets) SetAccelerateRegionId(v string) *ListIpSetsResponseBodyIpSets {
	s.AccelerateRegionId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetBandwidth(v int32) *ListIpSetsResponseBodyIpSets {
	s.Bandwidth = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpAddressList(v []*string) *ListIpSetsResponseBodyIpSets {
	s.IpAddressList = v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpSetId(v string) *ListIpSetsResponseBodyIpSets {
	s.IpSetId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpVersion(v string) *ListIpSetsResponseBodyIpSets {
	s.IpVersion = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIspType(v string) *ListIpSetsResponseBodyIpSets {
	s.IspType = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetServiceId(v string) *ListIpSetsResponseBodyIpSets {
	s.ServiceId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetServiceManaged(v bool) *ListIpSetsResponseBodyIpSets {
	s.ServiceManaged = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetServiceManagedInfos(v []*ListIpSetsResponseBodyIpSetsServiceManagedInfos) *ListIpSetsResponseBodyIpSets {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetState(v string) *ListIpSetsResponseBodyIpSets {
	s.State = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) Validate() error {
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

type ListIpSetsResponseBodyIpSetsServiceManagedInfos struct {
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
	// >  This parameter takes effect only if **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed.
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the actions on the managed instance.****
	//
	// 	- **false**: The specified actions are not managed, and users can perform the actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListIpSetsResponseBodyIpSetsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListIpSetsResponseBodyIpSetsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) SetAction(v string) *ListIpSetsResponseBodyIpSetsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) SetChildType(v string) *ListIpSetsResponseBodyIpSetsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) SetIsManaged(v bool) *ListIpSetsResponseBodyIpSetsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSetsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
