// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegionId(v string) *DescribeIpSetResponseBody
	GetAccelerateRegionId() *string
	SetAcceleratorId(v string) *DescribeIpSetResponseBody
	GetAcceleratorId() *string
	SetBandwidth(v int32) *DescribeIpSetResponseBody
	GetBandwidth() *int32
	SetIpAddressList(v []*string) *DescribeIpSetResponseBody
	GetIpAddressList() []*string
	SetIpSetId(v string) *DescribeIpSetResponseBody
	GetIpSetId() *string
	SetIpVersion(v string) *DescribeIpSetResponseBody
	GetIpVersion() *string
	SetIspType(v string) *DescribeIpSetResponseBody
	GetIspType() *string
	SetRequestId(v string) *DescribeIpSetResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeIpSetResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeIpSetResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeIpSetResponseBodyServiceManagedInfos) *DescribeIpSetResponseBody
	GetServiceManagedInfos() []*DescribeIpSetResponseBodyServiceManagedInfos
	SetState(v string) *DescribeIpSetResponseBody
	GetState() *string
}

type DescribeIpSetResponseBody struct {
	// The ID of the acceleration region.
	//
	// example:
	//
	// cn-hangzhou
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmqy****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth that is allocated to the acceleration region. Unit: Mbit/s.
	//
	// example:
	//
	// 3
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressList []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// ips-bp11ilwqjdkjeg9r7****
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
	// The ID of the request.
	//
	// example:
	//
	// 6D2BFF54-6AF2-4679-88C4-2F2E187F16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service that manages the instance.
	//
	// >  This parameter is returned only if the value of **ServiceManaged*	- is **true**.
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
	// > 	- This parameter is returned only if the value of **ServiceManaged*	- is **true**.
	//
	// >	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*DescribeIpSetResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the acceleration region. Valid values:
	//
	// 	- **init**: The acceleration region is being initialized.
	//
	// 	- **active**: The acceleration region is in the running state.
	//
	// 	- **updating**: The acceleration region is being configured.
	//
	// 	- **deleting**: The GA instance is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponseBody) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *DescribeIpSetResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeIpSetResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeIpSetResponseBody) GetIpAddressList() []*string {
	return s.IpAddressList
}

func (s *DescribeIpSetResponseBody) GetIpSetId() *string {
	return s.IpSetId
}

func (s *DescribeIpSetResponseBody) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeIpSetResponseBody) GetIspType() *string {
	return s.IspType
}

func (s *DescribeIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpSetResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeIpSetResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeIpSetResponseBody) GetServiceManagedInfos() []*DescribeIpSetResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeIpSetResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeIpSetResponseBody) SetAccelerateRegionId(v string) *DescribeIpSetResponseBody {
	s.AccelerateRegionId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetAcceleratorId(v string) *DescribeIpSetResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetBandwidth(v int32) *DescribeIpSetResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpAddressList(v []*string) *DescribeIpSetResponseBody {
	s.IpAddressList = v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpSetId(v string) *DescribeIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpVersion(v string) *DescribeIpSetResponseBody {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIspType(v string) *DescribeIpSetResponseBody {
	s.IspType = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetRequestId(v string) *DescribeIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetServiceId(v string) *DescribeIpSetResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetServiceManaged(v bool) *DescribeIpSetResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetServiceManagedInfos(v []*DescribeIpSetResponseBodyServiceManagedInfos) *DescribeIpSetResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeIpSetResponseBody) SetState(v string) *DescribeIpSetResponseBody {
	s.State = &v
	return s
}

func (s *DescribeIpSetResponseBody) Validate() error {
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

type DescribeIpSetResponseBodyServiceManagedInfos struct {
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
	// 	- **true**: The specified actions are managed, and users cannot perform the actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeIpSetResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpSetResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) SetAction(v string) *DescribeIpSetResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeIpSetResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeIpSetResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeIpSetResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
