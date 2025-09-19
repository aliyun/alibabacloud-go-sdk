// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindPortList(v []*CreateHoneypotProbeBindRequestBindPortList) *CreateHoneypotProbeBindRequest
	GetBindPortList() []*CreateHoneypotProbeBindRequestBindPortList
	SetHoneypotId(v string) *CreateHoneypotProbeBindRequest
	GetHoneypotId() *string
	SetLang(v string) *CreateHoneypotProbeBindRequest
	GetLang() *string
	SetProbeId(v string) *CreateHoneypotProbeBindRequest
	GetProbeId() *string
	SetServiceIpList(v []*string) *CreateHoneypotProbeBindRequest
	GetServiceIpList() []*string
}

type CreateHoneypotProbeBindRequest struct {
	// The ports that are bound to the probe.
	BindPortList []*CreateHoneypotProbeBindRequestBindPortList `json:"BindPortList,omitempty" xml:"BindPortList,omitempty" type:"Repeated"`
	// The honeypot ID.
	//
	// >  You can call the [ListHoneypot](~~ListHoneypot~~) operation to query the IDs of honeypots.
	//
	// example:
	//
	// dba7d44775be8e0e5888ee3b1a62554a93d2512247cabc38ddeac17a3b3f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
}

func (s CreateHoneypotProbeBindRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeBindRequest) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeBindRequest) GetBindPortList() []*CreateHoneypotProbeBindRequestBindPortList {
	return s.BindPortList
}

func (s *CreateHoneypotProbeBindRequest) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *CreateHoneypotProbeBindRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateHoneypotProbeBindRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *CreateHoneypotProbeBindRequest) GetServiceIpList() []*string {
	return s.ServiceIpList
}

func (s *CreateHoneypotProbeBindRequest) SetBindPortList(v []*CreateHoneypotProbeBindRequestBindPortList) *CreateHoneypotProbeBindRequest {
	s.BindPortList = v
	return s
}

func (s *CreateHoneypotProbeBindRequest) SetHoneypotId(v string) *CreateHoneypotProbeBindRequest {
	s.HoneypotId = &v
	return s
}

func (s *CreateHoneypotProbeBindRequest) SetLang(v string) *CreateHoneypotProbeBindRequest {
	s.Lang = &v
	return s
}

func (s *CreateHoneypotProbeBindRequest) SetProbeId(v string) *CreateHoneypotProbeBindRequest {
	s.ProbeId = &v
	return s
}

func (s *CreateHoneypotProbeBindRequest) SetServiceIpList(v []*string) *CreateHoneypotProbeBindRequest {
	s.ServiceIpList = v
	return s
}

func (s *CreateHoneypotProbeBindRequest) Validate() error {
	return dara.Validate(s)
}

type CreateHoneypotProbeBindRequestBindPortList struct {
	// Specifies whether to bind the port. Valid values:
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
	// 80
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// Specifies whether the port is a fixed port. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fixed *bool `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
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
	// 80
	StartPort *int32 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 8080
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s CreateHoneypotProbeBindRequestBindPortList) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeBindRequestBindPortList) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetBindPort() *bool {
	return s.BindPort
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetEndPort() *int32 {
	return s.EndPort
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetFixed() *bool {
	return s.Fixed
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetProto() *string {
	return s.Proto
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetStartPort() *int32 {
	return s.StartPort
}

func (s *CreateHoneypotProbeBindRequestBindPortList) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetBindPort(v bool) *CreateHoneypotProbeBindRequestBindPortList {
	s.BindPort = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetEndPort(v int32) *CreateHoneypotProbeBindRequestBindPortList {
	s.EndPort = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetFixed(v bool) *CreateHoneypotProbeBindRequestBindPortList {
	s.Fixed = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetProto(v string) *CreateHoneypotProbeBindRequestBindPortList {
	s.Proto = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetStartPort(v int32) *CreateHoneypotProbeBindRequestBindPortList {
	s.StartPort = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) SetTargetPort(v int32) *CreateHoneypotProbeBindRequestBindPortList {
	s.TargetPort = &v
	return s
}

func (s *CreateHoneypotProbeBindRequestBindPortList) Validate() error {
	return dara.Validate(s)
}
