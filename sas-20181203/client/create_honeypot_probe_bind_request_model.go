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
	// The list of port bindings.
	BindPortList []*CreateHoneypotProbeBindRequestBindPortList `json:"BindPortList,omitempty" xml:"BindPortList,omitempty" type:"Repeated"`
	// The honeypot ID.
	//
	// > Call the [ListHoneypot](~~ListHoneypot~~) operation to obtain this value.
	//
	// This parameter is required. If this parameter is not specified, the API returns InternalError (400).
	//
	// example:
	//
	// dba7d44775be8e0e5888ee3b1a62554a93d2512247cabc38ddeac17a3b3f****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The probe ID.
	//
	// >Call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to obtain this parameter.
	//
	// This parameter is required. If this parameter is not specified, the API returns InvalidParam (400).
	//
	// example:
	//
	// 36bad711-d1ac-4419-ac68-c1aa280f****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The list of listener IP addresses.
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
	if s.BindPortList != nil {
		for _, item := range s.BindPortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHoneypotProbeBindRequestBindPortList struct {
	// Specifies whether to bind the port. Valid values:
	//
	// - **true**: The port is bound.
	//
	// - **false**: The port is not bound.
	//
	// example:
	//
	// false
	BindPort *bool `json:"BindPort,omitempty" xml:"BindPort,omitempty"`
	// The end port of the probe listener.
	//
	// example:
	//
	// 80
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// Specifies whether the port is fixed. Valid values:
	//
	// - **true**: The port is fixed.
	//
	// - **false**: The port is not fixed.
	//
	// example:
	//
	// false
	Fixed *bool `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
	// The protocol type. Valid values:
	//
	// - **tcp**: TCP protocol.
	//
	// - **udp**: UDP protocol.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The start port of the probe listener.
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
