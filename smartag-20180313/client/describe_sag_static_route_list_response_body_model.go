// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagStaticRouteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagStaticRouteListResponseBody
	GetRequestId() *string
	SetStaticRoutes(v []*DescribeSagStaticRouteListResponseBodyStaticRoutes) *DescribeSagStaticRouteListResponseBody
	GetStaticRoutes() []*DescribeSagStaticRouteListResponseBodyStaticRoutes
	SetTaskStates(v []*DescribeSagStaticRouteListResponseBodyTaskStates) *DescribeSagStaticRouteListResponseBody
	GetTaskStates() []*DescribeSagStaticRouteListResponseBodyTaskStates
}

type DescribeSagStaticRouteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of static routes.
	StaticRoutes []*DescribeSagStaticRouteListResponseBodyStaticRoutes `json:"StaticRoutes,omitempty" xml:"StaticRoutes,omitempty" type:"Repeated"`
	// The status of the query tasks.
	TaskStates []*DescribeSagStaticRouteListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagStaticRouteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagStaticRouteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagStaticRouteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagStaticRouteListResponseBody) GetStaticRoutes() []*DescribeSagStaticRouteListResponseBodyStaticRoutes {
	return s.StaticRoutes
}

func (s *DescribeSagStaticRouteListResponseBody) GetTaskStates() []*DescribeSagStaticRouteListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagStaticRouteListResponseBody) SetRequestId(v string) *DescribeSagStaticRouteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBody) SetStaticRoutes(v []*DescribeSagStaticRouteListResponseBodyStaticRoutes) *DescribeSagStaticRouteListResponseBody {
	s.StaticRoutes = v
	return s
}

func (s *DescribeSagStaticRouteListResponseBody) SetTaskStates(v []*DescribeSagStaticRouteListResponseBodyTaskStates) *DescribeSagStaticRouteListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagStaticRouteListResponseBody) Validate() error {
	if s.StaticRoutes != nil {
		for _, item := range s.StaticRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskStates != nil {
		for _, item := range s.TaskStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagStaticRouteListResponseBodyStaticRoutes struct {
	// The destination CIDR block.
	//
	// example:
	//
	// 192.XX.XX.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The next hop.
	//
	// example:
	//
	// 192.XX.XX.1
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The port name.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The VLAN ID.
	//
	// example:
	//
	// 1
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DescribeSagStaticRouteListResponseBodyStaticRoutes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagStaticRouteListResponseBodyStaticRoutes) GoString() string {
	return s.String()
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) GetNextHop() *string {
	return s.NextHop
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) GetVlan() *string {
	return s.Vlan
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) SetDestinationCidr(v string) *DescribeSagStaticRouteListResponseBodyStaticRoutes {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) SetNextHop(v string) *DescribeSagStaticRouteListResponseBodyStaticRoutes {
	s.NextHop = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) SetPortName(v string) *DescribeSagStaticRouteListResponseBodyStaticRoutes {
	s.PortName = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) SetVlan(v string) *DescribeSagStaticRouteListResponseBodyStaticRoutes {
	s.Vlan = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyStaticRoutes) Validate() error {
	return dara.Validate(s)
}

type DescribeSagStaticRouteListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586857309000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device goes online.
	//
	// - **Succeed**: The query task is delivered successfully.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this feature.
	//
	// - **BuildRequestError**: The management and control system does not support the request.
	//
	// - **HardwareError**: The query task failed to be delivered due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered even after the device goes online.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagStaticRouteListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagStaticRouteListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagStaticRouteListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagStaticRouteListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagStaticRouteListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) SetState(v string) *DescribeSagStaticRouteListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagStaticRouteListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
