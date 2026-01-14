// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody
	GetDomains() []*ListDomainsResponseBodyDomains
	SetPageNumber(v int32) *ListDomainsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDomainsResponseBody
	GetTotalCount() *int32
}

type ListDomainsResponseBody struct {
	// A list of accelerated domain names.
	Domains []*ListDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
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
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) GetDomains() []*ListDomainsResponseBodyDomains {
	return s.Domains
}

func (s *ListDomainsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDomainsResponseBody) SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetPageNumber(v int32) *ListDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsResponseBody) SetPageSize(v int32) *ListDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) SetTotalCount(v int32) *ListDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDomainsResponseBody) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainsResponseBodyDomains struct {
	// A list of GA instances.
	Accelerators []*ListDomainsResponseBodyDomainsAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The accelerated domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ICP filing status of the accelerated domain name. Valid values:
	//
	// 	- **illegal:*	- The domain name is illegal.
	//
	// 	- **inactive:*	- The domain name has not completed ICP filing.
	//
	// 	- **active:*	- The domain name has a valid ICP number.
	//
	// 	- **unknown:*	- The ICP filing status is unknown.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomains) GetAccelerators() []*ListDomainsResponseBodyDomainsAccelerators {
	return s.Accelerators
}

func (s *ListDomainsResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *ListDomainsResponseBodyDomains) GetState() *string {
	return s.State
}

func (s *ListDomainsResponseBodyDomains) SetAccelerators(v []*ListDomainsResponseBodyDomainsAccelerators) *ListDomainsResponseBodyDomains {
	s.Accelerators = v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomain(v string) *ListDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetState(v string) *ListDomainsResponseBodyDomains {
	s.State = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) Validate() error {
	if s.Accelerators != nil {
		for _, item := range s.Accelerators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainsResponseBodyDomainsAccelerators struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The name of the GA instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 	- **true**: The GA instance is managed.
	//
	// 	- **false**: The GA instance is not managed.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that you can perform on the managed instance.
	//
	// >  This parameter takes effect only if **ServiceManaged*	- is set to **True**.
	//
	// 	- You can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
}

func (s ListDomainsResponseBodyDomainsAccelerators) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomainsAccelerators) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainsAccelerators) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListDomainsResponseBodyDomainsAccelerators) GetName() *string {
	return s.Name
}

func (s *ListDomainsResponseBodyDomainsAccelerators) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListDomainsResponseBodyDomainsAccelerators) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListDomainsResponseBodyDomainsAccelerators) GetServiceManagedInfos() []*ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListDomainsResponseBodyDomainsAccelerators) SetAcceleratorId(v string) *ListDomainsResponseBodyDomainsAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAccelerators) SetName(v string) *ListDomainsResponseBodyDomainsAccelerators {
	s.Name = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAccelerators) SetServiceId(v string) *ListDomainsResponseBodyDomainsAccelerators {
	s.ServiceId = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAccelerators) SetServiceManaged(v bool) *ListDomainsResponseBodyDomainsAccelerators {
	s.ServiceManaged = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAccelerators) SetServiceManagedInfos(v []*ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) *ListDomainsResponseBodyDomainsAccelerators {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListDomainsResponseBodyDomainsAccelerators) Validate() error {
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

type ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos struct {
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
	// >  This parameter takes effect only if **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and you cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and you can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) SetAction(v string) *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) SetChildType(v string) *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) SetIsManaged(v bool) *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListDomainsResponseBodyDomainsAcceleratorsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}
