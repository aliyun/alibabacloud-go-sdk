// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindNlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *BindNlbRequest
	GetAddressType() *string
	SetAppId(v string) *BindNlbRequest
	GetAppId() *string
	SetListeners(v string) *BindNlbRequest
	GetListeners() *string
	SetNlbId(v string) *BindNlbRequest
	GetNlbId() *string
	SetZoneMappings(v string) *BindNlbRequest
	GetZoneMappings() *string
}

type BindNlbRequest struct {
	// The address type of the NLB instance.
	//
	// - `Internet`: a public IP address.
	//
	// - `Intranet`: a private IP address.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the target application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The listeners, specified as a JSON-formatted string. Each listener object contains the following fields:
	//
	// - **Port**: Integer. Required. The listener port. Valid values: 0 to 65535.
	//
	// - **TargetPort**: Integer. Required. The port on the application instance that receives traffic. Valid values: 0 to 65535.
	//
	// - **Protocol**: String. Required. The listener protocol. Valid values: `TCP`, `UDP`, and `TCPSSL`.
	//
	// - **CertIds**: String. Optional. The server certificate IDs. This parameter is required only for `TCPSSL` listeners.
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// The ID of the NLB instance.
	//
	// example:
	//
	// nlb-7z7jjbzz44d82c9***
	NlbId *string `json:"NlbId,omitempty" xml:"NlbId,omitempty"`
	// The mappings between zones and vSwitches, specified as a JSON-formatted string. You can add up to 10 zones. If the current region supports two or more zones, you must specify at least two zones. Each `ZoneMapping` object contains the following fields:
	//
	// - **VSwitchId**: String. The ID of the vSwitch in the specified zone. Each zone can have only one vSwitch and one subnet.
	//
	// - ZoneId, String, the zone ID of the Network Load Balancer instance.
	//
	// example:
	//
	// vsw-sersdf****
	//
	// cn-hangzhou-a
	ZoneMappings *string `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty"`
}

func (s BindNlbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindNlbRequest) GoString() string {
	return s.String()
}

func (s *BindNlbRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *BindNlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindNlbRequest) GetListeners() *string {
	return s.Listeners
}

func (s *BindNlbRequest) GetNlbId() *string {
	return s.NlbId
}

func (s *BindNlbRequest) GetZoneMappings() *string {
	return s.ZoneMappings
}

func (s *BindNlbRequest) SetAddressType(v string) *BindNlbRequest {
	s.AddressType = &v
	return s
}

func (s *BindNlbRequest) SetAppId(v string) *BindNlbRequest {
	s.AppId = &v
	return s
}

func (s *BindNlbRequest) SetListeners(v string) *BindNlbRequest {
	s.Listeners = &v
	return s
}

func (s *BindNlbRequest) SetNlbId(v string) *BindNlbRequest {
	s.NlbId = &v
	return s
}

func (s *BindNlbRequest) SetZoneMappings(v string) *BindNlbRequest {
	s.ZoneMappings = &v
	return s
}

func (s *BindNlbRequest) Validate() error {
	return dara.Validate(s)
}
