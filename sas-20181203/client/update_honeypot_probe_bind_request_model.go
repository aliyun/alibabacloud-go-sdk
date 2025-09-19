// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindId(v string) *UpdateHoneypotProbeBindRequest
	GetBindId() *string
	SetBindPortList(v []*UpdateHoneypotProbeBindRequestBindPortList) *UpdateHoneypotProbeBindRequest
	GetBindPortList() []*UpdateHoneypotProbeBindRequestBindPortList
	SetBindType(v string) *UpdateHoneypotProbeBindRequest
	GetBindType() *string
	SetCurrentPage(v int32) *UpdateHoneypotProbeBindRequest
	GetCurrentPage() *int32
	SetHoneypotId(v string) *UpdateHoneypotProbeBindRequest
	GetHoneypotId() *string
	SetId(v int64) *UpdateHoneypotProbeBindRequest
	GetId() *int64
	SetLang(v string) *UpdateHoneypotProbeBindRequest
	GetLang() *string
	SetPageSize(v int32) *UpdateHoneypotProbeBindRequest
	GetPageSize() *int32
	SetPorts(v string) *UpdateHoneypotProbeBindRequest
	GetPorts() *string
	SetProbeId(v string) *UpdateHoneypotProbeBindRequest
	GetProbeId() *string
	SetServiceIpList(v []*string) *UpdateHoneypotProbeBindRequest
	GetServiceIpList() []*string
	SetSetStatus(v int32) *UpdateHoneypotProbeBindRequest
	GetSetStatus() *int32
}

type UpdateHoneypotProbeBindRequest struct {
	// The unique ID of the honeypot to which the probe is bound.
	//
	// example:
	//
	// f52e8624-e43c-473c-8312-e0fed384****
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The ports that are bound to the probe.
	BindPortList []*UpdateHoneypotProbeBindRequestBindPortList `json:"BindPortList,omitempty" xml:"BindPortList,omitempty" type:"Repeated"`
	// The operation that the probe performs. Valid values:
	//
	// 	- **forward_honey**: forward traffic to a honeypot
	//
	// 	- **scan_port**: monitor and scan
	//
	// example:
	//
	// forward_honey
	BindType *string `json:"BindType,omitempty" xml:"BindType,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The honeypot ID.
	//
	// >  You can call the [ListHoneypot](~~ListHoneypot~~) operation to obtain the IDs of honeypots.
	//
	// example:
	//
	// dba7d44775be8e0e5888ee3b1a62554a93d2512247cabc38ddeac17a3b3f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The port ID of the probe service.
	//
	// example:
	//
	// 1906
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ports that are monitored.
	//
	// example:
	//
	// {\\"tcp\\":\\"1-65535\\",\\"udp\\":\\"1-65535\\"}
	Ports *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	// The probe ID.
	//
	// >  You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// example:
	//
	// 36bad711-d1ac-4419-ac68-c1aa280f****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The IP addresses that are monitored.
	ServiceIpList []*string `json:"ServiceIpList,omitempty" xml:"ServiceIpList,omitempty" type:"Repeated"`
	// The status of the port.
	//
	// example:
	//
	// 0
	SetStatus *int32 `json:"SetStatus,omitempty" xml:"SetStatus,omitempty"`
}

func (s UpdateHoneypotProbeBindRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeBindRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeBindRequest) GetBindId() *string {
	return s.BindId
}

func (s *UpdateHoneypotProbeBindRequest) GetBindPortList() []*UpdateHoneypotProbeBindRequestBindPortList {
	return s.BindPortList
}

func (s *UpdateHoneypotProbeBindRequest) GetBindType() *string {
	return s.BindType
}

func (s *UpdateHoneypotProbeBindRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *UpdateHoneypotProbeBindRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *UpdateHoneypotProbeBindRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateHoneypotProbeBindRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateHoneypotProbeBindRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UpdateHoneypotProbeBindRequest) GetPorts() *string {
	return s.Ports
}

func (s *UpdateHoneypotProbeBindRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *UpdateHoneypotProbeBindRequest) GetServiceIpList() []*string {
	return s.ServiceIpList
}

func (s *UpdateHoneypotProbeBindRequest) GetSetStatus() *int32 {
	return s.SetStatus
}

func (s *UpdateHoneypotProbeBindRequest) SetBindId(v string) *UpdateHoneypotProbeBindRequest {
	s.BindId = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetBindPortList(v []*UpdateHoneypotProbeBindRequestBindPortList) *UpdateHoneypotProbeBindRequest {
	s.BindPortList = v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetBindType(v string) *UpdateHoneypotProbeBindRequest {
	s.BindType = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetCurrentPage(v int32) *UpdateHoneypotProbeBindRequest {
	s.CurrentPage = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetHoneypotId(v string) *UpdateHoneypotProbeBindRequest {
	s.HoneypotId = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetId(v int64) *UpdateHoneypotProbeBindRequest {
	s.Id = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetLang(v string) *UpdateHoneypotProbeBindRequest {
	s.Lang = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetPageSize(v int32) *UpdateHoneypotProbeBindRequest {
	s.PageSize = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetPorts(v string) *UpdateHoneypotProbeBindRequest {
	s.Ports = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetProbeId(v string) *UpdateHoneypotProbeBindRequest {
	s.ProbeId = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetServiceIpList(v []*string) *UpdateHoneypotProbeBindRequest {
	s.ServiceIpList = v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) SetSetStatus(v int32) *UpdateHoneypotProbeBindRequest {
	s.SetStatus = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateHoneypotProbeBindRequestBindPortList struct {
	// Specifies whether to bind a port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	BindPort *bool `json:"BindPort,omitempty" xml:"BindPort,omitempty"`
	// The end port on which the probe monitors.
	//
	// example:
	//
	// 81
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// Specifies whether the port is fixed. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	Fixed *bool `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
	// The UUID of the port.
	//
	// example:
	//
	// 3183
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The start port on which the probe monitors.
	//
	// example:
	//
	// 81
	StartPort *int32 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s UpdateHoneypotProbeBindRequestBindPortList) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeBindRequestBindPortList) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetBindPort() *bool {
	return s.BindPort
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetEndPort() *int32 {
	return s.EndPort
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetFixed() *bool {
	return s.Fixed
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetId() *int64 {
	return s.Id
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetProto() *string {
	return s.Proto
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetStartPort() *int32 {
	return s.StartPort
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetBindPort(v bool) *UpdateHoneypotProbeBindRequestBindPortList {
	s.BindPort = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetEndPort(v int32) *UpdateHoneypotProbeBindRequestBindPortList {
	s.EndPort = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetFixed(v bool) *UpdateHoneypotProbeBindRequestBindPortList {
	s.Fixed = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetId(v int64) *UpdateHoneypotProbeBindRequestBindPortList {
	s.Id = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetProto(v string) *UpdateHoneypotProbeBindRequestBindPortList {
	s.Proto = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetStartPort(v int32) *UpdateHoneypotProbeBindRequestBindPortList {
	s.StartPort = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) SetTargetPort(v int32) *UpdateHoneypotProbeBindRequestBindPortList {
	s.TargetPort = &v
	return s
}

func (s *UpdateHoneypotProbeBindRequestBindPortList) Validate() error {
	return dara.Validate(s)
}
