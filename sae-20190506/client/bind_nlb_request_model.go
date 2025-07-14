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
	// The type of the IP addresses. Valid values:
	//
	// 	- Internet: public endpoint.
	//
	// 	- Intranet: private endpoint.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the application to which the NLB instance is bound.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The listener that you want to manage. The value is a string that consists of JSON arrays. Each listener contains the following fields:
	//
	// 	- **port**: the port number of the NLB listener. This field is required. Data type: integer. Valid values: 0 to 65535.
	//
	// 	- **TargetPort**: the port number of the container listener. This field is required. Data type: integer. Valid values: 0 to 65535.
	//
	// 	- **Protocol**: the listener protocol. This field is required. Data type: string. Valid values: TCP, UDP, and TCPSSL.
	//
	// 	- **CertIds**: the IDs of the server certificates. This field is optional. Data type: string. This field is supported only by TCPSSL listeners.
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// The ID of the NLB instance.
	//
	// example:
	//
	// nlb-7z7jjbzz44d82c9***
	NlbId *string `json:"NlbId,omitempty" xml:"NlbId,omitempty"`
	// The mappings between zones and vSwitches. The value is a JSON string. You can specify at most 10 zones. If the region supports two or more zones, specify at least two zones. A ZoneMapping contains the following fields:
	//
	// 	- The ID of the vSwitch in the zone. Each zone can contain only one vSwitch and one subnet. Data type: string.
	//
	// 	- The zone ID of the NLB instance. Data type: string.
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
