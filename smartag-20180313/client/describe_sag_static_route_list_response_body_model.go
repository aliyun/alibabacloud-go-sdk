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
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the static route.
	StaticRoutes []*DescribeSagStaticRouteListResponseBodyStaticRoutes `json:"StaticRoutes,omitempty" xml:"StaticRoutes,omitempty" type:"Repeated"`
	// The state of the query task.
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
	// The name of the port.
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
	// The error code returned. A value of 200 indicates that the query task is successful.
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
	// The state of the query task. Valid values:
	//
	// 	- **Initialized**: The query task is initialized.
	//
	// 	- **Offline**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. After the SAG device is connected to Alibaba Cloud, Alibaba Cloud assigns the query task to the SAG device.
	//
	// 	- **Succeed**: Alibaba Cloud has assigned the query task to the SAG device.
	//
	// 	- **Processing**: Alibaba Cloud is assigning the query task to the SAG device.
	//
	// 	- **VersionNotSupport**: The query task is not supported by the current version of the SAG device.
	//
	// 	- **BuildRequestError**: The query task is not supported by the controller of the SAG device.
	//
	// 	- **HardwareError**: Alibaba Cloud failed to assign the query task to the SAG device because the SAG device is faulty.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. Alibaba Cloud does not assign the query task to the SAG device even after the SAG device is connected to Alibaba Cloud.
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
