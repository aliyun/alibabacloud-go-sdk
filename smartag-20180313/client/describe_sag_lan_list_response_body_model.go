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
	// The list of LAN port information.
	Lans []*DescribeSagLanListResponseBodyLans `json:"Lans,omitempty" xml:"Lans,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status information of the query task.
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
	// The end address of the DHCP address pool.
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
	// The connection type of the LAN port. Valid values:
	//
	// **DHCP**: Dynamic IP. Uses the DHCP protocol to dynamically assign IP addresses to connected devices.
	//
	// **STATIC**: Static IP. Statically specifies the IP address of the LAN port.
	//
	// example:
	//
	// DHCP
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The DHCP lease time. Unit: minutes.
	//
	// example:
	//
	// 7
	Lease *string `json:"Lease,omitempty" xml:"Lease,omitempty"`
	// The subnet mask.
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
	// The start IP address of the DHCP address pool.
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
	// The error code. 200 indicates that the query task succeeded.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. Successful indicates that the query task succeeded.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device comes online.
	//
	// - **Succeed**: The query task was delivered successfully.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this operation.
	//
	// - **BuildRequestError**: The request is not supported by the management system.
	//
	// - **HardwareError**: The query task failed to be delivered due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered even after the device comes online.
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
