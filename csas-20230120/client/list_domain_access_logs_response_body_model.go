// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainAccessLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogs(v []*ListDomainAccessLogsResponseBodyAccessLogs) *ListDomainAccessLogsResponseBody
	GetAccessLogs() []*ListDomainAccessLogsResponseBodyAccessLogs
	SetRequestId(v string) *ListDomainAccessLogsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDomainAccessLogsResponseBody
	GetTotalNum() *int32
}

type ListDomainAccessLogsResponseBody struct {
	// The list of access log records.
	AccessLogs []*ListDomainAccessLogsResponseBodyAccessLogs `json:"AccessLogs,omitempty" xml:"AccessLogs,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 019F68B6-A17E-5ECD-B053-820242E5ADBF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 122
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDomainAccessLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainAccessLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainAccessLogsResponseBody) GetAccessLogs() []*ListDomainAccessLogsResponseBodyAccessLogs {
	return s.AccessLogs
}

func (s *ListDomainAccessLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainAccessLogsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDomainAccessLogsResponseBody) SetAccessLogs(v []*ListDomainAccessLogsResponseBodyAccessLogs) *ListDomainAccessLogsResponseBody {
	s.AccessLogs = v
	return s
}

func (s *ListDomainAccessLogsResponseBody) SetRequestId(v string) *ListDomainAccessLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainAccessLogsResponseBody) SetTotalNum(v int32) *ListDomainAccessLogsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDomainAccessLogsResponseBody) Validate() error {
	if s.AccessLogs != nil {
		for _, item := range s.AccessLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainAccessLogsResponseBodyAccessLogs struct {
	// The action taken upon a rule hit.
	//
	// example:
	//
	// WhiteList
	BlockAction *string `json:"BlockAction,omitempty" xml:"BlockAction,omitempty"`
	// The department.
	//
	// example:
	//
	// IT department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The destination URL accessed.
	//
	// example:
	//
	// https://www.example.com/a
	DestAddress *string `json:"DestAddress,omitempty" xml:"DestAddress,omitempty"`
	// The event time.
	//
	// example:
	//
	// 2026-08-10 14:03:22
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The Layer 4 protocol type.
	//
	// example:
	//
	// tcp
	L4ProtocolType *string `json:"L4ProtocolType,omitempty" xml:"L4ProtocolType,omitempty"`
	// The name of the client process that initiated the access.
	//
	// example:
	//
	// chrome.exe
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 1.2.3.4
	RemoteAddress *string `json:"RemoteAddress,omitempty" xml:"RemoteAddress,omitempty"`
	// The destination domain name.
	//
	// example:
	//
	// www.example.com
	RemoteHost *string `json:"RemoteHost,omitempty" xml:"RemoteHost,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 443
	RemotePort *string `json:"RemotePort,omitempty" xml:"RemotePort,omitempty"`
	// The source address of the client.
	//
	// example:
	//
	// 10.0.0.5
	SrcAddress *string `json:"SrcAddress,omitempty" xml:"SrcAddress,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListDomainAccessLogsResponseBodyAccessLogs) String() string {
	return dara.Prettify(s)
}

func (s ListDomainAccessLogsResponseBodyAccessLogs) GoString() string {
	return s.String()
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetBlockAction() *string {
	return s.BlockAction
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetDepartment() *string {
	return s.Department
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetDestAddress() *string {
	return s.DestAddress
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetEventTime() *string {
	return s.EventTime
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetL4ProtocolType() *string {
	return s.L4ProtocolType
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetRemoteAddress() *string {
	return s.RemoteAddress
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetRemoteHost() *string {
	return s.RemoteHost
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetRemotePort() *string {
	return s.RemotePort
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetSrcAddress() *string {
	return s.SrcAddress
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) GetUsername() *string {
	return s.Username
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetBlockAction(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.BlockAction = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetDepartment(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.Department = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetDestAddress(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.DestAddress = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetEventTime(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.EventTime = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetL4ProtocolType(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.L4ProtocolType = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetProcessName(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.ProcessName = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetRemoteAddress(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.RemoteAddress = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetRemoteHost(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.RemoteHost = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetRemotePort(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.RemotePort = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetSrcAddress(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.SrcAddress = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) SetUsername(v string) *ListDomainAccessLogsResponseBodyAccessLogs {
	s.Username = &v
	return s
}

func (s *ListDomainAccessLogsResponseBodyAccessLogs) Validate() error {
	return dara.Validate(s)
}
