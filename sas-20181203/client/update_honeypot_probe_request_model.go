// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArp(v bool) *UpdateHoneypotProbeRequest
	GetArp() *bool
	SetDisplayName(v string) *UpdateHoneypotProbeRequest
	GetDisplayName() *string
	SetLang(v string) *UpdateHoneypotProbeRequest
	GetLang() *string
	SetPing(v bool) *UpdateHoneypotProbeRequest
	GetPing() *bool
	SetProbeId(v string) *UpdateHoneypotProbeRequest
	GetProbeId() *string
	SetServiceIpList(v []*string) *UpdateHoneypotProbeRequest
	GetServiceIpList() []*string
}

type UpdateHoneypotProbeRequest struct {
	// Specifies whether address resolution protocol (ARP) is enabled for the check type.
	//
	// example:
	//
	// false
	Arp *bool `json:"Arp,omitempty" xml:"Arp,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// svwsx-vpc-4430
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether ping is enabled for the check type.
	//
	// example:
	//
	// true
	Ping *bool `json:"Ping,omitempty" xml:"Ping,omitempty"`
	// The ID of the probe.
	//
	// > You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// This parameter is required.
	//
	// example:
	//
	// bbe7e382-956f-473e-beed-bc73a258****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The IP addresses that are monitored.
	ServiceIpList []*string `json:"ServiceIpList,omitempty" xml:"ServiceIpList,omitempty" type:"Repeated"`
}

func (s UpdateHoneypotProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeRequest) GetArp() *bool {
	return s.Arp
}

func (s *UpdateHoneypotProbeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateHoneypotProbeRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateHoneypotProbeRequest) GetPing() *bool {
	return s.Ping
}

func (s *UpdateHoneypotProbeRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *UpdateHoneypotProbeRequest) GetServiceIpList() []*string {
	return s.ServiceIpList
}

func (s *UpdateHoneypotProbeRequest) SetArp(v bool) *UpdateHoneypotProbeRequest {
	s.Arp = &v
	return s
}

func (s *UpdateHoneypotProbeRequest) SetDisplayName(v string) *UpdateHoneypotProbeRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateHoneypotProbeRequest) SetLang(v string) *UpdateHoneypotProbeRequest {
	s.Lang = &v
	return s
}

func (s *UpdateHoneypotProbeRequest) SetPing(v bool) *UpdateHoneypotProbeRequest {
	s.Ping = &v
	return s
}

func (s *UpdateHoneypotProbeRequest) SetProbeId(v string) *UpdateHoneypotProbeRequest {
	s.ProbeId = &v
	return s
}

func (s *UpdateHoneypotProbeRequest) SetServiceIpList(v []*string) *UpdateHoneypotProbeRequest {
	s.ServiceIpList = v
	return s
}

func (s *UpdateHoneypotProbeRequest) Validate() error {
	return dara.Validate(s)
}
