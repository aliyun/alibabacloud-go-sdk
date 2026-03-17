// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagLanListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLans(v []*DescribeSagLanListResponseBodyLans) *DescribeSagLanListResponseBody
	GetLans() []*DescribeSagLanListResponseBodyLans
	SetRequestId(v string) *DescribeSagLanListResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagLanListResponseBodyTaskStates) *DescribeSagLanListResponseBody
	GetTaskStates() []*DescribeSagLanListResponseBodyTaskStates
}

type DescribeSagLanListResponseBody struct {
	// The information about the LAN port.
	Lans []*DescribeSagLanListResponseBodyLans `json:"Lans,omitempty" xml:"Lans,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the query task.
	TaskStates []*DescribeSagLanListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagLanListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagLanListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagLanListResponseBody) GetLans() []*DescribeSagLanListResponseBodyLans {
	return s.Lans
}

func (s *DescribeSagLanListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagLanListResponseBody) GetTaskStates() []*DescribeSagLanListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagLanListResponseBody) SetLans(v []*DescribeSagLanListResponseBodyLans) *DescribeSagLanListResponseBody {
	s.Lans = v
	return s
}

func (s *DescribeSagLanListResponseBody) SetRequestId(v string) *DescribeSagLanListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagLanListResponseBody) SetTaskStates(v []*DescribeSagLanListResponseBodyTaskStates) *DescribeSagLanListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagLanListResponseBody) Validate() error {
	if s.Lans != nil {
		for _, item := range s.Lans {
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

type DescribeSagLanListResponseBodyLans struct {
	// The last IP address of the DHCP pool.
	//
	// example:
	//
	// 192.XX.XX.254
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The IP address of the LAN port.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The connection type of the LAN port.
	//
	// **DHCP**: a dynamic IP address. Uses the Dynamic Host Configuration Protocol (DHCP) to dynamically assign an IP address to a connected device.
	//
	// **STATIC**: a static IP address. Specifies a static IP address for the LAN port.
	//
	// example:
	//
	// DHCP
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The time duration that the IP address is retained after it is assigned through DHCP. Unit: minutes.
	//
	// example:
	//
	// 7
	Lease *string `json:"Lease,omitempty" xml:"Lease,omitempty"`
	// The subnet mask of the IP address of the port.
	//
	// example:
	//
	// 255.255.255.0
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name of the port.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The first IP address of the DHCP pool.
	//
	// example:
	//
	// 192.XX.XX.2
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s DescribeSagLanListResponseBodyLans) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagLanListResponseBodyLans) GoString() string {
	return s.String()
}

func (s *DescribeSagLanListResponseBodyLans) GetEndIp() *string {
	return s.EndIp
}

func (s *DescribeSagLanListResponseBodyLans) GetIP() *string {
	return s.IP
}

func (s *DescribeSagLanListResponseBodyLans) GetIPType() *string {
	return s.IPType
}

func (s *DescribeSagLanListResponseBodyLans) GetLease() *string {
	return s.Lease
}

func (s *DescribeSagLanListResponseBodyLans) GetMask() *string {
	return s.Mask
}

func (s *DescribeSagLanListResponseBodyLans) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagLanListResponseBodyLans) GetStartIp() *string {
	return s.StartIp
}

func (s *DescribeSagLanListResponseBodyLans) SetEndIp(v string) *DescribeSagLanListResponseBodyLans {
	s.EndIp = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetIP(v string) *DescribeSagLanListResponseBodyLans {
	s.IP = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetIPType(v string) *DescribeSagLanListResponseBodyLans {
	s.IPType = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetLease(v string) *DescribeSagLanListResponseBodyLans {
	s.Lease = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetMask(v string) *DescribeSagLanListResponseBodyLans {
	s.Mask = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetPortName(v string) *DescribeSagLanListResponseBodyLans {
	s.PortName = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) SetStartIp(v string) *DescribeSagLanListResponseBodyLans {
	s.StartIp = &v
	return s
}

func (s *DescribeSagLanListResponseBodyLans) Validate() error {
	return dara.Validate(s)
}

type DescribeSagLanListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586852928000
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

func (s DescribeSagLanListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagLanListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagLanListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagLanListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagLanListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagLanListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagLanListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagLanListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagLanListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagLanListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagLanListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagLanListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagLanListResponseBodyTaskStates) SetState(v string) *DescribeSagLanListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagLanListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
