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
	// The address type. Valid values:
	//
	// - Internet: public network address.
	//
	// - Intranet: private network address.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The JSON array string of the listeners to be configured. The structure of each listener is as follows:
	//
	// - **Port**: Integer type, required. The load balancing listener port. Valid values: 0 to 65535.
	//
	// - **TargetPort**: Integer type, required. The listening port of the container in the instance. Valid values: 0 to 65535.
	//
	// - **Protocol**: String type, required. The listening protocol. Valid values: TCP, UDP, and TCPSSL.
	//
	// - **CertIds**: String type, optional. The list of server certificates. This parameter takes effect only for TCPSSL listeners.
	//
	// example:
	//
	// [{"Port":80,"TargetPort":8080,"Protocol":"TCP"}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// The ID of the NLB instance.
	//
	// example:
	//
	// nlb-7z7jjbzz44d82c9***
	NlbId *string `json:"NlbId,omitempty" xml:"NlbId,omitempty"`
	// The JSON string of the zone and vSwitch mapping list. You can add up to 10 zones. If the current region supports two or more zones, you must add at least two zones. The structure of a ZoneMapping is as follows:
	//
	// - VSwitchId: String type. The vSwitch corresponding to the zone. Each zone can use only one vSwitch and one subnet.
	//
	// - ZoneId: String type. The zone ID of the Network Load Balancer (NLB) instance.
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
