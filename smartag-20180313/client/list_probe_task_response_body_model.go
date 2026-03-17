// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProbeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProbeTaskResponseBody
	GetCode() *string
	SetData(v []*ListProbeTaskResponseBodyData) *ListProbeTaskResponseBody
	GetData() []*ListProbeTaskResponseBodyData
	SetMessage(v string) *ListProbeTaskResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListProbeTaskResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProbeTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListProbeTaskResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProbeTaskResponseBody
	GetTotalCount() *int32
}

type ListProbeTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the probe task.
	//
	// example:
	//
	// probe-xxx
	Data []*ListProbeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProbeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProbeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListProbeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProbeTaskResponseBody) GetData() []*ListProbeTaskResponseBodyData {
	return s.Data
}

func (s *ListProbeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProbeTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProbeTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProbeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProbeTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProbeTaskResponseBody) SetCode(v string) *ListProbeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListProbeTaskResponseBody) SetData(v []*ListProbeTaskResponseBodyData) *ListProbeTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListProbeTaskResponseBody) SetMessage(v string) *ListProbeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListProbeTaskResponseBody) SetPageNumber(v int32) *ListProbeTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProbeTaskResponseBody) SetPageSize(v int32) *ListProbeTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProbeTaskResponseBody) SetRequestId(v string) *ListProbeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProbeTaskResponseBody) SetTotalCount(v int32) *ListProbeTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProbeTaskResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProbeTaskResponseBodyData struct {
	// The domain name that is probed by the task.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the probe task is enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the probe task was created.
	//
	// example:
	//
	// 2022-11-23 14:09
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the probe task was modified.
	//
	// example:
	//
	// 2022-11-23 14:09
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The number of probe packets transmitted by the probe task per minute.
	//
	// example:
	//
	// 10
	PacketNumber *int32 `json:"PacketNumber,omitempty" xml:"PacketNumber,omitempty"`
	// The port that is probed by the task.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the probe task.
	//
	// example:
	//
	// probe-****
	ProbeTaskId *string `json:"ProbeTaskId,omitempty" xml:"ProbeTaskId,omitempty"`
	// The source address of the probe task.
	//
	// example:
	//
	// 192.168.1.1
	ProbeTaskSourceAddress *string `json:"ProbeTaskSourceAddress,omitempty" xml:"ProbeTaskSourceAddress,omitempty"`
	// The protocol of the probe task. Valid values:
	//
	// 	- **ICMP**
	//
	// 	- **TCP**
	//
	// 	- **HTTP**
	//
	// > Tasks that probe private networks support only ICMP and TCP.
	//
	// example:
	//
	// ICMP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-asdfz6ac74oj5v****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sag****
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The name of the probe task.
	//
	// example:
	//
	// test-ping
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the probe task. Valid values:
	//
	// 	- **Internet**: probes a public network.
	//
	// 	- **Intranet**: probes a private network.
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProbeTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProbeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProbeTaskResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *ListProbeTaskResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *ListProbeTaskResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListProbeTaskResponseBodyData) GetGmtModify() *string {
	return s.GmtModify
}

func (s *ListProbeTaskResponseBodyData) GetPacketNumber() *int32 {
	return s.PacketNumber
}

func (s *ListProbeTaskResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListProbeTaskResponseBodyData) GetProbeTaskId() *string {
	return s.ProbeTaskId
}

func (s *ListProbeTaskResponseBodyData) GetProbeTaskSourceAddress() *string {
	return s.ProbeTaskSourceAddress
}

func (s *ListProbeTaskResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *ListProbeTaskResponseBodyData) GetSagId() *string {
	return s.SagId
}

func (s *ListProbeTaskResponseBodyData) GetSn() *string {
	return s.Sn
}

func (s *ListProbeTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListProbeTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListProbeTaskResponseBodyData) SetDomain(v string) *ListProbeTaskResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetEnable(v bool) *ListProbeTaskResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetGmtCreate(v string) *ListProbeTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetGmtModify(v string) *ListProbeTaskResponseBodyData {
	s.GmtModify = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetPacketNumber(v int32) *ListProbeTaskResponseBodyData {
	s.PacketNumber = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetPort(v int32) *ListProbeTaskResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetProbeTaskId(v string) *ListProbeTaskResponseBodyData {
	s.ProbeTaskId = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetProbeTaskSourceAddress(v string) *ListProbeTaskResponseBodyData {
	s.ProbeTaskSourceAddress = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetProtocol(v string) *ListProbeTaskResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetSagId(v string) *ListProbeTaskResponseBodyData {
	s.SagId = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetSn(v string) *ListProbeTaskResponseBodyData {
	s.Sn = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetTaskName(v string) *ListProbeTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) SetType(v string) *ListProbeTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListProbeTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
