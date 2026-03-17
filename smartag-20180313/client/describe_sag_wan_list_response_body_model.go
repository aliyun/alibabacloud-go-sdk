// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagWanListResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagWanListResponseBodyTaskStates) *DescribeSagWanListResponseBody
	GetTaskStates() []*DescribeSagWanListResponseBodyTaskStates
	SetWans(v []*DescribeSagWanListResponseBodyWans) *DescribeSagWanListResponseBody
	GetWans() []*DescribeSagWanListResponseBodyWans
}

type DescribeSagWanListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of and information about the query task.
	TaskStates []*DescribeSagWanListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
	// The settings of the WAN port.
	Wans []*DescribeSagWanListResponseBodyWans `json:"Wans,omitempty" xml:"Wans,omitempty" type:"Repeated"`
}

func (s DescribeSagWanListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagWanListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagWanListResponseBody) GetTaskStates() []*DescribeSagWanListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagWanListResponseBody) GetWans() []*DescribeSagWanListResponseBodyWans {
	return s.Wans
}

func (s *DescribeSagWanListResponseBody) SetRequestId(v string) *DescribeSagWanListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagWanListResponseBody) SetTaskStates(v []*DescribeSagWanListResponseBodyTaskStates) *DescribeSagWanListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagWanListResponseBody) SetWans(v []*DescribeSagWanListResponseBodyWans) *DescribeSagWanListResponseBody {
	s.Wans = v
	return s
}

func (s *DescribeSagWanListResponseBody) Validate() error {
	if s.TaskStates != nil {
		for _, item := range s.TaskStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Wans != nil {
		for _, item := range s.Wans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagWanListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586834861000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned to the query task. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned to the query task. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the query task. Valid values:
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

func (s DescribeSagWanListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagWanListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagWanListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagWanListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagWanListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagWanListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagWanListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagWanListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagWanListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagWanListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagWanListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagWanListResponseBodyTaskStates) SetState(v string) *DescribeSagWanListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagWanListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}

type DescribeSagWanListResponseBodyWans struct {
	// The bandwidth cap of the WAN port. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	BandWidth *int32 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The IP address of the gateway.
	//
	// example:
	//
	// 192.XX.XX.1
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the WAN port.
	//
	// example:
	//
	// 172.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The connection type of the WAN port. Valid values:
	//
	// 	- **DHCP**: The WAN port connects to the Internet through an IP address that is dynamically assigned by the Dynamic Host Configuration Protocol (DHCP) server.
	//
	// 	- **STATIC**: The WAN port connects to the Internet through a static IP address.
	//
	// 	- **PPPOE**: The WAN port connects to the Internet through dial-up connections.
	//
	// example:
	//
	// STATIC
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The Internet service provider (ISP) used by the WAN port.
	//
	// 	- **CT**: China Telecom
	//
	// 	- **CM**: China Mobile
	//
	// 	- **CU**: China Unicom
	//
	// 	- **Other**: Other ISPs
	//
	// example:
	//
	// CT
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The subnet mask of the WAN port IP address.
	//
	// example:
	//
	// 255.255.255.240
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The number of the WAN port.
	//
	// example:
	//
	// 1
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The priority of the WAN port.
	//
	// Valid values: **-1*	- and **1 to 50**. A smaller number represents a higher priority.
	//
	// >  A value of **-1*	- indicates that the WAN port is not used to forward network traffic.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of data transfer on the WAN port. Valid values:
	//
	// 	- **active**: The WAN port is the active port for data transfer.
	//
	// 	- **standby**: The WAN port is a standby port. If the active port is down, data transfer is switched to the WAN port.
	//
	// example:
	//
	// active
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The username of the PPPoE account.
	//
	// example:
	//
	// Usernamexx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The weight of the WAN port.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeSagWanListResponseBodyWans) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanListResponseBodyWans) GoString() string {
	return s.String()
}

func (s *DescribeSagWanListResponseBodyWans) GetBandWidth() *int32 {
	return s.BandWidth
}

func (s *DescribeSagWanListResponseBodyWans) GetGateway() *string {
	return s.Gateway
}

func (s *DescribeSagWanListResponseBodyWans) GetIP() *string {
	return s.IP
}

func (s *DescribeSagWanListResponseBodyWans) GetIPType() *string {
	return s.IPType
}

func (s *DescribeSagWanListResponseBodyWans) GetISP() *string {
	return s.ISP
}

func (s *DescribeSagWanListResponseBodyWans) GetMask() *string {
	return s.Mask
}

func (s *DescribeSagWanListResponseBodyWans) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagWanListResponseBodyWans) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeSagWanListResponseBodyWans) GetTrafficState() *string {
	return s.TrafficState
}

func (s *DescribeSagWanListResponseBodyWans) GetUsername() *string {
	return s.Username
}

func (s *DescribeSagWanListResponseBodyWans) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeSagWanListResponseBodyWans) SetBandWidth(v int32) *DescribeSagWanListResponseBodyWans {
	s.BandWidth = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetGateway(v string) *DescribeSagWanListResponseBodyWans {
	s.Gateway = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetIP(v string) *DescribeSagWanListResponseBodyWans {
	s.IP = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetIPType(v string) *DescribeSagWanListResponseBodyWans {
	s.IPType = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetISP(v string) *DescribeSagWanListResponseBodyWans {
	s.ISP = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetMask(v string) *DescribeSagWanListResponseBodyWans {
	s.Mask = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetPortName(v string) *DescribeSagWanListResponseBodyWans {
	s.PortName = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetPriority(v int32) *DescribeSagWanListResponseBodyWans {
	s.Priority = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetTrafficState(v string) *DescribeSagWanListResponseBodyWans {
	s.TrafficState = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetUsername(v string) *DescribeSagWanListResponseBodyWans {
	s.Username = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) SetWeight(v int32) *DescribeSagWanListResponseBodyWans {
	s.Weight = &v
	return s
}

func (s *DescribeSagWanListResponseBodyWans) Validate() error {
	return dara.Validate(s)
}
