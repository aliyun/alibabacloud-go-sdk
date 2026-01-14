// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetAcceleratorId() *string
	SetAccessLogSwitch(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetAccessLogSwitch() *string
	SetDescription(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetDescription() *string
	SetEnableAccessLog(v bool) *DescribeCustomRoutingEndpointGroupResponseBody
	GetEnableAccessLog() *bool
	SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetEndpointGroupIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetEndpointGroupIpList() []*string
	SetEndpointGroupRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetEndpointGroupRegion() *string
	SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetEndpointGroupUnconfirmedIpList() []*string
	SetListenerId(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetListenerId() *string
	SetName(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeCustomRoutingEndpointGroupResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointGroupResponseBody
	GetServiceManagedInfos() []*DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos
	SetSlsLogStoreName(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetSlsLogStoreName() *string
	SetSlsProjectName(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetSlsProjectName() *string
	SetSlsRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetSlsRegion() *string
	SetState(v string) *DescribeCustomRoutingEndpointGroupResponseBody
	GetState() *string
}

type DescribeCustomRoutingEndpointGroupResponseBody struct {
	// The GA instance ID.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Indicates the status of the binding between the Log Service project and the endpoint group. Valid values:
	//
	// 	- **on:*	- The endpoint group is bound to the Log Service project.
	//
	// 	- **off:*	- The endpoint group is not bound to the Log Service project.
	//
	// 	- **binding:*	- The endpoint group is being bound to the Log Service project.
	//
	// 	- **unbinding:*	- The endpoint group is being unbound from the Log Service project.
	//
	// example:
	//
	// on
	AccessLogSwitch *string `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	// The description of the endpoint group.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the access log feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableAccessLog *bool `json:"EnableAccessLog,omitempty" xml:"EnableAccessLog,omitempty"`
	// The endpoint group ID.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaua****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoint group IP addresses.
	EndpointGroupIpList []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	// The region ID of the endpoint group.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The endpoint group IP addresses that need to be confirmed after the GA instance is upgraded.
	EndpointGroupUnconfirmedIpList []*string `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	// The custom routing listener ID.
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
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service that manages the GA instance.
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
	// >  This parameter takes effect only if **ServiceManaged*	- is set to **True**.
	//
	// 	- Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The name of the Logstore.
	//
	// example:
	//
	// lsn-01
	SlsLogStoreName *string `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// pn-01
	SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	// The region of the logs that are created in Log Service.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegion *string `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
	// The status of the endpoint group. Valid values:
	//
	// 	- **init:*	- The endpoint group is being initialized.
	//
	// 	- **active:*	- The endpoint group is running normally.
	//
	// 	- **updating:*	- The endpoint group is being updated.
	//
	// 	- **deleting:*	- The ACL is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetAccessLogSwitch() *string {
	return s.AccessLogSwitch
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetEnableAccessLog() *bool {
	return s.EnableAccessLog
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetEndpointGroupIpList() []*string {
	return s.EndpointGroupIpList
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetEndpointGroupUnconfirmedIpList() []*string {
	return s.EndpointGroupUnconfirmedIpList
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetServiceManagedInfos() []*DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetSlsLogStoreName() *string {
	return s.SlsLogStoreName
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetSlsRegion() *string {
	return s.SlsRegion
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetAccessLogSwitch(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.AccessLogSwitch = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetDescription(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEnableAccessLog(v bool) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EnableAccessLog = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupIpList = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetServiceId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetServiceManaged(v bool) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetServiceManagedInfos(v []*DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetState(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.State = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) Validate() error {
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

type DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos struct {
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

func (s DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) SetAction(v string) *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
